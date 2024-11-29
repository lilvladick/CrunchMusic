package postgres

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
	dbname   = "MegaNews"
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

// func GetResultsJson(query string, dst interface{}, args ...interface{}) ([]byte, error) {
// 	err := MakeQuery(query, dst, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	jsonData, err := json.Marshal(dst)
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка при преобразовании результатов запроса в JSON: %w", err)
// 	}

// 	return jsonData, nil
// }

// func MakeQuery(query string, dst interface{}, args ...interface{}) error {
// 	if strings.HasPrefix(strings.ToUpper(query), "SELECT") {
// 		rows, err := db.Query(query, args...)
// 		if err != nil {
// 			return fmt.Errorf("ошибка выполнения SELECT запроса: %w", err)
// 		}
// 		defer rows.Close()

// 		err = RowsToStructs(rows, dst)
// 		if err != nil {
// 			return fmt.Errorf("ошибка при преобразовании результатов запроса в структуры: %w", err)
// 		}
// 		return nil
// 	} else {
// 		_, err := db.Exec(query, args...)
// 		if err != nil {
// 			return fmt.Errorf("ошибка выполнения запроса: %w", err)
// 		}
// 		return nil
// 	}
// }

// func RowsToStructs(rows *sql.Rows, dest interface{}) error {
// 	destValue := reflect.ValueOf(dest)

// 	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
// 		return fmt.Errorf("dest должен быть указателем на срез")
// 	}

// 	sliceValue := destValue.Elem()       // Ссылка на сам срез
// 	elemType := sliceValue.Type().Elem() // Тип элемента среза (структура)

// 	for rows.Next() {
// 		elem := reflect.New(elemType).Elem()

// 		scanArgs := make([]interface{}, elem.NumField())
// 		for i := 0; i < elem.NumField(); i++ {
// 			field := elem.Field(i)
// 			if field.CanAddr() && field.CanSet() {
// 				scanArgs[i] = field.Addr().Interface()
// 			} else {
// 				return fmt.Errorf("поле %d (%s) структуры %s не адресуемо", i, elemType.Field(i).Name, elemType.Name())
// 			}
// 		}

// 		if err := rows.Scan(scanArgs...); err != nil {
// 			return fmt.Errorf("ошибка сканирования строки: %w", err)
// 		}

// 		sliceValue.Set(reflect.Append(sliceValue, elem))
// 	}

// 	if err := rows.Err(); err != nil {
// 		return fmt.Errorf("ошибка при итерации по строкам: %w", err)
// 	}

// 	return nil
// }
