// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sacas/revenue/v1/events.proto

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

// EventRegisterRevenue is an event emitted when a contract is registered to receive a percentage of tx fees.
type EventRegisterRevenue struct {
	// deployer_address is the bech32 address of message sender. It must be the same as the origin EOA
	// sending the transaction which deploys the contract
	DeployerAddress string `protobuf:"bytes,1,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
	// contract_address in hex format
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// effective_withdrawer is the withdrawer address that is stored after the
	// revenue registration is completed. It defaults to the deployer address if
	// the withdrawer address in the msg is omitted. When omitted, the withdraw map
	// doesn't need to be set.
	EffectiveWithdrawer string `protobuf:"bytes,3,opt,name=effective_withdrawer,json=effectiveWithdrawer,proto3" json:"effective_withdrawer,omitempty"`
}

func (m *EventRegisterRevenue) Reset()         { *m = EventRegisterRevenue{} }
func (m *EventRegisterRevenue) String() string { return proto.CompactTextString(m) }
func (*EventRegisterRevenue) ProtoMessage()    {}
func (*EventRegisterRevenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9a4bb0e3643eef, []int{0}
}
func (m *EventRegisterRevenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegisterRevenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegisterRevenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegisterRevenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegisterRevenue.Merge(m, src)
}
func (m *EventRegisterRevenue) XXX_Size() int {
	return m.Size()
}
func (m *EventRegisterRevenue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegisterRevenue.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegisterRevenue proto.InternalMessageInfo

func (m *EventRegisterRevenue) GetDeployerAddress() string {
	if m != nil {
		return m.DeployerAddress
	}
	return ""
}

func (m *EventRegisterRevenue) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *EventRegisterRevenue) GetEffectiveWithdrawer() string {
	if m != nil {
		return m.EffectiveWithdrawer
	}
	return ""
}

// EventUpdateRevenue is an event emitted when a withdrawer address is updated for a contract.
type EventUpdateRevenue struct {
	// contract_address in hex format
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// deployer_address is the bech32 address of message sender. It must be the same as the origin EOA
	// sending the transaction which deploys the contract
	DeployerAddress string `protobuf:"bytes,2,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
	// withdrawer_address is the bech32 address of account receiving the transaction fees
	WithdrawerAddress string `protobuf:"bytes,3,opt,name=withdrawer_address,json=withdrawerAddress,proto3" json:"withdrawer_address,omitempty"`
}

func (m *EventUpdateRevenue) Reset()         { *m = EventUpdateRevenue{} }
func (m *EventUpdateRevenue) String() string { return proto.CompactTextString(m) }
func (*EventUpdateRevenue) ProtoMessage()    {}
func (*EventUpdateRevenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9a4bb0e3643eef, []int{1}
}
func (m *EventUpdateRevenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateRevenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateRevenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateRevenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateRevenue.Merge(m, src)
}
func (m *EventUpdateRevenue) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateRevenue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateRevenue.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateRevenue proto.InternalMessageInfo

func (m *EventUpdateRevenue) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *EventUpdateRevenue) GetDeployerAddress() string {
	if m != nil {
		return m.DeployerAddress
	}
	return ""
}

func (m *EventUpdateRevenue) GetWithdrawerAddress() string {
	if m != nil {
		return m.WithdrawerAddress
	}
	return ""
}

