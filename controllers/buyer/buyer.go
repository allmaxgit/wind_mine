package buyer

import (
	"errors"
	"strings"

	uErr "WindToken/errors"
	buyerGW "WindToken/api/buyer"
	"WindToken/configs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"golang.org/x/net/context"
)

type RPC struct {}

// AddBuyer adds new buyer.
func (r *RPC) GetBTCWallet(ctx context.Context, in *buyerGW.GetBTCWalletReq) (*buyerGW.GetBTCWalletResp, error) {
	ethAddrStr := in.EthAddress
	btcAddrStr := in.BtcAddress

	// Validate ETH address.
	if ethAddrStr == "" {
		return nil, errors.New(uErr.ErrorValidateETHAddress)
	}
	if !common.IsHexAddress(ethAddrStr) {
		return nil, errors.New(uErr.ErrorValidateETHAddress)
	}
	if common.EmptyHash(common.HexToHash(ethAddrStr)) {
		return nil, errors.New(uErr.ErrorValidateETHAddress)
	}
	ethAddr := common.HexToAddress(ethAddrStr)
	if strings.ToLower(ethAddrStr) != strings.ToLower(ethAddr.Hex()) {
		return nil, errors.New(uErr.ErrorValidateETHAddress)
	}

	// Define net configs.
	chainConf := &chaincfg.MainNetParams
	isTestnet := configs.GetConfigs().Crypto.BTCTestnet
	if isTestnet {
		chainConf = &chaincfg.TestNet3Params
	}

	// Validate BTC address.
	if btcAddrStr == "" {
		return nil, errors.New(uErr.ErrorValidateBTCAddress)
	}
	_, err := btcutil.DecodeAddress(btcAddrStr, chainConf)
	if err != nil {
		uErr.LogError(err, uErr.ErrorValidateBTCAddress)
		return nil, errors.New(uErr.ErrorValidateBTCAddress)
	}



	return nil, nil
}
