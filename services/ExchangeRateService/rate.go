package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ExchangeRateProvider struct {
	Name string
	URL  string
}

var (
	RateProviders = [3]*ExchangeRateProvider{
		{"CryptoCompare", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms="},
		{"Cryptonator", "https://api.cryptonator.com/api/ticker/eth-"},
		{"Coinmarketcap", "https://api.coinmarketcap.com/v1/ticker/ethereum/?convert="},
	}
)

func GetAverageRate(symbol string, c *Config) float64 {
	rates := make(chan float64, 3)

	go func() {
		rates <- GetCryptoCompareRate(symbol, c)
	}()

	go func() {
		rates <- GetCryptonatorRate(symbol, c)
	}()

	go func() {
		rates <- GetCoinMarketCapRate(symbol, c)
	}()

	count := 0.0
	sum := 0.0
	select {
	case val := <-rates:
		count++
		sum += val
		if count == 3 {
			break
		}
	}

	return sum / count
}

//GetCryptoCompareRate performs GET request to CryptoCompare for symbol and returns ETH/symbol exchange rate
func GetCryptoCompareRate(symbol string, c *Config) float64 {
	rate := -math.MaxFloat64
	if c == nil {
		return rate
	}
	fullUrl := RateProviders[0].URL + symbol

	// Perform c.Retries of request retries
	for retries := 0; retries < c.Retries; retries++ {
		resp, err := http.Get(fullUrl)
		if err != nil {
			fmt.Println("CryptoCompare request error: " + err.Error())
			continue
		}
		dec := json.NewDecoder(resp.Body)
		// Manually read JSON response, looking for key equal to fiat symbol
		for {
			t, tErr := dec.Token()
			if tErr == io.EOF {
				break
			}
			if tErr != nil {
				fmt.Println("JSON response decoding error: " + tErr.Error())
			}
			//if we have found fiat symbol in response, than next token will be exchange rate
			if fmt.Sprintf("%v", t) == symbol {
				val, _ := dec.Token()
				rate = val.(float64)
				break
			}
		}
		resp.Body.Close()
		//stop retrying requests if we have found some non-default value
		if rate != -math.MaxFloat64 {
			break
		}
	}

	return rate
}

type CryptonatorTicker struct {
	Base   string `json:"base"`
	Target string `json:"target"`
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Change string `json:"change"`
}

type CryptonatorResponse struct {
	Ticker    CryptonatorTicker `json:"ticker"`
	Timestamp int               `json:"timestamp"`
	Success   bool              `json:"success"`
	Error     string            `json:"error"`
}

func GetCryptonatorRate(symbol string, c *Config) float64 {
	rate := -math.MaxFloat64
	if c == nil {
		return rate
	}
	fullUrl := RateProviders[1].URL + strings.ToLower(symbol)

	// Perform c.Retries of request retries
	for retries := 0; retries < c.Retries; retries++ {
		resp, err := http.Get(fullUrl)
		if err != nil {
			fmt.Println("CryptoCompare request error: " + err.Error())
			continue
		}
		response := &CryptonatorResponse{}
		err = json.NewDecoder(resp.Body).Decode(response)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Failed to decode cryptonator response")
			continue
		}

		if response.Error != "" {
			fmt.Println("Cryptonator error: ", response.Error)
			continue
		}

		r, err := strconv.ParseFloat(response.Ticker.Price, 64)
		if err != nil {
			fmt.Println("Failed to parse float in response")
			continue
		}

		rate = r

		//stop retrying requests if we have found some non-default value
		if rate != -math.MaxFloat64 {
			break
		}
	}

	return rate
}

func GetCoinMarketCapRate(symbol string, c *Config) float64 {
	rate := -math.MaxFloat64
	if c == nil {
		return rate
	}
	fullUrl := RateProviders[2].URL + symbol
	responseField := "price_" + strings.ToLower(symbol)

	// Perform c.Retries of request retries
	for retries := 0; retries < c.Retries; retries++ {
		resp, err := http.Get(fullUrl)
		if err == nil {
			fmt.Println("CryptoCompare request error: " + err.Error())
			continue
		}
		dec := json.NewDecoder(resp.Body)
		// Manually read JSON response, looking for key equal to fiat symbol
		for {
			t, tErr := dec.Token()
			if tErr == io.EOF {
				break
			}
			if tErr != nil {
				fmt.Println("JSON response decoding error: " + tErr.Error())
			}
			//if we have found fiat symbol in response, than next token will be exchange rate
			if fmt.Sprintf("%v", t) == responseField {
				val, _ := dec.Token()
				strRate := fmt.Sprintf("%v", val)
				r, err := strconv.ParseFloat(strRate, 64)
				if err == nil {
					rate = r
				}
				break
			}
		}
		resp.Body.Close()
		//stop retrying requests if we have found some non-default value
		if rate != -math.MaxFloat64 {
			break
		}
	}

	return rate
}
