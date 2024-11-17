package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"fmt"
	"log"
)

func AllTracksFromPlaylists(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM playlist_tracks"
	Handler(query, w, r, &[]postgres.Playlist_tracks{})
}

func AddTrackToPlaylist(w Http.Response, r *Http.Request) {
	var playlist_tracks postgres.Playlist_tracks

	err := json.Unmarshal([]byte(r.Body), &playlist_tracks)
	if err != nil {
		log.Print("err: %v", err)
		w.WriteHeader(Http.StatusBadRequest)
		fmt.Println(Http.GetStatusText(Http.StatusBadRequest))
		return
	}

	query := "INSERT INTO playlist_tracks (playlist_id, track_id, added_at) VALUES ($1,$2,$3)"

	Handler(query, w, r, &postgres.Playlist_tracks{}, playlist_tracks.PlaylistID, playlist_tracks.TrackID, playlist_tracks.AddedAt)
}
