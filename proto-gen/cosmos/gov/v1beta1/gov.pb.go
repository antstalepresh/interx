// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cosmos/gov/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// VoteOption enumerates the valid vote options for a given governance proposal.
type VoteOption int32

const (
	// VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
	VoteOption_VOTE_OPTION_UNSPECIFIED VoteOption = 0
	// VOTE_OPTION_YES defines a yes vote option.
	VoteOption_VOTE_OPTION_YES VoteOption = 1
	// VOTE_OPTION_ABSTAIN defines an abstain vote option.
	VoteOption_VOTE_OPTION_ABSTAIN VoteOption = 2
	// VOTE_OPTION_NO defines a no vote option.
	VoteOption_VOTE_OPTION_NO VoteOption = 3
	// VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	VoteOption_VOTE_OPTION_NO_WITH_VETO VoteOption = 4
)

var VoteOption_name = map[int32]string{
	0: "VOTE_OPTION_UNSPECIFIED",
	1: "VOTE_OPTION_YES",
	2: "VOTE_OPTION_ABSTAIN",
	3: "VOTE_OPTION_NO",
	4: "VOTE_OPTION_NO_WITH_VETO",
}

var VoteOption_value = map[string]int32{
	"VOTE_OPTION_UNSPECIFIED":  0,
	"VOTE_OPTION_YES":          1,
	"VOTE_OPTION_ABSTAIN":      2,
	"VOTE_OPTION_NO":           3,
	"VOTE_OPTION_NO_WITH_VETO": 4,
}

func (x VoteOption) String() string {
	return proto.EnumName(VoteOption_name, int32(x))
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{0}
}

// ProposalStatus enumerates the valid statuses of a proposal.
type ProposalStatus int32

const (
	// PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status.
	ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED ProposalStatus = 0
	// PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
	// period.
	ProposalStatus_PROPOSAL_STATUS_DEPOSIT_PERIOD ProposalStatus = 1
	// PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
	// period.
	ProposalStatus_PROPOSAL_STATUS_VOTING_PERIOD ProposalStatus = 2
	// PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
	// passed.
	ProposalStatus_PROPOSAL_STATUS_PASSED ProposalStatus = 3
	// PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
	// been rejected.
	ProposalStatus_PROPOSAL_STATUS_REJECTED ProposalStatus = 4
	// PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
	// failed.
	ProposalStatus_PROPOSAL_STATUS_FAILED ProposalStatus = 5
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_UNSPECIFIED",
	1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	2: "PROPOSAL_STATUS_VOTING_PERIOD",
	3: "PROPOSAL_STATUS_PASSED",
	4: "PROPOSAL_STATUS_REJECTED",
	5: "PROPOSAL_STATUS_FAILED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_UNSPECIFIED":    0,
	"PROPOSAL_STATUS_DEPOSIT_PERIOD": 1,
	"PROPOSAL_STATUS_VOTING_PERIOD":  2,
	"PROPOSAL_STATUS_PASSED":         3,
	"PROPOSAL_STATUS_REJECTED":       4,
	"PROPOSAL_STATUS_FAILED":         5,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{1}
}

// WeightedVoteOption defines a unit of vote for vote split.
//
// Since: cosmos-sdk 0.43
type WeightedVoteOption struct {
	Option               VoteOption `protobuf:"varint,1,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`
	Weight               string     `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WeightedVoteOption) Reset()         { *m = WeightedVoteOption{} }
func (m *WeightedVoteOption) String() string { return proto.CompactTextString(m) }
func (*WeightedVoteOption) ProtoMessage()    {}
func (*WeightedVoteOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{0}
}

func (m *WeightedVoteOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedVoteOption.Unmarshal(m, b)
}
func (m *WeightedVoteOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedVoteOption.Marshal(b, m, deterministic)
}
func (m *WeightedVoteOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedVoteOption.Merge(m, src)
}
func (m *WeightedVoteOption) XXX_Size() int {
	return xxx_messageInfo_WeightedVoteOption.Size(m)
}
func (m *WeightedVoteOption) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedVoteOption.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedVoteOption proto.InternalMessageInfo

func (m *WeightedVoteOption) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

func (m *WeightedVoteOption) GetWeight() string {
	if m != nil {
		return m.Weight
	}
	return ""
}

// TextProposal defines a standard text proposal whose changes need to be
// manually updated in case of approval.
type TextProposal struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextProposal) Reset()         { *m = TextProposal{} }
func (m *TextProposal) String() string { return proto.CompactTextString(m) }
func (*TextProposal) ProtoMessage()    {}
func (*TextProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{1}
}

func (m *TextProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextProposal.Unmarshal(m, b)
}
func (m *TextProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextProposal.Marshal(b, m, deterministic)
}
func (m *TextProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextProposal.Merge(m, src)
}
func (m *TextProposal) XXX_Size() int {
	return xxx_messageInfo_TextProposal.Size(m)
}
func (m *TextProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TextProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TextProposal proto.InternalMessageInfo

func (m *TextProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TextProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Deposit defines an amount deposited by an account address to an active
// proposal.
type Deposit struct {
	ProposalId           uint64        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Depositor            string        `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount               []*types.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{2}
}

