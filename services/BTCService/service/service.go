package service

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"log"

	uErr "WindToken/errors"
	"WindToken/types"
	"WindToken/constants/messageTypes"
	"WindToken/services/BTCService/btc"
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

		btcWatcher.OnNewValue = func(value float64, from []string, txHash string) {
			var message bytes.Buffer
			enc := gob.NewEncoder(&message)
			enc.Encode(types.BTCServiceResp{
				Type:   messageTypes.VALUE_RECEIVED,
				Value:  value,
				From:   from,
				TXHash: txHash,
			})

			_, err = conn.Write(append(message.Bytes(), '\n'))
			if err != nil {
				uErr.Fatal(err, "failed to send response")
			}
		}

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
					err := btcWatcher.ReturnBTC(message.To, message.ReceiveTXHash, message.Value)
					if err != nil {
						var returnMsg bytes.Buffer
						enc := gob.NewEncoder(&returnMsg)
						enc.Encode(types.BTCServiceResp{
							Type:   messageTypes.RETURN_FAILED,
							Value:  message.Value,
							From:   []string{message.To},
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
