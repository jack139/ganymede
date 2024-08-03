// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ganymede/postoffice/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the postoffice module's genesis state.
type GenesisState struct {
	Params            Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId            string         `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	PostList          []Post         `protobuf:"bytes,3,rep,name=postList,proto3" json:"postList"`
	PostCount         uint64         `protobuf:"varint,4,opt,name=postCount,proto3" json:"postCount,omitempty"`
	SentPostList      []SentPost     `protobuf:"bytes,5,rep,name=sentPostList,proto3" json:"sentPostList"`
	SentPostCount     uint64         `protobuf:"varint,6,opt,name=sentPostCount,proto3" json:"sentPostCount,omitempty"`
	TimedoutPostList  []TimedoutPost `protobuf:"bytes,7,rep,name=timedoutPostList,proto3" json:"timedoutPostList"`
	TimedoutPostCount uint64         `protobuf:"varint,8,opt,name=timedoutPostCount,proto3" json:"timedoutPostCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc98e9e732f96af, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetPostList() []Post {
	if m != nil {
		return m.PostList
	}
	return nil
}

func (m *GenesisState) GetPostCount() uint64 {
	if m != nil {
		return m.PostCount
	}
	return 0
}

func (m *GenesisState) GetSentPostList() []SentPost {
	if m != nil {
		return m.SentPostList
	}
	return nil
}

func (m *GenesisState) GetSentPostCount() uint64 {
	if m != nil {
		return m.SentPostCount
	}
	return 0
}

func (m *GenesisState) GetTimedoutPostList() []TimedoutPost {
	if m != nil {
		return m.TimedoutPostList
	}
	return nil
}

