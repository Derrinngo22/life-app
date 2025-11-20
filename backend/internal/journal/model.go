package journal

import "time"

type JournalEntry struct {
	ID        int64          `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	Category  string         `db:"category"`
	Data      map[string]any `db:"data"`
	UserID    string         `db:"user_id"`
}
