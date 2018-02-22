package gotest

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/net/context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type User struct {
	Opts       *bind.TransactOpts
	PrivateKey string
}

func StartEthereumSimulator(count int) (*backends.SimulatedBackend, []*User) {
	params.TargetGasLimit = params.GenesisGasLimit

	genesisBlock := make(core.GenesisAlloc)
	initialBalance := new(big.Int)
	fmt.Sscan("1000000000000000000000000", initialBalance)

	accounts := make([]*User, count)

	for i := 0; i < count; i++ {
		key, _ := crypto.GenerateKey()
		hash := common.BigToHash(key.D)
		auth := bind.NewKeyedTransactor(key)
		auth.GasLimit = params.GenesisGasLimit
		auth.GasPrice = big.NewInt(1)
		genesisBlock[auth.From] = core.GenesisAccount{Balance: initialBalance}
		accounts[i] = &User{auth, hash.String()}
	}

	simulator := backends.NewSimulatedBackend(genesisBlock)
	simulator.Commit()

	return simulator, accounts
}

func GetReceipt(tx *types.Transaction, backend *backends.SimulatedBackend) *types.Receipt {
	backend.Commit()
	receipt, err := backend.TransactionReceipt(context.Background(), tx.Hash())
	//receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		panic("failed to get receipt")
	}
	return receipt
}

func sendTransaction(from *User, backend *backends.SimulatedBackend, to common.Address, value string) *types.Transaction {
	nonce, err := backend.NonceAt(context.Background(), from.Opts.From, nil)
	if err != nil {
		panic(err.Error())
	}
	gasprice, err := backend.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err.Error())
	}

	v := new(big.Int)
	_, err = fmt.Sscan(value, v)
	if err != nil {
		panic(err.Error())
	}

	tx := types.NewTransaction(
		nonce,
		to,
		v,
		params.GenesisGasLimit,
		gasprice,
		[]byte(`cooldata`),
	)
	signer := types.HomesteadSigner{}
	var signature []byte
	key, err := crypto.HexToECDSA(strings.TrimPrefix(from.PrivateKey, "0x"))
	if err != nil {
		panic(err)
	}
	signature, err = crypto.Sign(signer.Hash(tx).Bytes(), key)
	if err != nil {
		panic(err)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		panic(err)
	}
	if err := backend.SendTransaction(context.Background(), signedTx); err != nil {
		panic(err)
	}
	return signedTx
}

func DeployContract(user *User,
	backend *backends.SimulatedBackend,
	privateSaleStart,
	privateSaleDuration,
	preIcoDuration,
	icoDuration *big.Int,
	wallet, foundersWallet common.Address,
) common.Address {
	if user == nil || backend == nil || privateSaleStart == nil || privateSaleDuration == nil || preIcoDuration == nil || icoDuration == nil {
		panic("nil parameters")
	}
	addr, tx, _, err := DeployCrowdsale(user.Opts, backend, privateSaleStart, privateSaleDuration, preIcoDuration, icoDuration, wallet, foundersWallet)
	if err != nil {
		panic("failed to deploy crowdsale")
	}
	backend.Commit()
	receipt := GetReceipt(tx, backend)
	if receipt.Status != 1 {
		panic("crowdsale deploy tx failed")
	}
	return addr
}

func GetCrowdsaleSession(user *User, backend *backends.SimulatedBackend, contractAddress common.Address) *CrowdsaleSession {
	contract, err := NewCrowdsale(contractAddress, backend)
	if err != nil {
		panic("failed to get contract instance")
	}

	gasPrice, _ := backend.SuggestGasPrice(context.Background())
	session := &CrowdsaleSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    user.Opts.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     user.Opts.From,
			Signer:   user.Opts.Signer,
			Value:    user.Opts.Value,
			GasPrice: gasPrice,
			GasLimit: user.Opts.GasLimit,
		},
	}
	return session
}

func GetTokenSession(user *User, backend *backends.SimulatedBackend, crowdsaleAddress common.Address) *WindMineTokenSession {
	crowdsaleSession := GetCrowdsaleSession(user, backend, crowdsaleAddress)
	tokenAddress, err := crowdsaleSession.Token()
	if err != nil {
		return nil
	}
	contract, err := NewWindMineToken(tokenAddress, backend)
	if err != nil {
		panic("failed to get contract instance")
	}

	gasPrice, _ := backend.SuggestGasPrice(context.Background())
	session := &WindMineTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    user.Opts.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     user.Opts.From,
			Signer:   user.Opts.Signer,
			Value:    user.Opts.Value,
			GasPrice: gasPrice,
			GasLimit: user.Opts.GasLimit,
		},
	}
	return session
}
