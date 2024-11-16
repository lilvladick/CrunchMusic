package handlers

import (
	Http "CrunchServer/http"
)

func AllPlaylists(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM playlists"
	Handler(query, w, r)
}

func AllTracksFromPlaylists(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM playlist_tracks"
	Handler(query, w, r)
}
