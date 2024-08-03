// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ganymede/postoffice/tx.proto

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

type MsgSendIbcPost struct {
	Title            string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Content          string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	SentDate         string `protobuf:"bytes,7,opt,name=sentDate,proto3" json:"sentDate,omitempty"`
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
}

func (m *MsgSendIbcPost) Reset()         { *m = MsgSendIbcPost{} }
func (m *MsgSendIbcPost) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcPost) ProtoMessage()    {}
func (*MsgSendIbcPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_a12f96ccd0840fe6, []int{0}
}
func (m *MsgSendIbcPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcPost.Merge(m, src)
}
func (m *MsgSendIbcPost) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcPost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcPost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcPost proto.InternalMessageInfo

func (m *MsgSendIbcPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgSendIbcPost) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSendIbcPost) GetSentDate() string {
	if m != nil {
		return m.SentDate
	}
	return ""
}

func (m *MsgSendIbcPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendIbcPost) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendIbcPost) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendIbcPost) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

type MsgSendIbcPostResponse struct {
}

func (m *MsgSendIbcPostResponse) Reset()         { *m = MsgSendIbcPostResponse{} }
func (m *MsgSendIbcPostResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcPostResponse) ProtoMessage()    {}
func (*MsgSendIbcPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a12f96ccd0840fe6, []int{1}
}
func (m *MsgSendIbcPostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcPostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcPostResponse.Merge(m, src)
}
func (m *MsgSendIbcPostResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcPostResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendIbcPost)(nil), "ganymede.postoffice.MsgSendIbcPost")
	proto.RegisterType((*MsgSendIbcPostResponse)(nil), "ganymede.postoffice.MsgSendIbcPostResponse")
}

func init() { proto.RegisterFile("ganymede/postoffice/tx.proto", fileDescriptor_a12f96ccd0840fe6) }

var fileDescriptor_a12f96ccd0840fe6 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x9b, 0xaf, 0xff, 0xbe, 0x46, 0x10, 0x89, 0x22, 0xa1, 0x94, 0xa1, 0xd4, 0x4d, 0x51,
	0x9c, 0x01, 0x7d, 0x02, 0xa5, 0x9b, 0x2e, 0x0a, 0x52, 0x5d, 0xe9, 0x42, 0xd2, 0xe9, 0xed, 0x38,
	0x30, 0xc9, 0x0d, 0x93, 0x5b, 0xe8, 0xf8, 0x14, 0x3e, 0x96, 0xcb, 0x2e, 0x5c, 0xb8, 0x94, 0xf6,
	0x45, 0x84, 0xd4, 0xb6, 0x16, 0x5d, 0xb8, 0xcb, 0x39, 0xbf, 0xdc, 0x84, 0x73, 0x0f, 0x6f, 0x25,
	0xca, 0x14, 0x1a, 0xc6, 0x10, 0x59, 0x74, 0x84, 0x93, 0x49, 0x1a, 0x43, 0x44, 0xb3, 0xd0, 0xe6,
	0x48, 0x28, 0x0e, 0xd7, 0x34, 0xdc, 0xd2, 0xce, 0x1b, 0xe3, 0xfb, 0x03, 0x97, 0xdc, 0x82, 0x19,
	0xf7, 0x47, 0xf1, 0x0d, 0x3a, 0x12, 0x47, 0xbc, 0x4a, 0x29, 0x65, 0x20, 0xab, 0x6d, 0xd6, 0x6d,
	0x0c, 0x57, 0x42, 0x48, 0x5e, 0x8f, 0xd1, 0x10, 0x18, 0x92, 0x35, 0xef, 0xaf, 0xa5, 0x68, 0xf2,
	0xff, 0x0e, 0x0c, 0xf5, 0x14, 0x81, 0xac, 0x7b, 0xb4, 0xd1, 0x7e, 0x2a, 0x07, 0x45, 0x98, 0x4b,
	0xf6, 0x35, 0xb5, 0x92, 0x42, 0xf0, 0x8a, 0xc5, 0x9c, 0xe4, 0x3f, 0x6f, 0xfb, 0xb3, 0x68, 0xf1,
	0x46, 0xfc, 0xa4, 0x8c, 0x81, 0xac, 0xdf, 0x93, 0x65, 0x0f, 0xb6, 0x86, 0x38, 0xe5, 0x07, 0x94,
	0x6a, 0xc0, 0x29, 0xdd, 0xa5, 0x1a, 0x1c, 0x29, 0x6d, 0x65, 0xa5, 0xcd, 0xba, 0x95, 0xe1, 0x0f,
	0xbf, 0x23, 0xf9, 0xf1, 0x6e, 0xaa, 0x21, 0x38, 0x8b, 0xc6, 0xc1, 0xc5, 0x84, 0x97, 0x07, 0x2e,
	0x11, 0x8f, 0x7c, 0xef, 0x7b, 0xe6, 0x93, 0xf0, 0x97, 0xe5, 0x84, 0xbb, 0x4f, 0x34, 0xcf, 0xfe,
	0x70, 0x69, 0xfd, 0xcf, 0xf5, 0xc3, 0xeb, 0x22, 0x60, 0xf3, 0x45, 0xc0, 0x3e, 0x16, 0x01, 0x7b,
	0x59, 0x06, 0xa5, 0xf9, 0x32, 0x28, 0xbd, 0x2f, 0x83, 0xd2, 0xfd, 0x55, 0x92, 0x52, 0xa6, 0x46,
	0x61, 0x91, 0x3d, 0x5b, 0x55, 0x84, 0x31, 0xea, 0x28, 0x99, 0x2a, 0x43, 0x0a, 0xa3, 0x18, 0x9d,
	0x46, 0x77, 0x4e, 0xe0, 0x28, 0xda, 0x74, 0x39, 0xdb, 0x69, 0xb3, 0xb0, 0xe0, 0x46, 0x35, 0xdf,
	0xe8, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x42, 0xaa, 0xd0, 0xf1, 0x01, 0x00, 0x00,
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
	SendIbcPost(ctx context.Context, in *MsgSendIbcPost, opts ...grpc.CallOption) (*MsgSendIbcPostResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendIbcPost(ctx context.Context, in *MsgSendIbcPost, opts ...grpc.CallOption) (*MsgSendIbcPostResponse, error) {
	out := new(MsgSendIbcPostResponse)
	err := c.cc.Invoke(ctx, "/ganymede.postoffice.Msg/SendIbcPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendIbcPost(context.Context, *MsgSendIbcPost) (*MsgSendIbcPostResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendIbcPost(ctx context.Context, req *MsgSendIbcPost) (*MsgSendIbcPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIbcPost not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendIbcPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendIbcPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendIbcPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ganymede.postoffice.Msg/SendIbcPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendIbcPost(ctx, req.(*MsgSendIbcPost))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ganymede.postoffice.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIbcPost",
			Handler:    _Msg_SendIbcPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ganymede/postoffice/tx.proto",
}

func (m *MsgSendIbcPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SentDate) > 0 {
		i -= len(m.SentDate)
		copy(dAtA[i:], m.SentDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SentDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
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

func (m *MsgSendIbcPostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcPostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcPostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSendIbcPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SentDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendIbcPostResponse) Size() (n int) {
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
func (m *MsgSendIbcPost) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendIbcPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcPost: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
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
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
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
func (m *MsgSendIbcPostResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendIbcPostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcPostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
