package utils

import (
	"log"
	"os/exec"
	"os"
	"io"
	"path"
)

func RunCommand(cmd string, args ...string) {
	if output, err := exec.Command(cmd, args...).Output(); err != nil {
		os.Stdout.Write(output)
		os.Exit(1)
	}
}

func CopyFile(src, dst string) error {
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

func MakePath(outDirName string) {
	err := os.MkdirAll(path.Join("out", "builds", outDirName), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
