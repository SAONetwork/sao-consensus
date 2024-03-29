// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/sao/timeout_order.proto

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

type TimeoutOrder struct {
	Height    uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	OrderList []uint64 `protobuf:"varint,2,rep,packed,name=orderList,proto3" json:"orderList,omitempty"`
}

func (m *TimeoutOrder) Reset()         { *m = TimeoutOrder{} }
func (m *TimeoutOrder) String() string { return proto.CompactTextString(m) }
func (*TimeoutOrder) ProtoMessage()    {}
func (*TimeoutOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_42cefa198d5ff2bb, []int{0}
}
func (m *TimeoutOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeoutOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeoutOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeoutOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeoutOrder.Merge(m, src)
}
func (m *TimeoutOrder) XXX_Size() int {
	return m.Size()
}
func (m *TimeoutOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeoutOrder.DiscardUnknown(m)
}

var xxx_messageInfo_TimeoutOrder proto.InternalMessageInfo

func (m *TimeoutOrder) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TimeoutOrder) GetOrderList() []uint64 {
	if m != nil {
		return m.OrderList
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeoutOrder)(nil), "saonetwork.sao.sao.TimeoutOrder")
}

func init() { proto.RegisterFile("sao/sao/timeout_order.proto", fileDescriptor_42cefa198d5ff2bb) }

var fileDescriptor_42cefa198d5ff2bb = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0xcc, 0xd7,
	0x07, 0xe1, 0x92, 0xcc, 0xdc, 0xd4, 0xfc, 0xd2, 0x92, 0xf8, 0xfc, 0xa2, 0x94, 0xd4, 0x22, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xa1, 0xe2, 0xc4, 0xfc, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2,
	0x6c, 0xbd, 0xe2, 0xc4, 0x7c, 0x10, 0x56, 0x72, 0xe1, 0xe2, 0x09, 0x81, 0x28, 0xf5, 0x07, 0xa9,
	0x14, 0x12, 0xe3, 0x62, 0xcb, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x09, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0xc1, 0x46, 0xf9, 0x64, 0x16, 0x97, 0x48, 0x30, 0x29,
	0x30, 0x6b, 0xb0, 0x04, 0x21, 0x04, 0x9c, 0xec, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x35, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x38,
	0x31, 0xdf, 0x0f, 0x62, 0x3d, 0xd8, 0x89, 0x15, 0x10, 0x87, 0x56, 0x16, 0xa4, 0x16, 0x27, 0xb1,
	0x81, 0x5d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x53, 0x87, 0x69, 0xc0, 0x00, 0x00,
	0x00,
}

func (m *TimeoutOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeoutOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeoutOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OrderList) > 0 {
		dAtA2 := make([]byte, len(m.OrderList)*10)
		var j1 int
		for _, num := range m.OrderList {
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
		i = encodeVarintTimeoutOrder(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintTimeoutOrder(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimeoutOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimeoutOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimeoutOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTimeoutOrder(uint64(m.Height))
	}
	if len(m.OrderList) > 0 {
		l = 0
		for _, e := range m.OrderList {
			l += sovTimeoutOrder(uint64(e))
		}
		n += 1 + sovTimeoutOrder(uint64(l)) + l
	}
	return n
}

func sovTimeoutOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimeoutOrder(x uint64) (n int) {
	return sovTimeoutOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeoutOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeoutOrder
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
			return fmt.Errorf("proto: TimeoutOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeoutOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeoutOrder
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
						return ErrIntOverflowTimeoutOrder
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
				m.OrderList = append(m.OrderList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTimeoutOrder
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
					return ErrInvalidLengthTimeoutOrder
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTimeoutOrder
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
				if elementCount != 0 && len(m.OrderList) == 0 {
					m.OrderList = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTimeoutOrder
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
					m.OrderList = append(m.OrderList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimeoutOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimeoutOrder
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
func skipTimeoutOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimeoutOrder
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
					return 0, ErrIntOverflowTimeoutOrder
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
					return 0, ErrIntOverflowTimeoutOrder
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
				return 0, ErrInvalidLengthTimeoutOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimeoutOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimeoutOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimeoutOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimeoutOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimeoutOrder = fmt.Errorf("proto: unexpected end of group")
)
