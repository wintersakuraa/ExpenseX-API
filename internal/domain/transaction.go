package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Date       time.Time `json:"date"       db:"date"`
	CreatedAt  time.Time `json:"createdAt"  db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt"  db:"updated_at"`
	UserID     string    `json:"userId"     db:"user_id"`
	Type       Type      `json:"type"       db:"type"`
	Note       string    `json:"note"       db:"note"`
	Currency   Currency  `json:"currency"`
	Amount     float64   `json:"amount"     db:"amount"`
	ID         uuid.UUID `json:"id"         db:"id"`
	CategoryID uuid.UUID `json:"categoryId" db:"category_id"`
}
