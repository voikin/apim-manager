package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) GetApplication(
	ctx context.Context,
	req *managerpb.GetApplicationRequest,
) (*managerpb.GetApplicationResponse, error) {
	application, err := c.profileStoreClient.GetApplication(ctx, &profilestorepb.GetApplicationRequest{
		Id: req.GetId(),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "application %s not found", req.GetId())
		}

		return nil, fmt.Errorf("c.profileStoreClient.GetApplication: %w", err)
	}

	profiles, err := c.profileStoreClient.ListProfilesByApplication(ctx, &profilestorepb.ListProfilesByApplicationRequest{
		ApplicationId: req.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.profileStoreClient.ListProfilesByApplication: %w", err)
	}

	return &managerpb.GetApplicationResponse{
		Application: application.GetApplication(),
		Profiles:    profiles.GetProfiles(),
	}, nil
}
