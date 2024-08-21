package main

import (
	"auth/internal/handlers"
	"auth/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	Secret = "SECRET"
	Port   = "PORT"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	mux := http.NewServeMux()

	// Init services
	authService := services.NewAuthService(os.Getenv(Secret))

	// Init handlers
	authHandlers := handlers.NewAuthHandlers(authService)

	authHandlers.InstallRoutes(mux)

	err := http.ListenAndServe(os.Getenv(Port), mux)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
