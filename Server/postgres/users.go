package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InsertUser(db *sql.DB, id int, name string, login string, password []byte) (string, error) {
	sqlStatement := `
        INSERT INTO users (id, name, login, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	var userId string
	err := db.QueryRow(sqlStatement, id, name, login, password).Scan(&userId) //наверное на обычный Exec заменю,но пока пусть так будет
	return userId, err
}

func DeleteUser(db *sql.DB, id int) error {
	sqlStatement := `
        DELETE FROM users
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func UpdateUser(db *sql.DB, id int, name string, login string, password []byte) error {
	sqlStatement := `
        UPDATE users
        SET name = $2, login = $3, password = $4
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id, name, login, password)
	return err
}

func QueryUsers(db *sql.DB, login string) (*sql.Rows, error) {
	sqlStatement := `
        SELECT * FROM users
        WHERE login = $1;
    `
	return db.Query(sqlStatement, login)
}

func GetUserByID(db *sql.DB, id int) (*sql.Row, error) {
	sqlStatement := `
        SELECT * FROM users
        WHERE id = $1;
    `
	return db.QueryRow(sqlStatement, id), nil
}
