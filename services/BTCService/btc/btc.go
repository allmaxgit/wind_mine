package btc

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"WindToken/constants"
	uErr "WindToken/errors"
	"WindToken/services/BTCService/configs"
	"WindToken/services/BTCService/store"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type Watcher struct {
	mux    sync.Mutex
	client *rpcclient.Client

	ActiveWatcherId uint8

	OnNewValue func(float64, string, string)
}

const (
	ACCOUNT        = "mainAccount"
	LAST_BLOCK_KEY = "lastHandledBTCBlock"
)

// StartRPCConnection starts connection to node.
func StartRPCConnection(conf configs.Bitcoin) (watcher *Watcher, err error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         conf.RPCHost,
		User:         conf.User,
		Pass:         conf.Password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return
	}

	watcher = &Watcher{sync.Mutex{}, client, 0, nil}

	return
}

// IncreaseActiveWatcherId increases ActiveWatcherId.
func (w *Watcher) IncreaseActiveWatcherId() uint8 {
	w.mux.Lock()
	defer w.mux.Unlock()

	if w.ActiveWatcherId == 255 {
		w.ActiveWatcherId = 0
	} else {
		w.ActiveWatcherId++
	}
	return w.ActiveWatcherId
}

// CheckActiveWatcherIdChanged checks if ActiveWatcherId was changed.
func (w *Watcher) CheckActiveWatcherIdChanged(ownerId uint8) bool {
	w.mux.Lock()
	defer w.mux.Unlock()

	if w.ActiveWatcherId == ownerId {
		return true
	}
	return false
}

// StartWatchingAddress handles error from WatchAddress.
func (w *Watcher) StartWatchingAddress(addr string) {
	log.Println("Start watching address " + addr + "...")
	err := w.WatchAddress(addr)
	if err != nil {
		uErr.Fatal(err, "failed to watch address")
	}
}

// WatchAddress adds listener for address.
func (w *Watcher) WatchAddress(addr string) error {
	if addr == "" {
		return errors.New(uErr.ErrFindAddress)
	}

	var lastHandledBlock int64
	lastHandledBlockFromCache, found := store.Get(LAST_BLOCK_KEY)
	if found {
		lastHandledBlock = lastHandledBlockFromCache.(int64)
	}

	selfId := w.IncreaseActiveWatcherId()

	for {
		stillOwner := w.CheckActiveWatcherIdChanged(selfId)
		if !stillOwner {
			return nil
		}

		bCount, err := w.client.GetBlockCount()
		if err != nil {
			return err
		}
		log.Println("blocks count:", bCount)
		log.Println("last handled block:", lastHandledBlock)

		// If lastHandledBlock backward in time.
		backwardInTime := false
		if lastHandledBlock > 0 && bCount-lastHandledBlock > 1 {
			bCount = lastHandledBlock + 1
			backwardInTime = true
		} else if bCount == lastHandledBlock {
			sleep()
			continue
		}

		lastHash, err := w.client.GetBlockHash(bCount)
		if err != nil {
			return err
		}

		lastBlock, err := w.client.GetBlockVerbose(lastHash)
		if err != nil {
			return err
		}

		// Filter transactions.
		for _, txStr := range lastBlock.Tx {
			txHash, _ := chainhash.NewHashFromStr(txStr)

			// Get transaction.
			tx, err := w.client.GetRawTransactionVerbose(txHash)
			if err != nil {
				return err
			}

			// Filter addresses.
			for _, vout := range tx.Vout {
				for _, fAddr := range vout.ScriptPubKey.Addresses {
					if fAddr == addr { // TODO: Remove redundant logs
						log.Println("------------------")
						log.Println("Transaction for", addr)
						log.Println("txStr:", txStr)
						log.Println("tx:", tx)
						log.Println("vout:", vout)
						log.Println("------------------")

						// Detect TX owner.
						in := tx.Vin[0]
						ownerTxHash, _ := chainhash.NewHashFromStr(in.Txid)
						ownerTxOutIndex := in.Vout

						ownerTx, err := w.client.GetRawTransactionVerbose(ownerTxHash)
						if err != nil {
							return uErr.Combine(err, "failed to get owner tx")
						}

						ownerOut := ownerTx.Vout[ownerTxOutIndex]
						ownerAddr := ownerOut.ScriptPubKey.Addresses[0]
						if w.OnNewValue != nil {
							w.OnNewValue(vout.Value, ownerAddr, tx.Hash)
						}

						log.Println("tx owner:", ownerAddr, "tx value", vout.Value)
					}
				}
			}
		}

		// Save last handled block.
		store.Set(LAST_BLOCK_KEY, bCount)
		lastHandledBlock = bCount

		// Sleep for 1 minute.
		if !backwardInTime {
			sleep()
		}
	}

	return nil
}

