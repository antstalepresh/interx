package interx

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KiraCore/interx/tasks"
	"github.com/KiraCore/interx/test"
	"github.com/KiraCore/interx/types"
	pb "github.com/KiraCore/sekai/x/slashing/types"
	stakingTypes "github.com/KiraCore/sekai/x/staking/types"
	"github.com/stretchr/testify/suite"
	tmJsonRPCTypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"google.golang.org/grpc"
)

type ValidatorsTestSuite struct {
	suite.Suite

	dumpConsensusQuery tmJsonRPCTypes.RPCResponse
}

type kiraserver struct {
	pb.UnimplementedQueryServer
	pb.UnimplementedMsgServer
}

type ValidatorsStatus struct {
	Validators []types.QueryValidator `json:"validators,omitempty"`
	Pagination struct {
		Total int `json:"total,string,omitempty"`
	} `json:"pagination,omitempty"`
}

func (s *kiraserver) SigningInfos(ctx context.Context, req *pb.QuerySigningInfosRequest) (*pb.QuerySigningInfosResponse, error) {
	return &pb.QuerySigningInfosResponse{Validators: []stakingTypes.QueryValidator{
		{
			Address: "test_address",
		},
	}}, nil
}

func (suite *ValidatorsTestSuite) SetupTest() {
}

func (suite *ValidatorsTestSuite) TestDumpConsensusStateHandler() {
	response, _, statusCode := queryDumpConsensusStateHandler(nil, nil, test.TENDERMINT_RPC)
	suite.Require().EqualValues(response, "test")
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func (suite *ValidatorsTestSuite) TestValidatorInfosQuery() {
	r := httptest.NewRequest("GET", test.INTERX_RPC, nil)
	q := r.URL.Query()
	r.URL.RawQuery = q.Encode()

	gwCosmosmux, err := GetGrpcServeMux(*addr)
	if err != nil {
		panic("failed to serve grpc")
	}

	r.URL.Path = "/kira/slashing/v1beta1/signing_infos"
	response, _, statusCode := queryValidatorInfosHandle(r, gwCosmosmux)

	res := pb.QuerySigningInfosResponse{}
	bz, _ := json.Marshal(response)
	err = json.Unmarshal(bz, &res)
	if err != nil {
		suite.Assert()
	}

	suite.Require().EqualValues(res.Validators[0].Address, "test_address")
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func (suite *ValidatorsTestSuite) TestSnapInfoQuery() {
	tasks.AllValidators = types.AllValidators{
		Status: struct {
			ActiveValidators   int `json:"active_validators"`
			PausedValidators   int `json:"paused_validators"`
			InactiveValidators int `json:"inactive_validators"`
			JailedValidators   int `json:"jailed_validators"`
			TotalValidators    int `json:"total_validators"`
			WaitingValidators  int `json:"waiting_validators"`
		}{
			ActiveValidators: 10,
		},
		Validators: []types.QueryValidator{
			{
				Address:  "test_addr",
				Valkey:   "test_valkey",
				Pubkey:   "test_pubkey",
				Proposer: "test_proposer",
				Moniker:  "test_moniker",
				Status:   "test_status",
			},
		},
	}

	r := httptest.NewRequest("GET", test.INTERX_RPC, nil)
	q := r.URL.Query()
	q.Add("address", "test_addr")
	r.URL.RawQuery = q.Encode()
	response, _, statusCode := queryValidatorsHandle(r, nil, "")

	res := ValidatorsStatus{}
	bz, err := json.Marshal(response)
	if err != nil {
		panic("parse error")
	}

	err = json.Unmarshal(bz, &res)
	if err != nil {
		panic(err)
	}

	suite.Require().EqualValues(len(res.Validators), len(tasks.AllValidators.Validators))
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func TestValidatorsTestSuite(t *testing.T) {
	testSuite := new(ValidatorsTestSuite)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQueryServer(s, &kiraserver{})
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		_ = s.Serve(lis)
	}()

	testSuite.dumpConsensusQuery.Result, _ = json.Marshal("test")
	tmServer := http.Server{
		Addr: ":26657",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/dump_consensus_state" {
				response, _ := json.Marshal(testSuite.dumpConsensusQuery)
				w.Header().Set("Content-Type", "application/json")
				_, err := w.Write(response)
				if err != nil {
					panic(err)
				}
			}
		}),
	}
	go func() {
		_ = tmServer.ListenAndServe()
	}()

	suite.Run(t, testSuite)
	s.Stop()
	tmServer.Close()
}
