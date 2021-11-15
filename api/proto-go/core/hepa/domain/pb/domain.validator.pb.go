// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: domain.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/erda-project/erda-proto-go/core/hepa/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetTenantDomainsRequest) Validate() error {
	return nil
}
func (this *GetTenantDomainsResponse) Validate() error {
	return nil
}
func (this *ChangeRuntimeDomainsRequest) Validate() error {
	return nil
}
func (this *ChangeRuntimeDomainsResponse) Validate() error {
	return nil
}
func (this *GetRuntimeDomainsResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetRuntimeDomainsRequest) Validate() error {
	return nil
}
func (this *ChangeInnerIngressResponse) Validate() error {
	return nil
}
func (this *ChangeInnerIngressRequest) Validate() error {
	for _, item := range this.Routes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Routes", err)
			}
		}
	}
	if this.RouteOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RouteOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RouteOptions", err)
		}
	}
	return nil
}
func (this *RouteOptions) Validate() error {
	if this.EnableTls != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EnableTls); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EnableTls", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *IngressRoute) Validate() error {
	return nil
}
func (this *GetOrgDomainsRequest) Validate() error {
	return nil
}
func (this *GetOrgDomainsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
