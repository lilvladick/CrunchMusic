package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

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

func getResultsJson(query string) ([]byte, error) {
	var result interface{}
	var err error

	if strings.Contains(query, "tracks") {
		tracks, err := makeQuery[Track](query)
		if err != nil {
			return nil, err
		}
		result = tracks
	} else if strings.Contains(query, "users") {
		users, err := makeQuery[User](query)
		if err != nil {
			return nil, err
		}
		result = users
	} else if strings.Contains(query, "likes") {
		likes, err := makeQuery[Likes](query)
		if err != nil {
			return nil, err
		}
		result = likes
	} else {
		return nil, fmt.Errorf("unsupported query type")
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при преобразовании результатов запроса в JSON: %w", err)
	}

	return jsonData, nil
}

func makeQuery[T any](query string) (result []T, err error) {
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Закрываем результаты запроса

	err = rowsToStructs(rows, &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при преобразовании результатов запроса в структуры: %w", err)
	}
	return result, nil

}

func rowsToStructs(rows *sql.Rows, dest interface{}) error {
	destv := reflect.ValueOf(dest).Elem()

	elemType := destv.Type().Elem()

	for rows.Next() {
		rowp := reflect.New(elemType)
		rowv := rowp.Elem()

		args := make([]interface{}, rowv.NumField())
		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return rows.Err()
}
