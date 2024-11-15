package service

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
)

type Service struct {
	Authorization
	Team
	Participant
}

func New(repo *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Team:          NewTeamService(repo.Team),
		Participant:   NewParticipantService(repo.Participant),
	}
}

type Authorization interface {
	CreateUser(ctx context.Context, user entities.User) (int, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (int, error)
}

type Team interface {
	CreateTeam(ctx context.Context, t entities.Team) (int, error)
	GetAllTeams(ctx context.Context) ([]entities.Team, error)
	GetTeamById(ctx context.Context, id int) (entities.Team, error)
	GetTeamsByUserId(ctx context.Context, userId int) ([]entities.Team, error)
	UpdateTeam(ctx context.Context, t entities.Team) error
}

type Participant interface {
	CreateParticipant(ctx context.Context, p entities.Participant) (int, error)
	GetAllParticipants(ctx context.Context) ([]entities.Participant, error)
	GetParticipantById(ctx context.Context, id int) (entities.Participant, error)
	GetParticipantsByTeamId(ctx context.Context, teamId int) ([]entities.Participant, error)
	UpdateParticipant(ctx context.Context, p entities.Participant) error
}
