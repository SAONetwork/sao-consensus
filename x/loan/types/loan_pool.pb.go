// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/loan/loan_pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type LoanPool struct {
	Total         types.DecCoin                          `protobuf:"bytes,1,opt,name=total,proto3" json:"total"`
	LoanedOut     types.Coin                             `protobuf:"bytes,2,opt,name=loaned_out,json=loanedOut,proto3" json:"loaned_out"`
	TotalBonds    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=totalBonds,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalBonds"`
	LastChargedAt uint64                                 `protobuf:"varint,4,opt,name=last_charged_at,json=lastChargedAt,proto3" json:"last_charged_at,omitempty"`
}

func (m *LoanPool) Reset()         { *m = LoanPool{} }
func (m *LoanPool) String() string { return proto.CompactTextString(m) }
func (*LoanPool) ProtoMessage()    {}
func (*LoanPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e126dc3375a7678, []int{0}
}
func (m *LoanPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoanPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoanPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoanPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoanPool.Merge(m, src)
}
func (m *LoanPool) XXX_Size() int {
	return m.Size()
}
func (m *LoanPool) XXX_DiscardUnknown() {
	xxx_messageInfo_LoanPool.DiscardUnknown(m)
}

var xxx_messageInfo_LoanPool proto.InternalMessageInfo

func (m *LoanPool) GetTotal() types.DecCoin {
	if m != nil {
		return m.Total
	}
	return types.DecCoin{}
}

func (m *LoanPool) GetLoanedOut() types.Coin {
	if m != nil {
		return m.LoanedOut
	}
	return types.Coin{}
}

func (m *LoanPool) GetLastChargedAt() uint64 {
	if m != nil {
		return m.LastChargedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*LoanPool)(nil), "saonetwork.sao.loan.LoanPool")
}

func init() { proto.RegisterFile("sao/loan/loan_pool.proto", fileDescriptor_1e126dc3375a7678) }

