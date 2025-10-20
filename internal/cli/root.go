package cli

import (
	"context"
	"log"
	"os"

	"github.com/KrystofJan/polednice/internal/config"
	"github.com/KrystofJan/polednice/internal/db"
	"github.com/KrystofJan/polednice/internal/repository"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "polednice",
	Short: "Polednice - a simple CLI task tracker",
	Long:  `Track your active work sessions and move entries between tasks.`,
}

func Execute() {
	// NOTE: Need to ensure that the database is there
	database, err := db.New()
	if err != nil {
		log.Fatalf("There was a problem creating the database instance: %v", err)
		os.Exit(1)
	}
	err = db.MigrateUp()
	if err != nil {
		log.Fatalf("There was a problem migrating the database instance: %v", err)
		os.Exit(1)
	}

	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("There has been a problem getting the config file.\nIf you do not have a config file set in `~/.config/polednice`, please generate one using `polednice config set -n`\n%v", err)
		os.Exit(1)
	}

	ctx := context.Background()
	repo := repository.New(database.Instance)
	_, err = repo.FindTaskByName(ctx, cfg.DefaultTask)
	if err != nil {
		_, err := repo.AddTask(ctx, cfg.DefaultTask)
		if err != nil {
			log.Fatalf("Error querying the database, adding default task: %v", err.Error())
			os.Exit(1)
		}
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
