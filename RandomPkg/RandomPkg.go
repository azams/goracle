// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RandomPkg

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

// RandomPkgMetaData contains all meta data concerning the RandomPkg contract.
var RandomPkgMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NewRandomRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randomNumberMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomNumber\",\"type\":\"uint256\"}],\"name\":\"setRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RandomPkgABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomPkgMetaData.ABI instead.
var RandomPkgABI = RandomPkgMetaData.ABI

// RandomPkg is an auto generated Go binding around an Ethereum contract.
type RandomPkg struct {
	RandomPkgCaller     // Read-only binding to the contract
	RandomPkgTransactor // Write-only binding to the contract
	RandomPkgFilterer   // Log filterer for contract events
}

// RandomPkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomPkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomPkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomPkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomPkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomPkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomPkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomPkgSession struct {
	Contract     *RandomPkg        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomPkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomPkgCallerSession struct {
	Contract *RandomPkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RandomPkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomPkgTransactorSession struct {
	Contract     *RandomPkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RandomPkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomPkgRaw struct {
	Contract *RandomPkg // Generic contract binding to access the raw methods on
}

// RandomPkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomPkgCallerRaw struct {
	Contract *RandomPkgCaller // Generic read-only contract binding to access the raw methods on
}

// RandomPkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomPkgTransactorRaw struct {
	Contract *RandomPkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomPkg creates a new instance of RandomPkg, bound to a specific deployed contract.
func NewRandomPkg(address common.Address, backend bind.ContractBackend) (*RandomPkg, error) {
	contract, err := bindRandomPkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomPkg{RandomPkgCaller: RandomPkgCaller{contract: contract}, RandomPkgTransactor: RandomPkgTransactor{contract: contract}, RandomPkgFilterer: RandomPkgFilterer{contract: contract}}, nil
}

// NewRandomPkgCaller creates a new read-only instance of RandomPkg, bound to a specific deployed contract.
func NewRandomPkgCaller(address common.Address, caller bind.ContractCaller) (*RandomPkgCaller, error) {
	contract, err := bindRandomPkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomPkgCaller{contract: contract}, nil
}

// NewRandomPkgTransactor creates a new write-only instance of RandomPkg, bound to a specific deployed contract.
func NewRandomPkgTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomPkgTransactor, error) {
	contract, err := bindRandomPkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomPkgTransactor{contract: contract}, nil
}

// NewRandomPkgFilterer creates a new log filterer instance of RandomPkg, bound to a specific deployed contract.
func NewRandomPkgFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomPkgFilterer, error) {
	contract, err := bindRandomPkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomPkgFilterer{contract: contract}, nil
}

