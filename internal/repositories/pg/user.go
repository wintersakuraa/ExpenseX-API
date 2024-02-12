package pg

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/wintersakuraa/expense-x-api/internal/domain"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (
    id, 
    currency_id,
    created_at,
    updated_at
  ) VALUES (
    :id, 
    :currency_id,
    :created_at,
    :updated_at
  )`

	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	id string,
	input domain.UpdateUserInput,
) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}
