package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"fmt"
)

func AllTracks(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM tracks"
	jsonData, err := postgres.GetResultsJson(query)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		fmt.Println("Error fetching tracks")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonData)))
	w.Write(jsonData)
}
