package service

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
	"log/slog"
)

type TeamService struct {
	repo repo.Team
}

func NewTeamService(repo repo.Team) *TeamService {
	return &TeamService{
		repo: repo,
	}
}

func (s *TeamService) CreateTeam(ctx context.Context, t entities.Team) (int, error) {
	id, err := s.repo.CreateTeam(ctx, t)
	if err != nil {
		slog.Error(err.Error())
	}
	return id, err
}

func (s *TeamService) GetAllTeams(ctx context.Context) ([]entities.Team, error) {
	teams, err := s.repo.GetAllTeams(ctx)
	if err != nil {
		slog.Error(err.Error())
	}
	return teams, err
}

func (s *TeamService) GetTeamById(ctx context.Context, id int) (entities.Team, error) {
	team, err := s.repo.GetTeamById(ctx, id)
	if err != nil {
		slog.Error(err.Error())
	}
	return team, err
}

func (s *TeamService) GetTeamsByUserId(ctx context.Context, userId int) ([]entities.Team, error) {
	teams, err := s.repo.GetTeamsByUserId(ctx, userId)
	if err != nil {
		slog.Error(err.Error())
	}
	return teams, err
}

func (s *TeamService) UpdateTeam(ctx context.Context, t entities.Team) error {
	err := s.repo.UpdateTeam(ctx, t)
	if err != nil {
		slog.Error(err.Error())
	}
	return err
}
