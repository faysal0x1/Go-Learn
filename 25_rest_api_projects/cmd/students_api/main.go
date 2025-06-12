package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faysal0x1/Go-Learn/internal/config"
)

func main() {
	// Initialize the configuration

	cfg := config.MustLoad()

	// Database connection setup

	// Setup router

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Students API!"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Starting Students API server", cfg.Addr, slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath))

	// fmt.Println("Starting Students API server on", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Handle graceful shutdown
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Error shutting down server: ", slog.String("error", err.Error()))
	}

	slog.Info("Server gracefully stopped")
	// Setup Server

}

// NewStudentsAPI initializes a new instance of the Students API.
