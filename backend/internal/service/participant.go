package service

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/repo"
)

type ParticipantService struct {
	repo repo.Participant
}

func NewParticipantService(repo repo.Participant) *ParticipantService {
	return &ParticipantService{
		repo: repo,
	}
}

func (s *ParticipantService) CreateParticipant(ctx context.Context, p entities.Participant) (int, error) {
	id, err := s.repo.CreateParticipant(ctx, p)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ParticipantService) GetAllParticipants(ctx context.Context) ([]entities.Participant, error) {
	participants, err := s.repo.GetAllParticipants(ctx)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (s *ParticipantService) GetParticipantById(ctx context.Context, id int) (entities.Participant, error) {
	participant, err := s.repo.GetParticipantById(ctx, id)
	if err != nil {
		return entities.Participant{}, err
	}
	return participant, nil
}

func (s *ParticipantService) GetParticipantsByTeamId(ctx context.Context, teamId int) ([]entities.Participant, error) {
	participants, err := s.repo.GetParticipantsByTeamId(ctx, teamId)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (s *ParticipantService) UpdateParticipant(ctx context.Context, p entities.Participant) error {
	err := s.repo.UpdateParticipant(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (s *ParticipantService) DeleteParticipantById(ctx context.Context, id int) error {
	err := s.repo.DeleteParticipantById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ParticipantService) GetParticipantByCreatorId(ctx context.Context, creatorId int) ([]entities.Participant, error) {
	participants, err := s.repo.GetParticipantByCreatorId(ctx, creatorId)
	if err != nil {
		return nil, err
	}
	return participants, nil
}
