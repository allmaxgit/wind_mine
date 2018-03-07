package utils

import (
	"log"
	"os"
	"os/signal"
)

func RecoverWatcher(shutdownFunc func(interface{})) {
	if r := recover(); r != nil {
		log.Println("----------------------------------")
		log.Println("SESSION END WITH FATAL ERROR -", r)
		log.Println("----------------------------------")
		shutdownFunc(r)
	}
}

func ShutdownWatcher(shutdownFunc func(interface{})) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	select {
	case <-sig:
		shutdownFunc(nil)
	}
}
