// APACHE NOTICE: This file has been modified from the original source
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/oracle/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgAggregateExchangeRatePrevote represents a message to submit
// aggregate exchange rate prevote.
type MsgAggregateExchangeRatePrevote struct {
	Hash      string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty" yaml:"hash"`
	Feeder    string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty" yaml:"feeder"`
	Validator string `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty" yaml:"validator"`
}

func (m *MsgAggregateExchangeRatePrevote) Reset()         { *m = MsgAggregateExchangeRatePrevote{} }
func (m *MsgAggregateExchangeRatePrevote) String() string { return proto.CompactTextString(m) }
func (*MsgAggregateExchangeRatePrevote) ProtoMessage()    {}
func (*MsgAggregateExchangeRatePrevote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{0}
}
func (m *MsgAggregateExchangeRatePrevote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAggregateExchangeRatePrevote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAggregateExchangeRatePrevote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAggregateExchangeRatePrevote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAggregateExchangeRatePrevote.Merge(m, src)
}
func (m *MsgAggregateExchangeRatePrevote) XXX_Size() int {
	return m.Size()
}
func (m *MsgAggregateExchangeRatePrevote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAggregateExchangeRatePrevote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAggregateExchangeRatePrevote proto.InternalMessageInfo

// MsgAggregateExchangeRatePrevoteResponse defines the Msg/AggregateExchangeRatePrevote response type.
type MsgAggregateExchangeRatePrevoteResponse struct {
}

func (m *MsgAggregateExchangeRatePrevoteResponse) Reset() {
	*m = MsgAggregateExchangeRatePrevoteResponse{}
}
func (m *MsgAggregateExchangeRatePrevoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAggregateExchangeRatePrevoteResponse) ProtoMessage()    {}
func (*MsgAggregateExchangeRatePrevoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{1}
}
func (m *MsgAggregateExchangeRatePrevoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAggregateExchangeRatePrevoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAggregateExchangeRatePrevoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAggregateExchangeRatePrevoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAggregateExchangeRatePrevoteResponse.Merge(m, src)
}
func (m *MsgAggregateExchangeRatePrevoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAggregateExchangeRatePrevoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAggregateExchangeRatePrevoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAggregateExchangeRatePrevoteResponse proto.InternalMessageInfo

// MsgAggregateExchangeRateVote represents a message to submit
// aggregate exchange rate vote.
type MsgAggregateExchangeRateVote struct {
	Salt          string `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty" yaml:"salt"`
	ExchangeRates string `protobuf:"bytes,2,opt,name=exchange_rates,json=exchangeRates,proto3" json:"exchange_rates,omitempty" yaml:"exchange_rates"`
	Feeder        string `protobuf:"bytes,3,opt,name=feeder,proto3" json:"feeder,omitempty" yaml:"feeder"`
	Validator     string `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty" yaml:"validator"`
}

func (m *MsgAggregateExchangeRateVote) Reset()         { *m = MsgAggregateExchangeRateVote{} }
func (m *MsgAggregateExchangeRateVote) String() string { return proto.CompactTextString(m) }
func (*MsgAggregateExchangeRateVote) ProtoMessage()    {}
func (*MsgAggregateExchangeRateVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{2}
}
func (m *MsgAggregateExchangeRateVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAggregateExchangeRateVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAggregateExchangeRateVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAggregateExchangeRateVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAggregateExchangeRateVote.Merge(m, src)
}
func (m *MsgAggregateExchangeRateVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgAggregateExchangeRateVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAggregateExchangeRateVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAggregateExchangeRateVote proto.InternalMessageInfo

// MsgAggregateExchangeRateVoteResponse defines the Msg/AggregateExchangeRateVote response type.
type MsgAggregateExchangeRateVoteResponse struct {
}

func (m *MsgAggregateExchangeRateVoteResponse) Reset()         { *m = MsgAggregateExchangeRateVoteResponse{} }
func (m *MsgAggregateExchangeRateVoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAggregateExchangeRateVoteResponse) ProtoMessage()    {}
func (*MsgAggregateExchangeRateVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{3}
}
func (m *MsgAggregateExchangeRateVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAggregateExchangeRateVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAggregateExchangeRateVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAggregateExchangeRateVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAggregateExchangeRateVoteResponse.Merge(m, src)
}
func (m *MsgAggregateExchangeRateVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAggregateExchangeRateVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAggregateExchangeRateVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAggregateExchangeRateVoteResponse proto.InternalMessageInfo

// MsgDelegateFeedConsent represents a message to
// delegate oracle voting rights to another address.
type MsgDelegateFeedConsent struct {
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty" yaml:"operator"`
	Delegate string `protobuf:"bytes,2,opt,name=delegate,proto3" json:"delegate,omitempty" yaml:"delegate"`
}

func (m *MsgDelegateFeedConsent) Reset()         { *m = MsgDelegateFeedConsent{} }
func (m *MsgDelegateFeedConsent) String() string { return proto.CompactTextString(m) }
func (*MsgDelegateFeedConsent) ProtoMessage()    {}
func (*MsgDelegateFeedConsent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{4}
}
func (m *MsgDelegateFeedConsent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDelegateFeedConsent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDelegateFeedConsent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDelegateFeedConsent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDelegateFeedConsent.Merge(m, src)
}
func (m *MsgDelegateFeedConsent) XXX_Size() int {
	return m.Size()
}
func (m *MsgDelegateFeedConsent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDelegateFeedConsent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDelegateFeedConsent proto.InternalMessageInfo

// MsgDelegateFeedConsentResponse defines the Msg/DelegateFeedConsent response type.
type MsgDelegateFeedConsentResponse struct {
}

func (m *MsgDelegateFeedConsentResponse) Reset()         { *m = MsgDelegateFeedConsentResponse{} }
func (m *MsgDelegateFeedConsentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDelegateFeedConsentResponse) ProtoMessage()    {}
func (*MsgDelegateFeedConsentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade38ec3545c6da7, []int{5}
}
func (m *MsgDelegateFeedConsentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDelegateFeedConsentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDelegateFeedConsentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDelegateFeedConsentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDelegateFeedConsentResponse.Merge(m, src)
}
func (m *MsgDelegateFeedConsentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDelegateFeedConsentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDelegateFeedConsentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDelegateFeedConsentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAggregateExchangeRatePrevote)(nil), "terra.oracle.v1beta1.MsgAggregateExchangeRatePrevote")
	proto.RegisterType((*MsgAggregateExchangeRatePrevoteResponse)(nil), "terra.oracle.v1beta1.MsgAggregateExchangeRatePrevoteResponse")
	proto.RegisterType((*MsgAggregateExchangeRateVote)(nil), "terra.oracle.v1beta1.MsgAggregateExchangeRateVote")
	proto.RegisterType((*MsgAggregateExchangeRateVoteResponse)(nil), "terra.oracle.v1beta1.MsgAggregateExchangeRateVoteResponse")
	proto.RegisterType((*MsgDelegateFeedConsent)(nil), "terra.oracle.v1beta1.MsgDelegateFeedConsent")
	proto.RegisterType((*MsgDelegateFeedConsentResponse)(nil), "terra.oracle.v1beta1.MsgDelegateFeedConsentResponse")
}

func init() { proto.RegisterFile("terra/oracle/v1beta1/tx.proto", fileDescriptor_ade38ec3545c6da7) }

var fileDescriptor_ade38ec3545c6da7 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0x87, 0xbd, 0xdc, 0xe9, 0x74, 0xb7, 0xe8, 0x38, 0x30, 0x01, 0x05, 0x04, 0xde, 0xd3, 0x82,
	0x80, 0x93, 0xc0, 0xd6, 0x41, 0x77, 0x15, 0x1c, 0x70, 0x5d, 0x24, 0xb4, 0x05, 0x05, 0x0d, 0xda,
	0xd8, 0xc3, 0x3a, 0x92, 0x93, 0xb5, 0x76, 0x97, 0x28, 0xe9, 0x29, 0x28, 0x79, 0x84, 0xbc, 0x01,
	0xaf, 0x41, 0x99, 0x92, 0xca, 0x42, 0x49, 0x43, 0x45, 0xe1, 0x27, 0x40, 0xde, 0xb5, 0x8d, 0x11,
	0x7f, 0x4e, 0xe9, 0xa2, 0xf9, 0xbe, 0xc9, 0xcc, 0xfc, 0xe4, 0xc5, 0xb7, 0x0d, 0x28, 0xc5, 0x23,
	0xa9, 0x78, 0x9c, 0x41, 0x34, 0x3d, 0x1e, 0x82, 0xe1, 0xc7, 0x91, 0x99, 0x85, 0xb9, 0x92, 0x46,
	0xfa, 0x3d, 0x8b, 0x43, 0x87, 0xc3, 0x1a, 0xdf, 0xec, 0x09, 0x29, 0xa4, 0x15, 0xa2, 0xea, 0x97,
	0x73, 0xe9, 0x67, 0x84, 0xc9, 0x40, 0x8b, 0x67, 0x42, 0x28, 0x10, 0xdc, 0xc0, 0xcb, 0x59, 0x9c,
	0xf2, 0x89, 0x00, 0xc6, 0x0d, 0xbc, 0x52, 0x30, 0x95, 0x06, 0xfc, 0x3b, 0x78, 0x3b, 0xe5, 0x3a,
	0xed, 0xa3, 0x43, 0xf4, 0x60, 0xef, 0xf4, 0xa0, 0x2c, 0xc8, 0xc5, 0x39, 0x1f, 0x67, 0x27, 0xb4,
	0xaa, 0x52, 0x66, 0xa1, 0x7f, 0x84, 0x77, 0xde, 0x01, 0x24, 0xa0, 0xfa, 0x17, 0xac, 0x76, 0xa5,
	0x2c, 0xc8, 0xbe, 0xd3, 0x5c, 0x9d, 0xb2, 0x5a, 0xf0, 0x1f, 0xe3, 0xbd, 0x29, 0xcf, 0x46, 0x09,
	0x37, 0x52, 0xf5, 0xb7, 0xac, 0xdd, 0x2b, 0x0b, 0x72, 0xd9, 0xd9, 0x2d, 0xa2, 0xec, 0x97, 0x76,
	0xb2, 0xfb, 0x71, 0x41, 0xbc, 0xef, 0x0b, 0xe2, 0xd1, 0x23, 0x7c, 0xff, 0x9c, 0x85, 0x19, 0xe8,
	0x5c, 0x4e, 0x34, 0xd0, 0x1f, 0x08, 0xdf, 0xfa, 0x97, 0xfb, 0xba, 0xbe, 0x4c, 0xf3, 0xcc, 0xfc,
	0x79, 0x59, 0x55, 0xa5, 0xcc, 0x42, 0xff, 0x29, 0xbe, 0x04, 0x75, 0xe3, 0x5b, 0xc5, 0x0d, 0xe8,
	0xfa, 0xc2, 0x1b, 0x65, 0x41, 0xae, 0x39, 0xfd, 0x77, 0x4e, 0xd9, 0x3e, 0x74, 0x26, 0xe9, 0x4e,
	0x36, 0x5b, 0x1b, 0x65, 0xb3, 0xbd, 0x69, 0x36, 0xf7, 0xf0, 0xdd, 0xff, 0xdd, 0xdb, 0x06, 0xf3,
	0x01, 0xe1, 0xeb, 0x03, 0x2d, 0x5e, 0x40, 0x66, 0xbd, 0x33, 0x80, 0xe4, 0x79, 0x05, 0x26, 0xc6,
	0x8f, 0xf0, 0xae, 0xcc, 0x41, 0xd9, 0xf9, 0x2e, 0x96, 0xab, 0x65, 0x41, 0x0e, 0xdc, 0xfc, 0x86,
	0x50, 0xd6, 0x4a, 0x55, 0x43, 0x52, 0xff, 0x4f, 0x1d, 0x4c, 0xa7, 0xa1, 0x21, 0x94, 0xb5, 0x52,
	0x67, 0xdd, 0x43, 0x1c, 0xfc, 0x7d, 0x8b, 0x66, 0xd1, 0xd3, 0xb3, 0x2f, 0xab, 0x00, 0x2d, 0x57,
	0x01, 0xfa, 0xb6, 0x0a, 0xd0, 0xa7, 0x75, 0xe0, 0x2d, 0xd7, 0x81, 0xf7, 0x75, 0x1d, 0x78, 0x6f,
	0x1e, 0x8a, 0x91, 0x49, 0xdf, 0x0f, 0xc3, 0x58, 0x8e, 0xa3, 0x38, 0xe3, 0x5a, 0x8f, 0xe2, 0x47,
	0xee, 0x59, 0xc4, 0x52, 0x41, 0x34, 0x6b, 0x5e, 0x87, 0x99, 0xe7, 0xa0, 0x87, 0x3b, 0xf6, 0x6b,
	0x7f, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xf3, 0x39, 0xc1, 0x3a, 0x03, 0x00, 0x00,
}

func (m *MsgAggregateExchangeRatePrevote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAggregateExchangeRatePrevote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAggregateExchangeRatePrevote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAggregateExchangeRatePrevoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAggregateExchangeRatePrevoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAggregateExchangeRatePrevoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAggregateExchangeRateVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAggregateExchangeRateVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAggregateExchangeRateVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ExchangeRates) > 0 {
		i -= len(m.ExchangeRates)
		copy(dAtA[i:], m.ExchangeRates)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ExchangeRates)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAggregateExchangeRateVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAggregateExchangeRateVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAggregateExchangeRateVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDelegateFeedConsent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDelegateFeedConsent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDelegateFeedConsent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegate) > 0 {
		i -= len(m.Delegate)
		copy(dAtA[i:], m.Delegate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Delegate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDelegateFeedConsentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDelegateFeedConsentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDelegateFeedConsentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAggregateExchangeRatePrevote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAggregateExchangeRatePrevoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAggregateExchangeRateVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ExchangeRates)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAggregateExchangeRateVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDelegateFeedConsent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Delegate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDelegateFeedConsentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAggregateExchangeRatePrevote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAggregateExchangeRatePrevote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAggregateExchangeRatePrevote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAggregateExchangeRatePrevoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAggregateExchangeRatePrevoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAggregateExchangeRatePrevoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAggregateExchangeRateVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAggregateExchangeRateVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAggregateExchangeRateVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeRates = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAggregateExchangeRateVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAggregateExchangeRateVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAggregateExchangeRateVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDelegateFeedConsent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDelegateFeedConsent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDelegateFeedConsent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDelegateFeedConsentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDelegateFeedConsentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDelegateFeedConsentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
