package service

import (
	"context"
	"fmt"

	"github.com/KrystofJan/polednice/internal/db"
	"github.com/KrystofJan/polednice/internal/repository"
)

func FindAllEntries() ([]repository.Entry, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	entries, err := repo.FindAllEntries(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil
}
