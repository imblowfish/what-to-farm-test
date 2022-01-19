package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CryptoInfo struct {
	Symbol string
	Price  string
}

func GetBinanceSupportedCryptos() []string {
	return []string{
		"ETHUSDT",
		"BTCUSDT",
		// TODO @imblowfish: add new cryptos here...
	}
}

func GetServiceSupportedCryptos() []string {
	return []string{
		"ETH-USDT",
		"BTC-USDT",
		// TODO @imblowfish: add new cryptos here...
	}
}

func ConvertToBinanceSymbol(symbol string) string {
	converted, ok := map[string]string{
		"ETH-USDT": "ETHUSDT",
		"BTC-USDT": "BTCUSDT",
	}[symbol]
	if !ok {
		converted = ""
	}
	return converted
}

func ConvertFromBinanceSymbol(symbol string) string {
	converted, ok := map[string]string{
		"ETHUSDT": "ETH-USDT",
		"BTCUSDT": "BTC-USDT",
	}[symbol]
	if !ok {
		converted = ""
	}
	return converted
}

func MakeBinanceRequest() ([]CryptoInfo, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price")
	if err != nil {
		return nil, fmt.Errorf("Cannot exec GET request %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read response body %s", err)
	}
	var currenciesData []CryptoInfo
	if err := json.Unmarshal(body, &currenciesData); err != nil {
		return nil, fmt.Errorf("Cannot parse json response %s", err)
	}

	return currenciesData, nil
}
