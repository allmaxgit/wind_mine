package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"io"
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
	buildPkg("configs.toml", "")
	buildPkg(path.Join("services", "BTCService", "configs.toml"), "services", "BTCService")
	buildPkg(path.Join("services", "ExchangeRateService", "rate-conf.yaml"), "services", "ExchangeRateService")
}

func buildPkg(confPath string, pathToPkg ...string) {
	pkgPath := path.Join(pathToPkg...)
	_, confFileName := path.Split(confPath)

	if pkgPath == "" {
		mkPath(mainPkgName)

		fmt.Println("building root...")
		runCommand("go", "build", "-o", path.Join("out", "builds", mainPkgName, mainPkgName))

		err := Copy(confPath,  path.Join("out", "builds", mainPkgName, confFileName))
		if err != nil {
			log.Fatalf("failed to copy config file (%s/%s): %s", mainPkgName, confFileName, err)
		}
	} else {
		pkgPath = "." + string(os.PathSeparator) + pkgPath
		fmt.Println(pkgPath)

		_, file := path.Split(pkgPath)
		mkPath(file)

		fmt.Printf("building %s...\n", file)
		runCommand("go", "build", "-o", path.Join("out", "builds", file, file), pkgPath)

		err := Copy(confPath,  path.Join("out", "builds", file, confFileName))
		if err != nil {
			log.Fatalf("failed to copy config file (%s/%s): %s", file, confFileName, err)
		}
	}
}

func mkPath(outDirName string) {
	err := os.MkdirAll(path.Join("out", "builds", outDirName), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func runCommand(cmd string, args ...string) {
	if output, err := exec.Command(cmd, args...).Output(); err != nil {
		os.Stdout.Write(output)
		os.Exit(1)
	}
}


// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
