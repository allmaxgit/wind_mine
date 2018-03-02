//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

import (
	"log"
	"math"
	"math/big"
	"os"
	"time"
)

var lastWeiInFiatUnit *big.Int

func main() {
	conf := GetConfig()
	if conf == nil {
		panic("Failed to read config file")
	}
	conf.Logger = log.New(os.Stdout, "ExchangeRateService: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	conn := ConnectToEthereumNode(conf)
	if conn == nil {
		panic("Failed to connect to Ethereum node")
	}
	conf.EthConnection = conn

	for {
		conf.Logger.Println("Updating exchange rate...")
		rate := GetAverageRate(conf.FiatSymbol, conf)
		if rate == -math.MaxFloat64 {
			conf.Logger.Println("Failed to get exchange rate from 3 sources, skipping")
			time.Sleep(conf.UpdateRate * time.Minute)
			continue
		}

		conf.Logger.Println("Average ETH/"+conf.FiatSymbol+" exchange rate -", rate)
		last := GetWeiInFiatUnit(rate, 2)

		for retries := 0; retries < conf.Retries; retries++ {
			if lastWeiInFiatUnit != nil && last.Cmp(lastWeiInFiatUnit) == 0 {
				conf.Logger.Println("Exchange rate has not changed, skipping")
				break
			}
			last = UpdateExchangeRate(last, conf)
			if last != nil {
				lastWeiInFiatUnit = last
				conf.Logger.Println("Exchange rate is successfully updated")
				break
			}
		}

		conf.Logger.Println("Sleeping for", int64(conf.UpdateRate), "minutes..")
		time.Sleep(conf.UpdateRate * time.Minute)
	}
}
