package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllTracksFromPlaylists(w Http.Response, r *Http.Request) {
	playlists_tracks, err := postgres.GetPlaylistTracks()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	jsonPlaylists, err := json.Marshal(playlists_tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPlaylists)
}

func PlaylistsTracksByTrackID(w Http.Response, r *Http.Request) {
	var requestBody struct {
		TrackID int `json:"track_id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	TrackID := requestBody.TrackID

	tracks, err := postgres.GetPlaylistTracksByTrackID(TrackID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
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
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	playlistID := requestBody.PlaylistID

	tracks, err := postgres.GetPlaylistTracksByPlaylistID(playlistID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func AddTrackToPlaylist(w Http.Response, r *Http.Request) {
	var playlist_track postgres.Playlist_tracks

	err := json.Unmarshal([]byte(r.Body), &playlist_track)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	err = postgres.AddTrackToPlaylist(playlist_track.PlaylistID, playlist_track.TrackID, playlist_track.AddedAt)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte(Http.GetStatusText(Http.StatusCreated)))
}
