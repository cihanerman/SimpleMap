package main

import (
	"context"
	"errors"
	"github.com/cihanerman/SimpleMap/internal/routes"
	"github.com/cihanerman/SimpleMap/internal/simple_map"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Creating an HTTP server
	server := routes.NewServer()

	// Starting the server
	go func() {
		log.Println("Server started:", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Server startup error:", err)
		}
	}()

	// Saving the data when the program is stopped or panicked
	service := simple_map.NewStoreService()
	defer service.Save()
	// Creating a signal channel (corresponds to CTRL+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Shutting down the server by waiting for a signal
	<-sig
	log.Println("Shutdown request received. Closing the server...")

	// Set the shutdown timeout (optional)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Shutting down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error shutting down the server:", err)
	}

	log.Println("Server closed.")
}
