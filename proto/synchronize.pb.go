// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: proto/synchronize.proto

package synchronize

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_synchronize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_synchronize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_synchronize_proto_rawDescGZIP(), []int{0}
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_synchronize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_synchronize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_proto_synchronize_proto_rawDescGZIP(), []int{1}
}

type MsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Topic   string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Group   string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *MsgRequest) Reset() {
	*x = MsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_synchronize_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRequest) ProtoMessage() {}

func (x *MsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_synchronize_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRequest.ProtoReflect.Descriptor instead.
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return file_proto_synchronize_proto_rawDescGZIP(), []int{2}
}

func (x *MsgRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MsgRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *MsgRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_synchronize_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_synchronize_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_proto_synchronize_proto_rawDescGZIP(), []int{3}
}

func (x *CommonReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CommonReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_synchronize_proto protoreflect.FileDescriptor

var file_proto_synchronize_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x3e, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e, 0x63,
	0x12, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_synchronize_proto_rawDescOnce sync.Once
	file_proto_synchronize_proto_rawDescData = file_proto_synchronize_proto_rawDesc
)

func file_proto_synchronize_proto_rawDescGZIP() []byte {
	file_proto_synchronize_proto_rawDescOnce.Do(func() {
		file_proto_synchronize_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_synchronize_proto_rawDescData)
	})
	return file_proto_synchronize_proto_rawDescData
}

var file_proto_synchronize_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_synchronize_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: synchronize.PingRequest
	(*PongResponse)(nil), // 1: synchronize.PongResponse
	(*MsgRequest)(nil),   // 2: synchronize.MsgRequest
	(*CommonReply)(nil),  // 3: synchronize.CommonReply
}
var file_proto_synchronize_proto_depIdxs = []int32{
	2, // 0: synchronize.ServerSync.MsgSync:input_type -> synchronize.MsgRequest
	0, // 1: synchronize.ServerSync.Ping:input_type -> synchronize.PingRequest
	3, // 2: synchronize.ServerSync.MsgSync:output_type -> synchronize.CommonReply
	1, // 3: synchronize.ServerSync.Ping:output_type -> synchronize.PongResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_synchronize_proto_init() }
func file_proto_synchronize_proto_init() {
	if File_proto_synchronize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_synchronize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_synchronize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_synchronize_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_synchronize_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_synchronize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_synchronize_proto_goTypes,
		DependencyIndexes: file_proto_synchronize_proto_depIdxs,
		MessageInfos:      file_proto_synchronize_proto_msgTypes,
	}.Build()
	File_proto_synchronize_proto = out.File
	file_proto_synchronize_proto_rawDesc = nil
	file_proto_synchronize_proto_goTypes = nil
	file_proto_synchronize_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerSyncClient is the client API for ServerSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerSyncClient interface {
	MsgSync(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*CommonReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type serverSyncClient struct {
	cc grpc.ClientConnInterface
}

func NewServerSyncClient(cc grpc.ClientConnInterface) ServerSyncClient {
	return &serverSyncClient{cc}
}

func (c *serverSyncClient) MsgSync(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/synchronize.ServerSync/MsgSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverSyncClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/synchronize.ServerSync/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerSyncServer is the server API for ServerSync service.
type ServerSyncServer interface {
	MsgSync(context.Context, *MsgRequest) (*CommonReply, error)
	Ping(context.Context, *PingRequest) (*PongResponse, error)
}

// UnimplementedServerSyncServer can be embedded to have forward compatible implementations.
type UnimplementedServerSyncServer struct {
}

func (*UnimplementedServerSyncServer) MsgSync(context.Context, *MsgRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgSync not implemented")
}
func (*UnimplementedServerSyncServer) Ping(context.Context, *PingRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterServerSyncServer(s *grpc.Server, srv ServerSyncServer) {
	s.RegisterService(&_ServerSync_serviceDesc, srv)
}

func _ServerSync_MsgSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSyncServer).MsgSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronize.ServerSync/MsgSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSyncServer).MsgSync(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerSync_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSyncServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronize.ServerSync/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSyncServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synchronize.ServerSync",
	HandlerType: (*ServerSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgSync",
			Handler:    _ServerSync_MsgSync_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ServerSync_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/synchronize.proto",
}