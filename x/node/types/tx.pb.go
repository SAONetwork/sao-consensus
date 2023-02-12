// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgCreate struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgCreate) Reset()         { *m = MsgCreate{} }
func (m *MsgCreate) String() string { return proto.CompactTextString(m) }
func (*MsgCreate) ProtoMessage()    {}
func (*MsgCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{0}
}
func (m *MsgCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreate.Merge(m, src)
}
func (m *MsgCreate) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreate proto.InternalMessageInfo

func (m *MsgCreate) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgCreateResponse struct {
}

func (m *MsgCreateResponse) Reset()         { *m = MsgCreateResponse{} }
func (m *MsgCreateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResponse) ProtoMessage()    {}
func (*MsgCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{1}
}
func (m *MsgCreateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResponse.Merge(m, src)
}
func (m *MsgCreateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResponse proto.InternalMessageInfo

type MsgReset struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Peer    string `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	Status  uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *MsgReset) Reset()         { *m = MsgReset{} }
func (m *MsgReset) String() string { return proto.CompactTextString(m) }
func (*MsgReset) ProtoMessage()    {}
func (*MsgReset) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{2}
}
func (m *MsgReset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReset.Merge(m, src)
}
func (m *MsgReset) XXX_Size() int {
	return m.Size()
}
func (m *MsgReset) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReset.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReset proto.InternalMessageInfo

func (m *MsgReset) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgReset) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *MsgReset) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type MsgResetResponse struct {
}

func (m *MsgResetResponse) Reset()         { *m = MsgResetResponse{} }
func (m *MsgResetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgResetResponse) ProtoMessage()    {}
func (*MsgResetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{3}
}
func (m *MsgResetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgResetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgResetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgResetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgResetResponse.Merge(m, src)
}
func (m *MsgResetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgResetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgResetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgResetResponse proto.InternalMessageInfo

type MsgClaimReward struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgClaimReward) Reset()         { *m = MsgClaimReward{} }
func (m *MsgClaimReward) String() string { return proto.CompactTextString(m) }
func (*MsgClaimReward) ProtoMessage()    {}
func (*MsgClaimReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{4}
}
func (m *MsgClaimReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimReward.Merge(m, src)
}
func (m *MsgClaimReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimReward proto.InternalMessageInfo

func (m *MsgClaimReward) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgClaimRewardResponse struct {
	ClaimedReward uint64 `protobuf:"varint,1,opt,name=claimedReward,proto3" json:"claimedReward,omitempty"`
}

func (m *MsgClaimRewardResponse) Reset()         { *m = MsgClaimRewardResponse{} }
func (m *MsgClaimRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimRewardResponse) ProtoMessage()    {}
func (*MsgClaimRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaef138422732ffe, []int{5}
}
func (m *MsgClaimRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimRewardResponse.Merge(m, src)
}
func (m *MsgClaimRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimRewardResponse proto.InternalMessageInfo

func (m *MsgClaimRewardResponse) GetClaimedReward() uint64 {
	if m != nil {
		return m.ClaimedReward
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreate)(nil), "saonetwork.sao.node.MsgCreate")
	proto.RegisterType((*MsgCreateResponse)(nil), "saonetwork.sao.node.MsgCreateResponse")
	proto.RegisterType((*MsgReset)(nil), "saonetwork.sao.node.MsgReset")
	proto.RegisterType((*MsgResetResponse)(nil), "saonetwork.sao.node.MsgResetResponse")
	proto.RegisterType((*MsgClaimReward)(nil), "saonetwork.sao.node.MsgClaimReward")
	proto.RegisterType((*MsgClaimRewardResponse)(nil), "saonetwork.sao.node.MsgClaimRewardResponse")
}

func init() { proto.RegisterFile("sao/node/tx.proto", fileDescriptor_eaef138422732ffe) }

var fileDescriptor_eaef138422732ffe = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0x32, 0x51,
	0x14, 0x75, 0xd4, 0xcf, 0x2f, 0x6f, 0x18, 0xf9, 0x04, 0x19, 0x84, 0x1e, 0x32, 0xa5, 0x48, 0xc1,
	0x1b, 0xa8, 0x7d, 0x44, 0xad, 0x27, 0xe4, 0xb5, 0x6b, 0x13, 0x4f, 0xbd, 0x4c, 0x52, 0xfa, 0x86,
	0x77, 0x9f, 0x68, 0xbf, 0xa0, 0x6d, 0x3f, 0xab, 0xa5, 0xcb, 0x96, 0xa1, 0x7f, 0x24, 0x66, 0x74,
	0xa6, 0x84, 0x94, 0x76, 0xf7, 0xde, 0x73, 0xee, 0x39, 0x97, 0xc3, 0x85, 0x2a, 0x29, 0xed, 0x8f,
	0xf5, 0x00, 0x7d, 0x3b, 0x13, 0x91, 0xd1, 0x56, 0xb3, 0x1a, 0x29, 0x3d, 0x46, 0x3b, 0xd5, 0xe6,
	0x49, 0x90, 0xd2, 0x22, 0x46, 0xbd, 0x16, 0x94, 0x03, 0x0a, 0x6f, 0x0c, 0x2a, 0x8b, 0xcc, 0x85,
	0xff, 0xfd, 0xb8, 0xd2, 0xc6, 0x75, 0x9a, 0x4e, 0xa7, 0x2c, 0xd3, 0xd6, 0xab, 0x41, 0x35, 0xa3,
	0x49, 0xa4, 0x48, 0x8f, 0x09, 0xbd, 0x2e, 0xec, 0x05, 0x14, 0x4a, 0x24, 0xb4, 0xdb, 0x57, 0x19,
	0x83, 0x62, 0x84, 0x68, 0xdc, 0x7c, 0x32, 0x4e, 0x6a, 0x56, 0x87, 0x12, 0x59, 0x65, 0x27, 0xe4,
	0x16, 0x9a, 0x4e, 0xa7, 0x22, 0xd7, 0x9d, 0xc7, 0xe0, 0x30, 0x55, 0xcc, 0x5c, 0x4e, 0xe1, 0x20,
	0xb6, 0x7e, 0x56, 0xc3, 0x91, 0xc4, 0xa9, 0x32, 0x83, 0x1d, 0x67, 0x5e, 0x42, 0x7d, 0x93, 0x9b,
	0xaa, 0xb0, 0x13, 0xa8, 0xf4, 0xe3, 0x31, 0x0e, 0x56, 0x40, 0xb2, 0x59, 0x94, 0x9b, 0xc3, 0xf3,
	0xd7, 0x3c, 0x14, 0x02, 0x0a, 0x59, 0x17, 0x4a, 0xeb, 0x48, 0xb8, 0xf8, 0x25, 0x35, 0x91, 0x65,
	0xd1, 0x68, 0xef, 0xc6, 0x33, 0xff, 0x00, 0xfe, 0xad, 0x82, 0x3a, 0xda, 0xb6, 0x90, 0xc0, 0x8d,
	0xd6, 0x4e, 0x38, 0x93, 0x7b, 0x80, 0xfd, 0x9f, 0x89, 0x1c, 0x6f, 0xbd, 0xe2, 0x9b, 0xd4, 0x38,
	0xfb, 0x03, 0x29, 0x35, 0xb8, 0xbe, 0x7a, 0x5f, 0x70, 0x67, 0xbe, 0xe0, 0xce, 0xe7, 0x82, 0x3b,
	0x6f, 0x4b, 0x9e, 0x9b, 0x2f, 0x79, 0xee, 0x63, 0xc9, 0x73, 0xf7, 0xed, 0x70, 0x68, 0x1f, 0x27,
	0x3d, 0xd1, 0xd7, 0x23, 0xff, 0x4e, 0xe9, 0xdb, 0x95, 0xa0, 0x1f, 0xff, 0xdb, 0x6c, 0xfd, 0x71,
	0x2f, 0x11, 0x52, 0xaf, 0x94, 0x7c, 0xdd, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xaf,
	0xf2, 0x4a, 0x8a, 0x02, 0x00, 0x00,
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
	Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error)
	Reset(ctx context.Context, in *MsgReset, opts ...grpc.CallOption) (*MsgResetResponse, error)
	ClaimReward(ctx context.Context, in *MsgClaimReward, opts ...grpc.CallOption) (*MsgClaimRewardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error) {
	out := new(MsgCreateResponse)
	err := c.cc.Invoke(ctx, "/saonetwork.sao.node.Msg/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Reset(ctx context.Context, in *MsgReset, opts ...grpc.CallOption) (*MsgResetResponse, error) {
	out := new(MsgResetResponse)
	err := c.cc.Invoke(ctx, "/saonetwork.sao.node.Msg/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimReward(ctx context.Context, in *MsgClaimReward, opts ...grpc.CallOption) (*MsgClaimRewardResponse, error) {
	out := new(MsgClaimRewardResponse)
	err := c.cc.Invoke(ctx, "/saonetwork.sao.node.Msg/ClaimReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Create(context.Context, *MsgCreate) (*MsgCreateResponse, error)
	Reset(context.Context, *MsgReset) (*MsgResetResponse, error)
	ClaimReward(context.Context, *MsgClaimReward) (*MsgClaimRewardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Create(ctx context.Context, req *MsgCreate) (*MsgCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMsgServer) Reset(ctx context.Context, req *MsgReset) (*MsgResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (*UnimplementedMsgServer) ClaimReward(ctx context.Context, req *MsgClaimReward) (*MsgClaimRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saonetwork.sao.node.Msg/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Create(ctx, req.(*MsgCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saonetwork.sao.node.Msg/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Reset(ctx, req.(*MsgReset))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saonetwork.sao.node.Msg/ClaimReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimReward(ctx, req.(*MsgClaimReward))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "saonetwork.sao.node.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Msg_Create_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Msg_Reset_Handler,
		},
		{
			MethodName: "ClaimReward",
			Handler:    _Msg_ClaimReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sao/node/tx.proto",
}

func (m *MsgCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgReset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Peer) > 0 {
		i -= len(m.Peer)
		copy(dAtA[i:], m.Peer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Peer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgResetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgResetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgResetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClaimReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClaimedReward != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ClaimedReward))
		i--
		dAtA[i] = 0x8
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
func (m *MsgCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgReset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Peer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	return n
}

func (m *MsgResetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClaimReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClaimedReward != 0 {
		n += 1 + sovTx(uint64(m.ClaimedReward))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreate) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgReset) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgReset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
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
			m.Peer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgResetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgResetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgResetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgClaimReward) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
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
func (m *MsgClaimRewardResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgClaimRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedReward", wireType)
			}
			m.ClaimedReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimedReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
