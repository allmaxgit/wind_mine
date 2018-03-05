package main

import (
	"flag"
	"log"
	"fmt"

	"WindToken/utils"
	"WindToken/configs"
	"WindToken/services/BTCService/service"
	"WindToken/services/BTCService/btc"
)

func main() {
	// parse flags
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher()

	// parse configs
	conf, err := configs.ParseConfigs("./configs.toml", *prod)
	if err != nil {
		fmt.Println("failed to parse configs:", err.Error())
		return
	}

	// setup prod features
	if *prod {
		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("failed to setup log file:", err.Error())
			return
		}
	}

	// connect to BTC Node
	fmt.Println("Connecting to node...")
	err = btc.StartRPCConnection()
	if err != nil {
		fmt.Println("failed connect to node", err.Error())
		return
	}

	// start TCP Server
	fmt.Println("Launching BTCService...")
	if err := service.StartTCPServer(conf.Server.TCPPort); err != nil {
		log.Println("filed to start tcp server:", err.Error())
	}
}
