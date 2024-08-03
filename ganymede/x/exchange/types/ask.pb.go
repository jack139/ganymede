// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ganymede/exchange/ask.proto

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

type Ask struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Replier     string `protobuf:"bytes,3,opt,name=replier,proto3" json:"replier,omitempty"`
	Payload     string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	SentDate    string `protobuf:"bytes,5,opt,name=sentDate,proto3" json:"sentDate,omitempty"`
	Creator     string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	LinkSender  string `protobuf:"bytes,7,opt,name=linkSender,proto3" json:"linkSender,omitempty"`
	LinkReplier string `protobuf:"bytes,8,opt,name=linkReplier,proto3" json:"linkReplier,omitempty"`
}

func (m *Ask) Reset()         { *m = Ask{} }
func (m *Ask) String() string { return proto.CompactTextString(m) }
func (*Ask) ProtoMessage()    {}
func (*Ask) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c151aa70b29d69, []int{0}
}
func (m *Ask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ask.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ask.Merge(m, src)
}
func (m *Ask) XXX_Size() int {
	return m.Size()
}
func (m *Ask) XXX_DiscardUnknown() {
	xxx_messageInfo_Ask.DiscardUnknown(m)
}

var xxx_messageInfo_Ask proto.InternalMessageInfo

func (m *Ask) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Ask) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Ask) GetReplier() string {
	if m != nil {
		return m.Replier
	}
	return ""
}

func (m *Ask) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Ask) GetSentDate() string {
	if m != nil {
		return m.SentDate
	}
	return ""
}

func (m *Ask) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Ask) GetLinkSender() string {
	if m != nil {
		return m.LinkSender
	}
	return ""
}

func (m *Ask) GetLinkReplier() string {
	if m != nil {
		return m.LinkReplier
	}
	return ""
}

func init() {
	proto.RegisterType((*Ask)(nil), "ganymede.exchange.Ask")
}

func init() { proto.RegisterFile("ganymede/exchange/ask.proto", fileDescriptor_f7c151aa70b29d69) }

var fileDescriptor_f7c151aa70b29d69 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0x84, 0x30,
	0x14, 0x86, 0x29, 0x33, 0x32, 0x63, 0x4d, 0x4c, 0xec, 0xc2, 0x34, 0x9a, 0x34, 0xc4, 0xd5, 0x6c,
	0x84, 0x85, 0x07, 0x30, 0x1a, 0x4f, 0x80, 0x2b, 0xdd, 0xbd, 0x81, 0x17, 0x24, 0x40, 0x4b, 0xda,
	0x9a, 0x0c, 0x9e, 0xc2, 0x63, 0xb9, 0x9c, 0xa5, 0xee, 0x0c, 0x5c, 0xc4, 0x50, 0x60, 0x9c, 0xe5,
	0xf7, 0xbe, 0xf7, 0x27, 0x7f, 0x7e, 0x7a, 0x9d, 0x83, 0x6c, 0x6b, 0xcc, 0x30, 0xc6, 0x5d, 0xfa,
	0x06, 0x32, 0xc7, 0x18, 0x4c, 0x19, 0x35, 0x5a, 0x59, 0xc5, 0x2e, 0x66, 0x19, 0xcd, 0xf2, 0xe6,
	0x87, 0xd0, 0xc5, 0x83, 0x29, 0xd9, 0x39, 0xf5, 0x8b, 0x8c, 0x93, 0x90, 0x6c, 0x96, 0x89, 0x5f,
	0x64, 0xec, 0x92, 0x06, 0x06, 0x65, 0x86, 0x9a, 0xfb, 0x21, 0xd9, 0x9c, 0x26, 0x13, 0x31, 0x4e,
	0x57, 0x1a, 0x9b, 0xaa, 0x40, 0xcd, 0x17, 0x4e, 0xcc, 0x38, 0x98, 0x06, 0xda, 0x4a, 0x41, 0xc6,
	0x97, 0xa3, 0x99, 0x90, 0x5d, 0xd1, 0xb5, 0x41, 0x69, 0x9f, 0xc0, 0x22, 0x3f, 0x71, 0xea, 0xc0,
	0x43, 0x2a, 0xd5, 0x08, 0x56, 0x69, 0x1e, 0x8c, 0xa9, 0x09, 0x99, 0xa0, 0xb4, 0x2a, 0x64, 0xf9,
	0x3c, 0xb6, 0x58, 0x39, 0x79, 0x74, 0x61, 0x21, 0x3d, 0x1b, 0x28, 0x99, 0xda, 0xac, 0xdd, 0xc3,
	0xf1, 0xe9, 0xf1, 0xe5, 0xab, 0x13, 0x64, 0xdf, 0x09, 0xf2, 0xdb, 0x09, 0xf2, 0xd9, 0x0b, 0x6f,
	0xdf, 0x0b, 0xef, 0xbb, 0x17, 0xde, 0xeb, 0x7d, 0x5e, 0xd8, 0x0a, 0xb6, 0x51, 0x5b, 0x7d, 0x34,
	0xd0, 0x46, 0xa9, 0xaa, 0xe3, 0xfc, 0x1d, 0xa4, 0x05, 0x15, 0xa7, 0xca, 0xd4, 0xca, 0xdc, 0x5a,
	0x34, 0x36, 0x3e, 0x6c, 0xb9, 0xfb, 0x5f, 0xd3, 0xb6, 0x0d, 0x9a, 0x6d, 0xe0, 0x06, 0xbd, 0xfb,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xeb, 0xb4, 0x0b, 0x6f, 0x01, 0x00, 0x00,
}

func (m *Ask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LinkReplier) > 0 {
		i -= len(m.LinkReplier)
		copy(dAtA[i:], m.LinkReplier)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.LinkReplier)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.LinkSender) > 0 {
		i -= len(m.LinkSender)
		copy(dAtA[i:], m.LinkSender)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.LinkSender)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SentDate) > 0 {
		i -= len(m.SentDate)
		copy(dAtA[i:], m.SentDate)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.SentDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Replier) > 0 {
		i -= len(m.Replier)
		copy(dAtA[i:], m.Replier)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.Replier)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintAsk(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAsk(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsk(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAsk(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.Replier)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.SentDate)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.LinkSender)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	l = len(m.LinkReplier)
	if l > 0 {
		n += 1 + l + sovAsk(uint64(l))
	}
	return n
}

func sovAsk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsk(x uint64) (n int) {
	return sovAsk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsk
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
			return fmt.Errorf("proto: Ask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
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
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
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
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
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
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkReplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsk
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
				return ErrInvalidLengthAsk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkReplier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsk
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
func skipAsk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsk
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
					return 0, ErrIntOverflowAsk
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
					return 0, ErrIntOverflowAsk
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
				return 0, ErrInvalidLengthAsk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsk = fmt.Errorf("proto: unexpected end of group")
)
