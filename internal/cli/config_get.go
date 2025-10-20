package cli

import (
	"fmt"

	"github.com/KrystofJan/polednice/internal/config"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "print out the config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Get()
		if err != nil {
			return err
		}
		fmt.Print(cfg.ToString())
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
