package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wintersakuraa/expense-x-api/internal/domain"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category domain.Category) error {
	return nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	return []domain.Category{}, nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	return domain.Category{}, nil
}

func (r *categoryRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	input domain.UpdateCategoryInput,
) error {
	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
