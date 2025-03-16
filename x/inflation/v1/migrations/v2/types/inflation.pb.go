// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sacas/inflation/v1/inflation.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	cosmossdk_io_math "cosmossdk.io/math"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// V2InflationDistribution defines the distribution in which inflation is
// allocated through minting on each epoch (staking, incentives, community). It
// excludes the team vesting distribution, as this is minted once at genesis.
// The initial V2InflationDistribution can be calculated from the Sacas Token
// Model like this:
// mintDistribution1 = distribution1 / (1 - teamVestingDistribution)
// 0.5333333         = 40%           / (1 - 25%)
type V2InflationDistribution struct {
	// staking_rewards defines the proportion of the minted minted_denom that is
	// to be allocated as staking rewards
	StakingRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=staking_rewards,json=stakingRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_rewards"`
	// usage_incentives defines the proportion of the minted minted_denom that is
	// to be allocated to the incentives module address
	UsageIncentives cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=usage_incentives,json=usageIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"usage_incentives"`
	// community_pool defines the proportion of the minted minted_denom that is to
	// be allocated to the community pool
	CommunityPool cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool"`
}

func (m *V2InflationDistribution) Reset()         { *m = V2InflationDistribution{} }
func (m *V2InflationDistribution) String() string { return proto.CompactTextString(m) }
func (*V2InflationDistribution) ProtoMessage()    {}
func (*V2InflationDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_d064cb35c3ff7df8, []int{0}
}

func (m *V2InflationDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2InflationDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2InflationDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2InflationDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2InflationDistribution.Merge(m, src)
}

func (m *V2InflationDistribution) XXX_Size() int {
	return m.Size()
}

func (m *V2InflationDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_V2InflationDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_V2InflationDistribution proto.InternalMessageInfo

// V2ExponentialCalculation holds factors to calculate exponential inflation on
// each period. Calculation reference:
// periodProvision = exponentialDecay       *  bondingIncentive
// f(x)            = (a * (1 - r) ^ x + c)  *  (1 + max_variance - bondedRatio *
// (max_variance / bonding_target))
type V2ExponentialCalculation struct {
	// a defines the initial value
	A cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=a,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"a"`
	// r defines the reduction factor
	R cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=r,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"r"`
	// c defines the parameter for long term inflation
	C cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=c,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"c"`
	// bonding_target
	BondingTarget cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=bonding_target,json=bondingTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bonding_target"`
	// max_variance
	MaxVariance cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=max_variance,json=maxVariance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_variance"`
}

func (m *V2ExponentialCalculation) Reset()         { *m = V2ExponentialCalculation{} }
func (m *V2ExponentialCalculation) String() string { return proto.CompactTextString(m) }
func (*V2ExponentialCalculation) ProtoMessage()    {}
func (*V2ExponentialCalculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d064cb35c3ff7df8, []int{1}
}

func (m *V2ExponentialCalculation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V2ExponentialCalculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V2ExponentialCalculation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V2ExponentialCalculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2ExponentialCalculation.Merge(m, src)
}

func (m *V2ExponentialCalculation) XXX_Size() int {
	return m.Size()
}

func (m *V2ExponentialCalculation) XXX_DiscardUnknown() {
	xxx_messageInfo_V2ExponentialCalculation.DiscardUnknown(m)
}

var xxx_messageInfo_V2ExponentialCalculation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*V2InflationDistribution)(nil), "sacas.inflation.v1.V2InflationDistribution")
	proto.RegisterType((*V2ExponentialCalculation)(nil), "sacas.inflation.v1.V2ExponentialCalculation")
}

func init() {
	proto.RegisterFile("sacas/inflation/v1/inflation.proto", fileDescriptor_d064cb35c3ff7df8)
}

