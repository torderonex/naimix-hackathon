package postgres

import (
	"context"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	Db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		Db: db,
	}
}

func (r AuthRepo) CreateUser(ctx context.Context, u entities.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	row := r.Db.QueryRowContext(ctx, query, u.Username, u.Password, u.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

func (r AuthRepo) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	query := "SELECT * FROM users"
	err := r.Db.SelectContext(ctx, &users, query)
	return users, err
}

func (r AuthRepo) GetUserByCredentials(ctx context.Context, email, password string) (entities.User, error) {
	var user entities.User
	query := "SELECT * FROM users WHERE email = $1 AND password = $2"
	err := r.Db.GetContext(ctx, &user, query, email, password)
	return user, err
}

func (r AuthRepo) GetUserById(ctx context.Context, id int) (entities.User, error) {
	var user entities.User
	query := "SELECT * FROM users WHERE id = $1"
	err := r.Db.GetContext(ctx, &user, query, id)
	return user, err
}

func (r AuthRepo) DeleteUserById(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.Db.ExecContext(ctx, query, id)
	return err
}
