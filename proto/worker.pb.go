// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SealCommit2Request struct {
	Sector               *SectorID `protobuf:"bytes,1,opt,name=Sector,proto3" json:"Sector,omitempty"`
	Commit1Out           []byte    `protobuf:"bytes,2,opt,name=Commit1Out,proto3" json:"Commit1Out,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SealCommit2Request) Reset()         { *m = SealCommit2Request{} }
func (m *SealCommit2Request) String() string { return proto.CompactTextString(m) }
func (*SealCommit2Request) ProtoMessage()    {}
func (*SealCommit2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{0}
}

func (m *SealCommit2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealCommit2Request.Unmarshal(m, b)
}
func (m *SealCommit2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealCommit2Request.Marshal(b, m, deterministic)
}
func (m *SealCommit2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealCommit2Request.Merge(m, src)
}
func (m *SealCommit2Request) XXX_Size() int {
	return xxx_messageInfo_SealCommit2Request.Size(m)
}
func (m *SealCommit2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_SealCommit2Request.DiscardUnknown(m)
}

var xxx_messageInfo_SealCommit2Request proto.InternalMessageInfo

func (m *SealCommit2Request) GetSector() *SectorID {
	if m != nil {
		return m.Sector
	}
	return nil
}

func (m *SealCommit2Request) GetCommit1Out() []byte {
	if m != nil {
		return m.Commit1Out
	}
	return nil
}

type SealCommit2Response struct {
	Proof                []byte   `protobuf:"bytes,1,opt,name=Proof,proto3" json:"Proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SealCommit2Response) Reset()         { *m = SealCommit2Response{} }
func (m *SealCommit2Response) String() string { return proto.CompactTextString(m) }
func (*SealCommit2Response) ProtoMessage()    {}
func (*SealCommit2Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{1}
}

func (m *SealCommit2Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealCommit2Response.Unmarshal(m, b)
}
func (m *SealCommit2Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealCommit2Response.Marshal(b, m, deterministic)
}
func (m *SealCommit2Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealCommit2Response.Merge(m, src)
}
func (m *SealCommit2Response) XXX_Size() int {
	return xxx_messageInfo_SealCommit2Response.Size(m)
}
func (m *SealCommit2Response) XXX_DiscardUnknown() {
	xxx_messageInfo_SealCommit2Response.DiscardUnknown(m)
}

var xxx_messageInfo_SealCommit2Response proto.InternalMessageInfo

func (m *SealCommit2Response) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type SectorID struct {
	Miner                uint64   `protobuf:"varint,1,opt,name=Miner,proto3" json:"Miner,omitempty"`
	Number               uint64   `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SectorID) Reset()         { *m = SectorID{} }
func (m *SectorID) String() string { return proto.CompactTextString(m) }
func (*SectorID) ProtoMessage()    {}
func (*SectorID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{2}
}

func (m *SectorID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SectorID.Unmarshal(m, b)
}
func (m *SectorID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SectorID.Marshal(b, m, deterministic)
}
func (m *SectorID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SectorID.Merge(m, src)
}
func (m *SectorID) XXX_Size() int {
	return xxx_messageInfo_SectorID.Size(m)
}
func (m *SectorID) XXX_DiscardUnknown() {
	xxx_messageInfo_SectorID.DiscardUnknown(m)
}

var xxx_messageInfo_SectorID proto.InternalMessageInfo

func (m *SectorID) GetMiner() uint64 {
	if m != nil {
		return m.Miner
	}
	return 0
}

func (m *SectorID) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*SealCommit2Request)(nil), "proto.SealCommit2Request")
	proto.RegisterType((*SealCommit2Response)(nil), "proto.SealCommit2Response")
	proto.RegisterType((*SectorID)(nil), "proto.SectorID")
}

func init() { proto.RegisterFile("worker.proto", fileDescriptor_e4ff6184b07e587a) }

var fileDescriptor_e4ff6184b07e587a = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0x26, 0xd2, 0x06, 0x99, 0x46, 0x84, 0x55, 0xa4, 0xf6, 0x20, 0x25, 0x97, 0x56, 0xa4, 0x09,
	0xc6, 0x8b, 0x67, 0xed, 0xc5, 0x83, 0x0f, 0x36, 0x88, 0x20, 0x78, 0x30, 0x61, 0xd4, 0xc5, 0x26,
	0x13, 0x77, 0x67, 0xf5, 0xef, 0x8b, 0x3b, 0x55, 0xa2, 0x78, 0x1a, 0xbe, 0x99, 0xfd, 0x5e, 0x0b,
	0xc9, 0x07, 0xd9, 0x57, 0xb4, 0x59, 0x67, 0x89, 0x49, 0x0d, 0xc3, 0x48, 0x1f, 0x40, 0x95, 0xf8,
	0xb8, 0x3a, 0xa7, 0xa6, 0x31, 0x5c, 0x68, 0x7c, 0xf3, 0xe8, 0x58, 0xcd, 0x20, 0x2e, 0xb1, 0x66,
	0xb2, 0xe3, 0x68, 0x1a, 0xcd, 0x47, 0xc5, 0xb6, 0x90, 0x32, 0x59, 0x5e, 0x2c, 0xf5, 0xfa, 0xac,
	0x0e, 0x00, 0x84, 0x7a, 0x7c, 0xed, 0x79, 0xbc, 0x31, 0x8d, 0xe6, 0x89, 0xee, 0x6d, 0xd2, 0x23,
	0xd8, 0xf9, 0x25, 0xef, 0x3a, 0x6a, 0x1d, 0xaa, 0x5d, 0x18, 0xde, 0x58, 0xa2, 0xa7, 0x20, 0x9f,
	0x68, 0x01, 0xe9, 0x29, 0x6c, 0x7e, 0x1b, 0x7c, 0xbd, 0xb8, 0x34, 0x2d, 0x4a, 0x80, 0x81, 0x16,
	0xa0, 0xf6, 0x20, 0xbe, 0xf2, 0x4d, 0x85, 0x36, 0x58, 0x0d, 0xf4, 0x1a, 0x15, 0xb7, 0xb0, 0x75,
	0x17, 0xca, 0x95, 0x68, 0xdf, 0x4d, 0x8d, 0x6a, 0x09, 0xa3, 0x9e, 0xaf, 0xda, 0xff, 0xc9, 0xff,
	0xb7, 0xea, 0x64, 0xf2, 0xdf, 0x49, 0x62, 0x9e, 0x1d, 0xde, 0xcf, 0x9e, 0x0d, 0xbf, 0xf8, 0x2a,
	0xab, 0xa9, 0xc9, 0x9d, 0x73, 0x6e, 0xc1, 0xd4, 0xe5, 0xf2, 0x8f, 0x8b, 0x7a, 0x65, 0xb0, 0xe5,
	0x3c, 0xd0, 0xab, 0x38, 0x8c, 0x93, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x54, 0xe4, 0xdc,
	0x65, 0x01, 0x00, 0x00,
}
