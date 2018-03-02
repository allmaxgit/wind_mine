package main

import (
	"flag"
	"fmt"
	"log"

	"WindToken/configs"
	"WindToken/services/BTCService/btc"
	"WindToken/services/BTCService/service"
	"WindToken/utils"
)

func main() {
	// Parse flags
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher()

	// Parse configs
	conf, err := configs.ParseConfigs("./configs.toml", *prod)
	if err != nil {
		fmt.Println("failed to parse configs:", err.Error())
		return
	}

	// Setup prod features
	if *prod {
		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("failed setup log file:", err.Error())
			return
		}
	}

	// Connect to BTC Node
	fmt.Println("Connecting to node...")
	err = btc.StartRPCConnection()
	if err != nil {
		fmt.Println("failed connect to node", err.Error())
		return
	}

	// Start TCP Server
	fmt.Println("Launching BTCService...")
	if err := service.StartTCPServer(conf.Server.TCPPort); err != nil {
		log.Println("filed to start tcp server:", err.Error())
	}
}