var fileDescriptor_d064cb35c3ff7df8 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x4e, 0xe3, 0x30,
	0x10, 0xc7, 0xe3, 0xee, 0x87, 0xb4, 0xde, 0xdd, 0x76, 0x15, 0xed, 0xae, 0xa2, 0x3d, 0xa4, 0xab,
	0x1e, 0x10, 0x07, 0x48, 0xa8, 0xb8, 0x72, 0x2a, 0x05, 0xa9, 0x37, 0xa8, 0xf8, 0x10, 0x5c, 0x22,
	0xc7, 0x35, 0xc1, 0x6a, 0xe2, 0x89, 0x6c, 0x27, 0xa4, 0x6f, 0xc1, 0x33, 0xf0, 0x34, 0x3d, 0xf6,
	0x88, 0x38, 0x54, 0xa8, 0x7d, 0x0d, 0x0e, 0xc8, 0x49, 0x69, 0x7b, 0xce, 0x25, 0x19, 0xdb, 0xff,
	0xf9, 0xd9, 0xf3, 0xb7, 0x07, 0x77, 0x58, 0x9e, 0x80, 0xf2, 0xb9, 0xb8, 0x8b, 0x89, 0xe6, 0x20,
	0xfc, 0xbc, 0xbb, 0x19, 0x78, 0xa9, 0x04, 0x0d, 0xb6, 0x5d, 0x6a, 0xbc, 0xcd, 0x74, 0xde, 0xfd,
	0xf7, 0x3b, 0x82, 0x08, 0xca, 0x65, 0xdf, 0x44, 0x95, 0xb2, 0xf3, 0xd4, 0xc0, 0x7f, 0x06, 0x1f,
	0xb2, 0x3e, 0x57, 0x5a, 0xf2, 0x30, 0x33, 0xb1, 0x7d, 0x8d, 0x5b, 0x4a, 0x93, 0x31, 0x17, 0x51,
	0x20, 0xd9, 0x03, 0x91, 0x23, 0xe5, 0xa0, 0xff, 0x68, 0xf7, 0x5b, 0xcf, 0x9b, 0xce, 0xdb, 0xd6,
	0xcb, 0xbc, 0xbd, 0x13, 0x71, 0x7d, 0x9f, 0x85, 0x1e, 0x85, 0xc4, 0xa7, 0xa0, 0xcc, 0xa1, 0xaa,
	0xdf, 0xbe, 0x1a, 0x8d, 0x7d, 0x3d, 0x49, 0x99, 0xf2, 0xfa, 0x8c, 0x0e, 0x9b, 0x2b, 0xcc, 0xb0,
	0xa2, 0xd8, 0x37, 0xf8, 0x57, 0xa6, 0x48, 0xc4, 0x02, 0x2e, 0x28, 0x13, 0x9a, 0xe7, 0x4c, 0x39,
	0x8d, 0x5a, 0xe4, 0x56, 0xc9, 0x19, 0xac, 0x31, 0xf6, 0x25, 0x6e, 0x52, 0x48, 0x92, 0x4c, 0x70,
	0x3d, 0x09, 0x52, 0x80, 0xd8, 0xf9, 0x54, 0x0b, 0xfc, 0x73, 0x4d, 0x39, 0x03, 0x88, 0x3b, 0x6f,
	0x0d, 0xfc, 0xf7, 0xa4, 0x48, 0x41, 0x98, 0x7d, 0x48, 0x7c, 0x4c, 0x62, 0x9a, 0x55, 0x8e, 0xd9,
	0x47, 0x18, 0x91, 0x9a, 0xbe, 0x20, 0x62, 0xb2, 0x65, 0xcd, 0xda, 0x91, 0x34, 0xd9, 0xb4, 0x66,
	0x81, 0x88, 0x1a, 0xaf, 0x42, 0x10, 0x23, 0x73, 0xbf, 0x9a, 0xc8, 0x88, 0x69, 0xe7, 0x73, 0x3d,
	0xaf, 0x56, 0x94, 0x8b, 0x12, 0x62, 0x9f, 0xe3, 0x1f, 0x09, 0x29, 0x82, 0x9c, 0x48, 0x4e, 0x04,
	0x65, 0xce, 0x97, 0x5a, 0xd0, 0xef, 0x09, 0x29, 0xae, 0x56, 0x88, 0xde, 0xe9, 0x74, 0xe1, 0xa2,
	0xd9, 0xc2, 0x45, 0xaf, 0x0b, 0x17, 0x3d, 0x2e, 0x5d, 0x6b, 0xb6, 0x74, 0xad, 0xe7, 0xa5, 0x6b,
	0xdd, 0xee, 0x6d, 0xe1, 0xaa, 0xb6, 0xa8, 0xbe, 0x79, 0xf7, 0xc0, 0x2f, 0xb6, 0x5a, 0xa4, 0x04,
	0x87, 0x5f, 0xcb, 0x27, 0x7f, 0xf8, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x53, 0x51, 0xc3, 0x26, 0x42,
	0x03, 0x00, 0x00,
}

func (m *V2InflationDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2InflationDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2InflationDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UsageIncentives.Size()
		i -= size
		if _, err := m.UsageIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.StakingRewards.Size()
		i -= size
		if _, err := m.StakingRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *V2ExponentialCalculation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V2ExponentialCalculation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V2ExponentialCalculation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxVariance.Size()
		i -= size
		if _, err := m.MaxVariance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BondingTarget.Size()
		i -= size
		if _, err := m.BondingTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.C.Size()
		i -= size
		if _, err := m.C.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.R.Size()
		i -= size
		if _, err := m.R.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.A.Size()
		i -= size
		if _, err := m.A.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *V2InflationDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StakingRewards.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.UsageIncentives.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func (m *V2ExponentialCalculation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.A.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.R.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.C.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.BondingTarget.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.MaxVariance.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func sovInflation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozInflation(x uint64) (n int) {
	return sovInflation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V2InflationDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: V2InflationDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2InflationDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsageIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UsageIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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

func (m *V2ExponentialCalculation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: V2ExponentialCalculation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V2ExponentialCalculation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.R.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondingTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondingTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVariance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVariance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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

func skipInflation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
				return 0, ErrInvalidLengthInflation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflation = fmt.Errorf("proto: unexpected end of group")
)
