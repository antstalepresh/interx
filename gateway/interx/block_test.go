package interx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KiraCore/interx/test"

	"github.com/stretchr/testify/suite"
)

var (
	server *httptest.Server
	ext    test.External
)

type BlockQueryTestSuite struct {
	suite.Suite
}

func (suite *BlockQueryTestSuite) SetupTest() {
}

func (suite *BlockQueryTestSuite) TestInitBlockQuery() {
}

func TestBlockQueryTestSuite(t *testing.T) {
	fmt.Println("mocking server")
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mock here
	}))

	fmt.Println("mocking external")
	ext = test.New(test.TENDERMINT_RPC, http.DefaultClient, time.Second)

	fmt.Println("run tests")

	suite.Run(t, new(BlockQueryTestSuite))
}
