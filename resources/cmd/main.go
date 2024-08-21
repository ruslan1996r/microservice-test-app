package main

import (
	"log"
	"resources/internal/handlers"
	"resources/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	r := gin.New()

	// Init storages
	userRepo := storage.NewUserRepository()
	bookRepo := storage.NewBooksRepository()

	// Init mocks
	userRepo.CreateMocks(10)
	bookRepo.CreateMocks(10)

	// Init handlers
	userHandlers := handlers.NewUsersHandler(userRepo)
	bookHandlers := handlers.NewBooksHandler(bookRepo)

	// Setup router
	userHandlers.InstallRoutes(r)
	bookHandlers.InstallRoutes(r)

	err := r.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
