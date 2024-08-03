// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ganymede/exchange/reply.proto

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

type Reply struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AskId       string `protobuf:"bytes,2,opt,name=askId,proto3" json:"askId,omitempty"`
	Sender      string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Replier     string `protobuf:"bytes,4,opt,name=replier,proto3" json:"replier,omitempty"`
	Payload     string `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	SentDate    string `protobuf:"bytes,6,opt,name=sentDate,proto3" json:"sentDate,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	LinkSender  string `protobuf:"bytes,8,opt,name=linkSender,proto3" json:"linkSender,omitempty"`
	LinkReplier string `protobuf:"bytes,9,opt,name=linkReplier,proto3" json:"linkReplier,omitempty"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e28ed20e503a30, []int{0}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return m.Size()
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Reply) GetAskId() string {
	if m != nil {
		return m.AskId
	}
	return ""
}

func (m *Reply) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Reply) GetReplier() string {
	if m != nil {
		return m.Replier
	}
	return ""
}

func (m *Reply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Reply) GetSentDate() string {
	if m != nil {
		return m.SentDate
	}
	return ""
}

func (m *Reply) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Reply) GetLinkSender() string {
	if m != nil {
		return m.LinkSender
	}
	return ""
}

func (m *Reply) GetLinkReplier() string {
	if m != nil {
		return m.LinkReplier
	}
	return ""
}

func init() {
	proto.RegisterType((*Reply)(nil), "ganymede.exchange.Reply")
}

func init() { proto.RegisterFile("ganymede/exchange/reply.proto", fileDescriptor_39e28ed20e503a30) }

var fileDescriptor_39e28ed20e503a30 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x27, 0x75, 0xda, 0x99, 0x89, 0x20, 0x18, 0x44, 0x82, 0x60, 0x28, 0xae, 0x66, 0x63,
	0xbb, 0xf0, 0x00, 0x82, 0xb8, 0x71, 0x5b, 0x57, 0xba, 0xcb, 0x34, 0x8f, 0x5a, 0xa6, 0x4d, 0x4a,
	0x12, 0x61, 0xe2, 0x29, 0x3c, 0x96, 0xcb, 0x59, 0xba, 0x94, 0xf6, 0x04, 0xde, 0x40, 0x9a, 0xb6,
	0xe3, 0x2c, 0xbf, 0xff, 0x7b, 0x8f, 0xf7, 0xf8, 0xf1, 0x75, 0xc1, 0xa5, 0xab, 0x41, 0x40, 0x0a,
	0xbb, 0xfc, 0x8d, 0xcb, 0x02, 0x52, 0x0d, 0x4d, 0xe5, 0x92, 0x46, 0x2b, 0xab, 0xc8, 0xf9, 0xa4,
	0x93, 0x49, 0xdf, 0xfc, 0x22, 0x1c, 0x66, 0xfd, 0x08, 0x39, 0xc3, 0x41, 0x29, 0x28, 0x8a, 0xd1,
	0x7a, 0x9e, 0x05, 0xa5, 0x20, 0x17, 0x38, 0xe4, 0x66, 0xfb, 0x24, 0x68, 0x10, 0xa3, 0xf5, 0x2a,
	0x1b, 0x80, 0x5c, 0xe2, 0xc8, 0x80, 0x14, 0xa0, 0xe9, 0x89, 0x8f, 0x47, 0x22, 0x14, 0x2f, 0xfa,
	0x4b, 0x25, 0x68, 0x3a, 0xf7, 0x62, 0xc2, 0xde, 0x34, 0xdc, 0x55, 0x8a, 0x0b, 0x1a, 0x0e, 0x66,
	0x44, 0x72, 0x85, 0x97, 0x06, 0xa4, 0x7d, 0xe4, 0x16, 0x68, 0xe4, 0xd5, 0x81, 0xfb, 0xad, 0x5c,
	0x03, 0xb7, 0x4a, 0xd3, 0xc5, 0xb0, 0x35, 0x22, 0x61, 0x18, 0x57, 0xa5, 0xdc, 0x3e, 0x0f, 0x5f,
	0x2c, 0xbd, 0x3c, 0x4a, 0x48, 0x8c, 0x4f, 0x7b, 0xca, 0xc6, 0x6f, 0x56, 0x7e, 0xe0, 0x38, 0x7a,
	0x78, 0xf9, 0x6a, 0x19, 0xda, 0xb7, 0x0c, 0xfd, 0xb4, 0x0c, 0x7d, 0x76, 0x6c, 0xb6, 0xef, 0xd8,
	0xec, 0xbb, 0x63, 0xb3, 0xd7, 0xfb, 0xa2, 0xb4, 0x15, 0xdf, 0x24, 0xae, 0xfa, 0x68, 0xb8, 0x4b,
	0x72, 0x55, 0xa7, 0xc5, 0x3b, 0x97, 0x96, 0xab, 0x34, 0x57, 0xa6, 0x56, 0xe6, 0xd6, 0x82, 0xb1,
	0xe9, 0xa1, 0xe5, 0xdd, 0x7f, 0xcf, 0xd6, 0x35, 0x60, 0x36, 0x91, 0x2f, 0xfa, 0xee, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x89, 0x6d, 0x67, 0x89, 0x01, 0x00, 0x00,
}

func (m *Reply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LinkReplier) > 0 {
		i -= len(m.LinkReplier)
		copy(dAtA[i:], m.LinkReplier)
		i = encodeVarintReply(dAtA, i, uint64(len(m.LinkReplier)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.LinkSender) > 0 {
		i -= len(m.LinkSender)
		copy(dAtA[i:], m.LinkSender)
		i = encodeVarintReply(dAtA, i, uint64(len(m.LinkSender)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintReply(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SentDate) > 0 {
		i -= len(m.SentDate)
		copy(dAtA[i:], m.SentDate)
		i = encodeVarintReply(dAtA, i, uint64(len(m.SentDate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintReply(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Replier) > 0 {
		i -= len(m.Replier)
		copy(dAtA[i:], m.Replier)
		i = encodeVarintReply(dAtA, i, uint64(len(m.Replier)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintReply(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AskId) > 0 {
		i -= len(m.AskId)
		copy(dAtA[i:], m.AskId)
		i = encodeVarintReply(dAtA, i, uint64(len(m.AskId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintReply(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintReply(dAtA []byte, offset int, v uint64) int {
	offset -= sovReply(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Reply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovReply(uint64(m.Id))
	}
	l = len(m.AskId)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.Replier)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.SentDate)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.LinkSender)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	l = len(m.LinkReplier)
	if l > 0 {
		n += 1 + l + sovReply(uint64(l))
	}
	return n
}

func sovReply(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReply(x uint64) (n int) {
	return sovReply(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Reply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReply
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
			return fmt.Errorf("proto: Reply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReply
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
				return fmt.Errorf("proto: wrong wireType = %d for field AskId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
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
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
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
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
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
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
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
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkReplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReply
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
				return ErrInvalidLengthReply
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkReplier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReply(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReply
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
func skipReply(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReply
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
					return 0, ErrIntOverflowReply
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
					return 0, ErrIntOverflowReply
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
				return 0, ErrInvalidLengthReply
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReply
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReply
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReply        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReply          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReply = fmt.Errorf("proto: unexpected end of group")
)
