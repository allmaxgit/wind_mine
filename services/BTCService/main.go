package main

import (
	"flag"
	"log"
	"fmt"

	"WindToken/utils"
	"WindToken/configs"
	"WindToken/services/BTCService/service"
)

func main() {
	prod := flag.Bool("prod", false, "Run in production mode.")
	flag.Parse()

	defer utils.RecoverWatcher()

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
