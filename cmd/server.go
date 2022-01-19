package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/imblowfish/what-to-farm-test/crypto"
)

type safeMap struct {
	mu     sync.Mutex
	prices map[string]string
}

func NewSafeMap() *safeMap {
	return &safeMap{
		prices: make(map[string]string),
	}
}

func (m *safeMap) Set(symbol string, newPrice string) {
	m.mu.Lock()
	m.prices[symbol] = newPrice
	m.mu.Unlock()
}

func (m *safeMap) Value(symbol string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	price, ok := m.prices[symbol]
	if !ok {
		return ""
	}
	return price
}

func (m *safeMap) Get() map[string]string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.prices
}

func (m *safeMap) Print() {
	m.mu.Lock()
	fmt.Println(m.prices)
	m.mu.Unlock()
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Short server description",
	Long:  "Long server description",
	RunE: func(cmd *cobra.Command, args []string) error {
		const requestDuration = 5 * time.Second
		pricesMap := NewSafeMap()

		go func() {
			for {
				cryptosInfo, err := crypto.MakeBinanceRequest()
				if err != nil {
					os.Exit(1)
				}
				for _, v := range cryptosInfo {
					if contains(crypto.GetBinanceSupportedCryptos(), v.Symbol) {
						// save new price in the prices map
						pricesMap.Set(v.Symbol, v.Price)
					}
				}
				// pricesMap.Print()
				time.Sleep(requestDuration)
			}
		}()

		// run server
		http.HandleFunc("/api/v1/rates", func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "GET":
				if err := req.ParseForm(); err != nil {
					fmt.Println("Cannot parse form")
				}
				pairs := strings.Split(strings.Replace(req.Form.Get("pairs"), " ", "", 1), ",")
				fmt.Println(pairs)

				w.Header().Set("Content-Type", "application/json")
				resp := make(map[string]string)
				for _, pair := range pairs {
					priceValue := pricesMap.Value(crypto.ConvertToBinanceSymbol(pair))
					if len(priceValue) != 0 {
						resp[pair] = priceValue
					}
				}
				jsonResp, err := json.Marshal(resp)
				if err != nil {
					fmt.Errorf("Cannot create json response %s", err)
					return
				}
				w.Write(jsonResp)

			case "POST":
				body, err := ioutil.ReadAll(req.Body)
				if err != nil {
					fmt.Println("Cannot read body")
				}
				fmt.Println(string(body))
				var currencies map[string][]string
				if err := json.Unmarshal(body, &currencies); err != nil {
					fmt.Println("Cannot parse json", err)
				}

				w.Header().Set("Content-Type", "application/json")
				resp := make(map[string]string)
				for _, pair := range currencies["pairs"] {
					priceValue := pricesMap.Value(crypto.ConvertToBinanceSymbol(pair))
					if len(priceValue) != 0 {
						resp[pair] = priceValue
					}
				}

				jsonResp, err := json.Marshal(resp)
				if err != nil {
					fmt.Errorf("Cannot create json response %s", err)
					return
				}
				w.Write(jsonResp)

			default:
				http.NotFound(w, req)
			}
		})

		fmt.Println("Listening at port 3001...")
		err := http.ListenAndServe(":3001", nil)
		if err != nil {
			return fmt.Errorf("Cannot run server %s", err)
		}

		return nil
	},
}
