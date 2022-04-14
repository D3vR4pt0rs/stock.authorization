package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"authentication/internal/infrastructure/postgres"
	"authentication/internal/interfaces/handlers"
	"authentication/internal/interfaces/repository"
	"authentication/internal/usecases/storage"

	"github.com/D3vR4pt0rs/logger"

	"github.com/gorilla/mux"
)

func main() {

	config := postgres.Config{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Ip:       os.Getenv("POSTGRES_IP"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}

	psClient, err := postgres.New(config)
	if err != nil {
		os.Exit(1)
	}

	repo := repository.New(psClient, os.Getenv("SECRET_KEY"))

	application := storage.New(repo)

	router := mux.NewRouter()
	handlers.Make(router, application)
	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		fmt.Println("Received a shutdown signal:", <-listener)

		if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			logger.Error.Println("Failed to gracefully shutdown ", err)
		}
	}()

	logger.Info.Println("[*]  Listening...")
	if err := srv.ListenAndServe(); err != nil {
		logger.Error.Println("Failed to listen and serve ", err)
	}

	logger.Critical.Println("Server shutdown")
}
