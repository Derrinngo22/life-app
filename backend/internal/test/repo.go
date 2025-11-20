package test

import (
	"context"
	"lifeapp/internal/db"
	"log"
)

func TestAdd(test Test) error {
	query := `INSERT INTO tests (name) VALUES ($1)`

	_, err := db.Pool.Exec(context.Background(), query, test.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
