// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x6060604052341561000f57600080fd5b61025c8061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610080578063a9059cbb1461009f575b600080fd5b341561006657600080fd5b61006e6100d5565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e600160a060020a03600435166100db565b34156100aa57600080fd5b6100c1600160a060020a03600435166024356100f6565b604051901515815260200160405180910390f35b60015490565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561010d57600080fd5b600160a060020a03331660009081526020819052604090205482111561013257600080fd5b600160a060020a03331660009081526020819052604090205461015b908363ffffffff61020816565b600160a060020a033381166000908152602081905260408082209390935590851681522054610190908363ffffffff61021a16565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b60008282111561021457fe5b50900390565b60008282018381101561022957fe5b93925050505600a165627a7a72305820ff371cdc47f6d9a1d5f28c5d5f909e96652fa81653d15e8b210886903486e4580029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleABI is the input ABI used to generate the binding from.
const CrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"privateSaleHardCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoFinishDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_weiInFiat\",\"type\":\"uint256\"}],\"name\":\"updateWeiInFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFreezeTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"manualReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoFundsWithdrawn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"investorsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"setPreIcoStartDate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"setIcoStartDate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensSold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"setPrivateSaleStartDate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensOrdered\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preIcoPriceInFiatFracture\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoPriceInFiatFracture\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preIcoStartDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateSaleStartDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentHardCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"checkState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"receiveFrozenReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"addPrivateParticipant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateSaleWeiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"addToWhiteListMultiple\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalWhiteListed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privatePriceInFiatFracture\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"setIcoFinishDate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoWeiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiInFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInvestorsListLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preIcoHardCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preIcoWeiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"generalHardCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoHardCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveOrderedTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoStartDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawnState\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newWallet\",\"type\":\"address\"}],\"name\":\"setWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareCrowdsale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crowdsaleState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"preIcoFundsWithdrawn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frozenReserve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"privateSaleFundsWithdrawn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_privateSaleStartDate\",\"type\":\"uint256\"},{\"name\":\"_preIcoStartDate\",\"type\":\"uint256\"},{\"name\":\"_icoStartDate\",\"type\":\"uint256\"},{\"name\":\"_icoFinishDate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_foundersWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ManualTokenSend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldState\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_newState\",\"type\":\"uint8\"}],\"name\":\"StateHasChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_wallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_stage\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newWallet\",\"type\":\"address\"}],\"name\":\"WalletHasChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"PrivateSaleStartDateMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"PreIcoStartDateMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"IcoStartDateMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newDate\",\"type\":\"uint256\"}],\"name\":\"IcoFinishDateMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orderer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentWeiInFiat\",\"type\":\"uint256\"}],\"name\":\"TokensAreOrdered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_retriever\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokensAreRetrieved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_unfreezeTimestamp\",\"type\":\"uint256\"}],\"name\":\"FrozenReserveRetrieved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"whiteListedNum\",\"type\":\"uint256\"}],\"name\":\"LogWhiteListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"whiteListedNum\",\"type\":\"uint256\"}],\"name\":\"LogWhiteListedMultiple\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"RateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CrowdsaleBin is the compiled bytecode used for deploying new contracts.
const CrowdsaleBin = `0x60606040526016805460ff199081169091556018805482169055601a805482169055601c805460a060020a60ff021916905560228054909116905534156200004657600080fd5b60405160c080620032538339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519150604090508051908101604052600381527f4555520000000000000000000000000000000000000000000000000000000000602082015260008054600160a060020a03191633600160a060020a031617905560026001828051620000e7929160200190620001ce565b506002555060008611620000fa57600080fd5b8585116200010757600080fd5b8484116200011457600080fd5b600160a060020a03821615156200012a57600080fd5b600160a060020a03811615156200014057600080fd5b600f869055601085905560118490556012839055601c8054600160a060020a031916600160a060020a038416179055806200017a62000253565b600160a060020a039091168152602001604051809103906000f0801515620001a157600080fd5b601d8054600160a060020a031916600160a060020a03929092169190911790555062000284945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200021157805160ff191683800117855562000241565b8280016001018555821562000241579182015b828111156200024157825182559160200191906001019062000224565b506200024f92915062000264565b5090565b604051610a9880620027bb83390190565b6200028191905b808211156200024f57600081556001016200026b565b90565b61252780620002946000396000f30060606040526004361061026e5763ffffffff60e060020a6000350416630912b9b781146102795780630a3cb6631461029e5780630db111cb146102b15780631755bf3d146102c45780631abff816146102da5780631b0dbdaf146102ed57806327631d021461030f578063276ec457146103995780632ccc8727146103c0578063372c12b1146103f257806339c1904d146104115780633ccfd60b146104275780634042b66f1461043a57806347ee03941461044d5780634a1c13cd1461046c578063518ab2a814610482578063521eb2731461049557806353186c8c146104a85780635cac2917146104be57806369d3720c146104dd5780636ca27edc146104f057806379c3199d1461050357806386d2fe57146105165780638c585f5f146105295780638da5cb5b1461053c57806396dfcbea1461054f57806396f85f1614610562578063980a369114610581578063985585e1146105a05780639fec8e3b146105b3578063a07b206f146105d1578063a96d3141146105e4578063ae32ac45146105f7578063b12005691461060d578063b24a129014610620578063b2c560d214610633578063b62a86a814610646578063bf77566014610659578063d0b06f5d1461066c578063d122ca4e1461067f578063d33e225a14610692578063d3a106e5146106a5578063d73019e9146106b8578063db054c2f146106cb578063deaa59df146106e1578063e084a81914610700578063e7bb523314610713578063ec8ac4d81461074a578063efcdcb8b1461075e578063efe6980a14610771578063f2fde38b14610784578063f5890e46146107a3578063f80b99ef146107b6578063fc0c546a146107c9575b610277336107dc565b005b341561028457600080fd5b61028c610d0b565b60405190815260200160405180910390f35b34156102a957600080fd5b61028c610d11565b34156102bc57600080fd5b61028c610d19565b34156102cf57600080fd5b610277600435610d1f565b34156102e557600080fd5b61028c610d8c565b34156102f857600080fd5b610277600160a060020a0360043516602435610d92565b341561031a57600080fd5b610322611045565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561035e578082015183820152602001610346565b50505050905090810190601f16801561038b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103a457600080fd5b6103ac6110e3565b604051901515815260200160405180910390f35b34156103cb57600080fd5b6103d66004356110ec565b604051600160a060020a03909116815260200160405180910390f35b34156103fd57600080fd5b6103ac600160a060020a0360043516611114565b341561041c57600080fd5b610277600435611128565b341561043257600080fd5b6102776111f0565b341561044557600080fd5b61028c611570565b341561045857600080fd5b6103ac600160a060020a0360043516611576565b341561047757600080fd5b610277600435611647565b341561048d57600080fd5b61028c611712565b34156104a057600080fd5b6103d6611718565b34156104b357600080fd5b610277600435611727565b34156104c957600080fd5b61028c600160a060020a03600435166117bd565b34156104e857600080fd5b61028c6117cf565b34156104fb57600080fd5b61028c6117d5565b341561050e57600080fd5b61028c6117db565b341561052157600080fd5b61028c6117e1565b341561053457600080fd5b61028c6117e7565b341561054757600080fd5b6103d66117ed565b341561055a57600080fd5b6102776117fc565b341561056d57600080fd5b610277600160a060020a0360043516611b05565b341561058c57600080fd5b610277600160a060020a0360043516611caa565b34156105ab57600080fd5b61028c611cfe565b34156105be57600080fd5b6103ac6004803560248101910135611d04565b34156105dc57600080fd5b61028c611e12565b34156105ef57600080fd5b61028c611e18565b341561060257600080fd5b610277600435611e1e565b341561061857600080fd5b61028c611eb5565b341561062b57600080fd5b61028c611ebb565b341561063e57600080fd5b61028c611ec1565b341561065157600080fd5b61028c611ec8565b341561066457600080fd5b61028c611ece565b341561067757600080fd5b61028c611ed4565b341561068a57600080fd5b61028c611eda565b341561069d57600080fd5b61028c611ee0565b34156106b057600080fd5b610277611ee6565b34156106c357600080fd5b61028c612050565b34156106d657600080fd5b6103ac600435612056565b34156106ec57600080fd5b610277600160a060020a036004351661207d565b341561070b57600080fd5b610277612146565b341561071e57600080fd5b61072661234b565b6040518082600481111561073657fe5b60ff16815260200191505060405180910390f35b610277600160a060020a03600435166107dc565b341561076957600080fd5b6103ac61235b565b341561077c57600080fd5b61028c612364565b341561078f57600080fd5b610277600160a060020a036004351661236c565b34156107ae57600080fd5b61028c612407565b34156107c157600080fd5b6103ac61240d565b34156107d457600080fd5b6103d6612416565b6022546000908190819081908190819060ff16156107f957600080fd5b6022805460ff1916600117905561080e6117fc565b600160a060020a038716151561082357600080fd5b6000601c5460a060020a900460ff16600481111561083d57fe5b141561084857600080fd5b6004601c5460a060020a900460ff16600481111561086257fe5b141561086d57600080fd5b67016345785d8a000034101561088257600080fd5b6003546000901161089257600080fd5b6001601c5460a060020a900460ff1660048111156108ac57fe5b14156108d957600160a060020a03871660009081526021602052604090205460ff1615156108d957600080fd5b600160a060020a038716600090815260208052604090205460ff1615156108ff57600080fd5b60035434965061091690879063ffffffff61242516565b94506001601c5460a060020a900460ff16600481111561093257fe5b1415610942576007549250610992565b6002601c5460a060020a900460ff16600481111561095c57fe5b141561096c576008549250610992565b6003601c5460a060020a900460ff16600481111561098657fe5b14156109925760095492505b6109a2858463ffffffff61242516565b9350600084116109b157600080fd5b600b546013546109c7908663ffffffff61244116565b1115610a2157601354600b546109e29163ffffffff61245716565b600354909450610a08906109fc868663ffffffff61246916565b9063ffffffff61246916565b9050610a1a868263ffffffff61245716565b9150610a24565b50845b601354610a37908563ffffffff61244116565b601355601454610a4d908263ffffffff61244116565b6014556001601c5460a060020a900460ff166004811115610a6a57fe5b1415610a8b57601554610a83908263ffffffff61244116565b601555610afd565b6002601c5460a060020a900460ff166004811115610aa557fe5b1415610ac657601754610abe908263ffffffff61244116565b601755610afd565b6003601c5460a060020a900460ff166004811115610ae057fe5b1415610afd57601954610af9908263ffffffff61244116565b6019555b601d54600160a060020a03166370a082318860405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b4d57600080fd5b5af11515610b5a57600080fd5b505050604051805115159050610baf576006805460018101610b7c8382612494565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0389161790555b6000821115610be957600160a060020a03871682156108fc0283604051600060405180830381858888f193505050501515610be957600080fd5b601d54610c5290600160a060020a031663313ce5676040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610c2c57600080fd5b5af11515610c3957600080fd5b5050506040518051869150600a0a63ffffffff61246916565b600160a060020a0388166000908152601e6020526040902054909450610c7e908563ffffffff61244116565b600160a060020a0388166000908152601e602052604090819020919091556003547fb794cb9cc471caae752d175fddbc2f267fe684f1a107b2f19f44a95e602648b69189918791518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a150506022805460ff191690555050505050565b600c5481565b6301e1338081565b60125481565b60005433600160a060020a03908116911614610d3a57600080fd5b60008111610d4757600080fd5b7fb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d19166003548260405191825260208201526040908101905180910390a160035542600455565b60055481565b60005433600160a060020a03908116911614610dad57600080fd5b60225460ff1615610dbd57600080fd5b6022805460ff19166001179055610dd26117fc565b600160a060020a0382161515610de757600080fd5b60008111610df457600080fd5b6000601c5460a060020a900460ff166004811115610e0e57fe5b14158015610e3457506004601c5460a060020a900460ff166004811115610e3157fe5b14155b1515610e3f57600080fd5b600b54601354610e55908363ffffffff61244116565b10610e5f57600080fd5b6001601c5460a060020a900460ff166004811115610e7957fe5b1415610ea657600160a060020a03821660009081526021602052604090205460ff161515610ea657600080fd5b600160a060020a038216600090815260208052604090205460ff161515610ecc57600080fd5b601d54600160a060020a03166370a082318360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610f1c57600080fd5b5af11515610f2957600080fd5b505050604051805115159050610f7e576006805460018101610f4b8382612494565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b601354610f91908263ffffffff61244116565b601355600160a060020a0382166000908152601e6020526040902054610fbd908263ffffffff61244116565b600160a060020a0383166000908152601e602052604090819020919091556003547fb794cb9cc471caae752d175fddbc2f267fe684f1a107b2f19f44a95e602648b69184918491518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a150506022805460ff19169055565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110db5780601f106110b0576101008083540402835291602001916110db565b820191906000526020600020905b8154815290600101906020018083116110be57829003601f168201915b505050505081565b601a5460ff1681565b60068054829081106110fa57fe5b600091825260209091200154600160a060020a0316905081565b602080526000908152604090205460ff1681565b60005433600160a060020a0390811691161461114357600080fd5b61114b6117fc565b6000601c5460a060020a900460ff16600481111561116557fe5b148061118857506001601c5460a060020a900460ff16600481111561118657fe5b145b151561119357600080fd5b600f5481116111a157600080fd5b60115481106111af57600080fd5b7f9294d04a272cdf981934ca76de43865446ea70e949478d270dad1e6d90b5960e6010548260405191825260208201526040908101905180910390a1601055565b60005433600160a060020a0390811691161461120b57600080fd5b60225460ff161561121b57600080fd5b6022805460ff191660011790556112306117fc565b6000601c5460a060020a900460ff16600481111561124a57fe5b1161125457600080fd5b601b6112826001601c60149054906101000a900460ff16600481111561127657fe5b9063ffffffff61245716565b6004811061128c57fe5b602081049091015460ff601f9092166101000a900416156112ac57600080fd5b6001601c5460a060020a900460ff1660048111156112c657fe5b1180156112db5750601b54610100900460ff16155b1561139157601b805461ff001916610100179055601c546015547ffbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb491600160a060020a03169060016040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1601c54601554600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561139157600080fd5b6002601c5460a060020a900460ff1660048111156113ab57fe5b1180156113c15750601b5462010000900460ff16155b1561147957601b805462ff0000191662010000179055601c546017547ffbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb491600160a060020a03169060026040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1601c54601754600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561147957600080fd5b6003601c5460a060020a900460ff16600481111561149357fe5b1180156114aa5750601b546301000000900460ff16155b1561156457601b805463ff00000019166301000000179055601c546019547ffbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb491600160a060020a03169060036040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1601c54601954600160a060020a039091169080156108fc0290604051600060405180830381858888f19350505050151561156457600080fd5b6022805460ff19169055565b60145481565b6000805433600160a060020a0390811691161461159257600080fd5b600160a060020a03821615156115a757600080fd5b600160a060020a038216600090815260208052604090205460ff16151561163f57600160a060020a038216600090815260208052604090819020805460ff19166001908117909155601f8054909101908190557f88eb615cbff4540422d181389333a5c75e3d5eb98dd55fb176e3a615ddfd0f1f9184919051600160a060020a03909216825260208201526040908101905180910390a15b506001919050565b60005433600160a060020a0390811691161461166257600080fd5b61166a6117fc565b6003601c5460a060020a900460ff16600481111561168457fe5b141580156116aa57506004601c5460a060020a900460ff1660048111156116a757fe5b14155b15156116b557600080fd5b60105481116116c357600080fd5b60125481106116d157600080fd5b7f8d20f8b7089681de055ef410b1e22f13b3a44b8cda9145df73902b0505115c1c6011548260405191825260208201526040908101905180910390a1601155565b60135481565b601c54600160a060020a031681565b60005433600160a060020a0390811691161461174257600080fd5b61174a6117fc565b6000601c5460a060020a900460ff16600481111561176457fe5b1461176e57600080fd5b601054811061177c57600080fd5b7fefe29dae758f051a9c039b9f34ee371f674e06edf12f196cfeda8a42131fdfe5600f548260405191825260208201526040908101905180910390a1600f55565b601e6020526000908152604090205481565b60085481565b60095481565b60105481565b600f5481565b600b5481565b600054600160a060020a031681565b600f5442101561182957601c805474ff000000000000000000000000000000000000000019169055611b03565b600f54421015801561183c575060105442105b156118db576001601c5460a060020a900460ff16600481111561185b57fe5b146118d6576000805160206124dc833981519152600060016040518083600481111561188357fe5b60ff16815260200182600481111561189757fe5b60ff1681526020019250505060405180910390a1601c805474ff0000000000000000000000000000000000000000191660a060020a179055600c54600b555b611b03565b60105442101580156118ee575060115442105b1561199d576002601c5460a060020a900460ff16600481111561190d57fe5b146118d6576000805160206124dc833981519152600160026040518083600481111561193557fe5b60ff16815260200182600481111561194957fe5b60ff1681526020019250505060405180910390a1601c805474ff0000000000000000000000000000000000000000191674020000000000000000000000000000000000000000179055600d54600b55611b03565b60115442101580156119b0575060125442105b15611a5f576003601c5460a060020a900460ff1660048111156119cf57fe5b146118d6576000805160206124dc83398151915260026003604051808360048111156119f757fe5b60ff168152602001826004811115611a0b57fe5b60ff1681526020019250505060405180910390a1601c805474ff0000000000000000000000000000000000000000191674030000000000000000000000000000000000000000179055600e54600b55611b03565b6004601c5460a060020a900460ff166004811115611a7957fe5b14611b03576000805160206124dc8339815191526003600460405180836004811115611aa157fe5b60ff168152602001826004811115611ab557fe5b60ff1681526020019250505060405180910390a1601c805474ff0000000000000000000000000000000000000000191674040000000000000000000000000000000000000000179055426005555b565b60005433600160a060020a03908116911614611b2057600080fd5b60225460ff1615611b3057600080fd5b6022805460ff191660011790556004601c5460a060020a900460ff166004811115611b5757fe5b14611b6157600080fd5b6005546301e1338001421015611b7657600080fd5b600160a060020a0381161515611b8b57600080fd5b601d54600160a060020a031663a9059cbb82611bfe8363313ce5676040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611bd457600080fd5b5af11515611be157600080fd5b50505060405180516304c4b4009150600a0a63ffffffff61246916565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515611c4157600080fd5b5af11515611c4e57600080fd5b50505060405180519050507f39c2ff8a62bebaa5cd4af157ba107e234024dabe84f3e6f2cc416f50b7f380ab8142604051600160a060020a03909216825260208201526040908101905180910390a1506022805460ff19169055565b60005433600160a060020a03908116911614611cc557600080fd5b600160a060020a0381161515611cda57600080fd5b600160a060020a03166000908152602160205260409020805460ff19166001179055565b60155481565b600080548190819033600160a060020a03908116911614611d2457600080fd5b8391506096821115611d3557600080fd5b5060005b81811015611dd25760206000868684818110611d5157fe5b60209081029290920135600160a060020a03168352508101919091526040016000205460ff161515611dca57600160206000878785818110611d8f57fe5b60209081029290920135600160a060020a0316835250810191909152604001600020805460ff1916911515919091179055601f805460010190555b600101611d39565b7f9075b10cfdcbb4feaac9b18054833db2bdacc10263b692174de5b561d34e2f4f601f5460405190815260200160405180910390a1506001949350505050565b601f5481565b60075481565b60005433600160a060020a03908116911614611e3957600080fd5b611e416117fc565b6004601c5460a060020a900460ff166004811115611e5b57fe5b1415611e6657600080fd5b6011548111611e7457600080fd5b7f47c07aea40f2e4306adb071f865fcd76a81d353c213ea7eef819581989ff9f676012548260405191825260208201526040908101905180910390a1601255565b60195481565b60035481565b6006545b90565b600d5481565b60175481565b60045481565b600a5481565b600e5481565b60225460009060ff1615611ef957600080fd5b6022805460ff191660011790556004601c5460a060020a900460ff166004811115611f2057fe5b14611f2a57600080fd5b600160a060020a0333166000908152601e602052604081205411611f4d57600080fd5b600160a060020a033316600090815260208052604090205460ff161515611f7357600080fd5b50600160a060020a03338181166000908152601e60205260408082208054929055601d549193919091169163a9059cbb9184905160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515611fe757600080fd5b5af11515611ff457600080fd5b50505060405180519050507ff33726a859f09868cbfde46ebe9962f051984065979b92e9a7533b202a3c490e3382604051600160a060020a03909216825260208201526040908101905180910390a1506022805460ff19169055565b60115481565b601b816004811061206357fe5b60209182820401919006915054906101000a900460ff1681565b60005433600160a060020a0390811691161461209857600080fd5b6120a06117fc565b6004601c5460a060020a900460ff1660048111156120ba57fe5b14156120c557600080fd5b601c547f86b557e590894ebca3dce4dfc61800ecb963722c9cce83d7d0d1dd32e322420a90600160a060020a031682604051600160a060020a039283168152911660208201526040908101905180910390a1601c805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461216157600080fd5b601d54600160a060020a03166318160ddd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156121a057600080fd5b5af115156121ad57600080fd5b5050506040518051600a5550601d54600160a060020a031663313ce5676040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156121f857600080fd5b5af1151561220557600080fd5b5050506040518051600a0a6298968002600c819055601d54909150600160a060020a031663313ce5676040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561225c57600080fd5b5af1151561226957600080fd5b5050506040518051600a0a6301312d000291909101600d819055601d54909150600160a060020a031663313ce5676040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156122c557600080fd5b5af115156122d257600080fd5b50505060405180519050600a0a6301c9c3800201600e81905550612316606461230a6019600254600a0a61246990919063ffffffff16565b9063ffffffff61242516565b6007556002546123379060649061230a90600a0a603263ffffffff61246916565b600855600254600a0a600955611b036117fc565b601c5460a060020a900460ff1681565b60185460ff1681565b6304c4b40081565b60005433600160a060020a0390811691161461238757600080fd5b600160a060020a038116151561239c57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025481565b60165460ff1681565b601d54600160a060020a031681565b600080828481151561243357fe5b0490508091505b5092915050565b60008282018381101561245057fe5b9392505050565b60008282111561246357fe5b50900390565b60008083151561247c576000915061243a565b5082820282848281151561248c57fe5b041461245057fe5b8154818355818115116124b8576000838152602090206124b89181019083016124bd565b505050565b611ec591905b808211156124d757600081556001016124c3565b5090560040bc5f093899d560cbd784d85b697b577e63a2991330daf682bd4d6ade755b03a165627a7a72305820755955009c91a1878457b1efcc7769e23f21dc276a0d78c5e51d419b42d6a0f800296060604052341561000f57600080fd5b604051602080610a988339810160405280805160038054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b6638d7ea4c680000600181905560048054600160a060020a031916600160a060020a03841617905561009b9066071afd498d00006401000000006100d281026109511704565b600160a060020a033381166000908152602081905260408082209390935560045490911681522066071afd498d00009055506100e4565b6000828211156100de57fe5b50900390565b6109a5806100f36000396000f3006060604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100df578063095ea7b31461016957806318160ddd1461019f5780631bfaf155146101c457806323b872dd146101f3578063313ce5671461021b578063661884631461022e57806370a08231146102505780637b86120a1461026f5780638da5cb5b1461028257806395d89b4114610295578063a9059cbb146102a8578063d73dd623146102ca578063dd62ed3e146102ec578063f2fde38b14610311575b600080fd5b34156100ea57600080fd5b6100f2610332565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012e578082015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017457600080fd5b61018b600160a060020a0360043516602435610369565b604051901515815260200160405180910390f35b34156101aa57600080fd5b6101b26103d5565b60405190815260200160405180910390f35b34156101cf57600080fd5b6101d76103db565b604051600160a060020a03909116815260200160405180910390f35b34156101fe57600080fd5b61018b600160a060020a03600435811690602435166044356103ea565b341561022657600080fd5b6101b261056a565b341561023957600080fd5b61018b600160a060020a036004351660243561056f565b341561025b57600080fd5b6101b2600160a060020a0360043516610669565b341561027a57600080fd5b6101b2610684565b341561028d57600080fd5b6101d761068f565b34156102a057600080fd5b6100f261069e565b34156102b357600080fd5b61018b600160a060020a03600435166024356106d5565b34156102d557600080fd5b61018b600160a060020a03600435166024356107e7565b34156102f757600080fd5b6101b2600160a060020a036004358116906024351661088b565b341561031c57600080fd5b610330600160a060020a03600435166108b6565b005b60408051908101604052600881527f57696e644d696e65000000000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b600454600160a060020a031681565b6000600160a060020a038316151561040157600080fd5b600160a060020a03841660009081526020819052604090205482111561042657600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561045957600080fd5b600160a060020a038416600090815260208190526040902054610482908363ffffffff61095116565b600160a060020a0380861660009081526020819052604080822093909355908516815220546104b7908363ffffffff61096316565b600160a060020a03808516600090815260208181526040808320949094558783168252600281528382203390931682529190915220546104fd908363ffffffff61095116565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600881565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105cc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610603565b6105dc818463ffffffff61095116565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b66071afd498d000081565b600354600160a060020a031681565b60408051908101604052600381527f574d440000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a03831615156106ec57600080fd5b600160a060020a03331660009081526020819052604090205482111561071157600080fd5b600160a060020a03331660009081526020819052604090205461073a908363ffffffff61095116565b600160a060020a03338116600090815260208190526040808220939093559085168152205461076f908363ffffffff61096316565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a03338116600090815260026020908152604080832093861683529290529081205461081f908363ffffffff61096316565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146108d157600080fd5b600160a060020a03811615156108e657600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561095d57fe5b50900390565b60008282018381101561097257fe5b93925050505600a165627a7a7230582000d0f9300042a3129708c6249509a1f62188bf3c6e604dfdda625c084f313a0d0029`

