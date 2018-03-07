package main

import (
	"flag"
	"os"
	"log"
	"fmt"

	uErr "WindToken/errors"
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
		os.Exit(1)
	}

	// Setup prod env
	if *prod {
		conf.Common.Dev = false

		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("ERROR - failed to setup log file:", err.Error())
			os.Exit(1)
		}
	}

	// Start connection with BTCService
	for i := 0; i < 4; i++ {
		err = btc.Dial(conf)
		if err != nil {
			fmt.Println("ERROR - failed to dial tcp with BTC service")
			continue
		}
	}

	if err = btc.Dial(conf); err != nil {
		if err.Error() == uErr.ErrorConnectBTCService {
			uErr.Fatal(err, "failed to dial tcp")
		} else {
			fmt.Println("ERROR - failed to dial tcp", err.Error())
			os.Exit(1)
		}
	}
}

func shutdown(fatal bool, r interface{}) {
	log.Println("Shutdown")
	if fatal {
		os.Exit(1)
	}
	os.Exit(0)
}
