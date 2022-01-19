package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CryptoCurrency struct {
	ServerSideName  string
	BinanceSideName string
}

var supportedCryptoCurrencies = []CryptoCurrency{
	{"ETH-USDT", "ETHUSDT"},
	{"BTC-USDT", "BTCUSDT"},
	// TODO: Add new cryptocurrencies here...
}

func GetServiceSideSupportedCryptoNames() []string {
	values := make([]string, 0)
	for _, v := range supportedCryptoCurrencies {
		values = append(values, v.ServerSideName)
	}
	return values
}

func GetBinanceSideSupportedCryptoName() []string {
	values := make([]string, 0)
	for _, v := range supportedCryptoCurrencies {
		values = append(values, v.BinanceSideName)
	}
	return values
}

type CryptoCurrencyInfo struct {
	Symbol string
	Price  string
}

type CryptoExchangeRequester interface {
	MakeRequest() ([]CryptoCurrencyInfo, error)
}

type BinanceRequester struct{}

func (br BinanceRequester) MakeRequest() ([]CryptoCurrencyInfo, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price")
	if err != nil {
		return nil, fmt.Errorf("Cannot exec GET request %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read response body %s", err)
	}

	var currenciesData []CryptoCurrencyInfo
	if err := json.Unmarshal(body, &currenciesData); err != nil {
		return nil, fmt.Errorf("Cannot parse json response %s", err)
	}
	return currenciesData, nil
}
