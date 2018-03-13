package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var mainPkgOutPath = "./out/builds/WindToken"

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := strings.Split(pwd, "/")
	rootDir := path[len(path)-1]

	if rootDir != "WindToken" {
		log.Fatalln("please run it from 'WindToken' dir")
		return
	}

	// Run commands
	buildPkg("")
	buildPkg("./services/BTCService")
	buildPkg("./services/ExchangeRateService")

	fmt.Println("success")
}

func buildPkg(pkgPath string) {
	if pkgPath == "" {
		fmt.Println("building root...")
		runCommand("go", "build", "-o", mainPkgOutPath)
	} else {
		_, file := path.Split(pkgPath)
		fmt.Printf("building %s...\n", file)
		runCommand("go", "build", "-o", "./out/builds/" + file, pkgPath)
	}
}

func runCommand(cmd string, args ...string) {
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatalln("failed to exec command", err)
	}
}