func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deposit.Unmarshal(m, b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return xxx_messageInfo_Deposit.Size(m)
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Deposit) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Deposit) GetAmount() []*types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

// Proposal defines the core field members of a governance proposal.
type Proposal struct {
	ProposalId           uint64                 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Content              *anypb.Any             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status               ProposalStatus         `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"status,omitempty"`
	FinalTallyResult     *TallyResult           `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result,omitempty"`
	SubmitTime           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	DepositEndTime       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3" json:"deposit_end_time,omitempty"`
	TotalDeposit         []*types.Coin          `protobuf:"bytes,7,rep,name=total_deposit,json=totalDeposit,proto3" json:"total_deposit,omitempty"`
	VotingStartTime      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3" json:"voting_start_time,omitempty"`
	VotingEndTime        *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3" json:"voting_end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{3}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Proposal) GetContent() *anypb.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Proposal) GetStatus() ProposalStatus {
	if m != nil {
		return m.Status
	}
	return ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED
}

func (m *Proposal) GetFinalTallyResult() *TallyResult {
	if m != nil {
		return m.FinalTallyResult
	}
	return nil
}

func (m *Proposal) GetSubmitTime() *timestamppb.Timestamp {
	if m != nil {
		return m.SubmitTime
	}
	return nil
}

func (m *Proposal) GetDepositEndTime() *timestamppb.Timestamp {
	if m != nil {
		return m.DepositEndTime
	}
	return nil
}

func (m *Proposal) GetTotalDeposit() []*types.Coin {
	if m != nil {
		return m.TotalDeposit
	}
	return nil
}

func (m *Proposal) GetVotingStartTime() *timestamppb.Timestamp {
	if m != nil {
		return m.VotingStartTime
	}
	return nil
}

func (m *Proposal) GetVotingEndTime() *timestamppb.Timestamp {
	if m != nil {
		return m.VotingEndTime
	}
	return nil
}

// TallyResult defines a standard tally for a governance proposal.
type TallyResult struct {
	Yes                  string   `protobuf:"bytes,1,opt,name=yes,proto3" json:"yes,omitempty"`
	Abstain              string   `protobuf:"bytes,2,opt,name=abstain,proto3" json:"abstain,omitempty"`
	No                   string   `protobuf:"bytes,3,opt,name=no,proto3" json:"no,omitempty"`
	NoWithVeto           string   `protobuf:"bytes,4,opt,name=no_with_veto,json=noWithVeto,proto3" json:"no_with_veto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TallyResult) Reset()         { *m = TallyResult{} }
func (m *TallyResult) String() string { return proto.CompactTextString(m) }
func (*TallyResult) ProtoMessage()    {}
func (*TallyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{4}
}

func (m *TallyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TallyResult.Unmarshal(m, b)
}
func (m *TallyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TallyResult.Marshal(b, m, deterministic)
}
func (m *TallyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyResult.Merge(m, src)
}
func (m *TallyResult) XXX_Size() int {
	return xxx_messageInfo_TallyResult.Size(m)
}
func (m *TallyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyResult.DiscardUnknown(m)
}

var xxx_messageInfo_TallyResult proto.InternalMessageInfo

func (m *TallyResult) GetYes() string {
	if m != nil {
		return m.Yes
	}
	return ""
}

func (m *TallyResult) GetAbstain() string {
	if m != nil {
		return m.Abstain
	}
	return ""
}

func (m *TallyResult) GetNo() string {
	if m != nil {
		return m.No
	}
	return ""
}

func (m *TallyResult) GetNoWithVeto() string {
	if m != nil {
		return m.NoWithVeto
	}
	return ""
}

// Vote defines a vote on a governance proposal.
// A Vote consists of a proposal ID, the voter, and the vote option.
type Vote struct {
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// Deprecated: Prefer to use `options` instead. This field is set in queries
	// if and only if `len(options) == 1` and that option has weight 1. In all
	// other cases, this field will default to VOTE_OPTION_UNSPECIFIED.
	Option VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"` // Deprecated: Do not use.
	// Since: cosmos-sdk 0.43
	Options              []*WeightedVoteOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{5}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Vote) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

