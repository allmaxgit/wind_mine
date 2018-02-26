package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"io"

	"WindToken/types"
	"encoding/gob"
	"bytes"
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
	enc.Encode(types.ServicePayload{Type: messageTypes.SET_ADDRESS, Address: "address"})

	response := bufio.NewReader(conn)
	_, err = conn.Write(append(message.Bytes(), '\n'))
	if err != nil { panic(err) }

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