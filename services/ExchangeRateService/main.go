//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

import "fmt"

func main() {
	conf := GetConfig()
	if conf == nil {
		panic("Failed to read config file")
	}

	//conn := ConnectToEthereumNode(conf)
	//if conn == nil {
	//	panic("Failed to connect to Ethereum node")
	//}
	//conf.EthConnection = conn

	//fmt.Println(GetCryptoCompareRate("EUR", conf))
	//
	//fmt.Println(GetCoinMarketCapRate("EUR", conf))
	//
	//fmt.Println(GetCryptonatorRate("EUR", conf))

	fmt.Println(GetAverageRate("EUR", conf))
}