var fileDescriptor_1e126dc3375a7678 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0x02, 0x41,
	0x14, 0x85, 0x77, 0x10, 0x8d, 0x8c, 0x31, 0x26, 0xab, 0xc5, 0x42, 0xcc, 0x40, 0x2c, 0x08, 0x0d,
	0x33, 0x41, 0x1b, 0x0b, 0x63, 0x74, 0xa1, 0x34, 0x6a, 0xb0, 0x33, 0x26, 0x9b, 0xd9, 0xdd, 0xc9,
	0x42, 0x58, 0xf6, 0x12, 0xe6, 0xe2, 0xcf, 0x2b, 0x58, 0xf9, 0x30, 0x3e, 0x04, 0x25, 0xb1, 0x32,
	0x16, 0xc4, 0xc0, 0x8b, 0x98, 0xd9, 0xd9, 0x82, 0xc2, 0x66, 0x7e, 0xee, 0xb9, 0xdf, 0x99, 0x93,
	0xb9, 0xd4, 0xd3, 0x12, 0x44, 0x0a, 0x32, 0xcb, 0x97, 0x60, 0x02, 0x90, 0xf2, 0xc9, 0x14, 0x10,
	0xdc, 0x43, 0x2d, 0x21, 0x53, 0xf8, 0x02, 0xd3, 0x11, 0xd7, 0x12, 0xb8, 0xd1, 0x6b, 0x47, 0x09,
	0x24, 0x90, 0xeb, 0xc2, 0x9c, 0x6c, 0x6b, 0x8d, 0x45, 0xa0, 0xc7, 0xa0, 0x45, 0x28, 0xb5, 0x12,
	0xcf, 0x9d, 0x50, 0xa1, 0xec, 0x88, 0x08, 0x86, 0x59, 0xa1, 0x57, 0xad, 0x1e, 0x58, 0xd0, 0x5e,
	0xac, 0x74, 0xf2, 0x5e, 0xa2, 0xbb, 0x37, 0x20, 0xb3, 0x7b, 0x80, 0xd4, 0x3d, 0xa7, 0xdb, 0x08,
	0x28, 0x53, 0x8f, 0x34, 0x48, 0x6b, 0xef, 0xf4, 0x98, 0x17, 0xad, 0xc6, 0x97, 0x17, 0xbe, 0xbc,
	0xa7, 0xa2, 0x2e, 0x0c, 0x33, 0xbf, 0x3c, 0x5f, 0xd6, 0x9d, 0xbe, 0x05, 0xdc, 0x4b, 0x4a, 0x4d,
	0x3e, 0x15, 0x07, 0x30, 0x43, 0xaf, 0x94, 0xe3, 0xd5, 0x7f, 0xf1, 0x0d, 0xb6, 0x62, 0x91, 0xbb,
	0x19, 0xba, 0x4f, 0x94, 0xe6, 0x46, 0x3e, 0x64, 0xb1, 0xf6, 0xb6, 0x1a, 0xa4, 0x55, 0xf1, 0x2f,
	0x4c, 0xd3, 0xcf, 0xb2, 0xde, 0x4c, 0x86, 0x38, 0x98, 0x85, 0x3c, 0x82, 0x71, 0x91, 0xbd, 0xd8,
	0xda, 0x3a, 0x1e, 0x09, 0x7c, 0x9b, 0x28, 0x6d, 0x42, 0x7d, 0x7d, 0xb6, 0x69, 0xf1, 0x60, 0x4f,
	0x45, 0xfd, 0x0d, 0x3f, 0xb7, 0x49, 0x0f, 0x52, 0xa9, 0x31, 0x88, 0x06, 0x72, 0x9a, 0xa8, 0x38,
	0x90, 0xe8, 0x95, 0x1b, 0xa4, 0x55, 0xee, 0xef, 0x9b, 0x72, 0xd7, 0x56, 0xaf, 0xd1, 0xbf, 0x9a,
	0xaf, 0x18, 0x59, 0xac, 0x18, 0xf9, 0x5d, 0x31, 0xf2, 0xb1, 0x66, 0xce, 0x62, 0xcd, 0x9c, 0xef,
	0x35, 0x73, 0x1e, 0x37, 0x33, 0x3c, 0x48, 0xb8, 0xb5, 0x73, 0x11, 0x66, 0x78, 0xaf, 0x76, 0x7c,
	0x79, 0x8e, 0x70, 0x27, 0xff, 0xd5, 0xb3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x6e, 0xa1,
	0xc4, 0xd7, 0x01, 0x00, 0x00,
}

func (m *LoanPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoanPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoanPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastChargedAt != 0 {
		i = encodeVarintLoanPool(dAtA, i, uint64(m.LastChargedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.TotalBonds.Size()
		i -= size
		if _, err := m.TotalBonds.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLoanPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.LoanedOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLoanPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Total.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLoanPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLoanPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoanPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoanPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Total.Size()
	n += 1 + l + sovLoanPool(uint64(l))
	l = m.LoanedOut.Size()
	n += 1 + l + sovLoanPool(uint64(l))
	l = m.TotalBonds.Size()
	n += 1 + l + sovLoanPool(uint64(l))
	if m.LastChargedAt != 0 {
		n += 1 + sovLoanPool(uint64(m.LastChargedAt))
	}
	return n
}

func sovLoanPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoanPool(x uint64) (n int) {
	return sovLoanPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoanPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoanPool
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
			return fmt.Errorf("proto: LoanPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoanPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoanPool
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
				return ErrInvalidLengthLoanPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoanPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoanedOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoanPool
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
				return ErrInvalidLengthLoanPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoanPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LoanedOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBonds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoanPool
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
				return ErrInvalidLengthLoanPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoanPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBonds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastChargedAt", wireType)
			}
			m.LastChargedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoanPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastChargedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoanPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoanPool
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
func skipLoanPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoanPool
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
					return 0, ErrIntOverflowLoanPool
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
					return 0, ErrIntOverflowLoanPool
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
				return 0, ErrInvalidLengthLoanPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoanPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoanPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoanPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoanPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoanPool = fmt.Errorf("proto: unexpected end of group")
)
