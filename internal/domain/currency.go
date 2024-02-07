package domain

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	ID        uuid.UUID `json:"id"`
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
