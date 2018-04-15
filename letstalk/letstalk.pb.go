// Code generated by protoc-gen-go. DO NOT EDIT.
// source: letstalk.proto

/*
Package letstalk is a generated protocol buffer package.

It is generated from these files:
	letstalk.proto

It has these top-level messages:
	SignupRequest
	SignupReply
	LoginRequest
	LoginReply
	LogoutRequest
	LogoutReply
	CancelRequest
	CancelReply
	FollowRequest
	FollowReply
*/
package letstalk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignupRequest struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password1 string `protobuf:"bytes,2,opt,name=password1" json:"password1,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname" json:"lastname,omitempty"`
}

func (m *SignupRequest) Reset()                    { *m = SignupRequest{} }
func (m *SignupRequest) String() string            { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()               {}
func (*SignupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignupRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupRequest) GetPassword1() string {
	if m != nil {
		return m.Password1
	}
	return ""
}

func (m *SignupRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *SignupRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type SignupReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SignupReply) Reset()                    { *m = SignupReply{} }
func (m *SignupReply) String() string            { return proto.CompactTextString(m) }
func (*SignupReply) ProtoMessage()               {}
func (*SignupReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignupReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LoginRequest struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password1 string `protobuf:"bytes,2,opt,name=password1" json:"password1,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword1() string {
	if m != nil {
		return m.Password1
	}
	return ""
}

type LoginReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogoutRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogoutRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogoutReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *LogoutReply) Reset()                    { *m = LogoutReply{} }
func (m *LogoutReply) String() string            { return proto.CompactTextString(m) }
func (*LogoutReply) ProtoMessage()               {}
func (*LogoutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogoutReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CancelRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *CancelRequest) Reset()                    { *m = CancelRequest{} }
func (m *CancelRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()               {}
func (*CancelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CancelRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CancelReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *CancelReply) Reset()                    { *m = CancelReply{} }
func (m *CancelReply) String() string            { return proto.CompactTextString(m) }
func (*CancelReply) ProtoMessage()               {}
func (*CancelReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CancelReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FollowRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *FollowRequest) Reset()                    { *m = FollowRequest{} }
func (m *FollowRequest) String() string            { return proto.CompactTextString(m) }
func (*FollowRequest) ProtoMessage()               {}
func (*FollowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FollowRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type FollowReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *FollowReply) Reset()                    { *m = FollowReply{} }
func (m *FollowReply) String() string            { return proto.CompactTextString(m) }
func (*FollowReply) ProtoMessage()               {}
func (*FollowReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FollowReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SignupRequest)(nil), "letstalk.SignupRequest")
	proto.RegisterType((*SignupReply)(nil), "letstalk.SignupReply")
	proto.RegisterType((*LoginRequest)(nil), "letstalk.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "letstalk.LoginReply")
	proto.RegisterType((*LogoutRequest)(nil), "letstalk.LogoutRequest")
	proto.RegisterType((*LogoutReply)(nil), "letstalk.LogoutReply")
	proto.RegisterType((*CancelRequest)(nil), "letstalk.CancelRequest")
	proto.RegisterType((*CancelReply)(nil), "letstalk.CancelReply")
	proto.RegisterType((*FollowRequest)(nil), "letstalk.FollowRequest")
	proto.RegisterType((*FollowReply)(nil), "letstalk.FollowReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Signup service

type SignupClient interface {
	// Sends a signup request
	SendSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupReply, error)
}

type signupClient struct {
	cc *grpc.ClientConn
}

func NewSignupClient(cc *grpc.ClientConn) SignupClient {
	return &signupClient{cc}
}

func (c *signupClient) SendSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupReply, error) {
	out := new(SignupReply)
	err := grpc.Invoke(ctx, "/letstalk.Signup/SendSignup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Signup service

type SignupServer interface {
	// Sends a signup request
	SendSignup(context.Context, *SignupRequest) (*SignupReply, error)
}

func RegisterSignupServer(s *grpc.Server, srv SignupServer) {
	s.RegisterService(&_Signup_serviceDesc, srv)
}

func _Signup_SendSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).SendSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Signup/SendSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).SendSignup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "letstalk.Signup",
	HandlerType: (*SignupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSignup",
			Handler:    _Signup_SendSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "letstalk.proto",
}

// Client API for Login service

type LoginClient interface {
	// Sends a login request
	SendLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	SendLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) SendLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/letstalk.Login/SendLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) SendLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := grpc.Invoke(ctx, "/letstalk.Login/SendLogout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	// Sends a login request
	SendLogin(context.Context, *LoginRequest) (*LoginReply, error)
	SendLogout(context.Context, *LogoutRequest) (*LogoutReply, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_SendLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).SendLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Login/SendLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).SendLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_SendLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).SendLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Login/SendLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).SendLogout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "letstalk.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogin",
			Handler:    _Login_SendLogin_Handler,
		},
		{
			MethodName: "SendLogout",
			Handler:    _Login_SendLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "letstalk.proto",
}

// Client API for Cancel service

type CancelClient interface {
	// Sends a signup request
	SendCancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelReply, error)
}

type cancelClient struct {
	cc *grpc.ClientConn
}

func NewCancelClient(cc *grpc.ClientConn) CancelClient {
	return &cancelClient{cc}
}

func (c *cancelClient) SendCancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelReply, error) {
	out := new(CancelReply)
	err := grpc.Invoke(ctx, "/letstalk.Cancel/SendCancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cancel service

type CancelServer interface {
	// Sends a signup request
	SendCancel(context.Context, *CancelRequest) (*CancelReply, error)
}

func RegisterCancelServer(s *grpc.Server, srv CancelServer) {
	s.RegisterService(&_Cancel_serviceDesc, srv)
}

func _Cancel_SendCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CancelServer).SendCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Cancel/SendCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CancelServer).SendCancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cancel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "letstalk.Cancel",
	HandlerType: (*CancelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCancel",
			Handler:    _Cancel_SendCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "letstalk.proto",
}

// Client API for Follow service

type FollowClient interface {
	// Sends a follow request
	SendFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowReply, error)
}

type followClient struct {
	cc *grpc.ClientConn
}

func NewFollowClient(cc *grpc.ClientConn) FollowClient {
	return &followClient{cc}
}

func (c *followClient) SendFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := grpc.Invoke(ctx, "/letstalk.Follow/SendFollow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Follow service

type FollowServer interface {
	// Sends a follow request
	SendFollow(context.Context, *FollowRequest) (*FollowReply, error)
}

func RegisterFollowServer(s *grpc.Server, srv FollowServer) {
	s.RegisterService(&_Follow_serviceDesc, srv)
}

func _Follow_SendFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).SendFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Follow/SendFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).SendFollow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Follow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "letstalk.Follow",
	HandlerType: (*FollowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFollow",
			Handler:    _Follow_SendFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "letstalk.proto",
}

func init() { proto.RegisterFile("letstalk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x13, 0xa0, 0xa5, 0xb9, 0x12, 0x06, 0xab, 0x40, 0x14, 0x31, 0x20, 0x4b, 0x50, 0x58,
	0x2a, 0x51, 0x46, 0x24, 0x06, 0x90, 0x10, 0x43, 0xa7, 0xf6, 0x13, 0x18, 0x6a, 0xa2, 0x08, 0x27,
	0x0e, 0xb1, 0xa3, 0xaa, 0x03, 0x1b, 0x1f, 0x1c, 0xc5, 0x7f, 0x52, 0x47, 0x15, 0xe9, 0xd0, 0xf1,
	0xdd, 0x3b, 0xfd, 0xee, 0x72, 0xcf, 0x81, 0x53, 0x46, 0xa5, 0x90, 0x84, 0x7d, 0x4d, 0x8a, 0x92,
	0x4b, 0x8e, 0x06, 0x56, 0xe3, 0x1f, 0x08, 0x17, 0x69, 0x92, 0x57, 0xc5, 0x9c, 0x7e, 0x57, 0x54,
	0x48, 0x34, 0x82, 0x1e, 0xcd, 0x48, 0xca, 0x22, 0xff, 0xca, 0xbf, 0x0d, 0xe6, 0x5a, 0xa0, 0x4b,
	0x08, 0x0a, 0x22, 0xc4, 0x8a, 0x97, 0xcb, 0xfb, 0xe8, 0x40, 0x39, 0x9b, 0x42, 0xed, 0x7e, 0xa6,
	0xa5, 0x90, 0x39, 0xc9, 0x68, 0x74, 0xa8, 0xdd, 0xa6, 0x80, 0x62, 0x18, 0x30, 0x62, 0xcc, 0x23,
	0x65, 0x36, 0x1a, 0x8f, 0x61, 0x68, 0xc7, 0x17, 0x6c, 0x8d, 0x22, 0x38, 0xce, 0xa8, 0x10, 0x24,
	0xa1, 0x66, 0xbc, 0x95, 0xf8, 0x19, 0x4e, 0x66, 0x3c, 0x49, 0xf3, 0x3d, 0xd6, 0xc4, 0x37, 0x00,
	0x86, 0xd1, 0x3d, 0xeb, 0x0e, 0xc2, 0x19, 0x4f, 0x78, 0x25, 0xed, 0xb0, 0xff, 0x5b, 0xc7, 0x30,
	0xb4, 0xad, 0xdd, 0xcc, 0x6b, 0x08, 0x5f, 0x48, 0xfe, 0x41, 0x59, 0xe7, 0x07, 0xd4, 0x3c, 0xdb,
	0xb6, 0x93, 0xf7, 0xca, 0x19, 0xe3, 0xab, 0x9d, 0x3c, 0xdb, 0xd6, 0xc9, 0x9b, 0xbe, 0x41, 0x5f,
	0x07, 0x81, 0x9e, 0x00, 0x16, 0x34, 0x5f, 0x1a, 0x75, 0x31, 0x69, 0x9e, 0x4e, 0xeb, 0x9d, 0xc4,
	0x67, 0xdb, 0x46, 0xc1, 0xd6, 0xd8, 0x9b, 0xfe, 0xfa, 0xd0, 0x53, 0x67, 0x46, 0x8f, 0x10, 0xd4,
	0x24, 0x2d, 0xce, 0x37, 0xfd, 0x6e, 0x90, 0xf1, 0x68, 0xab, 0xae, 0x30, 0x76, 0x0d, 0x7d, 0x5d,
	0x77, 0x8d, 0x56, 0x34, 0xee, 0x1a, 0x4e, 0x10, 0xd8, 0xab, 0x3f, 0x48, 0x5f, 0xd2, 0x92, 0x8c,
	0x72, 0x48, 0xad, 0x40, 0x5c, 0x92, 0x13, 0x81, 0x26, 0xe9, 0x1b, 0x5a, 0x92, 0x51, 0x0e, 0xa9,
	0x15, 0x85, 0x4b, 0x72, 0x8e, 0x8f, 0xbd, 0xf7, 0xbe, 0xfa, 0xfb, 0x1e, 0xfe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0xa9, 0x0e, 0x24, 0x8f, 0x03, 0x00, 0x00,
}