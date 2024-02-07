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
