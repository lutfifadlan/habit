package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	UserName string `json:"user_name" validate:"required"`
}
