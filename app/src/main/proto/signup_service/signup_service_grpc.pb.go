// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: signup_service/signup_service.proto

package signup_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SignupServiceClient is the client API for SignupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignupServiceClient interface {
	SetPhone(ctx context.Context, in *SetPhoneRequest, opts ...grpc.CallOption) (*SetPhoneResponse, error)
	SetCode(ctx context.Context, in *SetCodeRequest, opts ...grpc.CallOption) (*SetCodeResponse, error)
	GetPartnerCode(ctx context.Context, in *GetPartnerCodeRequest, opts ...grpc.CallOption) (*GetPartnerCodeResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	AuthCheck(ctx context.Context, in *AuthCheckRequest, opts ...grpc.CallOption) (*AuthCheckResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	CodeAuthTransactionStart(ctx context.Context, in *CodeAuthTransactionStartRequest, opts ...grpc.CallOption) (*CodeAuthTransactionStartResponse, error)
	CodeAuthTransactionUpdate(ctx context.Context, in *CodeAuthTransactionUpdateRequest, opts ...grpc.CallOption) (*CodeAuthTransactionUpdateResponse, error)
	CodeAuthTransactionSetPhone(ctx context.Context, in *CodeAuthTransactionSetPhoneRequest, opts ...grpc.CallOption) (*CodeAuthTransactionSetPhoneResponse, error)
	CodeAuthTransactionSetCode(ctx context.Context, in *CodeAuthTransactionSetCodeRequest, opts ...grpc.CallOption) (*CodeAuthTransactionSetCodeResponse, error)
}

type signupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignupServiceClient(cc grpc.ClientConnInterface) SignupServiceClient {
	return &signupServiceClient{cc}
}

func (c *signupServiceClient) SetPhone(ctx context.Context, in *SetPhoneRequest, opts ...grpc.CallOption) (*SetPhoneResponse, error) {
	out := new(SetPhoneResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/SetPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) SetCode(ctx context.Context, in *SetCodeRequest, opts ...grpc.CallOption) (*SetCodeResponse, error) {
	out := new(SetCodeResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/SetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) GetPartnerCode(ctx context.Context, in *GetPartnerCodeRequest, opts ...grpc.CallOption) (*GetPartnerCodeResponse, error) {
	out := new(GetPartnerCodeResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/GetPartnerCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) AuthCheck(ctx context.Context, in *AuthCheckRequest, opts ...grpc.CallOption) (*AuthCheckResponse, error) {
	out := new(AuthCheckResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/AuthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) CodeAuthTransactionStart(ctx context.Context, in *CodeAuthTransactionStartRequest, opts ...grpc.CallOption) (*CodeAuthTransactionStartResponse, error) {
	out := new(CodeAuthTransactionStartResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/CodeAuthTransactionStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) CodeAuthTransactionUpdate(ctx context.Context, in *CodeAuthTransactionUpdateRequest, opts ...grpc.CallOption) (*CodeAuthTransactionUpdateResponse, error) {
	out := new(CodeAuthTransactionUpdateResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/CodeAuthTransactionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) CodeAuthTransactionSetPhone(ctx context.Context, in *CodeAuthTransactionSetPhoneRequest, opts ...grpc.CallOption) (*CodeAuthTransactionSetPhoneResponse, error) {
	out := new(CodeAuthTransactionSetPhoneResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/CodeAuthTransactionSetPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupServiceClient) CodeAuthTransactionSetCode(ctx context.Context, in *CodeAuthTransactionSetCodeRequest, opts ...grpc.CallOption) (*CodeAuthTransactionSetCodeResponse, error) {
	out := new(CodeAuthTransactionSetCodeResponse)
	err := c.cc.Invoke(ctx, "/signup_service.SignupService/CodeAuthTransactionSetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServiceServer is the server API for SignupService service.
// All implementations must embed UnimplementedSignupServiceServer
// for forward compatibility
type SignupServiceServer interface {
	SetPhone(context.Context, *SetPhoneRequest) (*SetPhoneResponse, error)
	SetCode(context.Context, *SetCodeRequest) (*SetCodeResponse, error)
	GetPartnerCode(context.Context, *GetPartnerCodeRequest) (*GetPartnerCodeResponse, error)
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	AuthCheck(context.Context, *AuthCheckRequest) (*AuthCheckResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	CodeAuthTransactionStart(context.Context, *CodeAuthTransactionStartRequest) (*CodeAuthTransactionStartResponse, error)
	CodeAuthTransactionUpdate(context.Context, *CodeAuthTransactionUpdateRequest) (*CodeAuthTransactionUpdateResponse, error)
	CodeAuthTransactionSetPhone(context.Context, *CodeAuthTransactionSetPhoneRequest) (*CodeAuthTransactionSetPhoneResponse, error)
	CodeAuthTransactionSetCode(context.Context, *CodeAuthTransactionSetCodeRequest) (*CodeAuthTransactionSetCodeResponse, error)
	mustEmbedUnimplementedSignupServiceServer()
}

// UnimplementedSignupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignupServiceServer struct {
}

func (UnimplementedSignupServiceServer) SetPhone(context.Context, *SetPhoneRequest) (*SetPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPhone not implemented")
}
func (UnimplementedSignupServiceServer) SetCode(context.Context, *SetCodeRequest) (*SetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCode not implemented")
}
func (UnimplementedSignupServiceServer) GetPartnerCode(context.Context, *GetPartnerCodeRequest) (*GetPartnerCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartnerCode not implemented")
}
func (UnimplementedSignupServiceServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedSignupServiceServer) AuthCheck(context.Context, *AuthCheckRequest) (*AuthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCheck not implemented")
}
func (UnimplementedSignupServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSignupServiceServer) CodeAuthTransactionStart(context.Context, *CodeAuthTransactionStartRequest) (*CodeAuthTransactionStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeAuthTransactionStart not implemented")
}
func (UnimplementedSignupServiceServer) CodeAuthTransactionUpdate(context.Context, *CodeAuthTransactionUpdateRequest) (*CodeAuthTransactionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeAuthTransactionUpdate not implemented")
}
func (UnimplementedSignupServiceServer) CodeAuthTransactionSetPhone(context.Context, *CodeAuthTransactionSetPhoneRequest) (*CodeAuthTransactionSetPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeAuthTransactionSetPhone not implemented")
}
func (UnimplementedSignupServiceServer) CodeAuthTransactionSetCode(context.Context, *CodeAuthTransactionSetCodeRequest) (*CodeAuthTransactionSetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeAuthTransactionSetCode not implemented")
}
func (UnimplementedSignupServiceServer) mustEmbedUnimplementedSignupServiceServer() {}

// UnsafeSignupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignupServiceServer will
// result in compilation errors.
type UnsafeSignupServiceServer interface {
	mustEmbedUnimplementedSignupServiceServer()
}

func RegisterSignupServiceServer(s grpc.ServiceRegistrar, srv SignupServiceServer) {
	s.RegisterService(&SignupService_ServiceDesc, srv)
}

func _SignupService_SetPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).SetPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/SetPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).SetPhone(ctx, req.(*SetPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_SetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).SetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/SetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).SetCode(ctx, req.(*SetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_GetPartnerCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartnerCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).GetPartnerCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/GetPartnerCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).GetPartnerCode(ctx, req.(*GetPartnerCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_AuthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).AuthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/AuthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).AuthCheck(ctx, req.(*AuthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_CodeAuthTransactionStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeAuthTransactionStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).CodeAuthTransactionStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/CodeAuthTransactionStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).CodeAuthTransactionStart(ctx, req.(*CodeAuthTransactionStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_CodeAuthTransactionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeAuthTransactionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).CodeAuthTransactionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/CodeAuthTransactionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).CodeAuthTransactionUpdate(ctx, req.(*CodeAuthTransactionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_CodeAuthTransactionSetPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeAuthTransactionSetPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).CodeAuthTransactionSetPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/CodeAuthTransactionSetPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).CodeAuthTransactionSetPhone(ctx, req.(*CodeAuthTransactionSetPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupService_CodeAuthTransactionSetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeAuthTransactionSetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServiceServer).CodeAuthTransactionSetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signup_service.SignupService/CodeAuthTransactionSetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServiceServer).CodeAuthTransactionSetCode(ctx, req.(*CodeAuthTransactionSetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignupService_ServiceDesc is the grpc.ServiceDesc for SignupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signup_service.SignupService",
	HandlerType: (*SignupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPhone",
			Handler:    _SignupService_SetPhone_Handler,
		},
		{
			MethodName: "SetCode",
			Handler:    _SignupService_SetCode_Handler,
		},
		{
			MethodName: "GetPartnerCode",
			Handler:    _SignupService_GetPartnerCode_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _SignupService_Auth_Handler,
		},
		{
			MethodName: "AuthCheck",
			Handler:    _SignupService_AuthCheck_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SignupService_Logout_Handler,
		},
		{
			MethodName: "CodeAuthTransactionStart",
			Handler:    _SignupService_CodeAuthTransactionStart_Handler,
		},
		{
			MethodName: "CodeAuthTransactionUpdate",
			Handler:    _SignupService_CodeAuthTransactionUpdate_Handler,
		},
		{
			MethodName: "CodeAuthTransactionSetPhone",
			Handler:    _SignupService_CodeAuthTransactionSetPhone_Handler,
		},
		{
			MethodName: "CodeAuthTransactionSetCode",
			Handler:    _SignupService_CodeAuthTransactionSetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signup_service/signup_service.proto",
}