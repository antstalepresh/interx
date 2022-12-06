package interx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/config"
	"github.com/KiraCore/interx/database"
	"github.com/KiraCore/interx/types"
	kiratypes "github.com/KiraCore/sekai/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmTypes "github.com/tendermint/tendermint/rpc/core/types"
	tmJsonRPCTypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"golang.org/x/exp/slices"
)

// RegisterInterxTxRoutes registers tx query routers.
func RegisterInterxTxRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryUnconfirmedTxs, QueryUnconfirmedTxs(rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryTransactions, QueryTransactions(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryKiraFunctions, "This is an API to query kira functions and metadata.", true)
	common.AddRPCMethod("GET", config.QueryUnconfirmedTxs, "This is an API to query unconfirmed transactions.", true)
	common.AddRPCMethod("GET", config.QueryTransactions, "This is an API to query transactions.", true)
}

func toSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// GetTransactionsWithSync is a function to sync user transactions and return result
func GetTransactionsWithSync(rpcAddr string, address string) (*tmTypes.ResultTxSearch, error) {
	var page = 1
	var limit = 100
	var limitPages = 100

	if address == "" {
		return &tmTypes.ResultTxSearch{}, nil
	}

	lastBlock := database.GetLastBlockFetched(address)

	for page < limitPages {
		var events = make([]string, 0, 5)
		events = append(events, fmt.Sprintf("message.sender='%s'", address))
		events = append(events, fmt.Sprintf("tx.Height>=%d", lastBlock))

		// search transactions
		endpoint := fmt.Sprintf("%s/tx_search?query=\"%s\"&page=%d&&per_page=%d&order_by=\"desc\"", rpcAddr, strings.Join(events, "%20AND%20"), page, limit)
		common.GetLogger().Info("[query-transaction] Entering transaction search: ", endpoint)

		resp, err := http.Get(endpoint)
		if err != nil {
			common.GetLogger().Error("[query-transaction] Unable to connect to ", endpoint)
			return nil, err
		}
		defer resp.Body.Close()

		respBody, _ := ioutil.ReadAll(resp.Body)

		response := new(tmJsonRPCTypes.RPCResponse)

		if err := json.Unmarshal(respBody, response); err != nil {
			common.GetLogger().Error("[query-transaction] Unable to decode response: ", err)
			return nil, err
		}

		if response.Error != nil {
			common.GetLogger().Error("[query-transaction] Error response:", response.Error.Message)
			break
		}

		result := new(tmTypes.ResultTxSearch)
		if err := tmjson.Unmarshal(response.Result, result); err != nil {
			common.GetLogger().Error("[query-transaction] Failed to unmarshal result:", err)
			return nil, fmt.Errorf("error unmarshalling result: %w", err)
		}

		database.SaveTransactions(address, *result)
		page++
	}

	return database.GetTransactions(address)
}

