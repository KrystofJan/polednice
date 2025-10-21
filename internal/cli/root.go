package cli

import (
	"log"
	"os"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/handlers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tempus",
	Short: "tempus - a simple CLI task tracker",
	Long:  `Track your active work sessions and move entries between tasks.`,
}

func Execute() {
	// NOTE: Need to ensure that the database is there
	if _, err := db.New(); err != nil {
		log.Fatalf("There was a problem creating the database instance: %v", err)
		os.Exit(1)
	}
	if err := db.MigrateUp(); err != nil {
		log.Fatalf("There was a problem migrating the database instance: %v", err)
		os.Exit(1)
	}

	if err := handlers.EnsureDefaultTaskExists(); err != nil {
		log.Fatalf("Error when ensuring the existance of default task: %v", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
