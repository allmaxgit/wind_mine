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
	"WindToken/service"
)

func main() {
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	// Run watchers.
	defer utils.RecoverWatcher(shutdown)
	go utils.ShutdownWatcher(shutdown)

	// Parse configs.
	conf, err := configs.ParseConfigs("./configs.toml")
	if err != nil {
		uErr.Fatal(err, "failed to parse configs")
	}

	// Initiate store
	//store.InitiateStore()

	// Setup prod env.
	envType := "dev"
	if *prod {
		envType = "prod"
		conf.Common.Dev = false

		err := utils.SetupLogFile("logPath")
		if err != nil {
			uErr.Fatal(err, "failed to setup log file")
		}
	}

	// Initiate database.
	err = db.Initiate(conf.DB[envType])
	if err != nil {
		uErr.Fatal(err, "failed to initialize db")
	}
	defer db.Instance.Close()

	// Start connection with ETH provider.
	err = eth.Dial(&conf.Crypto)
	if err != nil {
		uErr.Fatal(err, "failed ETH Dial")
	}

	// Start connection with ExchangeRateService.
	go func() {
		fmt.Println("Rate service connection...")
		err = crypto.Dial(9090)
		if err != nil {
			uErr.Fatal(err, "failed crypto Dial")
		}
	}()

	// Start connection with BTCService.
	go func() {
		fmt.Println("BTC service connection...")
		// TODO: Rename BTCAddr to BTCTrackedAddress
		if err = btc.Dial(conf.Services.BTCServicePort, conf.Crypto.BTCAddr); err != nil {
			uErr.Fatal(err, "failed to dial tcp")
		}
	}()

	// Start gRPC server.
	go func() {
		fmt.Println("Start gRPC server on", conf.Server.GRPCPort, "port")
		if err := service.StartGRPCServer(conf.Server.GRPCPort); err != nil {
			uErr.Fatal(err, "failed to start gRPC server")
		}
	}()

	// Start REST server.
	fmt.Println("Start REST server on", conf.Server.RESTPort, "port")
	err = service.StartRESTGetaway(conf.Server.RESTPort, conf.Server.GRPCPort)
	if err != nil {
		uErr.Fatal(err, "failed to start REST server")
	}
}

func shutdown(fatal bool, r interface{}) {
	if fatal {
		os.Exit(1)
	}

	log.Println("Shutdown")
	os.Exit(0)
}
