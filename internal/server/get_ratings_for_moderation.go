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

func (i *Implementation) GetRatingsForModeration(ctx context.Context, _ *emptypb.Empty) (*desc.GetRatingsForModerationResponse, error) {
	authData, err := i.validUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if authData.Role != user.RoleModerator {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	rest, err := i.restrictionsService.GetObjectRestrictions(ctx, restrictions.GetObjectRestrictions{
		RestrictionOn: restrictions.Rating,
		Type:          utils.Of(restrictions.ShadowBan),
	})

	ratingsID := make([]int64, len(rest))
	for j, r := range rest {
		ratingsID[j] = r.ObjectID
	}

	ratings, err := i.ratingService.GetByIDs(ctx, ratingsID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ratingsModeration := make([]*desc.RatingModeration, len(ratings))
	for j, r := range ratings {
		ratingsModeration[j] = &desc.RatingModeration{
			RatingId: r.ID,
			Score:    r.Score,
			Review:   r.Review,
			UserId:   r.UserID,
		}
	}

	return &desc.GetRatingsForModerationResponse{
		Ratings: ratingsModeration,
	}, nil
}
