// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/node.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
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

type Node struct {
	Creator        string  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Peer           string  `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	Reputation     float32 `protobuf:"fixed32,3,opt,name=reputation,proto3" json:"reputation,omitempty"`
	Status         uint32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	LastAliveHeigh int64   `protobuf:"varint,5,opt,name=lastAliveHeigh,proto3" json:"lastAliveHeigh,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_eddfe2c24a63425f, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Node) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *Node) GetReputation() float32 {
	if m != nil {
		return m.Reputation
	}
	return 0
}

func (m *Node) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Node) GetLastAliveHeigh() int64 {
	if m != nil {
		return m.LastAliveHeigh
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "saonetwork.sao.node.Node")
}

func init() { proto.RegisterFile("sao/node/node.proto", fileDescriptor_eddfe2c24a63425f) }

var fileDescriptor_eddfe2c24a63425f = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x46, 0x3d, 0x89, 0x09, 0x62, 0x24, 0x28, 0x16, 0x09, 0x6d, 0xb5, 0xb2, 0x28, 0x22, 0x57,
	0x76, 0xc1, 0x05, 0x80, 0x8a, 0x2a, 0x85, 0xe9, 0xe8, 0x36, 0xc9, 0x28, 0xb1, 0x08, 0x9e, 0xd5,
	0xee, 0x98, 0x9f, 0x5b, 0x50, 0x70, 0x28, 0xca, 0x94, 0x94, 0xc8, 0xbe, 0x08, 0xf2, 0x1a, 0x24,
	0x94, 0x66, 0x34, 0xef, 0x7d, 0xd5, 0xc3, 0xf3, 0x60, 0xb9, 0x6c, 0x78, 0x4d, 0xf1, 0x14, 0xce,
	0xb3, 0xb0, 0x1a, 0x64, 0x43, 0xf2, 0xc2, 0xfe, 0xb1, 0x08, 0x96, 0x8b, 0x61, 0xba, 0xfc, 0x00,
	0x4c, 0x17, 0xbc, 0x26, 0xa5, 0xf1, 0x78, 0xe5, 0xc9, 0x0a, 0x7b, 0x0d, 0x19, 0xe4, 0x27, 0xd5,
	0x1f, 0x2a, 0x85, 0xa9, 0x23, 0xf2, 0x7a, 0x12, 0x75, 0xfc, 0x95, 0x41, 0xf4, 0xe4, 0x5a, 0xb1,
	0x52, 0x73, 0xa3, 0xa7, 0x19, 0xe4, 0x93, 0xea, 0x9f, 0x51, 0x17, 0x38, 0x0b, 0x62, 0xa5, 0x0d,
	0x3a, 0xcd, 0x20, 0x3f, 0xad, 0x7e, 0x49, 0xcd, 0xf1, 0x6c, 0x67, 0x83, 0xdc, 0xec, 0xea, 0x67,
	0xba, 0xa3, 0x7a, 0xb3, 0xd5, 0x47, 0x19, 0xe4, 0xd3, 0xea, 0xc0, 0xde, 0x5e, 0x7f, 0x76, 0x06,
	0xf6, 0x9d, 0x81, 0xef, 0xce, 0xc0, 0x7b, 0x6f, 0x92, 0x7d, 0x6f, 0x92, 0xaf, 0xde, 0x24, 0x0f,
	0xf3, 0x4d, 0x2d, 0xdb, 0x76, 0x59, 0xac, 0xf8, 0xa9, 0xbc, 0xb7, 0xbc, 0x18, 0x83, 0xca, 0x21,
	0xf8, 0x75, 0x4c, 0x96, 0x37, 0x47, 0x61, 0x39, 0x8b, 0xd1, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7e, 0x6a, 0xf2, 0xe3, 0x0b, 0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastAliveHeigh != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.LastAliveHeigh))
		i--
		dAtA[i] = 0x28
	}
	if m.Status != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Reputation != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Reputation))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.Peer) > 0 {
		i -= len(m.Peer)
		copy(dAtA[i:], m.Peer)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Peer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Peer)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Reputation != 0 {
		n += 5
	}
	if m.Status != 0 {
		n += 1 + sovNode(uint64(m.Status))
	}
	if m.LastAliveHeigh != 0 {
		n += 1 + sovNode(uint64(m.LastAliveHeigh))
	}
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
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
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reputation", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Reputation = float32(math.Float32frombits(v))
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAliveHeigh", wireType)
			}
			m.LastAliveHeigh = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastAliveHeigh |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
