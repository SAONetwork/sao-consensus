// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/order/shard.proto

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

type Shard struct {
	Id        uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId   uint64     `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status    int32      `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Size_     uint64     `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Cid       string     `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
	Pledge    types.Coin `protobuf:"bytes,6,opt,name=pledge,proto3" json:"pledge"`
	From      string     `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	Sp        string     `protobuf:"bytes,8,opt,name=sp,proto3" json:"sp,omitempty"`
	Duration  uint64     `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt uint64     `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *Shard) Reset()         { *m = Shard{} }
func (m *Shard) String() string { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()    {}
func (*Shard) Descriptor() ([]byte, []int) {
	return fileDescriptor_b050d2e483ae2036, []int{0}
}
func (m *Shard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Shard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Shard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Shard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shard.Merge(m, src)
}
func (m *Shard) XXX_Size() int {
	return m.Size()
}
func (m *Shard) XXX_DiscardUnknown() {
	xxx_messageInfo_Shard.DiscardUnknown(m)
}

var xxx_messageInfo_Shard proto.InternalMessageInfo

func (m *Shard) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shard) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Shard) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Shard) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Shard) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Shard) GetPledge() types.Coin {
	if m != nil {
		return m.Pledge
	}
	return types.Coin{}
}

func (m *Shard) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Shard) GetSp() string {
	if m != nil {
		return m.Sp
	}
	return ""
}

func (m *Shard) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Shard) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Shard)(nil), "saonetwork.sao.order.Shard")
}

func init() { proto.RegisterFile("sao/order/shard.proto", fileDescriptor_b050d2e483ae2036) }

var fileDescriptor_b050d2e483ae2036 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0x45, 0x59, 0x96, 0x6d, 0x16, 0x28, 0x0a, 0xc2, 0x2d, 0x58, 0xa3, 0x50, 0x85, 0x2e,
	0xd5, 0x44, 0xc2, 0xed, 0x90, 0xd9, 0xce, 0x94, 0x25, 0x83, 0xbc, 0x65, 0xa3, 0x44, 0x46, 0x26,
	0x12, 0xeb, 0x0a, 0x24, 0x9d, 0xbf, 0x07, 0xc8, 0x9c, 0xc7, 0xf2, 0xe8, 0x31, 0x53, 0x10, 0xd8,
	0x2f, 0x12, 0x88, 0x52, 0x92, 0xed, 0xdc, 0xc3, 0xc3, 0x83, 0x0f, 0xf7, 0xe2, 0xef, 0x56, 0x00,
	0x07, 0x23, 0x95, 0xe1, 0x76, 0x2d, 0x8c, 0x64, 0x8d, 0x01, 0x07, 0x64, 0x6a, 0x05, 0xd4, 0xca,
	0xdd, 0x82, 0xb9, 0x62, 0x56, 0x00, 0xf3, 0x89, 0x59, 0x52, 0x82, 0xdd, 0x80, 0xe5, 0x85, 0xb0,
	0x8a, 0xdf, 0xcc, 0x0b, 0xe5, 0xc4, 0x9c, 0x97, 0xa0, 0xeb, 0xee, 0xd7, 0x6c, 0x5a, 0x41, 0x05,
	0x5e, 0xf2, 0x56, 0x75, 0xee, 0x9f, 0xc7, 0x10, 0x0f, 0x57, 0x6d, 0x37, 0xf9, 0x8a, 0x43, 0x2d,
	0x29, 0x4a, 0x51, 0x16, 0xe5, 0xa1, 0x96, 0x84, 0xe2, 0x91, 0x2f, 0x3e, 0x93, 0x34, 0xf4, 0xe6,
	0xfb, 0x48, 0x7e, 0xe0, 0xd8, 0x3a, 0xe1, 0xb6, 0x96, 0x0e, 0x52, 0x94, 0x0d, 0xf3, 0x7e, 0x22,
	0x04, 0x47, 0x56, 0x3f, 0x28, 0x1a, 0xf9, 0xb8, 0xd7, 0xe4, 0x1b, 0x1e, 0x94, 0x5a, 0xd2, 0x61,
	0x8a, 0xb2, 0x49, 0xde, 0x4a, 0x72, 0x82, 0xe3, 0xe6, 0x5a, 0xc9, 0x4a, 0xd1, 0x38, 0x45, 0xd9,
	0x97, 0x7f, 0x3f, 0x59, 0x07, 0xce, 0x5a, 0x70, 0xd6, 0x83, 0xb3, 0x53, 0xd0, 0xf5, 0x32, 0xda,
	0xbd, 0xfc, 0x0e, 0xf2, 0x3e, 0xde, 0xd6, 0x5f, 0x1a, 0xd8, 0xd0, 0x91, 0xef, 0xf2, 0xba, 0x85,
	0xb6, 0x0d, 0x1d, 0x7b, 0x27, 0xb4, 0x0d, 0x99, 0xe1, 0xb1, 0xdc, 0x1a, 0xe1, 0x34, 0xd4, 0x74,
	0xe2, 0x31, 0x3e, 0x66, 0xf2, 0x0b, 0x4f, 0x4a, 0xa3, 0x84, 0x53, 0x72, 0xe1, 0x28, 0xf6, 0x8f,
	0x9f, 0xc6, 0x72, 0xb1, 0x3b, 0x24, 0x68, 0x7f, 0x48, 0xd0, 0xeb, 0x21, 0x41, 0x4f, 0xc7, 0x24,
	0xd8, 0x1f, 0x93, 0xe0, 0xf9, 0x98, 0x04, 0x17, 0x7f, 0x2b, 0xed, 0xd6, 0xdb, 0x82, 0x95, 0xb0,
	0xe1, 0x2b, 0x01, 0xe7, 0xdd, 0xe6, 0x79, 0x7b, 0x9b, 0xbb, 0xfe, 0x3a, 0xee, 0xbe, 0x51, 0xb6,
	0x88, 0xfd, 0x4a, 0xff, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x85, 0x30, 0xb7, 0xef, 0xb7, 0x01,
	0x00, 0x00,
}

func (m *Shard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Shard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Shard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x50
	}
	if m.Duration != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Sp) > 0 {
		i -= len(m.Sp)
		copy(dAtA[i:], m.Sp)
		i = encodeVarintShard(dAtA, i, uint64(len(m.Sp)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintShard(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size, err := m.Pledge.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintShard(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintShard(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Size_ != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.OrderId != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShard(dAtA []byte, offset int, v uint64) int {
	offset -= sovShard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Shard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShard(uint64(m.Id))
	}
	if m.OrderId != 0 {
		n += 1 + sovShard(uint64(m.OrderId))
	}
	if m.Status != 0 {
		n += 1 + sovShard(uint64(m.Status))
	}
	if m.Size_ != 0 {
		n += 1 + sovShard(uint64(m.Size_))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovShard(uint64(l))
	}
	l = m.Pledge.Size()
	n += 1 + l + sovShard(uint64(l))
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovShard(uint64(l))
	}
	l = len(m.Sp)
	if l > 0 {
		n += 1 + l + sovShard(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovShard(uint64(m.Duration))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovShard(uint64(m.CreatedAt))
	}
	return n
}

func sovShard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShard(x uint64) (n int) {
	return sovShard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Shard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShard
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
			return fmt.Errorf("proto: Shard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Shard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
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
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pledge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
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
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pledge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
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
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
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
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShard
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
func skipShard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShard
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
					return 0, ErrIntOverflowShard
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
					return 0, ErrIntOverflowShard
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
				return 0, ErrInvalidLengthShard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShard = fmt.Errorf("proto: unexpected end of group")
)
