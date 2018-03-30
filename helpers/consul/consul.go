package main

import (
	"flag"
)

func main() {
	watch := flag.Bool("watch", false, "Start watching for services.")
	flag.Parse()

	if *watch {
		// TODO: Finish.
		return
	}
}
