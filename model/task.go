package model

import "time"

type Task struct {
	ID          int        `json:"id" db:"id"`
	UserID      int        `json:"user_id" db:"user_id"`
	Summary     string     `json:"summary" db:"summary"`
	CompletedAt *time.Time `json:"completed_at,omitempty" db:"completed_at"`
}
