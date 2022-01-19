package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/spf13/cobra"

	"github.com/imblowfish/what-to-farm-test/crypto"
	"github.com/imblowfish/what-to-farm-test/safemap"
	"github.com/imblowfish/what-to-farm-test/utils"
)

const binanceRequestDuration = 5 * time.Second

type errReason = string

func parseGetPairs(req *http.Request) ([]string, errReason) {
	if err := req.ParseForm(); err != nil {
		return nil, fmt.Sprint("Cannot parse params ", err)
	}
	strWithoutSpaces := strings.Replace(req.Form.Get("pairs"), " ", "", 1)
	return strings.Split(strWithoutSpaces, ","), ""
}

func parsePostPairs(req *http.Request) ([]string, errReason) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Sprint("Cannot read request body ", err)
	}
	var currencies map[string][]string
	if err := json.Unmarshal(body, &currencies); err != nil {
		return nil, fmt.Sprint("Cannot parse json ", err)
	}
	_, ok := currencies["pairs"]
	if !ok {
		return nil, `Cannot find "pairs" in request json`

	}
	return currencies["pairs"], ""
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Implements the server part of the application, runs and synchronizes data about the supported currency",
	RunE: func(cmd *cobra.Command, args []string) error {

		pricesMap := safemap.New()

		go func() {
			for {
				cryptosInfo, err := crypto.MakeBinanceRequest()
				if err != nil {
					os.Exit(1)
				}
				for _, v := range cryptosInfo {
					if utils.Contains(crypto.GetBinanceSupportedCryptos(), v.Symbol) {
						pricesMap.Set(v.Symbol, v.Price)
					}
				}
				time.Sleep(binanceRequestDuration)
			}
		}()

		http.HandleFunc("/api/v1/rates", func(w http.ResponseWriter, req *http.Request) {
			var pairs []string
			var errorReason errReason

			switch req.Method {
			case "GET":
				pairs, errorReason = parseGetPairs(req)

			case "POST":
				pairs, errorReason = parsePostPairs(req)

			default:
				http.NotFound(w, req)
				return
			}

			// send response
			w.Header().Set("Content-Type", "application/json")
			var resp = make(map[string]string)

			for _, pair := range pairs {
				priceValue := pricesMap.Value(crypto.ConvertToBinanceSymbol(pair))
				if len(priceValue) != 0 {
					resp[pair] = priceValue
				}
			}
			if len(resp) == 0 {
				resp["res"] = "Error"
				resp["reason"] = errorReason

				if len(resp["reason"]) == 0 {
					resp["reason"] = "Unknown error"
				}
			}

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				fmt.Errorf("Cannot create json response %s", err)
				return
			}
			w.Write(jsonResp)
		})

		fmt.Println("Listening at port 3001...")
		err := http.ListenAndServe(":3001", nil)
		if err != nil {
			return fmt.Errorf("Cannot run server %s", err)
		}

		return nil
	},
}
