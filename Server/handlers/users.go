package handlers

import (
	Http "CrunchServer/http"
)

func AllUsers(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM users"
	Handler(query, w, r)
}
