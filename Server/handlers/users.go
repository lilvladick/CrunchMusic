package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

type ctxKey struct{}

func AllUsers(w Http.Response, r *Http.Request) {
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

func GetUserById(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	userID := requestBody.ID

	user, err := postgres.GetUsersByID(userID)
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

func GetUserBylogin(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Login string `json:"login"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	login := requestBody.Login

	user, err := postgres.QueryUsers(login)
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

func GetUserByName(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	login := requestBody.Name

	user, err := postgres.GetUsrByName(login)
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
