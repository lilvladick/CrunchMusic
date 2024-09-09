package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

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
