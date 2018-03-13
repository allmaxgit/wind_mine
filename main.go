package main

import (
	"flag"
	"os"
	"log"
	"fmt"

	uErr "WindToken/errors"
	"WindToken/db"
	"WindToken/utils"
	"WindToken/configs"
	"WindToken/crypto"
	"WindToken/crypto/btc"
	"WindToken/crypto/eth"
)

func main() {
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher(shutdown)
	go utils.ShutdownWatcher(shutdown)

	// Initiate store
	//store.InitiateStore()

	// Parse configs.
	conf, err := configs.ParseConfigs("./configs.toml")
	if err != nil {
		fmt.Println("ERROR - failed to parse configs:", err.Error())
		os.Exit(1)
	}

	// Setup prod env.
	envType := "dev"
	if *prod {
		envType = "prod"
		conf.Common.Dev = false

		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("ERROR - failed to setup log file:", err.Error())
			os.Exit(1)
		}
	}

	// Initiate database.
	err = db.Initiate(conf.DB[envType])
	if err != nil {
		fmt.Println("ERROR - failed to initialize db:", err.Error())
		os.Exit(1)
	}
	defer db.Instance.Close()

	// Start connection with ETH provider.
	err = eth.Dial(&conf.Crypto)
	if err != nil {
		fmt.Println("ERROR - failed ETH Dial:", err.Error())
		os.Exit(1)
	}

	// Start connection with ExchangeRateService.
	go func() {
		err = crypto.Dial(8083)
		if err != nil {
			fmt.Println("ERROR - failed crypto Dial:", err.Error())
			os.Exit(1)
		}
	}()

	// Start connection with BTCService.
	fmt.Println("Start btc service connection...")
	// TODO: Rename BTCAddr to BTCTrackedAddress
	if err = btc.Dial(conf.Services.BTCServicePort, conf.Crypto.BTCAddr); err != nil {
		if err.Error() == uErr.ErrorConnectBTCService {
			uErr.Fatal(err, "failed to dial tcp")
		} else {
			// TODO: Replace to panic and use uErr.Combine method
			fmt.Println("ERROR - failed to dial tcp", err.Error())
			os.Exit(1)
		}
	}
}

func shutdown(fatal bool, r interface{}) {
	if fatal {
		os.Exit(1)
	}

	log.Println("Shutdown")
	os.Exit(0)
}
