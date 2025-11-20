package journal

import (
	"context"
	"lifeapp/internal/db"
	"log"
)

func AddJournalEntry(journal JournalEntry) error {
	query := `INSERT INTO tests (name) VALUES ($1)`

	_, err := db.Pool.Exec(context.Background(), query, journal.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
