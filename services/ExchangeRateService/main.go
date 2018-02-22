//go:generate abigen --pkg main -sol ../../contracts/UsingFiatPrice.sol -out ./contract.go

package main

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

}
