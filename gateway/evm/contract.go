package evm

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	jsonrpc2 "github.com/KeisukeYamashita/go-jsonrpc"
	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/gorilla/mux"

	// "github.com/powerman/rpc-codec/jsonrpc2"
	goeth "github.com/ethereum/go-ethereum/common"
)

// RegisterEVMContractRoutes registers read/write of smart contract.
func RegisterEVMContractRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryReadContract, QueryReadContractRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryReadContract, "This is an API to read smart contract.", true)
}

func ReadContractCall(nodeInfo config.EVMNodeConfig, call *EVMCall, method abi.Method) (interface{}, error) {
	client := jsonrpc2.NewRPCClient(nodeInfo.RPC + "/" + nodeInfo.RPCToken)
	if nodeInfo.RPCSecret != "" {
		client.SetBasicAuth(nodeInfo.RPCToken, nodeInfo.RPCSecret)
	}

	data, err := client.Call("eth_call", *call, "latest")
	if err != nil {
		return nil, err
	}

	result, err := data.GetString()
	if err != nil {
		return nil, err
	}

	bytes, err := hex.DecodeString(result[2:])
	if err != nil {
		return nil, err
	}

	res, err := method.Outputs.Unpack(bytes)

	return res, err
}

func decodeAbiArgument(param string, arg *abi.Argument) (v interface{}, err error) {
	param = strings.TrimSpace(param)
	switch arg.Type.T {
	case abi.StringTy:
		str_val := new(string)
		v = str_val
		err = json.Unmarshal([]byte(param), v)
	case abi.UintTy, abi.IntTy:
		val := big.NewInt(0)
		_, success := val.SetString(param, 10)
		if !success {
			err = errors.New(fmt.Sprintf("Invalid numeric (base 10) value: %v", param))
		}
		v = val
	case abi.AddressTy:
		if !((len(param) == (goeth.AddressLength*2 + 2)) || (len(param) == goeth.AddressLength*2)) {
			err = errors.New(fmt.Sprintf("Invalid address length (%v), must be 40 (unprefixed) or 42 (prefixed) chars", len(param)))
		} else {
			var addr goeth.Address
			if len(param) == (goeth.AddressLength*2 + 2) {
				addr = goeth.HexToAddress(param)
			} else {
				var data []byte
				data, err = hex.DecodeString(param)
				addr.SetBytes(data)
			}
			v = addr
		}
	case abi.HashTy:
		if !((len(param) == (goeth.HashLength*2 + 2)) || (len(param) == goeth.HashLength*2)) {
			err = errors.New(fmt.Sprintf("Invalid hash length, must be 64 (unprefixed) or 66 (prefixed) chars"))
		} else {
			var hash goeth.Hash
			if len(param) == (goeth.HashLength*2 + 2) {
				hash = goeth.HexToHash(param)
			} else {
				var data []byte
				data, err = hex.DecodeString(param)
				hash.SetBytes(data)
			}
			v = hash
		}
	case abi.BytesTy:
		if len(param) > 2 {
			if (param[0] == '0') && (param[1] == 'x') {
				param = param[2:] // cut 0x prefix
			}
		}
		decoded_bytes, tmperr := hex.DecodeString(param)
		v = decoded_bytes
		err = tmperr
	case abi.BoolTy:
		val := new(bool)
		v = val
		err = json.Unmarshal([]byte(param), v)
	default:
		err = errors.New(fmt.Sprintf("Not supported parameter type: %v", arg.Type))
	}
	return v, err
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
			return common.ServeError(0, "Invalid ABI JSON", err.Error(), http.StatusBadRequest)
		}

		abiDecoded, err = abi.JSON(strings.NewReader(abiJson))
		if err != nil {
			return common.ServeError(0, "Invalid ABI JSON", err.Error(), http.StatusBadRequest)
		}
	}

	method, exist := abiDecoded.Methods[functionName]
	if !exist {
		return common.ServeError(0, "", "function does not exist", http.StatusBadRequest)
	}

	argLen := len(method.Inputs)
	var args []interface{}
	for i := 0; i < argLen; i++ {
		arg := r.FormValue("key_" + strconv.Itoa(i+1))
		if len(arg) == 0 {
			return common.ServeError(0, "", "argument invalid", http.StatusBadRequest)
		}

		v, err := decodeAbiArgument(arg, &method.Inputs[i])
		if err != nil {
			return common.ServeError(0, "", "argument invalid", http.StatusBadRequest)
		}

		args = append(args, v)
	}

	bytes, err := abiDecoded.Pack(functionName, args...)
	if err != nil {
		return common.ServeError(0, "", err.Error(), http.StatusBadRequest)
	}

	call := new(EVMCall)
	call.To = contract
	call.Data = "0x" + hex.EncodeToString(bytes)

	res, err := ReadContractCall(chainConfig.QuickNode, call, method)
	if err != nil {
		res, err = ReadContractCall(chainConfig.Infura, call, method)
		if err != nil {
			res, err = ReadContractCall(chainConfig.Pokt, call, method)
		}
	}

	if err != nil {
		return common.ServeError(0, err.Error(), "contract call fail", http.StatusBadRequest)
	}

	return res, nil, http.StatusOK
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
