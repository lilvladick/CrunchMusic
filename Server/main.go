package main

import (
	"log"
	"net"

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
