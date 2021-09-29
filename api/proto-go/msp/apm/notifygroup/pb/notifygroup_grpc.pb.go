// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: notifygroup.proto

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

// NotifyGroupServiceClient is the client API for NotifyGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyGroupServiceClient interface {
	CreateNotifyGroup(ctx context.Context, in *CreateNotifyGroupRequest, opts ...grpc.CallOption) (*CreateNotifyGroupResponse, error)
	QueryNotifyGroup(ctx context.Context, in *QueryNotifyGroupRequest, opts ...grpc.CallOption) (*QueryNotifyGroupResponse, error)
	GetNotifyGroup(ctx context.Context, in *GetNotifyGroupRequest, opts ...grpc.CallOption) (*GetNotifyGroupResponse, error)
	UpdateNotifyGroup(ctx context.Context, in *UpdateNotifyGroupRequest, opts ...grpc.CallOption) (*UpdateNotifyGroupResponse, error)
	GetNotifyGroupDetail(ctx context.Context, in *GetNotifyGroupDetailRequest, opts ...grpc.CallOption) (*GetNotifyGroupDetailResponse, error)
	DeleteNotifyGroup(ctx context.Context, in *DeleteNotifyGroupRequest, opts ...grpc.CallOption) (*DeleteNotifyGroupResponse, error)
}

type notifyGroupServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewNotifyGroupServiceClient(cc grpc1.ClientConnInterface) NotifyGroupServiceClient {
	return &notifyGroupServiceClient{cc}
}

