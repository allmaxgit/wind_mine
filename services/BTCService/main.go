package main

import (
	"flag"
	"log"
	"fmt"
	"os"

	uErr "WindToken/errors"
	"WindToken/utils"
	"WindToken/services/BTCService/configs"
	"WindToken/services/BTCService/service"
	"WindToken/services/BTCService/btc"
	"WindToken/services/BTCService/store"
)

func main() {
	// Parse flags.
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher(shutdown)
	go utils.ShutdownWatcher(shutdown)

	// Parse configs.
	conf, err := configs.ParseConfigs("./configs.toml")
	if err != nil {
		uErr.Fatal(err, "failed to parse configs")
	}

	// Initiate store.
	store.InitiateStore(conf.Common.StorePath)

	// Setup prod env.
	if *prod {
		conf.Common.Dev = false

		err := utils.SetupLogFile(conf.Common.LogOutPath)
		if err != nil {
			uErr.Fatal(err, "failed to setup log file")
		}
	}

	// Connect to BTC Node.
	fmt.Println("Connecting to node...")
	btcWatcher, err := btc.StartRPCConnection(conf.Bitcoin)
	if err != nil {
		uErr.Fatal(err, "failed connect to node")
	}

	// Start TCP Server.
	fmt.Println("Launching TCP...")
	if err := service.StartTCPServer(conf.Server.TCPPort, btcWatcher); err != nil {
		uErr.Fatal(err, "filed to start tcp server")
	}
}

func shutdown(fatal bool, r interface{}) {
	if fatal {
		os.Exit(1)
	}

	err := store.Save()
	if err != nil {
		log.Println("FAILED TO SAVE STORE:", err)
	}

	log.Println("Shutdown")
	os.Exit(0)
}
