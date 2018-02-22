package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"log"
)

var outPath = "./cmd/WindToken"

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := strings.Split(pwd, "/")
	rootDir := path[len(path) - 1]

	if rootDir != "WindToken" {
		log.Fatalln("please run it from 'WindToken' dir")
		return
	}

	// Run commands
	buildPkg("")
	buildPkg("./services/BTCService")

	fmt.Println("success")
}

func buildPkg(pkgPath string) {
	if pkgPath == "" {
		runCommand("go", "build", "-o", outPath)
	} else {
		runCommand("go", "build", "-o", outPath, pkgPath)
	}
}

func runCommand(cmd string, args ...string) {
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatalln(os.Stderr, err)
	}
}
