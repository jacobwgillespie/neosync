package v1alpha1_jobservice

import (
	"context"

	"connectrpc.com/connect"
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

type Transformation string

const (
	Invalid     Transformation = "invalid"
	Passthrough Transformation = "passthrough"
	Uuid        Transformation = "uuid"
	FirstName   Transformation = "first_name"
	PhoneNumber Transformation = "phone_number"
	Email       Transformation = "email"
	Null        Transformation = "null"
)

func (s *Service) GetTransformers(
	ctx context.Context,
	req *connect.Request[mgmtv1alpha1.GetTransformersRequest],
) (*connect.Response[mgmtv1alpha1.GetTransformersResponse], error) {
	return connect.NewResponse(&mgmtv1alpha1.GetTransformersResponse{
		Transformers: []*mgmtv1alpha1.Transformer{
			{Value: string(Passthrough), Config: &mgmtv1alpha1.TransformerConfig{}},
			{Value: string(Uuid), Config: &mgmtv1alpha1.TransformerConfig{
				Config: &mgmtv1alpha1.TransformerConfig_UuidConfig{
					UuidConfig: &mgmtv1alpha1.Uuid{
						IncludeHyphen: true,
					},
				},
			},
			},
			{Value: string(FirstName), Config: &mgmtv1alpha1.TransformerConfig{}},
			{Value: string(PhoneNumber), Config: &mgmtv1alpha1.TransformerConfig{}},
			{Value: string(Null), Config: &mgmtv1alpha1.TransformerConfig{}},
			{
				Value: string(Email),
				Config: &mgmtv1alpha1.TransformerConfig{
					Config: &mgmtv1alpha1.TransformerConfig_EmailConfig{
						EmailConfig: &mgmtv1alpha1.EmailConfig{
							PreserveDomain: false,
							PreserveLength: false,
						},
					},
				}},
		},
	}), nil
}
