package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	uErr "WindToken/errors"
	"WindToken/types"
	"log"
	"time"
	"WindToken/constants"
)

// StartTCPServer starts TCP listening on certain port.
func StartTCPServer(conf *Config, w *Watcher) (err error) {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", conf.Port))
	if err != nil {
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			conf.Logger.Println("ERROR - accept connection", err.Error())
			continue
		}

		conf.Logger.Println("new connection")

		w.OnNewRate = func(currency uint8, fiatSymbol string, value float64) {
			var message bytes.Buffer
			enc := gob.NewEncoder(&message)
			enc.Encode(types.RateServiceResp{
				Currency:     currency,
				Value:        value,
				FiatCurrency: fiatSymbol,
			})

			_, err = conn.Write(append(message.Bytes(), '\n'))
			if err != nil {
				uErr.LogError(err, "failed to send response")
			}
		}

		// Waiting for rates.
		for GetBTCRate() == 0 || GetETHRate() == 0 {
			log.Println("wait-------------")
			time.Sleep(5 * time.Second)
		}
		w.OnNewRate(constants.BTC, conf.FiatSymbol, GetBTCRate())
		w.OnNewRate(constants.ETH, conf.FiatSymbol, GetETHRate())
	}

	return
}
