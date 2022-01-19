package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"github.com/imblowfish/what-to-farm-test/crypto"
	"github.com/imblowfish/what-to-farm-test/utils"
)

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Short rate description",
	Long:  "Long rate description",
	RunE: func(cmd *cobra.Command, args []string) error {
		pair, err := cmd.Flags().GetString("pair")
		if err != nil {
			return fmt.Errorf("Cannot get --pair arg %s", err)
		}
		if !utils.Contains(crypto.GetServiceSupportedCryptos(), pair) {
			return fmt.Errorf("Invalid --pair value %s, possible values: %s", pair, strings.Join(crypto.GetServiceSupportedCryptos(), ", "))
		}
		request := fmt.Sprintf("http://localhost:3001/api/v1/rates?pairs=%s", pair)
		resp, err := http.Get(request)
		if err != nil {
			return fmt.Errorf("Cannot exec GET request %s", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Cannot read response body %s", err)
		}
		var jsonResponse = make(map[string]string)
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			return fmt.Errorf("Cannot parse json response %s", err)
		}
		fmt.Println(jsonResponse[pair])

		return nil
	},
}

func init() {
	rateCmd.Flags().String("pair", "", "ETH-USDT or BTC-USDT")
	rateCmd.MarkFlagRequired("pair")
}
