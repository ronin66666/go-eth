// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subscribe

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

// SubscribeMetaData contains all meta data concerning the Subscribe contract.
var SubscribeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastLimitAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_oldSubScribe\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSubscribed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLimitAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldSubscribe\",\"outputs\":[{\"internalType\":\"contractIOldSubscribe\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldSubscribe\",\"type\":\"address\"}],\"name\":\"setOldSubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userSubscribeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SubscribeABI is the input ABI used to generate the binding from.
// Deprecated: Use SubscribeMetaData.ABI instead.
var SubscribeABI = SubscribeMetaData.ABI

// Subscribe is an auto generated Go binding around an Ethereum contract.
type Subscribe struct {
	SubscribeCaller     // Read-only binding to the contract
	SubscribeTransactor // Write-only binding to the contract
	SubscribeFilterer   // Log filterer for contract events
}

// SubscribeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscribeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscribeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscribeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscribeSession struct {
	Contract     *Subscribe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubscribeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscribeCallerSession struct {
	Contract *SubscribeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SubscribeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscribeTransactorSession struct {
	Contract     *SubscribeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SubscribeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscribeRaw struct {
	Contract *Subscribe // Generic contract binding to access the raw methods on
}

// SubscribeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscribeCallerRaw struct {
	Contract *SubscribeCaller // Generic read-only contract binding to access the raw methods on
}

// SubscribeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscribeTransactorRaw struct {
	Contract *SubscribeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscribe creates a new instance of Subscribe, bound to a specific deployed contract.
func NewSubscribe(address common.Address, backend bind.ContractBackend) (*Subscribe, error) {
	contract, err := bindSubscribe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subscribe{SubscribeCaller: SubscribeCaller{contract: contract}, SubscribeTransactor: SubscribeTransactor{contract: contract}, SubscribeFilterer: SubscribeFilterer{contract: contract}}, nil
}

// NewSubscribeCaller creates a new read-only instance of Subscribe, bound to a specific deployed contract.
func NewSubscribeCaller(address common.Address, caller bind.ContractCaller) (*SubscribeCaller, error) {
	contract, err := bindSubscribe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscribeCaller{contract: contract}, nil
}

// NewSubscribeTransactor creates a new write-only instance of Subscribe, bound to a specific deployed contract.
func NewSubscribeTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscribeTransactor, error) {
	contract, err := bindSubscribe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscribeTransactor{contract: contract}, nil
}

// NewSubscribeFilterer creates a new log filterer instance of Subscribe, bound to a specific deployed contract.
func NewSubscribeFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscribeFilterer, error) {
	contract, err := bindSubscribe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscribeFilterer{contract: contract}, nil
}

