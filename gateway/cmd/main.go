package main

import (
	"gateway/internal/handlers"
	"gateway/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	AuthUrl      = "AUTH_URL"
	ResourcesUrl = "RESOURCES_URL"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	mux := http.NewServeMux()

	// Init services
	fetchService := services.NewFetchService()
	authService := services.NewAuthService(fetchService, os.Getenv(AuthUrl))
	resourcesService := services.NewResourcesService(fetchService, os.Getenv(ResourcesUrl))

	// Init handlers
	authHandler := handlers.NewAuthHandlers(authService)
	resourcesHandler := handlers.NewResourcesHandlers(resourcesService, fetchService)

	authHandler.InstallRoutes(mux)
	resourcesHandler.InstallRoutes(mux)

	err := http.ListenAndServe(os.Getenv("PORT"), mux)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
