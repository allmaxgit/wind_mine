package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"io"
	"bytes"

	uErr "WindToken/errors"
	"WindToken/types"
	"encoding/gob"
	"WindToken/constants/messageTypes"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	var message bytes.Buffer
	enc := gob.NewEncoder(&message)
	enc.Encode(types.BTCServiceReq{Type: messageTypes.WATCH_ADDRESS, Address: "address"})

	response := bufio.NewReader(conn)
	_, err = conn.Write(append(message.Bytes(), '\n'))
	if err != nil {
		uErr.Fatal(err, "failed to send message to BTCService")
	}

	for {
		serverLine, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Print(string(serverLine))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Println("ERROR", err)
			os.Exit(2)
		}
	}
}