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

type TransactionParams struct {
	Date       time.Time
	Type       Type
	Note       string
	UserID     string
	Currency   Currency
	Amount     float64
	CategoryID uuid.UUID
}

func NewTransaction(params TransactionParams) Transaction {
	return Transaction{
		ID:         uuid.New(),
		CategoryID: params.CategoryID,
		UserID:     params.UserID,
		Amount:     params.Amount,
		Currency:   params.Currency,
		Note:       params.Note,
		Date:       params.Date,
		Type:       params.Type,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type UpdateTransactionInput struct {
	CategoryID *uuid.UUID
	Amount     *float64
	Currency   *Currency
	Note       *string
	Date       *time.Time
	Type       *Type
}
