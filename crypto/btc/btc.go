package btc

import (
	"net"
	"fmt"
	"bytes"
	"encoding/gob"
	"bufio"
	"io"

	uErr "WindToken/errors"
	dbTypes "WindToken/db/types"
	"WindToken/types"
	"WindToken/constants/messageTypes"
	"WindToken/db"
	"WindToken/db/models/transaction"
	"WindToken/db/models/buyer"
	"WindToken/crypto"
	"WindToken/crypto/eth"
	"log"
)

// Dial starts connection with BTCService via tcp.
func Dial(port uint, walletAddress string) (err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil { return }

	var message bytes.Buffer
	enc := gob.NewEncoder(&message)
	enc.Encode(types.BTCServiceReq{
		Type: messageTypes.WATCH_ADDRESS,
		Address: walletAddress,
	})

	response := bufio.NewReader(conn)
	_, err = conn.Write(append(message.Bytes(), '\n'))
	if err != nil { return }

	for {
		line, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			go handleMessage(line)
		case io.EOF:
			return uErr.Combine(nil, uErr.ErrConnectBTCService)
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
		updateBuyerBalance(message.Value, message.From, message.TXHash)
	default:
		return
	}
}

func updateBuyerBalance(value float64, buyerAddr string, txHash string) {
	log.Println("Updateting buyer balance...")
	// Waiting for rates.
	for crypto.GetBTCRate() == 0 || crypto.GetETHRate() == 0 {}

	// Check if transaction already exist.
	_, found, err := transaction.FindByHash(txHash)
	if found { return }

	// Find buyer in db.
	buyer, found, err := buyer.FindByBTCAddress(buyerAddr)
	if err != nil {
		uErr.Fatal(err, "failed to find buyer")
	}

	// Save transaction in DB.
	var ok bool
	if !found { // TODO: Save data in db
		uErr.LogError(nil, "failed to find buyer")
	} else {
		err := db.Instance.Insert(&dbTypes.Transaction{
			Buyer: buyer,
			From: buyerAddr,
			Amount: value,
			Hash: txHash,
		})
		if err != nil {
			uErr.LogError(err, "failed to insert Transaction")
		} else {
			ok = true
		}
	}

	// Save transaction as not handled if something went wrong.
	if !ok {
		err := db.Instance.Insert(&dbTypes.NotHandledTransaction{
			From: buyerAddr,
			Amount: value,
			Hash: txHash,
		})
		if err != nil {
			uErr.LogError(err, "failed to insert NotHandledTransaction from:", buyerAddr)
		}
	} else {
		// Update ICO state info.
		err := eth.GetTokenPrice()
		if err != nil && err.Error() != uErr.ErrICOFinished {
			uErr.Fatal(err, "failed to get token price while updating buyer balance")
		}

		// Convert BTC to tokens.
		tokensValue := eth.ConvertBTCToTokens(value)

		// Send tokens to buyer ETH address.
		err = eth.SendTokens(buyer.EthAddr, tokensValue)
		if err != nil {
			uErr.Fatal(err,"failed to send tokens while updating buyer balance")
		}
	}
}
