// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register/InstancePing.proto

package register

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	common "pb6/common"
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

type ServiceInstancePingPkg struct {
	ServiceInstanceId    int32    `protobuf:"varint,1,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	ServiceInstanceUUID  string   `protobuf:"bytes,3,opt,name=serviceInstanceUUID,proto3" json:"serviceInstanceUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInstancePingPkg) Reset()         { *m = ServiceInstancePingPkg{} }
func (m *ServiceInstancePingPkg) String() string { return proto.CompactTextString(m) }
func (*ServiceInstancePingPkg) ProtoMessage()    {}
func (*ServiceInstancePingPkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c10bfdff83d834ed, []int{0}
}

func (m *ServiceInstancePingPkg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstancePingPkg.Unmarshal(m, b)
}
func (m *ServiceInstancePingPkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstancePingPkg.Marshal(b, m, deterministic)
}
func (m *ServiceInstancePingPkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstancePingPkg.Merge(m, src)
}
func (m *ServiceInstancePingPkg) XXX_Size() int {
	return xxx_messageInfo_ServiceInstancePingPkg.Size(m)
}
func (m *ServiceInstancePingPkg) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstancePingPkg.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstancePingPkg proto.InternalMessageInfo

func (m *ServiceInstancePingPkg) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func (m *ServiceInstancePingPkg) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ServiceInstancePingPkg) GetServiceInstanceUUID() string {
	if m != nil {
		return m.ServiceInstanceUUID
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceInstancePingPkg)(nil), "ServiceInstancePingPkg")
}

func init() { proto.RegisterFile("register/InstancePing.proto", fileDescriptor_c10bfdff83d834ed) }

var fileDescriptor_c10bfdff83d834ed = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0xf7, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x0d, 0xc8, 0xcc,
	0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x12, 0x4e, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3,
	0x87, 0x50, 0x10, 0x41, 0xa5, 0x1e, 0x46, 0x2e, 0xb1, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x64, 0x2d, 0x01, 0xd9, 0xe9, 0x42, 0x3a, 0x5c, 0x82, 0xc5, 0xa8, 0x32, 0x9e, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x98, 0x12, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0x99, 0xb9, 0xa9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x90, 0x01, 0x97, 0x30, 0x9a, 0xc2, 0xd0, 0x50,
	0x4f, 0x17, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x6c, 0x52, 0x46, 0xae, 0x5c, 0xc2, 0x58,
	0x5c, 0x23, 0xa4, 0xc7, 0xc5, 0x96, 0x92, 0x0f, 0x66, 0x89, 0xeb, 0x61, 0x77, 0xad, 0x14, 0xa7,
	0x9e, 0x73, 0x7e, 0x6e, 0x6e, 0x62, 0x5e, 0x4a, 0xb1, 0x12, 0x83, 0x53, 0x0c, 0x97, 0x6e, 0x7e,
	0x51, 0xba, 0x5e, 0x62, 0x41, 0x62, 0x72, 0x46, 0xaa, 0x5e, 0x71, 0x76, 0x65, 0x79, 0x62, 0x4e,
	0x36, 0x28, 0x24, 0x12, 0x0b, 0x72, 0xf5, 0xf2, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x60,
	0x61, 0xa5, 0x57, 0x66, 0x14, 0xc0, 0x18, 0xc5, 0x53, 0x90, 0x64, 0xa6, 0x0f, 0x13, 0x5a, 0xc5,
	0x24, 0x15, 0x9c, 0x5d, 0x19, 0x0e, 0xd5, 0xe4, 0x07, 0xd1, 0x10, 0x00, 0x0a, 0xb0, 0xe4, 0xfc,
	0x9c, 0x24, 0x36, 0x70, 0xd0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x17, 0x08, 0x4b,
	0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceInstancePingClient is the client API for ServiceInstancePing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceInstancePingClient interface {
	DoPing(ctx context.Context, in *ServiceInstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error)
}

type serviceInstancePingClient struct {
	cc *grpc.ClientConn
}

func NewServiceInstancePingClient(cc *grpc.ClientConn) ServiceInstancePingClient {
	return &serviceInstancePingClient{cc}
}

func (c *serviceInstancePingClient) DoPing(ctx context.Context, in *ServiceInstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/ServiceInstancePing/doPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceInstancePingServer is the server API for ServiceInstancePing service.
type ServiceInstancePingServer interface {
	DoPing(context.Context, *ServiceInstancePingPkg) (*common.Commands, error)
}

// UnimplementedServiceInstancePingServer can be embedded to have forward compatible implementations.
type UnimplementedServiceInstancePingServer struct {
}

func (*UnimplementedServiceInstancePingServer) DoPing(ctx context.Context, req *ServiceInstancePingPkg) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoPing not implemented")
}

func RegisterServiceInstancePingServer(s *grpc.Server, srv ServiceInstancePingServer) {
	s.RegisterService(&_ServiceInstancePing_serviceDesc, srv)
}

func _ServiceInstancePing_DoPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInstancePingPkg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInstancePingServer).DoPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceInstancePing/DoPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInstancePingServer).DoPing(ctx, req.(*ServiceInstancePingPkg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceInstancePing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceInstancePing",
	HandlerType: (*ServiceInstancePingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doPing",
			Handler:    _ServiceInstancePing_DoPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register/InstancePing.proto",
}