package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"fmt"
)

func AllPlaylists(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM playlists"
	Handler(query, w, r, &[]postgres.Playlist{})
}

func CreatePlaylist(w Http.Response, r *Http.Request) {
	var playlist postgres.Playlist

	err := json.Unmarshal([]byte(r.Body), &playlist)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		fmt.Println(Http.GetStatusText(Http.StatusBadRequest))
		return
	}

	query := "INSERT INTO playlists (name,user_id) VALUES ($1,$2)"

	Handler(query, w, r, &postgres.Playlist{}, playlist.Name, playlist.UserID)
}
