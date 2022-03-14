package kira

import (
	"net/http"
	"strings"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterKiraSpendingRoutes registers kira spending query routers.
func RegisterKiraSpendingRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QuerySpendingPools, QuerySpendingPoolsRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QuerySpendingPools, "This is an API to query list of spending pool names.", true)
}

func QuerySpendingPoolsHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/spending-pools", "/kira/spending/pool_names", -1)
	return common.ServeGRPC(r, gwCosmosmux)
}

// QuerySpendingPoolsRequest is a function to query list of all spending pool names.
func QuerySpendingPoolsRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-spending-pool-names] Entering upgrade plan query")

		if !common.RPCMethods["GET"][config.QuerySpendingPools].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QuerySpendingPools].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-spending-pool-names] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = QuerySpendingPoolsHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QuerySpendingPools].CachingEnabled)
	}
}
