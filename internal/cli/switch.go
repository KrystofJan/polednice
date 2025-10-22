package cli

import (
	"fmt"

	"github.com/KrystofJan/tempus/internal/service"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "switches tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 0 {
			return fmt.Errorf("This command needs at least one argument")
		}
		taskName := args[0]
		switchService, err := service.NewSwitchService()
		if err != nil {
			return err
		}

		// TODO: Test this
		_, err = switchService.SwitchTask(taskName)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
