// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapper

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

// GopherConABI is the input ABI used to generate the binding from.
const GopherConABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numTargets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hitList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"hitMe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"Thump\",\"type\":\"event\"}]"

// GopherConBin is the compiled bytecode used for deploying new contracts.
const GopherConBin = `0x6060604052341561000f57600080fd5b6102708061001e6000396000f3006060604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663361f0b0e8114610066578063508704441461008b578063fbe8dde6146100bd578063fdf3a8f4146100d2575b600080fd5b341561007157600080fd5b6100796100f1565b60405190815260200160405180910390f35b341561009657600080fd5b6100a16004356100f8565b604051600160a060020a03909116815260200160405180910390f35b34156100c857600080fd5b6100d0610120565b005b34156100dd57600080fd5b610079600160a060020a03600435166101eb565b6001545b90565b600180548290811061010657fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a0333166000908152602081905260409020541515610183576001805480820161015083826101fd565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790555b33600160a060020a038116600090815260208190526040908190208054600101908190557f8562beded59a818307bdec7ab233f9b8b94d73b8b4bb1bbd7e055e049411503e929151600160a060020a03909216825260208201526040908101905180910390a1565b60006020819052908152604090205481565b81548183558181151161022157600083815260209020610221918101908301610226565b505050565b6100f591905b80821115610240576000815560010161022c565b50905600a165627a7a723058205988984a79eaeba1a6694a5ebc8f010615aba3cd34cc017b39aa9b75f653bc9c0029`

// DeployGopherCon deploys a new Ethereum contract, binding an instance of GopherCon to it.
func DeployGopherCon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GopherCon, error) {
	parsed, err := abi.JSON(strings.NewReader(GopherConABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GopherConBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GopherCon{GopherConCaller: GopherConCaller{contract: contract}, GopherConTransactor: GopherConTransactor{contract: contract}, GopherConFilterer: GopherConFilterer{contract: contract}}, nil
}

// GopherCon is an auto generated Go binding around an Ethereum contract.
type GopherCon struct {
	GopherConCaller     // Read-only binding to the contract
	GopherConTransactor // Write-only binding to the contract
	GopherConFilterer   // Log filterer for contract events
}

// GopherConCaller is an auto generated read-only Go binding around an Ethereum contract.
type GopherConCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GopherConTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GopherConTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GopherConFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GopherConFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GopherConSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GopherConSession struct {
	Contract     *GopherCon        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GopherConCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GopherConCallerSession struct {
	Contract *GopherConCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GopherConTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GopherConTransactorSession struct {
	Contract     *GopherConTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GopherConRaw is an auto generated low-level Go binding around an Ethereum contract.
type GopherConRaw struct {
	Contract *GopherCon // Generic contract binding to access the raw methods on
}

// GopherConCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GopherConCallerRaw struct {
	Contract *GopherConCaller // Generic read-only contract binding to access the raw methods on
}

// GopherConTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GopherConTransactorRaw struct {
	Contract *GopherConTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGopherCon creates a new instance of GopherCon, bound to a specific deployed contract.
func NewGopherCon(address common.Address, backend bind.ContractBackend) (*GopherCon, error) {
	contract, err := bindGopherCon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GopherCon{GopherConCaller: GopherConCaller{contract: contract}, GopherConTransactor: GopherConTransactor{contract: contract}, GopherConFilterer: GopherConFilterer{contract: contract}}, nil
}

// NewGopherConCaller creates a new read-only instance of GopherCon, bound to a specific deployed contract.
func NewGopherConCaller(address common.Address, caller bind.ContractCaller) (*GopherConCaller, error) {
	contract, err := bindGopherCon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GopherConCaller{contract: contract}, nil
}

// NewGopherConTransactor creates a new write-only instance of GopherCon, bound to a specific deployed contract.
func NewGopherConTransactor(address common.Address, transactor bind.ContractTransactor) (*GopherConTransactor, error) {
	contract, err := bindGopherCon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GopherConTransactor{contract: contract}, nil
}

// NewGopherConFilterer creates a new log filterer instance of GopherCon, bound to a specific deployed contract.
func NewGopherConFilterer(address common.Address, filterer bind.ContractFilterer) (*GopherConFilterer, error) {
	contract, err := bindGopherCon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GopherConFilterer{contract: contract}, nil
}

// bindGopherCon binds a generic wrapper to an already deployed contract.
func bindGopherCon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GopherConABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GopherCon *GopherConRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GopherCon.Contract.GopherConCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GopherCon *GopherConRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GopherCon.Contract.GopherConTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GopherCon *GopherConRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GopherCon.Contract.GopherConTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GopherCon *GopherConCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GopherCon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GopherCon *GopherConTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GopherCon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GopherCon *GopherConTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GopherCon.Contract.contract.Transact(opts, method, params...)
}

// HitList is a free data retrieval call binding the contract method 0x50870444.
//
// Solidity: function hitList( uint256) constant returns(address)
func (_GopherCon *GopherConCaller) HitList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GopherCon.contract.Call(opts, out, "hitList", arg0)
	return *ret0, err
}

// HitList is a free data retrieval call binding the contract method 0x50870444.
//
// Solidity: function hitList( uint256) constant returns(address)
func (_GopherCon *GopherConSession) HitList(arg0 *big.Int) (common.Address, error) {
	return _GopherCon.Contract.HitList(&_GopherCon.CallOpts, arg0)
}

// HitList is a free data retrieval call binding the contract method 0x50870444.
//
// Solidity: function hitList( uint256) constant returns(address)
func (_GopherCon *GopherConCallerSession) HitList(arg0 *big.Int) (common.Address, error) {
	return _GopherCon.Contract.HitList(&_GopherCon.CallOpts, arg0)
}

// Hits is a free data retrieval call binding the contract method 0xfdf3a8f4.
//
// Solidity: function hits( address) constant returns(uint256)
func (_GopherCon *GopherConCaller) Hits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GopherCon.contract.Call(opts, out, "hits", arg0)
	return *ret0, err
}

// Hits is a free data retrieval call binding the contract method 0xfdf3a8f4.
//
// Solidity: function hits( address) constant returns(uint256)
func (_GopherCon *GopherConSession) Hits(arg0 common.Address) (*big.Int, error) {
	return _GopherCon.Contract.Hits(&_GopherCon.CallOpts, arg0)
}

// Hits is a free data retrieval call binding the contract method 0xfdf3a8f4.
//
// Solidity: function hits( address) constant returns(uint256)
func (_GopherCon *GopherConCallerSession) Hits(arg0 common.Address) (*big.Int, error) {
	return _GopherCon.Contract.Hits(&_GopherCon.CallOpts, arg0)
}

// NumTargets is a free data retrieval call binding the contract method 0x361f0b0e.
//
// Solidity: function numTargets() constant returns(uint256)
func (_GopherCon *GopherConCaller) NumTargets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GopherCon.contract.Call(opts, out, "numTargets")
	return *ret0, err
}

