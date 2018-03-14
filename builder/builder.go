package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

var mainPkgName = "WindToken"

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, rootDir := path.Split(pwd)

	if rootDir != "WindToken" {
		log.Fatalln("please run it from 'WindToken' dir")
		return
	}

	Build()

	fmt.Println("success")
}

func Build() {
	buildPkg("")
	buildPkg("services", "BTCService")
	buildPkg("services", "ExchangeRateService")
}

func buildPkg(pathToPkg ...string) {
	pkgPath := path.Join(pathToPkg...)

	if pkgPath == "" {
		mkPath(mainPkgName)

		fmt.Println("building root...")
		runCommand("go", "build", "-o", path.Join("out", "builds", mainPkgName, mainPkgName))
	} else {
		pkgPath = "." + string(os.PathSeparator) + pkgPath
		fmt.Println(pkgPath)

		_, file := path.Split(pkgPath)
		mkPath(file)

		fmt.Printf("building %s...\n", file)
		runCommand("go", "build", "-o", path.Join("out", "builds", file, file), pkgPath)
	}
}

func mkPath(outDirName string) {
	err := os.MkdirAll(path.Join("out", "builds", outDirName), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func runCommand(cmd string, args ...string) {
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatalln("failed to exec command", err)
	}
}
