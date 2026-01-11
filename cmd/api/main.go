package main

import (
	"log"
	"net/http"
	"os"

	"github.com/astrocartics-xyz/Astrocartics-API/controller"
	"github.com/astrocartics-xyz/Astrocartics-API/dba"
	_ "github.com/astrocartics-xyz/Astrocartics-API/docs" // Import the generated docs
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// @title Astrocartics API
// @version 1.0
// @description The application programmer interface for Astrocartics for statistics about Eve Online.
// @host api.astrocartics.xyz
// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	dba.InitDB()

	r := chi.NewRouter()
	controller.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
