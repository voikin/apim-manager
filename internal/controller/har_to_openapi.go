package controller

import (
	"context"
	"fmt"

	harprofilerpb "github.com/voikin/apim-proto/gen/go/apim_har_profiler/v1"
	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) GenerateOpenAPISpecFromHAR(ctx context.Context, req *managerpb.GenerateOpenAPISpecFromHARRequest) (*managerpb.GenerateOpenAPISpecFromHARResponse, error) {
	_, err := c.profileStoreClient.GetApplication(ctx, &profilestorepb.GetApplicationRequest{
		Id: req.GetApplicationId(),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, fmt.Errorf("c.profileStoreClient.GetApplication: %w", err)
	}

	apiGraphRepsponse, err := c.harProfilerClient.BuildAPIGraph(ctx, &harprofilerpb.BuildAPIGraphRequest{
		HarJson: req.GetHarJson(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.harProfilerClient.BuildAPIGraph: %w", err)
	}

	_, err = c.profileStoreClient.AddProfile(ctx, &profilestorepb.AddProfileRequest{
		ApplicationId: req.GetApplicationId(),
		ApiGraph:      apiGraphRepsponse.GetGraph(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.profileStoreClient.AddProfile: %w", err)
	}

	return &managerpb.GenerateOpenAPISpecFromHARResponse{}, nil
}
