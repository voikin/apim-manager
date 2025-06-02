package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) DeleteProfile(
	ctx context.Context,
	req *managerpb.DeleteProfileRequest,
) (*managerpb.DeleteProfileResponse, error) {
	_, err := c.profileStoreClient.DeleteProfile(ctx, &profilestorepb.DeleteProfileRequest{
		Id: req.GetId(),
	})
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "profile %s not found", req.GetId())
		}

		return nil, fmt.Errorf("c.profileStoreClient.DeleteProfile: %w", err)
	}

	return &managerpb.DeleteProfileResponse{}, nil
}
