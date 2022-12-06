package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/KiraCore/interx/config"
	"github.com/KiraCore/interx/global"
	tmTypes "github.com/tendermint/tendermint/rpc/core/types"
)

// TransactionData is a struct for transaction details.
type TransactionData struct {
	Address string                 `json:"address"`
	Data    tmTypes.ResultTxSearch `json:"data"`
}

// GetTransactions is a function to get user transactions from cache
func GetTransactions(address string) (*tmTypes.ResultTxSearch, error) {

	filePath := fmt.Sprintf("%s/transactions/%s", config.GetDbCacheDir(), address)
	data := tmTypes.ResultTxSearch{}

	txs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &tmTypes.ResultTxSearch{}, err
	}

	err = json.Unmarshal([]byte(txs), &data)

	if err != nil {
		return &tmTypes.ResultTxSearch{}, err
	}

	return &data, nil
}

func GetLastBlockFetched(address string) int64 {
	data, err := GetTransactions(address)

	if err != nil {
		return 0
	}

	lastTx := data.Txs[len(data.Txs)-1]
	return lastTx.Height
}

// SaveTransactions is a function to save user transactions to cache
func SaveTransactions(address string, txsData tmTypes.ResultTxSearch) error {
	cachedData, err := GetTransactions(address)

	if cachedData.TotalCount > 0 {
		txsData.Txs = append(cachedData.Txs, txsData.Txs...)
		txsData.TotalCount = txsData.TotalCount + cachedData.TotalCount
	}

	data, err := json.Marshal(txsData)
	if err != nil {
		return err
	}

	folderPath := fmt.Sprintf("%s/transactions", config.GetDbCacheDir())
	filePath := fmt.Sprintf("%s/%s", folderPath, address)

	global.Mutex.Lock()
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		global.Mutex.Unlock()

		fmt.Println("[cache] Unable to create a folder: ", folderPath)
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	global.Mutex.Unlock()

	if err != nil {
		fmt.Println("[cache] Unable to save response: ", filePath)
	}

	return err
}
