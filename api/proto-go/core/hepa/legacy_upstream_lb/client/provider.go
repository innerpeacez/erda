// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: legacy_upstream_lb.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/legacy_upstream_lb/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.legacy_upstream_lb",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                 = reflect.TypeOf((*Client)(nil)).Elem()
	upstreamLbServiceClientType = reflect.TypeOf((*pb.UpstreamLbServiceClient)(nil)).Elem()
	upstreamLbServiceServerType = reflect.TypeOf((*pb.UpstreamLbServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.legacy_upstream_lb-client":
		return p.client
	case "erda.core.hepa.legacy_upstream_lb.UpstreamLbService":
		return &upstreamLbServiceWrapper{client: p.client.UpstreamLbService(), opts: opts}
	case "erda.core.hepa.legacy_upstream_lb.UpstreamLbService.client":
		return p.client.UpstreamLbService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case upstreamLbServiceClientType:
		return p.client.UpstreamLbService()
	case upstreamLbServiceServerType:
		return &upstreamLbServiceWrapper{client: p.client.UpstreamLbService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.legacy_upstream_lb-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.legacy_upstream_lb.UpstreamLbService",
			"erda.core.hepa.legacy_upstream_lb-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			upstreamLbServiceClientType,
			// server types
			upstreamLbServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
