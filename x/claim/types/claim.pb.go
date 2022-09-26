// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ConditionType defines the type of condition that a recipient must execute in order to receive a claimable amount.
type ConditionType int32

const (
	// CONDITION_TYPE_UNSPECIFIED specifies an unknown condition type
	ConditionTypeUnspecified ConditionType = 0
	// CONDITION_TYPE_DEPOSIT specifies deposit condition type
	ConditionTypeDeposit ConditionType = 1
	// CONDITION_TYPE_SWAP specifies swap condition type
	ConditionTypeSwap ConditionType = 2
	// CONDITION_TYPE_STAKE specifies staking condition
	ConditionTypeStake ConditionType = 3
	// CONDITION_TYPE_VOTE specifies governance vote condition type
	ConditionTypeVote ConditionType = 4
)

var ConditionType_name = map[int32]string{
	0: "CONDITION_TYPE_UNSPECIFIED",
	1: "CONDITION_TYPE_DEPOSIT",
	2: "CONDITION_TYPE_SWAP",
	3: "CONDITION_TYPE_STAKE",
	4: "CONDITION_TYPE_VOTE",
}

var ConditionType_value = map[string]int32{
	"CONDITION_TYPE_UNSPECIFIED": 0,
	"CONDITION_TYPE_DEPOSIT":     1,
	"CONDITION_TYPE_SWAP":        2,
	"CONDITION_TYPE_STAKE":       3,
	"CONDITION_TYPE_VOTE":        4,
}

func (x ConditionType) String() string {
	return proto.EnumName(ConditionType_name, int32(x))
}

func (ConditionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef2385a7eff0dbd5, []int{0}
}

// Airdrop defines airdrop information.
type Airdrop struct {
	// id specifies index of the airdrop
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// source_address defines the bech32-encoded source address where the source of coins from
	SourceAddress string `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// conditions specifies a list of conditions
	Conditions []ConditionType `protobuf:"varint,3,rep,packed,name=conditions,proto3,enum=ollo.claim.ConditionType" json:"conditions,omitempty"`
	// start_time specifies the start time of the airdrop
	StartTime time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end_time specifies the start time of the airdrop
	EndTime time.Time `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
}

