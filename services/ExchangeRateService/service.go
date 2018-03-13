package main

import (
	"net"
	"fmt"
	"bytes"
	"encoding/gob"

	uErr "WindToken/errors"
	"WindToken/types"
)

// StartTCPServer starts TCP listening on certain port.
func StartTCPServer(port uint, w *Watcher) (err error) {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil { return }
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		if w.OnNewRate == nil {
			w.OnNewRate = func(currency, fiatSymbol string, value float64) {
				var message bytes.Buffer
				enc := gob.NewEncoder(&message)
				enc.Encode(types.RateServiceResp{
					Currency: currency,
					Value: value,
					FiatCurrency: fiatSymbol,
				})

				_, err = conn.Write(append(message.Bytes(), '\n'))
				if err != nil { uErr.Fatal(err, "failed to send response") }
			}
		}
	}

	return
}