// GetFilteredTransactions is a function to filter transactions by various options
func GetFilteredTransactions(rpcAddr string, address string, types []string, directions []string, dateStart int, dateEnd int, statuses []string, sortBy string) (*tmTypes.ResultTxSearch, error) {
	result, _ := searchUnconfirmed(rpcAddr, "100")
	hashToStatusMap := make(map[string]string)
	for _, unconfirmedTx := range result.Txs {
		hashToStatusMap[string(unconfirmedTx.Hash())] = "unconfirmed"
	}

	cachedTxs, err := GetTransactionsWithSync(rpcAddr, address)
	if err != nil {
		return nil, err
	}

	// Filter transactions by date start/end timestamp and msg type
	filteredTxsByTimeMsgStatus := &tmTypes.ResultTxSearch{
		Txs:        []*tmTypes.ResultTx{},
		TotalCount: 0,
	}
	for _, cachedTx := range cachedTxs.Txs {
		// Filter by time
		txTime, err := common.GetBlockTime(rpcAddr, cachedTx.Height)
		if err != nil {
			common.GetLogger().Error("[query-transactions] Block not found: ", cachedTx.Height)
			continue
		}

		if (dateStart != -1 && txTime < int64(dateStart)) || (dateEnd != -1 && txTime > int64(dateEnd)) {
			continue
		}

		// Filter by msg
		tx, err := config.EncodingCg.TxConfig.TxDecoder()(cachedTx.Tx)
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to decode transaction: ", err)
			continue
		}

		contain := false
		for _, msg := range tx.GetMsgs() {
			txType := kiratypes.MsgType(msg)
			if slices.Contains(types, txType) {
				contain = true
				break
			}
		}

		if !contain {
			continue
		}

		// Filter by status
		if len(statuses) != 0 && !slices.Contains(statuses, "unconfirmed") && hashToStatusMap[cachedTx.Hash.String()] == "unconfirmed" {
			continue
		}

		if len(statuses) != 0 && !slices.Contains(statuses, "confirmed") && hashToStatusMap[cachedTx.Hash.String()] == "" {
			continue
		}

		filteredTxsByTimeMsgStatus.Txs = append(filteredTxsByTimeMsgStatus.Txs, cachedTx)
		filteredTxsByTimeMsgStatus.TotalCount++
	}

	return filteredTxsByTimeMsgStatus, nil

	// // Sort transactions
	// if sortBy == "amountASC" {

	// 	sort.Slice(filteredTxsByTimeMsgStatus.Txs, func(i, j int) bool {
	// 		return filteredTxsByTimeMsgStatus.Txs[i]. < filteredTxsByTimeMsgStatus.Txs[j].
	// 	})
	// }
}

