package main

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

type Config struct {
	Retries       int
	NetworkId     int
	InfuraToken   string
	FiatSymbol    string
	EthConnection *ethclient.Client
	Transactor    *bind.TransactOpts
	Contract      common.Address
	UpdateRate    time.Duration
}

var conf *Config

func GetConfig() *Config {
	viper.OnConfigChange(func(in fsnotify.Event) {
		if in.Op == fsnotify.Write {
			fmt.Println("Config file changed, reloading configuration...")
			conf = loadConfig()
		}
	})

	if conf == nil {
		conf = loadConfig()
	}
	return conf
}

func loadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("rate-conf")
	viper.SetConfigType("yaml")
	viper.SetDefault("fiat_symbol", "USD")
	viper.SetDefault("network_id", 3) // default value for network is Rinkeby
	viper.SetDefault("retries", 1)    // default value for number of retries is 1
	viper.ReadInConfig()

	c := &Config{}

	c.InfuraToken = viper.GetString("infura_token")
	c.NetworkId = viper.GetInt("network_id")

	//contractAddress := viper.GetString("contract_address")
	//if !common.IsHexAddress(contractAddress) {
	//	fmt.Println("Valid contract address is required")
	//	return nil
	//}
	//c.Contract = common.HexToAddress(contractAddress)

	c.UpdateRate = viper.GetDuration("update_rate")

	//key, err := crypto.HexToECDSA(viper.GetString("private_key"))
	//if err != nil {
	//	fmt.Println("Failed to parse secp256k1 key for transaction signing")
	//	return nil
	//}
	//c.Transactor = bind.NewKeyedTransactor(key)

	c.FiatSymbol = viper.GetString("fiat_symbol")
	c.Retries = viper.GetInt("retries")

	return c
}
