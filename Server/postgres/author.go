package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func InsertAuthor(author Author) error {
	sqlStatement := `
        INSERT INTO authors ( name, Email, password)
        VALUES ($1, $2, $3)
    `

	_, err := db.Exec(sqlStatement, author.Name, author.Email, author.Password)
	return err
}

func DeleteAuthor(id int) error {
	sqlStatement := `
        DELETE FROM authors
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func UpdateAuthor(id int, name string, Email string, password []byte) error {
	sqlStatement := `
        UPDATE authors
        SET name = $2, Email = $3, password = $4
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id, name, Email, password)
	return err
}

func QueryAuthors(Email string) (*Author, error) {
	query := "SELECT * FROM authors WHERE Email = $1"

	row := db.QueryRow(query, Email)
	var Author Author

	err := row.Scan(&Author.ID, &Author.Name, &Author.Email, &Author.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Author not found with Email %v", Email)
		}
		return nil, fmt.Errorf("error scanning Author row: %v", err)
	}

	return &Author, nil
}

func GetAuthorsByID(AuthorID int) (*Author, error) {
	query := "SELECT * FROM authors WHERE id = $1"

	row := db.QueryRow(query, AuthorID)
	var Author Author

	err := row.Scan(&Author.ID, &Author.Name, &Author.Email, &Author.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Author not found with ID %v", AuthorID)
		}
		return nil, fmt.Errorf("error scanning Author row: %v", err)
	}

	return &Author, nil
}

func GetAuthorByName(name string) (*Author, error) {
	query := "SELECT * FROM authors WHERE name = $1"

	row := db.QueryRow(query, name)
	var Author Author

	err := row.Scan(&Author.ID, &Author.Name, &Author.Email, &Author.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Author not found with name %v", name)
		}
		return nil, fmt.Errorf("error scanning Author row: %v", err)
	}

	return &Author, nil
}

func GetAuthors() ([]Author, error) {
	sqlStatement := `SELECT * FROM authors;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying Authors: %v", err)
	}
	defer rows.Close()

	var Authors []Author
	for rows.Next() {
		var Author Author
		err := rows.Scan(&Author.ID, &Author.Name, &Author.Email, &Author.Password)
		if err != nil {
			return nil, fmt.Errorf("error scanning Author row: %v", err)
		}
		Authors = append(Authors, Author)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over Authors rows: %v", err)
	}

	return Authors, nil
}
