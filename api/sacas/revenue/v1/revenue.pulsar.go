// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package revenuev1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Revenue                    protoreflect.MessageDescriptor
	fd_Revenue_contract_address   protoreflect.FieldDescriptor
	fd_Revenue_deployer_address   protoreflect.FieldDescriptor
	fd_Revenue_withdrawer_address protoreflect.FieldDescriptor
)

func init() {
	file_sacas_revenue_v1_revenue_proto_init()
	md_Revenue = File_sacas_revenue_v1_revenue_proto.Messages().ByName("Revenue")
	fd_Revenue_contract_address = md_Revenue.Fields().ByName("contract_address")
	fd_Revenue_deployer_address = md_Revenue.Fields().ByName("deployer_address")
	fd_Revenue_withdrawer_address = md_Revenue.Fields().ByName("withdrawer_address")
}

var _ protoreflect.Message = (*fastReflection_Revenue)(nil)

type fastReflection_Revenue Revenue

func (x *Revenue) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Revenue)(x)
}

func (x *Revenue) slowProtoReflect() protoreflect.Message {
	mi := &file_sacas_revenue_v1_revenue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Revenue_messageType fastReflection_Revenue_messageType
var _ protoreflect.MessageType = fastReflection_Revenue_messageType{}

type fastReflection_Revenue_messageType struct{}

func (x fastReflection_Revenue_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Revenue)(nil)
}
func (x fastReflection_Revenue_messageType) New() protoreflect.Message {
	return new(fastReflection_Revenue)
}
func (x fastReflection_Revenue_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Revenue
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Revenue) Descriptor() protoreflect.MessageDescriptor {
	return md_Revenue
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Revenue) Type() protoreflect.MessageType {
	return _fastReflection_Revenue_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Revenue) New() protoreflect.Message {
	return new(fastReflection_Revenue)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Revenue) Interface() protoreflect.ProtoMessage {
	return (*Revenue)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Revenue) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ContractAddress != "" {
		value := protoreflect.ValueOfString(x.ContractAddress)
		if !f(fd_Revenue_contract_address, value) {
			return
		}
	}
	if x.DeployerAddress != "" {
		value := protoreflect.ValueOfString(x.DeployerAddress)
		if !f(fd_Revenue_deployer_address, value) {
			return
		}
	}
	if x.WithdrawerAddress != "" {
		value := protoreflect.ValueOfString(x.WithdrawerAddress)
		if !f(fd_Revenue_withdrawer_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Revenue) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		return x.ContractAddress != ""
	case "sacas.revenue.v1.Revenue.deployer_address":
		return x.DeployerAddress != ""
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		return x.WithdrawerAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Revenue) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		x.ContractAddress = ""
	case "sacas.revenue.v1.Revenue.deployer_address":
		x.DeployerAddress = ""
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		x.WithdrawerAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Revenue) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		value := x.ContractAddress
		return protoreflect.ValueOfString(value)
	case "sacas.revenue.v1.Revenue.deployer_address":
		value := x.DeployerAddress
		return protoreflect.ValueOfString(value)
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		value := x.WithdrawerAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Revenue) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		x.ContractAddress = value.Interface().(string)
	case "sacas.revenue.v1.Revenue.deployer_address":
		x.DeployerAddress = value.Interface().(string)
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		x.WithdrawerAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Revenue) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		panic(fmt.Errorf("field contract_address of message sacas.revenue.v1.Revenue is not mutable"))
	case "sacas.revenue.v1.Revenue.deployer_address":
		panic(fmt.Errorf("field deployer_address of message sacas.revenue.v1.Revenue is not mutable"))
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		panic(fmt.Errorf("field withdrawer_address of message sacas.revenue.v1.Revenue is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Revenue) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sacas.revenue.v1.Revenue.contract_address":
		return protoreflect.ValueOfString("")
	case "sacas.revenue.v1.Revenue.deployer_address":
		return protoreflect.ValueOfString("")
	case "sacas.revenue.v1.Revenue.withdrawer_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sacas.revenue.v1.Revenue"))
		}
		panic(fmt.Errorf("message sacas.revenue.v1.Revenue does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Revenue) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sacas.revenue.v1.Revenue", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Revenue) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Revenue) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Revenue) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Revenue) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Revenue)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ContractAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DeployerAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.WithdrawerAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Revenue)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.WithdrawerAddress) > 0 {
			i -= len(x.WithdrawerAddress)
			copy(dAtA[i:], x.WithdrawerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.WithdrawerAddress)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.DeployerAddress) > 0 {
			i -= len(x.DeployerAddress)
			copy(dAtA[i:], x.DeployerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DeployerAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ContractAddress) > 0 {
			i -= len(x.ContractAddress)
			copy(dAtA[i:], x.ContractAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ContractAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Revenue)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Revenue: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Revenue: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ContractAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DeployerAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field WithdrawerAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.WithdrawerAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// SPDX-License-Identifier:LGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sacas/revenue/v1/revenue.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Revenue defines an instance that organizes fee distribution conditions for
// the owner of a given smart contract
type Revenue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// contract_address is the hex address of a registered contract
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// deployer_address is the bech32 address of message sender. It must be the same as the origin EOA
	// sending the transaction which deploys the contract
	DeployerAddress string `protobuf:"bytes,2,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
	// withdrawer_address is the bech32 address of account receiving the transaction fees it defaults to
	// deployer_address
	WithdrawerAddress string `protobuf:"bytes,3,opt,name=withdrawer_address,json=withdrawerAddress,proto3" json:"withdrawer_address,omitempty"`
}

func (x *Revenue) Reset() {
	*x = Revenue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sacas_revenue_v1_revenue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revenue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revenue) ProtoMessage() {}

// Deprecated: Use Revenue.ProtoReflect.Descriptor instead.
func (*Revenue) Descriptor() ([]byte, []int) {
	return file_sacas_revenue_v1_revenue_proto_rawDescGZIP(), []int{0}
}

func (x *Revenue) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *Revenue) GetDeployerAddress() string {
	if x != nil {
		return x.DeployerAddress
	}
	return ""
}

func (x *Revenue) GetWithdrawerAddress() string {
	if x != nil {
		return x.WithdrawerAddress
	}
	return ""
}

var File_sacas_revenue_v1_revenue_proto protoreflect.FileDescriptor

var file_sacas_revenue_v1_revenue_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x61, 0x63, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x61, 0x63, 0x61, 0x73, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0xb3, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x63, 0x61,
	0x73, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x52, 0x65,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x61, 0x63, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x52, 0x58, 0xaa,
	0x02, 0x10, 0x53, 0x61, 0x63, 0x61, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x10, 0x53, 0x61, 0x63, 0x61, 0x73, 0x5c, 0x52, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x53, 0x61, 0x63, 0x61, 0x73, 0x5c, 0x52, 0x65,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x61, 0x63, 0x61, 0x73, 0x3a, 0x3a, 0x52, 0x65,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sacas_revenue_v1_revenue_proto_rawDescOnce sync.Once
	file_sacas_revenue_v1_revenue_proto_rawDescData = file_sacas_revenue_v1_revenue_proto_rawDesc
)

func file_sacas_revenue_v1_revenue_proto_rawDescGZIP() []byte {
	file_sacas_revenue_v1_revenue_proto_rawDescOnce.Do(func() {
		file_sacas_revenue_v1_revenue_proto_rawDescData = protoimpl.X.CompressGZIP(file_sacas_revenue_v1_revenue_proto_rawDescData)
	})
	return file_sacas_revenue_v1_revenue_proto_rawDescData
}

var file_sacas_revenue_v1_revenue_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sacas_revenue_v1_revenue_proto_goTypes = []interface{}{
	(*Revenue)(nil), // 0: sacas.revenue.v1.Revenue
}
var file_sacas_revenue_v1_revenue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sacas_revenue_v1_revenue_proto_init() }
func file_sacas_revenue_v1_revenue_proto_init() {
	if File_sacas_revenue_v1_revenue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sacas_revenue_v1_revenue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revenue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sacas_revenue_v1_revenue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sacas_revenue_v1_revenue_proto_goTypes,
		DependencyIndexes: file_sacas_revenue_v1_revenue_proto_depIdxs,
		MessageInfos:      file_sacas_revenue_v1_revenue_proto_msgTypes,
	}.Build()
	File_sacas_revenue_v1_revenue_proto = out.File
	file_sacas_revenue_v1_revenue_proto_rawDesc = nil
	file_sacas_revenue_v1_revenue_proto_goTypes = nil
	file_sacas_revenue_v1_revenue_proto_depIdxs = nil
}
