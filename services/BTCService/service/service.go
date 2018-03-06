package service

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"bytes"
	"encoding/gob"

	"WindToken/types"
	"WindToken/constants/messageTypes"
	"WindToken/services/BTCService/btc"
	"log"
)

// StartTCPServer starts TCP listening on certain port
func StartTCPServer(port uint, btcWatcher *btc.Watcher) (err error) {
	l, err := net.Listen("tcp", "127.0.0.1:8082") // TODO: Use port from configs
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

		btcWatcher.OnNewValue = func(value float64, from string) {
			// TODO: Send data in gob
			log.Println("------------ amount", value, "from:", from)
			conn.Write([]byte("hello\n"))
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
			panic(err) // TODO: Improve
		}

		r := bytes.NewReader(line)
		var message types.ServicePayload
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
