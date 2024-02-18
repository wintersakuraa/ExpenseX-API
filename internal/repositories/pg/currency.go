package pg

import (
	"context"

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

	err := r.db.Select(&currencies, query)
	if err != nil {
		return []domain.Currency{}, err
	}

	return currencies, nil
}

func (r *currencyRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Currency, error) {
	return domain.Currency{}, nil
}

func (r *currencyRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	input domain.UpdateCurrencyInput,
) error {
	return nil
}

func (r *currencyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
