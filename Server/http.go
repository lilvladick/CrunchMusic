package main

import (
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
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, world!"
		syscall.Write(fd, []byte(response))
	} else {
		responce := "HTTP/1.1 404 Not Found\r\n" + "\r\n"
		syscall.Write(fd, []byte(responce))
	}
}
