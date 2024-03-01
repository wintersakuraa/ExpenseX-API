package pg

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wintersakuraa/expense-x-api/internal/domain"
)

type currencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *currencyRepository {
	return &currencyRepository{
		db: db,
	}
}

func (r *currencyRepository) Create(ctx context.Context, currency domain.Currency) error {
	query := `INSERT INTO currencies (
    id, 
    name,
    code,
    created_at,
    updated_at
  ) VALUES (
    :id, 
    :name,
    :code,
    :created_at,
    :updated_at
  )`

	_, err := r.db.NamedExecContext(ctx, query, currency)
	if err != nil {
		return err
	}

	return nil
}

func (r *currencyRepository) GetAll(ctx context.Context) ([]domain.Currency, error) {
	query := `SELECT * FROM currencies ORDER BY code`
	var currencies []domain.Currency

	err := r.db.SelectContext(ctx, &currencies, query)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r *currencyRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Currency, error) {
	query := `SELECT * FROM currencies WHERE id = $1`
	var currency domain.Currency
	err := r.db.GetContext(ctx, &currency, query, id)

	return currency, err
}

func (r *currencyRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	input domain.UpdateCurrencyInput,
) error {
	updates := newPGUpdates()

	if input.Code != nil {
		updates.add("code", input.Code)
	}

	if input.Name != nil {
		updates.add("name", input.Name)
	}

	updates.appendArg(id)
	query := fmt.Sprintf(
		"UPDATE currencies SET %s WHERE id = $%d",
		updates.getSetQuery(),
		updates.argId,
	)

	_, err := r.db.ExecContext(ctx, query, updates.args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *currencyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE from currencies WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
