package service

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
	"bytes"
	"encoding/gob"

	uErr "WindToken/errors"

	"WindToken/types"
	"WindToken/constants/messageTypes"
	"WindToken/services/BTCService/btc"
	"github.com/rcrowley/go-metrics"
)

// StartTCPServer starts TCP listening on certain port
func StartTCPServer(port uint) (err error) {
	l, err := net.Listen("tcp", "127.0.0.1:8082")
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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
			go btc.StartWatchingAddress(message.Address)
		default:
			continue
		}

		if err != nil {
			uErr.LogError(err, "filed to handle message with type:", message.Type)
			continue
		}

		fmt.Println(message, line)
	}
}
