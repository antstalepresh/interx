package database

import (
	"github.com/KiraCore/interx/config"
	"github.com/sonyarouje/simdb/db"
	tmTypes "github.com/tendermint/tendermint/rpc/core/types"
)

// TransactionData is a struct for transaction details.
type TransactionData struct {
	Address string                 `json:"address"`
	Data    tmTypes.ResultTxSearch `json:"data"`
}

// ID is a field for facuet claim struct.
func (c TransactionData) ID() (jsonField string, value interface{}) {
	value = c.Address
	jsonField = "address"
	return
}

func LoadTransactionDbDriver() {
	DisableStdout()
	driver, _ := db.New(config.GetDbCacheDir() + "/transaction")
	EnableStdout()

	transactionDb = driver
}

// GetTransactions is a function to get user transactions from cache
func GetTransactions(address string) (tmTypes.ResultTxSearch, error) {
	if transactionDb == nil {
		panic("cache dir not set")
	}

	data := TransactionData{}

	DisableStdout()
	err := transactionDb.Open(TransactionData{}).Where("address", "=", address).First().AsEntity(&data)
	EnableStdout()

	if err != nil {
		return tmTypes.ResultTxSearch{}, err
	}

	return data.Data, nil
}

// SaveTransactions is a function to save user transactions to cache
func SaveTransactions(address string, txsData tmTypes.ResultTxSearch) {
	if transactionDb == nil {
		panic("cache dir not set")
	}

	data := TransactionData{
		Address: address,
		Data:    txsData,
	}

	_, err := GetTransactions(address)

	if err != nil {
		dataToUpdate := TransactionData{}
		DisableStdout()
		err := transactionDb.Open(TransactionData{}).Where("address", "=", address).First().AsEntity(&dataToUpdate)
		if err != nil {
			err = transactionDb.Open(TransactionData{}).Insert(data)
		} else {
			dataToUpdate.Data = txsData
			transactionDb.Update(dataToUpdate)
		}
		EnableStdout()

		if err != nil {
			panic(err)
		}

	}
}

var (
	transactionDb *db.Driver
)
