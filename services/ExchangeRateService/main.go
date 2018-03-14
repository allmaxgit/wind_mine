//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

import (
	"log"
	"math"
	"math/big"
	"os"
	"time"
	"fmt"

	"WindToken/constants"
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

	// Create watcher.
	w := NewWatcher()

	// Run rate watchers.
	go watchETHRate(conf, w)
	go watchBTCRate(conf, w)

	// Start TCP Server.
	fmt.Println("Launching TCP...")
	// TODO: Take port from configs
	if err := StartTCPServer(9090, w); err != nil {
		fmt.Println("filed to start tcp server:", err.Error())
	}
}

func watchETHRate(conf *Config, w *Watcher) {
	for {
		conf.Logger.Println("Updating ETH rate...")
		rate := GetAverageRate(constants.ETH, conf.FiatSymbol, conf)
		if rate == -math.MaxFloat64 {
			conf.Logger.Println("Failed to get ETH rate from 3 sources, skipping")
			time.Sleep(conf.UpdateRate * time.Minute)
			continue
		}

		conf.Logger.Println("Average ETH/" + conf.FiatSymbol + " exchange rate -", rate)

		if w.OnNewRate != nil {
			w.OnNewRate(constants.ETH, conf.FiatSymbol, rate)
		}

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

func watchBTCRate(conf *Config, w *Watcher) {
	for {
		conf.Logger.Println("Updating BTC rate...")
		rate := GetAverageRate(constants.BTC, conf.FiatSymbol, conf)
		if rate == -math.MaxFloat64 {
			conf.Logger.Println("Failed to get BTC rate from 3 sources, skipping")
			time.Sleep(conf.UpdateRate * time.Minute)
			continue
		}

		conf.Logger.Println("Average BTC/" + conf.FiatSymbol + " exchange rate -", rate)

		if w.OnNewRate != nil {
			w.OnNewRate(constants.BTC, conf.FiatSymbol, rate)
		}

		conf.Logger.Println("Sleeping for", int64(conf.UpdateRate), "minutes..")
		time.Sleep(conf.UpdateRate * time.Minute)
	}
}
