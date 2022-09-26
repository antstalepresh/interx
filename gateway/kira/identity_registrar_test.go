package kira

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/KiraCore/interx/test"
	pb "github.com/KiraCore/sekai/x/gov/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

func (s *server) IdentityRecordsByAddress(ctx context.Context, in *pb.QueryIdentityRecordsByAddressRequest) (*pb.QueryIdentityRecordsByAddressResponse, error) {
	return &pb.QueryIdentityRecordsByAddressResponse{Records: []pb.IdentityRecord{
		{
			Address: "test_address",
		},
	}}, nil
}

func (s *server) AllIdentityRecords(ctx context.Context, in *pb.QueryAllIdentityRecordsRequest) (*pb.QueryAllIdentityRecordsResponse, error) {
	return &pb.QueryAllIdentityRecordsResponse{Records: []pb.IdentityRecord{
		{
			Address: "test_address",
		},
	}}, nil
}

func (s *server) IdentityRecordVerifyRequestsByRequester(ctx context.Context, in *pb.QueryIdentityRecordVerifyRequestsByRequester) (*pb.QueryIdentityRecordVerifyRequestsByRequesterResponse, error) {
	return &pb.QueryIdentityRecordVerifyRequestsByRequesterResponse{VerifyRecords: []pb.IdentityRecordsVerify{
		{
			Address: "test_address",
		},
	}}, nil
}

type tempIDRecord struct {
	Records []struct {
		Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
		Address   string    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
		Key       string    `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
		Value     string    `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
		Date      time.Time `protobuf:"bytes,5,opt,name=date,proto3,stdtime" json:"date"`
		Verifiers []string  `protobuf:"bytes,6,rep,name=verifiers,proto3" json:"verifiers,omitempty"`
	} `json:"records"`
}

type tempIDRecordByRequester struct {
	VerifyRecords []struct {
		Id       string `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
		Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
		Verifier string `protobuf:"bytes,3,opt,name=verifier,proto3" json:"verifier,omitempty"`
	} `json:"verify_records"`
}

type IdentityRecordsTestSuite struct {
	suite.Suite
}

func (suite *IdentityRecordsTestSuite) SetupTest() {
}

func (suite *IdentityRecordsTestSuite) TestAllIdentityHandler() {
	r := httptest.NewRequest("GET", test.INTERX_RPC, nil)
	query := r.URL.Query()
	query.Add("key", "test")
	query.Add("offset", "10")
	query.Add("limit", "10")
	query.Add("count_total", "100")

	gwCosmosmux, err := GetGrpcServeMux(*addr)
	if err != nil {
		panic("failed to serve grpc")
	}
	r.URL.Path = "/kira/gov/all_identity_records"
	response, _, statusCode := QueryAllIdentityRecordsHandler(r, gwCosmosmux)

	res := tempIDRecord{}
	bz, err := json.Marshal(response)
	err = json.Unmarshal(bz, &res)

	suite.Require().EqualValues(res.Records[0].Address, "test_address")
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func (suite *IdentityRecordsTestSuite) TestExecutionFeeHandler() {
	r := httptest.NewRequest("GET", test.INTERX_RPC, nil)
	query := r.URL.Query()

	query.Add("creator", "test")
	r.URL.RawQuery = query.Encode()

	gwCosmosmux, err := GetGrpcServeMux(*addr)
	if err != nil {
		panic("failed to serve grpc")
	}
	response, _, statusCode := QueryIdentityRecordsByAddressHandler(r, gwCosmosmux)

	res := tempIDRecord{}
	bz, err := json.Marshal(response)
	err = json.Unmarshal(bz, &res)

	suite.Require().EqualValues(res.Records[0].Address, "test_address")
	suite.Require().EqualValues(statusCode, http.StatusOK)
}

func (suite *IdentityRecordsTestSuite) TestIdentityRecordVerifyRequestsByRequesterHandler() {
	data := url.Values{}
	data.Set("requester", "cosmos18x8js8kfyrlmqtnzcqzfjs3qhackxep5ww4nx7")

	r := httptest.NewRequest("GET", test.INTERX_RPC, strings.NewReader(data.Encode()))
	query := r.URL.Query()
	query.Add("key", "test")
	query.Add("offset", "10")
	query.Add("limit", "10")
	query.Add("count_total", "100")
	r.URL.RawQuery = query.Encode()

	gwCosmosmux, err := GetGrpcServeMux(*addr)
	if err != nil {
		panic("failed to serve grpc")
	}
	response, _, _ := QueryIdentityRecordVerifyRequestsByRequesterHandler(r, gwCosmosmux)

	res := tempIDRecordByRequester{}
	bz, err := json.Marshal(response)
	err = json.Unmarshal(bz, &res)
}

func TestIdentityRecordsTestSuite(t *testing.T) {
	testSuite := new(IdentityRecordsTestSuite)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQueryServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	go s.Serve(lis)

	suite.Run(t, testSuite)
	s.Stop()
}
