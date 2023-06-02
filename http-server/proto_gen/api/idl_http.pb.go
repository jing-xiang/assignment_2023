// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idl_http.proto

package api

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

type Message struct {
	Chat                 string   `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Sender               string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	SendTime             int64    `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ff43e8df49a8be, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetChat() string {
	if m != nil {
		return m.Chat
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type SendRequest struct {
	Chat                 string   `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Sender               string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ff43e8df49a8be, []int{1}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetChat() string {
	if m != nil {
		return m.Chat
	}
	return ""
}

func (m *SendRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SendRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ff43e8df49a8be, []int{2}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

type PullRequest struct {
	Chat                 string   `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	Cursor               int64    `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Reverse              bool     `protobuf:"varint,4,opt,name=reverse,proto3" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullRequest) Reset()         { *m = PullRequest{} }
func (m *PullRequest) String() string { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()    {}
func (*PullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ff43e8df49a8be, []int{3}
}

func (m *PullRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequest.Unmarshal(m, b)
}
func (m *PullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequest.Marshal(b, m, deterministic)
}
func (m *PullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequest.Merge(m, src)
}
func (m *PullRequest) XXX_Size() int {
	return xxx_messageInfo_PullRequest.Size(m)
}
func (m *PullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequest proto.InternalMessageInfo

func (m *PullRequest) GetChat() string {
	if m != nil {
		return m.Chat
	}
	return ""
}

func (m *PullRequest) GetCursor() int64 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *PullRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PullRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

type PullResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	HasMore              bool       `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	NextCursor           int64      `protobuf:"varint,3,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PullResponse) Reset()         { *m = PullResponse{} }
func (m *PullResponse) String() string { return proto.CompactTextString(m) }
func (*PullResponse) ProtoMessage()    {}
func (*PullResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ff43e8df49a8be, []int{4}
}

func (m *PullResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullResponse.Unmarshal(m, b)
}
func (m *PullResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullResponse.Marshal(b, m, deterministic)
}
func (m *PullResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullResponse.Merge(m, src)
}
func (m *PullResponse) XXX_Size() int {
	return xxx_messageInfo_PullResponse.Size(m)
}
func (m *PullResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullResponse proto.InternalMessageInfo

func (m *PullResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *PullResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *PullResponse) GetNextCursor() int64 {
	if m != nil {
		return m.NextCursor
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "api.Message")
	proto.RegisterType((*SendRequest)(nil), "api.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "api.SendResponse")
	proto.RegisterType((*PullRequest)(nil), "api.PullRequest")
	proto.RegisterType((*PullResponse)(nil), "api.PullResponse")
}

func init() {
	proto.RegisterFile("idl_http.proto", fileDescriptor_94ff43e8df49a8be)
}

var fileDescriptor_94ff43e8df49a8be = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xc9, 0x9b, 0x34, 0x4d, 0x6f, 0x4a, 0x79, 0xbd, 0x48, 0x89, 0xba, 0xb0, 0x64, 0x15,
	0x28, 0x54, 0xa8, 0xff, 0x40, 0xd7, 0x05, 0x99, 0xba, 0x72, 0x13, 0xc6, 0xf6, 0x6a, 0x46, 0xf2,
	0xe5, 0xcc, 0xa4, 0xf4, 0xe7, 0xcb, 0x7c, 0x28, 0x59, 0xb9, 0x71, 0x77, 0xcf, 0x93, 0xe1, 0x9e,
	0x73, 0x0f, 0x81, 0x85, 0x38, 0xd6, 0x65, 0xa5, 0x75, 0xbf, 0xe9, 0x65, 0xa7, 0x3b, 0x0c, 0x79,
	0x2f, 0xf2, 0x37, 0x98, 0xee, 0x48, 0x29, 0xfe, 0x4e, 0x88, 0x10, 0x1d, 0x2a, 0xae, 0xb3, 0x60,
	0x15, 0x14, 0x33, 0x66, 0x67, 0xc3, 0x34, 0x9d, 0x75, 0xf6, 0xcf, 0x31, 0x33, 0xe3, 0x12, 0x62,
	0x45, 0xed, 0x91, 0x64, 0x16, 0x5a, 0xea, 0x15, 0xde, 0xc0, 0xcc, 0x4c, 0xa5, 0x16, 0x0d, 0x65,
	0xd1, 0x2a, 0x28, 0x42, 0x96, 0x18, 0xf0, 0x2c, 0x1a, 0xca, 0x77, 0x90, 0xee, 0xa9, 0x3d, 0x32,
	0xfa, 0x1c, 0x48, 0xe9, 0xbf, 0x7a, 0xe5, 0x0b, 0x98, 0xbb, 0x75, 0xaa, 0xef, 0x5a, 0x45, 0xb9,
	0x80, 0xf4, 0x69, 0xa8, 0xeb, 0xdf, 0xd6, 0x2f, 0x21, 0x3e, 0x0c, 0x52, 0x75, 0xd2, 0x1a, 0x84,
	0xcc, 0x2b, 0xbc, 0x84, 0x49, 0x2d, 0x1a, 0xa1, 0xad, 0xc3, 0x84, 0x39, 0x81, 0x19, 0x4c, 0x25,
	0x9d, 0x48, 0x2a, 0x77, 0x4a, 0xc2, 0xbe, 0x65, 0xae, 0x61, 0xee, 0xac, 0x9c, 0x35, 0x16, 0x90,
	0x34, 0xae, 0x41, 0x95, 0x05, 0xab, 0xb0, 0x48, 0xb7, 0xf3, 0x0d, 0xef, 0xc5, 0xc6, 0xd7, 0xca,
	0x7e, 0xbe, 0xe2, 0x15, 0x24, 0x15, 0x57, 0x65, 0xd3, 0x49, 0xb2, 0x19, 0x12, 0x36, 0xad, 0xb8,
	0xda, 0x75, 0x92, 0xf0, 0x16, 0xd2, 0x96, 0xce, 0xba, 0xf4, 0x09, 0x43, 0x9b, 0x10, 0x0c, 0x7a,
	0xb4, 0x64, 0xfb, 0x01, 0x0b, 0xbf, 0x70, 0x4f, 0xf2, 0x24, 0x0e, 0x84, 0x6b, 0x88, 0x4c, 0x05,
	0xf8, 0xdf, 0xba, 0x8d, 0xca, 0xbd, 0xbe, 0x18, 0x11, 0x1f, 0x72, 0x0d, 0x91, 0x09, 0xed, 0x1f,
	0x8f, 0xaa, 0xf2, 0x8f, 0xc7, 0x17, 0x3d, 0xc4, 0x2f, 0xd1, 0x1d, 0xef, 0xc5, 0x6b, 0x6c, 0xff,
	0x93, 0xfb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x99, 0x2a, 0x60, 0x39, 0x02, 0x00, 0x00,
}