func (m *GenesisState) GetTimedoutPostCount() uint64 {
	if m != nil {
		return m.TimedoutPostCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ganymede.postoffice.GenesisState")
}

func init() { proto.RegisterFile("ganymede/postoffice/genesis.proto", fileDescriptor_cdc98e9e732f96af) }

var fileDescriptor_cdc98e9e732f96af = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0xdb, 0x0b, 0xb7, 0xc0, 0x80, 0x89, 0x8e, 0x26, 0x56, 0xd4, 0x5a, 0xd4, 0xc4, 0x2e,
	0xb4, 0x4d, 0x70, 0x65, 0x5c, 0x89, 0x0b, 0x62, 0xe2, 0x82, 0x80, 0x2b, 0x5d, 0x90, 0x81, 0x0e,
	0x4d, 0x13, 0xda, 0x69, 0x98, 0x43, 0x62, 0x7d, 0x0a, 0xdf, 0xc4, 0xd7, 0x60, 0xc9, 0xd2, 0x95,
	0x31, 0xf0, 0x22, 0xa6, 0xff, 0xa0, 0x84, 0xd9, 0xcd, 0xcc, 0xf9, 0xce, 0xef, 0xeb, 0x69, 0x0e,
	0x6a, 0x38, 0xc4, 0x0f, 0x3d, 0x6a, 0x53, 0x2b, 0x60, 0x1c, 0xd8, 0x68, 0xe4, 0x0e, 0xa9, 0xe5,
	0x50, 0x9f, 0x72, 0x97, 0x9b, 0xc1, 0x84, 0x01, 0xc3, 0xfb, 0x19, 0x62, 0xae, 0x91, 0xfa, 0x81,
	0xc3, 0x1c, 0x16, 0xd7, 0xad, 0xe8, 0x94, 0xa0, 0x75, 0x5d, 0x94, 0x16, 0x90, 0x09, 0xf1, 0xd2,
	0xb0, 0xba, 0x26, 0x24, 0x18, 0x87, 0xb4, 0x7e, 0x21, 0xaa, 0x73, 0xea, 0x43, 0x3f, 0x07, 0x5d,
	0x89, 0x20, 0x70, 0x3d, 0x6a, 0xb3, 0x69, 0x1e, 0x3c, 0xff, 0x2a, 0xa0, 0x5a, 0x3b, 0x19, 0xa6,
	0x07, 0x04, 0x28, 0xbe, 0x43, 0x4a, 0xf2, 0x39, 0xaa, 0xac, 0xcb, 0x46, 0xb5, 0x79, 0x6c, 0x0a,
	0x86, 0x33, 0x3b, 0x31, 0xd2, 0x2a, 0xce, 0x7e, 0xce, 0xa4, 0x6e, 0xda, 0x80, 0x0f, 0x51, 0x29,
	0x60, 0x13, 0xe8, 0xbb, 0xb6, 0xfa, 0x4f, 0x97, 0x8d, 0x4a, 0x57, 0x89, 0xae, 0x4f, 0x36, 0xbe,
	0x47, 0xe5, 0xa8, 0xf7, 0xd9, 0xe5, 0xa0, 0x16, 0xf4, 0x82, 0x51, 0x6d, 0x1e, 0x89, 0x53, 0x19,
	0x87, 0x34, 0x73, 0xd5, 0x80, 0x4f, 0x50, 0x25, 0x3a, 0x3f, 0xb2, 0xa9, 0x0f, 0x6a, 0x51, 0x97,
	0x8d, 0x62, 0x77, 0xfd, 0x80, 0xdb, 0xa8, 0x16, 0xcd, 0xde, 0xc9, 0xe2, 0xff, 0xc7, 0xf1, 0xa7,
	0xc2, 0xf8, 0x5e, 0x0a, 0xa6, 0x8a, 0x8d, 0x46, 0x7c, 0x89, 0x76, 0xb2, 0x7b, 0xa2, 0x52, 0x62,
	0xd5, 0xe6, 0x23, 0xee, 0xa1, 0xdd, 0xec, 0x2f, 0xae, 0x94, 0xa5, 0x58, 0xd9, 0x10, 0x2a, 0x5f,
	0x72, 0x70, 0xaa, 0xdd, 0x0a, 0xc0, 0xd7, 0x68, 0x2f, 0xff, 0x96, 0xe8, 0xcb, 0xb1, 0x7e, 0xbb,
	0xd0, 0x7a, 0x9b, 0x2d, 0x34, 0x79, 0xbe, 0xd0, 0xe4, 0xdf, 0x85, 0x26, 0x7f, 0x2e, 0x35, 0x69,
	0xbe, 0xd4, 0xa4, 0xef, 0xa5, 0x26, 0xbd, 0x3e, 0x38, 0x2e, 0x8c, 0xc9, 0xc0, 0x0c, 0xc7, 0x1f,
	0x01, 0x09, 0xcd, 0x21, 0xf3, 0x2c, 0x67, 0x4a, 0x7c, 0x20, 0xcc, 0x1a, 0x32, 0xee, 0x31, 0x7e,
	0x03, 0x94, 0x83, 0xb5, 0x5a, 0x8d, 0xf7, 0x8d, 0xe5, 0x08, 0x03, 0xca, 0x07, 0x4a, 0xbc, 0x15,
	0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xf9, 0x53, 0xbb, 0xf5, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimedoutPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TimedoutPostCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.TimedoutPostList) > 0 {
		for iNdEx := len(m.TimedoutPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimedoutPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SentPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SentPostCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SentPostList) > 0 {
		for iNdEx := len(m.SentPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SentPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PostList) > 0 {
		for iNdEx := len(m.PostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.PostList) > 0 {
		for _, e := range m.PostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PostCount != 0 {
		n += 1 + sovGenesis(uint64(m.PostCount))
	}
	if len(m.SentPostList) > 0 {
		for _, e := range m.SentPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SentPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.SentPostCount))
	}
	if len(m.TimedoutPostList) > 0 {
		for _, e := range m.TimedoutPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TimedoutPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.TimedoutPostCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostList = append(m.PostList, Post{})
			if err := m.PostList[len(m.PostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostCount", wireType)
			}
			m.PostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentPostList = append(m.SentPostList, SentPost{})
			if err := m.SentPostList[len(m.SentPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostCount", wireType)
			}
			m.SentPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SentPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutPostList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimedoutPostList = append(m.TimedoutPostList, TimedoutPost{})
			if err := m.TimedoutPostList[len(m.TimedoutPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutPostCount", wireType)
			}
			m.TimedoutPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimedoutPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
