package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func InsertUser(id int, name string, login string, password []byte) (string, error) {
	sqlStatement := `
        INSERT INTO users (id, name, login, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	var userId string
	err := db.QueryRow(sqlStatement, id, name, login, password).Scan(&userId) //наверное на обычный Exec заменю,но пока пусть так будет
	return userId, err
}

func DeleteUser(id int) error {
	sqlStatement := `
        DELETE FROM users
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func UpdateUser(id int, name string, login string, password []byte) error {
	sqlStatement := `
        UPDATE users
        SET name = $2, login = $3, password = $4
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id, name, login, password)
	return err
}

func QueryUsers(login string) (*User, error) {
	query := "SELECT id, name, login, password FROM users WHERE login = $1"

	row := db.QueryRow(query, login)
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with login %v", login)
		}
		return nil, fmt.Errorf("error scanning user row: %v", err)
	}

	return &user, nil
}

func GetUsersByID(userID int) (*User, error) {
	query := "SELECT id, name, login, password FROM users WHERE id = $1"

	row := db.QueryRow(query, userID)
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with ID %v", userID)
		}
		return nil, fmt.Errorf("error scanning user row: %v", err)
	}

	return &user, nil
}

func GetUsrByName(name string) (*User, error) {
	query := "SELECT id, name, login, password FROM users WHERE name = $1"

	row := db.QueryRow(query, name)
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with name %v", name)
		}
		return nil, fmt.Errorf("error scanning user row: %v", err)
	}

	return &user, nil
}

func GetUsers() ([]User, error) {
	sqlStatement := `
        SELECT * FROM users;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Login, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users rows: %v", err)
	}

	return users, nil
}
