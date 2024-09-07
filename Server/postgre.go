package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Данные для подключения к БД
const (
	host     = "localhost"
	port     = 5431
	user     = "postgres"
	password = "admin"
	dbname   = "CrunchMusic"
)

var db *sql.DB

// Функция подключения к БД
func initDatabase() error {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("не удалось проверить подключение к базе данных: %w", err)
	}

	log.Println("Подключение к базе данных успешно установлено")
	return nil
}
