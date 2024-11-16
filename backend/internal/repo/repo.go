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
		Participant:   postgres.NewParticipantRepo(pg),
		Team:          postgres.NewTeamRepo(pg),
	}
}

type Repository struct {
	Authorization
	Participant
	Team
}

type Authorization interface {
	CreateUser(ctx context.Context, u entities.User) (int, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	GetUserById(ctx context.Context, id int) (entities.User, error)
	DeleteUserById(ctx context.Context, id int) error
	GetUserByCredentials(ctx context.Context, email, password string) (entities.User, error)
}

type Participant interface {
	CreateParticipant(ctx context.Context, p entities.Participant) (int, error)
	GetAllParticipants(ctx context.Context) ([]entities.Participant, error)
	GetParticipantById(ctx context.Context, id int) (entities.Participant, error)
	GetParticipantsByTeamId(ctx context.Context, teamId int) ([]entities.Participant, error)
	UpdateParticipant(ctx context.Context, p entities.Participant) error
	DeleteParticipantById(ctx context.Context, id int) error
	GetParticipantByCreatorId(ctx context.Context, creatorId int) ([]entities.Participant, error)
}

type Team interface {
	CreateTeam(ctx context.Context, t entities.Team) (int, error)
	GetAllTeams(ctx context.Context) ([]entities.Team, error)
	GetTeamById(ctx context.Context, id int) (entities.Team, error)
	GetTeamsByUserId(ctx context.Context, userId int) ([]entities.Team, error)
	UpdateTeam(ctx context.Context, t entities.Team) error
	DeleteTeamById(ctx context.Context, id int) error
}
