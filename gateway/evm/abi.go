package evm

import (
	"encoding/json"
	"net/http"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"
	// "github.com/powerman/rpc-codec/jsonrpc2"
)

// RegisterEVMAbiRoutes registers query abi of smart contract.
func RegisterEVMAbiRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryABI, QueryAbiRequests(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryABI, "This is an API to query abi.", true)
}

func queryAbiHandle(r *http.Request, chain string, contract string) (interface{}, interface{}, int) {
	isSupportedChain, chainConfig := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	result, err, statusCode := common.MakeGetRequest(chainConfig.Etherscan.API, "", "module=contract&action=getabi&address="+contract+"&apikey="+chainConfig.Etherscan.APIToken)
	abi := new(interface{})
	common.GetLogger().Info(result.(map[string]interface{})["result"])
	json.Unmarshal([]byte(result.(map[string]interface{})["result"].(string)), abi)

	type EVMAbi struct {
		Abi interface{} `json:"abi"`
	}
	response := new(EVMAbi)
	response.Abi = abi
	return response, err, statusCode
}

// QueryAbiRequests is a function to query abi of smart contract.
func QueryAbiRequests(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		chain := queries["chain"]
		contract := queries["contract"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-evm-abi] Entering abi query: ", chain)

		if !common.RPCMethods["GET"][config.QueryABI].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryABI].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-evm-abi] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryAbiHandle(r, chain, contract)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryABI].CachingEnabled)
	}
}
