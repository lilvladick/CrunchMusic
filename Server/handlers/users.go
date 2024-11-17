package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"fmt"
	"log"
)

func AllUsers(w Http.Response, r *Http.Request) {
	query := "SELECT * FROM users"
	Handler(query, w, r, &[]postgres.User{})
}

func GetUserById(w Http.Response, r *Http.Request) {
	var user postgres.User

	err := json.Unmarshal([]byte(r.Body), &user)
	if err != nil {
		log.Print("err: %v", err)
		w.WriteHeader(Http.StatusBadRequest)
		fmt.Println(Http.GetStatusText(Http.StatusBadRequest))
		return
	}
	query := "SELECT * FROM users WHERE id=$1"
	Handler(query, w, r, &[]postgres.User{})
}
