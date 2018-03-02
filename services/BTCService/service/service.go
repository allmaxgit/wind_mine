package service

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"

	uErr "WindToken/errors"

	"WindToken/constants/messageTypes"
	"WindToken/services/BTCService/btc"
	"WindToken/types"
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
			if err == io.EOF {
				break
			}
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
		case messageTypes.SET_ADDRESS:
			err = btc.WatchAddress(message.Address)
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
