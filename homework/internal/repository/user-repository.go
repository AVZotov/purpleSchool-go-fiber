package repository

import (
	"context"
	"news/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{}

	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, username, email, created_at, updated_at
	`

	if err = r.db.QueryRow(ctx, query, username, email, string(hashedPassword)).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, username, email, password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var hashedPassword string
	if err := r.db.QueryRow(ctx, query, email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&hashedPassword,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
