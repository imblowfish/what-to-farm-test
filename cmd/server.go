package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Short server description",
	Long:  "Long server description",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			for {
				currenciesData, err := BinanceRequester{}.MakeRequest()
				if err != nil {
					os.Exit(1)
				}
				for _, v := range currenciesData {
					if contains(GetBinanceSideSupportedCryptoName(), v.Symbol) {
						fmt.Println(v)
					}
				}
			}
		}()
		// TODO: Run server
		time.Sleep(15 * time.Second)
	},
}
