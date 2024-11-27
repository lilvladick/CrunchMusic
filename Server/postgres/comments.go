package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func InsertNewsComment(comment NewsComment) error {
	query := `INSERT INTO news_comments (news_id, author_id, comment_content) VALUES ($1, $2, $3) `
	_, err := db.Exec(query, comment.NewsID, comment.AuthorID, comment.CommentContent)
	return err
}

func GetNewsCommentByID(id int) (*NewsComment, error) {
	query := `SELECT * FROM news_comments WHERE id = $1`

	row := db.QueryRow(query, id)
	var comment NewsComment

	err := row.Scan(&comment.ID, &comment.NewsID, &comment.AuthorID, &comment.CommentContent, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Comment not found with ID %v", id)
		}
		return nil, fmt.Errorf("error scanning Comments row: %v", err)
	}

	return &comment, nil
}

func GetNewsCommentByAuthorID(id int) ([]NewsComment, error) {
	query := `SELECT * FROM news_comments WHERE author_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var news_comments []NewsComment
	for rows.Next() {
		var comment NewsComment
		err := rows.Scan(&comment.ID, &comment.NewsID, &comment.AuthorID, &comment.CommentContent, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		news_comments = append(news_comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments rows: %v", err)
	}

	return news_comments, nil
}

func GetNewsCommentByNewsID(id int) ([]NewsComment, error) {
	query := `SELECT * FROM news_comments WHERE news_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %v", err)
	}
	defer rows.Close()

	var news_comments []NewsComment
	for rows.Next() {
		var comment NewsComment
		err := rows.Scan(&comment.ID, &comment.NewsID, &comment.AuthorID, &comment.CommentContent, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		news_comments = append(news_comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments rows: %v", err)
	}

	return news_comments, nil
}

func GetNewsComments() ([]NewsComment, error) {
	sqlStatement := `
        SELECT * FROM news_comments;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying Comments: %v", err)
	}
	defer rows.Close()

	var NewsComments []NewsComment
	for rows.Next() {
		var NewsComment NewsComment
		err := rows.Scan(&NewsComment.ID, &NewsComment.NewsID, &NewsComment.AuthorID, &NewsComment.CommentContent, &NewsComment.CreatedAt, &NewsComment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning NewsComment row: %v", err)
		}
		NewsComments = append(NewsComments, NewsComment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over NewsComments rows: %v", err)
	}

	return NewsComments, nil
}
