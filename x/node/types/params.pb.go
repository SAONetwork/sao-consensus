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
	ShareThreshold        string     `protobuf:"bytes,6,opt,name=share_threshold,json=shareThreshold,proto3" json:"share_threshold,omitempty"`
	FishmenInfo           string     `protobuf:"bytes,7,opt,name=fishmen_info,json=fishmenInfo,proto3" json:"fishmen_info,omitempty"`
	PenaltyBase           uint64     `protobuf:"varint,8,opt,name=penalty_base,json=penaltyBase,proto3" json:"penalty_base,omitempty"`
	MaxPenalty            uint64     `protobuf:"varint,9,opt,name=max_penalty,json=maxPenalty,proto3" json:"max_penalty,omitempty"`
	VstorageThreshold     int64      `protobuf:"varint,10,opt,name=vstorage_threshold,json=vstorageThreshold,proto3" json:"vstorage_threshold,omitempty"`
	OfflineTriggerHeight  int64      `protobuf:"varint,11,opt,name=offline_trigger_height,json=offlineTriggerHeight,proto3" json:"offline_trigger_height,omitempty"`
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

func (m *Params) GetShareThreshold() string {
	if m != nil {
		return m.ShareThreshold
	}
	return ""
}

func (m *Params) GetFishmenInfo() string {
	if m != nil {
		return m.FishmenInfo
	}
	return ""
}

func (m *Params) GetPenaltyBase() uint64 {
	if m != nil {
		return m.PenaltyBase
	}
	return 0
}

func (m *Params) GetMaxPenalty() uint64 {
	if m != nil {
		return m.MaxPenalty
	}
	return 0
}

func (m *Params) GetVstorageThreshold() int64 {
	if m != nil {
		return m.VstorageThreshold
	}
	return 0
}

func (m *Params) GetOfflineTriggerHeight() int64 {
	if m != nil {
		return m.OfflineTriggerHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "saonetwork.sao.node.Params")
}

func init() { proto.RegisterFile("sao/node/params.proto", fileDescriptor_c6a39a6c05c26f45) }

