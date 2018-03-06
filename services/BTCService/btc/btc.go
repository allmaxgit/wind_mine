package btc

import (
	"errors"
	"log"
	"time"

	uErr "WindToken/errors"
	"WindToken/services/BTCService/store"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"sync"
)

type Info struct {
	mux              sync.Mutex
	ActiveWatcherId  uint8
}

const (
	ACCOUNT   = "mainAccount"
	LAST_BLOCK_KEY = "lastHandledBTCBlock"
)

var (
	client *rpcclient.Client

	info   Info
)

// IncreaseActiveWatcherId increases ActiveWatcherId
func (inf *Info) IncreaseActiveWatcherId() uint8 {
	inf.mux.Lock()
	defer inf.mux.Unlock()

	if inf.ActiveWatcherId == 255 {
		inf.ActiveWatcherId = 0
	} else {
		inf.ActiveWatcherId++
	}
	return inf.ActiveWatcherId
}

// CheckActiveWatcherIdChanged checks if ActiveWatcherId was changed
func (inf *Info) CheckActiveWatcherIdChanged(ownerId uint8) bool {
	inf.mux.Lock()
	defer inf.mux.Unlock()

	if inf.ActiveWatcherId == ownerId { return true }
	return false
}

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

	// TODO: Move it
	go WatchAddress("mxjq47LnJMfg3aLkmF943kLU1xYkxcZnk3")

	info = Info{sync.Mutex{}, 0}

	return
}

// StartWatchingAddress handles error from WatchAddress
func StartWatchingAddress(addr string) {
	err := WatchAddress(addr)
	if err != nil {
		uErr.Fatal(err, "failed to watch address")
	}
}

// WatchAddress adds listener for address
func WatchAddress(addr string) error {
	if addr == "" {
		return errors.New(uErr.ErrorFindAddress)
	}

	var lastHandledBlock int64
	lastHandledBlockFromCache, found := store.Get(LAST_BLOCK_KEY)
	if found {
		lastHandledBlock = lastHandledBlockFromCache.(int64)
	}

	selfId := info.IncreaseActiveWatcherId()

	for {
		stillOwner := info.CheckActiveWatcherIdChanged(selfId)
		if !stillOwner { return nil }

		// TODO: Check if it's new block by height
		bCount, err := client.GetBlockCount()
		if err != nil { return err }
		log.Println("blocks count:", bCount) // TODO: Remove

		// If lastHandledBlock backward in time
		backwardInTime := false
		if found && bCount - lastHandledBlock > 1 {
			bCount = lastHandledBlock + 1
			backwardInTime = true
		}

		lastHash, err := client.GetBlockHash(bCount)
		if err != nil { return err }
		log.Println("last block", lastHash)

		lastBlock, err := client.GetBlockVerbose(lastHash)
		if err != nil { return err }

		//txs, err := lastBlock.TxHashes()
		//if err != nil { return }

		// Filter transactions
		for _, txStr := range lastBlock.Tx {
			txHash, _ := chainhash.NewHashFromStr(txStr)

			// Get transaction
			tx, err := client.GetRawTransactionVerbose(txHash)
			if err != nil { return err }

			// Filter addresses
			for _, vout := range tx.Vout {
				for _, fAddr := range vout.ScriptPubKey.Addresses {
					log.Println("Address:", fAddr)
					if fAddr == addr { // TODO: Send data via tcp
						log.Println("------------------")
						log.Println("Transaction for", addr)
						log.Println("txStr:", txStr)
						log.Println("tx:", tx)
						log.Println("vout:", vout)
						log.Println("------------------")
					}
				}
			}
		}

		// Save last handled block
		store.Set(LAST_BLOCK_KEY, bCount)
		lastHandledBlock++

		// Sleep for 1 minute
		if !backwardInTime {
			log.Println("-------------------------Wait for new block")
			time.Sleep(1 * time.Minute)
		}
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