// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/did/did_balances.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type DidBalances struct {
	Did     string     `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Balance types.Coin `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance"`
}

func (m *DidBalances) Reset()         { *m = DidBalances{} }
func (m *DidBalances) String() string { return proto.CompactTextString(m) }
func (*DidBalances) ProtoMessage()    {}
func (*DidBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_09525d919cc87d81, []int{0}
}
func (m *DidBalances) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidBalances.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidBalances.Merge(m, src)
}
func (m *DidBalances) XXX_Size() int {
	return m.Size()
}
func (m *DidBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_DidBalances.DiscardUnknown(m)
}

var xxx_messageInfo_DidBalances proto.InternalMessageInfo

func (m *DidBalances) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *DidBalances) GetBalance() types.Coin {
	if m != nil {
		return m.Balance
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*DidBalances)(nil), "saonetwork.sao.did.DidBalances")
}

func init() { proto.RegisterFile("sao/did/did_balances.proto", fileDescriptor_09525d919cc87d81) }

var fileDescriptor_09525d919cc87d81 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0xcc, 0xd7,
	0x4f, 0xc9, 0x4c, 0x01, 0xe1, 0xf8, 0xa4, 0xc4, 0x9c, 0xc4, 0xbc, 0xe4, 0xd4, 0x62, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xa1, 0xe2, 0xc4, 0xfc, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c,
	0xbd, 0xe2, 0xc4, 0x7c, 0xbd, 0x94, 0xcc, 0x14, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb4,
	0x3e, 0x88, 0x05, 0x51, 0x29, 0x25, 0x97, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x94, 0x58,
	0x9c, 0xaa, 0x5f, 0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0xa8, 0x9f, 0x9c, 0x9f, 0x99, 0x07, 0x91,
	0x57, 0x8a, 0xe2, 0xe2, 0x76, 0xc9, 0x4c, 0x71, 0x82, 0x1a, 0x2f, 0x24, 0xc0, 0xc5, 0x9c, 0x92,
	0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x59, 0x72, 0xb1, 0x43, 0x2d,
	0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x18, 0xa9, 0x07, 0x32, 0x52, 0x0f,
	0x6a, 0xa4, 0x9e, 0x73, 0x7e, 0x66, 0x9e, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x30, 0xf5,
	0x4e, 0xf6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9a, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x1f, 0x9c, 0x98, 0xef, 0x07, 0xf1, 0x8a, 0x3e,
	0xc8, 0xc7, 0x15, 0x60, 0x3f, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x68, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x87, 0xa9, 0xf1, 0x5b, 0x0b, 0x01, 0x00, 0x00,
}

func (m *DidBalances) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidBalances) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidBalances) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDidBalances(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDidBalances(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDidBalances(dAtA []byte, offset int, v uint64) int {
	offset -= sovDidBalances(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidBalances) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDidBalances(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovDidBalances(uint64(l))
	return n
}

func sovDidBalances(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDidBalances(x uint64) (n int) {
	return sovDidBalances(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidBalances) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDidBalances
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
			return fmt.Errorf("proto: DidBalances: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidBalances: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidBalances
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
				return ErrInvalidLengthDidBalances
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidBalances
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
				return ErrInvalidLengthDidBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDidBalances(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDidBalances
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
func skipDidBalances(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDidBalances
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
					return 0, ErrIntOverflowDidBalances
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
					return 0, ErrIntOverflowDidBalances
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
				return 0, ErrInvalidLengthDidBalances
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDidBalances
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDidBalances
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDidBalances        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDidBalances          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDidBalances = fmt.Errorf("proto: unexpected end of group")
)
