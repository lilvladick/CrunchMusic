package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"database/sql"
	"encoding/json"
	"fmt"
)

var db *sql.DB

func AllTracks(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM tracks"
	Handler(query, w, r)
}

func Home(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM tracks LIMIT 100;"
	Handler(query, w, r)
}

func AddTracks(w Http.Response, r *Http.Request) {
	query := "INSERT INTO tracks VALUES ($1,$2,$3,$4,$5);"
	Handler(query, w, r)
}

func TracksFromPlaylist(w Http.Response, r *Http.Request) {
	query := "SELECT t.id, t.title, t.filepath, t.user_id, t.genre, t.duration FROM tracks t JOIN playlist_tracks pt ON t.id = pt.track_id WHERE pt.playlist_id = 1; "
	Handler(query, w, r)
}

func HandleAddTrack(w Http.Response, r *Http.Request) {
	var track postgres.Track
	err := json.Unmarshal([]byte(r.Body), &track)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		fmt.Println(Http.GetStatusText(Http.StatusBadRequest))
		return
	}
	query := "INSERT INTO tracks (title, filepath, user_id, genre, duration) VALUES ('Track 45', '/Server/music/track3.mp3', 1, 'Rap', '00:02:50');"
	_, err = db.Exec(query, track.Title, track.Filepath, track.UserID, track.Genre, track.Duration)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		fmt.Println(Http.GetStatusText(Http.StatusInternalServerError))
		return
	}
	w.WriteHeader(Http.StatusCreated)
	fmt.Println(Http.GetStatusText(Http.StatusCreated))
}
