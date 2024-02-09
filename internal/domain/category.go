package domain

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	Expense Type = "expense"
	Income  Type = "income"
)

type Category struct {
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	UserID    string    `json:"userId"    db:"user_id"`
	Name      string    `json:"name"      db:"name"`
	Type      Type      `json:"type"      db:"type"`
	ID        uuid.UUID `json:"id"        db:"id"`
}

type CategoryParams struct {
	Type   Type
	UserID string
	Name   string
}

func NewCategory(params CategoryParams) Category {
	return Category{
		ID:        uuid.New(),
		UserID:    params.UserID,
		Name:      params.Name,
		Type:      params.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UpdateCategoryInput struct {
	Name *string
	Type *Type
}
