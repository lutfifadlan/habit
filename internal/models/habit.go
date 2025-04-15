package models

import "time"

type Habit struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Habit           string    `json:"habit"`
	CompletionDates []string  `json:"completion_dates"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateHabitRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Habit  string `json:"habit" validate:"required"`
}
