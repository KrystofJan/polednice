package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/KrystofJan/tempus/internal/service"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var showEntryCmd = &cobra.Command{
	Use:   "show",
	Short: "Show entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return fmt.Errorf("PARAMETER ERROR: %v", err)
		}

		entryProvider, err := service.NewEntryProvider()
		if err != nil {
			return err
		}

		if all {
			log.Println("Looking for all entries")
			tasks, err := entryProvider.FindAllEntries()
			if err != nil {
				return fmt.Errorf("SERVICE ERROR: %v", err)
			}
			fmt.Println(tasks)
			return nil
		}

		id, err := cmd.Flags().GetInt64("id")
		if err != nil || id == 0 {
			return fmt.Errorf("PARAMETER ERROR: %v", err)
		}

		log.Printf("Looking for %d entry", id)
		tasks, err := entryProvider.FindEntryById(id)
		if err != nil {
			return fmt.Errorf("SERVICE ERROR: %v", err)
		}
		fmt.Println(tasks)
		return nil
	},
}

func init() {
	showEntryCmd.Flags().BoolP("all", "a", false, "Show all")
	showEntryCmd.Flags().Int64P("id", "i", 0, "entry id")
	entryCmd.AddCommand(showEntryCmd)
}
