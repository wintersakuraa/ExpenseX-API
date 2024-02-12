package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wintersakuraa/expense-x-api/internal/domain"
	"github.com/wintersakuraa/expense-x-api/internal/repositories/pg"
)

type User interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id string) (domain.User, error)
	Update(ctx context.Context, id string, input domain.UpdateUserInput) error
	Delete(ctx context.Context, id string) error
}

type Currency interface {
	Create(ctx context.Context, currency domain.Currency) error
	GetAll(ctx context.Context) ([]domain.Currency, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Currency, error)
	Update(ctx context.Context, id uuid.UUID, input domain.UpdateCurrencyInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Category interface {
	Create(ctx context.Context, category domain.Category) error
	GetAll(ctx context.Context) ([]domain.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Category, error)
	Update(ctx context.Context, id uuid.UUID, input domain.UpdateCategoryInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Transaction interface {
	Create(ctx context.Context, transaction domain.Transaction) error
	GetAll(ctx context.Context) ([]domain.Transaction, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Transaction, error)
	Update(ctx context.Context, id uuid.UUID, input domain.UpdateTransactionInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repositories struct {
	User        User
	Currency    Currency
	Category    Category
	Transaction Transaction
}

func New(db *sqlx.DB) *Repositories {
	return &Repositories{
		User:        pg.NewUserRepository(db),
		Currency:    pg.NewCurrencyRepository(db),
		Category:    pg.NewCategoryRepository(db),
		Transaction: pg.NewTransactionRepository(db),
	}
}
