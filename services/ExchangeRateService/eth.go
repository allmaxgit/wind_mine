package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

const (
	//MAINNET Infura main netowork id
	MAINNET = iota + 1
	//ROPSTEN Infura Ropsten network id
	ROPSTEN
	//RINKEBY Infura Rinkeby network id
	RINKEBY
	//KOVAN Infura Kovan network id
	KOVAN
	//INFURANET Infura Infuranet network id
	INFURANET
)

func ConnectToEthereumNode(conf *Config) *ethclient.Client {
	if conf == nil {
		return nil
	}
	var infuraNode string
	switch conf.NetworkId {
	case MAINNET:
		infuraNode = "https://mainnet.infura.io/"
	case ROPSTEN:
		infuraNode = "https://ropsten.infura.io/"
	case RINKEBY:
		infuraNode = "https://rinkeby.infura.io/"
	case KOVAN:
		infuraNode = "https://kovan.infura.io/"
	case INFURANET:
		infuraNode = "https://infuranet.infura.io/"
	default:
		infuraNode = "https://rinkeby.infura.io/"
	}
	infuraNode += conf.InfuraToken

	client, err := ethclient.Dial(infuraNode)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return client
}

func GetContractSession(conf *Config) *UsingFiatPriceSession {
	if conf == nil {
		return nil
	}
	contract, err := NewUsingFiatPrice(conf.Contract, conf.EthConnection)
	if err != nil {
		fmt.Println("Failed to create session: " + err.Error())
		return nil
	}

	session := &UsingFiatPriceSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    conf.Transactor.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     conf.Transactor.From,
			Signer:   conf.Transactor.Signer,
			GasLimit: 7900000,
			Value:    conf.Transactor.Value,
		},
	}
	return session
}

func GetReceipt(tx *types.Transaction, conf *Config) *types.Receipt {
	if conf == nil {
		return nil
	}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.UpdateRate)*time.Minute)
	//defer cancel()
	receipt, err := bind.WaitMined(context.Background(), conf.EthConnection, tx)
	if err != nil {
		fmt.Println("WaitMined error: " + err.Error())
		return nil
	}
	return receipt
}
