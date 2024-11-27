package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func InsertNews(news News) error {
	query := `INSERT INTO news (title, news_content, author_id, category_id) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, news.Title, news.NewsContent, news.AuthorID, news.CategoryID)
	return err
}

func GetNewsByID(id int) (*News, error) {
	query := "SELECT * FROM news WHERE id = $1"

	row := db.QueryRow(query, id)
	var news News

	err := row.Scan(&news.ID, &news.Title, &news.NewsContent, &news.AuthorID, &news.CategoryID, &news.PublishedAt, &news.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("News not found with ID %v", id)
		}
		return nil, fmt.Errorf("error scanning News row: %v", err)
	}

	return &news, nil
}

func GetNewsByTitle(title string) ([]News, error) {
	query := "SELECT * FROM news WHERE title = $1"

	rows, err := db.Query(query, title)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var newss []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.NewsContent, &news.AuthorID, &news.CategoryID, &news.PublishedAt, &news.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning News row: %v", err)
		}
		newss = append(newss, news)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over news rows: %v", err)
	}

	return newss, nil
}

func GetNewsByAuthorID(id int) ([]News, error) {
	query := `SELECT * FROM news WHERE author_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var newss []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.NewsContent, &news.AuthorID, &news.CategoryID, &news.PublishedAt, &news.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		newss = append(newss, news)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments rows: %v", err)
	}

	return newss, nil
}

func GetNewsByCategoryID(id int) ([]News, error) {
	query := `SELECT * FROM news WHERE author_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var newss []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.NewsContent, &news.AuthorID, &news.CategoryID, &news.PublishedAt, &news.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		newss = append(newss, news)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments rows: %v", err)
	}

	return newss, nil
}

func GetNews() ([]News, error) {
	query := "SELECT * FROM news"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var newss []News
	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.NewsContent, &news.AuthorID, &news.CategoryID, &news.PublishedAt, &news.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		newss = append(newss, news)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments rows: %v", err)
	}

	return newss, nil
}