//ReturnBTC returns `amount` BTC back to `to`, using `receiveHash` as input transaction
func (w *Watcher) ReturnBTC(to, receiveHash string, amount float64) error {
	log.Println("Sending", amount, "BTC back to", to, "...")
	txHash, _ := chainhash.NewHashFromStr(receiveHash)
	tx, err := w.client.GetRawTransactionVerbose(txHash)
	if err != nil {
		log.Println("Failed to get raw TX by hash", receiveHash, ":", err)
		return err
	}

	var txIn []btcjson.TransactionInput
	var rawTxIn []btcjson.RawTxInput
	var receivedBtc float64
	for voutIndex, vout := range tx.Vout {
		if math.Abs(vout.Value) == amount {
			receivedBtc = math.Abs(vout.Value)
			txIn = append(txIn, btcjson.TransactionInput{Txid: receiveHash, Vout: uint32(voutIndex)})
			rawTxIn = append(rawTxIn, btcjson.RawTxInput{Txid: receiveHash, Vout: uint32(voutIndex), ScriptPubKey: vout.ScriptPubKey.Hex})
			break
		}
	}

	amounts := make(map[btcutil.Address]btcutil.Amount)

	var magic int
	miningInfo, err := w.client.GetMiningInfo()
	if err != nil || miningInfo.TestNet {
		magic = constants.BitcoinTestnet3Magic
	}
	if !miningInfo.TestNet {
		magic = constants.BitcoinMainnetMagic
	}

	addr, err := btcutil.DecodeAddress(to, &chaincfg.Params{Net: wire.BitcoinNet(magic)})
	if err != nil {
		log.Println("Failed to decode BTC address:", err)
		return errors.New(uErr.ErrValidateBTCAddress)
	}

	//Get node info to get minimum relay fee
	info, err := w.client.GetNetworkInfo()
	if err != nil {
		log.Println("Failed to get node info:", err)
		return errors.New(uErr.ErrNodeInfo)
	}

	const bytesInTx = 200.0 //approximate size of transaction
	const satoshisInBytes = 5.0
	fee := satoshisInBytes * bytesInTx / 10E8
	if fee < info.RelayFee {
		fee = info.RelayFee + 0.00000001
	}
	btcWithFee := receivedBtc - fee
	amounts[addr], _ = btcutil.NewAmount(btcWithFee)

	rawTx, err := w.client.CreateRawTransaction(txIn, amounts, nil)
	if err != nil {
		log.Println("Failed to create raw transaction for returning BTC:", err)
		return errors.New(uErr.ErrCreateRawTX)
	}

	//decode private key
	wif, err := btcutil.DecodeWIF(configs.GetConfigs().Bitcoin.PrivateKey)
	if err != nil {
		log.Println("Failed to decode WIF from private key:", err)
		return errors.New(uErr.ErrPrivDecode)
	}

	//Import PriKey to node
	err = w.client.ImportPrivKey(wif)
	if err != nil {
		log.Println("Failed to import private key: ", err)
	}

	for _, out := range rawTx.TxOut {
		out.PkScript, err = txscript.PayToAddrScript(addr)
		if err != nil {
			log.Println("Failed to create scriptPubKey:", err)
			return errors.New(uErr.ErrPayToAddrScript)
		}
	}

	var wifs []string
	wifs = append(wifs, wif.String())
	// sign raw transaction and inputs for this transactions
	signedTx, allSigned, err := w.client.SignRawTransaction3(rawTx, rawTxIn, wifs)
	if err != nil {
		log.Println("SignRawTransaction failed:", err)
		return errors.New(uErr.ErrSignTx)
	}
	if !allSigned {
		log.Println("SignRawTransaction failed to sign one or more TX inputs")
		return errors.New(uErr.ErrSignIn)
	}

	// send raw transaction
	hash, err := w.client.SendRawTransaction(signedTx, false)
	if err != nil {
		log.Println("Failed to send raw transaction:", err)
		return errors.New(uErr.ErrSendRawTx)
	}

	log.Printf("Returning %f BTC to %s: %s", amount, to, hash.String())

	return nil
}

func sleep() {
	fmt.Println("sleeping for 2 minutes...")
	time.Sleep(2 * time.Minute)
}
