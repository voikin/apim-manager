package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
)

func (c *Controller) ListApplications(
	ctx context.Context,
	_ *managerpb.ListApplicationsRequest,
) (*managerpb.ListApplicationsResponse, error) {
	applications, err := c.profileStoreClient.ListApplications(ctx, &profilestorepb.ListApplicationsRequest{})
	if err != nil {
		return nil, fmt.Errorf("c.profileStoreClient.ListApplications: %w", err)
	}

	return &managerpb.ListApplicationsResponse{
		Applications: applications.GetApplications(),
	}, nil
}
