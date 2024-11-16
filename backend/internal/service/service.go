package service

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/webapi/sber"
	"time"
)

type Service struct {
	Authorization
	Team
	Participant
	Comparison
}

func New(repo *repo.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Team:          NewTeamService(repo.Team),
		Participant:   NewParticipantService(repo.Participant),
		Comparison:    NewComparisonService(sber.New()),
	}
}

type Authorization interface {
	CreateUser(ctx context.Context, user entities.User) (int, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (int, error)
	GetUserById(ctx context.Context, id int) (entities.User, error)
}

type Team interface {
	CreateTeam(ctx context.Context, t entities.Team) (int, error)
	GetAllTeams(ctx context.Context) ([]entities.Team, error)
	GetTeamById(ctx context.Context, id int) (entities.Team, error)
	GetTeamsByUserId(ctx context.Context, userId int) ([]entities.Team, error)
	UpdateTeam(ctx context.Context, t entities.Team) error
	DeleteTeamById(ctx context.Context, id int) error
	IsUserTeamCreator(ctx context.Context, userId, teamId int) (bool, error)
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

type Comparison interface {
	GetComparisonBetween2WithDescription(date1 time.Time, date2 time.Time) (entities.ComparisonResult, error)
	GetTeamComparison(part []entities.Participant, birthDate time.Time) ([]int, error)
}
