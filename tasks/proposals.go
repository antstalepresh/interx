package tasks

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/KiraCore/interx/common"
	"github.com/KiraCore/interx/types/kira/gov"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	sekaitypes "github.com/KiraCore/sekai/types"
)

var (
	AllProposals gov.AllProposals
)

const (
	// Undefined status
	VOTE_RESULT_PASSED string = "VOTE_RESULT_PASSED"
	// Active status
	VOTE_PENDING string = "VOTE_PENDING"
	// Inactive status
	VOTE_RESULT_ENACTMENT string = "VOTE_RESULT_ENACTMENT"
)

func QueryProposals(gwCosmosmux *runtime.ServeMux, gatewayAddr string) error {
	type ProposalsResponse = struct {
		Proposals  []gov.Proposal `json:"proposals,omitempty"`
		Pagination interface{}    `json:"pagination,omitempty"`
	}

	result := ProposalsResponse{}

	limit := sekaitypes.PageIterationLimit - 1
	offset := 0
	for {
		proposalsQueryRequest, _ := http.NewRequest("GET", "http://"+gatewayAddr+"/kira/gov/proposals?pagination.offset="+strconv.Itoa(offset)+"&pagination.limit="+strconv.Itoa(limit), nil)

		proposalsQueryResponse, failure, _ := common.ServeGRPC(proposalsQueryRequest, gwCosmosmux)

		if proposalsQueryResponse == nil {
			return errors.New(ToString(failure))
		}

		byteData, err := json.Marshal(proposalsQueryResponse)
		if err != nil {
			return err
		}

		subResult := ProposalsResponse{}
		err = json.Unmarshal(byteData, &subResult)
		if err != nil {
			return err
		}

		if len(subResult.Proposals) == 0 {
			break
		}

		result.Proposals = append(result.Proposals, subResult.Proposals...)

		offset += limit
	}

	allProposals := gov.AllProposals{}

	allProposals.Proposals = result.Proposals

	allProposals.Status.TotalProposals = len(result.Proposals)
	allProposals.Status.ActiveProposals = 0
	allProposals.Status.EnactingProposals = 0
	allProposals.Status.FinishedProposals = 0
	allProposals.Status.SuccessfulProposals = 0
	for _, proposal := range result.Proposals {
		if proposal.Result == VOTE_PENDING {
			allProposals.Status.ActiveProposals++
		}
		if proposal.Result == VOTE_RESULT_ENACTMENT {
			allProposals.Status.EnactingProposals++
		}
		if proposal.Result == VOTE_RESULT_PASSED {
			allProposals.Status.SuccessfulProposals++
		}
	}

	allProposals.Status.FinishedProposals = allProposals.Status.TotalProposals - allProposals.Status.ActiveProposals - allProposals.Status.EnactingProposals

	{
		request, _ := http.NewRequest("GET", "http://"+gatewayAddr+"/kira/gov/proposers_voters_count", nil)

		response, failure, _ := common.ServeGRPC(request, gwCosmosmux)

		if response == nil {
			return errors.New(ToString(failure))
		}

		byteData, err := json.Marshal(response)
		if err != nil {
			return err
		}
		result := gov.ProposalUserCount{}
		err = json.Unmarshal(byteData, &result)
		if err != nil {
			return err
		}

		allProposals.Users = result
	}

	AllProposals = allProposals

	return nil
}

func SyncProposals(gwCosmosmux *runtime.ServeMux, gatewayAddr string, isLog bool) {
	lastBlock := int64(0)
	for {
		if common.NodeStatus.Block != lastBlock {
			err := QueryProposals(gwCosmosmux, gatewayAddr)

			if err != nil && isLog {
				common.GetLogger().Error("[sync-validators] Failed to query validators: ", err)
			}

			lastBlock = common.NodeStatus.Block
		}

		time.Sleep(1 * time.Second)
	}
}
