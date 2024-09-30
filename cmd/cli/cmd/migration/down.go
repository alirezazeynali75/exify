package migration

import (
	"github.com/alirezazeynali75/exify/migrations"
	"github.com/spf13/cobra"
)

var down = &cobra.Command{
	Use:   "down",
	Short: "do migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.UndoMigrateDB(&mysqlConfig)
	},
}
