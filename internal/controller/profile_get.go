package controller

import (
	"context"
	"fmt"

	managerpb "github.com/voikin/apim-proto/gen/go/apim_manager/v1"
	profilestorepb "github.com/voikin/apim-proto/gen/go/apim_profile_store/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controller) GetProfileByID(
	ctx context.Context,
	req *managerpb.GetProfileByIDRequest,
) (*managerpb.GetProfileByIDResponse, error) {
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

	return &managerpb.GetProfileByIDResponse{
		Profile: profile.GetProfile(),
	}, nil
}
