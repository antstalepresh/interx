package kira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// RegisterKiraTokensRoutes registers kira tokens query routers.
func RegisterKiraTokensRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryKiraTokensAliases, QueryKiraTokensAliasesRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryKiraTokensRates, QueryKiraTokensRatesRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryKiraTokensAliases, "This is an API to query all tokens aliases.", true)
	common.AddRPCMethod("GET", config.QueryKiraTokensRates, "This is an API to query all tokens rates.", true)
}

func queryKiraTokensAliasesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	type TokenAliasesResult struct {
		Decimals int64    `json:"decimals"`
		Denoms   []string `json:"denoms"`
		Name     string   `json:"name"`
		Symbol   string   `json:"symbol"`
		Icon     string   `json:"icon"`
		Amount   int64    `json:"amount,string"`
	}

	tokens := common.GetTokenAliases(gwCosmosmux, r.Clone(r.Context()))
	tokensSupply := common.GetTokenSupply(gwCosmosmux, r.Clone(r.Context()))

	fmt.Println(tokens, tokensSupply)

	result := make([]TokenAliasesResult, 0)

	for _, token := range tokens {
		flag := false
		for _, denom := range token.Denoms {
			for _, supply := range tokensSupply {
				if denom == supply.Denom {
					result = append(result, TokenAliasesResult{
						Decimals: token.Decimals,
						Denoms:   token.Denoms,
						Name:     token.Name,
						Symbol:   token.Symbol,
						Icon:     token.Icon,
						Amount:   supply.Amount,
					})

					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}

	return result, nil, http.StatusOK
}

// QueryKiraTokensAliasesRequest is a function to query all tokens aliases.
func QueryKiraTokensAliasesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-aliases] Entering token aliases query")

		if !common.RPCMethods["GET"][config.QueryKiraTokensAliases].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryKiraTokensAliases].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-aliases] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryKiraTokensAliasesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryKiraTokensAliases].CachingEnabled)
	}
}

func convertRate(rateString string) string {
	rate, _ := strconv.ParseFloat(rateString, 64)
	rate = rate / 1000000000000000000.0
	rateString = fmt.Sprintf("%g", rate)
	if !strings.Contains(rateString, ".") {
		rateString = rateString + ".0"
	}
	return rateString
}

func queryKiraTokensRatesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/tokens", "/kira/tokens", -1)
	success, failure, status := common.ServeGRPC(r, gwCosmosmux)

	if success != nil {
		type TokenRate struct {
			Denom       string `json:"denom"`
			FeePayments bool   `json:"feePayments"`
			FeeRate     string `json:"feeRate"`
			StakeCap    string `json:"stakeCap"`
			StakeMin    string `json:"stakeMin"`
			StakeToken  bool   `json:"stakeToken"`
		}

		type TokenRatesResponse struct {
			Data []TokenRate `json:"data"`
		}
		result := TokenRatesResponse{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-token-rates] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-token-rates] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		for index, tokenRate := range result.Data {
			result.Data[index].FeeRate = convertRate(tokenRate.FeeRate)
			result.Data[index].StakeCap = convertRate(tokenRate.StakeCap)
			result.Data[index].StakeMin = convertRate(tokenRate.StakeMin)
		}

		success = result
	}

	return success, failure, status
}

// QueryKiraTokensRatesRequest is a function to query all tokens rates.
func QueryKiraTokensRatesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-rates] Entering token rates query")

		if !common.RPCMethods["GET"][config.QueryKiraTokensRates].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryKiraTokensRates].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-rates] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryKiraTokensRatesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryKiraTokensRates].CachingEnabled)
	}
}
