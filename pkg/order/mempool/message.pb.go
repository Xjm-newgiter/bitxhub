// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package mempool

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	pb "github.com/meshplus/bitxhub-model/pb"
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

type TxSlice struct {
	TxList               []*pb.Transaction `protobuf:"bytes,1,rep,name=TxList,proto3" json:"TxList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TxSlice) Reset()         { *m = TxSlice{} }
func (m *TxSlice) String() string { return proto.CompactTextString(m) }
func (*TxSlice) ProtoMessage()    {}
func (*TxSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *TxSlice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxSlice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxSlice.Merge(m, src)
}
func (m *TxSlice) XXX_Size() int {
	return m.Size()
}
func (m *TxSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_TxSlice.DiscardUnknown(m)
}

var xxx_messageInfo_TxSlice proto.InternalMessageInfo

func (m *TxSlice) GetTxList() []*pb.Transaction {
	if m != nil {
		return m.TxList
	}
	return nil
}

type FetchTxnRequest struct {
	ReplicaId            uint64            `protobuf:"varint,1,opt,name=replicaId,proto3" json:"replicaId,omitempty"`
	Height               uint64            `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	MissingTxHashes      map[uint64]string `protobuf:"bytes,3,rep,name=missing_tx_hashes,json=missingTxHashes,proto3" json:"missing_tx_hashes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FetchTxnRequest) Reset()         { *m = FetchTxnRequest{} }
func (m *FetchTxnRequest) String() string { return proto.CompactTextString(m) }
func (*FetchTxnRequest) ProtoMessage()    {}
func (*FetchTxnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}
func (m *FetchTxnRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchTxnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchTxnRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchTxnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTxnRequest.Merge(m, src)
}
func (m *FetchTxnRequest) XXX_Size() int {
	return m.Size()
}
func (m *FetchTxnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTxnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTxnRequest proto.InternalMessageInfo

func (m *FetchTxnRequest) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *FetchTxnRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FetchTxnRequest) GetMissingTxHashes() map[uint64]string {
	if m != nil {
		return m.MissingTxHashes
	}
	return nil
}

type FetchTxnResponse struct {
	ReplicaId            uint64                     `protobuf:"varint,1,opt,name=replicaId,proto3" json:"replicaId,omitempty"`
	Height               uint64                     `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	MissingTxnList       map[uint64]*pb.Transaction `protobuf:"bytes,3,rep,name=missing_txn_list,json=missingTxnList,proto3" json:"missing_txn_list,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FetchTxnResponse) Reset()         { *m = FetchTxnResponse{} }
func (m *FetchTxnResponse) String() string { return proto.CompactTextString(m) }
func (*FetchTxnResponse) ProtoMessage()    {}
func (*FetchTxnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}
func (m *FetchTxnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchTxnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchTxnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchTxnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTxnResponse.Merge(m, src)
}
func (m *FetchTxnResponse) XXX_Size() int {
	return m.Size()
}
func (m *FetchTxnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTxnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTxnResponse proto.InternalMessageInfo

func (m *FetchTxnResponse) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *FetchTxnResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FetchTxnResponse) GetMissingTxnList() map[uint64]*pb.Transaction {
	if m != nil {
		return m.MissingTxnList
	}
	return nil
}

func init() {
	proto.RegisterType((*TxSlice)(nil), "mempool.tx_slice")
	proto.RegisterType((*FetchTxnRequest)(nil), "mempool.fetch_txn_request")
	proto.RegisterMapType((map[uint64]string)(nil), "mempool.fetch_txn_request.MissingTxHashesEntry")
	proto.RegisterType((*FetchTxnResponse)(nil), "mempool.fetch_txn_response")
	proto.RegisterMapType((map[uint64]*pb.Transaction)(nil), "mempool.fetch_txn_response.MissingTxnListEntry")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xc7, 0x99, 0xf6, 0xde, 0xde, 0xdb, 0xf1, 0xa3, 0xed, 0x58, 0x24, 0x14, 0x29, 0xa5, 0x20,
	0x76, 0x63, 0x02, 0x16, 0x44, 0x5c, 0x16, 0x04, 0x05, 0xdd, 0x84, 0x6e, 0xc4, 0x45, 0x48, 0xd2,
	0x63, 0x66, 0x30, 0xf3, 0x61, 0x66, 0x22, 0xe9, 0x1b, 0xba, 0xf4, 0x11, 0xa4, 0x4b, 0xdf, 0xc0,
	0x9d, 0x34, 0x09, 0xa4, 0xc5, 0xe0, 0xc2, 0xdd, 0x9c, 0x8f, 0xff, 0xfc, 0xcf, 0xef, 0xcc, 0xe0,
	0x3d, 0x0e, 0x5a, 0xfb, 0x11, 0xd8, 0x2a, 0x91, 0x46, 0x92, 0x7f, 0x1c, 0xb8, 0x92, 0x32, 0x1e,
	0x9c, 0x47, 0xcc, 0xd0, 0x34, 0xb0, 0x43, 0xc9, 0x1d, 0x0e, 0x9a, 0xaa, 0x38, 0xd5, 0x4e, 0xc0,
	0x4c, 0x46, 0xd3, 0xe0, 0x94, 0xcb, 0x05, 0xc4, 0x8e, 0x0a, 0x1c, 0x93, 0xf8, 0x42, 0xfb, 0xa1,
	0x61, 0x52, 0x14, 0x17, 0x8c, 0xa7, 0xf8, 0xbf, 0xc9, 0x3c, 0x1d, 0xb3, 0x10, 0xc8, 0x09, 0x6e,
	0xcd, 0xb3, 0x5b, 0xa6, 0x8d, 0x85, 0x46, 0xcd, 0xc9, 0xce, 0x59, 0xc7, 0x56, 0x81, 0x3d, 0xaf,
	0x24, 0x6e, 0x59, 0x1e, 0x7f, 0x20, 0xdc, 0x7b, 0x04, 0x13, 0x52, 0xcf, 0x64, 0xc2, 0x4b, 0xe0,
	0x39, 0x05, 0x6d, 0xc8, 0x11, 0x6e, 0x27, 0xa0, 0x62, 0x16, 0xfa, 0x37, 0x0b, 0x0b, 0x8d, 0xd0,
	0xe4, 0x8f, 0x5b, 0x25, 0xc8, 0x21, 0x6e, 0x51, 0x60, 0x11, 0x35, 0x56, 0x23, 0x2f, 0x95, 0x11,
	0x79, 0xc0, 0x3d, 0xce, 0xb4, 0x66, 0x22, 0xf2, 0x4c, 0xe6, 0x51, 0x5f, 0x53, 0xd0, 0x56, 0x33,
	0xf7, 0x77, 0xec, 0x92, 0xce, 0xfe, 0x66, 0x66, 0xdf, 0x15, 0x9a, 0x79, 0x76, 0x9d, 0x2b, 0xae,
	0x84, 0x49, 0x96, 0x6e, 0x87, 0x6f, 0x67, 0x07, 0x33, 0xdc, 0xaf, 0x6b, 0x24, 0x5d, 0xdc, 0x7c,
	0x82, 0x65, 0x39, 0xe4, 0xfa, 0x48, 0xfa, 0xf8, 0xef, 0x8b, 0x1f, 0xa7, 0x90, 0x4f, 0xd7, 0x76,
	0x8b, 0xe0, 0xb2, 0x71, 0x81, 0xc6, 0x9f, 0x08, 0x93, 0x4d, 0x7f, 0xad, 0xa4, 0xd0, 0xf0, 0x4b,
	0xda, 0x7b, 0xdc, 0xad, 0x68, 0x85, 0x17, 0xaf, 0x97, 0xfd, 0x13, 0x6c, 0x61, 0x56, 0xd1, 0x8a,
	0xf5, 0xfe, 0x0b, 0xd8, 0x7d, 0xbe, 0x95, 0x1c, 0xb8, 0xf8, 0xa0, 0xa6, 0xad, 0x06, 0xf5, 0x78,
	0x13, 0xb5, 0xe6, 0x95, 0x2b, 0xf6, 0xd9, 0xee, 0xeb, 0x6a, 0x88, 0xde, 0x56, 0x43, 0xf4, 0xbe,
	0x1a, 0xa2, 0xa0, 0x95, 0x7f, 0x99, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc0, 0xd7,
	0xb5, 0x84, 0x02, 0x00, 0x00,
}

func (m *TxSlice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxSlice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxSlice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TxList) > 0 {
		for iNdEx := len(m.TxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FetchTxnRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchTxnRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchTxnRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MissingTxHashes) > 0 {
		for k := range m.MissingTxHashes {
			v := m.MissingTxHashes[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintMessage(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i = encodeVarintMessage(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMessage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.ReplicaId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ReplicaId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FetchTxnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchTxnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchTxnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MissingTxnList) > 0 {
		for k := range m.MissingTxnList {
			v := m.MissingTxnList[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintMessage(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintMessage(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMessage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.ReplicaId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ReplicaId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxSlice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxList) > 0 {
		for _, e := range m.TxList {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FetchTxnRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReplicaId != 0 {
		n += 1 + sovMessage(uint64(m.ReplicaId))
	}
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if len(m.MissingTxHashes) > 0 {
		for k, v := range m.MissingTxHashes {
			_ = k
			_ = v
			mapEntrySize := 1 + sovMessage(uint64(k)) + 1 + len(v) + sovMessage(uint64(len(v)))
			n += mapEntrySize + 1 + sovMessage(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FetchTxnResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReplicaId != 0 {
		n += 1 + sovMessage(uint64(m.ReplicaId))
	}
	if m.Height != 0 {
		n += 1 + sovMessage(uint64(m.Height))
	}
	if len(m.MissingTxnList) > 0 {
		for k, v := range m.MissingTxnList {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovMessage(uint64(l))
			}
			mapEntrySize := 1 + sovMessage(uint64(k)) + l
			n += mapEntrySize + 1 + sovMessage(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxSlice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: tx_slice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: tx_slice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxList = append(m.TxList, &pb.Transaction{})
			if err := m.TxList[len(m.TxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FetchTxnRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: fetch_txn_request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: fetch_txn_request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaId", wireType)
			}
			m.ReplicaId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicaId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissingTxHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MissingTxHashes == nil {
				m.MissingTxHashes = make(map[uint64]string)
			}
			var mapkey uint64
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
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
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthMessage
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthMessage
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMessage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMessage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MissingTxHashes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FetchTxnResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: fetch_txn_response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: fetch_txn_response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaId", wireType)
			}
			m.ReplicaId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicaId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissingTxnList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MissingTxnList == nil {
				m.MissingTxnList = make(map[uint64]*pb.Transaction)
			}
			var mapkey uint64
			var mapvalue *pb.Transaction
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
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
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
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
						return ErrInvalidLengthMessage
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMessage
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &pb.Transaction{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMessage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMessage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MissingTxnList[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)