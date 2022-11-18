// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package indux_deposit

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

// InduxDepositMetaData contains all meta data concerning the InduxDeposit contract.
var InduxDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_limitAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approveFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApproveFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLimitAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchSetOldBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setOldBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InduxDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use InduxDepositMetaData.ABI instead.
var InduxDepositABI = InduxDepositMetaData.ABI

// InduxDeposit is an auto generated Go binding around an Ethereum contract.
type InduxDeposit struct {
	InduxDepositCaller     // Read-only binding to the contract
	InduxDepositTransactor // Write-only binding to the contract
	InduxDepositFilterer   // Log filterer for contract events
}

// InduxDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type InduxDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InduxDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InduxDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InduxDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InduxDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InduxDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InduxDepositSession struct {
	Contract     *InduxDeposit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InduxDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InduxDepositCallerSession struct {
	Contract *InduxDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InduxDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InduxDepositTransactorSession struct {
	Contract     *InduxDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InduxDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type InduxDepositRaw struct {
	Contract *InduxDeposit // Generic contract binding to access the raw methods on
}

// InduxDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InduxDepositCallerRaw struct {
	Contract *InduxDepositCaller // Generic read-only contract binding to access the raw methods on
}

// InduxDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InduxDepositTransactorRaw struct {
	Contract *InduxDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInduxDeposit creates a new instance of InduxDeposit, bound to a specific deployed contract.
func NewInduxDeposit(address common.Address, backend bind.ContractBackend) (*InduxDeposit, error) {
	contract, err := bindInduxDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InduxDeposit{InduxDepositCaller: InduxDepositCaller{contract: contract}, InduxDepositTransactor: InduxDepositTransactor{contract: contract}, InduxDepositFilterer: InduxDepositFilterer{contract: contract}}, nil
}

// NewInduxDepositCaller creates a new read-only instance of InduxDeposit, bound to a specific deployed contract.
func NewInduxDepositCaller(address common.Address, caller bind.ContractCaller) (*InduxDepositCaller, error) {
	contract, err := bindInduxDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InduxDepositCaller{contract: contract}, nil
}

// NewInduxDepositTransactor creates a new write-only instance of InduxDeposit, bound to a specific deployed contract.
func NewInduxDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*InduxDepositTransactor, error) {
	contract, err := bindInduxDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InduxDepositTransactor{contract: contract}, nil
}

// NewInduxDepositFilterer creates a new log filterer instance of InduxDeposit, bound to a specific deployed contract.
func NewInduxDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*InduxDepositFilterer, error) {
	contract, err := bindInduxDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InduxDepositFilterer{contract: contract}, nil
}

