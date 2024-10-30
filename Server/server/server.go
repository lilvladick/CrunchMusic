package server

import (
	router "CrunchServer/Router"
	"CrunchServer/postgres"
	"context"
	"log"
)

func Run(ctx context.Context) error {
	if err := postgres.InitDatabase(); err != nil {
		log.Fatalf("data base init error: %v", err)
	}

	go func() {
		err := router.ListenAndServe()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}()

	<-ctx.Done()

	return nil
}
