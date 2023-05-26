// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sao/node/pool.proto

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

type Pool struct {
	TotalPledged       types.Coin    `protobuf:"bytes,1,opt,name=total_pledged,json=totalPledged,proto3" json:"total_pledged"`
	TotalReward        types.Coin    `protobuf:"bytes,2,opt,name=total_reward,json=totalReward,proto3" json:"total_reward"`
	AccPledgePerByte   types.DecCoin `protobuf:"bytes,3,opt,name=acc_pledge_per_byte,json=accPledgePerByte,proto3" json:"acc_pledge_per_byte"`
	AccRewardPerByte   types.DecCoin `protobuf:"bytes,4,opt,name=acc_reward_per_byte,json=accRewardPerByte,proto3" json:"acc_reward_per_byte"`
	RewardPerBlock     types.DecCoin `protobuf:"bytes,5,opt,name=reward_per_block,json=rewardPerBlock,proto3" json:"reward_per_block"`
	NextRewardPerBlock types.DecCoin `protobuf:"bytes,6,opt,name=next_reward_per_block,json=nextRewardPerBlock,proto3" json:"next_reward_per_block"`
	TotalStorage       int64         `protobuf:"varint,7,opt,name=total_storage,json=totalStorage,proto3" json:"total_storage,omitempty"`
	RewardedBlockCount int64         `protobuf:"varint,8,opt,name=rewarded_block_count,json=rewardedBlockCount,proto3" json:"rewarded_block_count,omitempty"`
	PendingStorage     int64         `protobuf:"varint,9,opt,name=pending_storage,json=pendingStorage,proto3" json:"pending_storage,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d299d3261ef03a4f, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetTotalPledged() types.Coin {
	if m != nil {
		return m.TotalPledged
	}
	return types.Coin{}
}

func (m *Pool) GetTotalReward() types.Coin {
	if m != nil {
		return m.TotalReward
	}
	return types.Coin{}
}

func (m *Pool) GetAccPledgePerByte() types.DecCoin {
	if m != nil {
		return m.AccPledgePerByte
	}
	return types.DecCoin{}
}

func (m *Pool) GetAccRewardPerByte() types.DecCoin {
	if m != nil {
		return m.AccRewardPerByte
	}
	return types.DecCoin{}
}

func (m *Pool) GetRewardPerBlock() types.DecCoin {
	if m != nil {
		return m.RewardPerBlock
	}
	return types.DecCoin{}
}

func (m *Pool) GetNextRewardPerBlock() types.DecCoin {
	if m != nil {
		return m.NextRewardPerBlock
	}
	return types.DecCoin{}
}

func (m *Pool) GetTotalStorage() int64 {
	if m != nil {
		return m.TotalStorage
	}
	return 0
}

func (m *Pool) GetRewardedBlockCount() int64 {
	if m != nil {
		return m.RewardedBlockCount
	}
	return 0
}

func (m *Pool) GetPendingStorage() int64 {
	if m != nil {
		return m.PendingStorage
	}
	return 0
}

func init() {
	proto.RegisterType((*Pool)(nil), "saonetwork.sao.node.Pool")
}

func init() { proto.RegisterFile("sao/node/pool.proto", fileDescriptor_d299d3261ef03a4f) }

var fileDescriptor_d299d3261ef03a4f = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x07, 0xf0, 0x84, 0x85, 0x01, 0x1e, 0x8c, 0xc9, 0x1b, 0x52, 0x98, 0x50, 0x98, 0x40, 0x82,
	0x9d, 0x6c, 0x06, 0x5f, 0x00, 0x65, 0x3b, 0x22, 0x54, 0x5a, 0x71, 0xe1, 0x12, 0x39, 0xce, 0x53,
	0x88, 0x9a, 0xfa, 0x45, 0x8e, 0x4b, 0xdb, 0x6f, 0xc1, 0xc7, 0xea, 0xb1, 0x47, 0x24, 0x24, 0x84,
	0xda, 0x2f, 0x82, 0x6c, 0x27, 0xa5, 0x08, 0x0e, 0xe5, 0x96, 0xf8, 0xbd, 0xff, 0xcf, 0x4f, 0xf2,
	0x23, 0xa7, 0xad, 0x40, 0xae, 0xb0, 0x00, 0xde, 0x20, 0xd6, 0xac, 0xd1, 0x68, 0x90, 0xda, 0x43,
	0x05, 0x66, 0x86, 0x7a, 0xcc, 0x5a, 0x81, 0xcc, 0xd6, 0xcf, 0xcf, 0x4a, 0x2c, 0xd1, 0xd5, 0xb9,
	0xfd, 0xf2, 0xad, 0xe7, 0x89, 0xc4, 0x76, 0x82, 0x2d, 0xcf, 0x45, 0x0b, 0xfc, 0xcb, 0x55, 0x0e,
	0x46, 0x5c, 0x71, 0x89, 0x95, 0xf2, 0xf5, 0x67, 0xdf, 0x23, 0x12, 0x0d, 0x10, 0x6b, 0x7a, 0x43,
	0x1e, 0x18, 0x34, 0xa2, 0xce, 0x9a, 0x1a, 0x8a, 0x12, 0x8a, 0x38, 0xbc, 0x08, 0x2f, 0x8f, 0x5e,
	0x3f, 0x66, 0x1e, 0x60, 0x16, 0x60, 0x1d, 0xc0, 0xae, 0xb1, 0x52, 0x69, 0xb4, 0xfc, 0xf1, 0x34,
	0x18, 0xde, 0x77, 0xa9, 0x81, 0x0f, 0xd1, 0x94, 0xf8, 0xff, 0x4c, 0xc3, 0x4c, 0xe8, 0x22, 0xbe,
	0xb5, 0x1f, 0x72, 0xe4, 0x42, 0x43, 0x97, 0xa1, 0x1f, 0xc8, 0xa9, 0x90, 0xb2, 0x9b, 0x23, 0x6b,
	0x40, 0x67, 0xf9, 0xc2, 0x40, 0x7c, 0xe0, 0xa8, 0x27, 0xff, 0xa4, 0x6e, 0x40, 0xee, 0x68, 0x27,
	0x42, 0x4a, 0x3f, 0xd0, 0x00, 0x74, 0xba, 0x30, 0xd0, 0x93, 0x7e, 0xa8, 0xdf, 0x64, 0xf4, 0x5f,
	0xa4, 0x1f, 0xaf, 0x27, 0xdf, 0x91, 0x93, 0x5d, 0xae, 0x46, 0x39, 0x8e, 0x6f, 0xef, 0xed, 0x1d,
	0xeb, 0x2d, 0x66, 0x93, 0xf4, 0x23, 0x79, 0xa4, 0x60, 0x6e, 0xb2, 0xbf, 0xc8, 0xc3, 0xbd, 0x49,
	0x6a, 0x81, 0xe1, 0x9f, 0xec, 0xf3, 0xfe, 0x51, 0x5b, 0x83, 0x5a, 0x94, 0x10, 0xdf, 0xb9, 0x08,
	0x2f, 0x0f, 0xba, 0x37, 0x1b, 0xf9, 0x33, 0xfa, 0x8a, 0x9c, 0xf9, 0x6b, 0xa1, 0xf0, 0x97, 0x66,
	0x12, 0xa7, 0xca, 0xc4, 0x77, 0x5d, 0x2f, 0xed, 0x6b, 0x4e, 0xbc, 0xb6, 0x15, 0xfa, 0x92, 0x3c,
	0x6c, 0x40, 0x15, 0x95, 0x2a, 0xb7, 0xf0, 0x3d, 0xd7, 0x7c, 0xdc, 0x1d, 0x77, 0x74, 0xfa, 0x76,
	0xb9, 0x4e, 0xc2, 0xd5, 0x3a, 0x09, 0x7f, 0xae, 0x93, 0xf0, 0xeb, 0x26, 0x09, 0x56, 0x9b, 0x24,
	0xf8, 0xb6, 0x49, 0x82, 0x4f, 0x2f, 0xca, 0xca, 0x7c, 0x9e, 0xe6, 0x4c, 0xe2, 0x84, 0x8f, 0x04,
	0xbe, 0xf7, 0xdb, 0xcc, 0xed, 0xb6, 0xcf, 0xfd, 0xbe, 0x9b, 0x45, 0x03, 0x6d, 0x7e, 0xe8, 0xd6,
	0xf4, 0xcd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x6a, 0xbe, 0xfa, 0x08, 0x03, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PendingStorage != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PendingStorage))
		i--
		dAtA[i] = 0x48
	}
	if m.RewardedBlockCount != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.RewardedBlockCount))
		i--
		dAtA[i] = 0x40
	}
	if m.TotalStorage != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.TotalStorage))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.NextRewardPerBlock.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.RewardPerBlock.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.AccRewardPerByte.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AccPledgePerByte.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.TotalReward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.TotalPledged.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalPledged.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TotalReward.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.AccPledgePerByte.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.AccRewardPerByte.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.RewardPerBlock.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.NextRewardPerBlock.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.TotalStorage != 0 {
		n += 1 + sovPool(uint64(m.TotalStorage))
	}
	if m.RewardedBlockCount != 0 {
		n += 1 + sovPool(uint64(m.RewardedBlockCount))
	}
	if m.PendingStorage != 0 {
		n += 1 + sovPool(uint64(m.PendingStorage))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPledged", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalPledged.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccPledgePerByte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccPledgePerByte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccRewardPerByte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccRewardPerByte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRewardPerBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NextRewardPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStorage", wireType)
			}
			m.TotalStorage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardedBlockCount", wireType)
			}
			m.RewardedBlockCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardedBlockCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingStorage", wireType)
			}
			m.PendingStorage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingStorage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
