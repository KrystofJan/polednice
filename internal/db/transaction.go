package db

import (
	"fmt"

	"github.com/KrystofJan/tempus/internal/repository"
)

type Transaction[T any] struct {
	Database *Database
}

func NewTransaction[T any]() (*Transaction[T], error) {
	database, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	return &Transaction[T]{
		Database: database,
	}, nil
}

func (tran *Transaction[T]) PerformTransaction(repo *repository.Queries, fn func(qtx *repository.Queries) (T, error)) (*T, error) {
	tx, err := tran.Database.Instance.Begin()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	qtx := repo.WithTx(tx)

	result, err := fn(qtx)
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return &result, nil
}
