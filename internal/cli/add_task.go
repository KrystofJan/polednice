package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"log"

	"github.com/KrystofJan/polednice/internal/service"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var runTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, nameErr := cmd.Flags().GetString("name")
		id, idErr := cmd.Flags().GetInt64("id")

		if nameErr != nil && (idErr != nil || id == 0) {
			return fmt.Errorf("PARAMETER ERROR: Neither id or name is set, at least of of these needs to be set\nidError: %v\nnameError: %v\n", idErr, nameErr)
		}

		if idErr != nil {
			log.Printf("Looking for %s task", name)
			tasks, err := service.FindTaskByName(name)
			if err != nil {
				log.Fatalf("SERVICE ERROR: %v", err)
			}
			fmt.Println(tasks)
			return nil
		}

		log.Printf("Looking for %d task", id)
		tasks, err := service.FindTaskById(id)
		if err != nil {
			log.Fatalf("SERVICE ERROR: %v", err)
		}
		fmt.Println(tasks)
		return nil
	},
}

func init() {
	runTaskCmd.Flags().BoolP("all", "a", false, "Show all")
	runTaskCmd.Flags().StringP("name", "n", "", "Task name")
	runTaskCmd.Flags().Int64P("id", "i", 0, "Task id")
	taskCmd.AddCommand(runTaskCmd)
}
