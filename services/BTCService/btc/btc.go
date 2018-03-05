package btc

import (
	"errors"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

const (
	ACCOUNT = "mainAccount"
)

var (
	client *rpcclient.Client
)

// StartRPCConnection starts connection to node
func StartRPCConnection() (err error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         "162.213.252.104:8332",
		User:         "bitcoin",
		Pass:         "local321",
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err = rpcclient.New(connCfg, nil)
	if err != nil { return }

	bCount, err := client.GetBlockCount()
	if err != nil { return }
	log.Println(bCount)

	lastHash, err := client.GetBlockHash(bCount)
	if err != nil { return }
	log.Println("last block", lastHash)

	lastBlock, err := client.GetBlockVerbose(lastHash)
	if err != nil { return }

	//txs, err := lastBlock.TxHashes()
	//if err != nil { return }

	for _, txStr := range lastBlock.Tx {
		txHash, _ := chainhash.NewHashFromStr(txStr)

		tx, err := client.GetRawTransactionVerbose(txHash)
		if err != nil {
			return err // TODO: Restart cycle
		}

		for _, vout := range tx.Vout {
			for _, addr := range vout.ScriptPubKey.Addresses {
				log.Println("Address:", addr)
			}
		}
	}


	return
}

// WatchAddress adds listener for address
func WatchAddress(addr string) error {
	if addr == "" {
		return errors.New("address not found")
	}

	return nil
}


//connCfg := &rpcclient.ConnConfig{
//Host:         "162.213.252.104:8332",
//User:         "bitcoin",
//Pass:         "local321",
//HTTPPostMode: true,
//DisableTLS:   true,
//}