package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Short app description",
		Long:  "Long app description",
	}

	rootCmd.AddCommand(serverCmd, rateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