// DeployCrowdsale deploys a new Ethereum contract, binding an instance of Crowdsale to it.
func DeployCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _privateSaleStartDate *big.Int, _preIcoStartDate *big.Int, _icoStartDate *big.Int, _icoFinishDate *big.Int, _wallet common.Address, _foundersWallet common.Address) (common.Address, *types.Transaction, *Crowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdsaleBin), backend, _privateSaleStartDate, _preIcoStartDate, _icoStartDate, _icoFinishDate, _wallet, _foundersWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}, CrowdsaleFilterer: CrowdsaleFilterer{contract: contract}}, nil
}

// Crowdsale is an auto generated Go binding around an Ethereum contract.
type Crowdsale struct {
	CrowdsaleCaller     // Read-only binding to the contract
	CrowdsaleTransactor // Write-only binding to the contract
	CrowdsaleFilterer   // Log filterer for contract events
}

// CrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdsaleSession struct {
	Contract     *Crowdsale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdsaleCallerSession struct {
	Contract *CrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdsaleTransactorSession struct {
	Contract     *CrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdsaleRaw struct {
	Contract *Crowdsale // Generic contract binding to access the raw methods on
}

// CrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdsaleCallerRaw struct {
	Contract *CrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdsaleTransactorRaw struct {
	Contract *CrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdsale creates a new instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsale(address common.Address, backend bind.ContractBackend) (*Crowdsale, error) {
	contract, err := bindCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}, CrowdsaleFilterer: CrowdsaleFilterer{contract: contract}}, nil
}

// NewCrowdsaleCaller creates a new read-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CrowdsaleCaller, error) {
	contract, err := bindCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleCaller{contract: contract}, nil
}

// NewCrowdsaleTransactor creates a new write-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdsaleTransactor, error) {
	contract, err := bindCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTransactor{contract: contract}, nil
}

