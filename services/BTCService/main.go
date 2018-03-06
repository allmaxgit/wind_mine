package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"

	"WindToken/utils"
	"WindToken/services/BTCService/configs"
	"WindToken/services/BTCService/service"
	"WindToken/services/BTCService/btc"
	"WindToken/services/BTCService/store"
)

func main() {
	// Parse flags
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher(shutdown)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-sig:
			shutdown(nil)
		}
	}()

	// Initiate store
	store.InitiateStore()

	// Parse configs
	conf, err := configs.ParseConfigs("./configs.toml", *prod)
	if err != nil {
		fmt.Println("ERROR - failed to parse configs:", err.Error())
		return
	}

	// Setup prod features
	if *prod {
		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("ERROR - failed to setup log file:", err.Error())
			return
		}
	}

	// Connect to BTC Node
	fmt.Println("Connecting to node...")
	btcWatcher, err := btc.StartRPCConnection()
	if err != nil {
		fmt.Println("ERROR - failed connect to node", err.Error())
		return
	}

	// Start TCP Server
	fmt.Println("Launching TCP...")
	if err := service.StartTCPServer(conf.Server.TCPPort, btcWatcher); err != nil {
		log.Println("filed to start tcp server:", err.Error())
	}
}

func shutdown(r interface{}) {
	log.Println("Shutdown")
	os.Exit(1)
}
