// APACHE NOTICE: This file has been modified from the original source
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/market/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// MsgSwap represents a message to swap coin to another denom.
type MsgSwap struct {
	Trader    string     `protobuf:"bytes,1,opt,name=trader,proto3" json:"trader,omitempty" yaml:"trader"`
	OfferCoin types.Coin `protobuf:"bytes,2,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin" yaml:"offer_coin"`
	AskDenom  string     `protobuf:"bytes,3,opt,name=ask_denom,json=askDenom,proto3" json:"ask_denom,omitempty" yaml:"ask_denom"`
}

func (m *MsgSwap) Reset()         { *m = MsgSwap{} }
func (m *MsgSwap) String() string { return proto.CompactTextString(m) }
func (*MsgSwap) ProtoMessage()    {}
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcd4b152743bd0f, []int{0}
}
func (m *MsgSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwap.Merge(m, src)
}
func (m *MsgSwap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwap proto.InternalMessageInfo

// MsgSwapResponse defines the Msg/Swap response type.
type MsgSwapResponse struct {
	SwapCoin types.Coin `protobuf:"bytes,1,opt,name=swap_coin,json=swapCoin,proto3" json:"swap_coin" yaml:"swap_coin"`
	SwapFee  types.Coin `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee" yaml:"swap_fee"`
}

func (m *MsgSwapResponse) Reset()         { *m = MsgSwapResponse{} }
func (m *MsgSwapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapResponse) ProtoMessage()    {}
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcd4b152743bd0f, []int{1}
}
func (m *MsgSwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapResponse.Merge(m, src)
}
func (m *MsgSwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapResponse proto.InternalMessageInfo

func (m *MsgSwapResponse) GetSwapCoin() types.Coin {
	if m != nil {
		return m.SwapCoin
	}
	return types.Coin{}
}

func (m *MsgSwapResponse) GetSwapFee() types.Coin {
	if m != nil {
		return m.SwapFee
	}
	return types.Coin{}
}

// MsgSwapSend represents a message to swap coin and send all result coin to recipient
type MsgSwapSend struct {
	FromAddress string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   string     `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	OfferCoin   types.Coin `protobuf:"bytes,3,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin" yaml:"offer_coin"`
	AskDenom    string     `protobuf:"bytes,4,opt,name=ask_denom,json=askDenom,proto3" json:"ask_denom,omitempty" yaml:"ask_denom"`
}

func (m *MsgSwapSend) Reset()         { *m = MsgSwapSend{} }
func (m *MsgSwapSend) String() string { return proto.CompactTextString(m) }
func (*MsgSwapSend) ProtoMessage()    {}
func (*MsgSwapSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcd4b152743bd0f, []int{2}
}
func (m *MsgSwapSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapSend.Merge(m, src)
}
func (m *MsgSwapSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapSend proto.InternalMessageInfo

// MsgSwapSendResponse defines the Msg/SwapSend response type.
type MsgSwapSendResponse struct {
	SwapCoin types.Coin `protobuf:"bytes,1,opt,name=swap_coin,json=swapCoin,proto3" json:"swap_coin" yaml:"swap_coin"`
	SwapFee  types.Coin `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee" yaml:"swap_fee"`
}

func (m *MsgSwapSendResponse) Reset()         { *m = MsgSwapSendResponse{} }
func (m *MsgSwapSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapSendResponse) ProtoMessage()    {}
func (*MsgSwapSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dcd4b152743bd0f, []int{3}
}
func (m *MsgSwapSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapSendResponse.Merge(m, src)
}
func (m *MsgSwapSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapSendResponse proto.InternalMessageInfo

func (m *MsgSwapSendResponse) GetSwapCoin() types.Coin {
	if m != nil {
		return m.SwapCoin
	}
	return types.Coin{}
}

func (m *MsgSwapSendResponse) GetSwapFee() types.Coin {
	if m != nil {
		return m.SwapFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgSwap)(nil), "terra.market.v1beta1.MsgSwap")
	proto.RegisterType((*MsgSwapResponse)(nil), "terra.market.v1beta1.MsgSwapResponse")
	proto.RegisterType((*MsgSwapSend)(nil), "terra.market.v1beta1.MsgSwapSend")
	proto.RegisterType((*MsgSwapSendResponse)(nil), "terra.market.v1beta1.MsgSwapSendResponse")
}

func init() { proto.RegisterFile("terra/market/v1beta1/tx.proto", fileDescriptor_7dcd4b152743bd0f) }

var fileDescriptor_7dcd4b152743bd0f = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xf6, 0xa6, 0xa8, 0x8d, 0x37, 0xa0, 0x52, 0x37, 0xa8, 0x69, 0x25, 0xec, 0x6a, 0x4f, 0x45,
	0x02, 0x5b, 0x01, 0x4e, 0xb9, 0x11, 0x50, 0x6f, 0x95, 0x90, 0x73, 0xe3, 0x12, 0x6d, 0xec, 0x71,
	0x88, 0x52, 0x7b, 0xac, 0xdd, 0x85, 0xb6, 0x6f, 0xc0, 0x11, 0xde, 0xa0, 0xcf, 0xc0, 0x81, 0x47,
	0x40, 0x3d, 0xf6, 0xc8, 0xc9, 0x42, 0xc9, 0x85, 0x73, 0x9e, 0x00, 0x79, 0x77, 0xf3, 0xc3, 0xa9,
	0xe2, 0xc0, 0x81, 0xdb, 0x8c, 0xbf, 0xf9, 0xbe, 0xf9, 0xbc, 0x9f, 0x86, 0x3e, 0x56, 0x20, 0x04,
	0x8f, 0x72, 0x2e, 0xa6, 0xa0, 0xa2, 0x8f, 0xdd, 0x11, 0x28, 0xde, 0x8d, 0xd4, 0x65, 0x58, 0x0a,
	0x54, 0xe8, 0xb5, 0x35, 0x1c, 0x1a, 0x38, 0xb4, 0xf0, 0x51, 0x7b, 0x8c, 0x63, 0xd4, 0x03, 0x51,
	0x5d, 0x99, 0xd9, 0x23, 0x3f, 0x41, 0x99, 0xa3, 0x8c, 0x46, 0x5c, 0xc2, 0x4a, 0x29, 0xc1, 0x49,
	0x61, 0x70, 0xf6, 0x9d, 0xd0, 0x9d, 0x33, 0x39, 0x1e, 0x5c, 0xf0, 0xd2, 0x7b, 0x42, 0xb7, 0x95,
	0xe0, 0x29, 0x88, 0x0e, 0x39, 0x26, 0x27, 0x6e, 0x7f, 0x6f, 0x51, 0x05, 0x0f, 0xae, 0x78, 0x7e,
	0xde, 0x63, 0xe6, 0x3b, 0x8b, 0xed, 0x80, 0x37, 0xa0, 0x14, 0xb3, 0x0c, 0xc4, 0xb0, 0x96, 0xea,
	0x34, 0x8e, 0xc9, 0x49, 0xeb, 0xf9, 0x61, 0x68, 0x76, 0x85, 0xf5, 0xae, 0xa5, 0xad, 0xf0, 0x35,
	0x4e, 0x8a, 0xfe, 0xe1, 0x4d, 0x15, 0x38, 0x8b, 0x2a, 0xd8, 0x33, 0x6a, 0x6b, 0x2a, 0x8b, 0x5d,
	0xdd, 0xd4, 0x53, 0x5e, 0x97, 0xba, 0x5c, 0x4e, 0x87, 0x29, 0x14, 0x98, 0x77, 0xb6, 0xb4, 0x85,
	0xf6, 0xa2, 0x0a, 0x1e, 0x1a, 0xd2, 0x0a, 0x62, 0x71, 0x93, 0xcb, 0xe9, 0x9b, 0xba, 0xec, 0x35,
	0x3f, 0x5d, 0x07, 0xce, 0xaf, 0xeb, 0xc0, 0x61, 0x5f, 0x09, 0xdd, 0xb5, 0x3f, 0x12, 0x83, 0x2c,
	0xb1, 0x90, 0xe0, 0xbd, 0xa5, 0xae, 0xbc, 0xe0, 0xa5, 0x31, 0x49, 0xee, 0x32, 0xd9, 0xb1, 0x26,
	0xed, 0xbe, 0x15, 0x93, 0xc5, 0xcd, 0xba, 0xd6, 0x16, 0xcf, 0xa8, 0xae, 0x87, 0x19, 0xc0, 0xdd,
	0x7f, 0x7d, 0x60, 0x05, 0x77, 0x37, 0x04, 0x33, 0x00, 0x16, 0xef, 0xd4, 0xe5, 0x29, 0x00, 0xfb,
	0xd2, 0xa0, 0x2d, 0x6b, 0x7a, 0x00, 0x45, 0xea, 0xf5, 0xe8, 0xfd, 0x4c, 0x60, 0x3e, 0xe4, 0x69,
	0x2a, 0x40, 0x4a, 0x9b, 0xc3, 0xc1, 0xa2, 0x0a, 0xf6, 0x8d, 0xc6, 0x26, 0xca, 0xe2, 0x56, 0xdd,
	0xbe, 0x32, 0x9d, 0xf7, 0x92, 0x52, 0x85, 0x2b, 0x66, 0x43, 0x33, 0x1f, 0xad, 0xdf, 0x7c, 0x8d,
	0xb1, 0xd8, 0x55, 0xb8, 0x64, 0xfd, 0x19, 0xe4, 0xd6, 0x3f, 0x08, 0xf2, 0xde, 0x5f, 0x06, 0xf9,
	0x8d, 0xd0, 0xfd, 0x8d, 0x37, 0xf9, 0x6f, 0xc2, 0xec, 0x9f, 0xde, 0xcc, 0x7c, 0x72, 0x3b, 0xf3,
	0xc9, 0xcf, 0x99, 0x4f, 0x3e, 0xcf, 0x7d, 0xe7, 0x76, 0xee, 0x3b, 0x3f, 0xe6, 0xbe, 0xf3, 0xee,
	0xe9, 0x78, 0xa2, 0xde, 0x7f, 0x18, 0x85, 0x09, 0xe6, 0x51, 0x72, 0xce, 0xa5, 0x9c, 0x24, 0xcf,
	0xcc, 0x89, 0x27, 0x28, 0x20, 0xba, 0x5c, 0x5e, 0xba, 0xba, 0x2a, 0x41, 0x8e, 0xb6, 0xf5, 0x65,
	0xbe, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x78, 0xeb, 0xe9, 0x06, 0x04, 0x00, 0x00,
}

func (m *MsgSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AskDenom) > 0 {
		i -= len(m.AskDenom)
		copy(dAtA[i:], m.AskDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AskDenom)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Trader) > 0 {
		i -= len(m.Trader)
		copy(dAtA[i:], m.Trader)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Trader)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SwapFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.SwapCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgSwapSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AskDenom) > 0 {
		i -= len(m.AskDenom)
		copy(dAtA[i:], m.AskDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AskDenom)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SwapFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.SwapCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *MsgSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Trader)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AskDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SwapFee.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AskDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapCoin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SwapFee.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trader", wireType)
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
			m.Trader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskDenom", wireType)
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
			m.AskDenom = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSwapResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSwapSend) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskDenom", wireType)
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
			m.AskDenom = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSwapSendResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
