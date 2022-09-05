package evm

import (
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gorilla/mux"
	// "github.com/powerman/rpc-codec/jsonrpc2"
)

// RegisterEVMContractRoutes registers read/write of smart contract.
func RegisterEVMContractRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryReadContract, QueryReadContractRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryReadContract, "This is an API to read smart contract.", true)
}

func queryReadSmartContractHandle(r *http.Request, chain string, contract string) (interface{}, interface{}, int) {
	isSupportedChain, chainConfig := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	_ = r.ParseForm()

	functionName := r.FormValue("function")
	if len(functionName) == 0 {
		return common.ServeError(0, "", "no function name", http.StatusBadRequest)
	}

	var abiDecoded abi.ABI
	abiJsonUrlEncoded := r.FormValue("abi")
	if len(abiJsonUrlEncoded) == 0 {
		result, err, statusCode := common.MakeGetRequest(chainConfig.Etherscan.API, "", "module=contract&action=getabi&address="+contract+"&apikey="+chainConfig.Etherscan.APIToken)
		if err != nil {
			return nil, err, statusCode
		}

		abiDecoded, err = abi.JSON(strings.NewReader(result.(map[string]interface{})["result"].(string)))
		if err != nil {
			return nil, err, http.StatusBadRequest
		}
	} else {
		abiJson, err := url.QueryUnescape(abiJsonUrlEncoded)
		if err != nil {
			return common.ServeError(0, "", err.Error(), http.StatusBadRequest)
		}

		abiDecoded, err = abi.JSON(strings.NewReader(abiJson))
		if err != nil {
			return common.ServeError(0, "", err.Error(), http.StatusBadRequest)
		}
	}

	method, exist := abiDecoded.Methods[functionName]
	if !exist {
		return common.ServeError(0, "", "function does not exist", http.StatusBadRequest)
	}
	argLen := len(method.Inputs)
	args := []interface{}{}
	for i := 0; i < argLen; i++ {
		arg := r.FormValue("key_" + strconv.Itoa(i))
		if len(arg) == 0 {
			return common.ServeError(0, "", "argument invalid", http.StatusBadRequest)
		}

		args = append(args, arg)
	}

	bytes, err := abiDecoded.Pack(functionName, args...)
	if err != nil {
		return common.ServeError(0, "", err.Error(), http.StatusBadRequest)
	}

	call := new(EVMCall)
	call.To = contract
	call.Data = "0x" + hex.EncodeToString(bytes)
	// data, err = client.Call("eth_call", *call, "latest")

	return nil, nil, http.StatusOK
}

// QueryReadContractRequest is a function to read smart contract.
func QueryReadContractRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		chain := queries["chain"]
		contract := queries["contract"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-evm-read-contract] Entering read smart contract: ", chain)

		if !common.RPCMethods["GET"][config.QueryReadContract].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryReadContract].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-evm-read-contract] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryReadSmartContractHandle(r, chain, contract)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryReadContract].CachingEnabled)
	}
}
