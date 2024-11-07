package postgres

import "github.com/lib/pq"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

type Track struct {
	ID       int         `json:"id"`
	Title    string      `json:"title"`
	Filepath string      `json:"filepath"`
	Artist   string      `json:"artist"`
	Genre    string      `json:"genre"`
	Duration pq.NullTime `json:"duration"`
}

type Likes struct {
	UserID  int `json:"user_id"`
	TrackID int `json:"track_id"`
}
