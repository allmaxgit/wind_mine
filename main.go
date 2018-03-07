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
	"WindToken/crypto/btc"
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
		panic(fmt.Sprintf("failed to initialize db: %s", err.Error()))
	}
	defer db.Instance.Close()

	// Start connection with BTCService.
	fmt.Println("Start btc service connection...")
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
