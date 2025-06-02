package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	openapiexporterpb "github.com/voikin/apim-proto/gen/go/apim_openapi_exporter/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) ExportOpenAPISpec(
	ctx context.Context,
	req *managerpb.ExportOpenAPISpecRequest,
) (*managerpb.ExportOpenAPISpecResponse, error) {
	profile, err := c.profileStoreClient.GetProfileByID(ctx, &profilestorepb.GetProfileByIDRequest{
		Id: req.GetId(),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "profile %s not found", req.GetId())
		}

		return nil, fmt.Errorf("c.profileStoreClient.GetProfileByID: %w", err)
	}

	spec, err := c.openapiExporterClient.BuildOpenAPISpec(ctx, &openapiexporterpb.BuildOpenAPISpecRequest{
		ApiGraph: profile.GetProfile().GetApiGraph(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.openapiExporterClient.BuildOpenAPISpec: %w", err)
	}

	return &managerpb.ExportOpenAPISpecResponse{
		SpecJson: spec.GetSpecJson(),
	}, nil
}
