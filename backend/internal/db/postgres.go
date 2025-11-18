package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	// connStr := os.Getenv("DATABASE_URL") // e.g. "postgres://lifeapp_user:password@localhost:5432/lifeapp_db"
	connStr := "postgres://lifeapp_user:lifeapp@localhost:5433/lifeapp_db"
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}
	Pool = pool
	log.Println("Connected to DB")
}
