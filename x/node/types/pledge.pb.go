// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/pledge.proto

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

type Pledge struct {
	Creator             string        `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TotalStoragePledged types.Coin    `protobuf:"bytes,2,opt,name=total_storage_pledged,json=totalStoragePledged,proto3" json:"total_storage_pledged"`
	TotalShardPledged   types.Coin    `protobuf:"bytes,3,opt,name=total_shard_pledged,json=totalShardPledged,proto3" json:"total_shard_pledged"`
	Reward              types.DecCoin `protobuf:"bytes,4,opt,name=reward,proto3" json:"reward"`
	RewardDebt          types.DecCoin `protobuf:"bytes,5,opt,name=reward_debt,json=rewardDebt,proto3" json:"reward_debt"`
	TotalStorage        int64         `protobuf:"varint,6,opt,name=total_storage,json=totalStorage,proto3" json:"total_storage,omitempty"`
	UsedStorage         int64         `protobuf:"varint,7,opt,name=used_storage,json=usedStorage,proto3" json:"used_storage,omitempty"`
}

func (m *Pledge) Reset()         { *m = Pledge{} }
func (m *Pledge) String() string { return proto.CompactTextString(m) }
func (*Pledge) ProtoMessage()    {}
func (*Pledge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44a3a94106aee56, []int{0}
}
func (m *Pledge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pledge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pledge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pledge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pledge.Merge(m, src)
}
func (m *Pledge) XXX_Size() int {
	return m.Size()
}
func (m *Pledge) XXX_DiscardUnknown() {
	xxx_messageInfo_Pledge.DiscardUnknown(m)
}

var xxx_messageInfo_Pledge proto.InternalMessageInfo

func (m *Pledge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Pledge) GetTotalStoragePledged() types.Coin {
	if m != nil {
		return m.TotalStoragePledged
	}
	return types.Coin{}
}

func (m *Pledge) GetTotalShardPledged() types.Coin {
	if m != nil {
		return m.TotalShardPledged
	}
	return types.Coin{}
}

func (m *Pledge) GetReward() types.DecCoin {
	if m != nil {
		return m.Reward
	}
	return types.DecCoin{}
}

func (m *Pledge) GetRewardDebt() types.DecCoin {
	if m != nil {
		return m.RewardDebt
	}
	return types.DecCoin{}
}

func (m *Pledge) GetTotalStorage() int64 {
	if m != nil {
		return m.TotalStorage
	}
	return 0
}

func (m *Pledge) GetUsedStorage() int64 {
	if m != nil {
		return m.UsedStorage
	}
	return 0
}

func init() {
	proto.RegisterType((*Pledge)(nil), "saonetwork.sao.node.Pledge")
}

func init() { proto.RegisterFile("sao/node/pledge.proto", fileDescriptor_c44a3a94106aee56) }

var fileDescriptor_c44a3a94106aee56 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x67, 0x3e, 0xf8, 0x86, 0x58, 0x70, 0xe1, 0x20, 0xc9, 0x48, 0x4c, 0x45, 0x4d, 0x0c,
	0xab, 0x36, 0xe8, 0xce, 0x95, 0x01, 0xd6, 0x6a, 0x60, 0xe7, 0x86, 0x74, 0xa6, 0x37, 0x03, 0x11,
	0xe6, 0x92, 0xb6, 0x88, 0xbe, 0x85, 0xbe, 0x15, 0x4b, 0x96, 0xae, 0x8c, 0x81, 0x17, 0x31, 0x33,
	0x1d, 0x08, 0x26, 0x2e, 0xd8, 0xdd, 0xf6, 0x9e, 0xf3, 0x3b, 0xfd, 0x73, 0x49, 0x4d, 0x0b, 0xe4,
	0x09, 0x4a, 0xe0, 0xd3, 0x31, 0xc8, 0x18, 0xd8, 0x54, 0xa1, 0x41, 0xbf, 0xaa, 0x05, 0x26, 0x60,
	0xe6, 0xa8, 0x9e, 0x99, 0x16, 0xc8, 0x52, 0x45, 0xfd, 0x38, 0xc6, 0x18, 0xb3, 0x3e, 0x4f, 0x2b,
	0x2b, 0xad, 0xd3, 0x08, 0xf5, 0x04, 0x35, 0x0f, 0x85, 0x06, 0xfe, 0xd2, 0x0a, 0xc1, 0x88, 0x16,
	0x8f, 0x70, 0x94, 0xd8, 0xfe, 0xc5, 0x47, 0x81, 0x78, 0x8f, 0x19, 0xdb, 0x0f, 0x48, 0x29, 0x52,
	0x20, 0x0c, 0xaa, 0xc0, 0x6d, 0xb8, 0xcd, 0x83, 0xde, 0x66, 0xe9, 0xf7, 0x49, 0xcd, 0xa0, 0x11,
	0xe3, 0x81, 0x36, 0xa8, 0x44, 0x0c, 0x03, 0x7b, 0x1a, 0x19, 0xfc, 0x6b, 0xb8, 0xcd, 0xf2, 0xf5,
	0x09, 0xb3, 0x21, 0x2c, 0x0d, 0x61, 0x79, 0x08, 0xeb, 0xe0, 0x28, 0x69, 0x17, 0x17, 0x5f, 0x67,
	0x4e, 0xaf, 0x9a, 0xb9, 0xfb, 0xd6, 0x6c, 0xd3, 0xa4, 0xff, 0x40, 0xaa, 0x39, 0x74, 0x28, 0x94,
	0xdc, 0x22, 0x0b, 0xfb, 0x21, 0x8f, 0x2c, 0x32, 0xb5, 0x6e, 0x80, 0xb7, 0xc4, 0x53, 0x30, 0x17,
	0x4a, 0x06, 0xc5, 0x8c, 0x71, 0xfa, 0x27, 0xa3, 0x0b, 0xd1, 0x0e, 0x26, 0x77, 0xf8, 0x1d, 0x52,
	0xb6, 0xd5, 0x40, 0x42, 0x68, 0x82, 0xff, 0x7b, 0x03, 0x88, 0xb5, 0x75, 0x21, 0x34, 0xfe, 0x25,
	0x39, 0xfc, 0xf5, 0x4c, 0x81, 0xd7, 0x70, 0x9b, 0x85, 0x5e, 0x65, 0xf7, 0xf6, 0xfe, 0x39, 0xa9,
	0xcc, 0x34, 0xc8, 0xad, 0xa6, 0x94, 0x69, 0xca, 0xe9, 0x5e, 0x2e, 0x69, 0xdf, 0x2d, 0x56, 0xd4,
	0x5d, 0xae, 0xa8, 0xfb, 0xbd, 0xa2, 0xee, 0xfb, 0x9a, 0x3a, 0xcb, 0x35, 0x75, 0x3e, 0xd7, 0xd4,
	0x79, 0xba, 0x8a, 0x47, 0x66, 0x38, 0x0b, 0x59, 0x84, 0x13, 0xde, 0x17, 0x78, 0x6f, 0x67, 0x80,
	0xa7, 0x53, 0xf2, 0x6a, 0xe7, 0xc4, 0xbc, 0x4d, 0x41, 0x87, 0x5e, 0xf6, 0xb9, 0x37, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0xf8, 0x3b, 0xb3, 0x40, 0x02, 0x00, 0x00,
}

func (m *Pledge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pledge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pledge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsedStorage != 0 {
		i = encodeVarintPledge(dAtA, i, uint64(m.UsedStorage))
		i--
		dAtA[i] = 0x38
	}
	if m.TotalStorage != 0 {
		i = encodeVarintPledge(dAtA, i, uint64(m.TotalStorage))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.RewardDebt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TotalShardPledged.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.TotalStoragePledged.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPledge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPledge(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPledge(dAtA []byte, offset int, v uint64) int {
	offset -= sovPledge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pledge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPledge(uint64(l))
	}
	l = m.TotalStoragePledged.Size()
	n += 1 + l + sovPledge(uint64(l))
	l = m.TotalShardPledged.Size()
	n += 1 + l + sovPledge(uint64(l))
	l = m.Reward.Size()
	n += 1 + l + sovPledge(uint64(l))
	l = m.RewardDebt.Size()
	n += 1 + l + sovPledge(uint64(l))
	if m.TotalStorage != 0 {
		n += 1 + sovPledge(uint64(m.TotalStorage))
	}
	if m.UsedStorage != 0 {
		n += 1 + sovPledge(uint64(m.UsedStorage))
	}
	return n
}

func sovPledge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPledge(x uint64) (n int) {
	return sovPledge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pledge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPledge
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
			return fmt.Errorf("proto: Pledge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pledge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStoragePledged", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalStoragePledged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShardPledged", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShardPledged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDebt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
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
				return ErrInvalidLengthPledge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPledge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStorage", wireType)
			}
			m.TotalStorage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalStorage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedStorage", wireType)
			}
			m.UsedStorage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPledge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedStorage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPledge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPledge
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
func skipPledge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPledge
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
					return 0, ErrIntOverflowPledge
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
					return 0, ErrIntOverflowPledge
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
				return 0, ErrInvalidLengthPledge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPledge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPledge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPledge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPledge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPledge = fmt.Errorf("proto: unexpected end of group")
)
