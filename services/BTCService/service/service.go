package service

import (
	"fmt"
	"log"
	"net"
	"bufio"
)

// StartTCPServer
func StartTCPServer(port uint) (err error) {
	ln, err := net.Listen("tcp", ":" + string(port))
	if err != nil { return }
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(fmt.Sprintf("filed to accept message: %s", err.Error()))
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(fmt.Sprintf("filed to read message: %s", err.Error()))
		}

		log.Println("Message Received: ", string(message))
	}
}
