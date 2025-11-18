package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func Connect() {
	godotenv.Load("../../.env")
	connStr := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}
	Pool = pool
	log.Println("Connected to DB")
}
