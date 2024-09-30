package main

import "github.com/lib/pq"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

type Track struct {
	ID       string      `json:"id"`
	Title    string      `json:"title"`
	Filepath string      `json:"filepath"`
	UserID   string      `json:"user_id"`
	Genre    string      `json:"genre"`
	Duration pq.NullTime `json:"duration"`
}

type Likes struct {
	UserID  string `json:"user_id"`
	TrackID string `json:"track_id"`
}