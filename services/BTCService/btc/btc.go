package btc

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
	"errors"
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
	defer client.Shutdown()

	// TODO: Remove
	// Get the current block count.
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
