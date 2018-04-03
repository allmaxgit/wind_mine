package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

	"WindToken/constants"
	uErr "WindToken/errors"
	"WindToken/types"
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

		ethRate := GetETHRate()
		btcRate := GetBTCRate()
		if ethRate != 0 {
			w.OnNewRate(constants.ETH, conf.FiatSymbol, ethRate)
		}
		if btcRate != 0 {
			w.OnNewRate(constants.BTC, conf.FiatSymbol, btcRate)
		}

	}

	return
}
