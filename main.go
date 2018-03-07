package main

import (
	"flag"
	"os"
	"log"
	"fmt"

	"WindToken/utils"
	"WindToken/configs"
	"WindToken/crypto/btc"
)

func main() {
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher(shutdown)
	go utils.ShutdownWatcher(shutdown)

	// Initiate store
	//store.InitiateStore()

	// Parse configs
	conf, err := configs.ParseConfigs("./configs.toml")
	if err != nil {
		fmt.Println("ERROR - failed to parse configs:", err.Error())
		return
	}

	// Setup prod env
	if *prod {
		conf.Common.Dev = false

		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("ERROR - failed to setup log file:", err.Error())
			return
		}
	}

	// Start connection with BTCService
	btc.Dial()
}

func shutdown(r interface{}) {
	log.Println("Shutdown")
	os.Exit(1)
}
