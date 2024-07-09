// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PricePkg

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PricePkgMetaData contains all meta data concerning the PricePkg contract.
var PricePkgMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"coin\",\"type\":\"string\"}],\"name\":\"NewTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_coin\",\"type\":\"string\"}],\"name\":\"getCoinPrice\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_coin\",\"type\":\"string\"}],\"name\":\"requestUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PricePkgABI is the input ABI used to generate the binding from.
// Deprecated: Use PricePkgMetaData.ABI instead.
var PricePkgABI = PricePkgMetaData.ABI

// PricePkg is an auto generated Go binding around an Ethereum contract.
type PricePkg struct {
	PricePkgCaller     // Read-only binding to the contract
	PricePkgTransactor // Write-only binding to the contract
	PricePkgFilterer   // Log filterer for contract events
}

// PricePkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type PricePkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricePkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PricePkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricePkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PricePkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricePkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricePkgSession struct {
	Contract     *PricePkg         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricePkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricePkgCallerSession struct {
	Contract *PricePkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PricePkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricePkgTransactorSession struct {
	Contract     *PricePkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PricePkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type PricePkgRaw struct {
	Contract *PricePkg // Generic contract binding to access the raw methods on
}

// PricePkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricePkgCallerRaw struct {
	Contract *PricePkgCaller // Generic read-only contract binding to access the raw methods on
}

// PricePkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricePkgTransactorRaw struct {
	Contract *PricePkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPricePkg creates a new instance of PricePkg, bound to a specific deployed contract.
func NewPricePkg(address common.Address, backend bind.ContractBackend) (*PricePkg, error) {
	contract, err := bindPricePkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PricePkg{PricePkgCaller: PricePkgCaller{contract: contract}, PricePkgTransactor: PricePkgTransactor{contract: contract}, PricePkgFilterer: PricePkgFilterer{contract: contract}}, nil
}

// NewPricePkgCaller creates a new read-only instance of PricePkg, bound to a specific deployed contract.
func NewPricePkgCaller(address common.Address, caller bind.ContractCaller) (*PricePkgCaller, error) {
	contract, err := bindPricePkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PricePkgCaller{contract: contract}, nil
}

// NewPricePkgTransactor creates a new write-only instance of PricePkg, bound to a specific deployed contract.
func NewPricePkgTransactor(address common.Address, transactor bind.ContractTransactor) (*PricePkgTransactor, error) {
	contract, err := bindPricePkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PricePkgTransactor{contract: contract}, nil
}

// NewPricePkgFilterer creates a new log filterer instance of PricePkg, bound to a specific deployed contract.
func NewPricePkgFilterer(address common.Address, filterer bind.ContractFilterer) (*PricePkgFilterer, error) {
	contract, err := bindPricePkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PricePkgFilterer{contract: contract}, nil
}

// bindPricePkg binds a generic wrapper to an already deployed contract.
func bindPricePkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PricePkgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PricePkg *PricePkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PricePkg.Contract.PricePkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PricePkg *PricePkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PricePkg.Contract.PricePkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PricePkg *PricePkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PricePkg.Contract.PricePkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PricePkg *PricePkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PricePkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PricePkg *PricePkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PricePkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PricePkg *PricePkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PricePkg.Contract.contract.Transact(opts, method, params...)
}

