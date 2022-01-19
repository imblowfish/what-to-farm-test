package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Short server description",
	Long:  "Long server description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run server mode")
	},
}
