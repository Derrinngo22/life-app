package journal

import (
	"context"
	"errors"
)

type JournalService struct {
	repo *JournalRepository
}

func NewJournalService(repo *JournalRepository) *JournalService {
	return &JournalService{repo: repo}
}

func (s *JournalService) CreateEntry(ctx context.Context, entry *JournalEntry) error {
	if entry.UserID == "" {
		return errors.New("user id is required")
	}
	if entry.Category == "" {
		return errors.New("category is required")
	}

	return s.repo.Create(ctx, entry)
}
