package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var requiredPairs = []string{
	"ETH-USDT",
	"BTC-USDT",
}

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Short rate description",
	Long:  "Long rate description",
	Args: func(cmd *cobra.Command, args []string) error {
		isValidPair := func(value string, requiredValues []string) bool {
			for _, reqValue := range requiredValues {
				if value == reqValue {
					return true
				}
			}
			return false
		}

		pair, err := cmd.Flags().GetString("pair")
		if err != nil {
			return fmt.Errorf("Cannot get --pair arg %s", err)
		}
		if !isValidPair(pair, requiredPairs) {
			return fmt.Errorf("Invalid --pair value %s, possible values: %s", pair, strings.Join(requiredPairs, ", "))
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
