package handlers

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB(connString string) {
	var err error
	DB, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