func (c *notifyGroupServiceClient) CreateNotifyGroup(ctx context.Context, in *CreateNotifyGroupRequest, opts ...grpc.CallOption) (*CreateNotifyGroupResponse, error) {
	out := new(CreateNotifyGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/CreateNotifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyGroupServiceClient) QueryNotifyGroup(ctx context.Context, in *QueryNotifyGroupRequest, opts ...grpc.CallOption) (*QueryNotifyGroupResponse, error) {
	out := new(QueryNotifyGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/QueryNotifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyGroupServiceClient) GetNotifyGroup(ctx context.Context, in *GetNotifyGroupRequest, opts ...grpc.CallOption) (*GetNotifyGroupResponse, error) {
	out := new(GetNotifyGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/GetNotifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyGroupServiceClient) UpdateNotifyGroup(ctx context.Context, in *UpdateNotifyGroupRequest, opts ...grpc.CallOption) (*UpdateNotifyGroupResponse, error) {
	out := new(UpdateNotifyGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/UpdateNotifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyGroupServiceClient) GetNotifyGroupDetail(ctx context.Context, in *GetNotifyGroupDetailRequest, opts ...grpc.CallOption) (*GetNotifyGroupDetailResponse, error) {
	out := new(GetNotifyGroupDetailResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/GetNotifyGroupDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyGroupServiceClient) DeleteNotifyGroup(ctx context.Context, in *DeleteNotifyGroupRequest, opts ...grpc.CallOption) (*DeleteNotifyGroupResponse, error) {
	out := new(DeleteNotifyGroupResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.notifygroup.NotifyGroupService/DeleteNotifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyGroupServiceServer is the server API for NotifyGroupService service.
// All implementations should embed UnimplementedNotifyGroupServiceServer
// for forward compatibility
type NotifyGroupServiceServer interface {
	CreateNotifyGroup(context.Context, *CreateNotifyGroupRequest) (*CreateNotifyGroupResponse, error)
	QueryNotifyGroup(context.Context, *QueryNotifyGroupRequest) (*QueryNotifyGroupResponse, error)
	GetNotifyGroup(context.Context, *GetNotifyGroupRequest) (*GetNotifyGroupResponse, error)
	UpdateNotifyGroup(context.Context, *UpdateNotifyGroupRequest) (*UpdateNotifyGroupResponse, error)
	GetNotifyGroupDetail(context.Context, *GetNotifyGroupDetailRequest) (*GetNotifyGroupDetailResponse, error)
	DeleteNotifyGroup(context.Context, *DeleteNotifyGroupRequest) (*DeleteNotifyGroupResponse, error)
}

// UnimplementedNotifyGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotifyGroupServiceServer struct {
}

func (*UnimplementedNotifyGroupServiceServer) CreateNotifyGroup(context.Context, *CreateNotifyGroupRequest) (*CreateNotifyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifyGroup not implemented")
}
func (*UnimplementedNotifyGroupServiceServer) QueryNotifyGroup(context.Context, *QueryNotifyGroupRequest) (*QueryNotifyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNotifyGroup not implemented")
}
func (*UnimplementedNotifyGroupServiceServer) GetNotifyGroup(context.Context, *GetNotifyGroupRequest) (*GetNotifyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyGroup not implemented")
}
func (*UnimplementedNotifyGroupServiceServer) UpdateNotifyGroup(context.Context, *UpdateNotifyGroupRequest) (*UpdateNotifyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifyGroup not implemented")
}
func (*UnimplementedNotifyGroupServiceServer) GetNotifyGroupDetail(context.Context, *GetNotifyGroupDetailRequest) (*GetNotifyGroupDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyGroupDetail not implemented")
}
func (*UnimplementedNotifyGroupServiceServer) DeleteNotifyGroup(context.Context, *DeleteNotifyGroupRequest) (*DeleteNotifyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifyGroup not implemented")
}

func RegisterNotifyGroupServiceServer(s grpc1.ServiceRegistrar, srv NotifyGroupServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_NotifyGroupService_serviceDesc(srv, opts...), srv)
}

var _NotifyGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.apm.notifygroup.NotifyGroupService",
	HandlerType: (*NotifyGroupServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "notifygroup.proto",
}

func _get_NotifyGroupService_serviceDesc(srv NotifyGroupServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_NotifyGroupService_CreateNotifyGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateNotifyGroup(ctx, req.(*CreateNotifyGroupRequest))
	}
	var _NotifyGroupService_CreateNotifyGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_CreateNotifyGroup_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "CreateNotifyGroup", srv)
		_NotifyGroupService_CreateNotifyGroup_Handler = h.Interceptor(_NotifyGroupService_CreateNotifyGroup_Handler)
	}

	_NotifyGroupService_QueryNotifyGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QueryNotifyGroup(ctx, req.(*QueryNotifyGroupRequest))
	}
	var _NotifyGroupService_QueryNotifyGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_QueryNotifyGroup_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "QueryNotifyGroup", srv)
		_NotifyGroupService_QueryNotifyGroup_Handler = h.Interceptor(_NotifyGroupService_QueryNotifyGroup_Handler)
	}

	_NotifyGroupService_GetNotifyGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyGroup(ctx, req.(*GetNotifyGroupRequest))
	}
	var _NotifyGroupService_GetNotifyGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_GetNotifyGroup_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "GetNotifyGroup", srv)
		_NotifyGroupService_GetNotifyGroup_Handler = h.Interceptor(_NotifyGroupService_GetNotifyGroup_Handler)
	}

	_NotifyGroupService_UpdateNotifyGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateNotifyGroup(ctx, req.(*UpdateNotifyGroupRequest))
	}
	var _NotifyGroupService_UpdateNotifyGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_UpdateNotifyGroup_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "UpdateNotifyGroup", srv)
		_NotifyGroupService_UpdateNotifyGroup_Handler = h.Interceptor(_NotifyGroupService_UpdateNotifyGroup_Handler)
	}

	_NotifyGroupService_GetNotifyGroupDetail_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetNotifyGroupDetail(ctx, req.(*GetNotifyGroupDetailRequest))
	}
	var _NotifyGroupService_GetNotifyGroupDetail_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_GetNotifyGroupDetail_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "GetNotifyGroupDetail", srv)
		_NotifyGroupService_GetNotifyGroupDetail_Handler = h.Interceptor(_NotifyGroupService_GetNotifyGroupDetail_Handler)
	}

	_NotifyGroupService_DeleteNotifyGroup_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteNotifyGroup(ctx, req.(*DeleteNotifyGroupRequest))
	}
	var _NotifyGroupService_DeleteNotifyGroup_info transport.ServiceInfo
	if h.Interceptor != nil {
		_NotifyGroupService_DeleteNotifyGroup_info = transport.NewServiceInfo("erda.msp.apm.notifygroup.NotifyGroupService", "DeleteNotifyGroup", srv)
		_NotifyGroupService_DeleteNotifyGroup_Handler = h.Interceptor(_NotifyGroupService_DeleteNotifyGroup_Handler)
	}

	var serviceDesc = _NotifyGroupService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "CreateNotifyGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateNotifyGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).CreateNotifyGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_CreateNotifyGroup_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_CreateNotifyGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/CreateNotifyGroup",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_CreateNotifyGroup_Handler)
			},
		},
		{
			MethodName: "QueryNotifyGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(QueryNotifyGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).QueryNotifyGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_QueryNotifyGroup_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_QueryNotifyGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/QueryNotifyGroup",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_QueryNotifyGroup_Handler)
			},
		},
		{
			MethodName: "GetNotifyGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).GetNotifyGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_GetNotifyGroup_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_GetNotifyGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/GetNotifyGroup",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_GetNotifyGroup_Handler)
			},
		},
		{
			MethodName: "UpdateNotifyGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateNotifyGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).UpdateNotifyGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_UpdateNotifyGroup_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_UpdateNotifyGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/UpdateNotifyGroup",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_UpdateNotifyGroup_Handler)
			},
		},
		{
			MethodName: "GetNotifyGroupDetail",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetNotifyGroupDetailRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).GetNotifyGroupDetail(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_GetNotifyGroupDetail_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_GetNotifyGroupDetail_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/GetNotifyGroupDetail",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_GetNotifyGroupDetail_Handler)
			},
		},
		{
			MethodName: "DeleteNotifyGroup",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteNotifyGroupRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(NotifyGroupServiceServer).DeleteNotifyGroup(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _NotifyGroupService_DeleteNotifyGroup_info)
				}
				if interceptor == nil {
					return _NotifyGroupService_DeleteNotifyGroup_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.notifygroup.NotifyGroupService/DeleteNotifyGroup",
				}
				return interceptor(ctx, in, info, _NotifyGroupService_DeleteNotifyGroup_Handler)
			},
		},
	}
	return &serviceDesc
}
