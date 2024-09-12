package film

import (
	"context"
	"fmt"
)

type Repository interface {
	Get(ctx context.Context, id int64) (filmModel, error)
	GetAll(ctx context.Context) ([]filmModel, error)
	AddRating(ctx context.Context, filmID, ratingID int64) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(ctx context.Context, id int64) (Film, error) {
	model, err := s.repo.Get(ctx, id)
	if err != nil {
		return Film{}, fmt.Errorf("repo.Get: %w", err)
	}
	var film Film
	film.fromModel(model)
	return film, nil
}

func (s *Service) GetAll(ctx context.Context) ([]Film, error) {
	models, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.Get: %w", err)
	}
	films := make([]Film, len(models))
	for i, model := range models {
		films[i].fromModel(model)
	}
	return films, nil
}

func (s *Service) AddRating(ctx context.Context, filmID, ratingID int64) error {
	err := s.repo.AddRating(ctx, filmID, ratingID)
	if err != nil {
		return fmt.Errorf("repo.AddRating: %w", err)
	}
	return nil
}
