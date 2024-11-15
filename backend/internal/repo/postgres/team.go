package postgres

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

type TeamRepo struct {
	Db *sqlx.DB
}

func NewTeamRepo(db *sqlx.DB) *TeamRepo {
	return &TeamRepo{
		Db: db,
	}
}

func (r TeamRepo) CreateTeam(ctx context.Context, t entities.Team) (int, error) {
	var id int
	query := "INSERT INTO teams (name, user_id, description) VALUES ($1, $2, $3) RETURNING id"
	row := r.Db.QueryRowContext(ctx, query, t.Name, t.UserID, t.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r TeamRepo) GetAllTeams(ctx context.Context) ([]entities.Team, error) {
	var teams []entities.Team
	query := "SELECT * FROM teams"
	err := r.Db.SelectContext(ctx, &teams, query)
	return teams, err
}

func (r TeamRepo) GetTeamById(ctx context.Context, id int) (entities.Team, error) {
	var team entities.Team
	query := "SELECT * FROM teams WHERE id = $1"
	err := r.Db.GetContext(ctx, &team, query, id)
	return team, err
}

func (r TeamRepo) GetTeamsByUserId(ctx context.Context, userId int) ([]entities.Team, error) {
	var teams []entities.Team
	query := "SELECT * FROM teams WHERE user_id = $1"
	err := r.Db.SelectContext(ctx, &teams, query, userId)
	return teams, err
}

func (r TeamRepo) UpdateTeam(ctx context.Context, t entities.Team) error {
	query := "UPDATE teams SET name = $1, description = $2 WHERE id = $3"
	_, err := r.Db.ExecContext(ctx, query, t.Name, t.Description, t.Id)
	return err
}

func (r TeamRepo) DeleteTeamById(ctx context.Context, id int) error {
	query := "DELETE FROM teams WHERE id = $1"
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
