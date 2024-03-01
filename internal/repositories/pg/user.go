package pg

import (
	"context"
	"fmt"
	"strings"

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
	query := `SELECT * FROM users WHERE id = $1`
	var user domain.User
	err := r.db.GetContext(ctx, &user, query, id)

	return user, err
}

func (r *userRepository) Update(
	ctx context.Context,
	id string,
	input domain.UpdateUserInput,
) error {
	setVals := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.CurrencyID != nil {
		args = append(args, input.CurrencyID)
		setVals = append(setVals, fmt.Sprintf("currency_id = $%d", argId))
		argId++
	}

	setQuery := strings.Join(setVals, ", ")
	args = append(args, id)
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", setQuery, argId)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}
