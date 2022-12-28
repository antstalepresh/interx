package bitcoin

import (
	"fmt"
	"net/http"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/gorilla/mux"

	"github.com/btcsuite/btcd/rpcclient"
)

// RegisterBitcoinStatusRoutes registers query status of Bitcoin chains.
func RegisterBitcoinStatusRoutes(r *mux.Router, rpcAddr string) {
	r.HandleFunc(config.QueryBitcoinStatus, QueryBitcoinStatusRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryBitcoinStatus, "This is an API to query status.", true)
}

type BitcoinStatus struct {
	NodeInfo struct {
		Network    string `json:"network"`
		RPCAddress string `json:"rpc_address"`
		Version    struct {
			Net      string `json:"net"`
			Sub      string `json:"sub"`
			Protocol string `json:"protocol"`
		} `json:"version"`
	} `json:"node_info"`
	SyncInfo struct {
		CatchingUp          bool   `json:"catching_up"`
		EarliestBlockHash   string `json:"earliest_block_hash"`
		EarliestBlockHeight uint64 `json:"earliest_block_height"`
		EarliestBlockTime   uint64 `json:"earliest_block_time"`
		LatestBlockHash     string `json:"latest_block_hash"`
		LatestBlockHeight   uint64 `json:"latest_block_height"`
		LatestBlockTime     uint64 `json:"latest_block_time"`
	} `json:"sync_info"`
	GasPrice    uint64 `json:"gas_price"`
	GasPriceAvg uint64 `json:"gas_price_avg"`
	GasPriceMin uint64 `json:"gas_price_min"`
	GasPriceInc uint64 `json:"gas_price_inc"`
}

func queryBitcoinStatusHandle(r *http.Request, chain string) (interface{}, interface{}, int) {

	isSupportedChain, _ := GetChainConfig(chain)
	if !isSupportedChain {
		return common.ServeError(0, "", "unsupported chain", http.StatusBadRequest)
	}

	config := rpcclient.ConnConfig{
		Host:         "65.109.5.45:18332",
		User:         "bitcoin",
		Pass:         "8c0b93fd94b0a2bfe8abcb0ecf5f9f1d43a4dfa9bc300fc3",
		DisableTLS:   true,
		HTTPPostMode: true,
	}

	client, err := rpcclient.New(&config, nil)
	if err != nil {
		fmt.Println(err)
		return common.ServeError(0, "failed to make initialize connection ", err.Error(), http.StatusInternalServerError)
	}

	response := BitcoinStatus{}
	chainInfo, err := client.GetBlockChainInfo()
	if err != nil {
		return common.ServeError(0, "failed to get blockchain info ", err.Error(), http.StatusInternalServerError)
	}
	_, err = client.GetNetworkInfo()
	if err != nil {
		return common.ServeError(0, "failed to get network info ", err.Error(), http.StatusInternalServerError)
	}

	response.NodeInfo.Network = chainInfo.Chain
	return response, nil, http.StatusOK
}

// QueryBitcoinStatusRequest is a function to query status of Bitcoin chains.
func QueryBitcoinStatusRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chain := "testnet"
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-bitcoin-status] Entering status query: ", chain)

		if !common.RPCMethods["GET"][config.QueryBitcoinStatus].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryBitcoinStatus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-bitcoin-status] Returning from the cache: ", chain)
					return
				}
			}

			response.Response, response.Error, statusCode = queryBitcoinStatusHandle(r, chain)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryBitcoinStatus].CachingEnabled)
	}
}