var fileDescriptor_c6a39a6c05c26f45 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xc7, 0xe3, 0x5f, 0xf3, 0x0b, 0xed, 0xb9, 0x14, 0x7a, 0xb4, 0x60, 0x3a, 0x38, 0x01, 0x09,
	0x88, 0x84, 0xb0, 0x55, 0x40, 0x0c, 0xb0, 0xa0, 0xb0, 0xc0, 0x82, 0x22, 0xd3, 0x05, 0x16, 0xeb,
	0x6c, 0x3f, 0xdb, 0x47, 0xed, 0x7b, 0xd6, 0xdd, 0x25, 0x4d, 0xfe, 0x8b, 0x8e, 0x8c, 0xfc, 0x39,
	0x1d, 0x3b, 0x32, 0x21, 0x94, 0xfc, 0x23, 0xe8, 0xce, 0x4e, 0xb3, 0xb2, 0x59, 0xdf, 0xcf, 0xe7,
	0xfb, 0xf4, 0x7c, 0x7a, 0xe4, 0x58, 0x31, 0x0c, 0x05, 0x66, 0x10, 0x36, 0x4c, 0xb2, 0x5a, 0x05,
	0x8d, 0x44, 0x8d, 0xf4, 0x9e, 0x62, 0x28, 0x40, 0x5f, 0xa0, 0x3c, 0x0f, 0x14, 0xc3, 0xc0, 0x18,
	0x27, 0x47, 0x05, 0x16, 0x68, 0x79, 0x68, 0xbe, 0x5a, 0xf5, 0xc4, 0x4f, 0x51, 0xd5, 0xa8, 0xc2,
	0x84, 0x29, 0x08, 0xe7, 0xa7, 0x09, 0x68, 0x76, 0x1a, 0xa6, 0xc8, 0x45, 0xcb, 0x1f, 0x5f, 0xf6,
	0xc9, 0x60, 0x6a, 0x67, 0xd3, 0x09, 0xd9, 0x4f, 0x2a, 0x4c, 0xcf, 0x63, 0x09, 0x17, 0x4c, 0x66,
	0x9e, 0x33, 0x72, 0xc6, 0xee, 0xcb, 0x87, 0x41, 0x3b, 0x21, 0x30, 0x13, 0x82, 0x6e, 0x42, 0xf0,
	0x01, 0xb9, 0x98, 0xf4, 0xaf, 0x7e, 0x0f, 0x7b, 0x91, 0x6b, 0x4b, 0x91, 0xed, 0xd0, 0x77, 0x64,
	0xd7, 0x78, 0x15, 0x17, 0xe0, 0xfd, 0xf7, 0x6f, 0xfd, 0x9b, 0x02, 0x7d, 0x43, 0x1e, 0x30, 0x21,
	0x66, 0xac, 0x8a, 0x1b, 0x90, 0x29, 0x08, 0xcd, 0x0a, 0x88, 0x97, 0x1c, 0xaa, 0xcc, 0xdb, 0x19,
	0x39, 0xe3, 0xbd, 0xe8, 0xb8, 0xc5, 0xd3, 0x1b, 0xfa, 0xd5, 0x40, 0xfa, 0x84, 0x1c, 0x94, 0xac,
	0x9a, 0x73, 0x51, 0x98, 0x22, 0xc7, 0xcc, 0xeb, 0x8f, 0x9c, 0xf1, 0x4e, 0x74, 0xbb, 0x4b, 0xa7,
	0x36, 0xa4, 0xcf, 0xc9, 0x21, 0xcb, 0xbe, 0xcf, 0x94, 0xae, 0x41, 0xe8, 0x8d, 0xf9, 0xbf, 0x35,
	0xef, 0x6e, 0x41, 0x27, 0x3f, 0x23, 0x77, 0x54, 0xc9, 0x24, 0xc4, 0xba, 0x94, 0xa0, 0x4a, 0xac,
	0x32, 0x6f, 0x60, 0x77, 0x38, 0xb0, 0xf1, 0xd9, 0x26, 0xa5, 0x8f, 0xc8, 0x7e, 0xce, 0x55, 0x59,
	0x83, 0x88, 0xb9, 0xc8, 0xd1, 0xbb, 0x65, 0x2d, 0xb7, 0xcb, 0x3e, 0x89, 0x1c, 0x8d, 0xd2, 0x80,
	0x60, 0x95, 0x5e, 0xc6, 0xe6, 0x5f, 0xbd, 0xdd, 0x91, 0x33, 0xee, 0x47, 0x6e, 0x97, 0x4d, 0x98,
	0x02, 0x3a, 0x24, 0x6e, 0xcd, 0x16, 0x71, 0x17, 0x79, 0x7b, 0xd6, 0x20, 0x35, 0x5b, 0x4c, 0xdb,
	0x84, 0xbe, 0x20, 0x74, 0xae, 0x34, 0x4a, 0xf3, 0x24, 0xdb, 0x95, 0x88, 0xdd, 0xfe, 0x70, 0x43,
	0xb6, 0x5b, 0xbd, 0x26, 0xf7, 0x31, 0xcf, 0xcd, 0xab, 0xc6, 0x5a, 0xf2, 0xa2, 0x00, 0x19, 0x97,
	0xc0, 0x8b, 0x52, 0x7b, 0xae, 0xad, 0x1c, 0x75, 0xf4, 0xac, 0x85, 0x1f, 0x2d, 0x7b, 0xdb, 0xff,
	0xf1, 0x73, 0xd8, 0x9b, 0xbc, 0xbf, 0x5a, 0xf9, 0xce, 0xf5, 0xca, 0x77, 0xfe, 0xac, 0x7c, 0xe7,
	0x72, 0xed, 0xf7, 0xae, 0xd7, 0x7e, 0xef, 0xd7, 0xda, 0xef, 0x7d, 0x7b, 0x5a, 0x70, 0x5d, 0xce,
	0x92, 0x20, 0xc5, 0x3a, 0xfc, 0xc2, 0xf0, 0x73, 0x7b, 0x82, 0xa1, 0x39, 0xd2, 0x45, 0x7b, 0xa6,
	0x7a, 0xd9, 0x80, 0x4a, 0x06, 0xf6, 0xb6, 0x5e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xfe,
	0x5e, 0xd1, 0xbf, 0x02, 0x00, 0x00,
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
	if m.OfflineTriggerHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OfflineTriggerHeight))
		i--
		dAtA[i] = 0x58
	}
	if m.VstorageThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VstorageThreshold))
		i--
		dAtA[i] = 0x50
	}
	if m.MaxPenalty != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPenalty))
		i--
		dAtA[i] = 0x48
	}
	if m.PenaltyBase != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PenaltyBase))
		i--
		dAtA[i] = 0x40
	}
	if len(m.FishmenInfo) > 0 {
		i -= len(m.FishmenInfo)
		copy(dAtA[i:], m.FishmenInfo)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FishmenInfo)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ShareThreshold) > 0 {
		i -= len(m.ShareThreshold)
		copy(dAtA[i:], m.ShareThreshold)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ShareThreshold)))
		i--
		dAtA[i] = 0x32
	}
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
	l = len(m.ShareThreshold)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.FishmenInfo)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.PenaltyBase != 0 {
		n += 1 + sovParams(uint64(m.PenaltyBase))
	}
	if m.MaxPenalty != 0 {
		n += 1 + sovParams(uint64(m.MaxPenalty))
	}
	if m.VstorageThreshold != 0 {
		n += 1 + sovParams(uint64(m.VstorageThreshold))
	}
	if m.OfflineTriggerHeight != 0 {
		n += 1 + sovParams(uint64(m.OfflineTriggerHeight))
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareThreshold", wireType)
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
			m.ShareThreshold = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FishmenInfo", wireType)
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
			m.FishmenInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PenaltyBase", wireType)
			}
			m.PenaltyBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PenaltyBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPenalty", wireType)
			}
			m.MaxPenalty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPenalty |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VstorageThreshold", wireType)
			}
			m.VstorageThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VstorageThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfflineTriggerHeight", wireType)
			}
			m.OfflineTriggerHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OfflineTriggerHeight |= int64(b&0x7F) << shift
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
