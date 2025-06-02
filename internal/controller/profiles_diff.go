package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) DiffProfiles(
	ctx context.Context,
	req *managerpb.DiffProfilesRequest,
) (*managerpb.DiffProfilesResponse, error) {
	res, err := c.profileStoreClient.DiffProfiles(ctx, &profilestorepb.DiffProfilesRequest{
		ApplicationId: req.GetApplicationId(),
		OldProfileId:  req.GetOldProfileId(),
		NewProfileId:  req.GetNewProfileId(),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			return nil, status.Errorf(
				codes.NotFound,
				"profiles %s or %s not found",
				req.GetOldProfileId(),
				req.GetNewProfileId(),
			)
		}

		return nil, fmt.Errorf("c.profileStoreClient.DiffProfiles: %w", err)
	}

	return &managerpb.DiffProfilesResponse{
		Added:   res.GetAdded(),
		Removed: res.GetRemoved(),
	}, nil
}
