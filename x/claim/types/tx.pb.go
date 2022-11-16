// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eve/claim/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgInitialClaim struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgInitialClaim) Reset()         { *m = MsgInitialClaim{} }
func (m *MsgInitialClaim) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaim) ProtoMessage()    {}
func (*MsgInitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d82cd90558068c0, []int{0}
}
func (m *MsgInitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaim.Merge(m, src)
}
func (m *MsgInitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaim proto.InternalMessageInfo

func (m *MsgInitialClaim) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgInitialClaimResponse struct {
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgInitialClaimResponse) Reset()         { *m = MsgInitialClaimResponse{} }
func (m *MsgInitialClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaimResponse) ProtoMessage()    {}
func (*MsgInitialClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d82cd90558068c0, []int{1}
}
func (m *MsgInitialClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaimResponse.Merge(m, src)
}
func (m *MsgInitialClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaimResponse proto.InternalMessageInfo

func (m *MsgInitialClaimResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

type MsgClaim struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgClaim) Reset()         { *m = MsgClaim{} }
func (m *MsgClaim) String() string { return proto.CompactTextString(m) }
func (*MsgClaim) ProtoMessage()    {}
func (*MsgClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d82cd90558068c0, []int{2}
}
func (m *MsgClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaim.Merge(m, src)
}
func (m *MsgClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaim proto.InternalMessageInfo

func (m *MsgClaim) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgClaim) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgClaimResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgClaimResponse) Reset()         { *m = MsgClaimResponse{} }
func (m *MsgClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimResponse) ProtoMessage()    {}
func (*MsgClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d82cd90558068c0, []int{3}
}
func (m *MsgClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimResponse.Merge(m, src)
}
func (m *MsgClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimResponse proto.InternalMessageInfo

func (m *MsgClaimResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInitialClaim)(nil), "evenetwork.eve.claim.v1beta1.MsgInitialClaim")
	proto.RegisterType((*MsgInitialClaimResponse)(nil), "evenetwork.eve.claim.v1beta1.MsgInitialClaimResponse")
	proto.RegisterType((*MsgClaim)(nil), "evenetwork.eve.claim.v1beta1.MsgClaim")
	proto.RegisterType((*MsgClaimResponse)(nil), "evenetwork.eve.claim.v1beta1.MsgClaimResponse")
}

func init() { proto.RegisterFile("eve/claim/v1beta1/tx.proto", fileDescriptor_8d82cd90558068c0) }

var fileDescriptor_8d82cd90558068c0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2d, 0x4b, 0xd5,
	0x4f, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0x2d, 0x4b, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4b, 0x2d, 0x4b, 0xd5, 0x03, 0x2b, 0xd3, 0x83, 0x2a, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x2b, 0xd4, 0x07, 0xb1, 0x20, 0x7a, 0xa4, 0xe4, 0x92, 0xf3, 0x8b, 0x73, 0xf3,
	0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0xe1, 0x26, 0x26, 0xe7, 0x67, 0xe6, 0x41, 0xe5, 0x55, 0x30,
	0xed, 0x03, 0xf3, 0xe2, 0x8b, 0x52, 0x93, 0xf3, 0x8b, 0x52, 0x20, 0xaa, 0x94, 0x34, 0xb9, 0xf8,
	0x7d, 0x8b, 0xd3, 0x3d, 0xf3, 0x32, 0x4b, 0x32, 0x13, 0x73, 0x9c, 0x41, 0xf2, 0x42, 0x62, 0x5c,
	0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e,
	0xd2, 0x72, 0x46, 0x2e, 0x71, 0x34, 0xb5, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0xdd, 0x8c, 0x5c, 0x7c, 0x60, 0xd3, 0x53, 0x53, 0xe2, 0x13, 0x73, 0xf3, 0x4b, 0xf3, 0x4a, 0x24,
	0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0x20, 0xce, 0xd4, 0x03, 0x39, 0x13, 0xe6, 0x23,
	0x3d, 0xe7, 0xfc, 0xcc, 0x3c, 0x27, 0xcf, 0x13, 0xf7, 0xe4, 0x19, 0x3e, 0xdd, 0x93, 0x17, 0xad,
	0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0x42, 0xd5, 0xae, 0xb4, 0xea, 0xbe, 0xbc, 0x46, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0xb3, 0x10, 0x4a, 0xb7, 0x38, 0x25, 0x5b,
	0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0x6c, 0x52, 0x71, 0x10, 0x2f, 0x54, 0xb3, 0x23, 0x44, 0xaf,
	0x0d, 0x17, 0x87, 0x6f, 0x71, 0x3a, 0x5e, 0xdf, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14,
	0xa5, 0x16, 0x17, 0x4b, 0x30, 0x81, 0x25, 0x60, 0x5c, 0xa5, 0x5d, 0x8c, 0x5c, 0x02, 0x30, 0xed,
	0x70, 0x0f, 0x22, 0x29, 0x67, 0x44, 0x51, 0x3e, 0xb8, 0xbc, 0x6e, 0x74, 0x87, 0x91, 0x8b, 0xd9,
	0xb7, 0x38, 0x5d, 0xa8, 0x84, 0x8b, 0x07, 0x25, 0x52, 0x75, 0xf5, 0xf0, 0x25, 0x31, 0x3d, 0xb4,
	0x78, 0x95, 0x32, 0x25, 0x49, 0x39, 0x3c, 0x94, 0xe2, 0xb9, 0x58, 0x21, 0xd6, 0xa9, 0x11, 0xd4,
	0x0f, 0xb1, 0x47, 0x8f, 0x38, 0x75, 0x30, 0x0b, 0x9c, 0x9c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0x0a, 0x39, 0xc4, 0x52, 0xcb, 0x52, 0x75, 0xa1, 0x86, 0x82, 0xd8, 0xfa,
	0x15, 0xd0, 0x7c, 0x00, 0x0e, 0xb7, 0x24, 0x36, 0x70, 0xca, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xaf, 0xff, 0xb2, 0x91, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error) {
	out := new(MsgInitialClaimResponse)
	err := c.cc.Invoke(ctx, "/evenetwork.eve.claim.v1beta1.Msg/InitialClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error) {
	out := new(MsgClaimResponse)
	err := c.cc.Invoke(ctx, "/evenetwork.eve.claim.v1beta1.Msg/Claim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InitialClaim(context.Context, *MsgInitialClaim) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	Claim(context.Context, *MsgClaim) (*MsgClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InitialClaim(ctx context.Context, req *MsgInitialClaim) (*MsgInitialClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialClaim not implemented")
}
func (*UnimplementedMsgServer) Claim(ctx context.Context, req *MsgClaim) (*MsgClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InitialClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitialClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitialClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evenetwork.eve.claim.v1beta1.Msg/InitialClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitialClaim(ctx, req.(*MsgInitialClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evenetwork.eve.claim.v1beta1.Msg/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Claim(ctx, req.(*MsgClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evenetwork.eve.claim.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitialClaim",
			Handler:    _Msg_InitialClaim_Handler,
		},
		{
			MethodName: "Claim",
			Handler:    _Msg_Claim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eve/claim/v1beta1/tx.proto",
}

func (m *MsgInitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitialClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgInitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitialClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitialClaim) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgInitialClaimResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInitialClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
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
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgClaim) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *MsgClaimResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
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
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
