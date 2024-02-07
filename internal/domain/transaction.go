package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UserID     string    `json:"userId"`
	Type       Type      `json:"type"`
	Note       string    `json:"note"`
	Currency   Currency  `json:"currency"`
	Amount     float64   `json:"amount"`
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `json:"categoryId"`
}
