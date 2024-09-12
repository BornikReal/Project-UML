package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"project4/internal/restrictions"
	"project4/internal/user"
	desc "project4/pkg/service-component/pb"
	"project4/pkg/utils"
)

func (i *Implementation) UnlockReview(ctx context.Context, req *desc.UnlockReviewRequest) (*emptypb.Empty, error) {
	authData, err := i.validUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if authData.Role != user.RoleModerator {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	rest, err := i.restrictionsService.GetObjectRestrictions(ctx, restrictions.GetObjectRestrictions{
		ObjectID:      &req.Id,
		RestrictionOn: restrictions.Rating,
		Type:          utils.Of(restrictions.ShadowBan),
	})

	if len(rest) == 0 {
		return nil, status.Error(codes.NotFound, "restrictions not found")
	}

	for _, r := range rest {
		err = i.restrictionsService.DeleteRestriction(ctx, r.ID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}
