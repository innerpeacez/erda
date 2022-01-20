// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: definition.proto

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

// DefinitionServiceClient is the client API for DefinitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefinitionServiceClient interface {
	Create(ctx context.Context, in *PipelineDefinitionCreateRequest, opts ...grpc.CallOption) (*PipelineDefinitionCreateResponse, error)
	Update(ctx context.Context, in *PipelineDefinitionUpdateRequest, opts ...grpc.CallOption) (*PipelineDefinitionUpdateResponse, error)
	Delete(ctx context.Context, in *PipelineDefinitionDeleteRequest, opts ...grpc.CallOption) (*PipelineDefinitionDeleteResponse, error)
	Get(ctx context.Context, in *PipelineDefinitionGetRequest, opts ...grpc.CallOption) (*PipelineDefinitionGetResponse, error)
	List(ctx context.Context, in *PipelineDefinitionListRequest, opts ...grpc.CallOption) (*PipelineDefinitionListResponse, error)
	StaticsGroupByRemote(ctx context.Context, in *PipelineDefinitionStaticsRequest, opts ...grpc.CallOption) (*PipelineDefinitionStaticsResponse, error)
}

type definitionServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewDefinitionServiceClient(cc grpc1.ClientConnInterface) DefinitionServiceClient {
	return &definitionServiceClient{cc}
}

