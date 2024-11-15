package service

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
)

type Service struct {
	Authorization
}

func New(repo *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}

type Authorization interface {
	CreateUser(ctx context.Context, user entities.User) (int, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (int, error)
}
