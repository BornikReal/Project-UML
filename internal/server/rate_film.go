package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"project4/internal/rating"
	"project4/internal/restrictions"
	"project4/internal/user"
	desc "project4/pkg/service-component/pb"
)

func (i *Implementation) RateFilm(ctx context.Context, req *desc.RateFilmRequest) (*emptypb.Empty, error) {
	authData, err := i.validUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	ratingDto := rating.Rating{
		Score:     req.Score,
		FilmID:    req.Id,
		UserID:    authData.ID,
		IsSpecial: authData.Role == user.RoleSpecialUser,
	}
	ratingDto.SetReview(req.Review)

	ratingID, err := i.ratingService.Create(ctx, ratingDto)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = i.filmService.AddRating(ctx, req.Id, ratingID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = i.restrictionsService.AddRestriction(ctx, restrictions.Restriction{
		Type:          restrictions.ShadowBan,
		ObjectID:      ratingID,
		RestrictionOn: restrictions.Rating,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
