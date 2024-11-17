package models

import (
	"time"
)

type Track struct {
	ID       int       `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Filepath string    `json:"filepath,omitempty"`
	UserID   int       `json:"user_id,omitempty"`
	Genre    string    `json:"genre,omitempty"`
	Duration time.Time `json:"duration,omitempty"`
}
