package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arnab-xyz/filefortress.git/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	// Upload handler
	r.Post("/upload", handler.UploadFile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting filefortress on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
