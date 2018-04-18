package btc

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"WindToken/constants/messageTypes"
	"WindToken/crypto"
	"WindToken/crypto/eth"
	"WindToken/db"
	"WindToken/db/models/buyer"
	"WindToken/db/models/transaction"
	dbTypes "WindToken/db/types"
	uErr "WindToken/errors"
	"WindToken/types"
	"math"
)

type ReturnBTCData struct {
	ReturnRequired bool
	BTCData        types.BTCServiceResp
}

// Dial starts connection with BTCService via tcp.
func Dial(port uint, walletAddress string) (err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return
	}

	var message bytes.Buffer
	enc := gob.NewEncoder(&message)
	enc.Encode(types.BTCServiceReq{
		Type:    messageTypes.WATCH_ADDRESS,
		Address: walletAddress,
	})

	response := bufio.NewReader(conn)
	_, err = conn.Write(append(message.Bytes(), '\n'))

	if err != nil {
		return
	}

	for {
		line, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			go func() {
				data := handleMessage(line)
				if data != nil && data.ReturnRequired {
					returnBtc(conn, &(data.BTCData))
				}
			}()
		case io.EOF:
			return uErr.Combine(nil, uErr.ErrConnectBTCService)
		default:
			return err
		}
	}
}

func returnBtc(conn net.Conn, data *types.BTCServiceResp) {
	if data == nil {
		log.Println("No data to return to BTCService")
		return
	}

	log.Println("BTC should be sent back, sending command to BTCService...")

	var message bytes.Buffer
	enc := gob.NewEncoder(&message)
	enc.Encode(types.BTCServiceReq{
		Type:          messageTypes.RETURN_REQUIRED,
		Value:         data.Value,
		ReceiveTXHash: data.TXHash,
		To:            data.From,
	})

	_, err := conn.Write(append(message.Bytes(), '\n'))
	if err != nil {
		err = saveUnhandledBtcReturn(data.From, data.Value, data.TXHash, false)
		if err != nil {
			log.Println("Failed to save unhandled BTC return transaction")
		}
	}
}

func handleMessage(line []byte) *ReturnBTCData {
	log.Println("handleMessage")
	var message types.BTCServiceResp

	r := bytes.NewReader(line)
	dec := gob.NewDecoder(r)
	err := dec.Decode(&message)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil
	}

	switch message.Type {
	case messageTypes.VALUE_RECEIVED:
		returnRequired := updateBuyerBalance(message.Value / math.Pow(10, 18), message.From, message.TXHash)
		if returnRequired {
			return &ReturnBTCData{ReturnRequired: returnRequired, BTCData: message}
		}
		return nil
	case messageTypes.RETURN_FAILED:
		log.Println("BTCService failed to send BTC back, saving unhandled BTC return transaction")
		err := saveUnhandledBtcReturn(message.From, message.Value, message.TXHash, false)
		if err != nil {
			log.Println("Failed to save unhandled BTC return transaction")
		}
		return nil
	default:
		return nil
	}
}

func updateBuyerBalance(value float64, buyerAddr string, txHash string) (btcReturnRequired bool) {
	log.Println("Start updating investor's balance...")
	// Waiting for rates.
	for crypto.GetBTCRate() == 0 || crypto.GetETHRate() == 0 {
		time.Sleep(5 * time.Second)
	}

	// Find investor in db.
	investor, found, err := buyer.FindByBTCAddress(buyerAddr)
	if err != nil || !found {
		if err != nil {
			uErr.LogError(err, "failed to find investor")
		}
		return
	}

	// Check if transaction already exist.
	_, found, err = transaction.FindByHash(txHash)
	if err != nil {
		uErr.Fatal(err, "failed to find transaction")
		btcReturnRequired = true
		return
	}
	if found { return }

	// Save transaction in DB.
	err = db.Instance.Insert(&dbTypes.Transaction{
		Buyer:   investor,
		BuyerId: investor.Id,
		From:    buyerAddr,
		Amount:  value,
		Hash:    txHash,
	})
	if err != nil {
		uErr.LogError(err, "failed to insert Transaction")
		btcReturnRequired = true
		return
	}

	// Check if contract is finished.
	finished, err := eth.IsCrowdsaleFinished()
	if err != nil || finished {
		log.Println("Crowdsale has finished, BTC should be returned")
		btcReturnRequired = true
		return
	}

	// Check if corresponding ETH address has passed KYC.
	kycPassed, err := eth.IsKycPassed(investor.EthAddr)
	if err != nil || !kycPassed {
		if err != nil {
			uErr.LogError(err, "failed to pass KYC")
		}
		log.Println("KYC is not passed, BTC should be returned")
		btcReturnRequired = true
		return
	}

	// Update ICO state info.
	err = eth.GetTokenPrice()
	if err != nil && err.Error() != uErr.ErrICOFinished {
		uErr.Fatal(err, "failed to get token price while updating investor balance")
	}

	// Convert BTC to tokens.
	tokensValue := eth.ConvertBTCToTokens(value)

	// Send tokens to investor ETH address.
	log.Println("Send", tokensValue.String(), "WMD tokens to investor:", investor.EthAddr)
	err = eth.SendTokens(investor.EthAddr, tokensValue)
	if err != nil {
		uErr.LogError(err, "failed to send tokens while updating investor balance")
		btcReturnRequired = true
		return
	}

	return
}

func saveUnhandledBtcReturn(from string, amount float64, hash string, kycFailed bool) error {
	err := db.Instance.Insert(&dbTypes.NotHandledBTCReturn{
		From:      from,
		Amount:    amount,
		Hash:      hash,
		KycFailed: kycFailed,
	})
	if err != nil {
		uErr.LogError(err, "failed to insert NotHandledBTCReturn from:", from)
		return errors.New("failed") //TODO return something sensible
	}

	return nil
}
