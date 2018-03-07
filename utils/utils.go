package utils

import (
	"log"
	"os"
	"os/signal"
)

func RecoverWatcher(shutdownFunc func(bool, interface{})) {
	if r := recover(); r != nil {
		log.Println("----------------------------------")
		log.Println("SESSION END WITH FATAL ERROR -", r)
		log.Println("----------------------------------")
		shutdownFunc(true, r)
	}
}

func ShutdownWatcher(shutdownFunc func(bool, interface{})) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	select {
	case <-sig:
		shutdownFunc(false, nil)
	}
}
