package handlers

import (
	service "CrunchServer/Service"
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func Home(w http.ResponseWriter, r *http.Request) {
	tracks, err := service.FetchAllTracks(db)
	if err != nil {
		http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Tracks</h1>")
	fmt.Fprintf(w, "<ul>")
	for _, track := range tracks {
		fmt.Fprintf(w, "<li>ID: %s - Title: %s - Genre: %s - Duration: %v</li>", track.ID, track.Title, track.Genre, track.Duration)
	}
	fmt.Fprintf(w, "</ul>")
}
