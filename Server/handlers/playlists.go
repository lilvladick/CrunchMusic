package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllPlaylists(w Http.Response, r *Http.Request) {
	playlists, err := postgres.GetPlaylists()
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonPlaylists, err := json.Marshal(playlists)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPlaylists)
}

func CreatePlaylist(w Http.Response, r *Http.Request) {
	var playlist postgres.Playlist

	err := json.Unmarshal([]byte(r.Body), &playlist)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	err = postgres.CreatePlaylist(playlist.Name, playlist.UserID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte("Playlist created successfully"))
}

func PlaylistsByName(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	Name := requestBody.Name

	tracks, err := postgres.GetPlaylistByName(Name)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func PlaylistsByUserID(w Http.Response, r *Http.Request) {
	var requestBody struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	UserID := requestBody.UserID

	tracks, err := postgres.GetPlaylistByUserID(UserID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func PlaylistsByID(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	ID := requestBody.ID

	tracks, err := postgres.GetPlaylistByID(ID)
	if err != nil {
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
