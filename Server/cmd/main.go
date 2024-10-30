package main

import (
	"CrunchServer/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	<-sigChan
	cancel()

	timeout := 5 * time.Second
	select {
	case <-ctx.Done():
		log.Print("Server stopped gracefully.")
	case <-time.After(timeout):
		log.Print("Timeout reached. Forcing shutdown.")
	}
}
