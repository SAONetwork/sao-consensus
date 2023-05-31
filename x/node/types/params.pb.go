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
	BlockReward           types.Coin    `protobuf:"bytes,1,opt,name=block_reward,json=blockReward,proto3" json:"block_reward"`
	Baseline              types.Coin    `protobuf:"bytes,2,opt,name=baseline,proto3" json:"baseline"`
	AnnualPercentageYield string        `protobuf:"bytes,3,opt,name=annual_percentage_yield,json=annualPercentageYield,proto3" json:"annual_percentage_yield,omitempty"`
	FishmenParam          *FishmenParam `protobuf:"bytes,4,opt,name=FishmenParam,proto3" json:"FishmenParam,omitempty"`
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

func (m *Params) GetFishmenParam() *FishmenParam {
	if m != nil {
		return m.FishmenParam
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "saonetwork.sao.node.Params")
}

func init() { proto.RegisterFile("sao/node/params.proto", fileDescriptor_c6a39a6c05c26f45) }

var fileDescriptor_c6a39a6c05c26f45 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xfa, 0x40,
	0x18, 0xc6, 0x7b, 0xfc, 0x09, 0xf9, 0x5b, 0x98, 0xaa, 0x68, 0x65, 0x38, 0xd0, 0xc1, 0x30, 0xdd,
	0x05, 0x4d, 0x1c, 0x74, 0x31, 0x98, 0xb8, 0x98, 0x18, 0x52, 0x27, 0x5d, 0xc8, 0xb5, 0xbc, 0x96,
	0x86, 0xf6, 0xde, 0xa6, 0x57, 0x44, 0xbe, 0x85, 0x83, 0x83, 0xa3, 0x1f, 0x87, 0x91, 0xd1, 0xc9,
	0x18, 0xf8, 0x22, 0xa6, 0x77, 0xa4, 0xd1, 0xc4, 0xc1, 0xad, 0xe9, 0xf3, 0xfc, 0x9e, 0xe7, 0xde,
	0xc7, 0x6e, 0x2a, 0x81, 0x5c, 0xe2, 0x08, 0x78, 0x2a, 0x32, 0x91, 0x28, 0x96, 0x66, 0x98, 0xa3,
	0xb3, 0xad, 0x04, 0x4a, 0xc8, 0x67, 0x98, 0x4d, 0x98, 0x12, 0xc8, 0x0a, 0x47, 0x6b, 0x27, 0xc4,
	0x10, 0xb5, 0xce, 0x8b, 0x2f, 0x63, 0x6d, 0xd1, 0x00, 0x55, 0x82, 0x8a, 0xfb, 0x42, 0x01, 0x7f,
	0xec, 0xf9, 0x90, 0x8b, 0x1e, 0x0f, 0x30, 0x92, 0x1b, 0x7d, 0xb7, 0x6c, 0x78, 0x88, 0xd4, 0x38,
	0x81, 0xcd, 0xff, 0xc3, 0x97, 0x8a, 0x5d, 0x1b, 0xe8, 0x4e, 0xa7, 0x6f, 0x37, 0xfc, 0x18, 0x83,
	0xc9, 0x30, 0x83, 0x99, 0xc8, 0x46, 0x2e, 0xe9, 0x90, 0x6e, 0xfd, 0x78, 0x9f, 0x99, 0x64, 0x56,
	0x24, 0xb3, 0x4d, 0x32, 0xbb, 0xc4, 0x48, 0xf6, 0xab, 0x8b, 0x8f, 0xb6, 0xe5, 0xd5, 0x35, 0xe4,
	0x69, 0xc6, 0x39, 0xb7, 0xff, 0x17, 0xbe, 0x38, 0x92, 0xe0, 0x56, 0xfe, 0xc6, 0x97, 0x80, 0x73,
	0x6a, 0xef, 0x09, 0x29, 0xa7, 0x22, 0x1e, 0xa6, 0x90, 0x05, 0x20, 0x73, 0x11, 0xc2, 0x70, 0x1e,
	0x41, 0x3c, 0x72, 0xff, 0x75, 0x48, 0x77, 0xcb, 0x6b, 0x1a, 0x79, 0x50, 0xaa, 0x77, 0x85, 0xe8,
	0x5c, 0xdb, 0x8d, 0x2b, 0x73, 0x94, 0xbe, 0xc4, 0xad, 0xea, 0xe2, 0x03, 0xf6, 0xcb, 0x7a, 0xec,
	0xbb, 0x51, 0x3f, 0x80, 0x78, 0x3f, 0xe0, 0xb3, 0xea, 0xeb, 0x5b, 0xdb, 0xea, 0x5f, 0x2c, 0x56,
	0x94, 0x2c, 0x57, 0x94, 0x7c, 0xae, 0x28, 0x79, 0x5e, 0x53, 0x6b, 0xb9, 0xa6, 0xd6, 0xfb, 0x9a,
	0x5a, 0xf7, 0x47, 0x61, 0x94, 0x8f, 0xa7, 0x3e, 0x0b, 0x30, 0xe1, 0xb7, 0x02, 0x6f, 0x4c, 0x01,
	0x2f, 0xe6, 0x7d, 0x32, 0x03, 0xe7, 0xf3, 0x14, 0x94, 0x5f, 0xd3, 0xfb, 0x9e, 0x7c, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x93, 0xc6, 0x51, 0x65, 0xdb, 0x01, 0x00, 0x00,
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
	if m.FishmenParam != nil {
		{
			size, err := m.FishmenParam.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
	if m.FishmenParam != nil {
		l = m.FishmenParam.Size()
		n += 1 + l + sovParams(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FishmenParam", wireType)
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
			if m.FishmenParam == nil {
				m.FishmenParam = &FishmenParam{}
			}
			if err := m.FishmenParam.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
