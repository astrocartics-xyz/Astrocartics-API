package main

import (
	"log"
	"net/http"
	"os"

	"github.com/astrocartics-xyz/Astrocartics-API/controller"
	"github.com/astrocartics-xyz/Astrocartics-API/dba"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	dba.InitDB()
	controller.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 