// Deprecated: Do not use.
func (m *Vote) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

func (m *Vote) GetOptions() []*WeightedVoteOption {
	if m != nil {
		return m.Options
	}
	return nil
}

// DepositParams defines the params for deposits on governance proposals.
type DepositParams struct {
	//  Minimum deposit for a proposal to enter voting period.
	MinDeposit []*types.Coin `protobuf:"bytes,1,rep,name=min_deposit,json=minDeposit,proto3" json:"min_deposit,omitempty"`
	//  Maximum period for Atom holders to deposit on a proposal. Initial value: 2
	//  months.
	MaxDepositPeriod     *durationpb.Duration `protobuf:"bytes,2,opt,name=max_deposit_period,json=maxDepositPeriod,proto3" json:"max_deposit_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DepositParams) Reset()         { *m = DepositParams{} }
func (m *DepositParams) String() string { return proto.CompactTextString(m) }
func (*DepositParams) ProtoMessage()    {}
func (*DepositParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{6}
}

func (m *DepositParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepositParams.Unmarshal(m, b)
}
func (m *DepositParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepositParams.Marshal(b, m, deterministic)
}
func (m *DepositParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositParams.Merge(m, src)
}
func (m *DepositParams) XXX_Size() int {
	return xxx_messageInfo_DepositParams.Size(m)
}
func (m *DepositParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositParams.DiscardUnknown(m)
}

var xxx_messageInfo_DepositParams proto.InternalMessageInfo

func (m *DepositParams) GetMinDeposit() []*types.Coin {
	if m != nil {
		return m.MinDeposit
	}
	return nil
}

func (m *DepositParams) GetMaxDepositPeriod() *durationpb.Duration {
	if m != nil {
		return m.MaxDepositPeriod
	}
	return nil
}

// VotingParams defines the params for voting on governance proposals.
type VotingParams struct {
	//  Length of the voting period.
	VotingPeriod         *durationpb.Duration `protobuf:"bytes,1,opt,name=voting_period,json=votingPeriod,proto3" json:"voting_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VotingParams) Reset()         { *m = VotingParams{} }
func (m *VotingParams) String() string { return proto.CompactTextString(m) }
func (*VotingParams) ProtoMessage()    {}
func (*VotingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{7}
}

func (m *VotingParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VotingParams.Unmarshal(m, b)
}
func (m *VotingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VotingParams.Marshal(b, m, deterministic)
}
func (m *VotingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingParams.Merge(m, src)
}
func (m *VotingParams) XXX_Size() int {
	return xxx_messageInfo_VotingParams.Size(m)
}
func (m *VotingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingParams.DiscardUnknown(m)
}

var xxx_messageInfo_VotingParams proto.InternalMessageInfo

func (m *VotingParams) GetVotingPeriod() *durationpb.Duration {
	if m != nil {
		return m.VotingPeriod
	}
	return nil
}

// TallyParams defines the params for tallying votes on governance proposals.
type TallyParams struct {
	//  Minimum percentage of total stake needed to vote for a result to be
	//  considered valid.
	Quorum []byte `protobuf:"bytes,1,opt,name=quorum,proto3" json:"quorum,omitempty"`
	//  Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
	Threshold []byte `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	//  Minimum value of Veto votes to Total votes ratio for proposal to be
	//  vetoed. Default value: 1/3.
	VetoThreshold        []byte   `protobuf:"bytes,3,opt,name=veto_threshold,json=vetoThreshold,proto3" json:"veto_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TallyParams) Reset()         { *m = TallyParams{} }
func (m *TallyParams) String() string { return proto.CompactTextString(m) }
func (*TallyParams) ProtoMessage()    {}
func (*TallyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{8}
}

func (m *TallyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TallyParams.Unmarshal(m, b)
}
func (m *TallyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TallyParams.Marshal(b, m, deterministic)
}
func (m *TallyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyParams.Merge(m, src)
}
func (m *TallyParams) XXX_Size() int {
	return xxx_messageInfo_TallyParams.Size(m)
}
func (m *TallyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyParams.DiscardUnknown(m)
}

var xxx_messageInfo_TallyParams proto.InternalMessageInfo

func (m *TallyParams) GetQuorum() []byte {
	if m != nil {
		return m.Quorum
	}
	return nil
}

func (m *TallyParams) GetThreshold() []byte {
	if m != nil {
		return m.Threshold
	}
	return nil
}

func (m *TallyParams) GetVetoThreshold() []byte {
	if m != nil {
		return m.VetoThreshold
	}
	return nil
}

func init() {
	proto.RegisterEnum("cosmos.gov.v1beta1.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterEnum("cosmos.gov.v1beta1.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterType((*WeightedVoteOption)(nil), "cosmos.gov.v1beta1.WeightedVoteOption")
	proto.RegisterType((*TextProposal)(nil), "cosmos.gov.v1beta1.TextProposal")
	proto.RegisterType((*Deposit)(nil), "cosmos.gov.v1beta1.Deposit")
	proto.RegisterType((*Proposal)(nil), "cosmos.gov.v1beta1.Proposal")
	proto.RegisterType((*TallyResult)(nil), "cosmos.gov.v1beta1.TallyResult")
	proto.RegisterType((*Vote)(nil), "cosmos.gov.v1beta1.Vote")
	proto.RegisterType((*DepositParams)(nil), "cosmos.gov.v1beta1.DepositParams")
	proto.RegisterType((*VotingParams)(nil), "cosmos.gov.v1beta1.VotingParams")
	proto.RegisterType((*TallyParams)(nil), "cosmos.gov.v1beta1.TallyParams")
}

func init() { proto.RegisterFile("cosmos/gov/v1beta1/gov.proto", fileDescriptor_6e82113c1a9a4b7c) }

var fileDescriptor_6e82113c1a9a4b7c = []byte{
	// 1430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x51, 0x68, 0xdb, 0x56,
	0x17, 0xb6, 0x6c, 0xc7, 0x89, 0xaf, 0x9d, 0x44, 0xbd, 0x49, 0x13, 0x47, 0x7f, 0x7f, 0x4b, 0xbf,
	0xfe, 0x9f, 0x12, 0x4a, 0xeb, 0xb4, 0xf9, 0x47, 0xc7, 0x52, 0xd8, 0x66, 0xc5, 0xca, 0xea, 0x51,
	0x62, 0x23, 0xab, 0x0e, 0xed, 0x1e, 0x84, 0x62, 0xdf, 0x3a, 0xda, 0x2c, 0x5d, 0xcf, 0xba, 0x4e,
	0x63, 0xf6, 0xb2, 0xc7, 0xe2, 0xc1, 0xd8, 0x63, 0x61, 0x18, 0x0a, 0x63, 0x2f, 0x7b, 0xde, 0xf3,
	0x9e, 0xcb, 0x18, 0x0c, 0xf6, 0x34, 0x36, 0x70, 0xe9, 0x06, 0xa3, 0xe4, 0x31, 0x0f, 0x7b, 0x1e,
	0xd2, 0xbd, 0xb2, 0x65, 0x3b, 0x2c, 0xf5, 0x9e, 0x22, 0x9d, 0x7b, 0xbe, 0xef, 0x9c, 0xf3, 0xf9,
	0x9c, 0x73, 0x15, 0x70, 0xa5, 0x86, 0x5d, 0x1b, 0xbb, 0x5b, 0x0d, 0x7c, 0xbc, 0x75, 0x7c, 0xeb,
	0x10, 0x11, 0xf3, 0x96, 0xf7, 0x9c, 0x6b, 0xb5, 0x31, 0xc1, 0x10, 0xd2, 0xd3, 0x9c, 0x67, 0x61,
	0xa7, 0x42, 0x96, 0x21, 0x0e, 0x4d, 0x17, 0x0d, 0x21, 0x35, 0x6c, 0x39, 0x14, 0x23, 0xac, 0x36,
	0x70, 0x03, 0xfb, 0x8f, 0x5b, 0xde, 0x13, 0xb3, 0x6e, 0x50, 0x94, 0x41, 0x0f, 0x18, 0x2d, 0x3d,
	0x12, 0x1b, 0x18, 0x37, 0x9a, 0x68, 0xcb, 0x7f, 0x3b, 0xec, 0x3c, 0xda, 0x22, 0x96, 0x8d, 0x5c,
	0x62, 0xda, 0xad, 0x00, 0x3b, 0xe9, 0x60, 0x3a, 0x5d, 0x76, 0x94, 0x9d, 0x3c, 0xaa, 0x77, 0xda,
	0x26, 0xb1, 0x30, 0x4b, 0x46, 0xfe, 0x9a, 0x03, 0xf0, 0x00, 0x59, 0x8d, 0x23, 0x82, 0xea, 0x55,
	0x4c, 0x50, 0xa9, 0xe5, 0x1d, 0xc2, 0xdb, 0x20, 0x81, 0xfd, 0xa7, 0x0c, 0x27, 0x71, 0x9b, 0x4b,
	0xdb, 0xd9, 0xdc, 0x74, 0xa1, 0xb9, 0x91, 0xbf, 0xc6, 0xbc, 0xe1, 0x01, 0x48, 0x3c, 0xf6, 0xd9,
	0x32, 0x51, 0x89, 0xdb, 0x4c, 0x2a, 0xef, 0x3c, 0x1f, 0x88, 0x91, 0x5f, 0x06, 0xe2, 0xd5, 0x86,
	0x45, 0x8e, 0x3a, 0x87, 0xb9, 0x1a, 0xb6, 0x59, 0x6d, 0xec, 0xcf, 0x0d, 0xb7, 0xfe, 0xd1, 0x16,
	0xe9, 0xb6, 0x90, 0x9b, 0x2b, 0xa0, 0xda, 0xd9, 0x40, 0x5c, 0xec, 0x9a, 0x76, 0x73, 0x47, 0xa6,
	0x2c, 0xb2, 0xc6, 0xe8, 0xe4, 0x03, 0x90, 0xd6, 0xd1, 0x09, 0x29, 0xb7, 0x71, 0x0b, 0xbb, 0x66,
	0x13, 0xae, 0x82, 0x39, 0x62, 0x91, 0x26, 0xf2, 0xf3, 0x4b, 0x6a, 0xf4, 0x05, 0x4a, 0x20, 0x55,
	0x47, 0x6e, 0xad, 0x6d, 0xd1, 0xdc, 0xfd, 0x1c, 0xb4, 0xb0, 0x69, 0x67, 0xf9, 0xd5, 0x33, 0x91,
	0xfb, 0xe9, 0xdb, 0x1b, 0xf3, 0xbb, 0xd8, 0x21, 0xc8, 0x21, 0xf2, 0x8f, 0x1c, 0x98, 0x2f, 0xa0,
	0x16, 0x76, 0x2d, 0x02, 0xdf, 0x04, 0xa9, 0x16, 0x0b, 0x60, 0x58, 0x75, 0x9f, 0x3a, 0xae, 0xac,
	0x9d, 0x0d, 0x44, 0x48, 0x93, 0x0a, 0x1d, 0xca, 0x1a, 0x08, 0xde, 0x8a, 0x75, 0x78, 0x05, 0x24,
	0xeb, 0x94, 0x03, 0xb7, 0x59, 0xd4, 0x91, 0x01, 0xd6, 0x40, 0xc2, 0xb4, 0x71, 0xc7, 0x21, 0x99,
	0x98, 0x14, 0xdb, 0x4c, 0x6d, 0x6f, 0x04, 0x62, 0x7a, 0x1d, 0x32, 0x54, 0x73, 0x17, 0x5b, 0x8e,
	0x72, 0xd3, 0xd3, 0xeb, 0x9b, 0x17, 0xe2, 0xe6, 0x6b, 0xe8, 0xe5, 0x01, 0x5c, 0x8d, 0x51, 0xef,
	0x2c, 0x3c, 0x79, 0x26, 0x46, 0x5e, 0x3d, 0x13, 0x23, 0xf2, 0x9f, 0x09, 0xb0, 0x30, 0xd4, 0xe9,
	0x8d, 0xf3, 0x4a, 0x5a, 0x39, 0x1d, 0x88, 0x51, 0xab, 0x7e, 0x36, 0x10, 0x93, 0xb4, 0xb0, 0xc9,
	0x7a, 0xee, 0x80, 0xf9, 0x1a, 0xd5, 0xc7, 0xaf, 0x26, 0xb5, 0xbd, 0x9a, 0xa3, 0x7d, 0x94, 0x0b,
	0xfa, 0x28, 0x97, 0x77, 0xba, 0x4a, 0xea, 0xfb, 0x91, 0x90, 0x5a, 0x80, 0x80, 0x55, 0x90, 0x70,
	0x89, 0x49, 0x3a, 0x6e, 0x26, 0xe6, 0xf7, 0x8e, 0x7c, 0x5e, 0xef, 0x04, 0x09, 0x56, 0x7c, 0x4f,
	0x45, 0x38, 0x1b, 0x88, 0x6b, 0x13, 0x22, 0x53, 0x12, 0x59, 0x63, 0x6c, 0xb0, 0x05, 0xe0, 0x23,
	0xcb, 0x31, 0x9b, 0x06, 0x31, 0x9b, 0xcd, 0xae, 0xd1, 0x46, 0x6e, 0xa7, 0x49, 0x32, 0x71, 0x3f,
	0x3f, 0xf1, 0xbc, 0x18, 0xba, 0xe7, 0xa7, 0xf9, 0x6e, 0xca, 0x7f, 0x3c, 0x61, 0xcf, 0x06, 0xe2,
	0x06, 0x0d, 0x32, 0x4d, 0x24, 0x6b, 0xbc, 0x6f, 0x0c, 0x81, 0xe0, 0x07, 0x20, 0xe5, 0x76, 0x0e,
	0x6d, 0x8b, 0x18, 0xde, 0xc4, 0x65, 0xe6, 0xfc, 0x50, 0xc2, 0x94, 0x14, 0x7a, 0x30, 0x8e, 0x4a,
	0x96, 0x45, 0x61, 0xfd, 0x12, 0x02, 0xcb, 0x5f, 0xbc, 0x10, 0x39, 0x0d, 0x50, 0x8b, 0x07, 0x80,
	0x16, 0xe0, 0x59, 0x8b, 0x18, 0xc8, 0xa9, 0xd3, 0x08, 0x89, 0x0b, 0x23, 0xfc, 0x97, 0x45, 0x58,
	0xa7, 0x11, 0x26, 0x19, 0x68, 0x98, 0x25, 0x66, 0x56, 0x9d, 0xba, 0x1f, 0xea, 0x09, 0x07, 0x16,
	0x09, 0x26, 0x66, 0xd3, 0x60, 0x07, 0x99, 0xf9, 0x8b, 0x1a, 0xf1, 0x2e, 0x8b, 0xb3, 0x4a, 0xe3,
	0x8c, 0xa1, 0xe5, 0x99, 0x1a, 0x34, 0xed, 0x63, 0x83, 0x11, 0x6b, 0x82, 0x4b, 0xc7, 0x98, 0x58,
	0x4e, 0xc3, 0xfb, 0x79, 0xdb, 0x4c, 0xd8, 0x85, 0x0b, 0xcb, 0xfe, 0x1f, 0x4b, 0x27, 0x43, 0xd3,
	0x99, 0xa2, 0xa0, 0x75, 0x2f, 0x53, 0x7b, 0xc5, 0x33, 0xfb, 0x85, 0x3f, 0x02, 0xcc, 0x34, 0x92,
	0x38, 0x79, 0x61, 0x2c, 0x99, 0xc5, 0x5a, 0x1b, 0x8b, 0x35, 0xae, 0xf0, 0x22, 0xb5, 0x32, 0x81,
	0x77, 0xe2, 0xde, 0x56, 0x91, 0x9f, 0x47, 0x41, 0x2a, 0xdc, 0x3e, 0xef, 0x82, 0x58, 0x17, 0xb9,
	0x74, 0x43, 0x29, 0xb9, 0x19, 0x36, 0x61, 0xd1, 0x21, 0x9a, 0x07, 0x85, 0x77, 0xc1, 0xbc, 0x79,
	0xe8, 0x12, 0xd3, 0x62, 0xbb, 0x6c, 0x66, 0x96, 0x00, 0x0e, 0xdf, 0x06, 0x51, 0x07, 0xfb, 0x03,
	0x39, 0x3b, 0x49, 0xd4, 0xc1, 0xb0, 0x01, 0xd2, 0x0e, 0x36, 0x1e, 0x5b, 0xe4, 0xc8, 0x38, 0x46,
	0x04, 0xfb, 0x63, 0x97, 0x54, 0xd4, 0xd9, 0x98, 0xce, 0x06, 0xe2, 0x0a, 0x15, 0x35, 0xcc, 0x25,
	0x6b, 0xc0, 0xc1, 0x07, 0x16, 0x39, 0xaa, 0x22, 0x82, 0x99, 0x94, 0xbf, 0x73, 0x20, 0xee, 0x5d,
	0x2f, 0xff, 0x7c, 0x25, 0xaf, 0x82, 0xb9, 0x63, 0x4c, 0x50, 0xb0, 0x8e, 0xe9, 0x0b, 0xdc, 0x19,
	0xde, 0x6b, 0xb1, 0xd7, 0xb9, 0xd7, 0x94, 0x68, 0x86, 0x1b, 0xde, 0x6d, 0x7b, 0x60, 0x9e, 0x3e,
	0xb9, 0x99, 0xb8, 0x3f, 0x3e, 0x57, 0xcf, 0x03, 0x4f, 0x5f, 0xa6, 0x4a, 0xdc, 0x53, 0x49, 0x0b,
	0xc0, 0x3b, 0x0b, 0x4f, 0x83, 0x4d, 0xfd, 0x5d, 0x14, 0x2c, 0xb2, 0xc1, 0x28, 0x9b, 0x6d, 0xd3,
	0x76, 0xe1, 0x97, 0x1c, 0x48, 0xd9, 0x96, 0x33, 0x9c, 0x53, 0xee, 0xa2, 0x39, 0x35, 0x3c, 0xee,
	0xd3, 0x81, 0x78, 0x39, 0x84, 0xba, 0x8e, 0x6d, 0x8b, 0x20, 0xbb, 0x45, 0xba, 0x23, 0x9d, 0x42,
	0xc7, 0xb3, 0x8d, 0x2f, 0xb0, 0x2d, 0x27, 0x18, 0xde, 0xcf, 0x39, 0x00, 0x6d, 0xf3, 0x24, 0x20,
	0x32, 0x5a, 0xa8, 0x6d, 0xe1, 0x3a, 0xbb, 0x22, 0x36, 0xa6, 0x46, 0xaa, 0xc0, 0x3e, 0x35, 0x68,
	0x9b, 0x9c, 0x0e, 0xc4, 0x2b, 0xd3, 0xe0, 0xb1, 0x5c, 0xd9, 0x72, 0x9e, 0xf6, 0x92, 0x9f, 0x7a,
	0x43, 0xc7, 0xdb, 0xe6, 0x49, 0x20, 0x17, 0x35, 0x7f, 0xc6, 0x81, 0x74, 0xd5, 0x9f, 0x44, 0xa6,
	0xdf, 0x27, 0x80, 0x4d, 0x66, 0x90, 0x1b, 0x77, 0x51, 0x6e, 0x77, 0x58, 0x6e, 0xeb, 0x63, 0xb8,
	0xb1, 0xb4, 0x56, 0xc7, 0x16, 0x41, 0x38, 0xa3, 0x34, 0xb5, 0xb1, 0x6c, 0x7e, 0x0d, 0xe6, 0x9f,
	0x25, 0xf3, 0x10, 0x24, 0x3e, 0xee, 0xe0, 0x76, 0xc7, 0xf6, 0xb3, 0x48, 0x2b, 0xca, 0x6c, 0x1f,
	0x43, 0xa7, 0x03, 0x91, 0xa7, 0xf8, 0x51, 0x36, 0x1a, 0x63, 0x84, 0x35, 0x90, 0x24, 0x47, 0x6d,
	0xe4, 0x1e, 0xe1, 0x26, 0xfd, 0x01, 0xd2, 0x33, 0x0d, 0x23, 0xa5, 0x5f, 0x19, 0x52, 0x84, 0x22,
	0x8c, 0x78, 0x61, 0x8f, 0x03, 0x4b, 0xde, 0x84, 0x1a, 0xa3, 0x50, 0x31, 0x3f, 0x54, 0x6d, 0xe6,
	0x50, 0x99, 0x71, 0x9e, 0x31, 0x7d, 0x2f, 0x33, 0x7d, 0xc7, 0x3c, 0x64, 0x6d, 0xd1, 0x33, 0xe8,
	0xc1, 0xfb, 0xb5, 0x3f, 0x38, 0x00, 0x42, 0x5f, 0xa8, 0xd7, 0xc1, 0x7a, 0xb5, 0xa4, 0xab, 0x46,
	0xa9, 0xac, 0x17, 0x4b, 0xfb, 0xc6, 0xfd, 0xfd, 0x4a, 0x59, 0xdd, 0x2d, 0xee, 0x15, 0xd5, 0x02,
	0x1f, 0x11, 0x96, 0x7b, 0x7d, 0x29, 0x45, 0x1d, 0x55, 0x2f, 0x08, 0x94, 0xc1, 0x72, 0xd8, 0xfb,
	0x81, 0x5a, 0xe1, 0x39, 0x61, 0xb1, 0xd7, 0x97, 0x92, 0xd4, 0xeb, 0x01, 0x72, 0xe1, 0x35, 0xb0,
	0x12, 0xf6, 0xc9, 0x2b, 0x15, 0x3d, 0x5f, 0xdc, 0xe7, 0xa3, 0xc2, 0xa5, 0x5e, 0x5f, 0x5a, 0xa4,
	0x7e, 0x79, 0xb6, 0x4e, 0x25, 0xb0, 0x14, 0xf6, 0xdd, 0x2f, 0xf1, 0x31, 0x21, 0xdd, 0xeb, 0x4b,
	0x0b, 0xd4, 0x6d, 0x1f, 0xc3, 0x6d, 0x90, 0x19, 0xf7, 0x30, 0x0e, 0x8a, 0xfa, 0x5d, 0xa3, 0xaa,
	0xea, 0x25, 0x3e, 0x2e, 0xac, 0xf6, 0xfa, 0x12, 0x1f, 0xf8, 0x06, 0xbb, 0x4f, 0x88, 0x3f, 0xf9,
	0x2a, 0x1b, 0xb9, 0xf6, 0x43, 0x14, 0x2c, 0x8d, 0x7f, 0x1e, 0xc1, 0x1c, 0xf8, 0x57, 0x59, 0x2b,
	0x95, 0x4b, 0x95, 0xfc, 0x3d, 0xa3, 0xa2, 0xe7, 0xf5, 0xfb, 0x95, 0x89, 0x82, 0xfd, 0x52, 0xa8,
	0xf3, 0xbe, 0xd5, 0x84, 0x77, 0x40, 0x76, 0xd2, 0xbf, 0xa0, 0x96, 0x4b, 0x95, 0xa2, 0x6e, 0x94,
	0x55, 0xad, 0x58, 0x2a, 0xf0, 0x9c, 0xb0, 0xde, 0xeb, 0x4b, 0x2b, 0x14, 0x32, 0x36, 0x54, 0xf0,
	0x2d, 0xf0, 0xef, 0x49, 0x70, 0xb5, 0xa4, 0x17, 0xf7, 0xdf, 0x0b, 0xb0, 0x51, 0x61, 0xad, 0xd7,
	0x97, 0x20, 0xc5, 0x56, 0x43, 0x13, 0x00, 0xaf, 0x83, 0xb5, 0x49, 0x68, 0x39, 0x5f, 0xa9, 0xa8,
	0x05, 0x3e, 0x26, 0xf0, 0xbd, 0xbe, 0x94, 0xa6, 0x98, 0xb2, 0xe9, 0xba, 0xa8, 0x0e, 0x6f, 0x82,
	0xcc, 0xa4, 0xb7, 0xa6, 0xbe, 0xaf, 0xee, 0xea, 0x6a, 0x81, 0x8f, 0x0b, 0xb0, 0xd7, 0x97, 0x96,
	0xa8, 0xbf, 0x86, 0x3e, 0x44, 0x35, 0x82, 0xce, 0xe5, 0xdf, 0xcb, 0x17, 0xef, 0xa9, 0x05, 0x7e,
	0x2e, 0xcc, 0xbf, 0x67, 0x5a, 0x4d, 0x54, 0xa7, 0x72, 0x2a, 0xb7, 0x9f, 0xbf, 0xcc, 0x46, 0x7e,
	0x7e, 0x99, 0x8d, 0x7c, 0xfa, 0x5b, 0x36, 0xf2, 0xf0, 0xef, 0x17, 0xdf, 0x89, 0xff, 0x6f, 0x9e,
	0xdf, 0xb7, 0x87, 0x09, 0x7f, 0x57, 0xfc, 0xff, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x12,
	0xbf, 0xfe, 0x01, 0x0e, 0x00, 0x00,
}
