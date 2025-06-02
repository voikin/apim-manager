package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
)

func (c *Controller) CreateApplication(
	ctx context.Context,
	req *managerpb.CreateApplicationRequest,
) (*managerpb.CreateApplicationResponse, error) {
	res, err := c.profileStoreClient.CreateApplication(ctx, &profilestorepb.CreateApplicationRequest{
		Name: req.GetName(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.profileStoreClient.CreateApplication: %w", err)
	}

	return &managerpb.CreateApplicationResponse{
		Application: res.GetApplication(),
	}, nil
}
