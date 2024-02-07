package domain

import "time"

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
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
