package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func insertUser(db *sql.DB, id string, name string, login string, password []byte) (string, error) {
	sqlStatement := `
        INSERT INTO users (id, name, login, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	var userId string
	err := db.QueryRow(sqlStatement, id, name, login, password).Scan(&userId) //наверное на обычный Exec заменю,но пока пусть так будет
	return userId, err
}

func deleteUser(db *sql.DB, id string) error {
	sqlStatement := `
        DELETE FROM users
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func updateUser(db *sql.DB, id string, name string, login string, password []byte) error {
	sqlStatement := `
        UPDATE users
        SET name = $2, login = $3, password = $4
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id, name, login, password)
	return err
}

func queryUsers(db *sql.DB, login string) (*sql.Rows, error) {
	sqlStatement := `
        SELECT * FROM users
        WHERE login = $1;
    `
	return db.Query(sqlStatement, login)
}

func getUserByID(db *sql.DB, id string) (*sql.Row, error) {
	sqlStatement := `
        SELECT * FROM users
        WHERE id = $1;
    `
	return db.QueryRow(sqlStatement, id), nil
}
