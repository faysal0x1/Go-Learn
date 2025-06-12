package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("Starting Students API server on", cfg.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	fmt.Println("Students API is running on", cfg.Addr)
	// Setup Server

}

// NewStudentsAPI initializes a new instance of the Students API.
