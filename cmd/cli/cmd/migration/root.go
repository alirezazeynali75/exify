package migration

import (
	"log"

	"github.com/alirezazeynali75/exify/internal/config"
	"github.com/spf13/cobra"
)

var mysqlConfig config.Mysql

var Root = &cobra.Command{
	Use:   "migrate",
	Short: "do/undo migration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cfg, err := config.Configure()
	if err != nil {
		log.Fatal(err.Error())
	}
	mysqlConfig = cfg.Mysql

	Root.AddCommand(up)
	Root.AddCommand(down)
}
