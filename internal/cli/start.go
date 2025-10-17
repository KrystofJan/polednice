package cli

import (
	"github.com/spf13/cobra"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "starts your day",
}

func init() {
	rootCmd.AddCommand(start)
}
