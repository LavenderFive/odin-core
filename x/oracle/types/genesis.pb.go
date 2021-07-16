// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params        Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DataSources   []DataSource   `protobuf:"bytes,2,rep,name=data_sources,json=dataSources,proto3" json:"data_sources"`
	OracleScripts []OracleScript `protobuf:"bytes,3,rep,name=oracle_scripts,json=oracleScripts,proto3" json:"oracle_scripts"`
	OraclePool    OraclePool     `protobuf:"bytes,4,opt,name=oracle_pool,json=oraclePool,proto3" json:"oracle_pool"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14b982a0a6345d1d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDataSources() []DataSource {
	if m != nil {
		return m.DataSources
	}
	return nil
}

func (m *GenesisState) GetOracleScripts() []OracleScript {
	if m != nil {
		return m.OracleScripts
	}
	return nil
}

func (m *GenesisState) GetOraclePool() OraclePool {
	if m != nil {
		return m.OraclePool
	}
	return OraclePool{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("oracle/v1/genesis.proto", fileDescriptor_14b982a0a6345d1d) }

var fileDescriptor_14b982a0a6345d1d = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x93, 0xb6, 0x14, 0xbc, 0x54, 0xc1, 0x60, 0x6d, 0xe8, 0x70, 0x16, 0xa7, 0x2e, 0xcd,
	0x91, 0xba, 0x8a, 0x43, 0x09, 0x74, 0x50, 0xb0, 0x98, 0xcd, 0xa5, 0x5c, 0x93, 0x23, 0x1e, 0x24,
	0x7d, 0xe1, 0xee, 0x5a, 0x74, 0xf6, 0x1f, 0xf0, 0xcf, 0xea, 0xd8, 0xd1, 0x49, 0x24, 0xf9, 0x47,
	0xc4, 0xcb, 0xd1, 0x06, 0xba, 0x3d, 0xbe, 0x3f, 0x3e, 0xef, 0xf1, 0xd0, 0x00, 0x04, 0x8d, 0x33,
	0x46, 0xb6, 0x01, 0x49, 0xd9, 0x9a, 0x49, 0x2e, 0xfd, 0x42, 0x80, 0x02, 0xf7, 0xac, 0x36, 0xfc,
	0x6d, 0x30, 0xbc, 0x4a, 0x21, 0x05, 0xad, 0x92, 0xff, 0xa9, 0x0e, 0x0c, 0xaf, 0x8f, 0x4d, 0x13,
	0x3d, 0xd1, 0x0b, 0x2a, 0x68, 0x6e, 0x80, 0xb7, 0x9f, 0x2d, 0xd4, 0x9b, 0xd7, 0x2b, 0x22, 0x45,
	0x15, 0x73, 0x09, 0xea, 0xd6, 0x01, 0xcf, 0x1e, 0xd9, 0x63, 0x67, 0x7a, 0xe9, 0x1f, 0x56, 0xfa,
	0x0b, 0x6d, 0xcc, 0x3a, 0xbb, 0x9f, 0x1b, 0xeb, 0xc5, 0xc4, 0xdc, 0x07, 0xd4, 0x4b, 0xa8, 0xa2,
	0x4b, 0x09, 0x1b, 0x11, 0x33, 0xe9, 0xb5, 0x46, 0xed, 0xb1, 0x33, 0xed, 0x37, 0x6a, 0x21, 0x55,
	0x34, 0xd2, 0xae, 0xa9, 0x3a, 0xc9, 0x41, 0x91, 0x6e, 0x88, 0x2e, 0xea, 0xe8, 0x52, 0xc6, 0x82,
	0x17, 0x4a, 0x7a, 0x6d, 0x4d, 0x18, 0x34, 0x08, 0xcf, 0x7a, 0x8a, 0xb4, 0x6f, 0x18, 0xe7, 0xd0,
	0xd0, 0xa4, 0x7b, 0x8f, 0x1c, 0x43, 0x29, 0x00, 0x32, 0xaf, 0xa3, 0x6f, 0xef, 0x9f, 0x20, 0x16,
	0x00, 0x99, 0x01, 0x20, 0x38, 0x2a, 0x8f, 0xbb, 0x12, 0xdb, 0xfb, 0x12, 0xdb, 0xbf, 0x25, 0xb6,
	0xbf, 0x2a, 0x6c, 0xed, 0x2b, 0x6c, 0x7d, 0x57, 0xd8, 0x7a, 0x0d, 0x52, 0xae, 0xde, 0x36, 0x2b,
	0x3f, 0x86, 0x9c, 0xcc, 0x19, 0x84, 0xb3, 0xc9, 0x13, 0xcf, 0xb9, 0x62, 0x09, 0x81, 0x84, 0xaf,
	0x27, 0x31, 0x08, 0x46, 0xde, 0xcd, 0xab, 0x89, 0xfa, 0x28, 0x98, 0x5c, 0x75, 0xf5, 0x67, 0xef,
	0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xbf, 0x80, 0x6c, 0xc5, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.OraclePool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.OracleScripts) > 0 {
		for iNdEx := len(m.OracleScripts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OracleScripts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DataSources) > 0 {
		for iNdEx := len(m.DataSources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataSources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DataSources) > 0 {
		for _, e := range m.DataSources {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OracleScripts) > 0 {
		for _, e := range m.OracleScripts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.OraclePool.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataSources = append(m.DataSources, DataSource{})
			if err := m.DataSources[len(m.DataSources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScripts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleScripts = append(m.OracleScripts, OracleScript{})
			if err := m.OracleScripts[len(m.OracleScripts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OraclePool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OraclePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
