package kira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	database "github.com/KiraCore/interx/database"
	"github.com/KiraCore/interx/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type NetworkProperties struct {
	MinTxFee                     string `json:"minTxFee"`
	MaxTxFee                     string `json:"maxTxFee"`
	VoteQuorum                   string `json:"voteQuorum"`
	MinimumProposalEndTime       string `json:"minimumProposalEndTime"`
	ProposalEnactmentTime        string `json:"proposalEnactmentTime"`
	MinProposalEndBlocks         string `json:"minProposalEndBlocks"`
	MinProposalEnactmentBlocks   string `json:"minProposalEnactmentBlocks"`
	EnableForeignFeePayments     bool   `json:"enableForeignFeePayments"`
	MischanceRankDecreaseAmount  string `json:"mischanceRankDecreaseAmount"`
	MaxMischance                 string `json:"maxMischance"`
	MischanceConfidence          string `json:"mischanceConfidence"`
	InactiveRankDecreasePercent  string `json:"inactiveRankDecreasePercent"`
	MinValidators                string `json:"minValidators"`
	PoorNetworkMaxBankSend       string `json:"poorNetworkMaxBankSend"`
	UnjailMaxTime                string `json:"unjailMaxTime"`
	EnableTokenWhitelist         bool   `json:"enableTokenWhitelist"`
	EnableTokenBlacklist         bool   `json:"enableTokenBlacklist"`
	MinIdentityApprovalTip       string `json:"minIdentityApprovalTip"`
	UniqueIdentityKeys           string `json:"uniqueIdentityKeys"`
	UbiHardcap                   string `json:"ubiHardcap"`
	ValidatorsFeeShare           string `json:"validatorsFeeShare"`
	InflationRate                string `json:"inflationRate"`
	InflationPeriod              string `json:"inflationPeriod"`
	UnstakingPeriod              string `json:"unstakingPeriod"`
	MaxDelegators                string `json:"maxDelegators"`
	MinDelegationPushout         string `json:"minDelegationPushout"`
	SlashingPeriod               string `json:"slashingPeriod"`
	MaxJailedPercentage          string `json:"maxJailedPercentage"`
	MaxSlashingPercentage        string `json:"maxSlashingPercentage"`
	MinCustodyReward             string `json:"minCustodyReward"`
	MaxCustodyBufferSize         string `json:"maxCustodyBufferSize"`
	MaxCustodyTxSize             string `json:"maxCustodyTxSize"`
	AbstentionRankDecreaseAmount string `json:"abstentionRankDecreaseAmount"`
	MaxAbstention                string `json:"maxAbstention"`
	MinCollectiveBond            string `json:"minCollectiveBond"`
	MinCollectiveBondingTime     string `json:"minCollectiveBondingTime"`
	MaxCollectiveOutputs         string `json:"maxCollectiveOutputs"`
	MinCollectiveClaimPeriod     string `json:"minCollectiveClaimPeriod"`
	ValidatorRecoveryBond        string `json:"validatorRecoveryBond"`
	MaxAnnualInflation           string `json:"maxAnnualInflation"`
	MaxProposalTitleSize         string `json:"maxProposalTitleSize"`
	MaxProposalDescriptionSize   string `json:"maxProposalDescriptionSize"`
	MaxProposalPollOptionSize    string `json:"maxProposalPollOptionSize"`
	MaxProposalPollOptionCount   string `json:"maxProposalPollOptionCount"`
	MaxProposalReferenceSize     string `json:"maxProposalReferenceSize"`
	MaxProposalChecksumSize      string `json:"maxProposalChecksumSize"`
	MinDappBond                  string `json:"minDappBond"`
	MaxDappBond                  string `json:"maxDappBond"`
	DappBondDuration             string `json:"dappBondDuration"`
	DappVerifierBond             string `json:"dappVerifierBond"`
	DappAutoDenounceTime         string `json:"dappAutoDenounceTime"`
}

type NetworkPropertiesResponse struct {
	Properties *NetworkProperties `json:"properties"`
}

// RegisterKiraGovRoutes registers kira gov query routers.
func RegisterKiraGovRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryDataReferenceKeys, QueryDataReferenceKeysRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryDataReference, QueryDataReferenceRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryNetworkProperties, QueryNetworkPropertiesRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryExecutionFee, QueryExecutionFeeRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryExecutionFees, QueryExecutionFeesRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryDataReferenceKeys, "This is an API to query all data reference keys.", true)
	common.AddRPCMethod("GET", config.QueryDataReference, "This is an API to query data reference by key.", true)
	common.AddRPCMethod("GET", config.QueryNetworkProperties, "This is an API to query network properties.", true)
	common.AddRPCMethod("GET", config.QueryExecutionFee, "This is an API to query execution fee by transaction type.", true)
	common.AddRPCMethod("GET", config.QueryExecutionFees, "This is an API to query all execution fees.", true)
}

func queryDataReferenceKeysHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	queries := r.URL.Query()
	key := queries["key"]
	offset := queries["offset"]
	limit := queries["limit"]
	countTotal := queries["count_total"]

	var events = make([]string, 0, 4)
	if len(key) == 1 {
		events = append(events, fmt.Sprintf("pagination.key=%s", key[0]))
	}
	if len(offset) == 1 {
		events = append(events, fmt.Sprintf("pagination.offset=%s", offset[0]))
	}
	if len(limit) == 1 {
		events = append(events, fmt.Sprintf("pagination.limit=%s", limit[0]))
	}
	if len(countTotal) == 1 {
		events = append(events, fmt.Sprintf("pagination.count_total=%s", countTotal[0]))
	}

	r.URL.RawQuery = strings.Join(events, "&")

	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/gov", "/kira/gov", -1)

	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryDataReferenceKeysRequest is a function to query data reference keys.
func QueryDataReferenceKeysRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-reference-keys] Entering data reference keys query")

		if !common.RPCMethods["GET"][config.QueryDataReferenceKeys].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryDataReferenceKeys].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-reference-keys] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryDataReferenceKeysHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryDataReferenceKeys].CachingEnabled)
	}
}

func queryDataReferenceHandle(r *http.Request, gwCosmosmux *runtime.ServeMux, key string) (interface{}, interface{}, int) {
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/gov", "/kira/gov", -1)
	success, failure, status := common.ServeGRPC(r, gwCosmosmux)

	if success != nil {
		type DataReferenceTempResponse struct {
			Data types.DataReferenceEntry `json:"data"`
		}
		result := DataReferenceTempResponse{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-reference] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-reference] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		success = result.Data

		filePath := key + filepath.Ext(result.Data.Reference)

		database.AddReference(key, result.Data.Reference, 0, time.Now().UTC(), config.DataReferenceRegistry+"/"+common.GetMD5Hash(filePath))
	}

	return success, failure, status
}

// QueryDataReferenceRequest is a function to query data reference by key.
func QueryDataReferenceRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		queries := mux.Vars(r)
		key := queries["key"]
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-reference] Entering data reference query by key: ", key)

		if !common.RPCMethods["GET"][config.QueryDataReference].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryDataReference].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-reference] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryDataReferenceHandle(r, gwCosmosmux, key)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryDataReference].CachingEnabled)
	}
}

func QueryNetworkPropertiesHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/gov", "/kira/gov", -1)
	success, failure, status := common.ServeGRPC(r, gwCosmosmux)
	if success != nil {
		result := NetworkPropertiesResponse{}
		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-network-properties] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-network-properties] Invalid response format", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		result.Properties.InactiveRankDecreasePercent = convertRate(result.Properties.InactiveRankDecreasePercent)
		result.Properties.ValidatorsFeeShare = convertRate(result.Properties.ValidatorsFeeShare)
		result.Properties.InflationRate = convertRate(result.Properties.InflationRate)
		result.Properties.MaxSlashingPercentage = convertRate(result.Properties.MaxSlashingPercentage)
		result.Properties.MaxAnnualInflation = convertRate(result.Properties.MaxAnnualInflation)
		result.Properties.DappVerifierBond = convertRate(result.Properties.DappVerifierBond)

		success = result
	}
	return success, failure, status
}

// QueryDataReferenceKeysRequest is a function to query data reference keys.
func QueryNetworkPropertiesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-network-properties] Entering network properties query")

		if !common.RPCMethods["GET"][config.QueryNetworkProperties].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryNetworkProperties].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-network-properties] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = QueryNetworkPropertiesHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryNetworkProperties].CachingEnabled)
	}
}

func QueryExecutionFeeHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	err := r.ParseForm()
	if err != nil {
		return common.ServeError(0, "failed to parse query parameters", err.Error(), http.StatusBadRequest)
	}

	message := r.FormValue("message")
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/gov/execution_fee", "/kira/gov/execution_fee/"+message, -1)
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryExecutionFeeRequest is a function to query execution fee by transaction type.
func QueryExecutionFeeRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-execution-fee] Entering execution fee query")

		if !common.RPCMethods["GET"][config.QueryExecutionFee].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryExecutionFee].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-execution-fee] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = QueryExecutionFeeHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryExecutionFee].CachingEnabled)
	}
}

func QueryExecutionFeesHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	r.URL.Path = strings.Replace(r.URL.Path, "/api/kira/gov/execution_fees", "/kira/gov/all_execution_fees", -1)
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryExecutionFeeRequest is a function to query execution fee by transaction type.
func QueryExecutionFeesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)

		common.GetLogger().Info("[query-execution-fees] Entering execution fees query")

		if !common.RPCMethods["GET"][config.QueryExecutionFees].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryExecutionFees].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-execution-fees] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = QueryExecutionFeesHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryExecutionFees].CachingEnabled)
	}
}
