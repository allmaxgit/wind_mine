//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

import (
	"math/big"
	"time"
)

var lastWeiInFiatUnit *big.Int

func main() {
	conf := GetConfig()
	if conf == nil {
		panic("Failed to read config file")
	}

	conn := ConnectToEthereumNode(conf)
	if conn == nil {
		panic("Failed to connect to Ethereum node")
	}
	conf.EthConnection = conn

	for {
		rate := GetAverageRate(conf.FiatSymbol, conf)

		for retries := 0; retries < conf.Retries; retries++ {
			last := GetWeiInFiatUnit(rate, 2)
			last = UpdateExchangeRate(last, conf)
			if last != nil {
				lastWeiInFiatUnit = last
				break
			}
		}

		time.Sleep(conf.UpdateRate * time.Minute)
	}
}
