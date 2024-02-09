package domain

import "time"

type User struct {
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	ID        string    `json:"id"        db:"id"`
	Currency  Currency  `json:"currency"`
}

func NewUser(id string, currency Currency) User {
	return User{
		ID:        id,
		Currency:  currency,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type UpdateUserInput struct {
	Currency Currency
}
