package models

import "time"

type Catheter struct {
	ID        string     `json:"id" db:"id"`
	UserID    string     `json:"userId" db:"user_id"`
	StartedAt time.Time  `json:"startedAt" db:"started_at"`
	EndedAt   *time.Time `json:"endedAt" db:"ended_at"`
}
