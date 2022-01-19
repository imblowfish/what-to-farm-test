package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Short rate description",
	Long:  "Long rate description",
	Args: func(cmd *cobra.Command, args []string) error {
		pair, err := cmd.Flags().GetString("pair")
		if err != nil {
			return fmt.Errorf("Cannot get --pair arg %s", err)
		}
		if !contains(GetServiceSideSupportedCryptoNames(), pair) {
			return fmt.Errorf("Invalid --pair value %s, possible values: %s", pair, strings.Join(GetServiceSideSupportedCryptoNames(), ", "))
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Run rate mode")
		return nil
	},
}

func init() {
	rateCmd.Flags().String("pair", "", "ETH-USDT or BTC-USDT")
	rateCmd.MarkFlagRequired("pair")
}
