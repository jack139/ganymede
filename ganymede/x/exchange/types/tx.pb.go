// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ganymede/exchange/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgNewAsk struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Replier  string `protobuf:"bytes,3,opt,name=replier,proto3" json:"replier,omitempty"`
	Payload  string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	SentDate string `protobuf:"bytes,5,opt,name=sentDate,proto3" json:"sentDate,omitempty"`
}

func (m *MsgNewAsk) Reset()         { *m = MsgNewAsk{} }
func (m *MsgNewAsk) String() string { return proto.CompactTextString(m) }
func (*MsgNewAsk) ProtoMessage()    {}
func (*MsgNewAsk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ea3010a1b44fa7, []int{0}
}
func (m *MsgNewAsk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewAsk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewAsk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewAsk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewAsk.Merge(m, src)
}
func (m *MsgNewAsk) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewAsk) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewAsk.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewAsk proto.InternalMessageInfo

func (m *MsgNewAsk) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgNewAsk) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgNewAsk) GetReplier() string {
	if m != nil {
		return m.Replier
	}
	return ""
}

func (m *MsgNewAsk) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *MsgNewAsk) GetSentDate() string {
	if m != nil {
		return m.SentDate
	}
	return ""
}

type MsgNewAskResponse struct {
}