// NewCrowdsaleFilterer creates a new log filterer instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdsaleFilterer, error) {
	contract, err := bindCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleFilterer{contract: contract}, nil
}

// bindCrowdsale binds a generic wrapper to an already deployed contract.
func bindCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.CrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transact(opts, method, params...)
}

// CrowdsaleState is a free data retrieval call binding the contract method 0xe7bb5233.
//
// Solidity: function crowdsaleState() constant returns(uint8)
func (_Crowdsale *CrowdsaleCaller) CrowdsaleState(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "crowdsaleState")
	return *ret0, err
}

// CrowdsaleState is a free data retrieval call binding the contract method 0xe7bb5233.
//
// Solidity: function crowdsaleState() constant returns(uint8)
func (_Crowdsale *CrowdsaleSession) CrowdsaleState() (uint8, error) {
	return _Crowdsale.Contract.CrowdsaleState(&_Crowdsale.CallOpts)
}

// CrowdsaleState is a free data retrieval call binding the contract method 0xe7bb5233.
//
// Solidity: function crowdsaleState() constant returns(uint8)
func (_Crowdsale *CrowdsaleCallerSession) CrowdsaleState() (uint8, error) {
	return _Crowdsale.Contract.CrowdsaleState(&_Crowdsale.CallOpts)
}

// CurrentHardCap is a free data retrieval call binding the contract method 0x8c585f5f.
//
// Solidity: function currentHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) CurrentHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "currentHardCap")
	return *ret0, err
}

// CurrentHardCap is a free data retrieval call binding the contract method 0x8c585f5f.
//
// Solidity: function currentHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) CurrentHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.CurrentHardCap(&_Crowdsale.CallOpts)
}

// CurrentHardCap is a free data retrieval call binding the contract method 0x8c585f5f.
//
// Solidity: function currentHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) CurrentHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.CurrentHardCap(&_Crowdsale.CallOpts)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) FiatDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "fiatDecimals")
	return *ret0, err
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) FiatDecimals() (*big.Int, error) {
	return _Crowdsale.Contract.FiatDecimals(&_Crowdsale.CallOpts)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) FiatDecimals() (*big.Int, error) {
	return _Crowdsale.Contract.FiatDecimals(&_Crowdsale.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_Crowdsale *CrowdsaleCaller) FiatSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "fiatSymbol")
	return *ret0, err
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_Crowdsale *CrowdsaleSession) FiatSymbol() (string, error) {
	return _Crowdsale.Contract.FiatSymbol(&_Crowdsale.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_Crowdsale *CrowdsaleCallerSession) FiatSymbol() (string, error) {
	return _Crowdsale.Contract.FiatSymbol(&_Crowdsale.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) FreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "freezePeriod")
	return *ret0, err
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) FreezePeriod() (*big.Int, error) {
	return _Crowdsale.Contract.FreezePeriod(&_Crowdsale.CallOpts)
}

// FreezePeriod is a free data retrieval call binding the contract method 0x0a3cb663.
//
// Solidity: function freezePeriod() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) FreezePeriod() (*big.Int, error) {
	return _Crowdsale.Contract.FreezePeriod(&_Crowdsale.CallOpts)
}

// FrozenReserve is a free data retrieval call binding the contract method 0xefe6980a.
//
// Solidity: function frozenReserve() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) FrozenReserve(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "frozenReserve")
	return *ret0, err
}

// FrozenReserve is a free data retrieval call binding the contract method 0xefe6980a.
//
// Solidity: function frozenReserve() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) FrozenReserve() (*big.Int, error) {
	return _Crowdsale.Contract.FrozenReserve(&_Crowdsale.CallOpts)
}

// FrozenReserve is a free data retrieval call binding the contract method 0xefe6980a.
//
// Solidity: function frozenReserve() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) FrozenReserve() (*big.Int, error) {
	return _Crowdsale.Contract.FrozenReserve(&_Crowdsale.CallOpts)
}

// GeneralHardCap is a free data retrieval call binding the contract method 0xd122ca4e.
//
// Solidity: function generalHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) GeneralHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "generalHardCap")
	return *ret0, err
}

// GeneralHardCap is a free data retrieval call binding the contract method 0xd122ca4e.
//
// Solidity: function generalHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) GeneralHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.GeneralHardCap(&_Crowdsale.CallOpts)
}

// GeneralHardCap is a free data retrieval call binding the contract method 0xd122ca4e.
//
// Solidity: function generalHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) GeneralHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.GeneralHardCap(&_Crowdsale.CallOpts)
}

// GetInvestorsListLength is a free data retrieval call binding the contract method 0xb2c560d2.
//
// Solidity: function getInvestorsListLength() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) GetInvestorsListLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "getInvestorsListLength")
	return *ret0, err
}

// GetInvestorsListLength is a free data retrieval call binding the contract method 0xb2c560d2.
//
// Solidity: function getInvestorsListLength() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) GetInvestorsListLength() (*big.Int, error) {
	return _Crowdsale.Contract.GetInvestorsListLength(&_Crowdsale.CallOpts)
}

// GetInvestorsListLength is a free data retrieval call binding the contract method 0xb2c560d2.
//
// Solidity: function getInvestorsListLength() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) GetInvestorsListLength() (*big.Int, error) {
	return _Crowdsale.Contract.GetInvestorsListLength(&_Crowdsale.CallOpts)
}

// IcoFinishDate is a free data retrieval call binding the contract method 0x0db111cb.
//
// Solidity: function icoFinishDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) IcoFinishDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoFinishDate")
	return *ret0, err
}

// IcoFinishDate is a free data retrieval call binding the contract method 0x0db111cb.
//
// Solidity: function icoFinishDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) IcoFinishDate() (*big.Int, error) {
	return _Crowdsale.Contract.IcoFinishDate(&_Crowdsale.CallOpts)
}

// IcoFinishDate is a free data retrieval call binding the contract method 0x0db111cb.
//
// Solidity: function icoFinishDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) IcoFinishDate() (*big.Int, error) {
	return _Crowdsale.Contract.IcoFinishDate(&_Crowdsale.CallOpts)
}

// IcoFundsWithdrawn is a free data retrieval call binding the contract method 0x276ec457.
//
// Solidity: function icoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) IcoFundsWithdrawn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoFundsWithdrawn")
	return *ret0, err
}

// IcoFundsWithdrawn is a free data retrieval call binding the contract method 0x276ec457.
//
// Solidity: function icoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleSession) IcoFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.IcoFundsWithdrawn(&_Crowdsale.CallOpts)
}

// IcoFundsWithdrawn is a free data retrieval call binding the contract method 0x276ec457.
//
// Solidity: function icoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) IcoFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.IcoFundsWithdrawn(&_Crowdsale.CallOpts)
}

// IcoHardCap is a free data retrieval call binding the contract method 0xd33e225a.
//
// Solidity: function icoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) IcoHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoHardCap")
	return *ret0, err
}

// IcoHardCap is a free data retrieval call binding the contract method 0xd33e225a.
//
// Solidity: function icoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) IcoHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.IcoHardCap(&_Crowdsale.CallOpts)
}

// IcoHardCap is a free data retrieval call binding the contract method 0xd33e225a.
//
// Solidity: function icoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) IcoHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.IcoHardCap(&_Crowdsale.CallOpts)
}

// IcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x6ca27edc.
//
// Solidity: function icoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) IcoPriceInFiatFracture(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoPriceInFiatFracture")
	return *ret0, err
}

// IcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x6ca27edc.
//
// Solidity: function icoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) IcoPriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.IcoPriceInFiatFracture(&_Crowdsale.CallOpts)
}

// IcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x6ca27edc.
//
// Solidity: function icoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) IcoPriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.IcoPriceInFiatFracture(&_Crowdsale.CallOpts)
}

// IcoStartDate is a free data retrieval call binding the contract method 0xd73019e9.
//
// Solidity: function icoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) IcoStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoStartDate")
	return *ret0, err
}

// IcoStartDate is a free data retrieval call binding the contract method 0xd73019e9.
//
// Solidity: function icoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) IcoStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.IcoStartDate(&_Crowdsale.CallOpts)
}

// IcoStartDate is a free data retrieval call binding the contract method 0xd73019e9.
//
// Solidity: function icoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) IcoStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.IcoStartDate(&_Crowdsale.CallOpts)
}

// IcoWeiRaised is a free data retrieval call binding the contract method 0xb1200569.
//
// Solidity: function icoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) IcoWeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "icoWeiRaised")
	return *ret0, err
}

// IcoWeiRaised is a free data retrieval call binding the contract method 0xb1200569.
//
// Solidity: function icoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) IcoWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.IcoWeiRaised(&_Crowdsale.CallOpts)
}

// IcoWeiRaised is a free data retrieval call binding the contract method 0xb1200569.
//
// Solidity: function icoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) IcoWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.IcoWeiRaised(&_Crowdsale.CallOpts)
}

// InvestorsList is a free data retrieval call binding the contract method 0x2ccc8727.
//
// Solidity: function investorsList( uint256) constant returns(address)
func (_Crowdsale *CrowdsaleCaller) InvestorsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "investorsList", arg0)
	return *ret0, err
}

// InvestorsList is a free data retrieval call binding the contract method 0x2ccc8727.
//
// Solidity: function investorsList( uint256) constant returns(address)
func (_Crowdsale *CrowdsaleSession) InvestorsList(arg0 *big.Int) (common.Address, error) {
	return _Crowdsale.Contract.InvestorsList(&_Crowdsale.CallOpts, arg0)
}

// InvestorsList is a free data retrieval call binding the contract method 0x2ccc8727.
//
// Solidity: function investorsList( uint256) constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) InvestorsList(arg0 *big.Int) (common.Address, error) {
	return _Crowdsale.Contract.InvestorsList(&_Crowdsale.CallOpts, arg0)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) LastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "lastUpdated")
	return *ret0, err
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) LastUpdated() (*big.Int, error) {
	return _Crowdsale.Contract.LastUpdated(&_Crowdsale.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) LastUpdated() (*big.Int, error) {
	return _Crowdsale.Contract.LastUpdated(&_Crowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Crowdsale *CrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Crowdsale *CrowdsaleSession) Owner() (common.Address, error) {
	return _Crowdsale.Contract.Owner(&_Crowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Owner() (common.Address, error) {
	return _Crowdsale.Contract.Owner(&_Crowdsale.CallOpts)
}

// PreIcoFundsWithdrawn is a free data retrieval call binding the contract method 0xefcdcb8b.
//
// Solidity: function preIcoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) PreIcoFundsWithdrawn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "preIcoFundsWithdrawn")
	return *ret0, err
}

// PreIcoFundsWithdrawn is a free data retrieval call binding the contract method 0xefcdcb8b.
//
// Solidity: function preIcoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleSession) PreIcoFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.PreIcoFundsWithdrawn(&_Crowdsale.CallOpts)
}

// PreIcoFundsWithdrawn is a free data retrieval call binding the contract method 0xefcdcb8b.
//
// Solidity: function preIcoFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) PreIcoFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.PreIcoFundsWithdrawn(&_Crowdsale.CallOpts)
}

// PreIcoHardCap is a free data retrieval call binding the contract method 0xb62a86a8.
//
// Solidity: function preIcoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PreIcoHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "preIcoHardCap")
	return *ret0, err
}

// PreIcoHardCap is a free data retrieval call binding the contract method 0xb62a86a8.
//
// Solidity: function preIcoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PreIcoHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoHardCap(&_Crowdsale.CallOpts)
}

// PreIcoHardCap is a free data retrieval call binding the contract method 0xb62a86a8.
//
// Solidity: function preIcoHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PreIcoHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoHardCap(&_Crowdsale.CallOpts)
}

// PreIcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x69d3720c.
//
// Solidity: function preIcoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PreIcoPriceInFiatFracture(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "preIcoPriceInFiatFracture")
	return *ret0, err
}

// PreIcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x69d3720c.
//
// Solidity: function preIcoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PreIcoPriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoPriceInFiatFracture(&_Crowdsale.CallOpts)
}

// PreIcoPriceInFiatFracture is a free data retrieval call binding the contract method 0x69d3720c.
//
// Solidity: function preIcoPriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PreIcoPriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoPriceInFiatFracture(&_Crowdsale.CallOpts)
}

// PreIcoStartDate is a free data retrieval call binding the contract method 0x79c3199d.
//
// Solidity: function preIcoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PreIcoStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "preIcoStartDate")
	return *ret0, err
}

// PreIcoStartDate is a free data retrieval call binding the contract method 0x79c3199d.
//
// Solidity: function preIcoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PreIcoStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoStartDate(&_Crowdsale.CallOpts)
}

// PreIcoStartDate is a free data retrieval call binding the contract method 0x79c3199d.
//
// Solidity: function preIcoStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PreIcoStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoStartDate(&_Crowdsale.CallOpts)
}

// PreIcoWeiRaised is a free data retrieval call binding the contract method 0xbf775660.
//
// Solidity: function preIcoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PreIcoWeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "preIcoWeiRaised")
	return *ret0, err
}

// PreIcoWeiRaised is a free data retrieval call binding the contract method 0xbf775660.
//
// Solidity: function preIcoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PreIcoWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoWeiRaised(&_Crowdsale.CallOpts)
}

// PreIcoWeiRaised is a free data retrieval call binding the contract method 0xbf775660.
//
// Solidity: function preIcoWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PreIcoWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.PreIcoWeiRaised(&_Crowdsale.CallOpts)
}

// PrivatePriceInFiatFracture is a free data retrieval call binding the contract method 0xa96d3141.
//
// Solidity: function privatePriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PrivatePriceInFiatFracture(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "privatePriceInFiatFracture")
	return *ret0, err
}

// PrivatePriceInFiatFracture is a free data retrieval call binding the contract method 0xa96d3141.
//
// Solidity: function privatePriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PrivatePriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.PrivatePriceInFiatFracture(&_Crowdsale.CallOpts)
}

