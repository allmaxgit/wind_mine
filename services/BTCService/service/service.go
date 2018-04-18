package service

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"math"

	"WindToken/constants/messageTypes"
	uErr "WindToken/errors"
	"WindToken/services/BTCService/btc"
	"WindToken/types"
)

// StartTCPServer starts TCP listening on certain port.
func StartTCPServer(port uint, btcWatcher *btc.Watcher) (err error) {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		btcWatcher.OnNewValue = func(value float64, from string, txHash string) {
			var message bytes.Buffer
			enc := gob.NewEncoder(&message)
			err := enc.Encode(types.BTCServiceResp{
				Type:   messageTypes.VALUE_RECEIVED,
				Value:  value * math.Pow(10, 18),
				From:   from,
				TXHash: txHash,
			})
			if err != nil {
				uErr.Fatal(err, "failed to encode in OnNewValue")
			}

			_, err = conn.Write(append(message.Bytes(), '\n'))
			if err != nil {
				uErr.Fatal(err, "failed to send response")
			}
		}

		btcWatcher.OnNewValue(1.51956470, "2NCD8CvpF4UAETvoKdQV3mKxqQW7Z6CEGKU", "31fe72f7b16ea14c84ec7bf467962814b0924cb4687e48f1ff68dd14c212436d")
		go handleConnection(conn, btcWatcher)
	}
}

func handleConnection(conn net.Conn, btcWatcher *btc.Watcher) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF {
				break
			}
			uErr.LogError(err, "failed to read message")
		}

		var message types.BTCServiceReq

		r := bytes.NewReader(line)
		dec := gob.NewDecoder(r)
		err = dec.Decode(&message)
		if err == nil {
			switch message.Type {
			case messageTypes.WATCH_ADDRESS:
				go btcWatcher.StartWatchingAddress(message.Address)
				continue
			case messageTypes.RETURN_REQUIRED:
				go func() {
					err := btcWatcher.ReturnBTC(message.To, message.ReceiveTXHash, message.Value / math.Pow(10, 18))
					if err != nil {
						uErr.LogError(err, "failed to return BTC")
						var returnMsg bytes.Buffer
						enc := gob.NewEncoder(&returnMsg)
						enc.Encode(types.BTCServiceResp{
							Type:   messageTypes.RETURN_FAILED,
							Value:  message.Value,
							From:   message.To,
							TXHash: message.ReceiveTXHash,
						})

						_, err = conn.Write(append(returnMsg.Bytes(), '\n'))
						if err != nil {
							log.Println("Failed to notify main service about failed BTC return")
						}
					}
				}()
				continue
			default:
				continue
			}
		}
	}
}
