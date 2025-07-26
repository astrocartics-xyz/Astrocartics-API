package dba

import (
	"database/sql"
	"log"
	"os"
	"time"

	// _ "github.com/lib/pq"
	import _ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")

	const maxRetries = 10
	const retryDelay = 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		var err error
		// db, err = sql.Open("postgres", connStr)
		db, err = sql.Open("pgx", connStr)
		if err != nil {
			log.Printf("Attempt %d/%d: Error opening database connection: %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Attempt %d/%d: Error connecting to the database: %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
			db.Close()
			time.Sleep(retryDelay)
			continue
		}

		log.Println("Successfully connected to PostgreSQL database!")
		return
	}

	log.Fatalf("Failed to connect to PostgreSQL database after %d retries.", maxRetries)
}

func GetDB() *sql.DB {
	return db
} 
