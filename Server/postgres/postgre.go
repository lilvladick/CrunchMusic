package postgres

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
	password = "Vlad_sosi"
	dbname   = "CrunchMusic"
)

var db *sql.DB

// Функция подключения к БД
func InitDatabase() error {
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

func GetResultsJson(query string) ([]byte, error) {
	var result interface{}
	var err error

	if strings.Contains(query, "playlists") {
		tracks, err := MakeQuery[Playlist](query)
		if err != nil {
			//log.Printf("Ошибка: %s\n", err)
			return nil, err
		}
		result = tracks
	} else if strings.Contains(query, "users") {
		users, err := MakeQuery[User](query)
		if err != nil {
			return nil, err
		}
		result = users
	} else if strings.Contains(query, "likes") {
		likes, err := MakeQuery[Likes](query)
		if err != nil {
			return nil, err
		}
		result = likes
	} else if strings.Contains(query, "playlist_tracks") {
		likes, err := MakeQuery[Playlist_tracks](query)
		if err != nil {
			return nil, err
		}
		result = likes
	} else if strings.Contains(query, "tracks") {
		likes, err := MakeQuery[Track](query)
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

func MakeQuery[T any](query string, args ...interface{}) (result []T, err error) {
	if strings.HasPrefix(strings.ToUpper(query), "SELECT") {
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, fmt.Errorf("ошибка выполнения SELECT запроса: %w", err)
		}
		defer rows.Close()

		err = RowsToStructs(rows, &result)
		if err != nil {
			return nil, fmt.Errorf("ошибка при преобразовании результатов запроса в структуры: %w", err)
		}
		return result, nil
	} else {
		_, err := db.Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
		}

		return nil, nil
	}
}

func RowsToStructs(rows *sql.Rows, dest interface{}) error {
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
