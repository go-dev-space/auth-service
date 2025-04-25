package store

import (
	"context"
	"errors"

	"github.com/auth-service/internal/auth/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreServicePostgresql struct {
	DB *pgxpool.Pool
}

func NewStoreServicePostrgresql(db *pgxpool.Pool) *StoreServicePostgresql {
	return &StoreServicePostgresql{
		DB: db,
	}
}

func (s *StoreServicePostgresql) Save(ctx context.Context, u *domain.User) error {

	tx, err := s.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var userID int
	err = tx.QueryRow(ctx, `INSERT INTO users 
	(username, email, password) VALUES ($1,$2,$3) RETURNING id`, u.Username, u.Email, u.Password).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "users_email_key":
				return errors.New("email already exists")
			case "users_username_key":
				return errors.New("username already exists")
			}
		}
		return err
	}

	_, err = tx.Exec(ctx, `INSERT INTO profiles (user_id) VALUES ($1)`, userID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