// PrivatePriceInFiatFracture is a free data retrieval call binding the contract method 0xa96d3141.
//
// Solidity: function privatePriceInFiatFracture() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PrivatePriceInFiatFracture() (*big.Int, error) {
	return _Crowdsale.Contract.PrivatePriceInFiatFracture(&_Crowdsale.CallOpts)
}

// PrivateSaleFundsWithdrawn is a free data retrieval call binding the contract method 0xf80b99ef.
//
// Solidity: function privateSaleFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) PrivateSaleFundsWithdrawn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "privateSaleFundsWithdrawn")
	return *ret0, err
}

// PrivateSaleFundsWithdrawn is a free data retrieval call binding the contract method 0xf80b99ef.
//
// Solidity: function privateSaleFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleSession) PrivateSaleFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.PrivateSaleFundsWithdrawn(&_Crowdsale.CallOpts)
}

// PrivateSaleFundsWithdrawn is a free data retrieval call binding the contract method 0xf80b99ef.
//
// Solidity: function privateSaleFundsWithdrawn() constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) PrivateSaleFundsWithdrawn() (bool, error) {
	return _Crowdsale.Contract.PrivateSaleFundsWithdrawn(&_Crowdsale.CallOpts)
}

// PrivateSaleHardCap is a free data retrieval call binding the contract method 0x0912b9b7.
//
// Solidity: function privateSaleHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PrivateSaleHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "privateSaleHardCap")
	return *ret0, err
}

// PrivateSaleHardCap is a free data retrieval call binding the contract method 0x0912b9b7.
//
// Solidity: function privateSaleHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PrivateSaleHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleHardCap(&_Crowdsale.CallOpts)
}

// PrivateSaleHardCap is a free data retrieval call binding the contract method 0x0912b9b7.
//
// Solidity: function privateSaleHardCap() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PrivateSaleHardCap() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleHardCap(&_Crowdsale.CallOpts)
}

// PrivateSaleStartDate is a free data retrieval call binding the contract method 0x86d2fe57.
//
// Solidity: function privateSaleStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PrivateSaleStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "privateSaleStartDate")
	return *ret0, err
}

// PrivateSaleStartDate is a free data retrieval call binding the contract method 0x86d2fe57.
//
// Solidity: function privateSaleStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PrivateSaleStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleStartDate(&_Crowdsale.CallOpts)
}

// PrivateSaleStartDate is a free data retrieval call binding the contract method 0x86d2fe57.
//
// Solidity: function privateSaleStartDate() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PrivateSaleStartDate() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleStartDate(&_Crowdsale.CallOpts)
}

// PrivateSaleWeiRaised is a free data retrieval call binding the contract method 0x985585e1.
//
// Solidity: function privateSaleWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) PrivateSaleWeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "privateSaleWeiRaised")
	return *ret0, err
}

// PrivateSaleWeiRaised is a free data retrieval call binding the contract method 0x985585e1.
//
// Solidity: function privateSaleWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) PrivateSaleWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleWeiRaised(&_Crowdsale.CallOpts)
}

// PrivateSaleWeiRaised is a free data retrieval call binding the contract method 0x985585e1.
//
// Solidity: function privateSaleWeiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) PrivateSaleWeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.PrivateSaleWeiRaised(&_Crowdsale.CallOpts)
}

// ReserveFreezeTimestamp is a free data retrieval call binding the contract method 0x1abff816.
//
// Solidity: function reserveFreezeTimestamp() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) ReserveFreezeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "reserveFreezeTimestamp")
	return *ret0, err
}

// ReserveFreezeTimestamp is a free data retrieval call binding the contract method 0x1abff816.
//
// Solidity: function reserveFreezeTimestamp() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) ReserveFreezeTimestamp() (*big.Int, error) {
	return _Crowdsale.Contract.ReserveFreezeTimestamp(&_Crowdsale.CallOpts)
}

// ReserveFreezeTimestamp is a free data retrieval call binding the contract method 0x1abff816.
//
// Solidity: function reserveFreezeTimestamp() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) ReserveFreezeTimestamp() (*big.Int, error) {
	return _Crowdsale.Contract.ReserveFreezeTimestamp(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// TokensOrdered is a free data retrieval call binding the contract method 0x5cac2917.
//
// Solidity: function tokensOrdered( address) constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) TokensOrdered(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "tokensOrdered", arg0)
	return *ret0, err
}

// TokensOrdered is a free data retrieval call binding the contract method 0x5cac2917.
//
// Solidity: function tokensOrdered( address) constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) TokensOrdered(arg0 common.Address) (*big.Int, error) {
	return _Crowdsale.Contract.TokensOrdered(&_Crowdsale.CallOpts, arg0)
}

// TokensOrdered is a free data retrieval call binding the contract method 0x5cac2917.
//
// Solidity: function tokensOrdered( address) constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) TokensOrdered(arg0 common.Address) (*big.Int, error) {
	return _Crowdsale.Contract.TokensOrdered(&_Crowdsale.CallOpts, arg0)
}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) TokensSold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "tokensSold")
	return *ret0, err
}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) TokensSold() (*big.Int, error) {
	return _Crowdsale.Contract.TokensSold(&_Crowdsale.CallOpts)
}

// TokensSold is a free data retrieval call binding the contract method 0x518ab2a8.
//
// Solidity: function tokensSold() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) TokensSold() (*big.Int, error) {
	return _Crowdsale.Contract.TokensSold(&_Crowdsale.CallOpts)
}

// TotalWhiteListed is a free data retrieval call binding the contract method 0xa07b206f.
//
// Solidity: function totalWhiteListed() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) TotalWhiteListed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "totalWhiteListed")
	return *ret0, err
}

// TotalWhiteListed is a free data retrieval call binding the contract method 0xa07b206f.
//
// Solidity: function totalWhiteListed() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) TotalWhiteListed() (*big.Int, error) {
	return _Crowdsale.Contract.TotalWhiteListed(&_Crowdsale.CallOpts)
}

// TotalWhiteListed is a free data retrieval call binding the contract method 0xa07b206f.
//
// Solidity: function totalWhiteListed() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) TotalWhiteListed() (*big.Int, error) {
	return _Crowdsale.Contract.TotalWhiteListed(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) WeiInFiat(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "weiInFiat")
	return *ret0, err
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) WeiInFiat() (*big.Int, error) {
	return _Crowdsale.Contract.WeiInFiat(&_Crowdsale.CallOpts)
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) WeiInFiat() (*big.Int, error) {
	return _Crowdsale.Contract.WeiInFiat(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList( address) constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "whiteList", arg0)
	return *ret0, err
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList( address) constant returns(bool)
func (_Crowdsale *CrowdsaleSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Crowdsale.Contract.WhiteList(&_Crowdsale.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList( address) constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _Crowdsale.Contract.WhiteList(&_Crowdsale.CallOpts, arg0)
}

// WithdrawnState is a free data retrieval call binding the contract method 0xdb054c2f.
//
// Solidity: function withdrawnState( uint256) constant returns(bool)
func (_Crowdsale *CrowdsaleCaller) WithdrawnState(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crowdsale.contract.Call(opts, out, "withdrawnState", arg0)
	return *ret0, err
}

// WithdrawnState is a free data retrieval call binding the contract method 0xdb054c2f.
//
// Solidity: function withdrawnState( uint256) constant returns(bool)
func (_Crowdsale *CrowdsaleSession) WithdrawnState(arg0 *big.Int) (bool, error) {
	return _Crowdsale.Contract.WithdrawnState(&_Crowdsale.CallOpts, arg0)
}

// WithdrawnState is a free data retrieval call binding the contract method 0xdb054c2f.
//
// Solidity: function withdrawnState( uint256) constant returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) WithdrawnState(arg0 *big.Int) (bool, error) {
	return _Crowdsale.Contract.WithdrawnState(&_Crowdsale.CallOpts, arg0)
}

// AddPrivateParticipant is a paid mutator transaction binding the contract method 0x980a3691.
//
// Solidity: function addPrivateParticipant(_participant address) returns()
func (_Crowdsale *CrowdsaleTransactor) AddPrivateParticipant(opts *bind.TransactOpts, _participant common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "addPrivateParticipant", _participant)
}

// AddPrivateParticipant is a paid mutator transaction binding the contract method 0x980a3691.
//
// Solidity: function addPrivateParticipant(_participant address) returns()
func (_Crowdsale *CrowdsaleSession) AddPrivateParticipant(_participant common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddPrivateParticipant(&_Crowdsale.TransactOpts, _participant)
}

// AddPrivateParticipant is a paid mutator transaction binding the contract method 0x980a3691.
//
// Solidity: function addPrivateParticipant(_participant address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) AddPrivateParticipant(_participant common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddPrivateParticipant(&_Crowdsale.TransactOpts, _participant)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_user address) returns(bool)
func (_Crowdsale *CrowdsaleTransactor) AddToWhiteList(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "addToWhiteList", _user)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_user address) returns(bool)
func (_Crowdsale *CrowdsaleSession) AddToWhiteList(_user common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddToWhiteList(&_Crowdsale.TransactOpts, _user)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_user address) returns(bool)
func (_Crowdsale *CrowdsaleTransactorSession) AddToWhiteList(_user common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddToWhiteList(&_Crowdsale.TransactOpts, _user)
}

// AddToWhiteListMultiple is a paid mutator transaction binding the contract method 0x9fec8e3b.
//
// Solidity: function addToWhiteListMultiple(_users address[]) returns(bool)
func (_Crowdsale *CrowdsaleTransactor) AddToWhiteListMultiple(opts *bind.TransactOpts, _users []common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "addToWhiteListMultiple", _users)
}

// AddToWhiteListMultiple is a paid mutator transaction binding the contract method 0x9fec8e3b.
//
// Solidity: function addToWhiteListMultiple(_users address[]) returns(bool)
func (_Crowdsale *CrowdsaleSession) AddToWhiteListMultiple(_users []common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddToWhiteListMultiple(&_Crowdsale.TransactOpts, _users)
}

// AddToWhiteListMultiple is a paid mutator transaction binding the contract method 0x9fec8e3b.
//
// Solidity: function addToWhiteListMultiple(_users address[]) returns(bool)
func (_Crowdsale *CrowdsaleTransactorSession) AddToWhiteListMultiple(_users []common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.AddToWhiteListMultiple(&_Crowdsale.TransactOpts, _users)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_sender address) returns()
func (_Crowdsale *CrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "buyTokens", _sender)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_sender address) returns()
func (_Crowdsale *CrowdsaleSession) BuyTokens(_sender common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, _sender)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_sender address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) BuyTokens(_sender common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, _sender)
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_Crowdsale *CrowdsaleTransactor) CheckState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "checkState")
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_Crowdsale *CrowdsaleSession) CheckState() (*types.Transaction, error) {
	return _Crowdsale.Contract.CheckState(&_Crowdsale.TransactOpts)
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_Crowdsale *CrowdsaleTransactorSession) CheckState() (*types.Transaction, error) {
	return _Crowdsale.Contract.CheckState(&_Crowdsale.TransactOpts)
}

// ManualReserve is a paid mutator transaction binding the contract method 0x1b0dbdaf.
//
// Solidity: function manualReserve(_receiver address, _amount uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) ManualReserve(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "manualReserve", _receiver, _amount)
}

// ManualReserve is a paid mutator transaction binding the contract method 0x1b0dbdaf.
//
// Solidity: function manualReserve(_receiver address, _amount uint256) returns()
func (_Crowdsale *CrowdsaleSession) ManualReserve(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.ManualReserve(&_Crowdsale.TransactOpts, _receiver, _amount)
}

// ManualReserve is a paid mutator transaction binding the contract method 0x1b0dbdaf.
//
// Solidity: function manualReserve(_receiver address, _amount uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) ManualReserve(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.ManualReserve(&_Crowdsale.TransactOpts, _receiver, _amount)
}

// PrepareCrowdsale is a paid mutator transaction binding the contract method 0xe084a819.
//
// Solidity: function prepareCrowdsale() returns()
func (_Crowdsale *CrowdsaleTransactor) PrepareCrowdsale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "prepareCrowdsale")
}

// PrepareCrowdsale is a paid mutator transaction binding the contract method 0xe084a819.
//
// Solidity: function prepareCrowdsale() returns()
func (_Crowdsale *CrowdsaleSession) PrepareCrowdsale() (*types.Transaction, error) {
	return _Crowdsale.Contract.PrepareCrowdsale(&_Crowdsale.TransactOpts)
}

// PrepareCrowdsale is a paid mutator transaction binding the contract method 0xe084a819.
//
// Solidity: function prepareCrowdsale() returns()
func (_Crowdsale *CrowdsaleTransactorSession) PrepareCrowdsale() (*types.Transaction, error) {
	return _Crowdsale.Contract.PrepareCrowdsale(&_Crowdsale.TransactOpts)
}

// ReceiveFrozenReserve is a paid mutator transaction binding the contract method 0x96f85f16.
//
// Solidity: function receiveFrozenReserve(_receiver address) returns()
func (_Crowdsale *CrowdsaleTransactor) ReceiveFrozenReserve(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "receiveFrozenReserve", _receiver)
}

// ReceiveFrozenReserve is a paid mutator transaction binding the contract method 0x96f85f16.
//
// Solidity: function receiveFrozenReserve(_receiver address) returns()
func (_Crowdsale *CrowdsaleSession) ReceiveFrozenReserve(_receiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.ReceiveFrozenReserve(&_Crowdsale.TransactOpts, _receiver)
}

// ReceiveFrozenReserve is a paid mutator transaction binding the contract method 0x96f85f16.
//
// Solidity: function receiveFrozenReserve(_receiver address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) ReceiveFrozenReserve(_receiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.ReceiveFrozenReserve(&_Crowdsale.TransactOpts, _receiver)
}

// ReceiveOrderedTokens is a paid mutator transaction binding the contract method 0xd3a106e5.
//
// Solidity: function receiveOrderedTokens() returns()
func (_Crowdsale *CrowdsaleTransactor) ReceiveOrderedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "receiveOrderedTokens")
}

