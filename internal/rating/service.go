package rating

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, rating Rating) (int64, error)
	GetByIDs(ctx context.Context, id []int64) ([]ratingModel, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, rating Rating) (int64, error) {
	id, err := s.repo.Create(ctx, rating)
	if err != nil {
		return 0, fmt.Errorf("repo.Create: %w", err)
	}
	return id, nil
}

func (s *Service) GetByIDs(ctx context.Context, id []int64) ([]Rating, error) {
	models, err := s.repo.GetByIDs(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByIDs: %w", err)
	}
	ratings := make([]Rating, len(models))
	for i, model := range models {
		ratings[i] = Rating{
			ID:        model.ID,
			Score:     model.Score,
			Review:    model.Review,
			FilmID:    model.FilmID,
			UserID:    model.UserID,
			IsSpecial: model.IsSpecial,
		}
	}
	return ratings, nil
}
