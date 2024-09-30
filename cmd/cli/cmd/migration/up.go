package migration

import (
	"log"

	"github.com/alirezazeynali75/exify/migrations"
	"github.com/spf13/cobra"
)

var up = &cobra.Command{
	Use:   "up",
	Short: "do migration",
	Run: func(cmd *cobra.Command, args []string) {
		err := migrations.MigrateDB(&mysqlConfig)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
