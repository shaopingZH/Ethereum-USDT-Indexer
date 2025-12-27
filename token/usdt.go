// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// UsdtMetaData contains all meta data concerning the Usdt contract.
var UsdtMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// UsdtABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdtMetaData.ABI instead.
var UsdtABI = UsdtMetaData.ABI

// Usdt is an auto generated Go binding around an Ethereum contract.
type Usdt struct {
	UsdtCaller     // Read-only binding to the contract
	UsdtTransactor // Write-only binding to the contract
	UsdtFilterer   // Log filterer for contract events
}

// UsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtSession struct {
	Contract     *Usdt             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtCallerSession struct {
	Contract *UsdtCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtTransactorSession struct {
	Contract     *UsdtTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtRaw struct {
	Contract *Usdt // Generic contract binding to access the raw methods on
}

// UsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtCallerRaw struct {
	Contract *UsdtCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtTransactorRaw struct {
	Contract *UsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdt creates a new instance of Usdt, bound to a specific deployed contract.
func NewUsdt(address common.Address, backend bind.ContractBackend) (*Usdt, error) {
	contract, err := bindUsdt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdt{UsdtCaller: UsdtCaller{contract: contract}, UsdtTransactor: UsdtTransactor{contract: contract}, UsdtFilterer: UsdtFilterer{contract: contract}}, nil
}

// NewUsdtCaller creates a new read-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtCaller(address common.Address, caller bind.ContractCaller) (*UsdtCaller, error) {
	contract, err := bindUsdt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtCaller{contract: contract}, nil
}

// NewUsdtTransactor creates a new write-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtTransactor, error) {
	contract, err := bindUsdt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtTransactor{contract: contract}, nil
}

// NewUsdtFilterer creates a new log filterer instance of Usdt, bound to a specific deployed contract.
func NewUsdtFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdtFilterer, error) {
	contract, err := bindUsdt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdtFilterer{contract: contract}, nil
}

// bindUsdt binds a generic wrapper to an already deployed contract.
func bindUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsdtMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.UsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtCallerSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtCallerSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtCallerSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtCallerSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdt *UsdtSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Usdt *UsdtTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, _from, _to, _value)
}

// UsdtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Usdt contract.
type UsdtTransferIterator struct {
	Event *UsdtTransfer // Event containing the contract specifics and raw log

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
func (it *UsdtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtTransfer)
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
		it.Event = new(UsdtTransfer)
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
func (it *UsdtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtTransfer represents a Transfer event raised by the Usdt contract.
type UsdtTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdt *UsdtFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UsdtTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UsdtTransferIterator{contract: _Usdt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdt *UsdtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UsdtTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Usdt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtTransfer)
				if err := _Usdt.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Usdt *UsdtFilterer) ParseTransfer(log types.Log) (*UsdtTransfer, error) {
	event := new(UsdtTransfer)
	if err := _Usdt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
