package evm

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"

	// "github.com/powerman/rpc-codec/jsonrpc2"
	jsonrpc2 "github.com/KeisukeYamashita/go-jsonrpc"
)

// RegisterEVMBlockRoutes registers query status of EVM chains.
func RegisterEVMBlockRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryEVMBlock, QueryEVMBlockRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryEVMBlock, "This is an API to query account address.", true)
}

func queryEVMBlockFromNode(nodeInfo config.EVMNodeConfig, blockHeightOrHash string) (interface{}, interface{}, int) {
	client := jsonrpc2.NewRPCClient(nodeInfo.RPC + "/" + nodeInfo.RPCToken)
	if nodeInfo.RPCSecret != "" {
		client.SetBasicAuth(nodeInfo.RPCToken, nodeInfo.RPCSecret)
	}

	response := new(interface{})

	if strings.HasPrefix(blockHeightOrHash, "0x") {
		data, err := client.Call("eth_getBlockByHash", blockHeightOrHash, true)
		if err != nil {
			return common.ServeError(0, "failed to get block by hash", err.Error(), http.StatusInternalServerError)
		}
		err = data.GetObject(response)
		if err != nil {
			return common.ServeError(0, "failed to get block by hash", err.Error(), http.StatusInternalServerError)
		}
	} else {
		blockHeight, err := strconv.ParseUint(blockHeightOrHash, 10, 64)

		data, err := client.Call("eth_getBlockByNumber", "0x"+fmt.Sprintf("%X", blockHeight), true)
		if err != nil {
			return common.ServeError(0, "failed to get block by number", err.Error(), http.StatusInternalServerError)
		}
		err = data.GetObject(response)
		if err != nil {
			return common.ServeError(0, "failed to get block by number", err.Error(), http.StatusInternalServerError)
		}
	}

	return response, nil, http.StatusOK
}

func queryEVMBlockRequestHandle(r *http.Request, chain string, blockHeightOrHash string) (interface{}, interface{}, int) {

	isSupportedChain, chainConfig := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	res, err, statusCode := queryEVMBlockFromNode(chainConfig.QuickNode, blockHeightOrHash)
	if err == nil {
		return res, err, statusCode
	}

	res, err, statusCode = queryEVMBlockFromNode(chainConfig.Infura, blockHeightOrHash)
	if err == nil {
		return res, err, statusCode
	}

	return queryEVMBlockFromNode(chainConfig.Pokt, blockHeightOrHash)
}

// QueryEVMBlockRequest is a function to query block of EVM chains.
func QueryEVMBlockRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		chain := queries["chain"]
		blockHeightOrHash := queries["identifier"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-evm-block] Entering account query: ", chain)

		if !common.RPCMethods["GET"][config.QueryEVMBlock].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryEVMBlock].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-evm-block] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryEVMBlockRequestHandle(r, chain, blockHeightOrHash)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryEVMBlock].CachingEnabled)
	}
}
