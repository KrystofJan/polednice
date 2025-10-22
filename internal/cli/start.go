package cli

import (
	"github.com/KrystofJan/tempus/internal/handlers"
	"github.com/KrystofJan/tempus/internal/service"
	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "starts your day",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Move all of this into a transaction
		entryProvider, err := service.NewEntryProvider()
		if err != nil {
			return err
		}
		taskProvider, err := service.NewTaskProvider()
		if err != nil {
			return err
		}
		if err = entryProvider.ClearEntries(); err != nil {
			return err
		}
		if err = taskProvider.ClearTasks(); err != nil {
			return err
		}

		task, err := handlers.EnsureDefaultTaskExists()
		if err != nil {
			return err
		}

		_, err = entryProvider.AddEntry(task.ID)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(start)
}
