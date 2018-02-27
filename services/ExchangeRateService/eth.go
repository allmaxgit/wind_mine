package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

const (
	//MAINNET Infura main network id
	MAINNET = iota + 1
	//ROPSTEN Infura Ropsten network id
	ROPSTEN
	//RINKEBY Infura Rinkeby network id
	RINKEBY
	//KOVAN Infura Kovan network id
	KOVAN
	//INFURANET Infuranet network id
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

//GetWeiInFiatUnit calculates how much wei one fiat unit is worth. This function rounds floats to `precision` in the process
func GetWeiInFiatUnit(fiatUnitsInEther float64, precision int) *big.Int {
	fiatUnitsInEther = round(fiatUnitsInEther, precision)

	weiInFiatUnit := new(big.Int)
	ether := new(big.Float)

	fmt.Sscan("1000000000000000000.0", ether)
	ether.Quo(ether, big.NewFloat(fiatUnitsInEther))
	roundBigFloat(ether, precision)
	ether.Int(weiInFiatUnit)

	return weiInFiatUnit
}

//UpdateExchangeRate updates `weiInFiat` field in the contract, specified in the config, using client and transactor from config,
//and returns value, which was passed as parameter to this function, if the op was successful or nil if op failed
func UpdateExchangeRate(weiInFiatUnit *big.Int, c *Config) *big.Int {
	if weiInFiatUnit == nil || c == nil {
		return nil
	}

	session := GetContractSession(c)
	if session == nil {
		return nil
	}

	decimals, err := session.FiatDecimals()
	if err != nil {
		fmt.Println("Failed to get fiat decimals from the contract")
		return nil
	}

	// calculating number of wei in minimal fiat currency fracture
	fiatInWei := weiInFiatUnit.Div(weiInFiatUnit, big.NewInt(int64(math.Pow(10, float64(decimals.Int64())))))

	tx, err := session.UpdateWeiInFiat(fiatInWei)
	if err != nil {
		fmt.Println("Failed to send TX UpdateWeiInFiat")
		return nil
	}

	receipt := GetReceipt(tx, c)
	if receipt == nil {
		fmt.Println("Failed to get receipt for TX UpdateWeiInFiat with hash", tx.Hash().Hex())
		return nil
	}

	if receipt.Status != 1 {
		fmt.Println("TX UpdateWeiInFiat failed")
		return nil
	}
	fmt.Println("Updated weiInFiat with value", fiatInWei.String())
	return weiInFiatUnit
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func roundBigFloat(x *big.Float, prec int) *big.Float {
	fl, _ := x.Float64()
	x.SetFloat64(round(fl, 2))
	return x
}
