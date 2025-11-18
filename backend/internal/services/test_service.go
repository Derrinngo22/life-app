package services

import (
	"context"
	"lifeapp/internal/db"
	"lifeapp/internal/models"
	"log"
)

func TestAdd(test models.Test) error {
	query := `INSERT INTO tests (name) VALUES ($1)`

	_, err := db.Pool.Exec(context.Background(), query, test.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
