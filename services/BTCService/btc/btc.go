package btc

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
	"errors"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

var (
	client *rpcclient.Client
)

// StartRPCConnection starts connection to node
func StartRPCConnection() (err error) {
	ntfnHandlers := rpcclient.NotificationHandlers{
		OnFilteredBlockConnected: func(height int32, header *wire.BlockHeader, txns []*btcutil.Tx) {
			log.Printf("Block connected: %v (%d) %v",
				header.BlockHash(), height, header.Timestamp)
		},
		OnFilteredBlockDisconnected: func(height int32, header *wire.BlockHeader) {
			log.Printf("Block disconnected: %v (%d) %v",
				header.BlockHash(), height, header.Timestamp)
		},
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         "162.213.252.104:8332",
		Endpoint:     "ws",
		User:         "bitcoin",
		Pass:         "local321",
	}
	client, err = rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil { panic(err) }

	if err := client.NotifyBlocks(); err != nil {
		panic(err)
	}
	log.Println("NotifyBlocks: Registration Complete")

	blockCount, err := client.GetBlockCount()
	if err != nil { return }
	log.Printf("Block count: %d", blockCount)

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