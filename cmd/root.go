package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "what-to-farm",
		Short: "An application that implements a server and a client to obtain data about a cryptocurrency with Binance API",
	}

	rootCmd.AddCommand(serverCmd, rateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
