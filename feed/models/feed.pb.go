// Code generated by protoc-gen-go.
// source: feed.proto
// DO NOT EDIT!

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	feed.proto

It has these top-level messages:
	Feed
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Feed struct {
	Id     int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Data   string `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	// action count
	LikeCount   int64 `protobuf:"varint,11,opt,name=like_count,json=likeCount" json:"like_count,omitempty"`
	ShareCount  int64 `protobuf:"varint,12,opt,name=share_count,json=shareCount" json:"share_count,omitempty"`
	CreatedUtc  int64 `protobuf:"varint,14,opt,name=created_utc,json=createdUtc" json:"created_utc,omitempty"`
	ModifiedUtc int64 `protobuf:"varint,15,opt,name=modified_utc,json=modifiedUtc" json:"modified_utc,omitempty"`
	Deleted     bool  `protobuf:"varint,16,opt,name=deleted" json:"deleted,omitempty"`
	DeletedUtc  int64 `protobuf:"varint,17,opt,name=deleted_utc,json=deletedUtc" json:"deleted_utc,omitempty"`
}

func (m *Feed) Reset()                    { *m = Feed{} }
func (m *Feed) String() string            { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()               {}
func (*Feed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Feed)(nil), "models.Feed")
}

func init() { proto.RegisterFile("feed.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x50, 0xb9, 0xf4, 0x52, 0x15, 0xb8, 0x05, 0x2f, 0x88, 0xc0, 0x94, 0x89, 0x85,
	0x47, 0x40, 0x42, 0x62, 0x8d, 0xd4, 0x39, 0x32, 0xbe, 0xab, 0xb0, 0x48, 0x31, 0x72, 0x2e, 0x2f,
	0xc0, 0x93, 0x23, 0x5f, 0x9c, 0x6e, 0xf7, 0x7f, 0xff, 0xe7, 0x7f, 0x30, 0xc0, 0x89, 0x99, 0x5e,
	0x7e, 0x53, 0x94, 0x88, 0xe6, 0x1c, 0x89, 0xc7, 0xe9, 0xf9, 0xaf, 0x86, 0xcd, 0x3b, 0x33, 0xe1,
	0x01, 0xea, 0x40, 0xb6, 0x6a, 0xab, 0xee, 0xaa, 0xaf, 0x03, 0xe1, 0x3d, 0x6c, 0xe7, 0x89, 0xd3,
	0x10, 0xc8, 0xd6, 0x0a, 0x4d, 0x8e, 0x1f, 0x84, 0x08, 0x1b, 0x72, 0xe2, 0xac, 0x69, 0xab, 0x6e,
	0xd7, 0xeb, 0x8d, 0x0f, 0x00, 0x63, 0xf8, 0xe6, 0xc1, 0xc7, 0xf9, 0x47, 0x6c, 0xa3, 0xfe, 0x2e,
	0x93, 0xb7, 0x0c, 0xf0, 0x11, 0x9a, 0xe9, 0xcb, 0xa5, 0xb5, 0xdf, 0x6b, 0x0f, 0x8a, 0x2e, 0x82,
	0x4f, 0xec, 0x84, 0x69, 0x98, 0xc5, 0xdb, 0xc3, 0x22, 0x14, 0x74, 0x14, 0x8f, 0x4f, 0xb0, 0x3f,
	0x47, 0x0a, 0xa7, 0x50, 0x8c, 0x1b, 0x35, 0x9a, 0x95, 0x65, 0xc5, 0xc2, 0x96, 0x78, 0x64, 0x61,
	0xb2, 0xb7, 0x6d, 0xd5, 0x5d, 0xf7, 0x6b, 0xcc, 0xeb, 0xe5, 0xd4, 0xb7, 0x77, 0xcb, 0x7a, 0x41,
	0x47, 0xf1, 0x9f, 0x46, 0xff, 0xe4, 0xf5, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xe3, 0x74, 0xed,
	0x21, 0x01, 0x00, 0x00,
}
