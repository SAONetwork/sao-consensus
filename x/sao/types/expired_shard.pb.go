// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/sao/expired_shard.proto

package types

import (
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

type ExpiredShard struct {
	Height    uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	ShardList []uint64 `protobuf:"varint,2,rep,packed,name=shardList,proto3" json:"shardList,omitempty"`
}

func (m *ExpiredShard) Reset()         { *m = ExpiredShard{} }
func (m *ExpiredShard) String() string { return proto.CompactTextString(m) }
func (*ExpiredShard) ProtoMessage()    {}
func (*ExpiredShard) Descriptor() ([]byte, []int) {
	return fileDescriptor_15db09f895192250, []int{0}
}
func (m *ExpiredShard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExpiredShard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExpiredShard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpiredShard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpiredShard.Merge(m, src)
}
func (m *ExpiredShard) XXX_Size() int {
	return m.Size()
}
func (m *ExpiredShard) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpiredShard.DiscardUnknown(m)
}

var xxx_messageInfo_ExpiredShard proto.InternalMessageInfo

func (m *ExpiredShard) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ExpiredShard) GetShardList() []uint64 {
	if m != nil {
		return m.ShardList
	}
	return nil
}

func init() {
	proto.RegisterType((*ExpiredShard)(nil), "saonetwork.sao.sao.ExpiredShard")
}

func init() { proto.RegisterFile("sao/sao/expired_shard.proto", fileDescriptor_15db09f895192250) }

var fileDescriptor_15db09f895192250 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0xcc, 0xd7,
	0x07, 0xe1, 0xd4, 0x8a, 0x82, 0xcc, 0xa2, 0xd4, 0x94, 0xf8, 0xe2, 0x8c, 0xc4, 0xa2, 0x14, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xa1, 0xe2, 0xc4, 0xfc, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2,
	0x6c, 0xbd, 0xe2, 0xc4, 0x7c, 0x10, 0x56, 0x72, 0xe1, 0xe2, 0x71, 0x85, 0x28, 0x0d, 0x06, 0xa9,
	0x14, 0x12, 0xe3, 0x62, 0xcb, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x09, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0xc1, 0x46, 0xf9, 0x64, 0x16, 0x97, 0x48, 0x30, 0x29,
	0x30, 0x6b, 0xb0, 0x04, 0x21, 0x04, 0x9c, 0xec, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x35, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x38,
	0x31, 0xdf, 0x0f, 0x62, 0x3d, 0xd8, 0x89, 0x15, 0x60, 0xb2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0xec, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x41, 0x17, 0xf3, 0xc0, 0x00,
	0x00, 0x00,
}

func (m *ExpiredShard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpiredShard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExpiredShard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShardList) > 0 {
		dAtA2 := make([]byte, len(m.ShardList)*10)
		var j1 int
		for _, num := range m.ShardList {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintExpiredShard(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintExpiredShard(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExpiredShard(dAtA []byte, offset int, v uint64) int {
	offset -= sovExpiredShard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExpiredShard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovExpiredShard(uint64(m.Height))
	}
	if len(m.ShardList) > 0 {
		l = 0
		for _, e := range m.ShardList {
			l += sovExpiredShard(uint64(e))
		}
		n += 1 + sovExpiredShard(uint64(l)) + l
	}
	return n
}

func sovExpiredShard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExpiredShard(x uint64) (n int) {
	return sovExpiredShard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExpiredShard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExpiredShard
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
			return fmt.Errorf("proto: ExpiredShard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpiredShard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExpiredShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExpiredShard
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ShardList = append(m.ShardList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExpiredShard
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthExpiredShard
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExpiredShard
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ShardList) == 0 {
					m.ShardList = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExpiredShard
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ShardList = append(m.ShardList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExpiredShard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExpiredShard
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
func skipExpiredShard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExpiredShard
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
					return 0, ErrIntOverflowExpiredShard
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
					return 0, ErrIntOverflowExpiredShard
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
				return 0, ErrInvalidLengthExpiredShard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExpiredShard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExpiredShard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExpiredShard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExpiredShard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExpiredShard = fmt.Errorf("proto: unexpected end of group")
)
