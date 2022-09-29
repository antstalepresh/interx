package interx

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/KiraCore/interx/config"
	"github.com/KiraCore/interx/test"
	"github.com/stretchr/testify/suite"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmRPCTypes "github.com/tendermint/tendermint/rpc/core/types"
	tmJsonRPCTypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	tmTypes "github.com/tendermint/tendermint/types"
)

type GenesisChecksumResponse struct {
	Checksum string `json:"checksum,omitempty"`
}

type GenesisQueryTestSuite struct {
	suite.Suite

	genesisQueryResponse tmJsonRPCTypes.RPCResponse
}

func (suite *GenesisQueryTestSuite) SetupTest() {
}

func (suite *GenesisQueryTestSuite) TestQueryGenesisSumHandler() {
	config.Config.Cache.CacheDir = "./"
	os.Mkdir("./reference", 0777)
	os.Create("./reference/genesis.json")
	response, _, statusCode := queryGenesisSumHandler(test.TENDERMINT_RPC)

	byteData, err := json.Marshal(response)
	if err != nil {
		suite.Assert()
	}

	result := GenesisChecksumResponse{}
	err = json.Unmarshal(byteData, &result)
	if err != nil {
		suite.Assert()
	}

	suite.Require().NoError(err)
	suite.Require().EqualValues(statusCode, http.StatusOK)
	os.RemoveAll("./reference")
}

func TestGenesisQueryTestSuite(t *testing.T) {
	testSuite := new(GenesisQueryTestSuite)
	resBytes, err := tmjson.Marshal(tmRPCTypes.ResultGenesis{
		Genesis: &tmTypes.GenesisDoc{
			GenesisTime:   time.Now(),
			ChainID:       "test",
			InitialHeight: 1,
		},
	})

	if err != nil {
		panic(err)
	}

	testSuite.genesisQueryResponse.Result = resBytes

	tendermintServer := http.Server{
		Addr: ":26657",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/genesis" {
				response, _ := tmjson.Marshal(testSuite.genesisQueryResponse)
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}),
	}
	go tendermintServer.ListenAndServe()

	suite.Run(t, testSuite)

	tendermintServer.Close()
}
