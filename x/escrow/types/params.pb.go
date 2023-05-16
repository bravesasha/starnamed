// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/escrow/v1beta1/params.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters of the escrow module
type Params struct {
	ModuleEnabled       bool     `protobuf:"varint,1,opt,name=module_enabled,json=moduleEnabled,proto3" json:"module_enabled,omitempty"`
	AllowedCustomTokens []string `protobuf:"bytes,2,rep,name=allowedCustomTokens,proto3" json:"allowedCustomTokens,omitempty"`
	CustomTokenFee      float64  `protobuf:"fixed64,3,opt,name=customTokenFee,proto3" json:"customTokenFee,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4af5fc7b8f17f7, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "starnamed.x.escrow.v1beta1.Params")
}

func init() { proto.RegisterFile("iov/escrow/v1beta1/params.proto", fileDescriptor_8e4af5fc7b8f17f7) }

var fileDescriptor_8e4af5fc7b8f17f7 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x16, 0x8a, 0x06, 0xec, 0x21, 0x7a, 0xa8, 0x39, 0xac, 0x41, 0x50, 0x02, 0x62,
	0xd6, 0xe2, 0x1b, 0x28, 0x0a, 0xde, 0x24, 0x78, 0xf2, 0x52, 0x76, 0x93, 0x31, 0x06, 0xb3, 0x3b,
	0x21, 0xbb, 0x49, 0xdb, 0x47, 0xf0, 0xe6, 0x63, 0xf5, 0xd8, 0xa3, 0x47, 0x4d, 0x5e, 0x44, 0xcc,
	0xb6, 0x15, 0xc4, 0xdb, 0xfc, 0xdf, 0x3f, 0xff, 0xb2, 0xf3, 0xbb, 0xc7, 0x39, 0x36, 0x0c, 0x74,
	0x52, 0xe1, 0x8c, 0x35, 0x13, 0x01, 0x86, 0x4f, 0x58, 0xc9, 0x2b, 0x2e, 0x75, 0x54, 0x56, 0x68,
	0xd0, 0xf3, 0xb5, 0xe1, 0x95, 0xe2, 0x12, 0xd2, 0x68, 0x1e, 0xd9, 0xc5, 0x68, 0xbd, 0xe8, 0xd3,
	0x04, 0xb5, 0x44, 0xcd, 0x04, 0xd7, 0xb0, 0x4d, 0x27, 0x98, 0x2b, 0x9b, 0xf5, 0x0f, 0x33, 0xcc,
	0xb0, 0x1f, 0xd9, 0xcf, 0xb4, 0xa6, 0x47, 0x19, 0x62, 0x56, 0x00, 0xeb, 0x95, 0xa8, 0x9f, 0x19,
	0x57, 0x8b, 0x8d, 0x65, 0x1f, 0x9c, 0xda, 0x8c, 0x15, 0xd6, 0x3a, 0x79, 0x23, 0xee, 0xf0, 0xa1,
	0xff, 0x98, 0x77, 0xea, 0x8e, 0x24, 0xa6, 0x75, 0x01, 0x53, 0x50, 0x5c, 0x14, 0x90, 0x8e, 0x49,
	0x40, 0xc2, 0xdd, 0x78, 0xdf, 0xd2, 0x5b, 0x0b, 0xbd, 0x4b, 0xf7, 0x80, 0x17, 0x05, 0xce, 0x20,
	0xbd, 0xa9, 0xb5, 0x41, 0xf9, 0x88, 0xaf, 0xa0, 0xf4, 0x78, 0x27, 0x18, 0x84, 0x7b, 0xf1, 0x7f,
	0x96, 0x77, 0xe6, 0x8e, 0x92, 0x5f, 0x7d, 0x07, 0x30, 0x1e, 0x04, 0x24, 0x24, 0xf1, 0x1f, 0x7a,
	0x7d, 0xbf, 0xfc, 0xa2, 0xce, 0xb2, 0xa5, 0x64, 0xd5, 0x52, 0xf2, 0xd9, 0x52, 0xf2, 0xde, 0x51,
	0x67, 0xd5, 0x51, 0xe7, 0xa3, 0xa3, 0xce, 0xd3, 0x79, 0x96, 0x9b, 0x97, 0x5a, 0x44, 0x09, 0x4a,
	0x96, 0x63, 0x73, 0x81, 0x0a, 0xd8, 0xb6, 0x44, 0x36, 0xdf, 0xb4, 0x6d, 0x16, 0x25, 0x68, 0x31,
	0xec, 0xaf, 0xbb, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x63, 0xcf, 0x15, 0x88, 0x01, 0x00,
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
	if m.CustomTokenFee != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CustomTokenFee))))
		i--
		dAtA[i] = 0x19
	}
	if len(m.AllowedCustomTokens) > 0 {
		for iNdEx := len(m.AllowedCustomTokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedCustomTokens[iNdEx])
			copy(dAtA[i:], m.AllowedCustomTokens[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowedCustomTokens[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ModuleEnabled {
		i--
		if m.ModuleEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
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
	if m.ModuleEnabled {
		n += 2
	}
	if len(m.AllowedCustomTokens) > 0 {
		for _, s := range m.AllowedCustomTokens {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.CustomTokenFee != 0 {
		n += 9
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.ModuleEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedCustomTokens", wireType)
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
			m.AllowedCustomTokens = append(m.AllowedCustomTokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomTokenFee", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CustomTokenFee = float64(math.Float64frombits(v))
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
