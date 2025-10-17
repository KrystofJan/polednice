package cli

import (
	"log"
	"os"

	"github.com/KrystofJan/polednice/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "polednice",
	Short: "Polednice - a simple CLI task tracker",
	Long:  `Track your active work sessions and move entries between tasks.`,
}

func Execute() {
	// NOTE: Need to ensure that the database is there
	_, err := db.New()
	if err != nil {
		log.Fatalf("There was a problem creating the database instance: %v", err)
		os.Exit(1)
	}
	err = db.MigrateUp()
	if err != nil {
		log.Fatalf("There was a problem migrating the database instance: %v", err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
