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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	Type      Type      `json:"type"`
	ID        uuid.UUID `json:"id"`
}
