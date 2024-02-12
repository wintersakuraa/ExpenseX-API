package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt  time.Time `json:"createdAt"   db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt"   db:"updated_at"`
	Currency   *Currency `json:"currency"`
	ID         string    `json:"id"          db:"id"`
	CurrencyID uuid.UUID `json:"currency_id" db:"currency_id"`
}

func NewUser(id string, currencyId uuid.UUID) User {
	return User{
		ID:         id,
		CurrencyID: currencyId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type UpdateUserInput struct {
	CurrencyID uuid.UUID
}
