package handlers

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	connString := os.Getenv("DB_CONNECTION")
	if connString == "" {
		log.Fatal("DB_CONNECTION not set")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Postgres pool connected")
}
