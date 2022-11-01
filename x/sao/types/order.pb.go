// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/order.proto

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

type Order struct {
	Creator  string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Owner    string            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Id       uint64            `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Provider string            `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Cid      string            `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
	Duration int32             `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Expire   int32             `protobuf:"varint,7,opt,name=expire,proto3" json:"expire,omitempty"`
	Status   int32             `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Replica  int32             `protobuf:"varint,9,opt,name=replica,proto3" json:"replica,omitempty"`
	Metadata string            `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Shards   map[string]*Shard `protobuf:"bytes,11,rep,name=shards,proto3" json:"shards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a8964a407d075e9, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Order) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Order) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Order) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Order) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetExpire() int32 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Order) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Order) GetShards() map[string]*Shard {
	if m != nil {
		return m.Shards
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "saonetwork.sao.sao.Order")
	proto.RegisterMapType((map[string]*Shard)(nil), "saonetwork.sao.sao.Order.ShardsEntry")
}

func init() { proto.RegisterFile("sao/order.proto", fileDescriptor_0a8964a407d075e9) }

var fileDescriptor_0a8964a407d075e9 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x76, 0xed, 0xb6, 0x14, 0x54, 0x82, 0x48, 0xdc, 0xa1, 0x14, 0x61, 0xd0, 0x53,
	0x0b, 0xf3, 0x22, 0x82, 0x08, 0x82, 0x57, 0x85, 0xce, 0x93, 0xb7, 0xac, 0x09, 0x2e, 0x6c, 0x6b,
	0x4a, 0x92, 0xee, 0xcf, 0x77, 0xf0, 0xe0, 0xc7, 0xf2, 0xb8, 0xa3, 0x47, 0xd9, 0xbe, 0x88, 0x24,
	0xd9, 0x86, 0xa0, 0x87, 0x96, 0xfc, 0x9e, 0x3c, 0xef, 0xdb, 0xb7, 0xcf, 0x0b, 0x4f, 0x15, 0x11,
	0xb9, 0x90, 0x94, 0xc9, 0xac, 0x96, 0x42, 0x0b, 0x84, 0x14, 0x11, 0x15, 0xd3, 0x4b, 0x21, 0xa7,
	0x99, 0x22, 0xc2, 0x3c, 0x7d, 0x6b, 0x52, 0x13, 0x22, 0xa9, 0x33, 0x5d, 0xbd, 0xfb, 0x30, 0x78,
	0x36, 0x45, 0x08, 0xc3, 0x4e, 0x29, 0x19, 0xd1, 0x42, 0x62, 0x90, 0x80, 0xb4, 0x57, 0x1c, 0x10,
	0x9d, 0xc3, 0x40, 0x2c, 0x2b, 0x26, 0xb1, 0x67, 0x75, 0x07, 0xe8, 0x04, 0x7a, 0x9c, 0x62, 0x3f,
	0x01, 0x69, 0xbb, 0xf0, 0x38, 0x45, 0x7d, 0xd8, 0xad, 0xa5, 0x58, 0x70, 0xca, 0x24, 0x6e, 0x5b,
	0xe3, 0x91, 0xd1, 0x19, 0xf4, 0x4b, 0x4e, 0x71, 0x60, 0x65, 0x73, 0x34, 0x6e, 0xda, 0x48, 0xa2,
	0xb9, 0xa8, 0x70, 0x98, 0x80, 0x34, 0x28, 0x8e, 0x8c, 0x2e, 0x60, 0xc8, 0x56, 0x35, 0x97, 0x0c,
	0x77, 0xec, 0xcd, 0x9e, 0x8c, 0xae, 0x34, 0xd1, 0x8d, 0xc2, 0x5d, 0xa7, 0x3b, 0x32, 0x93, 0x4b,
	0x56, 0xcf, 0x78, 0x49, 0x70, 0xcf, 0x5e, 0x1c, 0xd0, 0x7c, 0x65, 0xce, 0x34, 0xa1, 0x44, 0x13,
	0x0c, 0xdd, 0x4c, 0x07, 0x46, 0x77, 0x30, 0xb4, 0x41, 0x28, 0x1c, 0x25, 0x7e, 0x1a, 0x0d, 0x07,
	0xd9, 0xdf, 0xbc, 0x32, 0x1b, 0x4d, 0x36, 0xb2, 0xbe, 0xc7, 0x4a, 0xcb, 0x75, 0xb1, 0x2f, 0xea,
	0xbf, 0xc0, 0xe8, 0x97, 0x6c, 0xfe, 0x70, 0xca, 0xd6, 0xfb, 0xe4, 0xcc, 0x11, 0xe5, 0x30, 0x58,
	0x90, 0x59, 0xc3, 0x6c, 0x6a, 0xd1, 0xf0, 0xf2, 0xbf, 0xf6, 0xb6, 0x43, 0xe1, 0x7c, 0xb7, 0xde,
	0x0d, 0x78, 0xb8, 0xff, 0xdc, 0xc6, 0x60, 0xb3, 0x8d, 0xc1, 0xf7, 0x36, 0x06, 0x1f, 0xbb, 0xb8,
	0xb5, 0xd9, 0xc5, 0xad, 0xaf, 0x5d, 0xdc, 0x7a, 0x1d, 0xbc, 0x71, 0x3d, 0x69, 0xc6, 0x59, 0x29,
	0xe6, 0xf9, 0x88, 0x88, 0x27, 0xd7, 0x29, 0x37, 0xfb, 0x5c, 0xd9, 0xb7, 0x5e, 0xd7, 0x4c, 0x8d,
	0x43, 0xbb, 0xd6, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xd2, 0x4f, 0xc4, 0x0e, 0x02,
	0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Shards) > 0 {
		for k := range m.Shards {
			v := m.Shards[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintOrder(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintOrder(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintOrder(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x52
	}
	if m.Replica != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Replica))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.Expire != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Expire))
		i--
		dAtA[i] = 0x38
	}
	if m.Duration != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x22
	}
	if m.Id != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovOrder(uint64(m.Id))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovOrder(uint64(m.Duration))
	}
	if m.Expire != 0 {
		n += 1 + sovOrder(uint64(m.Expire))
	}
	if m.Status != 0 {
		n += 1 + sovOrder(uint64(m.Status))
	}
	if m.Replica != 0 {
		n += 1 + sovOrder(uint64(m.Replica))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if len(m.Shards) > 0 {
		for k, v := range m.Shards {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovOrder(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovOrder(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovOrder(uint64(mapEntrySize))
		}
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expire", wireType)
			}
			m.Expire = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expire |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			m.Replica = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Replica |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shards == nil {
				m.Shards = make(map[string]*Shard)
			}
			var mapkey string
			var mapvalue *Shard
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrder
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrder
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthOrder
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthOrder
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrder
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthOrder
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthOrder
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Shard{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOrder(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthOrder
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Shards[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
