package postgres

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ParticipantRepo struct {
	Db *sqlx.DB
}

func NewParticipantRepo(db *sqlx.DB) *ParticipantRepo {
	return &ParticipantRepo{
		Db: db,
	}
}

func (r ParticipantRepo) CreateParticipant(ctx context.Context, p entities.Participant) (int, error) {
	var id int
	query := "INSERT INTO participants (team_id, name, role, birthdate, birthplace) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := r.Db.QueryRowContext(ctx, query, p.TeamID, p.Name, p.Role, p.BirthDate, p.BirthPlace)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r ParticipantRepo) GetAllParticipants(ctx context.Context) ([]entities.Participant, error) {
	var participants []entities.Participant
	query := "SELECT * FROM participants"
	err := r.Db.SelectContext(ctx, &participants, query)
	return participants, err
}

func (r ParticipantRepo) GetParticipantById(ctx context.Context, id int) (entities.Participant, error) {
	var participant entities.Participant
	query := "SELECT * FROM participants WHERE id = $1"
	err := r.Db.GetContext(ctx, &participant, query, id)
	return participant, err
}

func (r ParticipantRepo) GetParticipantsByTeamId(ctx context.Context, teamId int) ([]entities.Participant, error) {
	var participants []entities.Participant
	query := "SELECT * FROM participants WHERE team_id = $1"
	err := r.Db.SelectContext(ctx, &participants, query, teamId)
	return participants, err
}

func (r ParticipantRepo) UpdateParticipant(ctx context.Context, p entities.Participant) error {
	query := "UPDATE participants SET name = $1, role = $2, birthdate = $3, birthplace = $4 WHERE id = $5"
	_, err := r.Db.ExecContext(ctx, query, p.Name, p.Role, p.BirthDate, p.BirthPlace, p.Id)
	return err
}

func (r ParticipantRepo) DeleteParticipantById(ctx context.Context, id int) error {
	query := "DELETE FROM participants WHERE id = $1"
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