// SearchTxHashHandle is a function to query transactions
func SearchTxHashHandle(rpcAddr string, sender string, recipient string, txType string, page int, limit int, txMinHeight int64, txMaxHeight int64, txHash string) (*tmTypes.ResultTxSearch, error) {
	var events = make([]string, 0, 5)

	if sender != "" {
		events = append(events, fmt.Sprintf("transfer.sender='%s'", sender))
	}

	if recipient != "" {
		events = append(events, fmt.Sprintf("transfer.recipient='%s'", recipient))
	}

	if txType != "all" && txType != "" {
		events = append(events, fmt.Sprintf("message.action='%s'", txType))
	}

	if txHash != "" {
		events = append(events, fmt.Sprintf("tx.hash='%s'", txHash))
	}

	if txMinHeight >= 0 {
		events = append(events, fmt.Sprintf("tx.height>=%d", txMinHeight))
	}

	if txMaxHeight >= 0 {
		events = append(events, fmt.Sprintf("tx.height<=%d", txMaxHeight))
	}

	// search transactions
	endpoint := fmt.Sprintf("%s/tx_search?query=\"%s\"&page=%d&&per_page=%d&order_by=\"desc\"", rpcAddr, strings.Join(events, "%20AND%20"), page, limit)
	if page == 0 {
		endpoint = fmt.Sprintf("%s/tx_search?query=\"%s\"&per_page=%d&order_by=\"desc\"", rpcAddr, strings.Join(events, "%20AND%20"), limit)
	}
	common.GetLogger().Info("[query-transaction] Entering transaction search: ", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		common.GetLogger().Error("[query-transaction] Unable to connect to ", endpoint)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	response := new(tmJsonRPCTypes.RPCResponse)

	if err := json.Unmarshal(respBody, response); err != nil {
		common.GetLogger().Error("[query-transaction] Unable to decode response: ", err)
		return nil, err
	}

	if response.Error != nil {
		common.GetLogger().Error("[query-transaction] Error response:", response.Error.Message)
		return nil, errors.New(response.Error.Message)
	}

	result := new(tmTypes.ResultTxSearch)
	if err := tmjson.Unmarshal(response.Result, result); err != nil {
		common.GetLogger().Error("[query-transaction] Failed to unmarshal result:", err)
		return nil, fmt.Errorf("error unmarshalling result: %w", err)
	}

	return result, nil
}

func getBlockHeight(rpcAddr string, hash string) (int64, error) {
	endpoint := fmt.Sprintf("%s/tx?hash=%s", rpcAddr, hash)
	common.GetLogger().Info("[query-block] Entering block query: ", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		common.GetLogger().Error("[query-block] Unable to connect to ", endpoint)
		return 0, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	response := new(tmJsonRPCTypes.RPCResponse)

	if err := json.Unmarshal(respBody, response); err != nil {
		common.GetLogger().Error("[query-block] Unable to decode response: ", err)
		return 0, err
	}
	if response.Error != nil {
		common.GetLogger().Error("[query-block] Error response:", response.Error.Message)
		return 0, errors.New(response.Error.Message)
	}

	result := new(tmTypes.ResultTx)
	if err := tmjson.Unmarshal(response.Result, result); err != nil {
		common.GetLogger().Error("[query-block] Failed to unmarshal result:", err)
		return 0, fmt.Errorf("error unmarshalling result: %w", err)
	}

	return result.Height, nil
}

func convertTxTypeToQueryString(txType string, address string) string {
	return fmt.Sprintf("message.action='%s'%20AND%20message.sender='%s'", config.MsgTypes[txType], address)
}

func QueryBlockTransactionsHandler(rpcAddr string, r *http.Request) (interface{}, interface{}, int) {
	err := r.ParseForm()
	if err != nil {
		common.GetLogger().Error("[query-transactions] Failed to parse query parameters:", err)
		return common.ServeError(0, "failed to parse query parameters", err.Error(), http.StatusBadRequest)
	}

	var (
		query      string = ""
		account    string = ""
		txTypes           = []string{}
		sender     string = ""
		recipient  string = ""
		directions        = []string{}
		statuses          = []string{}
		dateStart  int    = -1
		dateEnd    int    = -1
		sortBy     string = ""
		pageSize   int    = 10
		page       int    = 1
		limit      int    = 10
		offset     int    = 0
	)

	//------------ Type ------------
	txTypesParam := r.FormValue("type")
	txTypesArray := strings.Split(txTypesParam, ",")
	for _, txType := range txTypesArray {
		if config.MsgTypes[txType] != "" {
			txTypes = append(txTypes, config.MsgTypes[txType])
		}
	}

	//------------ Address ------------
	account = r.FormValue("address")
	if account == "" {
		common.GetLogger().Error("[query-transactions] 'address' is not set")
		return common.ServeError(0, "'address' is not set", "", http.StatusBadRequest)
	}

	//------------ Direction ------------
	directionsParam := r.FormValue("direction")
	directionsArray := strings.Split(directionsParam, ",")
	for _, drt := range directionsArray {
		if drt == "inbound" || drt == "outbound" {
			directions = append(directions, drt)
		}
	}

	//------------ Status ------------
	statusesParam := r.FormValue("status")
	statusesArray := strings.Split(statusesParam, ",")
	for _, sts := range statusesArray {
		if sts == "pending" || sts == "confirmed" || sts == "failed" {
			statuses = append(statuses, sts)
		}
	}

	//------------ Sort ------------
	sortParam := r.FormValue("sort")
	if sortParam == "amountASC" || sortParam == "amountDESC" || sortParam == "dateASC" || sortParam == "dateDESC" {
		sortBy = sortParam
	} else {
		sortBy = "dateDESC"
	}

	//------------ Timestamps ------------
	if dateStStr := r.FormValue("dateStart"); dateStStr != "" {
		if dateStart, err = strconv.Atoi(dateStStr); err != nil {
			layout := "01/02/2006 3:04:05 PM"
			t, err1 := time.Parse(layout, dateStStr+" 12:00:00 AM")
			if err1 != nil {
				common.GetLogger().Error("[query-transactions] Failed to parse parameter 'dateStart': ", err)
				return common.ServeError(0, "failed to parse parameter 'dateStart'", err.Error(), http.StatusBadRequest)
			}

			dateStart = int(t.Unix())
		}
	}

	if dateEdStr := r.FormValue("dateEnd"); dateEdStr != "" {
		if dateEnd, err = strconv.Atoi(dateEdStr); err != nil {
			layout := "01/02/2006 3:04:05 PM"
			t, err1 := time.Parse(layout, dateEdStr+" 12:00:00 AM")
			if err1 != nil {
				common.GetLogger().Error("[query-transactions] Failed to parse parameter 'dateEnd': ", err)
				return common.ServeError(0, "failed to parse parameter 'dateEnd'", err.Error(), http.StatusBadRequest)
			}

			dateEnd = int(t.Unix())
		}
	}

	//------------ Pagination ------------
	if pageSizeStr := r.FormValue("page_size"); pageSizeStr != "" {
		if pageSize, err = strconv.Atoi(pageSizeStr); err != nil {
			common.GetLogger().Error("[query-transactions] Failed to parse parameter 'page_size': ", err)
			return common.ServeError(0, "failed to parse parameter 'page_size'", err.Error(), http.StatusBadRequest)
		}
		if pageSize < 1 || pageSize > 100 {
			common.GetLogger().Error("[query-transactions] Invalid 'page_size' range: ", pageSize)
			return common.ServeError(0, "'page_size' should be 1 ~ 100", "", http.StatusBadRequest)
		}
	}

	if pageStr := r.FormValue("page"); pageStr != "" {
		if page, err = strconv.Atoi(pageStr); err != nil {
			common.GetLogger().Error("[query-transactions] Failed to parse parameter 'page': ", err)
			return common.ServeError(0, "failed to parse parameter 'page'", err.Error(), http.StatusBadRequest)
		}
	}

	if limitStr := r.FormValue("limit"); limitStr != "" {
		if limit, err = strconv.Atoi(limitStr); err != nil {
			common.GetLogger().Error("[query-transactions] Failed to parse parameter 'limit': ", err)
			return common.ServeError(0, "failed to parse parameter 'limit'", err.Error(), http.StatusBadRequest)
		}

		if limit < 1 || limit > 100 {
			common.GetLogger().Error("[query-transactions] Invalid 'limit' range: ", limit)
			return common.ServeError(0, "'limit' should be 1 ~ 100", "", http.StatusBadRequest)
		}
		pageSize = limit
	}

	if offsetStr := r.FormValue("offset"); offsetStr != "" {
		if offset, err = strconv.Atoi(offsetStr); err != nil {
			common.GetLogger().Error("[query-transactions] Failed to parse parameter 'offset': ", err)
			return common.ServeError(0, "failed to parse parameter 'offset'", err.Error(), http.StatusBadRequest)
		}

		page = offset / pageSize
		offset = offset % pageSize
	}

	filteredTxs, err := GetFilteredTransactions(rpcAddr, account, txTypes, directions, dateStart, dateEnd, statuses, sortBy)
	if err != nil {
		return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
	}

	var transactions []*tmTypes.ResultTx

	if limit > 0 || offset > 0 {
		searchResult, err := SearchTxHashHandle(rpcAddr, sender, recipient, txType, page, pageSize, -1, -1, "")
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to search transaction hash: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		transactions = searchResult.Txs[offset:]

		searchResult, err = SearchTxHashHandle(rpcAddr, sender, recipient, txType, page+1, pageSize, -1, -1, "")
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to search transaction hash: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		transactions = append(transactions, searchResult.Txs[:offset]...)
	} else {
		searchResult, err := SearchTxHashHandle(rpcAddr, sender, recipient, txType, page, pageSize, -1, -1, "")
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to search transaction hash: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}
		transactions = searchResult.Txs
	}

	var txResults = []types.TransactionResponse{}

	for _, transaction := range transactions {
		tx, err := config.EncodingCg.TxConfig.TxDecoder()(transaction.Tx)
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to decode transaction: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		blockTime, err := common.GetBlockTime(rpcAddr, transaction.Height)
		if err != nil {
			common.GetLogger().Error("[query-transactions] Block not found: ", transaction.Height)
			return common.ServeError(0, "", fmt.Sprintf("block not found: %d", transaction.Height), http.StatusInternalServerError)
		}

		logs, err := sdk.ParseABCILogs(transaction.TxResult.GetLog())
		if err != nil {
			common.GetLogger().Error("[query-transactions] Failed to parse ABCI logs: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		var txResponses []types.TransactionTxResult

		for index, msg := range tx.GetMsgs() {
			txType := kiratypes.MsgType(msg)

			var evMap = make(map[string]([]sdk.Attribute))
			for _, event := range logs[index].GetEvents() {
				evMap[event.GetType()] = event.GetAttributes()
			}

			if txType == "send" {
				msgSend := msg.(*bank.MsgSend)

				if isWithdraw && msgSend.FromAddress == account {
					for _, coin := range msgSend.Amount {
						txResponses = append(txResponses, types.DepositWithdrawTransaction{
							Address: msgSend.ToAddress,
							Type:    txType,
							Denom:   coin.GetDenom(),
							Amount:  coin.Amount.Int64(),
						})
					}
				}
				if !isWithdraw && msgSend.ToAddress == account {
					for _, coin := range msgSend.Amount {
						txResponses = append(txResponses, types.DepositWithdrawTransaction{
							Address: msgSend.FromAddress,
							Type:    txType,
							Denom:   coin.GetDenom(),
							Amount:  coin.Amount.Int64(),
						})
					}
				}
			} else if txType == "multisend" {
				msgMultiSend := msg.(*bank.MsgMultiSend)
				inputs := msgMultiSend.GetInputs()
				outputs := msgMultiSend.GetOutputs()
				if isWithdraw {
					for _, input := range inputs {
						if input.Address == account {
							if len(inputs) == 1 {
								for _, output := range outputs {
									for _, coin := range output.Coins {
										txResponses = append(txResponses, types.DepositWithdrawTransaction{
											Address: output.Address,
											Type:    txType,
											Denom:   coin.GetDenom(),
											Amount:  coin.Amount.Int64(),
										})
									}
								}
							} else if len(outputs) == 1 {
								for _, coin := range input.Coins {
									txResponses = append(txResponses, types.DepositWithdrawTransaction{
										Address: outputs[0].Address,
										Type:    txType,
										Denom:   coin.GetDenom(),
										Amount:  coin.Amount.Int64(),
									})
								}
							}
						}
					}
				} else {
					for _, output := range outputs {
						if output.Address == account {
							if len(inputs) == 1 {
								for _, coin := range output.Coins {
									txResponses = append(txResponses, types.DepositWithdrawTransaction{
										Address: inputs[0].Address,
										Type:    txType,
										Denom:   coin.GetDenom(),
										Amount:  coin.Amount.Int64(),
									})
								}
							} else if len(outputs) == 1 {
								for _, input := range inputs {
									for _, coin := range input.Coins {
										txResponses = append(txResponses, types.DepositWithdrawTransaction{
											Address: input.Address,
											Type:    txType,
											Denom:   coin.GetDenom(),
											Amount:  coin.Amount.Int64(),
										})
									}
								}
							}
						}
					}
				}
			} else if txType == "create_validator" {
				createValidatorMsg := msg.(*staking.MsgCreateValidator)

				if isWithdraw && createValidatorMsg.DelegatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: createValidatorMsg.ValidatorAddress,
						Type:    txType,
						Denom:   createValidatorMsg.Value.Denom,
						Amount:  createValidatorMsg.Value.Amount.Int64(),
					})
				} else if !isWithdraw && createValidatorMsg.ValidatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: createValidatorMsg.DelegatorAddress,
						Type:    txType,
						Denom:   createValidatorMsg.Value.Denom,
						Amount:  createValidatorMsg.Value.Amount.Int64(),
					})
				}
			} else if txType == "delegate" {
				delegateMsg := msg.(*staking.MsgDelegate)

				if isWithdraw && delegateMsg.DelegatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: delegateMsg.ValidatorAddress,
						Type:    txType,
						Denom:   delegateMsg.Amount.Denom,
						Amount:  delegateMsg.Amount.Amount.Int64(),
					})
				} else if !isWithdraw && delegateMsg.ValidatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: delegateMsg.DelegatorAddress,
						Type:    txType,
						Denom:   delegateMsg.Amount.Denom,
						Amount:  delegateMsg.Amount.Amount.Int64(),
					})
				}
			} else if txType == "begin_redelegate" {
				reDelegateMsg := msg.(*staking.MsgBeginRedelegate)

				if isWithdraw && reDelegateMsg.ValidatorSrcAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: reDelegateMsg.ValidatorDstAddress,
						Type:    txType,
						Denom:   reDelegateMsg.Amount.Denom,
						Amount:  reDelegateMsg.Amount.Amount.Int64(),
					})
				} else if !isWithdraw && reDelegateMsg.ValidatorDstAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: reDelegateMsg.ValidatorSrcAddress,
						Type:    txType,
						Denom:   reDelegateMsg.Amount.Denom,
						Amount:  reDelegateMsg.Amount.Amount.Int64(),
					})
				}
			} else if txType == "begin_unbonding" {
				unDelegateMsg := msg.(*staking.MsgUndelegate)

				if isWithdraw && unDelegateMsg.ValidatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: unDelegateMsg.DelegatorAddress,
						Type:    txType,
						Denom:   unDelegateMsg.Amount.Denom,
						Amount:  unDelegateMsg.Amount.Amount.Int64(),
					})
				} else if !isWithdraw && unDelegateMsg.DelegatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: unDelegateMsg.ValidatorAddress,
						Type:    txType,
						Denom:   unDelegateMsg.Amount.Denom,
						Amount:  unDelegateMsg.Amount.Amount.Int64(),
					})
				}
			} else if txType == "withdraw_delegator_reward" {
				var coin sdk.Coin
				if v, found := evMap["withdraw_rewards"]; found && len(v) >= 2 {
					if v[0].GetKey() == "amount" {
						coin, _ = sdk.ParseCoinNormalized(v[0].Value)
					} else if v[1].GetKey() == "amount" {
						coin, _ = sdk.ParseCoinNormalized(v[1].Value)
					}
				}

				withdrawDelegatorRewardMsg := msg.(*distribution.MsgWithdrawDelegatorReward)

				if isWithdraw && withdrawDelegatorRewardMsg.ValidatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: withdrawDelegatorRewardMsg.DelegatorAddress,
						Type:    txType,
						Denom:   coin.Denom,
						Amount:  coin.Amount.Int64(),
					})
				} else if !isWithdraw && withdrawDelegatorRewardMsg.DelegatorAddress == account {
					txResponses = append(txResponses, types.DepositWithdrawTransaction{
						Address: withdrawDelegatorRewardMsg.ValidatorAddress,
						Type:    txType,
						Denom:   coin.Denom,
						Amount:  coin.Amount.Int64(),
					})
				}
			}
		}

		if len(txResponses) == 0 {
			for _, event := range transaction.TxResult.GetEvents() {
				if event.GetType() == "transfer" {
					tx := types.TransactionTxResult{
						Address: "",
						Type:    txType,
						Tip: types.TransactionCoinSpentResult{
							Denom:  "",
							Amount: 0,
						},
					}

					attributes := event.GetAttributes()
					for _, attribute := range attributes {
						key := string(attribute.GetKey())
						value := string(attribute.GetValue())
						if key == "sender" {
							if isWithdraw {
								tx.Address = value
							}
						} else if key == "recipient" {
							if !isWithdraw {
								tx.Address = value
							}
						} else if key == "amount" {

							coin, err := sdk.ParseCoinNormalized(value)
							if err == nil {
								tx.Tip.Denom = coin.Denom
								tx.Tip.Amount = coin.Amount.Int64()
							}

						}
					}

					txResponses = append(txResponses, tx)
				}
			}
		}

		txResults = append(txResults, types.TransactionResponse{
			Time: blockTime,
			Hash: fmt.Sprintf("0x%X", transaction.Hash),
			Txs:  txResponses,
		})
	}

	res := struct {
		Transactions []types.TransactionResponse `json:"transactions"`
		TotalCount   int                         `json:"total_count"`
	}{}

	searchResult, err := SearchTxHashHandle(rpcAddr, sender, recipient, txType, 0, pageSize, -1, -1, "")
	res.TotalCount = searchResult.TotalCount
	res.Transactions = txResults

	return res, nil, http.StatusOK
}

