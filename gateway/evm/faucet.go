package evm

import (
	"crypto/ecdsa"
	"net/http"
	"strconv"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"

	// "github.com/powerman/rpc-codec/jsonrpc2"
	jsonrpc2 "github.com/KeisukeYamashita/go-jsonrpc"
	"github.com/ethereum/go-ethereum/crypto"
)

// RegisterEVMFaucetRoutes registers query status of EVM chains.
func RegisterEVMFaucetRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryEVMFaucet, RegisterEVMFaucetRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryEVMFaucet, "This is an API to faucet.", true)
}

func queryEVMFaucetFromNode(chainConfig *config.EVMConfig, nodeInfo config.EVMNodeConfig, address string, token string) (interface{}, interface{}, int) {
	client := jsonrpc2.NewRPCClient(nodeInfo.RPC + "/" + nodeInfo.RPCToken)
	if nodeInfo.RPCSecret != "" {
		client.SetBasicAuth(nodeInfo.RPCToken, nodeInfo.RPCSecret)
	}

	privateKey, _ := crypto.HexToECDSA(chainConfig.Faucet.PrivateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	faucetAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	faucetAccountBalance := uint64(0)
	Y := uint64(0) // userBalance

	X, ok := chainConfig.Faucet.FaucetAmounts[token]
	if !ok {
		return common.ServeError(0, "", "unsupported token", http.StatusBadRequest)
	}

	M, ok := chainConfig.Faucet.FaucetMinimumAmounts[token]
	if !ok {
		return common.ServeError(0, "", "unsupported token", http.StatusBadRequest)
	}

	if token == "0x00000000000000000000000000000000000000000000000000" {
		// faucet account eth balance
		data, err := client.Call("eth_getBalance", faucetAddress.String(), "latest")
		if err != nil {
			return common.ServeError(0, "failed to get eth balance of faucet account", err.Error(), http.StatusInternalServerError)
		}
		ethBalanceString, err := data.GetString()
		if err != nil {
			return common.ServeError(0, "failed to get eth balance of faucet account", err.Error(), http.StatusInternalServerError)
		}
		faucetAccountBalance, _ = strconv.ParseUint((ethBalanceString)[2:], 16, 64)

		// user eth balance
		data, err = client.Call("eth_getBalance", address, "latest")
		if err != nil {
			return common.ServeError(0, "failed to get eth balance of user ", err.Error(), http.StatusInternalServerError)
		}
		ethBalanceString, err = data.GetString()
		if err != nil {
			return common.ServeError(0, "failed to get eth balance of user", err.Error(), http.StatusInternalServerError)
		}
		Y, _ = strconv.ParseUint((ethBalanceString)[2:], 16, 64)
	} else {

		// faucet account token balance
		call := new(EVMCall)
		call.To = token
		call.Data = "0x70a08231000000000000000000000000" + faucetAddress.String()[2:] // balanceOf
		data, err := client.Call("eth_call", *call, "latest")
		if err != nil {
			return common.ServeError(0, "failed to get token balances", err.Error(), http.StatusInternalServerError)
		}
		tokenBalancesString, err := data.GetString()
		if err != nil {
			return common.ServeError(0, "failed to get token balances", err.Error(), http.StatusInternalServerError)
		}
		if tokenBalancesString != "0x" {
			faucetAccountBalance, _ = strconv.ParseUint((tokenBalancesString)[2:], 16, 64)
		}

		// user token balance
		call.To = token
		call.Data = "0x70a08231000000000000000000000000" + address[2:] // balanceOf
		data, err = client.Call("eth_call", *call, "latest")
		if err != nil {
			return common.ServeError(0, "failed to get token balances", err.Error(), http.StatusInternalServerError)
		}
		tokenBalancesString, err = data.GetString()
		if err != nil {
			return common.ServeError(0, "failed to get token balances", err.Error(), http.StatusInternalServerError)
		}
		if tokenBalancesString != "0x" {
			Y, _ = strconv.ParseUint((tokenBalancesString)[2:], 16, 64)
		}
	}

	if Y >= X {
		return common.ServeError(0, "", "the account already has enough balance", http.StatusInternalServerError)
	}

	Z := X - Y

	if Z > faucetAccountBalance {
		return common.ServeError(0, "", "faucet account doesn't have enough balance", http.StatusInternalServerError)
	}

	if Z < M {
		return common.ServeError(0, "", "transfer amount exceed the minimum amount", http.StatusInternalServerError)
	}

	return nil, nil, http.StatusOK
}

func queryEVMFaucetRequestHandle(r *http.Request, chain string) (interface{}, interface{}, int) {
	isSupportedChain, chainConfig := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	_ = r.ParseForm()

	address := r.FormValue("claim")
	if len(address) == 0 {
		return common.ServeError(0, "", "invalid address", http.StatusBadRequest)
	}

	token := r.FormValue("token")
	if len(token) == 0 {
		token = "0x0000000000000000000000000000000000000000"
	}

	res, err, statusCode := queryEVMFaucetFromNode(chainConfig, chainConfig.QuickNode, address, token)
	if err == nil {
		return res, err, statusCode
	}

	res, err, statusCode = queryEVMFaucetFromNode(chainConfig, chainConfig.Infura, address, token)
	if err == nil {
		return res, err, statusCode
	}

	return queryEVMFaucetFromNode(chainConfig, chainConfig.Pokt, address, token)
}

// RegisterEVMFaucetRequest is a function to faucet evm tokens
func RegisterEVMFaucetRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		chain := queries["chain"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-evm-faucet] Entering transactions execute: ", chain)

		if !common.RPCMethods["GET"][config.QueryEVMFaucet].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryEVMFaucet].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-evm-faucet] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryEVMFaucetRequestHandle(r, chain)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryEVMBalances].CachingEnabled)
	}
}
