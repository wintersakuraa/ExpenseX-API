package domain

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Name      string    `json:"name"      db:"name"`
	Code      string    `json:"code"      db:"code"`
	ID        uuid.UUID `json:"id"        db:"id"`
}

func NewCurrency(name string, code string) Currency {
	return Currency{
		ID:        uuid.New(),
		Name:      name,
		Code:      code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UpdateCurrencyInput struct {
	Name *string
	Code *string
}
