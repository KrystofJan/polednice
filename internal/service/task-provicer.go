package service

import (
	"context"
	"fmt"

	"github.com/KrystofJan/polednice/internal/db"
	"github.com/KrystofJan/polednice/internal/repository"
)

func FindAllTasks() ([]repository.Task, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	entries, err := repo.FindAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil

}

func FindTaskById(id int64) (repository.Task, error) {
	db, err := db.New()
	if err != nil {
		return repository.Task{}, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	tasks, err := repo.FindTaskById(ctx, id)
	if err != nil {
		return repository.Task{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return tasks, nil
}

func FindTaskByName(name string) (repository.Task, error) {
	db, err := db.New()
	if err != nil {
		return repository.Task{}, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	tasks, err := repo.FindTaskByName(ctx, name)
	if err != nil {
		return repository.Task{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return tasks, nil
}
