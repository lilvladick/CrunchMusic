package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/models"
	"CrunchServer/postgres"
	"encoding/json"
	"fmt"
)

func AllTracks(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM tracks"
	Handler(query, w, r, &[]postgres.Track{})
}

func Home(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM tracks LIMIT 100;"
	Handler(query, w, r, &[]postgres.Track{})
}

func TracksFromPlaylist(w Http.Response, r *Http.Request) {
	query := "SELECT t.* FROM tracks t JOIN playlist_tracks pt ON t.id = pt.track_id WHERE pt.playlist_id = $1; "
	Handler(query, w, r, &[]postgres.Track{})
}

func HandleAddTrack(w Http.Response, r *Http.Request) {
	var track models.Track

	err := json.Unmarshal([]byte(r.Body), &track)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		fmt.Println(Http.GetStatusText(Http.StatusBadRequest))
		return
	}
	tduration := models.Timestamp{Time: track.Duration}
	pgTrack := postgres.Track{
		Title:    track.Title,
		Filepath: track.Filepath,
		UserID:   track.UserID,
		Genre:    track.Genre,
		Duration: tduration.ToNullTime(),
	}

	query := "INSERT INTO tracks (title, filepath, user_id, genre, duration) VALUES ($1, $2, $3, $4, $5)"

	Handler(query, w, r, &postgres.Track{}, pgTrack.Title, pgTrack.Filepath, pgTrack.UserID, pgTrack.Genre, pgTrack.Duration)
}
