// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	BlockReward           types.Coin `protobuf:"bytes,1,opt,name=block_reward,json=blockReward,proto3" json:"block_reward"`
	Baseline              types.Coin `protobuf:"bytes,2,opt,name=baseline,proto3" json:"baseline"`
	AnnualPercentageYield string     `protobuf:"bytes,3,opt,name=annual_percentage_yield,json=annualPercentageYield,proto3" json:"annual_percentage_yield,omitempty"`
	HalvingPeriod         int64      `protobuf:"varint,4,opt,name=halving_period,json=halvingPeriod,proto3" json:"halving_period,omitempty"`
	AdjustmentPeriod      int64      `protobuf:"varint,5,opt,name=adjustment_period,json=adjustmentPeriod,proto3" json:"adjustment_period,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a39a6c05c26f45, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBlockReward() types.Coin {
	if m != nil {
		return m.BlockReward
	}
	return types.Coin{}
}

func (m *Params) GetBaseline() types.Coin {
	if m != nil {
		return m.Baseline
	}
	return types.Coin{}
}

func (m *Params) GetAnnualPercentageYield() string {
	if m != nil {
		return m.AnnualPercentageYield
	}
	return ""
}

func (m *Params) GetHalvingPeriod() int64 {
	if m != nil {
		return m.HalvingPeriod
	}
	return 0
}

func (m *Params) GetAdjustmentPeriod() int64 {
	if m != nil {
		return m.AdjustmentPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "saonetwork.sao.node.Params")
}

func init() { proto.RegisterFile("sao/node/params.proto", fileDescriptor_c6a39a6c05c26f45) }

var fileDescriptor_c6a39a6c05c26f45 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x18, 0x86, 0x5b, 0x40, 0xa2, 0x45, 0x8d, 0x56, 0x89, 0x95, 0xa1, 0x10, 0x13, 0x0d, 0x89, 0xc9,
	0x5d, 0xd0, 0xc4, 0x41, 0x17, 0x83, 0xbb, 0x21, 0x75, 0xd2, 0x85, 0x7c, 0x6d, 0x2f, 0xa5, 0xd2,
	0xde, 0xd7, 0xf4, 0x0e, 0x90, 0x7f, 0xe1, 0x68, 0xe2, 0xe2, 0xcf, 0x61, 0x64, 0x74, 0x32, 0x06,
	0xfe, 0x88, 0xe9, 0x1d, 0xe2, 0xea, 0x76, 0xf9, 0x9e, 0xe7, 0x7d, 0x87, 0x7b, 0xad, 0xba, 0x00,
	0xa4, 0x1c, 0x43, 0x46, 0x33, 0xc8, 0x21, 0x15, 0x24, 0xcb, 0x51, 0xa2, 0x7d, 0x20, 0x00, 0x39,
	0x93, 0x13, 0xcc, 0x87, 0x44, 0x00, 0x92, 0xc2, 0x68, 0x1c, 0x46, 0x18, 0xa1, 0xe2, 0xb4, 0x78,
	0x69, 0xb5, 0xe1, 0x06, 0x28, 0x52, 0x14, 0xd4, 0x07, 0xc1, 0xe8, 0xb8, 0xe3, 0x33, 0x09, 0x1d,
	0x1a, 0x60, 0xcc, 0x35, 0x3f, 0x79, 0x2f, 0x59, 0xd5, 0x9e, 0xea, 0xb6, 0xbb, 0xd6, 0xb6, 0x9f,
	0x60, 0x30, 0xec, 0xe7, 0x6c, 0x02, 0x79, 0xe8, 0x98, 0x2d, 0xb3, 0x5d, 0xbb, 0x38, 0x26, 0xba,
	0x81, 0x14, 0x0d, 0x64, 0xd5, 0x40, 0xee, 0x30, 0xe6, 0xdd, 0xca, 0xec, 0xab, 0x69, 0x78, 0x35,
	0x15, 0xf2, 0x54, 0xc6, 0xbe, 0xb1, 0x36, 0x0b, 0x2f, 0x89, 0x39, 0x73, 0x4a, 0xff, 0xcb, 0xaf,
	0x03, 0xf6, 0x95, 0x75, 0x04, 0x9c, 0x8f, 0x20, 0xe9, 0x67, 0x2c, 0x0f, 0x18, 0x97, 0x10, 0xb1,
	0xfe, 0x34, 0x66, 0x49, 0xe8, 0x94, 0x5b, 0x66, 0x7b, 0xcb, 0xab, 0x6b, 0xdc, 0x5b, 0xd3, 0xc7,
	0x02, 0xda, 0xa7, 0xd6, 0xee, 0x00, 0x92, 0x71, 0xcc, 0xa3, 0x22, 0x18, 0x63, 0xe8, 0x54, 0x5a,
	0x66, 0xbb, 0xec, 0xed, 0xac, 0xae, 0x3d, 0x75, 0xb4, 0xcf, 0xad, 0x7d, 0x08, 0x9f, 0x47, 0x42,
	0xa6, 0x8c, 0xcb, 0x5f, 0x73, 0x43, 0x99, 0x7b, 0x7f, 0x40, 0xcb, 0xd7, 0x95, 0xb7, 0x8f, 0xa6,
	0xd1, 0xbd, 0x9d, 0x2d, 0x5c, 0x73, 0xbe, 0x70, 0xcd, 0xef, 0x85, 0x6b, 0xbe, 0x2e, 0x5d, 0x63,
	0xbe, 0x74, 0x8d, 0xcf, 0xa5, 0x6b, 0x3c, 0x9d, 0x45, 0xb1, 0x1c, 0x8c, 0x7c, 0x12, 0x60, 0x4a,
	0x1f, 0x00, 0xef, 0xf5, 0x1a, 0xb4, 0xd8, 0xeb, 0x45, 0x2f, 0x26, 0xa7, 0x19, 0x13, 0x7e, 0x55,
	0x7d, 0xf3, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x7d, 0xdb, 0xf8, 0xca, 0x01, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AdjustmentPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AdjustmentPeriod))
		i--
		dAtA[i] = 0x28
	}
	if m.HalvingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HalvingPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AnnualPercentageYield) > 0 {
		i -= len(m.AnnualPercentageYield)
		copy(dAtA[i:], m.AnnualPercentageYield)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AnnualPercentageYield)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Baseline.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.BlockReward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BlockReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Baseline.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.AnnualPercentageYield)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.HalvingPeriod != 0 {
		n += 1 + sovParams(uint64(m.HalvingPeriod))
	}
	if m.AdjustmentPeriod != 0 {
		n += 1 + sovParams(uint64(m.AdjustmentPeriod))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Baseline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Baseline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualPercentageYield", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AnnualPercentageYield = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HalvingPeriod", wireType)
			}
			m.HalvingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HalvingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustmentPeriod", wireType)
			}
			m.AdjustmentPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjustmentPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
