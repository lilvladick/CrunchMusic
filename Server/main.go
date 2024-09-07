package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	if err := initDatabase(); err != nil {
		log.Fatalf("Ошибка при инициализации базы данных: %v", err)
	}

	addr := "localhost:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
	defer listener.Close()

	log.Println("Слушаем на " + addr)

	// Бесконечный цикл сервера
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Ошибка принятия соединения: %v", err)
			continue
		}
		// Обработка соединений
		go handleRequest(conn)
	}
}

// Функция обработки запроса
func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка чтения запроса: %v", err)
		return
	}

	// Логирование запроса
	log.Printf("Получен запрос: %s", request)

	// Разбор запроса
	lines := strings.Split(request, "\r\n")
	method := ""
	if len(lines) > 0 {
		parts := strings.Split(lines[0], " ")
		if len(parts) >= 2 {
			method = parts[0]
		}
	}

	// Формирование ответа
	var response string
	if method == "GET" {
		responseBody := "GET SUCCESS"
		response = "HTTP/1.1 200 OK\r\n"
		response += "Content-Type: text/html\r\n"
		response += "\r\n"
		response += responseBody
	} else {
		responseBody := "Method Not Allowed"
		response = "HTTP/1.1 405 Method Not Allowed\r\n"
		response += "Content-Type: text/html\r\n"
		response += "\r\n"
		response += responseBody
	}

	// Отправка ответа
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Ошибка записи ответа: %v", err)
		return
	}
}
