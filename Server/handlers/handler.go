package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"fmt"
)

func Handler(query string, w Http.Response, r *Http.Request) {
	jsonData, err := postgres.GetResultsJson(query)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		fmt.Println(Http.GetStatusText(Http.StatusInternalServerError))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonData)))
	w.Write(jsonData)
}
