// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: org_client.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/erda-project/erda-proto-go/core/hepa/openapi_rule/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ChangeClientLimitResponse) Validate() error {
	return nil
}
func (this *ChangeClientLimitRequest) Validate() error {
	for _, item := range this.Limits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Limits", err)
			}
		}
	}
	return nil
}
func (this *GrantEndpointRequest) Validate() error {
	return nil
}
func (this *GrantEndpointResponse) Validate() error {
	return nil
}
func (this *RevokeEndpointRequest) Validate() error {
	return nil
}
func (this *RevokeEndpointResponse) Validate() error {
	return nil
}
func (this *UpdateCredentialsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateCredentialsRequest) Validate() error {
	return nil
}
func (this *GetCredentialsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetCredentialsRequest) Validate() error {
	return nil
}
func (this *DeleteClientRequest) Validate() error {
	return nil
}
func (this *DeleteClientResponse) Validate() error {
	return nil
}
func (this *ClientInfo) Validate() error {
	return nil
}
func (this *CreateClientRequest) Validate() error {
	return nil
}
func (this *CreateClientResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
