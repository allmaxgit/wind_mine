// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
