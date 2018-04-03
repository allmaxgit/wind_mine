//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"WindToken/constants"
)

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
	fmt.Println("Launching TCP server...")
	if err := StartTCPServer(conf, w); err != nil {
		fmt.Println("failed to start tcp server:", err.Error())
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

		conf.Logger.Println("Average ETH/"+conf.FiatSymbol+" exchange rate -", rate)

		SetRate(constants.ETH, rate)

		if w.OnNewRate != nil {
			w.OnNewRate(constants.ETH, conf.FiatSymbol, rate)
		}

		last := GetWeiInFiatUnit(rate, 2)

		for retries := 0; retries < conf.Retries; retries++ {
			r := UpdateExchangeRate(last, conf)
			if r != nil {
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

		conf.Logger.Println("Average BTC/"+conf.FiatSymbol+" exchange rate -", rate)

		SetRate(constants.BTC, rate)

		if w.OnNewRate != nil {
			w.OnNewRate(constants.BTC, conf.FiatSymbol, rate)
		}

		conf.Logger.Println("Sleeping for", int64(conf.UpdateRate), "minutes..")
		time.Sleep(conf.UpdateRate * time.Minute)
	}
}
