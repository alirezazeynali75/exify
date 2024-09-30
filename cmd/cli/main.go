package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alirezazeynali75/exify/cmd/cli/cmd/migration"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "cli",
	Short: "exify command-Line interface",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("hello exify client")
	},
}

func init() {
	root.AddCommand(migration.Root)
}

func main() {
	if err := root.ExecuteContext(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
