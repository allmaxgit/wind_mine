package main

import (
	"flag"
	"log"
	"fmt"

	"WindToken/services/BTCService/utils"
	"WindToken/services/BTCService/configs"
	"WindToken/services/BTCService/service"
)

func main() {
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer recoverWatcher()

	conf, err := configs.ParseConfigs("./configs.toml", *prod)
	if err != nil {
		fmt.Println("filed to parse configs:", err.Error())
		return
	}

	if *prod {
		err := utils.SetupLogFile("logPath")
		if err != nil {
			fmt.Println("filed setup log file:", err.Error())
			return
		}
	}

	fmt.Println("Launching BTCService...")
	if err := service.StartTCPServer(conf.Server.TCPPort); err != nil {
		log.Println("filed to start tcp server:", err.Error())
	}
}

func recoverWatcher() {
	if r := recover(); r != nil {
		log.Println("----------------------------------")
		log.Println("SESSION END WITH FATAL ERROR -", r)
		log.Println("----------------------------------")
	}
}
