package main

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

// Функция обработки запроса
// syscall.Handle - файловый дескриптор (syscall.Handle для винды, а int для unix)
func handleRequest(fd int) {
	defer syscall.Close(fd)
	buf := make([]byte, 1024)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		log.Print("read error: ", err)
		return
	}

	request := string(buf[:n])
	log.Print(request)

	if strings.HasPrefix(request, "GET ") {
		if strings.Contains(request, "/tracks") {
			query := "SELECT * FROM tracks"

			jsonData, err := getResultsJson(query)
			if err != nil {
				log.Print("error fetching tracks: ", err)
				errorResponse := "HTTP/1.1 500 Internal Server Error\r\n" +
					"Content-Type: text/plain\r\n" +
					"\r\n" +
					"Error fetching tracks"
				syscall.Write(fd, []byte(errorResponse))
				return
			}

			response := "HTTP/1.1 200 OK\r\n" +
				"Content-Type: application/json\r\n" +
				"Content-Length: " + fmt.Sprintf("%d", len(jsonData)) + "\r\n" +
				"\r\n" +
				string(jsonData)

			syscall.Write(fd, []byte(response))
		} else {
			response := "HTTP/1.1 404 Not Found\r\n" + "\r\n"
			syscall.Write(fd, []byte(response))
		}
	} else {
		response := "HTTP/1.1 404 Not Found\r\n" + "\r\n"
		syscall.Write(fd, []byte(response))
	}
}
