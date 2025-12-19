package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Staff struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type Users struct {
	Students []Student `json:"students"`
	Staff    []Staff   `json:"staff"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	dbConn := os.Getenv("DB_CONNECTION")

	// connect to Postgres
	conn, err := pgx.Connect(context.Background(), dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	ctx := context.Background()

	// file dumb of test users u can have your own
	file, err := os.Open("./test/users.json")
	if err != nil {
		log.Fatal("Failed to open JSON file:", err)
	}
	defer file.Close()

	var users Users
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		log.Fatal("Failed to decode JSON:", err)
	}

	// insert students
	for _, s := range users.Students {
		_, err := conn.Exec(ctx, `
            INSERT INTO students (first_name, last_name, email)
            VALUES ($1, $2, $3)
            ON CONFLICT (email) DO NOTHING
        `, s.FirstName, s.LastName, s.Email)
		if err != nil {
			log.Println("Error inserting student:", err)
		}
	}

	// insert staff
	for _, s := range users.Staff {
		_, err := conn.Exec(ctx, `
            INSERT INTO staff (first_name, last_name, email, role)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (email) DO NOTHING
        `, s.FirstName, s.LastName, s.Email, s.Role)
		if err != nil {
			log.Println("Error inserting staff:", err)
		}
	}

	fmt.Println("Users imported successfully!")
}
