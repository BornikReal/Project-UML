package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "project4/pkg/service-component/pb"
)

func (i *Implementation) GetFilm(ctx context.Context, req *desc.GetFilmRequest) (*desc.Film, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Empty request")
	}

	film, err := i.filmService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.Film{
		Id:          film.ID,
		Name:        film.Name,
		Description: film.Description,
		Cover:       film.Cover,
		Genres:      film.Genres,
		Actors:      film.Actors,
		Images:      film.Images,
		Trailers:    film.Trailers,
		Ratings:     film.Ratings,
	}, nil
}
