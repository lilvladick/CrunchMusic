package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllAuthors(w Http.Response, r *Http.Request) {
	tracks, err := postgres.GetAuthors()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func GetAuthorById(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	AuthorID := requestBody.ID

	Author, err := postgres.GetAuthorsByID(AuthorID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte(Http.GetStatusText(Http.StatusNotFound)))
		return
	}

	jsonAuthor, err := json.Marshal(Author)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAuthor)
}

func GetAuthorByEmail(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Email string `json:"Email"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	Email := requestBody.Email

	Author, err := postgres.QueryAuthors(Email)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte(Http.GetStatusText(Http.StatusNotFound)))
		return
	}

	jsonAuthor, err := json.Marshal(Author)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAuthor)
}

func GetAuthorByName(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	Email := requestBody.Name

	Author, err := postgres.GetAuthorByName(Email)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte(Http.GetStatusText(Http.StatusNotFound)))
		return
	}

	jsonAuthor, err := json.Marshal(Author)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAuthor)
}

func AddAuthor(w Http.Response, r *Http.Request) {
	var author postgres.Author

	err := json.Unmarshal([]byte(r.Body), &author)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	err = postgres.InsertAuthor(author)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte(Http.GetStatusText(Http.StatusCreated)))
}