// QueryWithdraws is a function to query all transactions.
func QueryTransactions(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-transactions] Entering transactions query")

		fmt.Println(common.RPCMethods["GET"])
		if !common.RPCMethods["GET"][config.QueryTransactions].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryTransactions].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-transactions] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = QueryBlockTransactionsHandler(rpcAddr, r)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryStatus].CachingEnabled)
	}
}

func searchUnconfirmed(rpcAddr string, limit string) (*tmTypes.ResultUnconfirmedTxs, error) {
	endpoint := fmt.Sprintf("%s/unconfirmed_txs?limit=%s", rpcAddr, limit)
	fmt.Println(endpoint)
	common.GetLogger().Info("[query-unconfirmed-txs] Entering transaction search: ", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		common.GetLogger().Error("[query-unconfirmed-txs] Unable to connect to ", endpoint)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	response := new(tmJsonRPCTypes.RPCResponse)

	if err := json.Unmarshal(respBody, response); err != nil {
		common.GetLogger().Error("[query-unconfirmed-txs] Unable to decode response: ", err)
		return nil, err
	}

	if response.Error != nil {
		common.GetLogger().Error("[query-unconfirmed-txs] Error response:", response.Error.Message)
		return nil, errors.New(response.Error.Message)
	}

	result := new(tmTypes.ResultUnconfirmedTxs)
	if err := tmjson.Unmarshal(response.Result, result); err != nil {
		common.GetLogger().Error("[query-unconfirmed-txs] Failed to unmarshal result:", err)
		return nil, fmt.Errorf("error unmarshalling result: %w", err)
	}

	return result, nil
}

func queryUnconfirmedTransactionsHandler(rpcAddr string, r *http.Request) (interface{}, interface{}, int) {
	limit := r.FormValue("limit")
	result, err := searchUnconfirmed(rpcAddr, limit)
	if err != nil {
		common.GetLogger().Error("[query-unconfirmed-txs] Failed to query unconfirmed txs: %w ", err)
		return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
	}

	response := struct {
		Count      int                                  `json:"n_txs"`
		Total      int                                  `json:"total"`
		TotalBytes int64                                `json:"total_bytes"`
		Txs        []types.TransactionUnconfirmedResult `json:"txs"`
	}{}

	response.Count = result.Count
	response.Total = result.Total
	response.TotalBytes = result.TotalBytes
	response.Txs = make([]types.TransactionUnconfirmedResult, 0)

	for _, tx := range result.Txs {
		decodedTx, err := config.EncodingCg.TxConfig.TxDecoder()(tx)
		if err != nil {
			common.GetLogger().Error("[post-unconfirmed-txs] Failed to decode transaction: ", err)
			return common.ServeError(0, "failed to decode signed TX", err.Error(), http.StatusBadRequest)
		}

		txResult, ok := decodedTx.(signing.Tx)
		if !ok {
			common.GetLogger().Error("[post-unconfirmed-txs] Failed to decode transaction")
			return common.ServeError(0, "failed to decode signed TX", err.Error(), http.StatusBadRequest)
		}

		signature, _ := txResult.GetSignaturesV2()

		var msgs []types.TxMsg = make([]types.TxMsg, 0)

		for _, msg := range txResult.GetMsgs() {
			msgs = append(msgs, types.TxMsg{
				Type: kiratypes.MsgType(msg),
				Data: msg,
			})
		}

		response.Txs = append(response.Txs, types.TransactionUnconfirmedResult{
			Msgs:      msgs,
			Fees:      txResult.GetFee(),
			Gas:       txResult.GetGas(),
			Signature: signature,
			Memo:      txResult.GetMemo(),
		})
	}

	return response, nil, http.StatusOK
}

// QueryUnconfirmedTxs is a function to query unconfirmed transactions.
func QueryUnconfirmedTxs(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Error("[query-unconfirmed-txs] Entering query")

		if !common.RPCMethods["GET"][config.QueryUnconfirmedTxs].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			response.Response, response.Error, statusCode = queryUnconfirmedTransactionsHandler(rpcAddr, r)
		}

		common.WrapResponse(w, request, *response, statusCode, false)
	}
}