// EventCancelRevenue is an event emitted when a contract is unregistered from receiving tx fees.
type EventCancelRevenue struct {
	// deployer_address is the bech32 address of message sender. It must be the same as the origin EOA
	// sending the transaction which deploys the contract
	DeployerAddress string `protobuf:"bytes,1,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
	// contract_address in hex format
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *EventCancelRevenue) Reset()         { *m = EventCancelRevenue{} }
func (m *EventCancelRevenue) String() string { return proto.CompactTextString(m) }
func (*EventCancelRevenue) ProtoMessage()    {}
func (*EventCancelRevenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9a4bb0e3643eef, []int{2}
}
func (m *EventCancelRevenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCancelRevenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCancelRevenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCancelRevenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCancelRevenue.Merge(m, src)
}
func (m *EventCancelRevenue) XXX_Size() int {
	return m.Size()
}
func (m *EventCancelRevenue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCancelRevenue.DiscardUnknown(m)
}

var xxx_messageInfo_EventCancelRevenue proto.InternalMessageInfo

func (m *EventCancelRevenue) GetDeployerAddress() string {
	if m != nil {
		return m.DeployerAddress
	}
	return ""
}

func (m *EventCancelRevenue) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// EventDistributeRevenue is an event emitted when a contract receives a percentage of tx fees.
type EventDistributeRevenue struct {
	// sender is the address of message sender.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// contract address in hex format
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// withdrawer_address is the bech32 address of account receiving the transaction fees
	WithdrawerAddress string `protobuf:"bytes,3,opt,name=withdrawer_address,json=withdrawerAddress,proto3" json:"withdrawer_address,omitempty"`
	// amount of revenue distributed
	Amount string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventDistributeRevenue) Reset()         { *m = EventDistributeRevenue{} }
func (m *EventDistributeRevenue) String() string { return proto.CompactTextString(m) }
func (*EventDistributeRevenue) ProtoMessage()    {}
func (*EventDistributeRevenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9a4bb0e3643eef, []int{3}
}
func (m *EventDistributeRevenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDistributeRevenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDistributeRevenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDistributeRevenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDistributeRevenue.Merge(m, src)
}
func (m *EventDistributeRevenue) XXX_Size() int {
	return m.Size()
}
func (m *EventDistributeRevenue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDistributeRevenue.DiscardUnknown(m)
}

var xxx_messageInfo_EventDistributeRevenue proto.InternalMessageInfo

func (m *EventDistributeRevenue) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventDistributeRevenue) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *EventDistributeRevenue) GetWithdrawerAddress() string {
	if m != nil {
		return m.WithdrawerAddress
	}
	return ""
}

func (m *EventDistributeRevenue) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRegisterRevenue)(nil), "sacas.revenue.v1.EventRegisterRevenue")
	proto.RegisterType((*EventUpdateRevenue)(nil), "sacas.revenue.v1.EventUpdateRevenue")
	proto.RegisterType((*EventCancelRevenue)(nil), "sacas.revenue.v1.EventCancelRevenue")
	proto.RegisterType((*EventDistributeRevenue)(nil), "sacas.revenue.v1.EventDistributeRevenue")
}

func init() { proto.RegisterFile("sacas/revenue/v1/events.proto", fileDescriptor_9e9a4bb0e3643eef) }

var fileDescriptor_9e9a4bb0e3643eef = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xbb, 0x7d, 0x5f, 0x8a, 0xee, 0xc5, 0x1a, 0x4b, 0x29, 0x82, 0x41, 0x72, 0xd2, 0x83,
	0x09, 0x41, 0xfd, 0x00, 0xfe, 0xbb, 0x7a, 0x28, 0x88, 0xe0, 0xa5, 0x6c, 0x93, 0x69, 0x1b, 0x6d,
	0x77, 0xc3, 0xee, 0x24, 0xb5, 0xdf, 0x42, 0xf0, 0xe4, 0xc9, 0xaf, 0xe3, 0xb1, 0x47, 0x8f, 0xd2,
	0x7e, 0x11, 0xc9, 0xee, 0x26, 0x2d, 0xd8, 0x8b, 0x07, 0x6f, 0x3b, 0xf3, 0xfc, 0xd8, 0xe7, 0x99,
	0x61, 0xe8, 0x81, 0x62, 0x11, 0x53, 0x81, 0x84, 0x1c, 0x78, 0x06, 0x41, 0x1e, 0x06, 0xc5, 0x0b,
	0x95, 0x9f, 0x4a, 0x81, 0xc2, 0x69, 0x6a, 0xd9, 0xb7, 0xb2, 0x9f, 0x87, 0xde, 0x3b, 0xa1, 0xad,
	0x9b, 0x02, 0xe9, 0xc2, 0x30, 0x51, 0x08, 0xb2, 0x6b, 0x34, 0xe7, 0x98, 0x36, 0x63, 0x48, 0xc7,
	0x62, 0x06, 0xb2, 0xc7, 0xe2, 0x58, 0x82, 0x52, 0x1d, 0x72, 0x48, 0x8e, 0xb6, 0xbb, 0x3b, 0x65,
	0xff, 0xc2, 0xb4, 0x0b, 0x34, 0x12, 0x1c, 0x25, 0x8b, 0xb0, 0x42, 0xeb, 0x06, 0x2d, 0xfb, 0x25,
	0x1a, 0xd2, 0x16, 0x0c, 0x06, 0x10, 0x61, 0x92, 0x43, 0x6f, 0x9a, 0xe0, 0x28, 0x96, 0x6c, 0x0a,
	0xb2, 0xf3, 0x4f, 0xe3, 0x7b, 0x95, 0x76, 0x5f, 0x49, 0xde, 0x1b, 0xa1, 0x8e, 0x4e, 0x78, 0x97,
	0xc6, 0x0c, 0x61, 0x2d, 0xdf, 0x0f, 0x53, 0xb2, 0xd9, 0x74, 0xd3, 0x28, 0xf5, 0xcd, 0xa3, 0x9c,
	0x50, 0x67, 0x95, 0xaa, 0x82, 0x4d, 0xba, 0xdd, 0x95, 0x62, 0x71, 0xef, 0xd1, 0x46, 0xbb, 0x62,
	0x3c, 0x82, 0xf1, 0x9f, 0xae, 0xce, 0x7b, 0x25, 0xb4, 0xad, 0xcd, 0xae, 0x13, 0x85, 0x32, 0xe9,
	0x67, 0xab, 0x5d, 0xb4, 0x69, 0x43, 0x01, 0x8f, 0x41, 0x5a, 0x1b, 0x5b, 0x39, 0xfb, 0x74, 0xab,
	0xfc, 0xc5, 0xfe, 0x5a, 0xd5, 0xbf, 0x9c, 0xb4, 0xb0, 0x60, 0x13, 0x91, 0x71, 0xec, 0xfc, 0x37,
	0x16, 0xa6, 0xba, 0xbc, 0xfd, 0x58, 0xb8, 0x64, 0xbe, 0x70, 0xc9, 0xd7, 0xc2, 0x25, 0x2f, 0x4b,
	0xb7, 0x36, 0x5f, 0xba, 0xb5, 0xcf, 0xa5, 0x5b, 0x7b, 0x38, 0x1b, 0x26, 0x38, 0xca, 0xfa, 0x7e,
	0x24, 0x26, 0x81, 0x3e, 0x3b, 0x0e, 0x38, 0x15, 0xf2, 0xc9, 0x14, 0x41, 0x1e, 0x9e, 0x07, 0xcf,
	0xeb, 0x87, 0x8a, 0xb3, 0x14, 0x54, 0xbf, 0xa1, 0x0f, 0xf5, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x9a, 0x70, 0x1e, 0x00, 0xc9, 0x02, 0x00, 0x00,
}

func (m *EventRegisterRevenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegisterRevenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegisterRevenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EffectiveWithdrawer) > 0 {
		i -= len(m.EffectiveWithdrawer)
		copy(dAtA[i:], m.EffectiveWithdrawer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EffectiveWithdrawer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeployerAddress) > 0 {
		i -= len(m.DeployerAddress)
		copy(dAtA[i:], m.DeployerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DeployerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateRevenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateRevenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateRevenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawerAddress) > 0 {
		i -= len(m.WithdrawerAddress)
		copy(dAtA[i:], m.WithdrawerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.WithdrawerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeployerAddress) > 0 {
		i -= len(m.DeployerAddress)
		copy(dAtA[i:], m.DeployerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DeployerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCancelRevenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCancelRevenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCancelRevenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeployerAddress) > 0 {
		i -= len(m.DeployerAddress)
		copy(dAtA[i:], m.DeployerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DeployerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventDistributeRevenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDistributeRevenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDistributeRevenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.WithdrawerAddress) > 0 {
		i -= len(m.WithdrawerAddress)
		copy(dAtA[i:], m.WithdrawerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.WithdrawerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
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
func (m *EventRegisterRevenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeployerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.EffectiveWithdrawer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateRevenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.DeployerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.WithdrawerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCancelRevenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeployerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventDistributeRevenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.WithdrawerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
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
func (m *EventRegisterRevenue) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventRegisterRevenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegisterRevenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
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
			m.DeployerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveWithdrawer", wireType)
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
			m.EffectiveWithdrawer = string(dAtA[iNdEx:postIndex])
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
func (m *EventUpdateRevenue) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventUpdateRevenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateRevenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
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
			m.DeployerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawerAddress", wireType)
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
			m.WithdrawerAddress = string(dAtA[iNdEx:postIndex])
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
func (m *EventCancelRevenue) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCancelRevenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCancelRevenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
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
			m.DeployerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *EventDistributeRevenue) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventDistributeRevenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDistributeRevenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawerAddress", wireType)
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
			m.WithdrawerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
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
