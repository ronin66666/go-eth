// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdt_deposit

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
)

// UsdtDepositMetaData contains all meta data concerning the UsdtDeposit contract.
var UsdtDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiveAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveAddr\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpasue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UsdtDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use UsdtDepositMetaData.ABI instead.
var UsdtDepositABI = UsdtDepositMetaData.ABI

// UsdtDeposit is an auto generated Go binding around an Ethereum contract.
type UsdtDeposit struct {
	UsdtDepositCaller     // Read-only binding to the contract
	UsdtDepositTransactor // Write-only binding to the contract
	UsdtDepositFilterer   // Log filterer for contract events
}

// UsdtDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdtDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtDepositSession struct {
	Contract     *UsdtDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtDepositCallerSession struct {
	Contract *UsdtDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UsdtDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtDepositTransactorSession struct {
	Contract     *UsdtDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UsdtDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtDepositRaw struct {
	Contract *UsdtDeposit // Generic contract binding to access the raw methods on
}

// UsdtDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtDepositCallerRaw struct {
	Contract *UsdtDepositCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtDepositTransactorRaw struct {
	Contract *UsdtDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdtDeposit creates a new instance of UsdtDeposit, bound to a specific deployed contract.
func NewUsdtDeposit(address common.Address, backend bind.ContractBackend) (*UsdtDeposit, error) {
	contract, err := bindUsdtDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsdtDeposit{UsdtDepositCaller: UsdtDepositCaller{contract: contract}, UsdtDepositTransactor: UsdtDepositTransactor{contract: contract}, UsdtDepositFilterer: UsdtDepositFilterer{contract: contract}}, nil
}

// NewUsdtDepositCaller creates a new read-only instance of UsdtDeposit, bound to a specific deployed contract.
func NewUsdtDepositCaller(address common.Address, caller bind.ContractCaller) (*UsdtDepositCaller, error) {
	contract, err := bindUsdtDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtDepositCaller{contract: contract}, nil
}

// NewUsdtDepositTransactor creates a new write-only instance of UsdtDeposit, bound to a specific deployed contract.
func NewUsdtDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtDepositTransactor, error) {
	contract, err := bindUsdtDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtDepositTransactor{contract: contract}, nil
}

// NewUsdtDepositFilterer creates a new log filterer instance of UsdtDeposit, bound to a specific deployed contract.
func NewUsdtDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdtDepositFilterer, error) {
	contract, err := bindUsdtDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdtDepositFilterer{contract: contract}, nil
}

// bindUsdtDeposit binds a generic wrapper to an already deployed contract.
func bindUsdtDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdtDeposit *UsdtDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdtDeposit.Contract.UsdtDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdtDeposit *UsdtDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.UsdtDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdtDeposit *UsdtDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.UsdtDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsdtDeposit *UsdtDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsdtDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsdtDeposit *UsdtDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsdtDeposit *UsdtDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UsdtDeposit *UsdtDepositCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UsdtDeposit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UsdtDeposit *UsdtDepositSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UsdtDeposit.Contract.BalanceOf(&_UsdtDeposit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UsdtDeposit *UsdtDepositCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UsdtDeposit.Contract.BalanceOf(&_UsdtDeposit.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtDeposit *UsdtDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtDeposit *UsdtDepositSession) Owner() (common.Address, error) {
	return _UsdtDeposit.Contract.Owner(&_UsdtDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UsdtDeposit *UsdtDepositCallerSession) Owner() (common.Address, error) {
	return _UsdtDeposit.Contract.Owner(&_UsdtDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtDeposit *UsdtDepositCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UsdtDeposit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtDeposit *UsdtDepositSession) Paused() (bool, error) {
	return _UsdtDeposit.Contract.Paused(&_UsdtDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UsdtDeposit *UsdtDepositCallerSession) Paused() (bool, error) {
	return _UsdtDeposit.Contract.Paused(&_UsdtDeposit.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_UsdtDeposit *UsdtDepositCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtDeposit.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_UsdtDeposit *UsdtDepositSession) Receiver() (common.Address, error) {
	return _UsdtDeposit.Contract.Receiver(&_UsdtDeposit.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_UsdtDeposit *UsdtDepositCallerSession) Receiver() (common.Address, error) {
	return _UsdtDeposit.Contract.Receiver(&_UsdtDeposit.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UsdtDeposit *UsdtDepositCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsdtDeposit.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UsdtDeposit *UsdtDepositSession) TokenAddress() (common.Address, error) {
	return _UsdtDeposit.Contract.TokenAddress(&_UsdtDeposit.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UsdtDeposit *UsdtDepositCallerSession) TokenAddress() (common.Address, error) {
	return _UsdtDeposit.Contract.TokenAddress(&_UsdtDeposit.CallOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_UsdtDeposit *UsdtDepositTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_UsdtDeposit *UsdtDepositSession) Attack() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Attack(&_UsdtDeposit.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) Attack() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Attack(&_UsdtDeposit.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_UsdtDeposit *UsdtDepositTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_UsdtDeposit *UsdtDepositSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Deposit(&_UsdtDeposit.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Deposit(&_UsdtDeposit.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtDeposit *UsdtDepositTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtDeposit *UsdtDepositSession) Pause() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Pause(&_UsdtDeposit.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) Pause() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Pause(&_UsdtDeposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UsdtDeposit *UsdtDepositTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UsdtDeposit *UsdtDepositSession) RenounceOwnership() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.RenounceOwnership(&_UsdtDeposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.RenounceOwnership(&_UsdtDeposit.TransactOpts)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiveAddr) returns()
func (_UsdtDeposit *UsdtDepositTransactor) SetReceiver(opts *bind.TransactOpts, _receiveAddr common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "setReceiver", _receiveAddr)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiveAddr) returns()
func (_UsdtDeposit *UsdtDepositSession) SetReceiver(_receiveAddr common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.SetReceiver(&_UsdtDeposit.TransactOpts, _receiveAddr)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiveAddr) returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) SetReceiver(_receiveAddr common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.SetReceiver(&_UsdtDeposit.TransactOpts, _receiveAddr)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_UsdtDeposit *UsdtDepositTransactor) SetToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "setToken", _tokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_UsdtDeposit *UsdtDepositSession) SetToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.SetToken(&_UsdtDeposit.TransactOpts, _tokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) SetToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.SetToken(&_UsdtDeposit.TransactOpts, _tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtDeposit *UsdtDepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtDeposit *UsdtDepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.TransferOwnership(&_UsdtDeposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UsdtDeposit.Contract.TransferOwnership(&_UsdtDeposit.TransactOpts, newOwner)
}

// Unpasue is a paid mutator transaction binding the contract method 0xcd5f76a0.
//
// Solidity: function unpasue() returns()
func (_UsdtDeposit *UsdtDepositTransactor) Unpasue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsdtDeposit.contract.Transact(opts, "unpasue")
}

// Unpasue is a paid mutator transaction binding the contract method 0xcd5f76a0.
//
// Solidity: function unpasue() returns()
func (_UsdtDeposit *UsdtDepositSession) Unpasue() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Unpasue(&_UsdtDeposit.TransactOpts)
}

// Unpasue is a paid mutator transaction binding the contract method 0xcd5f76a0.
//
// Solidity: function unpasue() returns()
func (_UsdtDeposit *UsdtDepositTransactorSession) Unpasue() (*types.Transaction, error) {
	return _UsdtDeposit.Contract.Unpasue(&_UsdtDeposit.TransactOpts)
}

// UsdtDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the UsdtDeposit contract.
type UsdtDepositDepositedIterator struct {
	Event *UsdtDepositDeposited // Event containing the contract specifics and raw log

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
func (it *UsdtDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtDepositDeposited)
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
		it.Event = new(UsdtDepositDeposited)
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
func (it *UsdtDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtDepositDeposited represents a Deposited event raised by the UsdtDeposit contract.
type UsdtDepositDeposited struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_UsdtDeposit *UsdtDepositFilterer) FilterDeposited(opts *bind.FilterOpts) (*UsdtDepositDepositedIterator, error) {

	logs, sub, err := _UsdtDeposit.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &UsdtDepositDepositedIterator{contract: _UsdtDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_UsdtDeposit *UsdtDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *UsdtDepositDeposited) (event.Subscription, error) {

	logs, sub, err := _UsdtDeposit.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtDepositDeposited)
				if err := _UsdtDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_UsdtDeposit *UsdtDepositFilterer) ParseDeposited(log types.Log) (*UsdtDepositDeposited, error) {
	event := new(UsdtDepositDeposited)
	if err := _UsdtDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtDepositOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UsdtDeposit contract.
type UsdtDepositOwnershipTransferredIterator struct {
	Event *UsdtDepositOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UsdtDepositOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtDepositOwnershipTransferred)
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
		it.Event = new(UsdtDepositOwnershipTransferred)
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
func (it *UsdtDepositOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtDepositOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtDepositOwnershipTransferred represents a OwnershipTransferred event raised by the UsdtDeposit contract.
type UsdtDepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UsdtDeposit *UsdtDepositFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsdtDepositOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UsdtDeposit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsdtDepositOwnershipTransferredIterator{contract: _UsdtDeposit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UsdtDeposit *UsdtDepositFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsdtDepositOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UsdtDeposit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtDepositOwnershipTransferred)
				if err := _UsdtDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UsdtDeposit *UsdtDepositFilterer) ParseOwnershipTransferred(log types.Log) (*UsdtDepositOwnershipTransferred, error) {
	event := new(UsdtDepositOwnershipTransferred)
	if err := _UsdtDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtDepositPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UsdtDeposit contract.
type UsdtDepositPausedIterator struct {
	Event *UsdtDepositPaused // Event containing the contract specifics and raw log

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
func (it *UsdtDepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtDepositPaused)
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
		it.Event = new(UsdtDepositPaused)
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
func (it *UsdtDepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtDepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtDepositPaused represents a Paused event raised by the UsdtDeposit contract.
type UsdtDepositPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) FilterPaused(opts *bind.FilterOpts) (*UsdtDepositPausedIterator, error) {

	logs, sub, err := _UsdtDeposit.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UsdtDepositPausedIterator{contract: _UsdtDeposit.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UsdtDepositPaused) (event.Subscription, error) {

	logs, sub, err := _UsdtDeposit.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtDepositPaused)
				if err := _UsdtDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) ParsePaused(log types.Log) (*UsdtDepositPaused, error) {
	event := new(UsdtDepositPaused)
	if err := _UsdtDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtDepositUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UsdtDeposit contract.
type UsdtDepositUnpausedIterator struct {
	Event *UsdtDepositUnpaused // Event containing the contract specifics and raw log

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
func (it *UsdtDepositUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtDepositUnpaused)
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
		it.Event = new(UsdtDepositUnpaused)
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
func (it *UsdtDepositUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtDepositUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtDepositUnpaused represents a Unpaused event raised by the UsdtDeposit contract.
type UsdtDepositUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UsdtDepositUnpausedIterator, error) {

	logs, sub, err := _UsdtDeposit.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UsdtDepositUnpausedIterator{contract: _UsdtDeposit.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UsdtDepositUnpaused) (event.Subscription, error) {

	logs, sub, err := _UsdtDeposit.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtDepositUnpaused)
				if err := _UsdtDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UsdtDeposit *UsdtDepositFilterer) ParseUnpaused(log types.Log) (*UsdtDepositUnpaused, error) {
	event := new(UsdtDepositUnpaused)
	if err := _UsdtDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
