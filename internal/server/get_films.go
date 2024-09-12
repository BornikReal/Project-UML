package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "project4/pkg/service-component/pb"
)

func (i *Implementation) GetFilms(ctx context.Context, _ *emptypb.Empty) (*desc.GetFilmsResponse, error) {
	films, err := i.filmService.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	protoFilms := make([]*desc.Film, len(films))
	for j, film := range films {
		protoFilms[j] = &desc.Film{
			Id:          film.ID,
			Name:        film.Name,
			Description: film.Description,
			Cover:       film.Cover,
			Genres:      film.Genres,
			Actors:      film.Actors,
			Images:      film.Images,
			Trailers:    film.Trailers,
			Ratings:     film.Ratings,
		}
	}

	return &desc.GetFilmsResponse{
		Film: protoFilms,
	}, nil
}
