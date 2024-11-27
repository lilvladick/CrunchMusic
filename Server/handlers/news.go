package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllNews(w Http.Response, r *Http.Request) {
	tracks, err := postgres.GetNews()
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	jsonTracks, err := json.Marshal(tracks)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTracks)
}

func GetNewsById(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	ID := requestBody.ID

	Author, err := postgres.GetNewsByID(ID)
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

func GetNewsByAuthorId(w Http.Response, r *Http.Request) {
	var requestBody struct {
		AuthorID int `json:"author_id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	AuthorID := requestBody.AuthorID

	Author, err := postgres.GetNewsByAuthorID(AuthorID)
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

func GetNewsByCategoryId(w Http.Response, r *Http.Request) {
	var requestBody struct {
		CategoryID int `json:"category_id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	CategoryID := requestBody.CategoryID

	Author, err := postgres.GetNewsByCategoryID(CategoryID)
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

func GetNewsByTitle(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Title string `json:"title"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	Title := requestBody.Title

	Author, err := postgres.GetNewsByTitle(Title)
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

func WriteNews(w Http.Response, r *Http.Request) {
	var news postgres.News

	err := json.Unmarshal([]byte(r.Body), &news)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	err = postgres.InsertNews(news)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte(Http.GetStatusText(Http.StatusCreated)))
}
