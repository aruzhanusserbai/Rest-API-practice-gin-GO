package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=book_store sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("DB Ping Error: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
}