// ReceiveOrderedTokens is a paid mutator transaction binding the contract method 0xd3a106e5.
//
// Solidity: function receiveOrderedTokens() returns()
func (_Crowdsale *CrowdsaleSession) ReceiveOrderedTokens() (*types.Transaction, error) {
	return _Crowdsale.Contract.ReceiveOrderedTokens(&_Crowdsale.TransactOpts)
}

// ReceiveOrderedTokens is a paid mutator transaction binding the contract method 0xd3a106e5.
//
// Solidity: function receiveOrderedTokens() returns()
func (_Crowdsale *CrowdsaleTransactorSession) ReceiveOrderedTokens() (*types.Transaction, error) {
	return _Crowdsale.Contract.ReceiveOrderedTokens(&_Crowdsale.TransactOpts)
}

// SetIcoFinishDate is a paid mutator transaction binding the contract method 0xae32ac45.
//
// Solidity: function setIcoFinishDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) SetIcoFinishDate(opts *bind.TransactOpts, _newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "setIcoFinishDate", _newDate)
}

// SetIcoFinishDate is a paid mutator transaction binding the contract method 0xae32ac45.
//
// Solidity: function setIcoFinishDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleSession) SetIcoFinishDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetIcoFinishDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetIcoFinishDate is a paid mutator transaction binding the contract method 0xae32ac45.
//
// Solidity: function setIcoFinishDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) SetIcoFinishDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetIcoFinishDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetIcoStartDate is a paid mutator transaction binding the contract method 0x4a1c13cd.
//
// Solidity: function setIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) SetIcoStartDate(opts *bind.TransactOpts, _newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "setIcoStartDate", _newDate)
}

// SetIcoStartDate is a paid mutator transaction binding the contract method 0x4a1c13cd.
//
// Solidity: function setIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleSession) SetIcoStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetIcoStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetIcoStartDate is a paid mutator transaction binding the contract method 0x4a1c13cd.
//
// Solidity: function setIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) SetIcoStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetIcoStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetPreIcoStartDate is a paid mutator transaction binding the contract method 0x39c1904d.
//
// Solidity: function setPreIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) SetPreIcoStartDate(opts *bind.TransactOpts, _newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "setPreIcoStartDate", _newDate)
}

// SetPreIcoStartDate is a paid mutator transaction binding the contract method 0x39c1904d.
//
// Solidity: function setPreIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleSession) SetPreIcoStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetPreIcoStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetPreIcoStartDate is a paid mutator transaction binding the contract method 0x39c1904d.
//
// Solidity: function setPreIcoStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) SetPreIcoStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetPreIcoStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetPrivateSaleStartDate is a paid mutator transaction binding the contract method 0x53186c8c.
//
// Solidity: function setPrivateSaleStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) SetPrivateSaleStartDate(opts *bind.TransactOpts, _newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "setPrivateSaleStartDate", _newDate)
}

