package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllTracks(w Http.Response, r *Http.Request) {
	tracks, err := postgres.GetTracks()
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

func Home(w Http.Response, r *Http.Request) {
	tracks, err := postgres.Get100Tracks()
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

func TracksFromPlaylist(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	playlistID := requestBody.ID

	tracks, err := postgres.GetTracksFromPlaylist(playlistID)
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

func TracksByGenre(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Genre string `json:"genre"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	Genre := requestBody.Genre

	tracks, err := postgres.GetTrackByGenre(Genre)
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

func TracksByTitle(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Title string `json:"title"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	Title := requestBody.Title

	tracks, err := postgres.GetTrackByTitle(Title)
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

func GetTrackById(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	trackID := requestBody.ID

	user, err := postgres.GetTrackByID(trackID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}

func HandleAddTrack(w Http.Response, r *Http.Request) {
	var track postgres.Track

	err := json.Unmarshal([]byte(r.Body), &track)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	err = postgres.UploadTrack(track.Title, track.Filepath, track.UserID, track.Genre, track.Duration)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte("Track created successfully"))
}
