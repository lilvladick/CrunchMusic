package main

import (
	"CrunchServer/postgres"
	"log"
	"os"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0) // AF_INET = IPv4, SOCK_STREAM = TCP
	if err != nil {
		log.Print("socket creating errorrror: ", err)
		os.Exit(1)
	}
	defer syscall.Close(fd)

	if err := postgres.InitDatabase(); err != nil {
		log.Fatalf("data base init error: %v", err)
	}

	addr := syscall.SockaddrInet4{Port: 8080}
	copy(addr.Addr[:], []byte{0, 0, 0, 0})

	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Print("bind error: ", err)
		os.Exit(1)
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		log.Print("listen error: ", err)
		os.Exit(1)
	}
	log.Print("server started")

	for {
		connFd, _, err := syscall.Accept(fd) // AF_INET = IPv4, SOCK_STREAM = TCP
		if err != nil {
			log.Print("accept error: ", err)
			os.Exit(1)
		}

		go handleRequest(connFd)
	}
}

// func main() {
//     // Start TCP socket server in a goroutine
//     go startTCPServer()

//     // Start HTTP server
//     http.Handle("/", http.HandlerFunc(router.Serve))
//     log.Fatal(http.ListenAndServe(":8081", nil))
// }

// func startTCPServer() {
//     fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
//     if err != nil {
//         log.Print("socket creating error: ", err)
//         os.Exit(1)
//     }
//     defer syscall.Close(fd)

//     if err := postgres.InitDatabase(); err != nil {
//         log.Fatalf("database init error: %v", err)
//     }

//     addr := syscall.SockaddrInet4{Port: 8080}
//     copy(addr.Addr[:], []byte{0, 0, 0, 0})

//     err = syscall.Bind(fd, &addr)
//     if err != nil {
//         log.Print("bind error: ", err)
//         os.Exit(1)
//     }

//     err = syscall.Listen(fd, 10)
//     if err != nil {
//         log.Print("listen error: ", err)
//         os.Exit(1)
//     }
//     log.Print("TCP server started")

//     for {
//         connFd, _, err := syscall.Accept(fd)
//         if err != nil {
//             log.Print("accept error: ", err)
//             os.Exit(1)
//         }

//         go handleRequest(connFd)
//     }
// }
