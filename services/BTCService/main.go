package main

import (
	"flag"
	"log"
	"fmt"
	"os"

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

	// Initiate store.
	store.InitiateStore()

	// Parse configs.
	conf, err := configs.ParseConfigs("./configs.toml")
	if err != nil {
		fmt.Println("ERROR - failed to parse configs:", err.Error())
		os.Exit(1)
	}

	// Setup prod env.
	if *prod {
		conf.Common.Dev = false

		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("ERROR - failed to setup log file:", err.Error())
			os.Exit(1)
		}
	}

	// Connect to BTC Node.
	fmt.Println("Connecting to node...")
	btcWatcher, err := btc.StartRPCConnection(conf.Bitcoin)
	if err != nil {
		fmt.Println("ERROR - failed connect to node", err.Error())
		os.Exit(1)
	}

	// Start TCP Server.
	fmt.Println("Launching TCP...")
	if err := service.StartTCPServer(conf.Server.TCPPort, btcWatcher); err != nil {
		fmt.Println("filed to start tcp server:", err.Error())
	}
}

func shutdown(fatal bool, r interface{}) {
	if fatal {
		err := store.Save()
		if err != nil {
			log.Println("FAILED TO SAVE STORE:", err)
		}
		os.Exit(1)
	}

	log.Println("Shutdown")
	os.Exit(0)
}