// SetPrivateSaleStartDate is a paid mutator transaction binding the contract method 0x53186c8c.
//
// Solidity: function setPrivateSaleStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleSession) SetPrivateSaleStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetPrivateSaleStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetPrivateSaleStartDate is a paid mutator transaction binding the contract method 0x53186c8c.
//
// Solidity: function setPrivateSaleStartDate(_newDate uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) SetPrivateSaleStartDate(_newDate *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetPrivateSaleStartDate(&_Crowdsale.TransactOpts, _newDate)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_Crowdsale *CrowdsaleTransactor) SetWallet(opts *bind.TransactOpts, _newWallet common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "setWallet", _newWallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_Crowdsale *CrowdsaleSession) SetWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetWallet(&_Crowdsale.TransactOpts, _newWallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_newWallet address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) SetWallet(_newWallet common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.SetWallet(&_Crowdsale.TransactOpts, _newWallet)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Crowdsale *CrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Crowdsale *CrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.TransferOwnership(&_Crowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Crowdsale *CrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.TransferOwnership(&_Crowdsale.TransactOpts, newOwner)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_Crowdsale *CrowdsaleTransactor) UpdateWeiInFiat(opts *bind.TransactOpts, _weiInFiat *big.Int) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "updateWeiInFiat", _weiInFiat)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_Crowdsale *CrowdsaleSession) UpdateWeiInFiat(_weiInFiat *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.UpdateWeiInFiat(&_Crowdsale.TransactOpts, _weiInFiat)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_Crowdsale *CrowdsaleTransactorSession) UpdateWeiInFiat(_weiInFiat *big.Int) (*types.Transaction, error) {
	return _Crowdsale.Contract.UpdateWeiInFiat(&_Crowdsale.TransactOpts, _weiInFiat)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Crowdsale *CrowdsaleTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Crowdsale *CrowdsaleSession) Withdraw() (*types.Transaction, error) {
	return _Crowdsale.Contract.Withdraw(&_Crowdsale.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Crowdsale *CrowdsaleTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Crowdsale.Contract.Withdraw(&_Crowdsale.TransactOpts)
}

// CrowdsaleFrozenReserveRetrievedIterator is returned from FilterFrozenReserveRetrieved and is used to iterate over the raw logs and unpacked data for FrozenReserveRetrieved events raised by the Crowdsale contract.
type CrowdsaleFrozenReserveRetrievedIterator struct {
	Event *CrowdsaleFrozenReserveRetrieved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleFrozenReserveRetrievedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleFrozenReserveRetrieved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleFrozenReserveRetrieved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleFrozenReserveRetrievedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleFrozenReserveRetrievedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleFrozenReserveRetrieved represents a FrozenReserveRetrieved event raised by the Crowdsale contract.
type CrowdsaleFrozenReserveRetrieved struct {
	Receiver          common.Address
	UnfreezeTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFrozenReserveRetrieved is a free log retrieval operation binding the contract event 0x39c2ff8a62bebaa5cd4af157ba107e234024dabe84f3e6f2cc416f50b7f380ab.
//
// Solidity: event FrozenReserveRetrieved(_receiver address, _unfreezeTimestamp uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterFrozenReserveRetrieved(opts *bind.FilterOpts) (*CrowdsaleFrozenReserveRetrievedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "FrozenReserveRetrieved")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleFrozenReserveRetrievedIterator{contract: _Crowdsale.contract, event: "FrozenReserveRetrieved", logs: logs, sub: sub}, nil
}

// WatchFrozenReserveRetrieved is a free log subscription operation binding the contract event 0x39c2ff8a62bebaa5cd4af157ba107e234024dabe84f3e6f2cc416f50b7f380ab.
//
// Solidity: event FrozenReserveRetrieved(_receiver address, _unfreezeTimestamp uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchFrozenReserveRetrieved(opts *bind.WatchOpts, sink chan<- *CrowdsaleFrozenReserveRetrieved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "FrozenReserveRetrieved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleFrozenReserveRetrieved)
				if err := _Crowdsale.contract.UnpackLog(event, "FrozenReserveRetrieved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Crowdsale contract.
type CrowdsaleFundsWithdrawnIterator struct {
	Event *CrowdsaleFundsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleFundsWithdrawn represents a FundsWithdrawn event raised by the Crowdsale contract.
type CrowdsaleFundsWithdrawn struct {
	Wallet common.Address
	Amount *big.Int
	Stage  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(_wallet address, _amount uint256, _stage uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts) (*CrowdsaleFundsWithdrawnIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleFundsWithdrawnIterator{contract: _Crowdsale.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(_wallet address, _amount uint256, _stage uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *CrowdsaleFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleFundsWithdrawn)
				if err := _Crowdsale.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleIcoFinishDateMovedIterator is returned from FilterIcoFinishDateMoved and is used to iterate over the raw logs and unpacked data for IcoFinishDateMoved events raised by the Crowdsale contract.
type CrowdsaleIcoFinishDateMovedIterator struct {
	Event *CrowdsaleIcoFinishDateMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleIcoFinishDateMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleIcoFinishDateMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleIcoFinishDateMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleIcoFinishDateMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleIcoFinishDateMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleIcoFinishDateMoved represents a IcoFinishDateMoved event raised by the Crowdsale contract.
type CrowdsaleIcoFinishDateMoved struct {
	OldDate *big.Int
	NewDate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIcoFinishDateMoved is a free log retrieval operation binding the contract event 0x47c07aea40f2e4306adb071f865fcd76a81d353c213ea7eef819581989ff9f67.
//
// Solidity: event IcoFinishDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterIcoFinishDateMoved(opts *bind.FilterOpts) (*CrowdsaleIcoFinishDateMovedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "IcoFinishDateMoved")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleIcoFinishDateMovedIterator{contract: _Crowdsale.contract, event: "IcoFinishDateMoved", logs: logs, sub: sub}, nil
}

// WatchIcoFinishDateMoved is a free log subscription operation binding the contract event 0x47c07aea40f2e4306adb071f865fcd76a81d353c213ea7eef819581989ff9f67.
//
// Solidity: event IcoFinishDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchIcoFinishDateMoved(opts *bind.WatchOpts, sink chan<- *CrowdsaleIcoFinishDateMoved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "IcoFinishDateMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleIcoFinishDateMoved)
				if err := _Crowdsale.contract.UnpackLog(event, "IcoFinishDateMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleIcoStartDateMovedIterator is returned from FilterIcoStartDateMoved and is used to iterate over the raw logs and unpacked data for IcoStartDateMoved events raised by the Crowdsale contract.
type CrowdsaleIcoStartDateMovedIterator struct {
	Event *CrowdsaleIcoStartDateMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleIcoStartDateMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleIcoStartDateMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleIcoStartDateMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleIcoStartDateMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleIcoStartDateMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleIcoStartDateMoved represents a IcoStartDateMoved event raised by the Crowdsale contract.
type CrowdsaleIcoStartDateMoved struct {
	OldDate *big.Int
	NewDate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIcoStartDateMoved is a free log retrieval operation binding the contract event 0x8d20f8b7089681de055ef410b1e22f13b3a44b8cda9145df73902b0505115c1c.
//
// Solidity: event IcoStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterIcoStartDateMoved(opts *bind.FilterOpts) (*CrowdsaleIcoStartDateMovedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "IcoStartDateMoved")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleIcoStartDateMovedIterator{contract: _Crowdsale.contract, event: "IcoStartDateMoved", logs: logs, sub: sub}, nil
}

// WatchIcoStartDateMoved is a free log subscription operation binding the contract event 0x8d20f8b7089681de055ef410b1e22f13b3a44b8cda9145df73902b0505115c1c.
//
// Solidity: event IcoStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchIcoStartDateMoved(opts *bind.WatchOpts, sink chan<- *CrowdsaleIcoStartDateMoved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "IcoStartDateMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleIcoStartDateMoved)
				if err := _Crowdsale.contract.UnpackLog(event, "IcoStartDateMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleLogWhiteListedIterator is returned from FilterLogWhiteListed and is used to iterate over the raw logs and unpacked data for LogWhiteListed events raised by the Crowdsale contract.
type CrowdsaleLogWhiteListedIterator struct {
	Event *CrowdsaleLogWhiteListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleLogWhiteListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleLogWhiteListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleLogWhiteListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleLogWhiteListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleLogWhiteListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleLogWhiteListed represents a LogWhiteListed event raised by the Crowdsale contract.
type CrowdsaleLogWhiteListed struct {
	User           common.Address
	WhiteListedNum *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogWhiteListed is a free log retrieval operation binding the contract event 0x88eb615cbff4540422d181389333a5c75e3d5eb98dd55fb176e3a615ddfd0f1f.
//
// Solidity: event LogWhiteListed(user address, whiteListedNum uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterLogWhiteListed(opts *bind.FilterOpts) (*CrowdsaleLogWhiteListedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "LogWhiteListed")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleLogWhiteListedIterator{contract: _Crowdsale.contract, event: "LogWhiteListed", logs: logs, sub: sub}, nil
}

// WatchLogWhiteListed is a free log subscription operation binding the contract event 0x88eb615cbff4540422d181389333a5c75e3d5eb98dd55fb176e3a615ddfd0f1f.
//
// Solidity: event LogWhiteListed(user address, whiteListedNum uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchLogWhiteListed(opts *bind.WatchOpts, sink chan<- *CrowdsaleLogWhiteListed) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "LogWhiteListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleLogWhiteListed)
				if err := _Crowdsale.contract.UnpackLog(event, "LogWhiteListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleLogWhiteListedMultipleIterator is returned from FilterLogWhiteListedMultiple and is used to iterate over the raw logs and unpacked data for LogWhiteListedMultiple events raised by the Crowdsale contract.
type CrowdsaleLogWhiteListedMultipleIterator struct {
	Event *CrowdsaleLogWhiteListedMultiple // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleLogWhiteListedMultipleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleLogWhiteListedMultiple)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleLogWhiteListedMultiple)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleLogWhiteListedMultipleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleLogWhiteListedMultipleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleLogWhiteListedMultiple represents a LogWhiteListedMultiple event raised by the Crowdsale contract.
type CrowdsaleLogWhiteListedMultiple struct {
	WhiteListedNum *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogWhiteListedMultiple is a free log retrieval operation binding the contract event 0x9075b10cfdcbb4feaac9b18054833db2bdacc10263b692174de5b561d34e2f4f.
//
// Solidity: event LogWhiteListedMultiple(whiteListedNum uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterLogWhiteListedMultiple(opts *bind.FilterOpts) (*CrowdsaleLogWhiteListedMultipleIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "LogWhiteListedMultiple")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleLogWhiteListedMultipleIterator{contract: _Crowdsale.contract, event: "LogWhiteListedMultiple", logs: logs, sub: sub}, nil
}

// WatchLogWhiteListedMultiple is a free log subscription operation binding the contract event 0x9075b10cfdcbb4feaac9b18054833db2bdacc10263b692174de5b561d34e2f4f.
//
// Solidity: event LogWhiteListedMultiple(whiteListedNum uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchLogWhiteListedMultiple(opts *bind.WatchOpts, sink chan<- *CrowdsaleLogWhiteListedMultiple) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "LogWhiteListedMultiple")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleLogWhiteListedMultiple)
				if err := _Crowdsale.contract.UnpackLog(event, "LogWhiteListedMultiple", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleManualTokenSendIterator is returned from FilterManualTokenSend and is used to iterate over the raw logs and unpacked data for ManualTokenSend events raised by the Crowdsale contract.
type CrowdsaleManualTokenSendIterator struct {
	Event *CrowdsaleManualTokenSend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleManualTokenSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleManualTokenSend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleManualTokenSend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleManualTokenSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleManualTokenSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleManualTokenSend represents a ManualTokenSend event raised by the Crowdsale contract.
type CrowdsaleManualTokenSend struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterManualTokenSend is a free log retrieval operation binding the contract event 0xc1c7fcfd21f7aae09db30246e22bda2d3a2baf012df58fa845af5e1e4cc70fc6.
//
// Solidity: event ManualTokenSend(_receiver address, _amount uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterManualTokenSend(opts *bind.FilterOpts) (*CrowdsaleManualTokenSendIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "ManualTokenSend")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleManualTokenSendIterator{contract: _Crowdsale.contract, event: "ManualTokenSend", logs: logs, sub: sub}, nil
}

// WatchManualTokenSend is a free log subscription operation binding the contract event 0xc1c7fcfd21f7aae09db30246e22bda2d3a2baf012df58fa845af5e1e4cc70fc6.
//
// Solidity: event ManualTokenSend(_receiver address, _amount uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchManualTokenSend(opts *bind.WatchOpts, sink chan<- *CrowdsaleManualTokenSend) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "ManualTokenSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleManualTokenSend)
				if err := _Crowdsale.contract.UnpackLog(event, "ManualTokenSend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crowdsale contract.
type CrowdsaleOwnershipTransferredIterator struct {
	Event *CrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the Crowdsale contract.
type CrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Crowdsale *CrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleOwnershipTransferredIterator{contract: _Crowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Crowdsale *CrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleOwnershipTransferred)
				if err := _Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsalePreIcoStartDateMovedIterator is returned from FilterPreIcoStartDateMoved and is used to iterate over the raw logs and unpacked data for PreIcoStartDateMoved events raised by the Crowdsale contract.
type CrowdsalePreIcoStartDateMovedIterator struct {
	Event *CrowdsalePreIcoStartDateMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsalePreIcoStartDateMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsalePreIcoStartDateMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsalePreIcoStartDateMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsalePreIcoStartDateMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsalePreIcoStartDateMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsalePreIcoStartDateMoved represents a PreIcoStartDateMoved event raised by the Crowdsale contract.
type CrowdsalePreIcoStartDateMoved struct {
	OldDate *big.Int
	NewDate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPreIcoStartDateMoved is a free log retrieval operation binding the contract event 0x9294d04a272cdf981934ca76de43865446ea70e949478d270dad1e6d90b5960e.
//
// Solidity: event PreIcoStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterPreIcoStartDateMoved(opts *bind.FilterOpts) (*CrowdsalePreIcoStartDateMovedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "PreIcoStartDateMoved")
	if err != nil {
		return nil, err
	}
	return &CrowdsalePreIcoStartDateMovedIterator{contract: _Crowdsale.contract, event: "PreIcoStartDateMoved", logs: logs, sub: sub}, nil
}

// WatchPreIcoStartDateMoved is a free log subscription operation binding the contract event 0x9294d04a272cdf981934ca76de43865446ea70e949478d270dad1e6d90b5960e.
//
// Solidity: event PreIcoStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchPreIcoStartDateMoved(opts *bind.WatchOpts, sink chan<- *CrowdsalePreIcoStartDateMoved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "PreIcoStartDateMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsalePreIcoStartDateMoved)
				if err := _Crowdsale.contract.UnpackLog(event, "PreIcoStartDateMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsalePrivateSaleStartDateMovedIterator is returned from FilterPrivateSaleStartDateMoved and is used to iterate over the raw logs and unpacked data for PrivateSaleStartDateMoved events raised by the Crowdsale contract.
type CrowdsalePrivateSaleStartDateMovedIterator struct {
	Event *CrowdsalePrivateSaleStartDateMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsalePrivateSaleStartDateMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsalePrivateSaleStartDateMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsalePrivateSaleStartDateMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsalePrivateSaleStartDateMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsalePrivateSaleStartDateMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsalePrivateSaleStartDateMoved represents a PrivateSaleStartDateMoved event raised by the Crowdsale contract.
type CrowdsalePrivateSaleStartDateMoved struct {
	OldDate *big.Int
	NewDate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrivateSaleStartDateMoved is a free log retrieval operation binding the contract event 0xefe29dae758f051a9c039b9f34ee371f674e06edf12f196cfeda8a42131fdfe5.
//
// Solidity: event PrivateSaleStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterPrivateSaleStartDateMoved(opts *bind.FilterOpts) (*CrowdsalePrivateSaleStartDateMovedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "PrivateSaleStartDateMoved")
	if err != nil {
		return nil, err
	}
	return &CrowdsalePrivateSaleStartDateMovedIterator{contract: _Crowdsale.contract, event: "PrivateSaleStartDateMoved", logs: logs, sub: sub}, nil
}

// WatchPrivateSaleStartDateMoved is a free log subscription operation binding the contract event 0xefe29dae758f051a9c039b9f34ee371f674e06edf12f196cfeda8a42131fdfe5.
//
// Solidity: event PrivateSaleStartDateMoved(_oldDate uint256, _newDate uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchPrivateSaleStartDateMoved(opts *bind.WatchOpts, sink chan<- *CrowdsalePrivateSaleStartDateMoved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "PrivateSaleStartDateMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsalePrivateSaleStartDateMoved)
				if err := _Crowdsale.contract.UnpackLog(event, "PrivateSaleStartDateMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleRateUpdatedIterator is returned from FilterRateUpdated and is used to iterate over the raw logs and unpacked data for RateUpdated events raised by the Crowdsale contract.
type CrowdsaleRateUpdatedIterator struct {
	Event *CrowdsaleRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleRateUpdated represents a RateUpdated event raised by the Crowdsale contract.
type CrowdsaleRateUpdated struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateUpdated is a free log retrieval operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(_oldRate uint256, _newRate uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterRateUpdated(opts *bind.FilterOpts) (*CrowdsaleRateUpdatedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleRateUpdatedIterator{contract: _Crowdsale.contract, event: "RateUpdated", logs: logs, sub: sub}, nil
}

// WatchRateUpdated is a free log subscription operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(_oldRate uint256, _newRate uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchRateUpdated(opts *bind.WatchOpts, sink chan<- *CrowdsaleRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleRateUpdated)
				if err := _Crowdsale.contract.UnpackLog(event, "RateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleStateHasChangedIterator is returned from FilterStateHasChanged and is used to iterate over the raw logs and unpacked data for StateHasChanged events raised by the Crowdsale contract.
type CrowdsaleStateHasChangedIterator struct {
	Event *CrowdsaleStateHasChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleStateHasChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleStateHasChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleStateHasChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleStateHasChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleStateHasChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleStateHasChanged represents a StateHasChanged event raised by the Crowdsale contract.
type CrowdsaleStateHasChanged struct {
	OldState uint8
	NewState uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStateHasChanged is a free log retrieval operation binding the contract event 0x40bc5f093899d560cbd784d85b697b577e63a2991330daf682bd4d6ade755b03.
//
// Solidity: event StateHasChanged(_oldState uint8, _newState uint8)
func (_Crowdsale *CrowdsaleFilterer) FilterStateHasChanged(opts *bind.FilterOpts) (*CrowdsaleStateHasChangedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "StateHasChanged")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleStateHasChangedIterator{contract: _Crowdsale.contract, event: "StateHasChanged", logs: logs, sub: sub}, nil
}

// WatchStateHasChanged is a free log subscription operation binding the contract event 0x40bc5f093899d560cbd784d85b697b577e63a2991330daf682bd4d6ade755b03.
//
// Solidity: event StateHasChanged(_oldState uint8, _newState uint8)
func (_Crowdsale *CrowdsaleFilterer) WatchStateHasChanged(opts *bind.WatchOpts, sink chan<- *CrowdsaleStateHasChanged) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "StateHasChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleStateHasChanged)
				if err := _Crowdsale.contract.UnpackLog(event, "StateHasChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleTokensAreOrderedIterator is returned from FilterTokensAreOrdered and is used to iterate over the raw logs and unpacked data for TokensAreOrdered events raised by the Crowdsale contract.
type CrowdsaleTokensAreOrderedIterator struct {
	Event *CrowdsaleTokensAreOrdered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleTokensAreOrderedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleTokensAreOrdered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleTokensAreOrdered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleTokensAreOrderedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleTokensAreOrderedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleTokensAreOrdered represents a TokensAreOrdered event raised by the Crowdsale contract.
type CrowdsaleTokensAreOrdered struct {
	Orderer          common.Address
	Amount           *big.Int
	CurrentWeiInFiat *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokensAreOrdered is a free log retrieval operation binding the contract event 0xb794cb9cc471caae752d175fddbc2f267fe684f1a107b2f19f44a95e602648b6.
//
// Solidity: event TokensAreOrdered(_orderer address, _amount uint256, _currentWeiInFiat uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterTokensAreOrdered(opts *bind.FilterOpts) (*CrowdsaleTokensAreOrderedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "TokensAreOrdered")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTokensAreOrderedIterator{contract: _Crowdsale.contract, event: "TokensAreOrdered", logs: logs, sub: sub}, nil
}

// WatchTokensAreOrdered is a free log subscription operation binding the contract event 0xb794cb9cc471caae752d175fddbc2f267fe684f1a107b2f19f44a95e602648b6.
//
// Solidity: event TokensAreOrdered(_orderer address, _amount uint256, _currentWeiInFiat uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchTokensAreOrdered(opts *bind.WatchOpts, sink chan<- *CrowdsaleTokensAreOrdered) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "TokensAreOrdered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleTokensAreOrdered)
				if err := _Crowdsale.contract.UnpackLog(event, "TokensAreOrdered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleTokensAreRetrievedIterator is returned from FilterTokensAreRetrieved and is used to iterate over the raw logs and unpacked data for TokensAreRetrieved events raised by the Crowdsale contract.
type CrowdsaleTokensAreRetrievedIterator struct {
	Event *CrowdsaleTokensAreRetrieved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleTokensAreRetrievedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleTokensAreRetrieved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleTokensAreRetrieved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleTokensAreRetrievedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleTokensAreRetrievedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleTokensAreRetrieved represents a TokensAreRetrieved event raised by the Crowdsale contract.
type CrowdsaleTokensAreRetrieved struct {
	Retriever common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensAreRetrieved is a free log retrieval operation binding the contract event 0xf33726a859f09868cbfde46ebe9962f051984065979b92e9a7533b202a3c490e.
//
// Solidity: event TokensAreRetrieved(_retriever address, _amount uint256)
func (_Crowdsale *CrowdsaleFilterer) FilterTokensAreRetrieved(opts *bind.FilterOpts) (*CrowdsaleTokensAreRetrievedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "TokensAreRetrieved")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTokensAreRetrievedIterator{contract: _Crowdsale.contract, event: "TokensAreRetrieved", logs: logs, sub: sub}, nil
}

// WatchTokensAreRetrieved is a free log subscription operation binding the contract event 0xf33726a859f09868cbfde46ebe9962f051984065979b92e9a7533b202a3c490e.
//
// Solidity: event TokensAreRetrieved(_retriever address, _amount uint256)
func (_Crowdsale *CrowdsaleFilterer) WatchTokensAreRetrieved(opts *bind.WatchOpts, sink chan<- *CrowdsaleTokensAreRetrieved) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "TokensAreRetrieved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleTokensAreRetrieved)
				if err := _Crowdsale.contract.UnpackLog(event, "TokensAreRetrieved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CrowdsaleWalletHasChangedIterator is returned from FilterWalletHasChanged and is used to iterate over the raw logs and unpacked data for WalletHasChanged events raised by the Crowdsale contract.
type CrowdsaleWalletHasChangedIterator struct {
	Event *CrowdsaleWalletHasChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrowdsaleWalletHasChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleWalletHasChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrowdsaleWalletHasChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrowdsaleWalletHasChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleWalletHasChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleWalletHasChanged represents a WalletHasChanged event raised by the Crowdsale contract.
type CrowdsaleWalletHasChanged struct {
	OldWallet common.Address
	NewWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletHasChanged is a free log retrieval operation binding the contract event 0x86b557e590894ebca3dce4dfc61800ecb963722c9cce83d7d0d1dd32e322420a.
//
// Solidity: event WalletHasChanged(_oldWallet address, _newWallet address)
func (_Crowdsale *CrowdsaleFilterer) FilterWalletHasChanged(opts *bind.FilterOpts) (*CrowdsaleWalletHasChangedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "WalletHasChanged")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleWalletHasChangedIterator{contract: _Crowdsale.contract, event: "WalletHasChanged", logs: logs, sub: sub}, nil
}

// WatchWalletHasChanged is a free log subscription operation binding the contract event 0x86b557e590894ebca3dce4dfc61800ecb963722c9cce83d7d0d1dd32e322420a.
//
// Solidity: event WalletHasChanged(_oldWallet address, _newWallet address)
func (_Crowdsale *CrowdsaleFilterer) WatchWalletHasChanged(opts *bind.WatchOpts, sink chan<- *CrowdsaleWalletHasChanged) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "WalletHasChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleWalletHasChanged)
				if err := _Crowdsale.contract.UnpackLog(event, "WalletHasChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BasicTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101768061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461007f575b600080fd5b341561005b57600080fd5b6100636100a0565b604051600160a060020a03909116815260200160405180910390f35b341561008a57600080fd5b61009e600160a060020a03600435166100af565b005b600054600160a060020a031681565b60005433600160a060020a039081169116146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058203d4af7836893dbb7923591c9734082ea6028a2d20d98991dfff64032a9c579ca0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a723058201fc74fca81597a65ade5a1633aea369a71b081c55395594ac6006de0231c29760029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b6106fb8061001e6000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100c857806323b872dd146100ed578063661884631461011557806370a0823114610137578063a9059cbb14610156578063d73dd62314610178578063dd62ed3e1461019a575b600080fd5b341561009d57600080fd5b6100b4600160a060020a03600435166024356101bf565b604051901515815260200160405180910390f35b34156100d357600080fd5b6100db61022b565b60405190815260200160405180910390f35b34156100f857600080fd5b6100b4600160a060020a0360043581169060243516604435610231565b341561012057600080fd5b6100b4600160a060020a03600435166024356103b1565b341561014257600080fd5b6100db600160a060020a03600435166104ab565b341561016157600080fd5b6100b4600160a060020a03600435166024356104c6565b341561018357600080fd5b6100b4600160a060020a03600435166024356105d8565b34156101a557600080fd5b6100db600160a060020a036004358116906024351661067c565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b6000600160a060020a038316151561024857600080fd5b600160a060020a03841660009081526020819052604090205482111561026d57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156102a057600080fd5b600160a060020a0384166000908152602081905260409020546102c9908363ffffffff6106a716565b600160a060020a0380861660009081526020819052604080822093909355908516815220546102fe908363ffffffff6106b916565b600160a060020a0380851660009081526020818152604080832094909455878316825260028152838220339093168252919091522054610344908363ffffffff6106a716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561040e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610445565b61041e818463ffffffff6106a716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a03831615156104dd57600080fd5b600160a060020a03331660009081526020819052604090205482111561050257600080fd5b600160a060020a03331660009081526020819052604090205461052b908363ffffffff6106a716565b600160a060020a033381166000908152602081905260408082209390935590851681522054610560908363ffffffff6106b916565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610610908363ffffffff6106b916565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b6000828211156106b357fe5b50900390565b6000828201838110156106c857fe5b93925050505600a165627a7a72305820a5d8e61592dda1c8579ade30cc76b5a830fb5fd834e8d7cab2849945f0e73a750029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsingFiatPriceABI is the input ABI used to generate the binding from.
const UsingFiatPriceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_weiInFiat\",\"type\":\"uint256\"}],\"name\":\"updateWeiInFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiInFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_fiatSymbol\",\"type\":\"string\"},{\"name\":\"_fiatDecimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"RateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// UsingFiatPriceBin is the compiled bytecode used for deploying new contracts.
const UsingFiatPriceBin = `0x6060604052341561000f57600080fd5b6040516104e53803806104e58339810160405280805182019190602001805160008054600160a060020a03191633600160a060020a03161790559150506002811080159061005e575060128111155b151561006957600080fd5b600182805161007c929160200190610086565b5060025550610121565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c757805160ff19168380011785556100f4565b828001600101855582156100f4579182015b828111156100f45782518255916020019190600101906100d9565b50610100929150610104565b5090565b61011e91905b80821115610100576000815560010161010a565b90565b6103b5806101306000396000f3006060604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631755bf3d811461008757806327631d021461009f5780638da5cb5b14610129578063b24a129014610158578063d0b06f5d1461017d578063f2fde38b14610190578063f5890e46146101af575b600080fd5b341561009257600080fd5b61009d6004356101c2565b005b34156100aa57600080fd5b6100b261022f565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100ee5780820151838201526020016100d6565b50505050905090810190601f16801561011b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013457600080fd5b61013c6102cd565b604051600160a060020a03909116815260200160405180910390f35b341561016357600080fd5b61016b6102dc565b60405190815260200160405180910390f35b341561018857600080fd5b61016b6102e2565b341561019b57600080fd5b61009d600160a060020a03600435166102e8565b34156101ba57600080fd5b61016b610383565b60005433600160a060020a039081169116146101dd57600080fd5b600081116101ea57600080fd5b7fb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d19166003548260405191825260208201526040908101905180910390a160035542600455565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c55780601f1061029a576101008083540402835291602001916102c5565b820191906000526020600020905b8154815290600101906020018083116102a857829003601f168201915b505050505081565b600054600160a060020a031681565b60035481565b60045481565b60005433600160a060020a0390811691161461030357600080fd5b600160a060020a038116151561031857600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254815600a165627a7a723058202892c40f15672b6d7ef4cfc0fe03eb54b43880d9a6eb4ca564f3fdbee11a8d6a0029`

// DeployUsingFiatPrice deploys a new Ethereum contract, binding an instance of UsingFiatPrice to it.
func DeployUsingFiatPrice(auth *bind.TransactOpts, backend bind.ContractBackend, _fiatSymbol string, _fiatDecimals *big.Int) (common.Address, *types.Transaction, *UsingFiatPrice, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingFiatPriceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingFiatPriceBin), backend, _fiatSymbol, _fiatDecimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingFiatPrice{UsingFiatPriceCaller: UsingFiatPriceCaller{contract: contract}, UsingFiatPriceTransactor: UsingFiatPriceTransactor{contract: contract}, UsingFiatPriceFilterer: UsingFiatPriceFilterer{contract: contract}}, nil
}

// UsingFiatPrice is an auto generated Go binding around an Ethereum contract.
type UsingFiatPrice struct {
	UsingFiatPriceCaller     // Read-only binding to the contract
	UsingFiatPriceTransactor // Write-only binding to the contract
	UsingFiatPriceFilterer   // Log filterer for contract events
}

// UsingFiatPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingFiatPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingFiatPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingFiatPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingFiatPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsingFiatPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingFiatPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingFiatPriceSession struct {
	Contract     *UsingFiatPrice   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingFiatPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingFiatPriceCallerSession struct {
	Contract *UsingFiatPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UsingFiatPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingFiatPriceTransactorSession struct {
	Contract     *UsingFiatPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UsingFiatPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingFiatPriceRaw struct {
	Contract *UsingFiatPrice // Generic contract binding to access the raw methods on
}

// UsingFiatPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingFiatPriceCallerRaw struct {
	Contract *UsingFiatPriceCaller // Generic read-only contract binding to access the raw methods on
}

// UsingFiatPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingFiatPriceTransactorRaw struct {
	Contract *UsingFiatPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingFiatPrice creates a new instance of UsingFiatPrice, bound to a specific deployed contract.
func NewUsingFiatPrice(address common.Address, backend bind.ContractBackend) (*UsingFiatPrice, error) {
	contract, err := bindUsingFiatPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingFiatPrice{UsingFiatPriceCaller: UsingFiatPriceCaller{contract: contract}, UsingFiatPriceTransactor: UsingFiatPriceTransactor{contract: contract}, UsingFiatPriceFilterer: UsingFiatPriceFilterer{contract: contract}}, nil
}

// NewUsingFiatPriceCaller creates a new read-only instance of UsingFiatPrice, bound to a specific deployed contract.
func NewUsingFiatPriceCaller(address common.Address, caller bind.ContractCaller) (*UsingFiatPriceCaller, error) {
	contract, err := bindUsingFiatPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsingFiatPriceCaller{contract: contract}, nil
}

// NewUsingFiatPriceTransactor creates a new write-only instance of UsingFiatPrice, bound to a specific deployed contract.
func NewUsingFiatPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingFiatPriceTransactor, error) {
	contract, err := bindUsingFiatPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsingFiatPriceTransactor{contract: contract}, nil
}

// NewUsingFiatPriceFilterer creates a new log filterer instance of UsingFiatPrice, bound to a specific deployed contract.
func NewUsingFiatPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*UsingFiatPriceFilterer, error) {
	contract, err := bindUsingFiatPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsingFiatPriceFilterer{contract: contract}, nil
}

// bindUsingFiatPrice binds a generic wrapper to an already deployed contract.
func bindUsingFiatPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingFiatPriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingFiatPrice *UsingFiatPriceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingFiatPrice.Contract.UsingFiatPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingFiatPrice *UsingFiatPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.UsingFiatPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingFiatPrice *UsingFiatPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.UsingFiatPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingFiatPrice *UsingFiatPriceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingFiatPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingFiatPrice *UsingFiatPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingFiatPrice *UsingFiatPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.contract.Transact(opts, method, params...)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCaller) FiatDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsingFiatPrice.contract.Call(opts, out, "fiatDecimals")
	return *ret0, err
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceSession) FiatDecimals() (*big.Int, error) {
	return _UsingFiatPrice.Contract.FiatDecimals(&_UsingFiatPrice.CallOpts)
}

// FiatDecimals is a free data retrieval call binding the contract method 0xf5890e46.
//
// Solidity: function fiatDecimals() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCallerSession) FiatDecimals() (*big.Int, error) {
	return _UsingFiatPrice.Contract.FiatDecimals(&_UsingFiatPrice.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_UsingFiatPrice *UsingFiatPriceCaller) FiatSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UsingFiatPrice.contract.Call(opts, out, "fiatSymbol")
	return *ret0, err
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_UsingFiatPrice *UsingFiatPriceSession) FiatSymbol() (string, error) {
	return _UsingFiatPrice.Contract.FiatSymbol(&_UsingFiatPrice.CallOpts)
}

// FiatSymbol is a free data retrieval call binding the contract method 0x27631d02.
//
// Solidity: function fiatSymbol() constant returns(string)
func (_UsingFiatPrice *UsingFiatPriceCallerSession) FiatSymbol() (string, error) {
	return _UsingFiatPrice.Contract.FiatSymbol(&_UsingFiatPrice.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCaller) LastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsingFiatPrice.contract.Call(opts, out, "lastUpdated")
	return *ret0, err
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceSession) LastUpdated() (*big.Int, error) {
	return _UsingFiatPrice.Contract.LastUpdated(&_UsingFiatPrice.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCallerSession) LastUpdated() (*big.Int, error) {
	return _UsingFiatPrice.Contract.LastUpdated(&_UsingFiatPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UsingFiatPrice *UsingFiatPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UsingFiatPrice.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UsingFiatPrice *UsingFiatPriceSession) Owner() (common.Address, error) {
	return _UsingFiatPrice.Contract.Owner(&_UsingFiatPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UsingFiatPrice *UsingFiatPriceCallerSession) Owner() (common.Address, error) {
	return _UsingFiatPrice.Contract.Owner(&_UsingFiatPrice.CallOpts)
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCaller) WeiInFiat(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsingFiatPrice.contract.Call(opts, out, "weiInFiat")
	return *ret0, err
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceSession) WeiInFiat() (*big.Int, error) {
	return _UsingFiatPrice.Contract.WeiInFiat(&_UsingFiatPrice.CallOpts)
}

// WeiInFiat is a free data retrieval call binding the contract method 0xb24a1290.
//
// Solidity: function weiInFiat() constant returns(uint256)
func (_UsingFiatPrice *UsingFiatPriceCallerSession) WeiInFiat() (*big.Int, error) {
	return _UsingFiatPrice.Contract.WeiInFiat(&_UsingFiatPrice.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UsingFiatPrice *UsingFiatPriceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UsingFiatPrice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UsingFiatPrice *UsingFiatPriceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.TransferOwnership(&_UsingFiatPrice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UsingFiatPrice *UsingFiatPriceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.TransferOwnership(&_UsingFiatPrice.TransactOpts, newOwner)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_UsingFiatPrice *UsingFiatPriceTransactor) UpdateWeiInFiat(opts *bind.TransactOpts, _weiInFiat *big.Int) (*types.Transaction, error) {
	return _UsingFiatPrice.contract.Transact(opts, "updateWeiInFiat", _weiInFiat)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_UsingFiatPrice *UsingFiatPriceSession) UpdateWeiInFiat(_weiInFiat *big.Int) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.UpdateWeiInFiat(&_UsingFiatPrice.TransactOpts, _weiInFiat)
}

// UpdateWeiInFiat is a paid mutator transaction binding the contract method 0x1755bf3d.
//
// Solidity: function updateWeiInFiat(_weiInFiat uint256) returns()
func (_UsingFiatPrice *UsingFiatPriceTransactorSession) UpdateWeiInFiat(_weiInFiat *big.Int) (*types.Transaction, error) {
	return _UsingFiatPrice.Contract.UpdateWeiInFiat(&_UsingFiatPrice.TransactOpts, _weiInFiat)
}

// UsingFiatPriceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UsingFiatPrice contract.
type UsingFiatPriceOwnershipTransferredIterator struct {
	Event *UsingFiatPriceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsingFiatPriceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsingFiatPriceOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsingFiatPriceOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsingFiatPriceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsingFiatPriceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsingFiatPriceOwnershipTransferred represents a OwnershipTransferred event raised by the UsingFiatPrice contract.
type UsingFiatPriceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_UsingFiatPrice *UsingFiatPriceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsingFiatPriceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UsingFiatPrice.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsingFiatPriceOwnershipTransferredIterator{contract: _UsingFiatPrice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_UsingFiatPrice *UsingFiatPriceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsingFiatPriceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UsingFiatPrice.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsingFiatPriceOwnershipTransferred)
				if err := _UsingFiatPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsingFiatPriceRateUpdatedIterator is returned from FilterRateUpdated and is used to iterate over the raw logs and unpacked data for RateUpdated events raised by the UsingFiatPrice contract.
type UsingFiatPriceRateUpdatedIterator struct {
	Event *UsingFiatPriceRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsingFiatPriceRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsingFiatPriceRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsingFiatPriceRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsingFiatPriceRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsingFiatPriceRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsingFiatPriceRateUpdated represents a RateUpdated event raised by the UsingFiatPrice contract.
type UsingFiatPriceRateUpdated struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateUpdated is a free log retrieval operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(_oldRate uint256, _newRate uint256)
func (_UsingFiatPrice *UsingFiatPriceFilterer) FilterRateUpdated(opts *bind.FilterOpts) (*UsingFiatPriceRateUpdatedIterator, error) {

	logs, sub, err := _UsingFiatPrice.contract.FilterLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return &UsingFiatPriceRateUpdatedIterator{contract: _UsingFiatPrice.contract, event: "RateUpdated", logs: logs, sub: sub}, nil
}

// WatchRateUpdated is a free log subscription operation binding the contract event 0xb38780ddde1f073d91c150de2696f3f7085883648ba21cc5ef01029cb21d1916.
//
// Solidity: event RateUpdated(_oldRate uint256, _newRate uint256)
func (_UsingFiatPrice *UsingFiatPriceFilterer) WatchRateUpdated(opts *bind.WatchOpts, sink chan<- *UsingFiatPriceRateUpdated) (event.Subscription, error) {

	logs, sub, err := _UsingFiatPrice.contract.WatchLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsingFiatPriceRateUpdated)
				if err := _UsingFiatPrice.contract.UnpackLog(event, "RateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WindMineTokenABI is the input ABI used to generate the binding from.
const WindMineTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundersWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundersReserve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_foundersWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// WindMineTokenBin is the compiled bytecode used for deploying new contracts.
const WindMineTokenBin = `0x6060604052341561000f57600080fd5b604051602080610a988339810160405280805160038054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b6638d7ea4c680000600181905560048054600160a060020a031916600160a060020a03841617905561009b9066071afd498d00006401000000006100d281026109511704565b600160a060020a033381166000908152602081905260408082209390935560045490911681522066071afd498d00009055506100e4565b6000828211156100de57fe5b50900390565b6109a5806100f36000396000f3006060604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100df578063095ea7b31461016957806318160ddd1461019f5780631bfaf155146101c457806323b872dd146101f3578063313ce5671461021b578063661884631461022e57806370a08231146102505780637b86120a1461026f5780638da5cb5b1461028257806395d89b4114610295578063a9059cbb146102a8578063d73dd623146102ca578063dd62ed3e146102ec578063f2fde38b14610311575b600080fd5b34156100ea57600080fd5b6100f2610332565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012e578082015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017457600080fd5b61018b600160a060020a0360043516602435610369565b604051901515815260200160405180910390f35b34156101aa57600080fd5b6101b26103d5565b60405190815260200160405180910390f35b34156101cf57600080fd5b6101d76103db565b604051600160a060020a03909116815260200160405180910390f35b34156101fe57600080fd5b61018b600160a060020a03600435811690602435166044356103ea565b341561022657600080fd5b6101b261056a565b341561023957600080fd5b61018b600160a060020a036004351660243561056f565b341561025b57600080fd5b6101b2600160a060020a0360043516610669565b341561027a57600080fd5b6101b2610684565b341561028d57600080fd5b6101d761068f565b34156102a057600080fd5b6100f261069e565b34156102b357600080fd5b61018b600160a060020a03600435166024356106d5565b34156102d557600080fd5b61018b600160a060020a03600435166024356107e7565b34156102f757600080fd5b6101b2600160a060020a036004358116906024351661088b565b341561031c57600080fd5b610330600160a060020a03600435166108b6565b005b60408051908101604052600881527f57696e644d696e65000000000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60015490565b600454600160a060020a031681565b6000600160a060020a038316151561040157600080fd5b600160a060020a03841660009081526020819052604090205482111561042657600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561045957600080fd5b600160a060020a038416600090815260208190526040902054610482908363ffffffff61095116565b600160a060020a0380861660009081526020819052604080822093909355908516815220546104b7908363ffffffff61096316565b600160a060020a03808516600090815260208181526040808320949094558783168252600281528382203390931682529190915220546104fd908363ffffffff61095116565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600881565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054808311156105cc57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610603565b6105dc818463ffffffff61095116565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526020819052604090205490565b66071afd498d000081565b600354600160a060020a031681565b60408051908101604052600381527f574d440000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a03831615156106ec57600080fd5b600160a060020a03331660009081526020819052604090205482111561071157600080fd5b600160a060020a03331660009081526020819052604090205461073a908363ffffffff61095116565b600160a060020a03338116600090815260208190526040808220939093559085168152205461076f908363ffffffff61096316565b60008085600160a060020a0316600160a060020a031681526020019081526020016000208190555082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a03338116600090815260026020908152604080832093861683529290529081205461081f908363ffffffff61096316565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146108d157600080fd5b600160a060020a03811615156108e657600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561095d57fe5b50900390565b60008282018381101561097257fe5b93925050505600a165627a7a7230582000d0f9300042a3129708c6249509a1f62188bf3c6e604dfdda625c084f313a0d0029`

// DeployWindMineToken deploys a new Ethereum contract, binding an instance of WindMineToken to it.
func DeployWindMineToken(auth *bind.TransactOpts, backend bind.ContractBackend, _foundersWallet common.Address) (common.Address, *types.Transaction, *WindMineToken, error) {
	parsed, err := abi.JSON(strings.NewReader(WindMineTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WindMineTokenBin), backend, _foundersWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WindMineToken{WindMineTokenCaller: WindMineTokenCaller{contract: contract}, WindMineTokenTransactor: WindMineTokenTransactor{contract: contract}, WindMineTokenFilterer: WindMineTokenFilterer{contract: contract}}, nil
}

// WindMineToken is an auto generated Go binding around an Ethereum contract.
type WindMineToken struct {
	WindMineTokenCaller     // Read-only binding to the contract
	WindMineTokenTransactor // Write-only binding to the contract
	WindMineTokenFilterer   // Log filterer for contract events
}

// WindMineTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WindMineTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WindMineTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WindMineTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WindMineTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WindMineTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WindMineTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WindMineTokenSession struct {
	Contract     *WindMineToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WindMineTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WindMineTokenCallerSession struct {
	Contract *WindMineTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WindMineTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WindMineTokenTransactorSession struct {
	Contract     *WindMineTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WindMineTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WindMineTokenRaw struct {
	Contract *WindMineToken // Generic contract binding to access the raw methods on
}

// WindMineTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WindMineTokenCallerRaw struct {
	Contract *WindMineTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WindMineTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WindMineTokenTransactorRaw struct {
	Contract *WindMineTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWindMineToken creates a new instance of WindMineToken, bound to a specific deployed contract.
func NewWindMineToken(address common.Address, backend bind.ContractBackend) (*WindMineToken, error) {
	contract, err := bindWindMineToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WindMineToken{WindMineTokenCaller: WindMineTokenCaller{contract: contract}, WindMineTokenTransactor: WindMineTokenTransactor{contract: contract}, WindMineTokenFilterer: WindMineTokenFilterer{contract: contract}}, nil
}

// NewWindMineTokenCaller creates a new read-only instance of WindMineToken, bound to a specific deployed contract.
func NewWindMineTokenCaller(address common.Address, caller bind.ContractCaller) (*WindMineTokenCaller, error) {
	contract, err := bindWindMineToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenCaller{contract: contract}, nil
}

// NewWindMineTokenTransactor creates a new write-only instance of WindMineToken, bound to a specific deployed contract.
func NewWindMineTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WindMineTokenTransactor, error) {
	contract, err := bindWindMineToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenTransactor{contract: contract}, nil
}

// NewWindMineTokenFilterer creates a new log filterer instance of WindMineToken, bound to a specific deployed contract.
func NewWindMineTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WindMineTokenFilterer, error) {
	contract, err := bindWindMineToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenFilterer{contract: contract}, nil
}

// bindWindMineToken binds a generic wrapper to an already deployed contract.
func bindWindMineToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WindMineTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WindMineToken *WindMineTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WindMineToken.Contract.WindMineTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WindMineToken *WindMineTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WindMineToken.Contract.WindMineTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WindMineToken *WindMineTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WindMineToken.Contract.WindMineTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WindMineToken *WindMineTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WindMineToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WindMineToken *WindMineTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WindMineToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WindMineToken *WindMineTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WindMineToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_WindMineToken *WindMineTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_WindMineToken *WindMineTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WindMineToken.Contract.Allowance(&_WindMineToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_WindMineToken *WindMineTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WindMineToken.Contract.Allowance(&_WindMineToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_WindMineToken *WindMineTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_WindMineToken *WindMineTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _WindMineToken.Contract.BalanceOf(&_WindMineToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_WindMineToken *WindMineTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _WindMineToken.Contract.BalanceOf(&_WindMineToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_WindMineToken *WindMineTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_WindMineToken *WindMineTokenSession) Decimals() (*big.Int, error) {
	return _WindMineToken.Contract.Decimals(&_WindMineToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_WindMineToken *WindMineTokenCallerSession) Decimals() (*big.Int, error) {
	return _WindMineToken.Contract.Decimals(&_WindMineToken.CallOpts)
}

// FoundersReserve is a free data retrieval call binding the contract method 0x7b86120a.
//
// Solidity: function foundersReserve() constant returns(uint256)
func (_WindMineToken *WindMineTokenCaller) FoundersReserve(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "foundersReserve")
	return *ret0, err
}

// FoundersReserve is a free data retrieval call binding the contract method 0x7b86120a.
//
// Solidity: function foundersReserve() constant returns(uint256)
func (_WindMineToken *WindMineTokenSession) FoundersReserve() (*big.Int, error) {
	return _WindMineToken.Contract.FoundersReserve(&_WindMineToken.CallOpts)
}

// FoundersReserve is a free data retrieval call binding the contract method 0x7b86120a.
//
// Solidity: function foundersReserve() constant returns(uint256)
func (_WindMineToken *WindMineTokenCallerSession) FoundersReserve() (*big.Int, error) {
	return _WindMineToken.Contract.FoundersReserve(&_WindMineToken.CallOpts)
}

// FoundersWallet is a free data retrieval call binding the contract method 0x1bfaf155.
//
// Solidity: function foundersWallet() constant returns(address)
func (_WindMineToken *WindMineTokenCaller) FoundersWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "foundersWallet")
	return *ret0, err
}

// FoundersWallet is a free data retrieval call binding the contract method 0x1bfaf155.
//
// Solidity: function foundersWallet() constant returns(address)
func (_WindMineToken *WindMineTokenSession) FoundersWallet() (common.Address, error) {
	return _WindMineToken.Contract.FoundersWallet(&_WindMineToken.CallOpts)
}

// FoundersWallet is a free data retrieval call binding the contract method 0x1bfaf155.
//
// Solidity: function foundersWallet() constant returns(address)
func (_WindMineToken *WindMineTokenCallerSession) FoundersWallet() (common.Address, error) {
	return _WindMineToken.Contract.FoundersWallet(&_WindMineToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WindMineToken *WindMineTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WindMineToken *WindMineTokenSession) Name() (string, error) {
	return _WindMineToken.Contract.Name(&_WindMineToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WindMineToken *WindMineTokenCallerSession) Name() (string, error) {
	return _WindMineToken.Contract.Name(&_WindMineToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WindMineToken *WindMineTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WindMineToken *WindMineTokenSession) Owner() (common.Address, error) {
	return _WindMineToken.Contract.Owner(&_WindMineToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WindMineToken *WindMineTokenCallerSession) Owner() (common.Address, error) {
	return _WindMineToken.Contract.Owner(&_WindMineToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WindMineToken *WindMineTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WindMineToken *WindMineTokenSession) Symbol() (string, error) {
	return _WindMineToken.Contract.Symbol(&_WindMineToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_WindMineToken *WindMineTokenCallerSession) Symbol() (string, error) {
	return _WindMineToken.Contract.Symbol(&_WindMineToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WindMineToken *WindMineTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WindMineToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WindMineToken *WindMineTokenSession) TotalSupply() (*big.Int, error) {
	return _WindMineToken.Contract.TotalSupply(&_WindMineToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_WindMineToken *WindMineTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WindMineToken.Contract.TotalSupply(&_WindMineToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.Approve(&_WindMineToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.Approve(&_WindMineToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.DecreaseApproval(&_WindMineToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.DecreaseApproval(&_WindMineToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.IncreaseApproval(&_WindMineToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.IncreaseApproval(&_WindMineToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.Transfer(&_WindMineToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.Transfer(&_WindMineToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.TransferFrom(&_WindMineToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_WindMineToken *WindMineTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WindMineToken.Contract.TransferFrom(&_WindMineToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WindMineToken *WindMineTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WindMineToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WindMineToken *WindMineTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WindMineToken.Contract.TransferOwnership(&_WindMineToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WindMineToken *WindMineTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WindMineToken.Contract.TransferOwnership(&_WindMineToken.TransactOpts, newOwner)
}

// WindMineTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WindMineToken contract.
type WindMineTokenApprovalIterator struct {
	Event *WindMineTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WindMineTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WindMineTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WindMineTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WindMineTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WindMineTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WindMineTokenApproval represents a Approval event raised by the WindMineToken contract.
type WindMineTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_WindMineToken *WindMineTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WindMineTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WindMineToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenApprovalIterator{contract: _WindMineToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_WindMineToken *WindMineTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WindMineTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WindMineToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WindMineTokenApproval)
				if err := _WindMineToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WindMineTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WindMineToken contract.
type WindMineTokenOwnershipTransferredIterator struct {
	Event *WindMineTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WindMineTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WindMineTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WindMineTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WindMineTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WindMineTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WindMineTokenOwnershipTransferred represents a OwnershipTransferred event raised by the WindMineToken contract.
type WindMineTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_WindMineToken *WindMineTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WindMineTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WindMineToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenOwnershipTransferredIterator{contract: _WindMineToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_WindMineToken *WindMineTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WindMineTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WindMineToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WindMineTokenOwnershipTransferred)
				if err := _WindMineToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// WindMineTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WindMineToken contract.
type WindMineTokenTransferIterator struct {
	Event *WindMineTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WindMineTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WindMineTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WindMineTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WindMineTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WindMineTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WindMineTokenTransfer represents a Transfer event raised by the WindMineToken contract.
type WindMineTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_WindMineToken *WindMineTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WindMineTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WindMineToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WindMineTokenTransferIterator{contract: _WindMineToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_WindMineToken *WindMineTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WindMineTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WindMineToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WindMineTokenTransfer)
				if err := _WindMineToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
