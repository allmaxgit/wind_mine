package utils

import (
	"log"
)

func RecoverWatcher(shutdownFunc func(interface{})) {
	if r := recover(); r != nil {
		log.Println("----------------------------------")
		log.Println("SESSION END WITH FATAL ERROR -", r)
		log.Println("----------------------------------")
		shutdownFunc(r)
	}
}
