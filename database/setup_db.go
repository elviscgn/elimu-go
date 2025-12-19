package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	dbConn := os.Getenv("DB_CONNECTION")

	// Connect to Postgres
	conn, err := pgx.Connect(context.Background(), dbConn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer conn.Close(context.Background())

	ctx := context.Background()

	// Drop tables if they exist
	_, err = conn.Exec(ctx, `
        DROP TABLE IF EXISTS students;
        DROP TABLE IF EXISTS staff;
    `)
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	// Create tables
	_, err = conn.Exec(ctx, `
        CREATE TABLE students (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            created_at TIMESTAMP DEFAULT NOW()
        );

        CREATE TABLE staff (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            role VARCHAR(50),
            created_at TIMESTAMP DEFAULT NOW()
        );
    `)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	log.Println("Database setup complete!")
}
