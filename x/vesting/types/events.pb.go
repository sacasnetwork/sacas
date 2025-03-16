// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sacas/vesting/v2/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventCreateClawbackVestingAccount defines the event type
// for creating a clawback vesting account
type EventCreateClawbackVestingAccount struct {
	// funder is the address of the funder
	Funder string `protobuf:"bytes,1,opt,name=funder,proto3" json:"funder,omitempty"`
	// vesting_account is the address of the created vesting account
	VestingAccount string `protobuf:"bytes,2,opt,name=vesting_account,json=vestingAccount,proto3" json:"vesting_account,omitempty"`
}

func (m *EventCreateClawbackVestingAccount) Reset()         { *m = EventCreateClawbackVestingAccount{} }
func (m *EventCreateClawbackVestingAccount) String() string { return proto.CompactTextString(m) }
func (*EventCreateClawbackVestingAccount) ProtoMessage()    {}
func (*EventCreateClawbackVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_29fa0a11592c0b4a, []int{0}
}
func (m *EventCreateClawbackVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateClawbackVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateClawbackVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateClawbackVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateClawbackVestingAccount.Merge(m, src)
}
func (m *EventCreateClawbackVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateClawbackVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateClawbackVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateClawbackVestingAccount proto.InternalMessageInfo

func (m *EventCreateClawbackVestingAccount) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

func (m *EventCreateClawbackVestingAccount) GetVestingAccount() string {
	if m != nil {
		return m.VestingAccount
	}
	return ""
}

// EventFundVestingAccount defines the event type for funding a vesting account
type EventFundVestingAccount struct {
	// funder is the address of the funder
	Funder string `protobuf:"bytes,1,opt,name=funder,proto3" json:"funder,omitempty"`
	// coins to be vested
	Coins string `protobuf:"bytes,2,opt,name=coins,proto3" json:"coins,omitempty"`
	// start_time is the time when the coins start to vest
	StartTime string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// vesting_account is the address of the recipient
	VestingAccount string `protobuf:"bytes,5,opt,name=vesting_account,json=vestingAccount,proto3" json:"vesting_account,omitempty"`
}

func (m *EventFundVestingAccount) Reset()         { *m = EventFundVestingAccount{} }
func (m *EventFundVestingAccount) String() string { return proto.CompactTextString(m) }
func (*EventFundVestingAccount) ProtoMessage()    {}
func (*EventFundVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_29fa0a11592c0b4a, []int{1}
}
func (m *EventFundVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFundVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFundVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFundVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFundVestingAccount.Merge(m, src)
}
func (m *EventFundVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *EventFundVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFundVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EventFundVestingAccount proto.InternalMessageInfo

func (m *EventFundVestingAccount) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

func (m *EventFundVestingAccount) GetCoins() string {
	if m != nil {
		return m.Coins
	}
	return ""
}

func (m *EventFundVestingAccount) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *EventFundVestingAccount) GetVestingAccount() string {
	if m != nil {
		return m.VestingAccount
	}
	return ""
}

// EventClawback defines the event type for clawback
type EventClawback struct {
	// funder is the address of the funder
	Funder string `protobuf:"bytes,1,opt,name=funder,proto3" json:"funder,omitempty"`
	// account is the address of the account
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// destination is the address of the destination
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (m *EventClawback) Reset()         { *m = EventClawback{} }
func (m *EventClawback) String() string { return proto.CompactTextString(m) }
func (*EventClawback) ProtoMessage()    {}
func (*EventClawback) Descriptor() ([]byte, []int) {
	return fileDescriptor_29fa0a11592c0b4a, []int{2}
}
func (m *EventClawback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClawback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClawback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClawback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClawback.Merge(m, src)
}
func (m *EventClawback) XXX_Size() int {
	return m.Size()
}
func (m *EventClawback) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClawback.DiscardUnknown(m)
}

var xxx_messageInfo_EventClawback proto.InternalMessageInfo

func (m *EventClawback) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

func (m *EventClawback) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EventClawback) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

// EventUpdateVestingFunder defines the event type for updating the vesting funder
type EventUpdateVestingFunder struct {
	// funder is the address of the funder
	Funder string `protobuf:"bytes,1,opt,name=funder,proto3" json:"funder,omitempty"`
	// account is the address of the account
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// new_funder is the address of the new funder
	NewFunder string `protobuf:"bytes,3,opt,name=new_funder,json=newFunder,proto3" json:"new_funder,omitempty"`
}

func (m *EventUpdateVestingFunder) Reset()         { *m = EventUpdateVestingFunder{} }
func (m *EventUpdateVestingFunder) String() string { return proto.CompactTextString(m) }
func (*EventUpdateVestingFunder) ProtoMessage()    {}
func (*EventUpdateVestingFunder) Descriptor() ([]byte, []int) {
	return fileDescriptor_29fa0a11592c0b4a, []int{3}
}
func (m *EventUpdateVestingFunder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateVestingFunder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateVestingFunder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateVestingFunder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateVestingFunder.Merge(m, src)
}
func (m *EventUpdateVestingFunder) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateVestingFunder) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateVestingFunder.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateVestingFunder proto.InternalMessageInfo

func (m *EventUpdateVestingFunder) GetFunder() string {
	if m != nil {
		return m.Funder
	}
	return ""
}

func (m *EventUpdateVestingFunder) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *EventUpdateVestingFunder) GetNewFunder() string {
	if m != nil {
		return m.NewFunder
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreateClawbackVestingAccount)(nil), "sacas.vesting.v2.EventCreateClawbackVestingAccount")
	proto.RegisterType((*EventFundVestingAccount)(nil), "sacas.vesting.v2.EventFundVestingAccount")
	proto.RegisterType((*EventClawback)(nil), "sacas.vesting.v2.EventClawback")
	proto.RegisterType((*EventUpdateVestingFunder)(nil), "sacas.vesting.v2.EventUpdateVestingFunder")
}