func (c *definitionServiceClient) Create(ctx context.Context, in *PipelineDefinitionCreateRequest, opts ...grpc.CallOption) (*PipelineDefinitionCreateResponse, error) {
	out := new(PipelineDefinitionCreateResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionServiceClient) Update(ctx context.Context, in *PipelineDefinitionUpdateRequest, opts ...grpc.CallOption) (*PipelineDefinitionUpdateResponse, error) {
	out := new(PipelineDefinitionUpdateResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionServiceClient) Delete(ctx context.Context, in *PipelineDefinitionDeleteRequest, opts ...grpc.CallOption) (*PipelineDefinitionDeleteResponse, error) {
	out := new(PipelineDefinitionDeleteResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionServiceClient) Get(ctx context.Context, in *PipelineDefinitionGetRequest, opts ...grpc.CallOption) (*PipelineDefinitionGetResponse, error) {
	out := new(PipelineDefinitionGetResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionServiceClient) List(ctx context.Context, in *PipelineDefinitionListRequest, opts ...grpc.CallOption) (*PipelineDefinitionListResponse, error) {
	out := new(PipelineDefinitionListResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *definitionServiceClient) StaticsGroupByRemote(ctx context.Context, in *PipelineDefinitionStaticsRequest, opts ...grpc.CallOption) (*PipelineDefinitionStaticsResponse, error) {
	out := new(PipelineDefinitionStaticsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.definition.DefinitionService/StaticsGroupByRemote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefinitionServiceServer is the server API for DefinitionService service.
// All implementations should embed UnimplementedDefinitionServiceServer
// for forward compatibility
type DefinitionServiceServer interface {
	Create(context.Context, *PipelineDefinitionCreateRequest) (*PipelineDefinitionCreateResponse, error)
	Update(context.Context, *PipelineDefinitionUpdateRequest) (*PipelineDefinitionUpdateResponse, error)
	Delete(context.Context, *PipelineDefinitionDeleteRequest) (*PipelineDefinitionDeleteResponse, error)
	Get(context.Context, *PipelineDefinitionGetRequest) (*PipelineDefinitionGetResponse, error)
	List(context.Context, *PipelineDefinitionListRequest) (*PipelineDefinitionListResponse, error)
	StaticsGroupByRemote(context.Context, *PipelineDefinitionStaticsRequest) (*PipelineDefinitionStaticsResponse, error)
}

// UnimplementedDefinitionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDefinitionServiceServer struct {
}

func (*UnimplementedDefinitionServiceServer) Create(context.Context, *PipelineDefinitionCreateRequest) (*PipelineDefinitionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDefinitionServiceServer) Update(context.Context, *PipelineDefinitionUpdateRequest) (*PipelineDefinitionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDefinitionServiceServer) Delete(context.Context, *PipelineDefinitionDeleteRequest) (*PipelineDefinitionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDefinitionServiceServer) Get(context.Context, *PipelineDefinitionGetRequest) (*PipelineDefinitionGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDefinitionServiceServer) List(context.Context, *PipelineDefinitionListRequest) (*PipelineDefinitionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDefinitionServiceServer) StaticsGroupByRemote(context.Context, *PipelineDefinitionStaticsRequest) (*PipelineDefinitionStaticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaticsGroupByRemote not implemented")
}

func RegisterDefinitionServiceServer(s grpc1.ServiceRegistrar, srv DefinitionServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_DefinitionService_serviceDesc(srv, opts...), srv)
}

var _DefinitionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.pipeline.definition.DefinitionService",
	HandlerType: (*DefinitionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "definition.proto",
}

func _get_DefinitionService_serviceDesc(srv DefinitionServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_DefinitionService_Create_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Create(ctx, req.(*PipelineDefinitionCreateRequest))
	}
	var _DefinitionService_Create_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_Create_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "Create", srv)
		_DefinitionService_Create_Handler = h.Interceptor(_DefinitionService_Create_Handler)
	}

	_DefinitionService_Update_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Update(ctx, req.(*PipelineDefinitionUpdateRequest))
	}
	var _DefinitionService_Update_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_Update_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "Update", srv)
		_DefinitionService_Update_Handler = h.Interceptor(_DefinitionService_Update_Handler)
	}

	_DefinitionService_Delete_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Delete(ctx, req.(*PipelineDefinitionDeleteRequest))
	}
	var _DefinitionService_Delete_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_Delete_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "Delete", srv)
		_DefinitionService_Delete_Handler = h.Interceptor(_DefinitionService_Delete_Handler)
	}

	_DefinitionService_Get_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Get(ctx, req.(*PipelineDefinitionGetRequest))
	}
	var _DefinitionService_Get_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_Get_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "Get", srv)
		_DefinitionService_Get_Handler = h.Interceptor(_DefinitionService_Get_Handler)
	}

	_DefinitionService_List_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.List(ctx, req.(*PipelineDefinitionListRequest))
	}
	var _DefinitionService_List_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_List_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "List", srv)
		_DefinitionService_List_Handler = h.Interceptor(_DefinitionService_List_Handler)
	}

	_DefinitionService_StaticsGroupByRemote_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.StaticsGroupByRemote(ctx, req.(*PipelineDefinitionStaticsRequest))
	}
	var _DefinitionService_StaticsGroupByRemote_info transport.ServiceInfo
	if h.Interceptor != nil {
		_DefinitionService_StaticsGroupByRemote_info = transport.NewServiceInfo("erda.core.pipeline.definition.DefinitionService", "StaticsGroupByRemote", srv)
		_DefinitionService_StaticsGroupByRemote_Handler = h.Interceptor(_DefinitionService_StaticsGroupByRemote_Handler)
	}

	var serviceDesc = _DefinitionService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionCreateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).Create(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_Create_info)
				}
				if interceptor == nil {
					return _DefinitionService_Create_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/Create",
				}
				return interceptor(ctx, in, info, _DefinitionService_Create_Handler)
			},
		},
		{
			MethodName: "Update",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionUpdateRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).Update(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_Update_info)
				}
				if interceptor == nil {
					return _DefinitionService_Update_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/Update",
				}
				return interceptor(ctx, in, info, _DefinitionService_Update_Handler)
			},
		},
		{
			MethodName: "Delete",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionDeleteRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).Delete(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_Delete_info)
				}
				if interceptor == nil {
					return _DefinitionService_Delete_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/Delete",
				}
				return interceptor(ctx, in, info, _DefinitionService_Delete_Handler)
			},
		},
		{
			MethodName: "Get",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionGetRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).Get(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_Get_info)
				}
				if interceptor == nil {
					return _DefinitionService_Get_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/Get",
				}
				return interceptor(ctx, in, info, _DefinitionService_Get_Handler)
			},
		},
		{
			MethodName: "List",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionListRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).List(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_List_info)
				}
				if interceptor == nil {
					return _DefinitionService_List_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/List",
				}
				return interceptor(ctx, in, info, _DefinitionService_List_Handler)
			},
		},
		{
			MethodName: "StaticsGroupByRemote",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PipelineDefinitionStaticsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(DefinitionServiceServer).StaticsGroupByRemote(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _DefinitionService_StaticsGroupByRemote_info)
				}
				if interceptor == nil {
					return _DefinitionService_StaticsGroupByRemote_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.definition.DefinitionService/StaticsGroupByRemote",
				}
				return interceptor(ctx, in, info, _DefinitionService_StaticsGroupByRemote_Handler)
			},
		},
	}
	return &serviceDesc
}
