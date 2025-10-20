package cli

import (
	"fmt"

	"github.com/KrystofJan/polednice/internal/config"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "config",
	RunE: func(cmd *cobra.Command, args []string) error {
		generateNewConfig, err := cmd.Flags().GetBool("new")
		if err != nil {
			return err
		}

		if generateNewConfig {
			if err = config.Delete(); err != nil {
				return err
			}
			conf, err := config.New()
			if err != nil {
				return err
			}
			fmt.Print(conf)
			return nil
		}

		defaultTask, err := cmd.Flags().GetString("defaultTask")
		if err != nil {
			return err
		}
		if defaultTask == "" {
			return fmt.Errorf("Wrong input, default task cannot be empty")
		}

		cfg, err := config.Get()
		if err != nil {
			return err
		}
		cfg.DefaultTask = defaultTask
		_, err = cfg.Save()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	configSetCmd.Flags().BoolP("new", "n", false, "Generate a new config file")
	configSetCmd.Flags().String("defaultTask", "", "Generate a new config file")

	configCmd.AddCommand(configSetCmd)
}
