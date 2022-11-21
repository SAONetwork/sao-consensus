// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: order/metadata.proto

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

type Metadata struct {
	DataId     string   `protobuf:"bytes,1,opt,name=dataId,proto3" json:"dataId,omitempty"`
	Owner      string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Alias      string   `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	GroupId    string   `protobuf:"bytes,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	OrderId    uint64   `protobuf:"varint,5,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Tags       []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Cid        string   `protobuf:"bytes,7,opt,name=cid,proto3" json:"cid,omitempty"`
	Commits    []string `protobuf:"bytes,8,rep,name=commits,proto3" json:"commits,omitempty"`
	ExtendInfo string   `protobuf:"bytes,9,opt,name=extendInfo,proto3" json:"extendInfo,omitempty"`
	Update     bool     `protobuf:"varint,10,opt,name=update,proto3" json:"update,omitempty"`
	Commit     string   `protobuf:"bytes,11,opt,name=commit,proto3" json:"commit,omitempty"`
	Rule       string   `protobuf:"bytes,12,opt,name=rule,proto3" json:"rule,omitempty"`
	Duration   uint64   `protobuf:"varint,13,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt  uint64   `protobuf:"varint,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a3996f887f605c, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *Metadata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Metadata) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Metadata) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Metadata) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Metadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Metadata) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Metadata) GetCommits() []string {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *Metadata) GetExtendInfo() string {
	if m != nil {
		return m.ExtendInfo
	}
	return ""
}

func (m *Metadata) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

func (m *Metadata) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *Metadata) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *Metadata) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Metadata) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Metadata)(nil), "saonetwork.sao.order.Metadata")
}

func init() { proto.RegisterFile("order/metadata.proto", fileDescriptor_c9a3996f887f605c) }

var fileDescriptor_c9a3996f887f605c = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x10, 0xc0, 0x9b, 0xfe, 0x4d, 0xfc, 0xfd, 0xd1, 0x27, 0xab, 0xfa, 0x74, 0x42, 0x28, 0x8a, 0x58,
	0xc8, 0x94, 0x0c, 0x3c, 0x41, 0xd9, 0x32, 0xc0, 0x10, 0x36, 0x36, 0x37, 0x36, 0x25, 0xa2, 0xc9,
	0x45, 0xb6, 0xa3, 0x96, 0xb7, 0xe0, 0xb1, 0x18, 0x3b, 0x32, 0x56, 0xed, 0x8b, 0x20, 0x9f, 0x5b,
	0x60, 0xca, 0xfd, 0x7e, 0xbe, 0xbb, 0xdc, 0xe9, 0xd8, 0x1c, 0xb5, 0x54, 0x3a, 0x6f, 0x94, 0x15,
	0x52, 0x58, 0x91, 0x75, 0x1a, 0x2d, 0xf2, 0xb9, 0x11, 0xd8, 0x2a, 0xbb, 0x41, 0xfd, 0x92, 0x19,
	0x81, 0x19, 0x25, 0x5d, 0xed, 0x87, 0x2c, 0xbc, 0x3b, 0x25, 0xf2, 0xff, 0x6c, 0xea, 0xbe, 0x85,
	0x84, 0x20, 0x09, 0xd2, 0xa8, 0x3c, 0x11, 0x9f, 0xb3, 0x09, 0x6e, 0x5a, 0xa5, 0x61, 0x48, 0xda,
	0x83, 0xb3, 0x62, 0x5d, 0x0b, 0x03, 0x23, 0x6f, 0x09, 0x38, 0xb0, 0xd9, 0x4a, 0x63, 0xdf, 0x15,
	0x12, 0xc6, 0xe4, 0xcf, 0xe8, 0x5e, 0xe8, 0x9f, 0x85, 0x84, 0x49, 0x12, 0xa4, 0xe3, 0xf2, 0x8c,
	0x9c, 0xb3, 0xb1, 0x15, 0x2b, 0x03, 0xd3, 0x64, 0x94, 0x46, 0x25, 0xc5, 0xfc, 0x1f, 0x1b, 0x55,
	0xb5, 0x84, 0x19, 0xf5, 0x70, 0xa1, 0xab, 0xaf, 0xb0, 0x69, 0x6a, 0x6b, 0x20, 0xa4, 0xc4, 0x33,
	0xf2, 0x98, 0x31, 0xb5, 0xb5, 0xaa, 0x95, 0x45, 0xfb, 0x84, 0x10, 0x51, 0xc9, 0x0f, 0xe3, 0xf6,
	0xea, 0x3b, 0x29, 0xac, 0x02, 0x96, 0x04, 0x69, 0x58, 0x9e, 0xc8, 0x79, 0xdf, 0x02, 0x7e, 0xf9,
	0x7d, 0x3d, 0xb9, 0x79, 0x74, 0xbf, 0x56, 0xf0, 0x9b, 0x2c, 0xc5, 0xfc, 0x82, 0x85, 0xb2, 0xd7,
	0xc2, 0xd6, 0xd8, 0xc2, 0x1f, 0x1a, 0xff, 0x8b, 0xf9, 0x25, 0x8b, 0x2a, 0xad, 0x84, 0x55, 0x72,
	0x61, 0xe1, 0x2f, 0x3d, 0x7e, 0x8b, 0xdb, 0xc5, 0xfb, 0x21, 0x0e, 0x76, 0x87, 0x38, 0xd8, 0x1f,
	0xe2, 0xe0, 0xed, 0x18, 0x0f, 0x76, 0xc7, 0x78, 0xf0, 0x71, 0x8c, 0x07, 0x8f, 0xd7, 0xab, 0xda,
	0x3e, 0xf7, 0xcb, 0xac, 0xc2, 0x26, 0x7f, 0x10, 0x78, 0xef, 0xaf, 0x93, 0x1b, 0x81, 0xf9, 0x36,
	0xf7, 0x47, 0xb4, 0xaf, 0x9d, 0x32, 0xcb, 0x29, 0x9d, 0xf0, 0xe6, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x08, 0x96, 0x28, 0xda, 0x01, 0x00, 0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x70
	}
	if m.Duration != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Rule) > 0 {
		i -= len(m.Rule)
		copy(dAtA[i:], m.Rule)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Rule)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Commit) > 0 {
		i -= len(m.Commit)
		copy(dAtA[i:], m.Commit)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Commit)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Update {
		i--
		if m.Update {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if len(m.ExtendInfo) > 0 {
		i -= len(m.ExtendInfo)
		copy(dAtA[i:], m.ExtendInfo)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.ExtendInfo)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Commits) > 0 {
		for iNdEx := len(m.Commits) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Commits[iNdEx])
			copy(dAtA[i:], m.Commits[iNdEx])
			i = encodeVarintMetadata(dAtA, i, uint64(len(m.Commits[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintMetadata(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.OrderId != 0 {
		i = encodeVarintMetadata(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Alias) > 0 {
		i -= len(m.Alias)
		copy(dAtA[i:], m.Alias)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Alias)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DataId) > 0 {
		i -= len(m.DataId)
		copy(dAtA[i:], m.DataId)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.DataId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DataId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.OrderId != 0 {
		n += 1 + sovMetadata(uint64(m.OrderId))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Commits) > 0 {
		for _, s := range m.Commits {
			l = len(s)
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	l = len(m.ExtendInfo)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.Update {
		n += 2
	}
	l = len(m.Commit)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Rule)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovMetadata(uint64(m.Duration))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovMetadata(uint64(m.CreatedAt))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commits = append(m.Commits, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtendInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Update = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
