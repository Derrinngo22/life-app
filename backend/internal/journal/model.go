package journal

import "time"

type JournalEntry struct {
	ID        int64          `json:"id" db:"id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	Category  string         `json:"category" db:"category"`
	Data      map[string]any `json:"data" db:"data"`
	UserID    string         `json:"user_id" db:"user_id"`
}
