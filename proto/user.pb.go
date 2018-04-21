// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package letstalk is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	Talk
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
	TalkRequest
	TalkReply
	GetUsers
	GetTalks
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

type User struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password1 string `protobuf:"bytes,2,opt,name=password1" json:"password1,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname" json:"lastname,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword1() string {
	if m != nil {
		return m.Password1
	}
	return ""
}

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type Talk struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Talk  string `protobuf:"bytes,2,opt,name=talk" json:"talk,omitempty"`
	Date  string `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
}

func (m *Talk) Reset()                    { *m = Talk{} }
func (m *Talk) String() string            { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()               {}
func (*Talk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Talk) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Talk) GetTalk() string {
	if m != nil {
		return m.Talk
	}
	return ""
}

func (m *Talk) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type SignupRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *SignupRequest) Reset()                    { *m = SignupRequest{} }
func (m *SignupRequest) String() string            { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()               {}
func (*SignupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SignupRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SignupReply struct {
	Message   string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Sessionid string `protobuf:"bytes,2,opt,name=sessionid" json:"sessionid,omitempty"`
}

func (m *SignupReply) Reset()                    { *m = SignupReply{} }
func (m *SignupReply) String() string            { return proto.CompactTextString(m) }
func (*SignupReply) ProtoMessage()               {}
func (*SignupReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SignupReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SignupReply) GetSessionid() string {
	if m != nil {
		return m.Sessionid
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
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	Message   string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoginReply) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type LogoutRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LogoutRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type LogoutReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *LogoutReply) Reset()                    { *m = LogoutReply{} }
func (m *LogoutReply) String() string            { return proto.CompactTextString(m) }
func (*LogoutReply) ProtoMessage()               {}
func (*LogoutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*CancelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
func (*CancelReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
func (*FollowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FollowRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type FollowReply struct {
	Message  string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Userlist []string `protobuf:"bytes,2,rep,name=userlist" json:"userlist,omitempty"`
}

func (m *FollowReply) Reset()                    { *m = FollowReply{} }
func (m *FollowReply) String() string            { return proto.CompactTextString(m) }
func (*FollowReply) ProtoMessage()               {}
func (*FollowReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FollowReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FollowReply) GetUserlist() []string {
	if m != nil {
		return m.Userlist
	}
	return nil
}

type TalkRequest struct {
	Talk    *Talk  `protobuf:"bytes,1,opt,name=talk" json:"talk,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *TalkRequest) Reset()                    { *m = TalkRequest{} }
func (m *TalkRequest) String() string            { return proto.CompactTextString(m) }
func (*TalkRequest) ProtoMessage()               {}
func (*TalkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TalkRequest) GetTalk() *Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

func (m *TalkRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TalkReply struct {
	Message string  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Talk    []*Talk `protobuf:"bytes,2,rep,name=talk" json:"talk,omitempty"`
}

func (m *TalkReply) Reset()                    { *m = TalkReply{} }
func (m *TalkReply) String() string            { return proto.CompactTextString(m) }
func (*TalkReply) ProtoMessage()               {}
func (*TalkReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TalkReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TalkReply) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type GetUsers struct {
	User []*User `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
}

func (m *GetUsers) Reset()                    { *m = GetUsers{} }
func (m *GetUsers) String() string            { return proto.CompactTextString(m) }
func (*GetUsers) ProtoMessage()               {}
func (*GetUsers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetUsers) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetTalks struct {
	Talk []*Talk `protobuf:"bytes,1,rep,name=talk" json:"talk,omitempty"`
}

func (m *GetTalks) Reset()                    { *m = GetTalks{} }
func (m *GetTalks) String() string            { return proto.CompactTextString(m) }
func (*GetTalks) ProtoMessage()               {}
func (*GetTalks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetTalks) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "letstalk.User")
	proto.RegisterType((*Talk)(nil), "letstalk.Talk")
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
	proto.RegisterType((*TalkRequest)(nil), "letstalk.TalkRequest")
	proto.RegisterType((*TalkReply)(nil), "letstalk.TalkReply")
	proto.RegisterType((*GetUsers)(nil), "letstalk.GetUsers")
	proto.RegisterType((*GetTalks)(nil), "letstalk.GetTalks")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Letstalk service

type LetstalkClient interface {
	// Sends a signup request
	SendSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupReply, error)
	// Sends a login request
	SendLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// Sends a logout request
	SendLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// Sends a cancel request
	SendCancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelReply, error)
	// Sends a follow request
	SendFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowReply, error)
	// Sends a talk request
	SendTalk(ctx context.Context, in *TalkRequest, opts ...grpc.CallOption) (*TalkReply, error)
}

type letstalkClient struct {
	cc *grpc.ClientConn
}

func NewLetstalkClient(cc *grpc.ClientConn) LetstalkClient {
	return &letstalkClient{cc}
}

func (c *letstalkClient) SendSignup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupReply, error) {
	out := new(SignupReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendSignup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letstalkClient) SendLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letstalkClient) SendLogout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendLogout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letstalkClient) SendCancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelReply, error) {
	out := new(CancelReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendCancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letstalkClient) SendFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendFollow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letstalkClient) SendTalk(ctx context.Context, in *TalkRequest, opts ...grpc.CallOption) (*TalkReply, error) {
	out := new(TalkReply)
	err := grpc.Invoke(ctx, "/letstalk.Letstalk/SendTalk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Letstalk service

type LetstalkServer interface {
	// Sends a signup request
	SendSignup(context.Context, *SignupRequest) (*SignupReply, error)
	// Sends a login request
	SendLogin(context.Context, *LoginRequest) (*LoginReply, error)
	// Sends a logout request
	SendLogout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// Sends a cancel request
	SendCancel(context.Context, *CancelRequest) (*CancelReply, error)
	// Sends a follow request
	SendFollow(context.Context, *FollowRequest) (*FollowReply, error)
	// Sends a talk request
	SendTalk(context.Context, *TalkRequest) (*TalkReply, error)
}

func RegisterLetstalkServer(s *grpc.Server, srv LetstalkServer) {
	s.RegisterService(&_Letstalk_serviceDesc, srv)
}

func _Letstalk_SendSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendSignup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Letstalk_SendLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Letstalk_SendLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendLogout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Letstalk_SendCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendCancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Letstalk_SendFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendFollow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Letstalk_SendTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetstalkServer).SendTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letstalk.Letstalk/SendTalk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetstalkServer).SendTalk(ctx, req.(*TalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Letstalk_serviceDesc = grpc.ServiceDesc{
	ServiceName: "letstalk.Letstalk",
	HandlerType: (*LetstalkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSignup",
			Handler:    _Letstalk_SendSignup_Handler,
		},
		{
			MethodName: "SendLogin",
			Handler:    _Letstalk_SendLogin_Handler,
		},
		{
			MethodName: "SendLogout",
			Handler:    _Letstalk_SendLogout_Handler,
		},
		{
			MethodName: "SendCancel",
			Handler:    _Letstalk_SendCancel_Handler,
		},
		{
			MethodName: "SendFollow",
			Handler:    _Letstalk_SendFollow_Handler,
		},
		{
			MethodName: "SendTalk",
			Handler:    _Letstalk_SendTalk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x77, 0x9b, 0x00, 0xc9, 0x84, 0xe5, 0x60, 0x76, 0x21, 0x8a, 0x38, 0xac, 0x2c, 0x21,
	0x38, 0x45, 0x62, 0xf7, 0x82, 0x84, 0xc4, 0x81, 0x16, 0x50, 0x45, 0x4f, 0x2d, 0x3c, 0x80, 0x21,
	0x26, 0x8a, 0xea, 0xc6, 0x21, 0x76, 0x54, 0xf5, 0xc8, 0x9b, 0x23, 0xdb, 0x71, 0xe2, 0xb4, 0x4a,
	0x8a, 0xf6, 0x16, 0xcf, 0xcc, 0x3f, 0xff, 0x8c, 0xfb, 0xd5, 0x00, 0x8d, 0xa0, 0x75, 0x5a, 0xd5,
	0x5c, 0x72, 0x14, 0x30, 0x2a, 0x85, 0x24, 0x6c, 0x8b, 0x25, 0xf8, 0x3f, 0x04, 0xad, 0xd1, 0x35,
	0x3c, 0xa2, 0x3b, 0x52, 0xb0, 0xf8, 0xf2, 0xf6, 0xf2, 0x6d, 0xb8, 0x36, 0x07, 0xf4, 0x0a, 0xc2,
	0x8a, 0x08, 0xb1, 0xe7, 0x75, 0xf6, 0x2e, 0x9e, 0xe9, 0x4c, 0x1f, 0x50, 0xd9, 0xdf, 0x45, 0x2d,
	0x64, 0x49, 0x76, 0x34, 0xf6, 0x4c, 0xb6, 0x0b, 0xa0, 0x04, 0x02, 0x46, 0xda, 0xa4, 0xaf, 0x93,
	0xdd, 0x19, 0x2f, 0xc0, 0xff, 0x4e, 0xd8, 0x76, 0xc4, 0x15, 0x81, 0xaf, 0x66, 0x6b, 0x0d, 0xf5,
	0xb7, 0x8a, 0x65, 0x44, 0x5a, 0x1b, 0xfd, 0x8d, 0xef, 0xe1, 0x6a, 0x53, 0xe4, 0x65, 0x53, 0xad,
	0xe9, 0x9f, 0x86, 0x0a, 0x89, 0x30, 0xf8, 0x6a, 0x49, 0xdd, 0x2d, 0xba, 0x7b, 0x96, 0xda, 0x2d,
	0x53, 0xb5, 0xe2, 0x5a, 0xe7, 0xf0, 0x67, 0x88, 0xac, 0xa8, 0x62, 0x07, 0x14, 0xc3, 0x93, 0x1d,
	0x15, 0x82, 0xe4, 0xb4, 0x9d, 0xc1, 0x1e, 0xd5, 0x76, 0x82, 0x0a, 0x51, 0xf0, 0xb2, 0xc8, 0xec,
	0xee, 0x5d, 0x00, 0x7f, 0x82, 0xa7, 0x2b, 0x9e, 0x17, 0xa5, 0xb5, 0x7e, 0xc0, 0xfd, 0xe1, 0x05,
	0x40, 0xdb, 0xe3, 0x7f, 0x27, 0x59, 0x1e, 0x4f, 0xb2, 0xcc, 0xf0, 0x6b, 0xb8, 0x5a, 0xf1, 0x9c,
	0x37, 0x72, 0x72, 0x14, 0xfc, 0x06, 0x22, 0x5b, 0x36, 0xe9, 0xa6, 0xfa, 0xcd, 0x49, 0xf9, 0x8b,
	0xb2, 0xb3, 0xfd, 0x6c, 0xd9, 0xd9, 0x7e, 0x5f, 0x38, 0x63, 0x7c, 0x3f, 0xdd, 0x6f, 0x0e, 0x91,
	0x2d, 0x9b, 0xbe, 0x8d, 0x04, 0x02, 0xf5, 0x43, 0xb2, 0x42, 0xc8, 0x78, 0x76, 0xeb, 0x29, 0xae,
	0xec, 0x19, 0x7f, 0x83, 0x48, 0x71, 0xe5, 0xf0, 0xa0, 0x41, 0x3a, 0xe1, 0x41, 0x17, 0x19, 0xb0,
	0x1c, 0xa3, 0xd9, 0x70, 0xf0, 0x25, 0x84, 0xa6, 0xd9, 0xf4, 0x3c, 0xb8, 0xa3, 0xd5, 0x1b, 0x33,
	0xc1, 0x29, 0x04, 0x5f, 0xa9, 0x54, 0x14, 0x0a, 0x07, 0x52, 0x6f, 0x14, 0x52, 0x53, 0xaf, 0x1a,
	0x08, 0x67, 0x89, 0xd1, 0xfe, 0x77, 0x7f, 0x3d, 0x08, 0x56, 0x6d, 0x1c, 0x7d, 0x04, 0xd8, 0xd0,
	0x32, 0x33, 0x94, 0xa3, 0x97, 0xbd, 0x60, 0xf0, 0x67, 0x49, 0x6e, 0x4e, 0x13, 0x15, 0x3b, 0xe0,
	0x0b, 0xf4, 0x01, 0x42, 0xa5, 0xd7, 0x68, 0xa2, 0x17, 0x7d, 0x95, 0xcb, 0x7b, 0x72, 0x7d, 0x12,
	0x37, 0xe2, 0xd6, 0xdc, 0xa0, 0xe6, 0x9a, 0x0f, 0x18, 0x75, 0xcd, 0x1d, 0x2a, 0x7b, 0xbd, 0x41,
	0xcb, 0xd5, 0x0f, 0x98, 0x74, 0xf5, 0x0e, 0x85, 0xbd, 0xde, 0xa0, 0xe4, 0xea, 0x07, 0x0c, 0xba,
	0x7a, 0x87, 0x3a, 0x7c, 0x81, 0xde, 0x43, 0xa0, 0xf4, 0xfa, 0x75, 0xba, 0x39, 0xba, 0xeb, 0x56,
	0xfb, 0xfc, 0x38, 0xac, 0x95, 0x3f, 0x1f, 0xeb, 0xa7, 0xf5, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x57, 0x1c, 0x67, 0x68, 0x05, 0x00, 0x00,
}