func init() { proto.RegisterFile("sacas/vesting/v2/events.proto", fileDescriptor_29fa0a11592c0b4a) }

var fileDescriptor_29fa0a11592c0b4a = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x1b, 0x50, 0x8b, 0x7a, 0x88, 0x3f, 0x8a, 0x10, 0x64, 0x69, 0x54, 0xba, 0xc0, 0x14,
	0x43, 0x79, 0x02, 0xa8, 0xe8, 0xc2, 0x86, 0x80, 0x81, 0xa5, 0x72, 0x9d, 0xa3, 0x58, 0xa5, 0x76,
	0x15, 0x5f, 0x12, 0x78, 0x0a, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0xa8, 0x79, 0x11, 0x84, 0x6d, 0x10,
	0x95, 0x8a, 0x04, 0xe3, 0xf9, 0x3e, 0xff, 0xee, 0xfb, 0xa4, 0x0f, 0x5a, 0x86, 0x0b, 0x6e, 0x58,
	0x81, 0x86, 0xa4, 0x1a, 0xb1, 0xa2, 0xcb, 0xb0, 0x40, 0x45, 0x26, 0x99, 0x66, 0x9a, 0x74, 0xb8,
	0x6d, 0xd7, 0x89, 0x5f, 0x27, 0x45, 0xb7, 0x93, 0xc2, 0xfe, 0xf9, 0xa7, 0xa2, 0x97, 0x21, 0x27,
	0xec, 0x3d, 0xf0, 0x72, 0xc8, 0xc5, 0xf8, 0xc6, 0x09, 0x4e, 0x85, 0xd0, 0xb9, 0xa2, 0x70, 0x17,
	0x1a, 0x77, 0xb9, 0x4a, 0x31, 0x8b, 0x82, 0x76, 0x70, 0xd8, 0xbc, 0xf4, 0x53, 0x78, 0x00, 0x5b,
	0x1e, 0x35, 0xe0, 0x4e, 0x1a, 0xad, 0x58, 0xc1, 0x66, 0xb1, 0x00, 0xe8, 0x3c, 0x07, 0xb0, 0x67,
	0xcf, 0xf4, 0x73, 0x95, 0xfe, 0x11, 0xbe, 0x03, 0x75, 0xa1, 0xa5, 0x32, 0x1e, 0xe9, 0x86, 0xb0,
	0x05, 0x60, 0x88, 0x67, 0x34, 0x20, 0x39, 0xc1, 0x68, 0xd5, 0xae, 0x9a, 0xf6, 0xe5, 0x4a, 0x4e,
	0x70, 0x99, 0xa3, 0xfa, 0x52, 0x47, 0x02, 0x36, 0x5c, 0x6e, 0x9f, 0xf8, 0x57, 0x1b, 0x11, 0xac,
	0x2d, 0x66, 0xfb, 0x1a, 0xc3, 0x36, 0xac, 0xa7, 0x16, 0xca, 0x49, 0x6a, 0xe5, 0xbd, 0xfc, 0x7c,
	0xea, 0x8c, 0x21, 0xb2, 0x47, 0xae, 0xa7, 0x29, 0x27, 0xf4, 0xb9, 0xfb, 0x8e, 0xfb, 0xff, 0x7b,
	0x2d, 0x00, 0x85, 0xe5, 0xc0, 0xff, 0xf2, 0xd1, 0x15, 0x96, 0x0e, 0x78, 0x76, 0xf1, 0x3a, 0x8f,
	0x83, 0xd9, 0x3c, 0x0e, 0xde, 0xe7, 0x71, 0xf0, 0x52, 0xc5, 0xb5, 0x59, 0x15, 0xd7, 0xde, 0xaa,
	0xb8, 0x76, 0x7b, 0x3c, 0x92, 0x74, 0x9f, 0x0f, 0x13, 0xa1, 0x27, 0xcc, 0x16, 0x40, 0x21, 0x95,
	0x3a, 0x1b, 0x33, 0x5f, 0x96, 0xee, 0x11, 0x7b, 0xfc, 0xae, 0x0c, 0x3d, 0x4d, 0xd1, 0x0c, 0x1b,
	0xb6, 0x2f, 0x27, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x03, 0xca, 0xf5, 0x50, 0x02, 0x00,
	0x00,
}

func (m *EventCreateClawbackVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateClawbackVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateClawbackVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingAccount) > 0 {
		i -= len(m.VestingAccount)
		copy(dAtA[i:], m.VestingAccount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.VestingAccount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFundVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFundVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFundVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingAccount) > 0 {
		i -= len(m.VestingAccount)
		copy(dAtA[i:], m.VestingAccount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.VestingAccount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StartTime) > 0 {
		i -= len(m.StartTime)
		copy(dAtA[i:], m.StartTime)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.StartTime)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Coins) > 0 {
		i -= len(m.Coins)
		copy(dAtA[i:], m.Coins)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Coins)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventClawback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClawback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClawback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateVestingFunder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateVestingFunder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateVestingFunder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewFunder) > 0 {
		i -= len(m.NewFunder)
		copy(dAtA[i:], m.NewFunder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewFunder)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Funder) > 0 {
		i -= len(m.Funder)
		copy(dAtA[i:], m.Funder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Funder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateClawbackVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.VestingAccount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventFundVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Coins)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.StartTime)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.VestingAccount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventClawback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateVestingFunder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Funder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewFunder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateClawbackVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateClawbackVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateClawbackVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventFundVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFundVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFundVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventClawback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventClawback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClawback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateVestingFunder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateVestingFunder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateVestingFunder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewFunder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewFunder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
