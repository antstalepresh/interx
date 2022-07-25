package interx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KiraCore/interx/test"
	"github.com/KiraCore/interx/types"

	"github.com/stretchr/testify/suite"
)

type BlockQueryResult struct {
	BlockNumber int `json:"block_number"`
}

type BlockQueryTestSuite struct {
	suite.Suite

	blockQueryResponse types.RPCResponse
}

func (suite *BlockQueryTestSuite) SetupTest() {
}

func (suite *BlockQueryTestSuite) TestInitBlockQuery() {
	r := httptest.NewRequest("GET", test.INTERX_RPC, nil)
	response, error, statusCode := queryBlocksHandle(test.TENDERMINT_RPC, r)

	byteData, err := json.Marshal(response)
	if err != nil {
		suite.Assert()
	}

	result := BlockQueryResult{}
	err = json.Unmarshal(byteData, &result)
	if err != nil {
		suite.Assert()
	}

	suite.Require().EqualValues(result, suite.blockQueryResponse.Result.(BlockQueryResult))
	suite.Require().Nil(error)
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func TestBlockQueryTestSuite(t *testing.T) {
	testSuite := new(BlockQueryTestSuite)
	testSuite.blockQueryResponse.Result = BlockQueryResult{BlockNumber: 1}

	tendermintServer := http.Server{
		Addr: ":26657",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(testSuite.blockQueryResponse)
		}),
	}
	go tendermintServer.ListenAndServe()

	suite.Run(t, testSuite)

	tendermintServer.Close()
}