// NumTargets is a free data retrieval call binding the contract method 0x361f0b0e.
//
// Solidity: function numTargets() constant returns(uint256)
func (_GopherCon *GopherConSession) NumTargets() (*big.Int, error) {
	return _GopherCon.Contract.NumTargets(&_GopherCon.CallOpts)
}

// NumTargets is a free data retrieval call binding the contract method 0x361f0b0e.
//
// Solidity: function numTargets() constant returns(uint256)
func (_GopherCon *GopherConCallerSession) NumTargets() (*big.Int, error) {
	return _GopherCon.Contract.NumTargets(&_GopherCon.CallOpts)
}

// HitMe is a paid mutator transaction binding the contract method 0xfbe8dde6.
//
// Solidity: function hitMe() returns()
func (_GopherCon *GopherConTransactor) HitMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GopherCon.contract.Transact(opts, "hitMe")
}

// HitMe is a paid mutator transaction binding the contract method 0xfbe8dde6.
//
// Solidity: function hitMe() returns()
func (_GopherCon *GopherConSession) HitMe() (*types.Transaction, error) {
	return _GopherCon.Contract.HitMe(&_GopherCon.TransactOpts)
}

// HitMe is a paid mutator transaction binding the contract method 0xfbe8dde6.
//
// Solidity: function hitMe() returns()
func (_GopherCon *GopherConTransactorSession) HitMe() (*types.Transaction, error) {
	return _GopherCon.Contract.HitMe(&_GopherCon.TransactOpts)
}

// GopherConThumpIterator is returned from FilterThump and is used to iterate over the raw logs and unpacked data for Thump events raised by the GopherCon contract.
type GopherConThumpIterator struct {
	Event *GopherConThump // Event containing the contract specifics and raw log

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
func (it *GopherConThumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GopherConThump)
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
		it.Event = new(GopherConThump)
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
func (it *GopherConThumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GopherConThumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GopherConThump represents a Thump event raised by the GopherCon contract.
type GopherConThump struct {
	Who   common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterThump is a free log retrieval operation binding the contract event 0x8562beded59a818307bdec7ab233f9b8b94d73b8b4bb1bbd7e055e049411503e.
//
// Solidity: event Thump(who address, count uint256)
func (_GopherCon *GopherConFilterer) FilterThump(opts *bind.FilterOpts) (*GopherConThumpIterator, error) {

	logs, sub, err := _GopherCon.contract.FilterLogs(opts, "Thump")
	if err != nil {
		return nil, err
	}
	return &GopherConThumpIterator{contract: _GopherCon.contract, event: "Thump", logs: logs, sub: sub}, nil
}

// WatchThump is a free log subscription operation binding the contract event 0x8562beded59a818307bdec7ab233f9b8b94d73b8b4bb1bbd7e055e049411503e.
//
// Solidity: event Thump(who address, count uint256)
func (_GopherCon *GopherConFilterer) WatchThump(opts *bind.WatchOpts, sink chan<- *GopherConThump) (event.Subscription, error) {

	logs, sub, err := _GopherCon.contract.WatchLogs(opts, "Thump")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GopherConThump)
				if err := _GopherCon.contract.UnpackLog(event, "Thump", log); err != nil {
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
