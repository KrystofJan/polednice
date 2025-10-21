package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"log"

	"github.com/KrystofJan/tempus/internal/service"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var showTaskCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return fmt.Errorf("PARAMETER ERROR: %v", err)
		}

		if all {
			log.Println("Looking for all tasks")
			tasks, err := service.FindAllEntries()
			if err != nil {
				return fmt.Errorf("SERVICE ERROR: %v", err)
			}
			fmt.Println(tasks)
			return nil
		}

		name, nameErr := cmd.Flags().GetString("name")
		id, idErr := cmd.Flags().GetInt64("id")

		taskProvider, err := service.NewTaskProvider()
		if err != nil {
			return err
		}

		if nameErr != nil && (idErr != nil || id == 0) {
			return fmt.Errorf("PARAMETER ERROR: Neither id or name is set, at least of of these needs to be set\nidError: %v\nnameError: %v\n", idErr, nameErr)
		}

		if idErr != nil {
			log.Printf("Looking for %s task", name)
			tasks, err := taskProvider.FindTaskByName(name)
			if err != nil {
				log.Fatalf("SERVICE ERROR: %v", err)
			}
			fmt.Println(tasks)
			return nil
		}

		log.Printf("Looking for %d task", id)
		tasks, err := taskProvider.FindTaskById(id)
		if err != nil {
			log.Fatalf("SERVICE ERROR: %v", err)
		}
		fmt.Println(tasks)
		return nil
	},
}

func init() {
	showTaskCmd.Flags().BoolP("all", "a", false, "Show all")
	showTaskCmd.Flags().StringP("name", "n", "", "Task name")
	showTaskCmd.Flags().Int64P("id", "i", 0, "Task id")
	taskCmd.AddCommand(showTaskCmd)
}
