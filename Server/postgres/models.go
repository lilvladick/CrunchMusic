package postgres

import (
	"github.com/lib/pq"
)

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
	UserID   int         `json:"user_id"`
	Genre    string      `json:"genre"`
	Duration pq.NullTime `json:"duration"`
}

type Likes struct {
	UserID  int `json:"user_id"`
	TrackID int `json:"track_id"`
}

type Playlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type Playlist_tracks struct {
	PlaylistID int         `json:"playlist_id"`
	TrackID    int         `json:"track_id"`
	AddedAt    pq.NullTime `json:"added_at"`
}