// bindInduxDeposit binds a generic wrapper to an already deployed contract.
func bindInduxDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InduxDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InduxDeposit *InduxDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InduxDeposit.Contract.InduxDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InduxDeposit *InduxDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.Contract.InduxDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InduxDeposit *InduxDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InduxDeposit.Contract.InduxDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InduxDeposit *InduxDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InduxDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InduxDeposit *InduxDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InduxDeposit *InduxDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InduxDeposit.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InduxDeposit *InduxDepositCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InduxDeposit *InduxDepositSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InduxDeposit.Contract.BalanceOf(&_InduxDeposit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InduxDeposit *InduxDepositCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InduxDeposit.Contract.BalanceOf(&_InduxDeposit.CallOpts, account)
}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_InduxDeposit *InduxDepositCaller) IsApproveFor(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "isApproveFor", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_InduxDeposit *InduxDepositSession) IsApproveFor(operator common.Address) (bool, error) {
	return _InduxDeposit.Contract.IsApproveFor(&_InduxDeposit.CallOpts, operator)
}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_InduxDeposit *InduxDepositCallerSession) IsApproveFor(operator common.Address) (bool, error) {
	return _InduxDeposit.Contract.IsApproveFor(&_InduxDeposit.CallOpts, operator)
}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_InduxDeposit *InduxDepositCaller) LastLimitAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "lastLimitAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_InduxDeposit *InduxDepositSession) LastLimitAmount() (*big.Int, error) {
	return _InduxDeposit.Contract.LastLimitAmount(&_InduxDeposit.CallOpts)
}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_InduxDeposit *InduxDepositCallerSession) LastLimitAmount() (*big.Int, error) {
	return _InduxDeposit.Contract.LastLimitAmount(&_InduxDeposit.CallOpts)
}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_InduxDeposit *InduxDepositCaller) LastTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "lastTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_InduxDeposit *InduxDepositSession) LastTotal() (*big.Int, error) {
	return _InduxDeposit.Contract.LastTotal(&_InduxDeposit.CallOpts)
}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_InduxDeposit *InduxDepositCallerSession) LastTotal() (*big.Int, error) {
	return _InduxDeposit.Contract.LastTotal(&_InduxDeposit.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_InduxDeposit *InduxDepositCaller) Members(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_InduxDeposit *InduxDepositSession) Members(arg0 common.Address) (bool, error) {
	return _InduxDeposit.Contract.Members(&_InduxDeposit.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_InduxDeposit *InduxDepositCallerSession) Members(arg0 common.Address) (bool, error) {
	return _InduxDeposit.Contract.Members(&_InduxDeposit.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InduxDeposit *InduxDepositCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InduxDeposit *InduxDepositSession) Name() (string, error) {
	return _InduxDeposit.Contract.Name(&_InduxDeposit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InduxDeposit *InduxDepositCallerSession) Name() (string, error) {
	return _InduxDeposit.Contract.Name(&_InduxDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InduxDeposit *InduxDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InduxDeposit *InduxDepositSession) Owner() (common.Address, error) {
	return _InduxDeposit.Contract.Owner(&_InduxDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InduxDeposit *InduxDepositCallerSession) Owner() (common.Address, error) {
	return _InduxDeposit.Contract.Owner(&_InduxDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InduxDeposit *InduxDepositCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InduxDeposit *InduxDepositSession) Paused() (bool, error) {
	return _InduxDeposit.Contract.Paused(&_InduxDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InduxDeposit *InduxDepositCallerSession) Paused() (bool, error) {
	return _InduxDeposit.Contract.Paused(&_InduxDeposit.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_InduxDeposit *InduxDepositCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_InduxDeposit *InduxDepositSession) Price() (*big.Int, error) {
	return _InduxDeposit.Contract.Price(&_InduxDeposit.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_InduxDeposit *InduxDepositCallerSession) Price() (*big.Int, error) {
	return _InduxDeposit.Contract.Price(&_InduxDeposit.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_InduxDeposit *InduxDepositCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_InduxDeposit *InduxDepositSession) Receiver() (common.Address, error) {
	return _InduxDeposit.Contract.Receiver(&_InduxDeposit.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_InduxDeposit *InduxDepositCallerSession) Receiver() (common.Address, error) {
	return _InduxDeposit.Contract.Receiver(&_InduxDeposit.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_InduxDeposit *InduxDepositCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_InduxDeposit *InduxDepositSession) TokenAddress() (common.Address, error) {
	return _InduxDeposit.Contract.TokenAddress(&_InduxDeposit.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_InduxDeposit *InduxDepositCallerSession) TokenAddress() (common.Address, error) {
	return _InduxDeposit.Contract.TokenAddress(&_InduxDeposit.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_InduxDeposit *InduxDepositCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_InduxDeposit *InduxDepositSession) Total() (*big.Int, error) {
	return _InduxDeposit.Contract.Total(&_InduxDeposit.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_InduxDeposit *InduxDepositCallerSession) Total() (*big.Int, error) {
	return _InduxDeposit.Contract.Total(&_InduxDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_InduxDeposit *InduxDepositCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InduxDeposit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_InduxDeposit *InduxDepositSession) Version() (string, error) {
	return _InduxDeposit.Contract.Version(&_InduxDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_InduxDeposit *InduxDepositCallerSession) Version() (string, error) {
	return _InduxDeposit.Contract.Version(&_InduxDeposit.CallOpts)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_InduxDeposit *InduxDepositTransactor) ApproveFor(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "approveFor", operator, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_InduxDeposit *InduxDepositSession) ApproveFor(operator common.Address, approved bool) (*types.Transaction, error) {
	return _InduxDeposit.Contract.ApproveFor(&_InduxDeposit.TransactOpts, operator, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_InduxDeposit *InduxDepositTransactorSession) ApproveFor(operator common.Address, approved bool) (*types.Transaction, error) {
	return _InduxDeposit.Contract.ApproveFor(&_InduxDeposit.TransactOpts, operator, approved)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_InduxDeposit *InduxDepositTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_InduxDeposit *InduxDepositSession) Attack() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Attack(&_InduxDeposit.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_InduxDeposit *InduxDepositTransactorSession) Attack() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Attack(&_InduxDeposit.TransactOpts)
}

// BatchSetOldBalance is a paid mutator transaction binding the contract method 0xe117baca.
//
// Solidity: function batchSetOldBalance(address[] accounts, uint256[] amounts) returns()
func (_InduxDeposit *InduxDepositTransactor) BatchSetOldBalance(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "batchSetOldBalance", accounts, amounts)
}

// BatchSetOldBalance is a paid mutator transaction binding the contract method 0xe117baca.
//
// Solidity: function batchSetOldBalance(address[] accounts, uint256[] amounts) returns()
func (_InduxDeposit *InduxDepositSession) BatchSetOldBalance(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.BatchSetOldBalance(&_InduxDeposit.TransactOpts, accounts, amounts)
}

// BatchSetOldBalance is a paid mutator transaction binding the contract method 0xe117baca.
//
// Solidity: function batchSetOldBalance(address[] accounts, uint256[] amounts) returns()
func (_InduxDeposit *InduxDepositTransactorSession) BatchSetOldBalance(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.BatchSetOldBalance(&_InduxDeposit.TransactOpts, accounts, amounts)
}

// Deposit is a paid mutator transaction binding the contract method 0x7d7dd076.
//
// Solidity: function deposit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_InduxDeposit *InduxDepositTransactor) Deposit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "deposit", owner, spender, value, deadline, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0x7d7dd076.
//
// Solidity: function deposit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_InduxDeposit *InduxDepositSession) Deposit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InduxDeposit.Contract.Deposit(&_InduxDeposit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0x7d7dd076.
//
// Solidity: function deposit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_InduxDeposit *InduxDepositTransactorSession) Deposit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InduxDeposit.Contract.Deposit(&_InduxDeposit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InduxDeposit *InduxDepositTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InduxDeposit *InduxDepositSession) Pause() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Pause(&_InduxDeposit.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InduxDeposit *InduxDepositTransactorSession) Pause() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Pause(&_InduxDeposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InduxDeposit *InduxDepositTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InduxDeposit *InduxDepositSession) RenounceOwnership() (*types.Transaction, error) {
	return _InduxDeposit.Contract.RenounceOwnership(&_InduxDeposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InduxDeposit *InduxDepositTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InduxDeposit.Contract.RenounceOwnership(&_InduxDeposit.TransactOpts)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_InduxDeposit *InduxDepositTransactor) SetLimitAmount(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "setLimitAmount", _limit)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_InduxDeposit *InduxDepositSession) SetLimitAmount(_limit *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetLimitAmount(&_InduxDeposit.TransactOpts, _limit)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_InduxDeposit *InduxDepositTransactorSession) SetLimitAmount(_limit *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetLimitAmount(&_InduxDeposit.TransactOpts, _limit)
}

// SetOldBalance is a paid mutator transaction binding the contract method 0x24e9f9a9.
//
// Solidity: function setOldBalance(address account, uint256 amount) returns()
func (_InduxDeposit *InduxDepositTransactor) SetOldBalance(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "setOldBalance", account, amount)
}

// SetOldBalance is a paid mutator transaction binding the contract method 0x24e9f9a9.
//
// Solidity: function setOldBalance(address account, uint256 amount) returns()
func (_InduxDeposit *InduxDepositSession) SetOldBalance(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetOldBalance(&_InduxDeposit.TransactOpts, account, amount)
}

// SetOldBalance is a paid mutator transaction binding the contract method 0x24e9f9a9.
//
// Solidity: function setOldBalance(address account, uint256 amount) returns()
func (_InduxDeposit *InduxDepositTransactorSession) SetOldBalance(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetOldBalance(&_InduxDeposit.TransactOpts, account, amount)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_InduxDeposit *InduxDepositTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_InduxDeposit *InduxDepositSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetPrice(&_InduxDeposit.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_InduxDeposit *InduxDepositTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetPrice(&_InduxDeposit.TransactOpts, _price)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_InduxDeposit *InduxDepositTransactor) SetReceiver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "setReceiver", _account)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_InduxDeposit *InduxDepositSession) SetReceiver(_account common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetReceiver(&_InduxDeposit.TransactOpts, _account)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_InduxDeposit *InduxDepositTransactorSession) SetReceiver(_account common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetReceiver(&_InduxDeposit.TransactOpts, _account)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_InduxDeposit *InduxDepositTransactor) SetToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "setToken", _tokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_InduxDeposit *InduxDepositSession) SetToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetToken(&_InduxDeposit.TransactOpts, _tokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _tokenAddress) returns()
func (_InduxDeposit *InduxDepositTransactorSession) SetToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.SetToken(&_InduxDeposit.TransactOpts, _tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InduxDeposit *InduxDepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InduxDeposit *InduxDepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.TransferOwnership(&_InduxDeposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InduxDeposit *InduxDepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InduxDeposit.Contract.TransferOwnership(&_InduxDeposit.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InduxDeposit *InduxDepositTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InduxDeposit.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InduxDeposit *InduxDepositSession) Unpause() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Unpause(&_InduxDeposit.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InduxDeposit *InduxDepositTransactorSession) Unpause() (*types.Transaction, error) {
	return _InduxDeposit.Contract.Unpause(&_InduxDeposit.TransactOpts)
}

// InduxDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the InduxDeposit contract.
type InduxDepositDepositedIterator struct {
	Event *InduxDepositDeposited // Event containing the contract specifics and raw log

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
func (it *InduxDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InduxDepositDeposited)
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
		it.Event = new(InduxDepositDeposited)
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
func (it *InduxDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InduxDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InduxDepositDeposited represents a Deposited event raised by the InduxDeposit contract.
type InduxDepositDeposited struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_InduxDeposit *InduxDepositFilterer) FilterDeposited(opts *bind.FilterOpts) (*InduxDepositDepositedIterator, error) {

	logs, sub, err := _InduxDeposit.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &InduxDepositDepositedIterator{contract: _InduxDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_InduxDeposit *InduxDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *InduxDepositDeposited) (event.Subscription, error) {

	logs, sub, err := _InduxDeposit.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InduxDepositDeposited)
				if err := _InduxDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_InduxDeposit *InduxDepositFilterer) ParseDeposited(log types.Log) (*InduxDepositDeposited, error) {
	event := new(InduxDepositDeposited)
	if err := _InduxDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InduxDepositOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InduxDeposit contract.
type InduxDepositOwnershipTransferredIterator struct {
	Event *InduxDepositOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InduxDepositOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InduxDepositOwnershipTransferred)
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
		it.Event = new(InduxDepositOwnershipTransferred)
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
func (it *InduxDepositOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InduxDepositOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InduxDepositOwnershipTransferred represents a OwnershipTransferred event raised by the InduxDeposit contract.
type InduxDepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InduxDeposit *InduxDepositFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InduxDepositOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InduxDeposit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InduxDepositOwnershipTransferredIterator{contract: _InduxDeposit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InduxDeposit *InduxDepositFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InduxDepositOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InduxDeposit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InduxDepositOwnershipTransferred)
				if err := _InduxDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InduxDeposit *InduxDepositFilterer) ParseOwnershipTransferred(log types.Log) (*InduxDepositOwnershipTransferred, error) {
	event := new(InduxDepositOwnershipTransferred)
	if err := _InduxDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InduxDepositPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the InduxDeposit contract.
type InduxDepositPausedIterator struct {
	Event *InduxDepositPaused // Event containing the contract specifics and raw log

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
func (it *InduxDepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InduxDepositPaused)
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
		it.Event = new(InduxDepositPaused)
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
func (it *InduxDepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InduxDepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InduxDepositPaused represents a Paused event raised by the InduxDeposit contract.
type InduxDepositPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InduxDeposit *InduxDepositFilterer) FilterPaused(opts *bind.FilterOpts) (*InduxDepositPausedIterator, error) {

	logs, sub, err := _InduxDeposit.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InduxDepositPausedIterator{contract: _InduxDeposit.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InduxDeposit *InduxDepositFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InduxDepositPaused) (event.Subscription, error) {

	logs, sub, err := _InduxDeposit.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InduxDepositPaused)
				if err := _InduxDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_InduxDeposit *InduxDepositFilterer) ParsePaused(log types.Log) (*InduxDepositPaused, error) {
	event := new(InduxDepositPaused)
	if err := _InduxDeposit.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InduxDepositUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the InduxDeposit contract.
type InduxDepositUnpausedIterator struct {
	Event *InduxDepositUnpaused // Event containing the contract specifics and raw log

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
func (it *InduxDepositUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InduxDepositUnpaused)
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
		it.Event = new(InduxDepositUnpaused)
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
func (it *InduxDepositUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InduxDepositUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InduxDepositUnpaused represents a Unpaused event raised by the InduxDeposit contract.
type InduxDepositUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InduxDeposit *InduxDepositFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InduxDepositUnpausedIterator, error) {

	logs, sub, err := _InduxDeposit.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InduxDepositUnpausedIterator{contract: _InduxDeposit.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InduxDeposit *InduxDepositFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InduxDepositUnpaused) (event.Subscription, error) {

	logs, sub, err := _InduxDeposit.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InduxDepositUnpaused)
				if err := _InduxDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_InduxDeposit *InduxDepositFilterer) ParseUnpaused(log types.Log) (*InduxDepositUnpaused, error) {
	event := new(InduxDepositUnpaused)
	if err := _InduxDeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
