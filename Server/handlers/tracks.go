package handlers

import (
	"CrunchServer/postgres"
	"fmt"
	"net/http"
)

func AllTracks(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM tracks"
	jsonData, err := postgres.GetResultsJson(query)
	if err != nil {
		http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonData)))
	fmt.Fprint(w, string(jsonData))
}
