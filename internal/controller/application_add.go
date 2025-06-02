package controller

import (
	"context"
	"fmt"

	harprofilerpb "github.com/voikin/apim-proto/gen/go/apim_har_profiler/v1"
	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
)

func (c *Controller) AddProfile(
	ctx context.Context,
	req *managerpb.AddProfileRequest,
) (*managerpb.AddProfileResponse, error) {
	profile, err := c.harProfilerClient.BuildAPIGraph(ctx, &harprofilerpb.BuildAPIGraphRequest{
		HarFiles: req.GetHarFiles(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.harProfilerClient.BuildAPIGraph: %w", err)
	}

	res, err := c.profileStoreClient.AddProfile(ctx, &profilestorepb.AddProfileRequest{
		ApplicationId: req.GetApplicationId(),
		ApiGraph:      profile.GetGraph(),
	})
	if err != nil {
		return nil, fmt.Errorf("c.profileStoreClient.AddProfile: %w", err)
	}

	return &managerpb.AddProfileResponse{
		Profile: res.GetProfile(),
	}, nil
}
