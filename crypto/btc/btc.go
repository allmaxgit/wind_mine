package btc

import (
	"net"
	"fmt"
	"bytes"
	"encoding/gob"
	"bufio"
	"io"

	uErr "WindToken/errors"
	"WindToken/configs"
	"WindToken/types"
	"WindToken/constants/messageTypes"
)

// Dial starts connection with BTCService via tcp.
func Dial(conf *configs.Configs) (err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", conf.Services.BTCServicePort))
	if err != nil { return }

	var message bytes.Buffer
	enc := gob.NewEncoder(&message)
	enc.Encode(types.BTCServiceReq{
		Type: messageTypes.WATCH_ADDRESS,
		Address: conf.Crypto.BTCAddr,
	})

	response := bufio.NewReader(conn)
	_, err = conn.Write(append(message.Bytes(), '\n'))
	if err != nil { return }

	for {
		message, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Println("Message from server", string(message))
		case io.EOF:
			return uErr.Combine(nil, uErr.ErrorConnectBTCService)
		default:
			return err
		}
	}
}
