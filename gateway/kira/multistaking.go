package kira

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/KiraCore/interx/tasks"
	"github.com/KiraCore/interx/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type QueryStakingPoolDelegatorsResponse struct {
	Pool       types.ValidatorPool `json:"pool"`
	Delegators []string            `json:"delegators,omitempty"`
}

func RegisterKiraMultiStakingRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryStakingPool, QueryStakingPoolRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryStakingPool, "This is an API to query staking pool.", true)
}

func queryStakingPoolHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	queries := r.URL.Query()
	account := queries["validatorAddress"]
	tokenPath := ""

	if len(account) == 1 {
		valAddr, found := tasks.AddrToValidator[account[0]]
		if found {
			r.URL.RawQuery = ""
			tokenPath = strings.Replace(r.URL.Path, "/api/kira/staking-pool", "/kira/tokens/rates", -1)
			r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/staking-pool", "/kira/multistaking/v1beta1/staking_pool_delegators/"+valAddr, -1)
		}
	}

	success, failure, status := common.ServeGRPC(r, gwCosmosmux)
	if success != nil {
		result := QueryStakingPoolDelegatorsResponse{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-staking-pool] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-staking-pool] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		response := types.ValidatorPoolResult{}
		response.ID = result.Pool.ID
		response.Slashed = convertRate(result.Pool.Slashed)
		response.Commission = convertRate(result.Pool.Commission)
		response.VotingPower = result.Pool.TotalStakingTokens
		response.TotalDelegators = int64(len(result.Delegators))
		response.Tokens = []string{}

		r.URL.RawQuery = ""
		r.URL.Path = tokenPath
		successTokens, _, _ := common.ServeGRPC(r, gwCosmosmux)
		if successTokens != nil {
			type TokenRatesResponse struct {
				Data []types.TokenRate `json:"data"`
			}

			res := TokenRatesResponse{}
			byteData, err := json.Marshal(successTokens)
			if err != nil {
				common.GetLogger().Error("[query-staking-pool(token-rates)] Invalid response format", err)
				return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
			}
			err = json.Unmarshal(byteData, &res)
			if err != nil {
				common.GetLogger().Error("[query-staking-pool(token-rates)] Invalid response format", err)
				return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
			}

			for _, tokenRate := range res.Data {
				response.Tokens = append(response.Tokens, tokenRate.Denom)
			}
		}
		success = response
	}
	return success, failure, status
}

// QueryStakingPoolRequest is a function to query staking pool with given validator address.
func QueryStakingPoolRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-staking-pool] Entering staking pool query")

		if !common.RPCMethods["GET"][config.QueryStakingPool].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryStakingPool].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-staking-pool] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryStakingPoolHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryStakingPool].CachingEnabled)
	}
}
