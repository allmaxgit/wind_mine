package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	EthConnection *ethclient.Client
	Transactor    *bind.TransactOpts
	InfuraToken   string
	NetworkId     int
	Contract      common.Address
	UpdateRate    time.Duration
}

var conf *Config

func GetConfig() *Config {
	if conf == nil {
		conf = loadConfig()
	}
	return conf
}

func loadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("rate-conf")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	c := &Config{}

	c.InfuraToken = viper.GetString("infura_token")
	c.NetworkId = viper.GetInt("network_id")

	contractAddress := viper.GetString("contract_address")
	if !common.IsHexAddress(contractAddress) {
		return nil
	}
	c.Contract = common.HexToAddress(contractAddress)

	c.UpdateRate = viper.GetDuration("update_rate")

	key, err := crypto.HexToECDSA(viper.GetString("private_key"))
	if err != nil {
		fmt.Println("Failed to parse secp256k1 key for transaction signing")
		return nil
	}
	c.Transactor = bind.NewKeyedTransactor(key)

	return c
}