// bindSubscribe binds a generic wrapper to an already deployed contract.
func bindSubscribe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscribeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscribe *SubscribeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscribe.Contract.SubscribeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscribe *SubscribeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscribe.Contract.SubscribeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscribe *SubscribeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscribe.Contract.SubscribeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscribe *SubscribeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscribe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscribe *SubscribeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscribe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscribe *SubscribeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscribe.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeSession) ADMINROLE() ([32]byte, error) {
	return _Subscribe.Contract.ADMINROLE(&_Subscribe.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCallerSession) ADMINROLE() ([32]byte, error) {
	return _Subscribe.Contract.ADMINROLE(&_Subscribe.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Subscribe.Contract.DEFAULTADMINROLE(&_Subscribe.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Subscribe.Contract.DEFAULTADMINROLE(&_Subscribe.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeSession) UPGRADERROLE() ([32]byte, error) {
	return _Subscribe.Contract.UPGRADERROLE(&_Subscribe.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Subscribe *SubscribeCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Subscribe.Contract.UPGRADERROLE(&_Subscribe.CallOpts)
}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _lastLimitAmount, uint256 _lastTotal)
func (_Subscribe *SubscribeCaller) ConfigInfo(opts *bind.CallOpts) (struct {
	StartTime       *big.Int
	LastLimitAmount *big.Int
	LastTotal       *big.Int
}, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "configInfo")

	outstruct := new(struct {
		StartTime       *big.Int
		LastLimitAmount *big.Int
		LastTotal       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLimitAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _lastLimitAmount, uint256 _lastTotal)
func (_Subscribe *SubscribeSession) ConfigInfo() (struct {
	StartTime       *big.Int
	LastLimitAmount *big.Int
	LastTotal       *big.Int
}, error) {
	return _Subscribe.Contract.ConfigInfo(&_Subscribe.CallOpts)
}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _lastLimitAmount, uint256 _lastTotal)
func (_Subscribe *SubscribeCallerSession) ConfigInfo() (struct {
	StartTime       *big.Int
	LastLimitAmount *big.Int
	LastTotal       *big.Int
}, error) {
	return _Subscribe.Contract.ConfigInfo(&_Subscribe.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subscribe *SubscribeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subscribe *SubscribeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Subscribe.Contract.GetRoleAdmin(&_Subscribe.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subscribe *SubscribeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Subscribe.Contract.GetRoleAdmin(&_Subscribe.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subscribe *SubscribeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subscribe *SubscribeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Subscribe.Contract.HasRole(&_Subscribe.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subscribe *SubscribeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Subscribe.Contract.HasRole(&_Subscribe.CallOpts, role, account)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_Subscribe *SubscribeCaller) IsSubscribed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "isSubscribed", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_Subscribe *SubscribeSession) IsSubscribed(account common.Address) (bool, error) {
	return _Subscribe.Contract.IsSubscribed(&_Subscribe.CallOpts, account)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_Subscribe *SubscribeCallerSession) IsSubscribed(account common.Address) (bool, error) {
	return _Subscribe.Contract.IsSubscribed(&_Subscribe.CallOpts, account)
}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_Subscribe *SubscribeCaller) LastLimitAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "lastLimitAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_Subscribe *SubscribeSession) LastLimitAmount() (*big.Int, error) {
	return _Subscribe.Contract.LastLimitAmount(&_Subscribe.CallOpts)
}

// LastLimitAmount is a free data retrieval call binding the contract method 0xf611c52f.
//
// Solidity: function lastLimitAmount() view returns(uint256)
func (_Subscribe *SubscribeCallerSession) LastLimitAmount() (*big.Int, error) {
	return _Subscribe.Contract.LastLimitAmount(&_Subscribe.CallOpts)
}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_Subscribe *SubscribeCaller) LastTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "lastTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_Subscribe *SubscribeSession) LastTotal() (*big.Int, error) {
	return _Subscribe.Contract.LastTotal(&_Subscribe.CallOpts)
}

// LastTotal is a free data retrieval call binding the contract method 0x4812418d.
//
// Solidity: function lastTotal() view returns(uint256)
func (_Subscribe *SubscribeCallerSession) LastTotal() (*big.Int, error) {
	return _Subscribe.Contract.LastTotal(&_Subscribe.CallOpts)
}

// OldSubscribe is a free data retrieval call binding the contract method 0xea3577aa.
//
// Solidity: function oldSubscribe() view returns(address)
func (_Subscribe *SubscribeCaller) OldSubscribe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "oldSubscribe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldSubscribe is a free data retrieval call binding the contract method 0xea3577aa.
//
// Solidity: function oldSubscribe() view returns(address)
func (_Subscribe *SubscribeSession) OldSubscribe() (common.Address, error) {
	return _Subscribe.Contract.OldSubscribe(&_Subscribe.CallOpts)
}

// OldSubscribe is a free data retrieval call binding the contract method 0xea3577aa.
//
// Solidity: function oldSubscribe() view returns(address)
func (_Subscribe *SubscribeCallerSession) OldSubscribe() (common.Address, error) {
	return _Subscribe.Contract.OldSubscribe(&_Subscribe.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subscribe *SubscribeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subscribe *SubscribeSession) Paused() (bool, error) {
	return _Subscribe.Contract.Paused(&_Subscribe.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subscribe *SubscribeCallerSession) Paused() (bool, error) {
	return _Subscribe.Contract.Paused(&_Subscribe.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Subscribe *SubscribeCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Subscribe *SubscribeSession) Price() (*big.Int, error) {
	return _Subscribe.Contract.Price(&_Subscribe.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Subscribe *SubscribeCallerSession) Price() (*big.Int, error) {
	return _Subscribe.Contract.Price(&_Subscribe.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Subscribe *SubscribeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Subscribe *SubscribeSession) ProxiableUUID() ([32]byte, error) {
	return _Subscribe.Contract.ProxiableUUID(&_Subscribe.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Subscribe *SubscribeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Subscribe.Contract.ProxiableUUID(&_Subscribe.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Subscribe *SubscribeCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Subscribe *SubscribeSession) Receiver() (common.Address, error) {
	return _Subscribe.Contract.Receiver(&_Subscribe.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Subscribe *SubscribeCallerSession) Receiver() (common.Address, error) {
	return _Subscribe.Contract.Receiver(&_Subscribe.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Subscribe *SubscribeCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Subscribe *SubscribeSession) StartTime() (*big.Int, error) {
	return _Subscribe.Contract.StartTime(&_Subscribe.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Subscribe *SubscribeCallerSession) StartTime() (*big.Int, error) {
	return _Subscribe.Contract.StartTime(&_Subscribe.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subscribe *SubscribeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subscribe *SubscribeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subscribe.Contract.SupportsInterface(&_Subscribe.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subscribe *SubscribeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subscribe.Contract.SupportsInterface(&_Subscribe.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Subscribe *SubscribeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Subscribe *SubscribeSession) Token() (common.Address, error) {
	return _Subscribe.Contract.Token(&_Subscribe.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Subscribe *SubscribeCallerSession) Token() (common.Address, error) {
	return _Subscribe.Contract.Token(&_Subscribe.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_Subscribe *SubscribeCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_Subscribe *SubscribeSession) Total() (*big.Int, error) {
	return _Subscribe.Contract.Total(&_Subscribe.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_Subscribe *SubscribeCallerSession) Total() (*big.Int, error) {
	return _Subscribe.Contract.Total(&_Subscribe.CallOpts)
}

// UserSubscribeAmount is a free data retrieval call binding the contract method 0x8ebb1895.
//
// Solidity: function userSubscribeAmount(address account) view returns(uint256 amount)
func (_Subscribe *SubscribeCaller) UserSubscribeAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Subscribe.contract.Call(opts, &out, "userSubscribeAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSubscribeAmount is a free data retrieval call binding the contract method 0x8ebb1895.
//
// Solidity: function userSubscribeAmount(address account) view returns(uint256 amount)
func (_Subscribe *SubscribeSession) UserSubscribeAmount(account common.Address) (*big.Int, error) {
	return _Subscribe.Contract.UserSubscribeAmount(&_Subscribe.CallOpts, account)
}

// UserSubscribeAmount is a free data retrieval call binding the contract method 0x8ebb1895.
//
// Solidity: function userSubscribeAmount(address account) view returns(uint256 amount)
func (_Subscribe *SubscribeCallerSession) UserSubscribeAmount(account common.Address) (*big.Int, error) {
	return _Subscribe.Contract.UserSubscribeAmount(&_Subscribe.CallOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.GrantRole(&_Subscribe.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.GrantRole(&_Subscribe.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _token, address _receiver, uint256 _price, address _oldSubScribe) returns()
func (_Subscribe *SubscribeTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _price *big.Int, _oldSubScribe common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "initialize", _token, _receiver, _price, _oldSubScribe)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _token, address _receiver, uint256 _price, address _oldSubScribe) returns()
func (_Subscribe *SubscribeSession) Initialize(_token common.Address, _receiver common.Address, _price *big.Int, _oldSubScribe common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.Initialize(&_Subscribe.TransactOpts, _token, _receiver, _price, _oldSubScribe)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _token, address _receiver, uint256 _price, address _oldSubScribe) returns()
func (_Subscribe *SubscribeTransactorSession) Initialize(_token common.Address, _receiver common.Address, _price *big.Int, _oldSubScribe common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.Initialize(&_Subscribe.TransactOpts, _token, _receiver, _price, _oldSubScribe)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subscribe *SubscribeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subscribe *SubscribeSession) Pause() (*types.Transaction, error) {
	return _Subscribe.Contract.Pause(&_Subscribe.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subscribe *SubscribeTransactorSession) Pause() (*types.Transaction, error) {
	return _Subscribe.Contract.Pause(&_Subscribe.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.RenounceRole(&_Subscribe.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.RenounceRole(&_Subscribe.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.RevokeRole(&_Subscribe.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subscribe *SubscribeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.RevokeRole(&_Subscribe.TransactOpts, role, account)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_Subscribe *SubscribeTransactor) SetLimitAmount(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setLimitAmount", _limit)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_Subscribe *SubscribeSession) SetLimitAmount(_limit *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetLimitAmount(&_Subscribe.TransactOpts, _limit)
}

// SetLimitAmount is a paid mutator transaction binding the contract method 0x1f107a45.
//
// Solidity: function setLimitAmount(uint256 _limit) returns()
func (_Subscribe *SubscribeTransactorSession) SetLimitAmount(_limit *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetLimitAmount(&_Subscribe.TransactOpts, _limit)
}

// SetOldSubscribe is a paid mutator transaction binding the contract method 0xaa696ba2.
//
// Solidity: function setOldSubscribe(address _oldSubscribe) returns()
func (_Subscribe *SubscribeTransactor) SetOldSubscribe(opts *bind.TransactOpts, _oldSubscribe common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setOldSubscribe", _oldSubscribe)
}

// SetOldSubscribe is a paid mutator transaction binding the contract method 0xaa696ba2.
//
// Solidity: function setOldSubscribe(address _oldSubscribe) returns()
func (_Subscribe *SubscribeSession) SetOldSubscribe(_oldSubscribe common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetOldSubscribe(&_Subscribe.TransactOpts, _oldSubscribe)
}

// SetOldSubscribe is a paid mutator transaction binding the contract method 0xaa696ba2.
//
// Solidity: function setOldSubscribe(address _oldSubscribe) returns()
func (_Subscribe *SubscribeTransactorSession) SetOldSubscribe(_oldSubscribe common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetOldSubscribe(&_Subscribe.TransactOpts, _oldSubscribe)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Subscribe *SubscribeTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Subscribe *SubscribeSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetPrice(&_Subscribe.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Subscribe *SubscribeTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetPrice(&_Subscribe.TransactOpts, _price)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_Subscribe *SubscribeTransactor) SetReceiver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setReceiver", _account)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_Subscribe *SubscribeSession) SetReceiver(_account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetReceiver(&_Subscribe.TransactOpts, _account)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _account) returns()
func (_Subscribe *SubscribeTransactorSession) SetReceiver(_account common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetReceiver(&_Subscribe.TransactOpts, _account)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _startTime) returns()
func (_Subscribe *SubscribeTransactor) SetStartTime(opts *bind.TransactOpts, _startTime *big.Int) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setStartTime", _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _startTime) returns()
func (_Subscribe *SubscribeSession) SetStartTime(_startTime *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetStartTime(&_Subscribe.TransactOpts, _startTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _startTime) returns()
func (_Subscribe *SubscribeTransactorSession) SetStartTime(_startTime *big.Int) (*types.Transaction, error) {
	return _Subscribe.Contract.SetStartTime(&_Subscribe.TransactOpts, _startTime)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Subscribe *SubscribeTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Subscribe *SubscribeSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetToken(&_Subscribe.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Subscribe *SubscribeTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.SetToken(&_Subscribe.TransactOpts, _token)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Subscribe *SubscribeTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Subscribe *SubscribeSession) Subscribe() (*types.Transaction, error) {
	return _Subscribe.Contract.Subscribe(&_Subscribe.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Subscribe *SubscribeTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Subscribe.Contract.Subscribe(&_Subscribe.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subscribe *SubscribeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subscribe *SubscribeSession) Unpause() (*types.Transaction, error) {
	return _Subscribe.Contract.Unpause(&_Subscribe.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subscribe *SubscribeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Subscribe.Contract.Unpause(&_Subscribe.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Subscribe *SubscribeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Subscribe *SubscribeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.UpgradeTo(&_Subscribe.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Subscribe *SubscribeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Subscribe.Contract.UpgradeTo(&_Subscribe.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Subscribe *SubscribeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Subscribe.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Subscribe *SubscribeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Subscribe.Contract.UpgradeToAndCall(&_Subscribe.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Subscribe *SubscribeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Subscribe.Contract.UpgradeToAndCall(&_Subscribe.TransactOpts, newImplementation, data)
}

// SubscribeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Subscribe contract.
type SubscribeAdminChangedIterator struct {
	Event *SubscribeAdminChanged // Event containing the contract specifics and raw log

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
func (it *SubscribeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeAdminChanged)
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
		it.Event = new(SubscribeAdminChanged)
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
func (it *SubscribeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeAdminChanged represents a AdminChanged event raised by the Subscribe contract.
type SubscribeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Subscribe *SubscribeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SubscribeAdminChangedIterator, error) {

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SubscribeAdminChangedIterator{contract: _Subscribe.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Subscribe *SubscribeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SubscribeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeAdminChanged)
				if err := _Subscribe.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Subscribe *SubscribeFilterer) ParseAdminChanged(log types.Log) (*SubscribeAdminChanged, error) {
	event := new(SubscribeAdminChanged)
	if err := _Subscribe.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Subscribe contract.
type SubscribeBeaconUpgradedIterator struct {
	Event *SubscribeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SubscribeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeBeaconUpgraded)
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
		it.Event = new(SubscribeBeaconUpgraded)
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
func (it *SubscribeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeBeaconUpgraded represents a BeaconUpgraded event raised by the Subscribe contract.
type SubscribeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Subscribe *SubscribeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SubscribeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeBeaconUpgradedIterator{contract: _Subscribe.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Subscribe *SubscribeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SubscribeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeBeaconUpgraded)
				if err := _Subscribe.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Subscribe *SubscribeFilterer) ParseBeaconUpgraded(log types.Log) (*SubscribeBeaconUpgraded, error) {
	event := new(SubscribeBeaconUpgraded)
	if err := _Subscribe.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Subscribe contract.
type SubscribeDepositedIterator struct {
	Event *SubscribeDeposited // Event containing the contract specifics and raw log

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
func (it *SubscribeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeDeposited)
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
		it.Event = new(SubscribeDeposited)
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
func (it *SubscribeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeDeposited represents a Deposited event raised by the Subscribe contract.
type SubscribeDeposited struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_Subscribe *SubscribeFilterer) FilterDeposited(opts *bind.FilterOpts) (*SubscribeDepositedIterator, error) {

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &SubscribeDepositedIterator{contract: _Subscribe.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address from, address to, uint256 amount)
func (_Subscribe *SubscribeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *SubscribeDeposited) (event.Subscription, error) {

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeDeposited)
				if err := _Subscribe.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Subscribe *SubscribeFilterer) ParseDeposited(log types.Log) (*SubscribeDeposited, error) {
	event := new(SubscribeDeposited)
	if err := _Subscribe.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Subscribe contract.
type SubscribeInitializedIterator struct {
	Event *SubscribeInitialized // Event containing the contract specifics and raw log

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
func (it *SubscribeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeInitialized)
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
		it.Event = new(SubscribeInitialized)
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
func (it *SubscribeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeInitialized represents a Initialized event raised by the Subscribe contract.
type SubscribeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Subscribe *SubscribeFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubscribeInitializedIterator, error) {

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubscribeInitializedIterator{contract: _Subscribe.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Subscribe *SubscribeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubscribeInitialized) (event.Subscription, error) {

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeInitialized)
				if err := _Subscribe.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Subscribe *SubscribeFilterer) ParseInitialized(log types.Log) (*SubscribeInitialized, error) {
	event := new(SubscribeInitialized)
	if err := _Subscribe.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Subscribe contract.
type SubscribePausedIterator struct {
	Event *SubscribePaused // Event containing the contract specifics and raw log

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
func (it *SubscribePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribePaused)
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
		it.Event = new(SubscribePaused)
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
func (it *SubscribePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribePaused represents a Paused event raised by the Subscribe contract.
type SubscribePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Subscribe *SubscribeFilterer) FilterPaused(opts *bind.FilterOpts) (*SubscribePausedIterator, error) {

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SubscribePausedIterator{contract: _Subscribe.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Subscribe *SubscribeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SubscribePaused) (event.Subscription, error) {

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribePaused)
				if err := _Subscribe.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Subscribe *SubscribeFilterer) ParsePaused(log types.Log) (*SubscribePaused, error) {
	event := new(SubscribePaused)
	if err := _Subscribe.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Subscribe contract.
type SubscribeRoleAdminChangedIterator struct {
	Event *SubscribeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SubscribeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeRoleAdminChanged)
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
		it.Event = new(SubscribeRoleAdminChanged)
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
func (it *SubscribeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeRoleAdminChanged represents a RoleAdminChanged event raised by the Subscribe contract.
type SubscribeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Subscribe *SubscribeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SubscribeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeRoleAdminChangedIterator{contract: _Subscribe.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Subscribe *SubscribeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SubscribeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeRoleAdminChanged)
				if err := _Subscribe.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Subscribe *SubscribeFilterer) ParseRoleAdminChanged(log types.Log) (*SubscribeRoleAdminChanged, error) {
	event := new(SubscribeRoleAdminChanged)
	if err := _Subscribe.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Subscribe contract.
type SubscribeRoleGrantedIterator struct {
	Event *SubscribeRoleGranted // Event containing the contract specifics and raw log

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
func (it *SubscribeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeRoleGranted)
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
		it.Event = new(SubscribeRoleGranted)
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
func (it *SubscribeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeRoleGranted represents a RoleGranted event raised by the Subscribe contract.
type SubscribeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubscribeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeRoleGrantedIterator{contract: _Subscribe.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SubscribeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeRoleGranted)
				if err := _Subscribe.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) ParseRoleGranted(log types.Log) (*SubscribeRoleGranted, error) {
	event := new(SubscribeRoleGranted)
	if err := _Subscribe.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Subscribe contract.
type SubscribeRoleRevokedIterator struct {
	Event *SubscribeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SubscribeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeRoleRevoked)
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
		it.Event = new(SubscribeRoleRevoked)
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
func (it *SubscribeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeRoleRevoked represents a RoleRevoked event raised by the Subscribe contract.
type SubscribeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubscribeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeRoleRevokedIterator{contract: _Subscribe.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SubscribeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeRoleRevoked)
				if err := _Subscribe.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subscribe *SubscribeFilterer) ParseRoleRevoked(log types.Log) (*SubscribeRoleRevoked, error) {
	event := new(SubscribeRoleRevoked)
	if err := _Subscribe.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Subscribe contract.
type SubscribeUnpausedIterator struct {
	Event *SubscribeUnpaused // Event containing the contract specifics and raw log

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
func (it *SubscribeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeUnpaused)
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
		it.Event = new(SubscribeUnpaused)
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
func (it *SubscribeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeUnpaused represents a Unpaused event raised by the Subscribe contract.
type SubscribeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Subscribe *SubscribeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SubscribeUnpausedIterator, error) {

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SubscribeUnpausedIterator{contract: _Subscribe.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Subscribe *SubscribeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SubscribeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeUnpaused)
				if err := _Subscribe.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Subscribe *SubscribeFilterer) ParseUnpaused(log types.Log) (*SubscribeUnpaused, error) {
	event := new(SubscribeUnpaused)
	if err := _Subscribe.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Subscribe contract.
type SubscribeUpgradedIterator struct {
	Event *SubscribeUpgraded // Event containing the contract specifics and raw log

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
func (it *SubscribeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeUpgraded)
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
		it.Event = new(SubscribeUpgraded)
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
func (it *SubscribeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeUpgraded represents a Upgraded event raised by the Subscribe contract.
type SubscribeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Subscribe *SubscribeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SubscribeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Subscribe.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeUpgradedIterator{contract: _Subscribe.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Subscribe *SubscribeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SubscribeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Subscribe.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeUpgraded)
				if err := _Subscribe.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Subscribe *SubscribeFilterer) ParseUpgraded(log types.Log) (*SubscribeUpgraded, error) {
	event := new(SubscribeUpgraded)
	if err := _Subscribe.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
