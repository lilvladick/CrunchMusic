package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"fmt"
	"log"
)

func AllTracksFromPlaylists(w Http.Response, r *Http.Request) {
	playlists_tracks, err := postgres.GetPlaylistTracks()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonPlaylists, err := json.Marshal(playlists_tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPlaylists)
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

func PlaylistsTracksByTrackID(w Http.Response, r *Http.Request) {
	var requestBody struct {
		TrackID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	TrackID := requestBody.TrackID

	tracks, err := postgres.GetPlaylistTracksByTrackID(TrackID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func PlaylistsTracksByPlaylistID(w Http.Response, r *Http.Request) {
	var requestBody struct {
		PlaylistID int `json:"playlist_id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	playlistID := requestBody.PlaylistID

	tracks, err := postgres.GetPlaylistTracksByPlaylistID(playlistID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}
