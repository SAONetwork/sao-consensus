// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/did/did_binding_proof.proto

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

type DidBingingProof struct {
	AccountId string        `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Proof     *BindingProof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *DidBingingProof) Reset()         { *m = DidBingingProof{} }
func (m *DidBingingProof) String() string { return proto.CompactTextString(m) }
func (*DidBingingProof) ProtoMessage()    {}
func (*DidBingingProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_203c3540b8479edc, []int{0}
}
func (m *DidBingingProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidBingingProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidBingingProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidBingingProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidBingingProof.Merge(m, src)
}
func (m *DidBingingProof) XXX_Size() int {
	return m.Size()
}
func (m *DidBingingProof) XXX_DiscardUnknown() {
	xxx_messageInfo_DidBingingProof.DiscardUnknown(m)
}

var xxx_messageInfo_DidBingingProof proto.InternalMessageInfo

func (m *DidBingingProof) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *DidBingingProof) GetProof() *BindingProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*DidBingingProof)(nil), "saonetwork.sao.did.DidBingingProof")
}

func init() { proto.RegisterFile("sao/did/did_binding_proof.proto", fileDescriptor_203c3540b8479edc) }

var fileDescriptor_203c3540b8479edc = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4e, 0xcc, 0xd7,
	0x4f, 0xc9, 0x4c, 0x01, 0xe1, 0xf8, 0xa4, 0xcc, 0xbc, 0x94, 0xcc, 0xbc, 0xf4, 0xf8, 0x82, 0xa2,
	0xfc, 0xfc, 0x34, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xa1, 0xe2, 0xc4, 0xfc, 0xbc, 0xd4,
	0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0xbd, 0xe2, 0xc4, 0x7c, 0xbd, 0x94, 0xcc, 0x14, 0x29, 0x69, 0x98,
	0x26, 0x2c, 0x1a, 0x94, 0xd2, 0xb9, 0xf8, 0x5d, 0x32, 0x53, 0x9c, 0x32, 0xf3, 0xd2, 0x33, 0xf3,
	0xd2, 0x03, 0x40, 0x12, 0x42, 0x32, 0x5c, 0x9c, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x9e,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x33, 0x2e, 0x56, 0xb0, 0x7e,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x05, 0x3d, 0x4c, 0x1b, 0xf5, 0x9c, 0x20, 0x16, 0x81,
	0x8d, 0x0b, 0x82, 0x28, 0x77, 0xb2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xd5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe0, 0xc4, 0x7c,
	0x3f, 0x88, 0x61, 0xfa, 0x20, 0x57, 0x57, 0x80, 0xdd, 0x5d, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x76, 0xb0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x96, 0x2c, 0xda, 0xc3, 0x04, 0x01, 0x00,
	0x00,
}

func (m *DidBingingProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidBingingProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidBingingProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDidBindingProof(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountId) > 0 {
		i -= len(m.AccountId)
		copy(dAtA[i:], m.AccountId)
		i = encodeVarintDidBindingProof(dAtA, i, uint64(len(m.AccountId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDidBindingProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovDidBindingProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidBingingProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountId)
	if l > 0 {
		n += 1 + l + sovDidBindingProof(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovDidBindingProof(uint64(l))
	}
	return n
}

func sovDidBindingProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDidBindingProof(x uint64) (n int) {
	return sovDidBindingProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidBingingProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDidBindingProof
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
			return fmt.Errorf("proto: DidBingingProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidBingingProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidBindingProof
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
				return ErrInvalidLengthDidBindingProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidBindingProof
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
				return ErrInvalidLengthDidBindingProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidBindingProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &BindingProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDidBindingProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDidBindingProof
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
func skipDidBindingProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDidBindingProof
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
					return 0, ErrIntOverflowDidBindingProof
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
					return 0, ErrIntOverflowDidBindingProof
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
				return 0, ErrInvalidLengthDidBindingProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDidBindingProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDidBindingProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDidBindingProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDidBindingProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDidBindingProof = fmt.Errorf("proto: unexpected end of group")
)