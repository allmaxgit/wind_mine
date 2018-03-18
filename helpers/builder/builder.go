package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"WindToken/helpers/utils"
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

	build()

	fmt.Println("build success")
}

func build() {
	// TODO: Use main pkg name here.
	buildPkg("configs.toml", "")
	buildPkg(path.Join("services", "BTCService", "configs.toml"), "services", "BTCService")
	buildPkg(path.Join("services", "ExchangeRateService", "rate-conf.yaml"), "services", "ExchangeRateService")
	buildConsul()
}

func buildPkg(confPath string, pathToPkg ...string) {
	pkgPath := path.Join(pathToPkg...)
	_, confFileName := path.Split(confPath)

	if pkgPath == "" {
		// TODO: Remove it.
		utils.MakePath(mainPkgName)

		fmt.Println("building root...")
		utils.RunCommand("go", "build", "-o", path.Join("out", "builds", mainPkgName, mainPkgName))

		err := utils.CopyFile(confPath,  path.Join("out", "builds", mainPkgName, confFileName))
		if err != nil {
			log.Fatalf("failed to copy config file (%s/%s): %s", mainPkgName, confFileName, err)
		}
	} else {
		pkgPath = "." + string(os.PathSeparator) + pkgPath
		fmt.Println(pkgPath)

		_, file := path.Split(pkgPath)
		utils.MakePath(file)

		fmt.Printf("building %s...\n", file)
		utils.RunCommand("go", "build", "-o", path.Join("out", "builds", file, file), pkgPath)

		err := utils.CopyFile(confPath,  path.Join("out", "builds", file, confFileName))
		if err != nil {
			log.Fatalf("failed to copy config file (%s/%s): %s", file, confFileName, err)
		}
	}
}

func buildConsul() {
	consulPath := "." + string(os.PathSeparator) + path.Join("helpers", "consul")
	utils.RunCommand("go", "build", "-o", path.Join("out", "builds", "consul"), consulPath)
}
