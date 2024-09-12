package restrictions

import (
	"context"
	"fmt"
	"project4/pkg/utils"
)

type Repository interface {
	GetObjectRestrictions(ctx context.Context, req GetObjectRestrictions) ([]restrictionModel, error)
	AddRestriction(ctx context.Context, restriction Restriction) (int64, error)
	DeleteRestriction(ctx context.Context, id int64) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) IsUserBanned(ctx context.Context, id int64) (bool, error) {
	models, err := s.repo.GetObjectRestrictions(ctx, GetObjectRestrictions{
		ObjectID:      &id,
		RestrictionOn: User,
		Type:          utils.Of(Ban),
	})
	if err != nil {
		return false, fmt.Errorf("repo.GetObjectRestrictions: %w", err)
	}
	return len(models) > 0, nil
}

func (s *Service) GetObjectRestrictions(ctx context.Context, req GetObjectRestrictions) ([]Restriction, error) {
	models, err := s.repo.GetObjectRestrictions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("repo.GetObjectRestrictions: %w", err)
	}
	restrictions := make([]Restriction, len(models))
	for i, model := range models {
		restrictions[i] = Restriction{
			ID:            model.ID,
			ValidUntil:    model.ValidUntil,
			Type:          RestrictionType(model.Type),
			ObjectID:      model.ObjectID,
			RestrictionOn: RestrictionObject(model.RestrictionOn),
		}
	}
	return restrictions, nil
}

func (s *Service) AddRestriction(ctx context.Context, restriction Restriction) (int64, error) {
	id, err := s.repo.AddRestriction(ctx, restriction)
	if err != nil {
		return 0, fmt.Errorf("repo.AddRestriction: %w", err)
	}
	return id, nil
}

func (s *Service) DeleteRestriction(ctx context.Context, id int64) error {
	err := s.repo.DeleteRestriction(ctx, id)
	if err != nil {
		return fmt.Errorf("repo.DeleteRestriction: %w", err)
	}
	return nil
}
