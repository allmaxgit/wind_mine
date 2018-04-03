package eth

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"

	"WindToken/configs"
	uCrypto "WindToken/crypto"
	uErr "WindToken/errors"
	"WindToken/gocontracts"
	"WindToken/utils"
)

var (
	client *ethclient.Client

	auth      *bind.TransactOpts
	dAuth     *bind.TransactOpts
	session   *gocontracts.CrowdsaleSession
	euroCents *big.Int
)

// Dial connects to ETH provider create sessions for contracts.
func Dial(conf *configs.Crypto) (err error) {
	providerURL := utils.GetInfuraProviderUrl(conf.ETHNetworkId, conf.InfuraToken)

	client, _ = ethclient.Dial(providerURL)
	if err != nil {
		return
	}

	//gasPrice, err := client.SuggestGasPrice(context.TODO())
	//if err != nil {
	//	return uErr.Combine(err, "failed to get suggested gas price")
	//}

	// Create TransactOpts.
	auth, err = createAuth(conf.OwnerPrivateKey, &conf.GasPrice, conf.GasLimit)
	if err != nil {
		return uErr.Combine(err, "failed to create auth")
	}

	// Create Crowdsale session.
	session, err = createCrowdsaleSession(conf.CrowdsaleAddress)
	if err != nil {
		return uErr.Combine(err, "failed to create Crowdsale session")
	}

	err = prepareContracts(conf.OwnerAddress)
	if err != nil && err.Error() != uErr.ErrICOFinished {
		return uErr.Combine(err, "failed to prepare contracts")
	}

	return nil
}

// UpdateGasLimit asks for new gas limit value
// and sets it for TransactOpts.
func UpdateGasLimit() (err error) {
	gasLimit, err := getGasLimit()
	if err != nil {
		return
	}

	session.TransactOpts.GasLimit = gasLimit
	return
}

// GetTokenPrice determines token price in euro cents
// depending on ICO period.
//
// How match euro cents in one token.
func GetTokenPrice() (err error) {
	fmt.Println("Getting token price...")
	state, err := session.CrowdsaleState()
	if err != nil {
		return
	}

	if state == 0 {
		tx, err := session.CheckState()
		if err != nil {
			return err
		}

		receipt, err := getReceipt(tx, false)

		if receipt.Status == 0 {
			return errors.New(uErr.ErrReceiptStatus)
		}
	}

	state, err = session.CrowdsaleState()
	if err != nil {
		return uErr.Combine(err, "failed get Crowdsale state in second time")
	}

	switch state {
	case 0:
		return errors.New(uErr.ErrICONotStarted)
	case 1:
		euroCents, err = session.PrivatePriceInFiatFracture()
	case 2:
		euroCents, err = session.PreIcoPriceInFiatFracture()
	case 3:
		euroCents, err = session.IcoPriceInFiatFracture()
	case 4:
		return errors.New(uErr.ErrICOFinished)
	default:
		return errors.New(uErr.ErrUnknown)
	}
	return
}

// ConvertBTCToTokens converts BTC to ICO tokens.
func ConvertBTCToTokens(btcValue float64) *big.Int {
	tokens := new(big.Int)
	inEuroCents := new(big.Float).SetFloat64(btcValue * uCrypto.GetBTCRate() * 100 * math.Pow10(8))
	inEuroCents.Int(tokens)
	tokens.Quo(tokens, euroCents)

	return tokens
}

// SendTokens sends tokens to ETH address.
func SendTokens(addr string, amount *big.Int) error {
	err := utils.DoNTimeBeforeComplete(10, func(i int) (err error) {
		log.Println("session.TransactOpts.GasPrice:", session.TransactOpts.GasPrice)
		_, err = manualReserve(addr, amount)
		if err != nil {
			if err.Error() == core.ErrGasLimit.Error() {
				err = UpdateGasLimit()
				if err != nil {
					return uErr.Combine(err, uErr.ErrUpdateGasLimit)
				}
				return uErr.Combine(nil, "gas limit was updated")
			} else if err.Error() == core.ErrReplaceUnderpriced.Error() || err.Error() == uErr.ErrTXTimedOut {
				utils.IncreaseGasPrice(session.TransactOpts.GasPrice)
				return uErr.Combine(nil, "gas price was updated")
			}

			return err
		}

		return
	})

	return err
}

