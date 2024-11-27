package handlers

import (
	Http "CrunchServer/http"
	"CrunchServer/postgres"
	"encoding/json"
	"log"
)

func AllCategories(w Http.Response, r *Http.Request) {
	tracks, err := postgres.GetCategories()
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

func GetCategoryById(w Http.Response, r *Http.Request) {
	var requestBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	CategoryID := requestBody.ID

	Category, err := postgres.GetCategoriesByID(CategoryID)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte(Http.GetStatusText(Http.StatusNotFound)))
		return
	}

	jsonCategory, err := json.Marshal(Category)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCategory)
}

func GetCategoryByName(w Http.Response, r *Http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	Name := requestBody.Name

	Category, err := postgres.GetCategoryByName(Name)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusNotFound)
		w.Write([]byte(Http.GetStatusText(Http.StatusNotFound)))
		return
	}

	jsonCategory, err := json.Marshal(Category)
	if err != nil {
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCategory)
}

func AddCategory(w Http.Response, r *Http.Request) {
	var category postgres.Category

	err := json.Unmarshal([]byte(r.Body), &category)
	if err != nil {
		w.WriteHeader(Http.StatusBadRequest)
		w.Write([]byte(Http.GetStatusText(Http.StatusBadRequest)))
		return
	}

	err = postgres.InsertCategory(category)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(Http.StatusInternalServerError)
		w.Write([]byte(Http.GetStatusText(Http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(Http.StatusCreated)
	w.Write([]byte(Http.GetStatusText(Http.StatusCreated)))
}
