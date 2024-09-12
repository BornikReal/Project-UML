package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Repository interface {
	Create(ctx context.Context, req CreateUserRequest) (int64, error)
	Get(ctx context.Context, id int64) (userModel, error)
	Update(ctx context.Context, req UpdateUserRequest) error
	Login(ctx context.Context, username, password string) (authDataModel, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateUserRequest) (int64, error) {
	id, err := s.repo.Create(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("repo.Create: %w", err)
	}
	return id, nil
}

func (s *Service) Get(ctx context.Context, id int64) (User, error) {
	model, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("repo.Get: %w", err)
	}
	var user User
	user.fromModel(model)
	return user, nil
}

func (s *Service) Update(ctx context.Context, req UpdateUserRequest) error {
	err := s.repo.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("repo.Update: %w", err)
	}
	return nil
}

func (s *Service) Login(ctx context.Context, username, password string) (string, error) {
	authData, err := s.repo.Login(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("repo.Login: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"aud": authData.Role,
		"id":  authData.ID,
		"exp": time.Now().UTC().Add(time.Hour * 72).Unix(),
	})
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}
	return t, nil
}

func (s *Service) GetAuthData(token string) (AuthData, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return AuthData{}, fmt.Errorf("jwt.Parse: %w", err)
	}
	if !parsedToken.Valid {
		return AuthData{}, errors.New("unauthorized")
	}
	expirationTime, err := parsedToken.Claims.GetExpirationTime()
	if err != nil {
		return AuthData{}, fmt.Errorf("parsedToken.Claims.GetExpirationTime: %w", err)
	}
	if expirationTime == nil {
		return AuthData{}, errors.New("expirationTime is nil")
	}
	if expirationTime.Before(time.Now().UTC()) {
		return AuthData{}, errors.New("token is expired")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return AuthData{}, errors.New("claims is invalid")
	}
	role, ok := claims["aud"].(string)
	if !ok {
		return AuthData{}, errors.New("role is invalid")
	}
	id, ok := claims["id"].(int64)
	if !ok {
		idFloat, floatOk := claims["id"].(float64)
		if !floatOk {
			return AuthData{}, errors.New("id is invalid")
		}
		id = int64(idFloat)
	}

	return AuthData{
		ID:   id,
		Role: role,
	}, nil
}