// ManualReserve sends tokens on particular address.
func manualReserve(addrStr string, amount *big.Int) (receipt *types.Receipt, err error) {
	addr := common.HexToAddress(addrStr)

	tx, err := session.ManualReserve(addr, amount)
	if err != nil {
		return
	}

	log.Println("manualReserve tx:", tx.Hash().Hex())

	receipt, err = getReceipt(tx, true)
	if err != nil {
		return
	}

	if receipt.Status == 0 {
		return nil, errors.New(uErr.ErrReceiptStatus)
	}

	return
}

//CheckIsKycPassed check in contract if address has passed KYC
func CheckIsKycPassed(addrStr string) (bool, error) {
	addr := common.HexToAddress(addrStr)

	whiteListed, err := session.WhiteList(addr)
	if err != nil {
		log.Println("whiteList() error:", err)
		return false, errors.New(uErr.ErrDataRetrieve)
	}

	return whiteListed, nil
}

func CheckCrowdsaleNotFinished() (bool, error) {
	state, err := session.CrowdsaleState()
	if err != nil {
		log.Println("Failed to get current crowdsale state")
		return false, errors.New(uErr.ErrICOFinished)
	}

	// Crowdsale state is described by enum, where 4 is FINISHED
	return state == 4, nil
}

func createAuth(key string, gasPrice *big.Int, gasLimit uint64) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	log.Println("GasPrice is:", gasPrice, "GasLimit is:", gasLimit)

	auth = bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	return
}

func createCrowdsaleSession(contractAddr string) (*gocontracts.CrowdsaleSession, error) {
	addr := common.HexToAddress(contractAddr)
	contract, err := gocontracts.NewCrowdsale(addr, client)
	if err != nil {
		return nil, err
	}

	session := &gocontracts.CrowdsaleSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: auth.GasPrice,
			GasLimit: auth.GasLimit,
			Value:    auth.Value,
		},
	}

	return session, nil
}

func prepareContracts(addrStr string) (err error) {
	// Check contracts owner.
	ownerAddress, err := session.Owner()
	if err != nil {
		return
	}

	if addrStr != ownerAddress.Hex() {
		return errors.New(uErr.ErrOwner)
	}

	// Check if contract not started or finished.
	err = GetTokenPrice()
	if err != nil {
		return
	}

	// Prepare contracts.
	err = prepareCrowdsale()
	return
}

func prepareCrowdsale() (err error) {
	fmt.Println("Prepare Crowdsale...")
	privateSaleHardCap, err := session.PrivateSaleHardCap()
	if err != nil {
		return uErr.Combine(err, "failed to get privateSaleHardCap")
	}

	if privateSaleHardCap.BitLen() == 0 {

		err := utils.DoNTimeBeforeComplete(10, func(i int) (err error) {
			fmt.Println("PrepareCrowdsale GasPrice:", session.TransactOpts.GasPrice) // TODO: Remove

			_, err = prepareCrowdsaleSync()
			if err != nil {
				if err.Error() == core.ErrGasLimit.Error() {
					err = UpdateGasLimit()
					if err != nil {
						return uErr.Combine(err, uErr.ErrUpdateGasLimit)
					}
					return uErr.Combine(nil, "gas limit was updated")
				} else if err.Error() == core.ErrReplaceUnderpriced.Error() || err.Error() == uErr.ErrTXTimedOut {
					utils.IncreaseGasPrice(session.TransactOpts.GasPrice)
					return uErr.Combine(nil, "gas price was updated")
				}

				return
			}

			return
		})

		if err != nil {
			return uErr.Combine(err, "failed to call prepareCrowdsale of contract")
		}

	}

	fmt.Println("Crowdsale preparation was completed")
	return nil
}

func getGasLimit() (uint64, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}

	return block.GasLimit(), nil
}

func getReceipt(tx *types.Transaction, useTimeout bool) (receipt *types.Receipt, err error) {
	ctxB := context.Background()
	var ctx context.Context
	var cancel context.CancelFunc

	if useTimeout {
		ctx, cancel = context.WithTimeout(ctxB, 60*time.Minute)
	} else {
		ctx, cancel = context.WithCancel(ctxB)
	}

	defer cancel()

	receipt, err = bind.WaitMined(ctx, client, tx)
	if ctx.Err() != nil && err == ctx.Err() {
		return nil, errors.New(uErr.ErrTXTimedOut)
	} else if err != nil {
		return nil, err
	}

	return
}

func prepareCrowdsaleSync() (receipt *types.Receipt, err error) {
	tx, err := session.PrepareCrowdsale()
	if err != nil {
		return
	}

	log.Println("tx:", tx.Hash())

	receipt, err = getReceipt(tx, true)
	if err != nil {
		return
	}

	if receipt.Status == 0 {
		return nil, errors.New(uErr.ErrReceiptStatus)
	}

	return
}
