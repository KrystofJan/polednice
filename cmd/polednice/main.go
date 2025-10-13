package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/KrystofJan/polednice/repository"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Print("Hello")
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "polednice.db")
	if err != nil {
		log.Fatal("cannot open db:", err)
	}

	repo := repository.New(db)
	entries, err := repo.FindAllEntries(ctx)
	if err != nil {
		panic("SOmething went wrong")
	}

	fmt.Println(entries)
}
