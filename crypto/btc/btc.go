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
	"WindToken/db/models/buyer"
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
			go handleMessage(message)
		case io.EOF:
			return uErr.Combine(nil, uErr.ErrorConnectBTCService)
		default:
			return err
		}
	}
}

func handleMessage(line []byte) {
	var message types.BTCServiceResp

	r := bytes.NewReader(line)
	dec := gob.NewDecoder(r)
	err := dec.Decode(&message)
	if err != nil { return }

	switch message.Type {
	case messageTypes.VALUE_RECEIVED:
		updateBuyerBalance(message.Value, message.From)
	default:
		return
	}
}

func updateBuyerBalance(value float64, buyerAddr string) {
	// TODO: Check transaction if exist
	buyer.FindByBTCAddress()

}
