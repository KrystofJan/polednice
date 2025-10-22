package service

import (
	"context"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/repository"
)

type CurrentEntryProvider struct {
	ctx  context.Context
	repo *repository.Queries
}

func NewCurrentEntryProvider() (*CurrentEntryProvider, error) {
	db, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	return &CurrentEntryProvider{
		ctx:  ctx,
		repo: repo,
	}, nil
}
