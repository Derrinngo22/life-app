package journal

import (
	"context"
	"lifeapp/internal/db"
)

type JournalRepository struct{}

func NewJournalRepository() *JournalRepository {
	return &JournalRepository{}
}

func (r *JournalRepository) Create(ctx context.Context, entry *JournalEntry) error {
	query := `
		INSERT INTO journal_entries (category, data, user_id) 
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return db.Pool.QueryRow(
		ctx,
		query,
		entry.Category,
		entry.Data,
		entry.UserID,
	).Scan(&entry.ID, &entry.CreatedAt)

}