// GetCoinPrice is a free data retrieval call binding the contract method 0xa9d524e3.
//
// Solidity: function getCoinPrice(string _coin) view returns(string)
func (_PricePkg *PricePkgCaller) GetCoinPrice(opts *bind.CallOpts, _coin string) (string, error) {
	var out []interface{}
	err := _PricePkg.contract.Call(opts, &out, "getCoinPrice", _coin)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCoinPrice is a free data retrieval call binding the contract method 0xa9d524e3.
//
// Solidity: function getCoinPrice(string _coin) view returns(string)
func (_PricePkg *PricePkgSession) GetCoinPrice(_coin string) (string, error) {
	return _PricePkg.Contract.GetCoinPrice(&_PricePkg.CallOpts, _coin)
}

// GetCoinPrice is a free data retrieval call binding the contract method 0xa9d524e3.
//
// Solidity: function getCoinPrice(string _coin) view returns(string)
func (_PricePkg *PricePkgCallerSession) GetCoinPrice(_coin string) (string, error) {
	return _PricePkg.Contract.GetCoinPrice(&_PricePkg.CallOpts, _coin)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_PricePkg *PricePkgCaller) GetContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PricePkg.contract.Call(opts, &out, "getContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_PricePkg *PricePkgSession) GetContractAddress() (common.Address, error) {
	return _PricePkg.Contract.GetContractAddress(&_PricePkg.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_PricePkg *PricePkgCallerSession) GetContractAddress() (common.Address, error) {
	return _PricePkg.Contract.GetContractAddress(&_PricePkg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PricePkg *PricePkgCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PricePkg.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PricePkg *PricePkgSession) Owner() (common.Address, error) {
	return _PricePkg.Contract.Owner(&_PricePkg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PricePkg *PricePkgCallerSession) Owner() (common.Address, error) {
	return _PricePkg.Contract.Owner(&_PricePkg.CallOpts)
}

// RequestUpdate is a paid mutator transaction binding the contract method 0x69b43d07.
//
// Solidity: function requestUpdate(string _coin) returns()
func (_PricePkg *PricePkgTransactor) RequestUpdate(opts *bind.TransactOpts, _coin string) (*types.Transaction, error) {
	return _PricePkg.contract.Transact(opts, "requestUpdate", _coin)
}

// RequestUpdate is a paid mutator transaction binding the contract method 0x69b43d07.
//
// Solidity: function requestUpdate(string _coin) returns()
func (_PricePkg *PricePkgSession) RequestUpdate(_coin string) (*types.Transaction, error) {
	return _PricePkg.Contract.RequestUpdate(&_PricePkg.TransactOpts, _coin)
}

// RequestUpdate is a paid mutator transaction binding the contract method 0x69b43d07.
//
// Solidity: function requestUpdate(string _coin) returns()
func (_PricePkg *PricePkgTransactorSession) RequestUpdate(_coin string) (*types.Transaction, error) {
	return _PricePkg.Contract.RequestUpdate(&_PricePkg.TransactOpts, _coin)
}

// PricePkgNewTransactionIterator is returned from FilterNewTransaction and is used to iterate over the raw logs and unpacked data for NewTransaction events raised by the PricePkg contract.
type PricePkgNewTransactionIterator struct {
	Event *PricePkgNewTransaction // Event containing the contract specifics and raw log

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
func (it *PricePkgNewTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricePkgNewTransaction)
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
		it.Event = new(PricePkgNewTransaction)
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
func (it *PricePkgNewTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricePkgNewTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricePkgNewTransaction represents a NewTransaction event raised by the PricePkg contract.
type PricePkgNewTransaction struct {
	Coin string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTransaction is a free log retrieval operation binding the contract event 0x08f81b6e9109fad00df4c1e4c8017ba4cb67013034439d71e57399c83aa23f3c.
//
// Solidity: event NewTransaction(string coin)
func (_PricePkg *PricePkgFilterer) FilterNewTransaction(opts *bind.FilterOpts) (*PricePkgNewTransactionIterator, error) {

	logs, sub, err := _PricePkg.contract.FilterLogs(opts, "NewTransaction")
	if err != nil {
		return nil, err
	}
	return &PricePkgNewTransactionIterator{contract: _PricePkg.contract, event: "NewTransaction", logs: logs, sub: sub}, nil
}

// WatchNewTransaction is a free log subscription operation binding the contract event 0x08f81b6e9109fad00df4c1e4c8017ba4cb67013034439d71e57399c83aa23f3c.
//
// Solidity: event NewTransaction(string coin)
func (_PricePkg *PricePkgFilterer) WatchNewTransaction(opts *bind.WatchOpts, sink chan<- *PricePkgNewTransaction) (event.Subscription, error) {

	logs, sub, err := _PricePkg.contract.WatchLogs(opts, "NewTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PricePkgNewTransaction)
				if err := _PricePkg.contract.UnpackLog(event, "NewTransaction", log); err != nil {
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

// ParseNewTransaction is a log parse operation binding the contract event 0x08f81b6e9109fad00df4c1e4c8017ba4cb67013034439d71e57399c83aa23f3c.
//
// Solidity: event NewTransaction(string coin)
func (_PricePkg *PricePkgFilterer) ParseNewTransaction(log types.Log) (*PricePkgNewTransaction, error) {
	event := new(PricePkgNewTransaction)
	if err := _PricePkg.contract.UnpackLog(event, "NewTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
