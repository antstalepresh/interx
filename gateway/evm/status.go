package evm

import (
	"context"
	"net/http"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// RegisterEVMStatusRoutes registers query status of EVM chains.
func RegisterEVMStatusRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryEVMStatus, QueryEVMStatusRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryEVMStatus, "This is an API to query account address.", true)
}

type EVMStatus struct {
	NodeInfo struct {
		Network string `json:"network"`
	} `json:"node_info"`
}

func queryEVMStatusHandle(r *http.Request, chain string) (interface{}, interface{}, int) {

	isSupportedChain, chainConfig := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	response := EVMStatus{}

	// infuraClient, err := ethclient.Dial(chainConfig.Infura.RPC + "/" + chainConfig.Infura.RPCToken)
	// if err != nil {
	// 	return common.ServeError(0, "failed to get infura client", err.Error(), http.StatusInternalServerError)
	// }

	ctx := context.WithValue(context.Background(), "user", chainConfig.Pokt.RPCSecret)
	poktClient, err := ethclient.Dial(chainConfig.Pokt.RPC + "/" + chainConfig.Pokt.RPCToken)
	if err != nil {
		return common.ServeError(0, "failed to get pokt client", err.Error(), http.StatusInternalServerError)
	}
	// https: //stackoverflow.com/questions/24230500/go-http-post-request-with-basic-auth-and-formvalue
	chainId, err := poktClient.ChainID(ctx)
	if err != nil {
		return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
	}

	response.NodeInfo.Network = chainId.String()

	return nil, nil, http.StatusOK
}

// QueryEVMStatusRequest is a function to query status of EVM chains.
func QueryEVMStatusRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := mux.Vars(r)
		chain := queries["chain"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-evm-status] Entering account query: ", chain)

		if !common.RPCMethods["GET"][config.QueryEVMStatus].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryEVMStatus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-evm-status] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryEVMStatusHandle(r, chain)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryEVMStatus].CachingEnabled)
	}
}