// bindRandomPkg binds a generic wrapper to an already deployed contract.
func bindRandomPkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomPkgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomPkg *RandomPkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomPkg.Contract.RandomPkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomPkg *RandomPkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomPkg.Contract.RandomPkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomPkg *RandomPkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomPkg.Contract.RandomPkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomPkg *RandomPkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomPkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomPkg *RandomPkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomPkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomPkg *RandomPkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomPkg.Contract.contract.Transact(opts, method, params...)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_RandomPkg *RandomPkgCaller) GetContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomPkg.contract.Call(opts, &out, "getContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_RandomPkg *RandomPkgSession) GetContractAddress() (common.Address, error) {
	return _RandomPkg.Contract.GetContractAddress(&_RandomPkg.CallOpts)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x32a2c5d0.
//
// Solidity: function getContractAddress() view returns(address)
func (_RandomPkg *RandomPkgCallerSession) GetContractAddress() (common.Address, error) {
	return _RandomPkg.Contract.GetContractAddress(&_RandomPkg.CallOpts)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0xb37217a4.
//
// Solidity: function getRandomNumber(uint256 requestId) view returns(uint256)
func (_RandomPkg *RandomPkgCaller) GetRandomNumber(opts *bind.CallOpts, requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomPkg.contract.Call(opts, &out, "getRandomNumber", requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandomNumber is a free data retrieval call binding the contract method 0xb37217a4.
//
// Solidity: function getRandomNumber(uint256 requestId) view returns(uint256)
func (_RandomPkg *RandomPkgSession) GetRandomNumber(requestId *big.Int) (*big.Int, error) {
	return _RandomPkg.Contract.GetRandomNumber(&_RandomPkg.CallOpts, requestId)
}

// GetRandomNumber is a free data retrieval call binding the contract method 0xb37217a4.
//
// Solidity: function getRandomNumber(uint256 requestId) view returns(uint256)
func (_RandomPkg *RandomPkgCallerSession) GetRandomNumber(requestId *big.Int) (*big.Int, error) {
	return _RandomPkg.Contract.GetRandomNumber(&_RandomPkg.CallOpts, requestId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomPkg *RandomPkgCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomPkg.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomPkg *RandomPkgSession) Owner() (common.Address, error) {
	return _RandomPkg.Contract.Owner(&_RandomPkg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomPkg *RandomPkgCallerSession) Owner() (common.Address, error) {
	return _RandomPkg.Contract.Owner(&_RandomPkg.CallOpts)
}

// RandomNumberMapping is a free data retrieval call binding the contract method 0x34e52fa8.
//
// Solidity: function randomNumberMapping(uint256 ) view returns(uint256)
func (_RandomPkg *RandomPkgCaller) RandomNumberMapping(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomPkg.contract.Call(opts, &out, "randomNumberMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomNumberMapping is a free data retrieval call binding the contract method 0x34e52fa8.
//
// Solidity: function randomNumberMapping(uint256 ) view returns(uint256)
func (_RandomPkg *RandomPkgSession) RandomNumberMapping(arg0 *big.Int) (*big.Int, error) {
	return _RandomPkg.Contract.RandomNumberMapping(&_RandomPkg.CallOpts, arg0)
}

// RandomNumberMapping is a free data retrieval call binding the contract method 0x34e52fa8.
//
// Solidity: function randomNumberMapping(uint256 ) view returns(uint256)
func (_RandomPkg *RandomPkgCallerSession) RandomNumberMapping(arg0 *big.Int) (*big.Int, error) {
	return _RandomPkg.Contract.RandomNumberMapping(&_RandomPkg.CallOpts, arg0)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0x8678a7b2.
//
// Solidity: function requestRandomNumber() returns()
func (_RandomPkg *RandomPkgTransactor) RequestRandomNumber(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomPkg.contract.Transact(opts, "requestRandomNumber")
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0x8678a7b2.
//
// Solidity: function requestRandomNumber() returns()
func (_RandomPkg *RandomPkgSession) RequestRandomNumber() (*types.Transaction, error) {
	return _RandomPkg.Contract.RequestRandomNumber(&_RandomPkg.TransactOpts)
}

// RequestRandomNumber is a paid mutator transaction binding the contract method 0x8678a7b2.
//
// Solidity: function requestRandomNumber() returns()
func (_RandomPkg *RandomPkgTransactorSession) RequestRandomNumber() (*types.Transaction, error) {
	return _RandomPkg.Contract.RequestRandomNumber(&_RandomPkg.TransactOpts)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x5da4fabc.
//
// Solidity: function setRandomNumber(uint256 requestId, uint256 _randomNumber) returns()
func (_RandomPkg *RandomPkgTransactor) SetRandomNumber(opts *bind.TransactOpts, requestId *big.Int, _randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomPkg.contract.Transact(opts, "setRandomNumber", requestId, _randomNumber)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x5da4fabc.
//
// Solidity: function setRandomNumber(uint256 requestId, uint256 _randomNumber) returns()
func (_RandomPkg *RandomPkgSession) SetRandomNumber(requestId *big.Int, _randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomPkg.Contract.SetRandomNumber(&_RandomPkg.TransactOpts, requestId, _randomNumber)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x5da4fabc.
//
// Solidity: function setRandomNumber(uint256 requestId, uint256 _randomNumber) returns()
func (_RandomPkg *RandomPkgTransactorSession) SetRandomNumber(requestId *big.Int, _randomNumber *big.Int) (*types.Transaction, error) {
	return _RandomPkg.Contract.SetRandomNumber(&_RandomPkg.TransactOpts, requestId, _randomNumber)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newAddress) returns()
func (_RandomPkg *RandomPkgTransactor) UpdateOwner(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _RandomPkg.contract.Transact(opts, "updateOwner", newAddress)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newAddress) returns()
func (_RandomPkg *RandomPkgSession) UpdateOwner(newAddress common.Address) (*types.Transaction, error) {
	return _RandomPkg.Contract.UpdateOwner(&_RandomPkg.TransactOpts, newAddress)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address newAddress) returns()
func (_RandomPkg *RandomPkgTransactorSession) UpdateOwner(newAddress common.Address) (*types.Transaction, error) {
	return _RandomPkg.Contract.UpdateOwner(&_RandomPkg.TransactOpts, newAddress)
}

// RandomPkgNewRandomRequestIterator is returned from FilterNewRandomRequest and is used to iterate over the raw logs and unpacked data for NewRandomRequest events raised by the RandomPkg contract.
type RandomPkgNewRandomRequestIterator struct {
	Event *RandomPkgNewRandomRequest // Event containing the contract specifics and raw log

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
func (it *RandomPkgNewRandomRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomPkgNewRandomRequest)
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
		it.Event = new(RandomPkgNewRandomRequest)
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
func (it *RandomPkgNewRandomRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomPkgNewRandomRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomPkgNewRandomRequest represents a NewRandomRequest event raised by the RandomPkg contract.
type RandomPkgNewRandomRequest struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRandomRequest is a free log retrieval operation binding the contract event 0xa151a34dd00e4be61a468428732ad695c6c3b69f300a82003df4aae1a760a5fe.
//
// Solidity: event NewRandomRequest(uint256 requestId)
func (_RandomPkg *RandomPkgFilterer) FilterNewRandomRequest(opts *bind.FilterOpts) (*RandomPkgNewRandomRequestIterator, error) {

	logs, sub, err := _RandomPkg.contract.FilterLogs(opts, "NewRandomRequest")
	if err != nil {
		return nil, err
	}
	return &RandomPkgNewRandomRequestIterator{contract: _RandomPkg.contract, event: "NewRandomRequest", logs: logs, sub: sub}, nil
}

// WatchNewRandomRequest is a free log subscription operation binding the contract event 0xa151a34dd00e4be61a468428732ad695c6c3b69f300a82003df4aae1a760a5fe.
//
// Solidity: event NewRandomRequest(uint256 requestId)
func (_RandomPkg *RandomPkgFilterer) WatchNewRandomRequest(opts *bind.WatchOpts, sink chan<- *RandomPkgNewRandomRequest) (event.Subscription, error) {

	logs, sub, err := _RandomPkg.contract.WatchLogs(opts, "NewRandomRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomPkgNewRandomRequest)
				if err := _RandomPkg.contract.UnpackLog(event, "NewRandomRequest", log); err != nil {
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

// ParseNewRandomRequest is a log parse operation binding the contract event 0xa151a34dd00e4be61a468428732ad695c6c3b69f300a82003df4aae1a760a5fe.
//
// Solidity: event NewRandomRequest(uint256 requestId)
func (_RandomPkg *RandomPkgFilterer) ParseNewRandomRequest(log types.Log) (*RandomPkgNewRandomRequest, error) {
	event := new(RandomPkgNewRandomRequest)
	if err := _RandomPkg.contract.UnpackLog(event, "NewRandomRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