func (m *MsgNewAskResponse) Reset()         { *m = MsgNewAskResponse{} }
func (m *MsgNewAskResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewAskResponse) ProtoMessage()    {}
func (*MsgNewAskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ea3010a1b44fa7, []int{1}
}
func (m *MsgNewAskResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewAskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewAskResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewAskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewAskResponse.Merge(m, src)
}
func (m *MsgNewAskResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewAskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewAskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewAskResponse proto.InternalMessageInfo

type MsgNewReply struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	AskId    string `protobuf:"bytes,2,opt,name=askId,proto3" json:"askId,omitempty"`
	Sender   string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Replier  string `protobuf:"bytes,4,opt,name=replier,proto3" json:"replier,omitempty"`
	Payload  string `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	SentDate string `protobuf:"bytes,6,opt,name=sentDate,proto3" json:"sentDate,omitempty"`
}

func (m *MsgNewReply) Reset()         { *m = MsgNewReply{} }
func (m *MsgNewReply) String() string { return proto.CompactTextString(m) }
func (*MsgNewReply) ProtoMessage()    {}
func (*MsgNewReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ea3010a1b44fa7, []int{2}
}
func (m *MsgNewReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewReply.Merge(m, src)
}
func (m *MsgNewReply) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewReply.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewReply proto.InternalMessageInfo

func (m *MsgNewReply) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgNewReply) GetAskId() string {
	if m != nil {
		return m.AskId
	}
	return ""
}

func (m *MsgNewReply) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgNewReply) GetReplier() string {
	if m != nil {
		return m.Replier
	}
	return ""
}

func (m *MsgNewReply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *MsgNewReply) GetSentDate() string {
	if m != nil {
		return m.SentDate
	}
	return ""
}

type MsgNewReplyResponse struct {
}

func (m *MsgNewReplyResponse) Reset()         { *m = MsgNewReplyResponse{} }
func (m *MsgNewReplyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewReplyResponse) ProtoMessage()    {}
func (*MsgNewReplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ea3010a1b44fa7, []int{3}
}
func (m *MsgNewReplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewReplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewReplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewReplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewReplyResponse.Merge(m, src)
}
func (m *MsgNewReplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewReplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewReplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewReplyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgNewAsk)(nil), "ganymede.exchange.MsgNewAsk")
	proto.RegisterType((*MsgNewAskResponse)(nil), "ganymede.exchange.MsgNewAskResponse")
	proto.RegisterType((*MsgNewReply)(nil), "ganymede.exchange.MsgNewReply")
	proto.RegisterType((*MsgNewReplyResponse)(nil), "ganymede.exchange.MsgNewReplyResponse")
}

func init() { proto.RegisterFile("ganymede/exchange/tx.proto", fileDescriptor_e8ea3010a1b44fa7) }

var fileDescriptor_e8ea3010a1b44fa7 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x3b, 0x5f, 0xdb, 0x7c, 0xed, 0x71, 0xd5, 0x54, 0x25, 0x04, 0x19, 0xa4, 0x88, 0xb8,
	0x31, 0x01, 0xbd, 0x00, 0x51, 0xdc, 0x08, 0xd6, 0x45, 0x76, 0xba, 0x9b, 0x26, 0x87, 0x58, 0x9a,
	0x66, 0x86, 0x9c, 0x11, 0x1b, 0xef, 0x41, 0xf0, 0x26, 0xc4, 0x5b, 0x71, 0xd9, 0xa5, 0x4b, 0x69,
	0x6f, 0x44, 0x9a, 0x3f, 0x8a, 0x92, 0x2c, 0xdf, 0x3c, 0x9c, 0xc3, 0x93, 0x33, 0x2f, 0xd8, 0xa1,
	0x88, 0xd3, 0x39, 0x06, 0xe8, 0xe2, 0xc2, 0x7f, 0x14, 0x71, 0x88, 0xae, 0x5e, 0x38, 0x2a, 0x91,
	0x5a, 0x9a, 0x83, 0x92, 0x39, 0x25, 0x1b, 0xbd, 0x32, 0xe8, 0x8f, 0x29, 0xbc, 0xc3, 0xe7, 0x4b,
	0x9a, 0x99, 0x16, 0xfc, 0xf7, 0x13, 0x14, 0x5a, 0x26, 0x16, 0x3b, 0x64, 0x27, 0x7d, 0xaf, 0x8c,
	0xe6, 0x3e, 0x18, 0x84, 0x71, 0x80, 0x89, 0xf5, 0x2f, 0x03, 0x45, 0xda, 0x4c, 0x24, 0xa8, 0xa2,
	0x29, 0x26, 0x56, 0x3b, 0x9f, 0x28, 0xe2, 0x86, 0x28, 0x91, 0x46, 0x52, 0x04, 0x56, 0x27, 0x27,
	0x45, 0x34, 0x6d, 0xe8, 0x11, 0xc6, 0xfa, 0x5a, 0x68, 0xb4, 0xba, 0x19, 0xaa, 0xf2, 0x68, 0x08,
	0x83, 0x4a, 0xc7, 0x43, 0x52, 0x32, 0x26, 0x1c, 0xbd, 0x33, 0xd8, 0xc9, 0xbf, 0x7a, 0xa8, 0xa2,
	0xb4, 0x41, 0x73, 0x17, 0xba, 0x82, 0x66, 0x37, 0x41, 0x61, 0x99, 0x87, 0x2d, 0xf9, 0x76, 0x9d,
	0x7c, 0xa7, 0x56, 0xbe, 0x5b, 0x2f, 0x6f, 0xfc, 0x92, 0xdf, 0x83, 0xe1, 0x96, 0x66, 0xa9, 0x7f,
	0xf6, 0xc1, 0xa0, 0x3d, 0xa6, 0xd0, 0xbc, 0x05, 0xa3, 0xb8, 0xf3, 0x81, 0xf3, 0xe7, 0x25, 0x9c,
	0xea, 0xb7, 0xed, 0xa3, 0x26, 0x5a, 0x6e, 0x35, 0x3d, 0xe8, 0x55, 0x07, 0xe1, 0xb5, 0x13, 0x19,
	0xb7, 0x8f, 0x9b, 0x79, 0xb9, 0xf3, 0xea, 0xfe, 0x73, 0xc5, 0xd9, 0x72, 0xc5, 0xd9, 0xf7, 0x8a,
	0xb3, 0xb7, 0x35, 0x6f, 0x2d, 0xd7, 0xbc, 0xf5, 0xb5, 0xe6, 0xad, 0x87, 0x8b, 0x70, 0xaa, 0x23,
	0x31, 0x71, 0xd2, 0xe8, 0x45, 0x89, 0xd4, 0xf1, 0xe5, 0xdc, 0x0d, 0x9f, 0x44, 0xac, 0x85, 0x74,
	0x7d, 0x49, 0x73, 0x49, 0xa7, 0x1a, 0x49, 0xbb, 0x55, 0xf9, 0x16, 0x5b, 0xf5, 0x4b, 0x15, 0xd2,
	0xc4, 0xc8, 0x2a, 0x78, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x24, 0x3e, 0x02, 0xa0, 0x02,
	0x00, 0x00,
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
	NewAsk(ctx context.Context, in *MsgNewAsk, opts ...grpc.CallOption) (*MsgNewAskResponse, error)
	NewReply(ctx context.Context, in *MsgNewReply, opts ...grpc.CallOption) (*MsgNewReplyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) NewAsk(ctx context.Context, in *MsgNewAsk, opts ...grpc.CallOption) (*MsgNewAskResponse, error) {
	out := new(MsgNewAskResponse)
	err := c.cc.Invoke(ctx, "/ganymede.exchange.Msg/NewAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewReply(ctx context.Context, in *MsgNewReply, opts ...grpc.CallOption) (*MsgNewReplyResponse, error) {
	out := new(MsgNewReplyResponse)
	err := c.cc.Invoke(ctx, "/ganymede.exchange.Msg/NewReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	NewAsk(context.Context, *MsgNewAsk) (*MsgNewAskResponse, error)
	NewReply(context.Context, *MsgNewReply) (*MsgNewReplyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) NewAsk(ctx context.Context, req *MsgNewAsk) (*MsgNewAskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAsk not implemented")
}
func (*UnimplementedMsgServer) NewReply(ctx context.Context, req *MsgNewReply) (*MsgNewReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewReply not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_NewAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewAsk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymede.exchange.Msg/NewAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewAsk(ctx, req.(*MsgNewAsk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymede.exchange.Msg/NewReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewReply(ctx, req.(*MsgNewReply))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ganymede.exchange.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAsk",
			Handler:    _Msg_NewAsk_Handler,
		},
		{
			MethodName: "NewReply",
			Handler:    _Msg_NewReply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ganymede/exchange/tx.proto",
}

func (m *MsgNewAsk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewAsk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewAsk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SentDate) > 0 {
		i -= len(m.SentDate)
		copy(dAtA[i:], m.SentDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SentDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Replier) > 0 {
		i -= len(m.Replier)
		copy(dAtA[i:], m.Replier)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Replier)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
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

func (m *MsgNewAskResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewAskResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewAskResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgNewReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SentDate) > 0 {
		i -= len(m.SentDate)
		copy(dAtA[i:], m.SentDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SentDate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Replier) > 0 {
		i -= len(m.Replier)
		copy(dAtA[i:], m.Replier)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Replier)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AskId) > 0 {
		i -= len(m.AskId)
		copy(dAtA[i:], m.AskId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AskId)))
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

func (m *MsgNewReplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewReplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewReplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgNewAsk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Replier)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SentDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewAskResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgNewReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AskId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Replier)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SentDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewReplyResponse) Size() (n int) {
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
func (m *MsgNewAsk) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewAsk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewAsk: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replier", wireType)
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
			m.Replier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
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
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentDate", wireType)
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
			m.SentDate = string(dAtA[iNdEx:postIndex])
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
func (m *MsgNewAskResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewAskResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewAskResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgNewReply) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewReply: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field AskId", wireType)
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
			m.AskId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replier", wireType)
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
			m.Replier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
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
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentDate", wireType)
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
			m.SentDate = string(dAtA[iNdEx:postIndex])
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
func (m *MsgNewReplyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgNewReplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewReplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
