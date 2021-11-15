// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: api_policy.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ApiPolicyServiceClient is the client API for ApiPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiPolicyServiceClient interface {
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error)
	SetPolicy(ctx context.Context, in *SetPolicyRequest, opts ...grpc.CallOption) (*SetPolicyResponse, error)
}

type apiPolicyServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewApiPolicyServiceClient(cc grpc1.ClientConnInterface) ApiPolicyServiceClient {
	return &apiPolicyServiceClient{cc}
}

func (c *apiPolicyServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error) {
	out := new(GetPolicyResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.api_policy.ApiPolicyService/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiPolicyServiceClient) SetPolicy(ctx context.Context, in *SetPolicyRequest, opts ...grpc.CallOption) (*SetPolicyResponse, error) {
	out := new(SetPolicyResponse)
	err := c.cc.Invoke(ctx, "/erda.core.hepa.api_policy.ApiPolicyService/SetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiPolicyServiceServer is the server API for ApiPolicyService service.
// All implementations should embed UnimplementedApiPolicyServiceServer
// for forward compatibility
type ApiPolicyServiceServer interface {
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	SetPolicy(context.Context, *SetPolicyRequest) (*SetPolicyResponse, error)
}

// UnimplementedApiPolicyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedApiPolicyServiceServer struct {
}

func (*UnimplementedApiPolicyServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (*UnimplementedApiPolicyServiceServer) SetPolicy(context.Context, *SetPolicyRequest) (*SetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPolicy not implemented")
}

func RegisterApiPolicyServiceServer(s grpc1.ServiceRegistrar, srv ApiPolicyServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ApiPolicyService_serviceDesc(srv, opts...), srv)
}

var _ApiPolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.hepa.api_policy.ApiPolicyService",
	HandlerType: (*ApiPolicyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api_policy.proto",
}

func _get_ApiPolicyService_serviceDesc(srv ApiPolicyServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ApiPolicyService_GetPolicy_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	var _ApiPolicyService_GetPolicy_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ApiPolicyService_GetPolicy_info = transport.NewServiceInfo("erda.core.hepa.api_policy.ApiPolicyService", "GetPolicy", srv)
		_ApiPolicyService_GetPolicy_Handler = h.Interceptor(_ApiPolicyService_GetPolicy_Handler)
	}

	_ApiPolicyService_SetPolicy_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.SetPolicy(ctx, req.(*SetPolicyRequest))
	}
	var _ApiPolicyService_SetPolicy_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ApiPolicyService_SetPolicy_info = transport.NewServiceInfo("erda.core.hepa.api_policy.ApiPolicyService", "SetPolicy", srv)
		_ApiPolicyService_SetPolicy_Handler = h.Interceptor(_ApiPolicyService_SetPolicy_Handler)
	}

	var serviceDesc = _ApiPolicyService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetPolicy",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetPolicyRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ApiPolicyServiceServer).GetPolicy(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ApiPolicyService_GetPolicy_info)
				}
				if interceptor == nil {
					return _ApiPolicyService_GetPolicy_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.api_policy.ApiPolicyService/GetPolicy",
				}
				return interceptor(ctx, in, info, _ApiPolicyService_GetPolicy_Handler)
			},
		},
		{
			MethodName: "SetPolicy",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(SetPolicyRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ApiPolicyServiceServer).SetPolicy(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ApiPolicyService_SetPolicy_info)
				}
				if interceptor == nil {
					return _ApiPolicyService_SetPolicy_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.hepa.api_policy.ApiPolicyService/SetPolicy",
				}
				return interceptor(ctx, in, info, _ApiPolicyService_SetPolicy_Handler)
			},
		},
	}
	return &serviceDesc
}
