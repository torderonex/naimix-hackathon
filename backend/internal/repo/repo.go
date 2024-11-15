package repo

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/config"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo/postgres"
)

func New(config *config.Config) *Repository {
	pg := postgres.New(config.Postgres)
	return &Repository{
		Authorization: postgres.NewAuthRepo(pg),
	}
}

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(ctx context.Context, u entities.User) (int, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	GetUserById(ctx context.Context, id int) (entities.User, error)
	DeleteUserById(ctx context.Context, id int) error
	GetUserByNickNPass(ctx context.Context, nickname, password string) (entities.User, error)
}
