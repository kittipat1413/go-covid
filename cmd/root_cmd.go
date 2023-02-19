package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bunbot",
	Short: "Starts the bunbot application",
}

func init() {
	rootCmd.AddCommand(
		migrateCmd,
		newMigrationCmd,
		printConfigCmd,
		rollbackCmd,
		serveCmd,
	)
}
