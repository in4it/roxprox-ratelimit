// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

package cache

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GetCacheRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCacheRequest) Reset()         { *m = GetCacheRequest{} }
func (m *GetCacheRequest) String() string { return proto.CompactTextString(m) }
func (*GetCacheRequest) ProtoMessage()    {}
func (*GetCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}

func (m *GetCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCacheRequest.Unmarshal(m, b)
}
func (m *GetCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCacheRequest.Marshal(b, m, deterministic)
}
func (m *GetCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheRequest.Merge(m, src)
}
func (m *GetCacheRequest) XXX_Size() int {
	return xxx_messageInfo_GetCacheRequest.Size(m)
}
func (m *GetCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheRequest proto.InternalMessageInfo

type GetCacheReply struct {
	CacheItems           []*GetCacheReply_CacheItem `protobuf:"bytes,1,rep,name=cacheItems,proto3" json:"cacheItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetCacheReply) Reset()         { *m = GetCacheReply{} }
func (m *GetCacheReply) String() string { return proto.CompactTextString(m) }
func (*GetCacheReply) ProtoMessage()    {}
func (*GetCacheReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

func (m *GetCacheReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCacheReply.Unmarshal(m, b)
}
func (m *GetCacheReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCacheReply.Marshal(b, m, deterministic)
}
func (m *GetCacheReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheReply.Merge(m, src)
}
func (m *GetCacheReply) XXX_Size() int {
	return xxx_messageInfo_GetCacheReply.Size(m)
}
func (m *GetCacheReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheReply proto.InternalMessageInfo

func (m *GetCacheReply) GetCacheItems() []*GetCacheReply_CacheItem {
	if m != nil {
		return m.CacheItems
	}
	return nil
}

type GetCacheReply_CacheItem struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCacheReply_CacheItem) Reset()         { *m = GetCacheReply_CacheItem{} }
func (m *GetCacheReply_CacheItem) String() string { return proto.CompactTextString(m) }
func (*GetCacheReply_CacheItem) ProtoMessage()    {}
func (*GetCacheReply_CacheItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1, 0}
}

func (m *GetCacheReply_CacheItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCacheReply_CacheItem.Unmarshal(m, b)
}
func (m *GetCacheReply_CacheItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCacheReply_CacheItem.Marshal(b, m, deterministic)
}
func (m *GetCacheReply_CacheItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheReply_CacheItem.Merge(m, src)
}
func (m *GetCacheReply_CacheItem) XXX_Size() int {
	return xxx_messageInfo_GetCacheReply_CacheItem.Size(m)
}
func (m *GetCacheReply_CacheItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheReply_CacheItem.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheReply_CacheItem proto.InternalMessageInfo

func (m *GetCacheReply_CacheItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetCacheReply_CacheItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCacheRequest)(nil), "GetCacheRequest")
	proto.RegisterType((*GetCacheReply)(nil), "GetCacheReply")
	proto.RegisterType((*GetCacheReply_CacheItem)(nil), "GetCacheReply.CacheItem")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe4, 0xe2, 0x77, 0x4f, 0x2d, 0x71, 0x06,
	0x89, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xd5, 0x71, 0xf1, 0x22, 0x84, 0x0a, 0x72,
	0x2a, 0x85, 0x2c, 0xb8, 0xb8, 0xc0, 0x5a, 0x3c, 0x4b, 0x52, 0x73, 0x8b, 0x25, 0x18, 0x15, 0x98,
	0x35, 0xb8, 0x8d, 0x24, 0xf4, 0x50, 0xd4, 0xe8, 0x39, 0xc3, 0x14, 0x04, 0x21, 0xa9, 0x95, 0x32,
	0xe6, 0xe2, 0x84, 0x4b, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60,
	0x31, 0x08, 0xc7, 0xc8, 0x9c, 0x8b, 0x15, 0xac, 0x49, 0x48, 0x8f, 0x8b, 0x03, 0x66, 0x89, 0x90,
	0x80, 0x1e, 0x9a, 0x33, 0xa5, 0xf8, 0x50, 0x5d, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0x92, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x82, 0x4e, 0xbb, 0xd5, 0xe1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheReply, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheReply, error) {
	out := new(GetCacheReply)
	err := c.cc.Invoke(ctx, "/Cache/GetCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	GetCache(context.Context, *GetCacheRequest) (*GetCacheReply, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/GetCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetCache(ctx, req.(*GetCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCache",
			Handler:    _Cache_GetCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}