func (m *Airdrop) Reset()         { *m = Airdrop{} }
func (m *Airdrop) String() string { return proto.CompactTextString(m) }
func (*Airdrop) ProtoMessage()    {}
func (*Airdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2385a7eff0dbd5, []int{0}
}
func (m *Airdrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Airdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Airdrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Airdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airdrop.Merge(m, src)
}
func (m *Airdrop) XXX_Size() int {
	return m.Size()
}
func (m *Airdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Airdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Airdrop proto.InternalMessageInfo

// ClaimRecord defines claim record that corresponds to the airdrop.
type ClaimRecord struct {
	// airdrop_id specifies airdrop id
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	// recipient specifies the bech32-encoded address that is eligible to claim airdrop
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// initial_claimable_coins specifies the initial claimable coins
	InitialClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_claimable_coins,json=initialClaimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_coins"`
	// claimable_coins specifies the unclaimed claimable coins
	ClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=claimable_coins,json=claimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimable_coins"`
	// claimed_conditions specifies a list of condition types
	// initial values are empty and each condition type gets appended when claim is successfully executed
	ClaimedConditions []ConditionType `protobuf:"varint,5,rep,packed,name=claimed_conditions,json=claimedConditions,proto3,enum=ollo.claim.ConditionType" json:"claimed_conditions,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2385a7eff0dbd5, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.claim.ConditionType", ConditionType_name, ConditionType_value)
	proto.RegisterType((*Airdrop)(nil), "ollo.claim.Airdrop")
	proto.RegisterType((*ClaimRecord)(nil), "ollo.claim.ClaimRecord")
}

func init() { proto.RegisterFile("claim/claim.proto", fileDescriptor_ef2385a7eff0dbd5) }

var fileDescriptor_ef2385a7eff0dbd5 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6f, 0xd3, 0x3e,
	0x1c, 0x4d, 0xda, 0xee, 0xbb, 0xd5, 0xd3, 0xfa, 0xed, 0xcc, 0x36, 0xba, 0x68, 0xa4, 0xd1, 0x24,
	0xa4, 0x0a, 0x09, 0x67, 0x1b, 0x5c, 0x90, 0x90, 0x50, 0x97, 0x16, 0x51, 0x21, 0xad, 0x55, 0x9b,
	0x0d, 0xc1, 0x25, 0x4a, 0x63, 0xaf, 0x58, 0x4b, 0xe3, 0x28, 0xf6, 0x80, 0x1d, 0x11, 0x17, 0xd4,
	0xd3, 0xee, 0xa8, 0x27, 0x6e, 0xfc, 0x25, 0x3b, 0xee, 0xc8, 0x89, 0xc1, 0xfa, 0x8f, 0xa0, 0xc4,
	0xe9, 0xb6, 0x54, 0x93, 0x10, 0x12, 0x97, 0xfc, 0x78, 0xfe, 0x3c, 0xbf, 0xcf, 0xe7, 0x3d, 0x27,
	0x60, 0xd9, 0xf3, 0x5d, 0x3a, 0x34, 0x93, 0x2b, 0x0a, 0x23, 0x26, 0x18, 0x04, 0xcc, 0xf7, 0x19,
	0x4a, 0x10, 0x6d, 0x65, 0xc0, 0x06, 0x2c, 0x81, 0xcd, 0xf8, 0x49, 0x56, 0x68, 0xd5, 0x01, 0x63,
	0x03, 0x9f, 0x98, 0xc9, 0x5b, 0xff, 0xf8, 0xd0, 0x14, 0x74, 0x48, 0xb8, 0x70, 0x87, 0x61, 0x5a,
	0xa0, 0x7b, 0x8c, 0x0f, 0x19, 0x37, 0xfb, 0x2e, 0x27, 0xe6, 0xbb, 0xed, 0x3e, 0x11, 0xee, 0xb6,
	0xe9, 0x31, 0x1a, 0xc8, 0xf5, 0xcd, 0x8f, 0x39, 0x30, 0x5f, 0xa7, 0x11, 0x8e, 0x58, 0x08, 0x4b,
	0x20, 0x47, 0x71, 0x45, 0x35, 0xd4, 0x5a, 0xa1, 0x9b, 0xa3, 0x18, 0xde, 0x07, 0x25, 0xce, 0x8e,
	0x23, 0x8f, 0x38, 0x2e, 0xc6, 0x11, 0xe1, 0xbc, 0x92, 0x33, 0xd4, 0x5a, 0xb1, 0xbb, 0x24, 0xd1,
	0xba, 0x04, 0xe1, 0x13, 0x00, 0x3c, 0x16, 0x60, 0x2a, 0x28, 0x0b, 0x78, 0x25, 0x6f, 0xe4, 0x6b,
	0xa5, 0x9d, 0x75, 0x74, 0xdd, 0x3a, 0xb2, 0xa6, 0xab, 0xf6, 0x49, 0x48, 0xba, 0x37, 0x8a, 0xa1,
	0x05, 0x00, 0x17, 0x6e, 0x24, 0x9c, 0xb8, 0xed, 0x4a, 0xc1, 0x50, 0x6b, 0x8b, 0x3b, 0x1a, 0x92,
	0x33, 0xa1, 0xe9, 0x4c, 0xc8, 0x9e, 0xce, 0xb4, 0xbb, 0x70, 0xf6, 0xa3, 0xaa, 0x9c, 0x5e, 0x54,
	0xd5, 0x6e, 0x31, 0xe1, 0xc5, 0x2b, 0xf0, 0x19, 0x58, 0x20, 0x01, 0x96, 0x5b, 0xcc, 0xfd, 0xc5,
	0x16, 0xf3, 0x24, 0xc0, 0x31, 0xbe, 0x79, 0x9a, 0x07, 0x8b, 0x56, 0xdc, 0x69, 0x97, 0x78, 0x2c,
	0xc2, 0xf0, 0x1e, 0x00, 0xae, 0xb4, 0xc4, 0xb9, 0xf2, 0xa3, 0x98, 0x22, 0x2d, 0x0c, 0x37, 0x40,
	0x31, 0x22, 0x1e, 0x0d, 0x29, 0x09, 0x44, 0xea, 0xc8, 0x35, 0x00, 0x3f, 0xa9, 0xe0, 0x2e, 0x0d,
	0xa8, 0xa0, 0xae, 0xef, 0x24, 0xe3, 0xbb, 0x7d, 0x9f, 0x38, 0xb1, 0xe3, 0xd2, 0x9b, 0xc5, 0x9d,
	0x75, 0x24, 0x33, 0x41, 0x71, 0x26, 0x28, 0xcd, 0x04, 0x59, 0x8c, 0x06, 0xbb, 0x5b, 0x71, 0x73,
	0xdf, 0x2e, 0xaa, 0xb5, 0x01, 0x15, 0x6f, 0x8f, 0xfb, 0xc8, 0x63, 0x43, 0x33, 0x0d, 0x50, 0xde,
	0x1e, 0x72, 0x7c, 0x64, 0x8a, 0x93, 0x90, 0xf0, 0x84, 0xc0, 0xbb, 0xab, 0xa9, 0x96, 0x35, 0x95,
	0x4a, 0x60, 0x28, 0xc0, 0xff, 0xb3, 0xe2, 0x85, 0x7f, 0x2f, 0x5e, 0xf2, 0xb2, 0xaa, 0x2f, 0x00,
	0x4c, 0x10, 0x82, 0x9d, 0x1b, 0x27, 0x62, 0xee, 0x4f, 0x27, 0x62, 0x39, 0x25, 0x5d, 0xa1, 0xfc,
	0xc1, 0x97, 0x1c, 0x58, 0xca, 0x14, 0xc1, 0xa7, 0x40, 0xb3, 0xda, 0x7b, 0x8d, 0x96, 0xdd, 0x6a,
	0xef, 0x39, 0xf6, 0xeb, 0x4e, 0xd3, 0xd9, 0xdf, 0xeb, 0x75, 0x9a, 0x56, 0xeb, 0x79, 0xab, 0xd9,
	0x28, 0x2b, 0xda, 0xc6, 0x68, 0x6c, 0x54, 0x32, 0x94, 0xfd, 0x80, 0x87, 0xc4, 0xa3, 0x87, 0x94,
	0x60, 0xf8, 0x18, 0xac, 0xcd, 0xb0, 0x1b, 0xcd, 0x4e, 0xbb, 0xd7, 0xb2, 0xcb, 0xaa, 0x56, 0x19,
	0x8d, 0x8d, 0x95, 0x0c, 0xb3, 0x41, 0x42, 0xc6, 0xa9, 0x80, 0x08, 0xdc, 0x99, 0x61, 0xf5, 0x5e,
	0xd5, 0x3b, 0xe5, 0x9c, 0xb6, 0x3a, 0x1a, 0x1b, 0xcb, 0x19, 0x4a, 0xef, 0xbd, 0x1b, 0xc2, 0x2d,
	0xb0, 0x32, 0x5b, 0x6f, 0xd7, 0x5f, 0x36, 0xcb, 0x79, 0x6d, 0x6d, 0x34, 0x36, 0x60, 0x96, 0x20,
	0xdc, 0x23, 0x72, 0x8b, 0xc2, 0x41, 0xdb, 0x6e, 0x96, 0x0b, 0xb7, 0x28, 0x1c, 0x30, 0x41, 0xb4,
	0xc2, 0xe7, 0xaf, 0xba, 0xb2, 0xbb, 0x75, 0xf6, 0x4b, 0x57, 0xce, 0x2e, 0x75, 0xf5, 0xfc, 0x52,
	0x57, 0x7f, 0x5e, 0xea, 0xea, 0xe9, 0x44, 0x57, 0xce, 0x27, 0xba, 0xf2, 0x7d, 0xa2, 0x2b, 0x6f,
	0x60, 0x6c, 0xb4, 0xf9, 0x41, 0xfe, 0x49, 0x64, 0x5e, 0xfd, 0xff, 0x92, 0x2f, 0xe1, 0xd1, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x51, 0x60, 0xd9, 0x65, 0x04, 0x00, 0x00,
}

func (m *Airdrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Airdrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Airdrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClaim(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintClaim(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Conditions) > 0 {
		dAtA4 := make([]byte, len(m.Conditions)*10)
		var j3 int
		for _, num := range m.Conditions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaim(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedConditions) > 0 {
		dAtA6 := make([]byte, len(m.ClaimedConditions)*10)
		var j5 int
		for _, num := range m.ClaimedConditions {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintClaim(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimableCoins) > 0 {
		for iNdEx := len(m.ClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitialClaimableCoins) > 0 {
		for iNdEx := len(m.InitialClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Airdrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClaim(uint64(m.Id))
	}
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.Conditions) > 0 {
		l = 0
		for _, e := range m.Conditions {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovClaim(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovClaim(uint64(l))
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovClaim(uint64(m.AirdropId))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableCoins) > 0 {
		for _, e := range m.InitialClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ClaimableCoins) > 0 {
		for _, e := range m.ClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ClaimedConditions) > 0 {
		l = 0
		for _, e := range m.ClaimedConditions {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Airdrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Airdrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Airdrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v ConditionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ConditionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Conditions = append(m.Conditions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Conditions) == 0 {
					m.Conditions = make([]ConditionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ConditionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ConditionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Conditions = append(m.Conditions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableCoins = append(m.InitialClaimableCoins, types.Coin{})
			if err := m.InitialClaimableCoins[len(m.InitialClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimableCoins = append(m.ClaimableCoins, types.Coin{})
			if err := m.ClaimableCoins[len(m.ClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType == 0 {
				var v ConditionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ConditionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedConditions = append(m.ClaimedConditions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.ClaimedConditions) == 0 {
					m.ClaimedConditions = make([]ConditionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ConditionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ConditionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedConditions = append(m.ClaimedConditions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedConditions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
