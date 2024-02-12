package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wintersakuraa/expense-x-api/internal/domain"
)

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Create(ctx context.Context, transaction domain.Transaction) error {
	return nil
}

func (r *transactionRepository) GetAll(ctx context.Context) ([]domain.Transaction, error) {
	return []domain.Transaction{}, nil
}

func (r *transactionRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (r *transactionRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	input domain.UpdateTransactionInput,
) error {
	return nil
}

func (r *transactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
