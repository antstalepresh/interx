// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cosmos/distribution/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

// DelegatorWithdrawInfo is the address for where distributions rewards are
// withdrawn to by default this struct is only used at genesis to feed in
// default withdraw addresses.
type DelegatorWithdrawInfo struct {
	// delegator_address is the address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// withdraw_address is the address to withdraw the delegation rewards to.
	WithdrawAddress      string   `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegatorWithdrawInfo) Reset()         { *m = DelegatorWithdrawInfo{} }
func (m *DelegatorWithdrawInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorWithdrawInfo) ProtoMessage()    {}
func (*DelegatorWithdrawInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{0}
}

func (m *DelegatorWithdrawInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegatorWithdrawInfo.Unmarshal(m, b)
}
func (m *DelegatorWithdrawInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegatorWithdrawInfo.Marshal(b, m, deterministic)
}
func (m *DelegatorWithdrawInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorWithdrawInfo.Merge(m, src)
}
func (m *DelegatorWithdrawInfo) XXX_Size() int {
	return xxx_messageInfo_DelegatorWithdrawInfo.Size(m)
}
func (m *DelegatorWithdrawInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorWithdrawInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorWithdrawInfo proto.InternalMessageInfo

func (m *DelegatorWithdrawInfo) GetDelegatorAddress() string {
	if m != nil {
		return m.DelegatorAddress
	}
	return ""
}

func (m *DelegatorWithdrawInfo) GetWithdrawAddress() string {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

// ValidatorOutstandingRewardsRecord is used for import/export via genesis json.
type ValidatorOutstandingRewardsRecord struct {
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// outstanding_rewards represents the oustanding rewards of a validator.
	OutstandingRewards   []*types.DecCoin `protobuf:"bytes,2,rep,name=outstanding_rewards,json=outstandingRewards,proto3" json:"outstanding_rewards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ValidatorOutstandingRewardsRecord) Reset()         { *m = ValidatorOutstandingRewardsRecord{} }
func (m *ValidatorOutstandingRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorOutstandingRewardsRecord) ProtoMessage()    {}
func (*ValidatorOutstandingRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{1}
}

func (m *ValidatorOutstandingRewardsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorOutstandingRewardsRecord.Unmarshal(m, b)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorOutstandingRewardsRecord.Marshal(b, m, deterministic)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorOutstandingRewardsRecord.Merge(m, src)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_Size() int {
	return xxx_messageInfo_ValidatorOutstandingRewardsRecord.Size(m)
}
func (m *ValidatorOutstandingRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorOutstandingRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorOutstandingRewardsRecord proto.InternalMessageInfo

func (m *ValidatorOutstandingRewardsRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorOutstandingRewardsRecord) GetOutstandingRewards() []*types.DecCoin {
	if m != nil {
		return m.OutstandingRewards
	}
	return nil
}

// ValidatorAccumulatedCommissionRecord is used for import / export via genesis
// json.
type ValidatorAccumulatedCommissionRecord struct {
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// accumulated is the accumulated commission of a validator.
	Accumulated          *ValidatorAccumulatedCommission `protobuf:"bytes,2,opt,name=accumulated,proto3" json:"accumulated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ValidatorAccumulatedCommissionRecord) Reset()         { *m = ValidatorAccumulatedCommissionRecord{} }
func (m *ValidatorAccumulatedCommissionRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorAccumulatedCommissionRecord) ProtoMessage()    {}
func (*ValidatorAccumulatedCommissionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{2}
}

func (m *ValidatorAccumulatedCommissionRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Unmarshal(m, b)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Marshal(b, m, deterministic)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Merge(m, src)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_Size() int {
	return xxx_messageInfo_ValidatorAccumulatedCommissionRecord.Size(m)
}
func (m *ValidatorAccumulatedCommissionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorAccumulatedCommissionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorAccumulatedCommissionRecord proto.InternalMessageInfo

func (m *ValidatorAccumulatedCommissionRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorAccumulatedCommissionRecord) GetAccumulated() *ValidatorAccumulatedCommission {
	if m != nil {
		return m.Accumulated
	}
	return nil
}

// ValidatorHistoricalRewardsRecord is used for import / export via genesis
// json.
type ValidatorHistoricalRewardsRecord struct {
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// period defines the period the historical rewards apply to.
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// rewards defines the historical rewards of a validator.
	Rewards              *ValidatorHistoricalRewards `protobuf:"bytes,3,opt,name=rewards,proto3" json:"rewards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValidatorHistoricalRewardsRecord) Reset()         { *m = ValidatorHistoricalRewardsRecord{} }
func (m *ValidatorHistoricalRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorHistoricalRewardsRecord) ProtoMessage()    {}
func (*ValidatorHistoricalRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{3}
}

func (m *ValidatorHistoricalRewardsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorHistoricalRewardsRecord.Unmarshal(m, b)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorHistoricalRewardsRecord.Marshal(b, m, deterministic)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorHistoricalRewardsRecord.Merge(m, src)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_Size() int {
	return xxx_messageInfo_ValidatorHistoricalRewardsRecord.Size(m)
}
func (m *ValidatorHistoricalRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorHistoricalRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorHistoricalRewardsRecord proto.InternalMessageInfo

func (m *ValidatorHistoricalRewardsRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorHistoricalRewardsRecord) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *ValidatorHistoricalRewardsRecord) GetRewards() *ValidatorHistoricalRewards {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// ValidatorCurrentRewardsRecord is used for import / export via genesis json.
type ValidatorCurrentRewardsRecord struct {
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// rewards defines the current rewards of a validator.
	Rewards              *ValidatorCurrentRewards `protobuf:"bytes,2,opt,name=rewards,proto3" json:"rewards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ValidatorCurrentRewardsRecord) Reset()         { *m = ValidatorCurrentRewardsRecord{} }
func (m *ValidatorCurrentRewardsRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorCurrentRewardsRecord) ProtoMessage()    {}
func (*ValidatorCurrentRewardsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{4}
}

func (m *ValidatorCurrentRewardsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorCurrentRewardsRecord.Unmarshal(m, b)
}
func (m *ValidatorCurrentRewardsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorCurrentRewardsRecord.Marshal(b, m, deterministic)
}
func (m *ValidatorCurrentRewardsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorCurrentRewardsRecord.Merge(m, src)
}
func (m *ValidatorCurrentRewardsRecord) XXX_Size() int {
	return xxx_messageInfo_ValidatorCurrentRewardsRecord.Size(m)
}
func (m *ValidatorCurrentRewardsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorCurrentRewardsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorCurrentRewardsRecord proto.InternalMessageInfo

func (m *ValidatorCurrentRewardsRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorCurrentRewardsRecord) GetRewards() *ValidatorCurrentRewards {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// DelegatorStartingInfoRecord used for import / export via genesis json.
type DelegatorStartingInfoRecord struct {
	// delegator_address is the address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// starting_info defines the starting info of a delegator.
	StartingInfo         *DelegatorStartingInfo `protobuf:"bytes,3,opt,name=starting_info,json=startingInfo,proto3" json:"starting_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DelegatorStartingInfoRecord) Reset()         { *m = DelegatorStartingInfoRecord{} }
func (m *DelegatorStartingInfoRecord) String() string { return proto.CompactTextString(m) }
func (*DelegatorStartingInfoRecord) ProtoMessage()    {}
func (*DelegatorStartingInfoRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{5}
}

func (m *DelegatorStartingInfoRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegatorStartingInfoRecord.Unmarshal(m, b)
}
func (m *DelegatorStartingInfoRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegatorStartingInfoRecord.Marshal(b, m, deterministic)
}
func (m *DelegatorStartingInfoRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorStartingInfoRecord.Merge(m, src)
}
func (m *DelegatorStartingInfoRecord) XXX_Size() int {
	return xxx_messageInfo_DelegatorStartingInfoRecord.Size(m)
}
func (m *DelegatorStartingInfoRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorStartingInfoRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorStartingInfoRecord proto.InternalMessageInfo

func (m *DelegatorStartingInfoRecord) GetDelegatorAddress() string {
	if m != nil {
		return m.DelegatorAddress
	}
	return ""
}

func (m *DelegatorStartingInfoRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *DelegatorStartingInfoRecord) GetStartingInfo() *DelegatorStartingInfo {
	if m != nil {
		return m.StartingInfo
	}
	return nil
}

// ValidatorSlashEventRecord is used for import / export via genesis json.
type ValidatorSlashEventRecord struct {
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// height defines the block height at which the slash event occured.
	Height uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// period is the period of the slash event.
	Period uint64 `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	// validator_slash_event describes the slash event.
	ValidatorSlashEvent  *ValidatorSlashEvent `protobuf:"bytes,4,opt,name=validator_slash_event,json=validatorSlashEvent,proto3" json:"validator_slash_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValidatorSlashEventRecord) Reset()         { *m = ValidatorSlashEventRecord{} }
func (m *ValidatorSlashEventRecord) String() string { return proto.CompactTextString(m) }
func (*ValidatorSlashEventRecord) ProtoMessage()    {}
func (*ValidatorSlashEventRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{6}
}

func (m *ValidatorSlashEventRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSlashEventRecord.Unmarshal(m, b)
}
func (m *ValidatorSlashEventRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSlashEventRecord.Marshal(b, m, deterministic)
}
func (m *ValidatorSlashEventRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSlashEventRecord.Merge(m, src)
}
func (m *ValidatorSlashEventRecord) XXX_Size() int {
	return xxx_messageInfo_ValidatorSlashEventRecord.Size(m)
}
func (m *ValidatorSlashEventRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSlashEventRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSlashEventRecord proto.InternalMessageInfo

func (m *ValidatorSlashEventRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorSlashEventRecord) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ValidatorSlashEventRecord) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *ValidatorSlashEventRecord) GetValidatorSlashEvent() *ValidatorSlashEvent {
	if m != nil {
		return m.ValidatorSlashEvent
	}
	return nil
}

// GenesisState defines the distribution module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// fee_pool defines the fee pool at genesis.
	FeePool *FeePool `protobuf:"bytes,2,opt,name=fee_pool,json=feePool,proto3" json:"fee_pool,omitempty"`
	// fee_pool defines the delegator withdraw infos at genesis.
	DelegatorWithdrawInfos []*DelegatorWithdrawInfo `protobuf:"bytes,3,rep,name=delegator_withdraw_infos,json=delegatorWithdrawInfos,proto3" json:"delegator_withdraw_infos,omitempty"`
	// fee_pool defines the previous proposer at genesis.
	PreviousProposer string `protobuf:"bytes,4,opt,name=previous_proposer,json=previousProposer,proto3" json:"previous_proposer,omitempty"`
	// fee_pool defines the outstanding rewards of all validators at genesis.
	OutstandingRewards []*ValidatorOutstandingRewardsRecord `protobuf:"bytes,5,rep,name=outstanding_rewards,json=outstandingRewards,proto3" json:"outstanding_rewards,omitempty"`
	// fee_pool defines the accumulated commisions of all validators at genesis.
	ValidatorAccumulatedCommissions []*ValidatorAccumulatedCommissionRecord `protobuf:"bytes,6,rep,name=validator_accumulated_commissions,json=validatorAccumulatedCommissions,proto3" json:"validator_accumulated_commissions,omitempty"`
	// fee_pool defines the historical rewards of all validators at genesis.
	ValidatorHistoricalRewards []*ValidatorHistoricalRewardsRecord `protobuf:"bytes,7,rep,name=validator_historical_rewards,json=validatorHistoricalRewards,proto3" json:"validator_historical_rewards,omitempty"`
	// fee_pool defines the current rewards of all validators at genesis.
	ValidatorCurrentRewards []*ValidatorCurrentRewardsRecord `protobuf:"bytes,8,rep,name=validator_current_rewards,json=validatorCurrentRewards,proto3" json:"validator_current_rewards,omitempty"`
	// fee_pool defines the delegator starting infos at genesis.
	DelegatorStartingInfos []*DelegatorStartingInfoRecord `protobuf:"bytes,9,rep,name=delegator_starting_infos,json=delegatorStartingInfos,proto3" json:"delegator_starting_infos,omitempty"`
	// fee_pool defines the validator slash events at genesis.
	ValidatorSlashEvents []*ValidatorSlashEventRecord `protobuf:"bytes,10,rep,name=validator_slash_events,json=validatorSlashEvents,proto3" json:"validator_slash_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_76eed0f9489db580, []int{7}
}

func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetFeePool() *FeePool {
	if m != nil {
		return m.FeePool
	}
	return nil
}

func (m *GenesisState) GetDelegatorWithdrawInfos() []*DelegatorWithdrawInfo {
	if m != nil {
		return m.DelegatorWithdrawInfos
	}
	return nil
}

func (m *GenesisState) GetPreviousProposer() string {
	if m != nil {
		return m.PreviousProposer
	}
	return ""
}

func (m *GenesisState) GetOutstandingRewards() []*ValidatorOutstandingRewardsRecord {
	if m != nil {
		return m.OutstandingRewards
	}
	return nil
}

func (m *GenesisState) GetValidatorAccumulatedCommissions() []*ValidatorAccumulatedCommissionRecord {
	if m != nil {
		return m.ValidatorAccumulatedCommissions
	}
	return nil
}

func (m *GenesisState) GetValidatorHistoricalRewards() []*ValidatorHistoricalRewardsRecord {
	if m != nil {
		return m.ValidatorHistoricalRewards
	}
	return nil
}

func (m *GenesisState) GetValidatorCurrentRewards() []*ValidatorCurrentRewardsRecord {
	if m != nil {
		return m.ValidatorCurrentRewards
	}
	return nil
}

func (m *GenesisState) GetDelegatorStartingInfos() []*DelegatorStartingInfoRecord {
	if m != nil {
		return m.DelegatorStartingInfos
	}
	return nil
}

func (m *GenesisState) GetValidatorSlashEvents() []*ValidatorSlashEventRecord {
	if m != nil {
		return m.ValidatorSlashEvents
	}
	return nil
}

func init() {
	proto.RegisterType((*DelegatorWithdrawInfo)(nil), "cosmos.distribution.v1beta1.DelegatorWithdrawInfo")
	proto.RegisterType((*ValidatorOutstandingRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord")
	proto.RegisterType((*ValidatorAccumulatedCommissionRecord)(nil), "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord")
	proto.RegisterType((*ValidatorHistoricalRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord")
	proto.RegisterType((*ValidatorCurrentRewardsRecord)(nil), "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord")
	proto.RegisterType((*DelegatorStartingInfoRecord)(nil), "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord")
	proto.RegisterType((*ValidatorSlashEventRecord)(nil), "cosmos.distribution.v1beta1.ValidatorSlashEventRecord")
	proto.RegisterType((*GenesisState)(nil), "cosmos.distribution.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/genesis.proto", fileDescriptor_76eed0f9489db580)
}

var fileDescriptor_76eed0f9489db580 = []byte{
	// 996 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xce, 0x3a, 0x25, 0x49, 0xc7, 0x29, 0x0d, 0xdb, 0x7c, 0xb8, 0x4e, 0xda, 0x4d, 0xa7, 0x45,
	0x04, 0x55, 0xac, 0x9b, 0x80, 0x28, 0x0a, 0x02, 0x29, 0x9b, 0x52, 0xe8, 0x89, 0x30, 0x91, 0x00,
	0x71, 0xb1, 0xc6, 0xbb, 0x63, 0x7b, 0x84, 0xbd, 0x63, 0xcd, 0x8c, 0x1d, 0xc2, 0x1f, 0x80, 0x3b,
	0xe2, 0x54, 0x0e, 0x39, 0x22, 0xc4, 0xb1, 0x77, 0xae, 0xfc, 0x08, 0x14, 0x24, 0x4e, 0x9c, 0x73,
	0xe0, 0xc0, 0x09, 0xed, 0xcc, 0xec, 0x97, 0xbd, 0x36, 0x4e, 0x9a, 0x9c, 0x12, 0x8f, 0xdf, 0x7d,
	0x9e, 0xe7, 0x7d, 0xe6, 0xfd, 0x58, 0x83, 0x37, 0x7d, 0x26, 0xba, 0x4c, 0xd4, 0x02, 0x2a, 0x24,
	0xa7, 0x8d, 0xbe, 0xa4, 0x2c, 0xac, 0x0d, 0xb6, 0x1b, 0x44, 0xe2, 0xed, 0x5a, 0x8b, 0x84, 0x44,
	0x50, 0xe1, 0xf6, 0x38, 0x93, 0xcc, 0x5e, 0xd7, 0xa1, 0x6e, 0x36, 0xd4, 0x35, 0xa1, 0xd5, 0xe5,
	0x16, 0x6b, 0x31, 0x15, 0x57, 0x8b, 0xfe, 0xd3, 0x8f, 0x54, 0xef, 0x1a, 0xf4, 0x06, 0x16, 0x24,
	0x41, 0xf5, 0x19, 0x0d, 0xcd, 0xf7, 0xee, 0x24, 0xf6, 0x1c, 0x8f, 0x8a, 0x87, 0x2f, 0x2c, 0xb0,
	0xf2, 0x84, 0x74, 0x48, 0x0b, 0x4b, 0xc6, 0xbf, 0xa0, 0xb2, 0x1d, 0x70, 0x7c, 0xf4, 0x2c, 0x6c,
	0x32, 0xfb, 0x19, 0x78, 0x2d, 0x88, 0xbf, 0xa8, 0xe3, 0x20, 0xe0, 0x44, 0x88, 0x8a, 0xb5, 0x69,
	0x6d, 0x5d, 0xf7, 0x36, 0xce, 0x4e, 0x9d, 0xca, 0x31, 0xee, 0x76, 0x76, 0xe1, 0x48, 0x08, 0x44,
	0x4b, 0xc9, 0xd9, 0x9e, 0x3e, 0xb2, 0x9f, 0x82, 0xa5, 0x23, 0x03, 0x9d, 0x20, 0x95, 0x14, 0xd2,
	0xfa, 0xd9, 0xa9, 0xb3, 0xa6, 0x91, 0x86, 0x23, 0x20, 0xba, 0x19, 0x1f, 0x19, 0x9c, 0xdd, 0x85,
	0xef, 0x4f, 0x9c, 0x99, 0xbf, 0x4f, 0x9c, 0x19, 0xf8, 0xbc, 0x04, 0xee, 0x7d, 0x8e, 0x3b, 0x34,
	0x88, 0x68, 0x3e, 0xed, 0x4b, 0x21, 0x71, 0x18, 0xd0, 0xb0, 0x85, 0xc8, 0x11, 0xe6, 0x81, 0x40,
	0xc4, 0x67, 0x3c, 0x88, 0x52, 0x18, 0xc4, 0x41, 0xe3, 0x53, 0x18, 0x09, 0x81, 0x68, 0x29, 0x39,
	0x8b, 0x53, 0x38, 0xb1, 0xc0, 0x2d, 0x96, 0xf2, 0xd4, 0xb9, 0x26, 0xaa, 0x94, 0x36, 0x67, 0xb7,
	0xca, 0x3b, 0x1b, 0xc6, 0x76, 0x37, 0xba, 0x96, 0xf8, 0x06, 0xdd, 0x27, 0xc4, 0xdf, 0x67, 0x34,
	0xf4, 0x3e, 0xfb, 0xfd, 0xd4, 0x99, 0x39, 0x3b, 0x75, 0xaa, 0x9a, 0xaf, 0x00, 0x06, 0xfe, 0xf2,
	0xa7, 0xf3, 0xb0, 0x45, 0x65, 0xbb, 0xdf, 0x70, 0x7d, 0xd6, 0xad, 0x99, 0x4b, 0xd4, 0x7f, 0xde,
	0x12, 0xc1, 0xd7, 0x35, 0x79, 0xdc, 0x23, 0x22, 0x46, 0x14, 0xc8, 0x66, 0x23, 0x39, 0x67, 0xdc,
	0xf9, 0xc7, 0x02, 0x0f, 0x12, 0x77, 0xf6, 0x7c, 0xbf, 0xdf, 0xed, 0x77, 0xb0, 0x24, 0xc1, 0x3e,
	0xeb, 0x76, 0xa9, 0x10, 0x94, 0x85, 0x97, 0x6f, 0xd0, 0x31, 0x28, 0xe3, 0x94, 0x49, 0x5d, 0x6f,
	0x79, 0xe7, 0x7d, 0x77, 0x42, 0x85, 0xbb, 0x93, 0x25, 0x7a, 0x55, 0x63, 0x9b, 0xad, 0x55, 0x64,
	0xd0, 0x21, 0xca, 0x72, 0x65, 0x12, 0xff, 0xd7, 0x02, 0x9b, 0x09, 0xea, 0x27, 0x54, 0x48, 0xc6,
	0xa9, 0x8f, 0x3b, 0x57, 0x56, 0x15, 0xab, 0x60, 0xae, 0x47, 0x38, 0x65, 0x3a, 0xdf, 0x6b, 0xc8,
	0x7c, 0xb2, 0x29, 0x98, 0x8f, 0x0b, 0x64, 0x56, 0x19, 0xf1, 0x78, 0x3a, 0x23, 0x46, 0x24, 0x7b,
	0xab, 0xc6, 0x84, 0x57, 0xb5, 0xaa, 0xb8, 0x5e, 0x50, 0x8c, 0x9f, 0x49, 0xfe, 0x0f, 0x0b, 0xdc,
	0x49, 0x90, 0xf6, 0xfb, 0x9c, 0x93, 0x50, 0x5e, 0x59, 0xe6, 0xcd, 0x34, 0x43, 0x7d, 0xd5, 0xef,
	0x4c, 0x97, 0x61, 0x5e, 0xd7, 0x79, 0xd2, 0x7b, 0x51, 0x02, 0xeb, 0xc9, 0xa4, 0x3a, 0x94, 0x98,
	0x4b, 0x1a, 0xb6, 0xa2, 0x49, 0x95, 0x26, 0x77, 0x59, 0xf3, 0xaa, 0xd0, 0xa7, 0xd2, 0x85, 0x7c,
	0xea, 0x83, 0x1b, 0xc2, 0x68, 0xad, 0xd3, 0xb0, 0xc9, 0x4c, 0x3d, 0xec, 0x4c, 0x74, 0xab, 0x30,
	0x4d, 0x6f, 0xc3, 0x78, 0xb5, 0xac, 0xe9, 0x73, 0xb0, 0x10, 0x2d, 0x8a, 0x4c, 0x6c, 0xc6, 0xb6,
	0x9f, 0x4a, 0xe0, 0x76, 0xe2, 0xfe, 0x61, 0x07, 0x8b, 0xf6, 0x47, 0x03, 0x75, 0x01, 0x57, 0xd0,
	0x0b, 0x6d, 0x42, 0x5b, 0x6d, 0x19, 0xf7, 0x82, 0xfe, 0x94, 0xe9, 0x91, 0xd9, 0x5c, 0x8f, 0x7c,
	0x0b, 0x56, 0x52, 0x5c, 0x11, 0x09, 0xab, 0x93, 0x48, 0x59, 0xe5, 0x9a, 0x72, 0xe8, 0xd1, 0x74,
	0xf5, 0x94, 0x66, 0xe4, 0x2d, 0x1b, 0x7f, 0x16, 0xb5, 0x68, 0x05, 0x06, 0xd1, 0xad, 0xc1, 0x68,
	0x68, 0xc6, 0x9e, 0xef, 0xca, 0x60, 0xf1, 0x63, 0xbd, 0x94, 0x0f, 0x25, 0x96, 0xc4, 0x46, 0x60,
	0xae, 0x87, 0x39, 0xee, 0x6a, 0x1b, 0xca, 0x3b, 0xf7, 0x27, 0xea, 0x38, 0x50, 0xa1, 0xde, 0x8a,
	0xa1, 0xbe, 0xa1, 0xa9, 0x35, 0x00, 0x44, 0x06, 0xc9, 0xfe, 0x12, 0x2c, 0x34, 0x09, 0xa9, 0xf7,
	0x18, 0xeb, 0x98, 0x6e, 0x79, 0x30, 0x11, 0xf5, 0x29, 0x21, 0x07, 0x8c, 0x75, 0xbc, 0x35, 0x03,
	0x7b, 0x53, 0xc3, 0xc6, 0x18, 0x10, 0xcd, 0x37, 0x75, 0x84, 0xfd, 0xa3, 0x05, 0x2a, 0x69, 0x49,
	0x27, 0x2b, 0x34, 0x2a, 0x89, 0x68, 0xf4, 0xcc, 0x4e, 0x5f, 0x6a, 0xd9, 0xdd, 0xef, 0xbd, 0x61,
	0x88, 0x9d, 0xe1, 0xa6, 0xc9, 0x33, 0x40, 0xb4, 0x1a, 0x14, 0x3d, 0xaf, 0x3a, 0xa8, 0xc7, 0xc9,
	0x80, 0xb2, 0xbe, 0xa8, 0xf7, 0x38, 0xeb, 0x31, 0x41, 0xb8, 0xba, 0xd8, 0x5c, 0x5d, 0x8d, 0x84,
	0x40, 0xb4, 0x14, 0x9f, 0x1d, 0x98, 0x23, 0xfb, 0x87, 0x31, 0x9b, 0xf7, 0x15, 0x95, 0xdd, 0x87,
	0xd3, 0x95, 0xc9, 0xb8, 0x57, 0x04, 0x0f, 0xfe, 0xff, 0x6e, 0x2e, 0x5a, 0xb6, 0xf6, 0x6f, 0x16,
	0xb8, 0x97, 0x69, 0x8b, 0x74, 0x1b, 0xd5, 0xfd, 0x64, 0x83, 0x89, 0xca, 0x9c, 0xd2, 0xb8, 0xf7,
	0x12, 0x5b, 0xd0, 0xc8, 0x7c, 0x64, 0x64, 0x6e, 0x8d, 0x34, 0x64, 0x31, 0x33, 0x44, 0xce, 0x60,
	0x22, 0xae, 0xb0, 0x7f, 0xb5, 0xc0, 0x46, 0x8a, 0xd3, 0x4e, 0x36, 0x4f, 0x62, 0xf0, 0xbc, 0x12,
	0xff, 0xc1, 0x05, 0x37, 0x97, 0x11, 0xfe, 0xd0, 0x08, 0xbf, 0x3f, 0x2c, 0x7c, 0x94, 0x10, 0xa2,
	0xea, 0x60, 0x2c, 0x5c, 0xf4, 0x02, 0x76, 0x3b, 0x7d, 0xda, 0xd7, 0x6b, 0x24, 0xd1, 0xba, 0xa0,
	0xb4, 0xee, 0x5e, 0x64, 0x07, 0x19, 0xa1, 0x5b, 0x46, 0xe8, 0xe6, 0xb0, 0xd0, 0x21, 0x2a, 0x88,
	0xd6, 0x06, 0xc5, 0x40, 0xf6, 0xf3, 0x5c, 0x33, 0xe6, 0xe6, 0xb3, 0xa8, 0x5c, 0x57, 0x0a, 0xdf,
	0x3b, 0xff, 0xdc, 0x37, 0xfa, 0xc6, 0xb6, 0x64, 0x9e, 0x27, 0xdb, 0x92, 0x59, 0x14, 0x11, 0xf5,
	0xd1, 0x6a, 0xe1, 0xc0, 0x15, 0x15, 0xa0, 0xb4, 0xbd, 0x7b, 0xde, 0x89, 0x6b, 0x94, 0xbd, 0x6e,
	0x94, 0xdd, 0x19, 0x76, 0x2e, 0xcb, 0x01, 0xd1, 0x72, 0xc1, 0x20, 0xce, 0xec, 0x77, 0xef, 0xf1,
	0xcf, 0x7f, 0xdd, 0xb5, 0xbe, 0xda, 0x9e, 0xf8, 0x16, 0xfc, 0x4d, 0xfe, 0x77, 0x8d, 0x7a, 0x29,
	0x6e, 0xcc, 0xa9, 0x5f, 0x32, 0x6f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x34, 0xe8, 0x18,
	0x79, 0x0d, 0x00, 0x00,
}
