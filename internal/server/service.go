package server

import (
	"context"
	"project4/internal/film"
	"project4/internal/rating"
	"project4/internal/restrictions"

	"project4/internal/user"
	desc "project4/pkg/service-component/pb"
)

type UserService interface {
	Create(ctx context.Context, req user.CreateUserRequest) (int64, error)
	Get(ctx context.Context, id int64) (user.User, error)
	Update(ctx context.Context, req user.UpdateUserRequest) error
	Login(ctx context.Context, username, password string) (string, error)
	GetAuthData(token string) (user.AuthData, error)
}

type FilmService interface {
	Get(ctx context.Context, id int64) (film.Film, error)
	GetAll(ctx context.Context) ([]film.Film, error)
	AddRating(ctx context.Context, filmID, ratingID int64) error
}

type RestrictionsService interface {
	IsUserBanned(ctx context.Context, id int64) (bool, error)
	GetObjectRestrictions(ctx context.Context, req restrictions.GetObjectRestrictions) ([]restrictions.Restriction, error)
	AddRestriction(ctx context.Context, restriction restrictions.Restriction) (int64, error)
	DeleteRestriction(ctx context.Context, id int64) error
}

type RatingService interface {
	Create(ctx context.Context, rating rating.Rating) (int64, error)
	GetByIDs(ctx context.Context, id []int64) ([]rating.Rating, error)
}

type Implementation struct {
	desc.UnimplementedRTServiceServer

	userService         UserService
	filmService         FilmService
	restrictionsService RestrictionsService
	ratingService       RatingService
}

func NewImplementation(userService UserService, filmService FilmService,
	restrictionsService RestrictionsService, ratingService RatingService) *Implementation {
	return &Implementation{
		userService:         userService,
		filmService:         filmService,
		restrictionsService: restrictionsService,
		ratingService:       ratingService,
	}
}
