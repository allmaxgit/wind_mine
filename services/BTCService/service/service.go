package service

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"bytes"
	"encoding/gob"
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
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		if btcWatcher.OnNewValue == nil {
			btcWatcher.OnNewValue = func(value float64, from string, txHash string) {
				var message bytes.Buffer
				enc := gob.NewEncoder(&message)
				enc.Encode(types.BTCServiceResp{
					Type: messageTypes.VALUE_RECEIVED,
					Value: value,
					From: from,
					TXHash: txHash,
				})

				// TODO: Remove log
				log.Println("------------ amount", value, "from:", from)

				_, err = conn.Write(append(message.Bytes(), '\n'))
				if err != nil { uErr.Fatal(err, "failed to send response") }

			}
		}

		go handleConnection(conn,  btcWatcher)
	}
}

func handleConnection(conn net.Conn,  btcWatcher *btc.Watcher) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF { break }
			uErr.LogError(err, "filed to read message")
		}

		var message types.BTCServiceReq

		r := bytes.NewReader(line)
		dec := gob.NewDecoder(r)
		err = dec.Decode(&message)
		if err != nil {
			continue
		}

		switch message.Type {
		case messageTypes.WATCH_ADDRESS:
			go btcWatcher.StartWatchingAddress(message.Address)
		default:
			continue
		}

		//if err != nil {
		//	uErr.LogError(err, "filed to handle message with type:", message.Type)
		//	continue
		//}

		fmt.Println(message, line)
	}
}
