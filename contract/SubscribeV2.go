// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SubscribeV2ReserveInfo is an auto generated low-level Go binding around an user-defined struct.
type SubscribeV2ReserveInfo struct {
	SubscribeId *big.Int
	User        common.Address
	Amount      *big.Int
	Timestamp   *big.Int
}

// SubscribeV2MetaData contains all meta data concerning the SubscribeV2 contract.
var SubscribeV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Reserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMPLOY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limitReserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserveCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserveId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"currentReserveInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subscribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSubscribeV2.ReserveInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserveId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isReserve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSubscribed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitReserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserveId\",\"type\":\"uint256\"}],\"name\":\"reserveCountAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"subscribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"reserveInfoAt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subscribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSubscribeV2.ReserveInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"reserveable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limitReserveAmount\",\"type\":\"uint256\"}],\"name\":\"setLimitReserveAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddr\",\"type\":\"address\"}],\"name\":\"setNFTAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"setReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limitReserveAmount\",\"type\":\"uint256\"}],\"name\":\"setTimeAndLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userReserveInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subscribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSubscribeV2.ReserveInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SubscribeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SubscribeV2MetaData.ABI instead.
var SubscribeV2ABI = SubscribeV2MetaData.ABI

// SubscribeV2 is an auto generated Go binding around an Ethereum contract.
type SubscribeV2 struct {
	SubscribeV2Caller     // Read-only binding to the contract
	SubscribeV2Transactor // Write-only binding to the contract
	SubscribeV2Filterer   // Log filterer for contract events
}

// SubscribeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SubscribeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscribeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscribeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscribeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscribeV2Session struct {
	Contract     *SubscribeV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubscribeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscribeV2CallerSession struct {
	Contract *SubscribeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SubscribeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscribeV2TransactorSession struct {
	Contract     *SubscribeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SubscribeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SubscribeV2Raw struct {
	Contract *SubscribeV2 // Generic contract binding to access the raw methods on
}

// SubscribeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscribeV2CallerRaw struct {
	Contract *SubscribeV2Caller // Generic read-only contract binding to access the raw methods on
}

// SubscribeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscribeV2TransactorRaw struct {
	Contract *SubscribeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscribeV2 creates a new instance of SubscribeV2, bound to a specific deployed contract.
func NewSubscribeV2(address common.Address, backend bind.ContractBackend) (*SubscribeV2, error) {
	contract, err := bindSubscribeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2{SubscribeV2Caller: SubscribeV2Caller{contract: contract}, SubscribeV2Transactor: SubscribeV2Transactor{contract: contract}, SubscribeV2Filterer: SubscribeV2Filterer{contract: contract}}, nil
}

// NewSubscribeV2Caller creates a new read-only instance of SubscribeV2, bound to a specific deployed contract.
func NewSubscribeV2Caller(address common.Address, caller bind.ContractCaller) (*SubscribeV2Caller, error) {
	contract, err := bindSubscribeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2Caller{contract: contract}, nil
}

// NewSubscribeV2Transactor creates a new write-only instance of SubscribeV2, bound to a specific deployed contract.
func NewSubscribeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*SubscribeV2Transactor, error) {
	contract, err := bindSubscribeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2Transactor{contract: contract}, nil
}

// NewSubscribeV2Filterer creates a new log filterer instance of SubscribeV2, bound to a specific deployed contract.
func NewSubscribeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*SubscribeV2Filterer, error) {
	contract, err := bindSubscribeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2Filterer{contract: contract}, nil
}

// bindSubscribeV2 binds a generic wrapper to an already deployed contract.
func bindSubscribeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscribeV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscribeV2 *SubscribeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscribeV2.Contract.SubscribeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscribeV2 *SubscribeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SubscribeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscribeV2 *SubscribeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SubscribeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscribeV2 *SubscribeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscribeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscribeV2 *SubscribeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscribeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscribeV2 *SubscribeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscribeV2.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) ADMINROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.ADMINROLE(&_SubscribeV2.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) ADMINROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.ADMINROLE(&_SubscribeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.DEFAULTADMINROLE(&_SubscribeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.DEFAULTADMINROLE(&_SubscribeV2.CallOpts)
}

// EMPLOYROLE is a free data retrieval call binding the contract method 0x4244ddf3.
//
// Solidity: function EMPLOY_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) EMPLOYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "EMPLOY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPLOYROLE is a free data retrieval call binding the contract method 0x4244ddf3.
//
// Solidity: function EMPLOY_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) EMPLOYROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.EMPLOYROLE(&_SubscribeV2.CallOpts)
}

// EMPLOYROLE is a free data retrieval call binding the contract method 0x4244ddf3.
//
// Solidity: function EMPLOY_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) EMPLOYROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.EMPLOYROLE(&_SubscribeV2.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) UPGRADERROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.UPGRADERROLE(&_SubscribeV2.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) UPGRADERROLE() ([32]byte, error) {
	return _SubscribeV2.Contract.UPGRADERROLE(&_SubscribeV2.CallOpts)
}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount, bool _isPaused)
func (_SubscribeV2 *SubscribeV2Caller) ConfigInfo(opts *bind.CallOpts) (struct {
	StartTime          *big.Int
	EndTime            *big.Int
	LimitReserveAmount *big.Int
	IsPaused           bool
}, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "configInfo")

	outstruct := new(struct {
		StartTime          *big.Int
		EndTime            *big.Int
		LimitReserveAmount *big.Int
		IsPaused           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LimitReserveAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsPaused = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount, bool _isPaused)
func (_SubscribeV2 *SubscribeV2Session) ConfigInfo() (struct {
	StartTime          *big.Int
	EndTime            *big.Int
	LimitReserveAmount *big.Int
	IsPaused           bool
}, error) {
	return _SubscribeV2.Contract.ConfigInfo(&_SubscribeV2.CallOpts)
}

// ConfigInfo is a free data retrieval call binding the contract method 0x670017f4.
//
// Solidity: function configInfo() view returns(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount, bool _isPaused)
func (_SubscribeV2 *SubscribeV2CallerSession) ConfigInfo() (struct {
	StartTime          *big.Int
	EndTime            *big.Int
	LimitReserveAmount *big.Int
	IsPaused           bool
}, error) {
	return _SubscribeV2.Contract.ConfigInfo(&_SubscribeV2.CallOpts)
}

// CurrentReserveCount is a free data retrieval call binding the contract method 0x97bdac80.
//
// Solidity: function currentReserveCount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) CurrentReserveCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "currentReserveCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserveCount is a free data retrieval call binding the contract method 0x97bdac80.
//
// Solidity: function currentReserveCount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) CurrentReserveCount() (*big.Int, error) {
	return _SubscribeV2.Contract.CurrentReserveCount(&_SubscribeV2.CallOpts)
}

// CurrentReserveCount is a free data retrieval call binding the contract method 0x97bdac80.
//
// Solidity: function currentReserveCount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) CurrentReserveCount() (*big.Int, error) {
	return _SubscribeV2.Contract.CurrentReserveCount(&_SubscribeV2.CallOpts)
}

// CurrentReserveId is a free data retrieval call binding the contract method 0x5306f102.
//
// Solidity: function currentReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) CurrentReserveId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "currentReserveId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserveId is a free data retrieval call binding the contract method 0x5306f102.
//
// Solidity: function currentReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) CurrentReserveId() (*big.Int, error) {
	return _SubscribeV2.Contract.CurrentReserveId(&_SubscribeV2.CallOpts)
}

// CurrentReserveId is a free data retrieval call binding the contract method 0x5306f102.
//
// Solidity: function currentReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) CurrentReserveId() (*big.Int, error) {
	return _SubscribeV2.Contract.CurrentReserveId(&_SubscribeV2.CallOpts)
}

// CurrentReserveInfo is a free data retrieval call binding the contract method 0xce12d4f7.
//
// Solidity: function currentReserveInfo(uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Caller) CurrentReserveInfo(opts *bind.CallOpts, _index *big.Int) (SubscribeV2ReserveInfo, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "currentReserveInfo", _index)

	if err != nil {
		return *new(SubscribeV2ReserveInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SubscribeV2ReserveInfo)).(*SubscribeV2ReserveInfo)

	return out0, err

}

// CurrentReserveInfo is a free data retrieval call binding the contract method 0xce12d4f7.
//
// Solidity: function currentReserveInfo(uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Session) CurrentReserveInfo(_index *big.Int) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.CurrentReserveInfo(&_SubscribeV2.CallOpts, _index)
}

// CurrentReserveInfo is a free data retrieval call binding the contract method 0xce12d4f7.
//
// Solidity: function currentReserveInfo(uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2CallerSession) CurrentReserveInfo(_index *big.Int) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.CurrentReserveInfo(&_SubscribeV2.CallOpts, _index)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) EndTime() (*big.Int, error) {
	return _SubscribeV2.Contract.EndTime(&_SubscribeV2.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) EndTime() (*big.Int, error) {
	return _SubscribeV2.Contract.EndTime(&_SubscribeV2.CallOpts)
}

// GetReserveId is a free data retrieval call binding the contract method 0x4c3efb3c.
//
// Solidity: function getReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) GetReserveId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "getReserveId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveId is a free data retrieval call binding the contract method 0x4c3efb3c.
//
// Solidity: function getReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) GetReserveId() (*big.Int, error) {
	return _SubscribeV2.Contract.GetReserveId(&_SubscribeV2.CallOpts)
}

// GetReserveId is a free data retrieval call binding the contract method 0x4c3efb3c.
//
// Solidity: function getReserveId() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) GetReserveId() (*big.Int, error) {
	return _SubscribeV2.Contract.GetReserveId(&_SubscribeV2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SubscribeV2.Contract.GetRoleAdmin(&_SubscribeV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SubscribeV2.Contract.GetRoleAdmin(&_SubscribeV2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SubscribeV2.Contract.HasRole(&_SubscribeV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SubscribeV2.Contract.HasRole(&_SubscribeV2.CallOpts, role, account)
}

// IsRefund is a free data retrieval call binding the contract method 0x36cd7f3a.
//
// Solidity: function isRefund(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) IsRefund(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "isRefund", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRefund is a free data retrieval call binding the contract method 0x36cd7f3a.
//
// Solidity: function isRefund(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) IsRefund(arg0 common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsRefund(&_SubscribeV2.CallOpts, arg0)
}

// IsRefund is a free data retrieval call binding the contract method 0x36cd7f3a.
//
// Solidity: function isRefund(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) IsRefund(arg0 common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsRefund(&_SubscribeV2.CallOpts, arg0)
}

// IsReserve is a free data retrieval call binding the contract method 0x7a2b0587.
//
// Solidity: function isReserve(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) IsReserve(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "isReserve", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserve is a free data retrieval call binding the contract method 0x7a2b0587.
//
// Solidity: function isReserve(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) IsReserve(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsReserve(&_SubscribeV2.CallOpts, account)
}

// IsReserve is a free data retrieval call binding the contract method 0x7a2b0587.
//
// Solidity: function isReserve(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) IsReserve(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsReserve(&_SubscribeV2.CallOpts, account)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) IsSubscribed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "isSubscribed", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) IsSubscribed(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsSubscribed(&_SubscribeV2.CallOpts, account)
}

// IsSubscribed is a free data retrieval call binding the contract method 0xb92ae87c.
//
// Solidity: function isSubscribed(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) IsSubscribed(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsSubscribed(&_SubscribeV2.CallOpts, account)
}

// IsWin is a free data retrieval call binding the contract method 0xb08a4f53.
//
// Solidity: function isWin(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) IsWin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "isWin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWin is a free data retrieval call binding the contract method 0xb08a4f53.
//
// Solidity: function isWin(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) IsWin(arg0 common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsWin(&_SubscribeV2.CallOpts, arg0)
}

// IsWin is a free data retrieval call binding the contract method 0xb08a4f53.
//
// Solidity: function isWin(address ) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) IsWin(arg0 common.Address) (bool, error) {
	return _SubscribeV2.Contract.IsWin(&_SubscribeV2.CallOpts, arg0)
}

// LimitReserveAmount is a free data retrieval call binding the contract method 0x70eb3e3d.
//
// Solidity: function limitReserveAmount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) LimitReserveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "limitReserveAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitReserveAmount is a free data retrieval call binding the contract method 0x70eb3e3d.
//
// Solidity: function limitReserveAmount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) LimitReserveAmount() (*big.Int, error) {
	return _SubscribeV2.Contract.LimitReserveAmount(&_SubscribeV2.CallOpts)
}

// LimitReserveAmount is a free data retrieval call binding the contract method 0x70eb3e3d.
//
// Solidity: function limitReserveAmount() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) LimitReserveAmount() (*big.Int, error) {
	return _SubscribeV2.Contract.LimitReserveAmount(&_SubscribeV2.CallOpts)
}

// NftAddr is a free data retrieval call binding the contract method 0x740f1e18.
//
// Solidity: function nftAddr() view returns(address)
func (_SubscribeV2 *SubscribeV2Caller) NftAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "nftAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddr is a free data retrieval call binding the contract method 0x740f1e18.
//
// Solidity: function nftAddr() view returns(address)
func (_SubscribeV2 *SubscribeV2Session) NftAddr() (common.Address, error) {
	return _SubscribeV2.Contract.NftAddr(&_SubscribeV2.CallOpts)
}

// NftAddr is a free data retrieval call binding the contract method 0x740f1e18.
//
// Solidity: function nftAddr() view returns(address)
func (_SubscribeV2 *SubscribeV2CallerSession) NftAddr() (common.Address, error) {
	return _SubscribeV2.Contract.NftAddr(&_SubscribeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) Paused() (bool, error) {
	return _SubscribeV2.Contract.Paused(&_SubscribeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) Paused() (bool, error) {
	return _SubscribeV2.Contract.Paused(&_SubscribeV2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) Price() (*big.Int, error) {
	return _SubscribeV2.Contract.Price(&_SubscribeV2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) Price() (*big.Int, error) {
	return _SubscribeV2.Contract.Price(&_SubscribeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2Session) ProxiableUUID() ([32]byte, error) {
	return _SubscribeV2.Contract.ProxiableUUID(&_SubscribeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SubscribeV2 *SubscribeV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _SubscribeV2.Contract.ProxiableUUID(&_SubscribeV2.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_SubscribeV2 *SubscribeV2Caller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_SubscribeV2 *SubscribeV2Session) Receiver() (common.Address, error) {
	return _SubscribeV2.Contract.Receiver(&_SubscribeV2.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_SubscribeV2 *SubscribeV2CallerSession) Receiver() (common.Address, error) {
	return _SubscribeV2.Contract.Receiver(&_SubscribeV2.CallOpts)
}

// ReserveCountAt is a free data retrieval call binding the contract method 0x3e73dc9f.
//
// Solidity: function reserveCountAt(uint256 _reserveId) view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) ReserveCountAt(opts *bind.CallOpts, _reserveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "reserveCountAt", _reserveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveCountAt is a free data retrieval call binding the contract method 0x3e73dc9f.
//
// Solidity: function reserveCountAt(uint256 _reserveId) view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) ReserveCountAt(_reserveId *big.Int) (*big.Int, error) {
	return _SubscribeV2.Contract.ReserveCountAt(&_SubscribeV2.CallOpts, _reserveId)
}

// ReserveCountAt is a free data retrieval call binding the contract method 0x3e73dc9f.
//
// Solidity: function reserveCountAt(uint256 _reserveId) view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) ReserveCountAt(_reserveId *big.Int) (*big.Int, error) {
	return _SubscribeV2.Contract.ReserveCountAt(&_SubscribeV2.CallOpts, _reserveId)
}

// ReserveInfo is a free data retrieval call binding the contract method 0xe78fb663.
//
// Solidity: function reserveInfo(address ) view returns(uint256 subscribeId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Caller) ReserveInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	SubscribeId *big.Int
	User        common.Address
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "reserveInfo", arg0)

	outstruct := new(struct {
		SubscribeId *big.Int
		User        common.Address
		Amount      *big.Int
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SubscribeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReserveInfo is a free data retrieval call binding the contract method 0xe78fb663.
//
// Solidity: function reserveInfo(address ) view returns(uint256 subscribeId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Session) ReserveInfo(arg0 common.Address) (struct {
	SubscribeId *big.Int
	User        common.Address
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	return _SubscribeV2.Contract.ReserveInfo(&_SubscribeV2.CallOpts, arg0)
}

// ReserveInfo is a free data retrieval call binding the contract method 0xe78fb663.
//
// Solidity: function reserveInfo(address ) view returns(uint256 subscribeId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2CallerSession) ReserveInfo(arg0 common.Address) (struct {
	SubscribeId *big.Int
	User        common.Address
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	return _SubscribeV2.Contract.ReserveInfo(&_SubscribeV2.CallOpts, arg0)
}

// ReserveInfoAt is a free data retrieval call binding the contract method 0x7f949038.
//
// Solidity: function reserveInfoAt(uint256 _reserveId, uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Caller) ReserveInfoAt(opts *bind.CallOpts, _reserveId *big.Int, _index *big.Int) (SubscribeV2ReserveInfo, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "reserveInfoAt", _reserveId, _index)

	if err != nil {
		return *new(SubscribeV2ReserveInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SubscribeV2ReserveInfo)).(*SubscribeV2ReserveInfo)

	return out0, err

}

// ReserveInfoAt is a free data retrieval call binding the contract method 0x7f949038.
//
// Solidity: function reserveInfoAt(uint256 _reserveId, uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Session) ReserveInfoAt(_reserveId *big.Int, _index *big.Int) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.ReserveInfoAt(&_SubscribeV2.CallOpts, _reserveId, _index)
}

// ReserveInfoAt is a free data retrieval call binding the contract method 0x7f949038.
//
// Solidity: function reserveInfoAt(uint256 _reserveId, uint256 _index) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2CallerSession) ReserveInfoAt(_reserveId *big.Int, _index *big.Int) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.ReserveInfoAt(&_SubscribeV2.CallOpts, _reserveId, _index)
}

// Reserveable is a free data retrieval call binding the contract method 0x7c8f1916.
//
// Solidity: function reserveable(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) Reserveable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "reserveable", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Reserveable is a free data retrieval call binding the contract method 0x7c8f1916.
//
// Solidity: function reserveable(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) Reserveable(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.Reserveable(&_SubscribeV2.CallOpts, account)
}

// Reserveable is a free data retrieval call binding the contract method 0x7c8f1916.
//
// Solidity: function reserveable(address account) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) Reserveable(account common.Address) (bool, error) {
	return _SubscribeV2.Contract.Reserveable(&_SubscribeV2.CallOpts, account)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Caller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2Session) StartTime() (*big.Int, error) {
	return _SubscribeV2.Contract.StartTime(&_SubscribeV2.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SubscribeV2 *SubscribeV2CallerSession) StartTime() (*big.Int, error) {
	return _SubscribeV2.Contract.StartTime(&_SubscribeV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubscribeV2 *SubscribeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubscribeV2 *SubscribeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubscribeV2.Contract.SupportsInterface(&_SubscribeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubscribeV2 *SubscribeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubscribeV2.Contract.SupportsInterface(&_SubscribeV2.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SubscribeV2 *SubscribeV2Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SubscribeV2 *SubscribeV2Session) Token() (common.Address, error) {
	return _SubscribeV2.Contract.Token(&_SubscribeV2.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SubscribeV2 *SubscribeV2CallerSession) Token() (common.Address, error) {
	return _SubscribeV2.Contract.Token(&_SubscribeV2.CallOpts)
}

// UserReserveInfo is a free data retrieval call binding the contract method 0xa63b9bc4.
//
// Solidity: function userReserveInfo(address account) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Caller) UserReserveInfo(opts *bind.CallOpts, account common.Address) (SubscribeV2ReserveInfo, error) {
	var out []interface{}
	err := _SubscribeV2.contract.Call(opts, &out, "userReserveInfo", account)

	if err != nil {
		return *new(SubscribeV2ReserveInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SubscribeV2ReserveInfo)).(*SubscribeV2ReserveInfo)

	return out0, err

}

// UserReserveInfo is a free data retrieval call binding the contract method 0xa63b9bc4.
//
// Solidity: function userReserveInfo(address account) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2Session) UserReserveInfo(account common.Address) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.UserReserveInfo(&_SubscribeV2.CallOpts, account)
}

// UserReserveInfo is a free data retrieval call binding the contract method 0xa63b9bc4.
//
// Solidity: function userReserveInfo(address account) view returns((uint256,address,uint256,uint256))
func (_SubscribeV2 *SubscribeV2CallerSession) UserReserveInfo(account common.Address) (SubscribeV2ReserveInfo, error) {
	return _SubscribeV2.Contract.UserReserveInfo(&_SubscribeV2.CallOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.GrantRole(&_SubscribeV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.GrantRole(&_SubscribeV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _token, address _nftAddr, address _receiver, uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2Transactor) Initialize(opts *bind.TransactOpts, _token common.Address, _nftAddr common.Address, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "initialize", _token, _nftAddr, _receiver, _price)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _token, address _nftAddr, address _receiver, uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2Session) Initialize(_token common.Address, _nftAddr common.Address, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.Initialize(&_SubscribeV2.TransactOpts, _token, _nftAddr, _receiver, _price)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _token, address _nftAddr, address _receiver, uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) Initialize(_token common.Address, _nftAddr common.Address, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.Initialize(&_SubscribeV2.TransactOpts, _token, _nftAddr, _receiver, _price)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SubscribeV2 *SubscribeV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SubscribeV2 *SubscribeV2Session) Pause() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Pause(&_SubscribeV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) Pause() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Pause(&_SubscribeV2.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_SubscribeV2 *SubscribeV2Transactor) Refund(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "refund", account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_SubscribeV2 *SubscribeV2Session) Refund(account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.Refund(&_SubscribeV2.TransactOpts, account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) Refund(account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.Refund(&_SubscribeV2.TransactOpts, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.RenounceRole(&_SubscribeV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.RenounceRole(&_SubscribeV2.TransactOpts, role, account)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_SubscribeV2 *SubscribeV2Transactor) Reserve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "reserve")
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_SubscribeV2 *SubscribeV2Session) Reserve() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Reserve(&_SubscribeV2.TransactOpts)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) Reserve() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Reserve(&_SubscribeV2.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.RevokeRole(&_SubscribeV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.RevokeRole(&_SubscribeV2.TransactOpts, role, account)
}

// SetLimitReserveAmount is a paid mutator transaction binding the contract method 0x235816da.
//
// Solidity: function setLimitReserveAmount(uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetLimitReserveAmount(opts *bind.TransactOpts, _limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setLimitReserveAmount", _limitReserveAmount)
}

// SetLimitReserveAmount is a paid mutator transaction binding the contract method 0x235816da.
//
// Solidity: function setLimitReserveAmount(uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2Session) SetLimitReserveAmount(_limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetLimitReserveAmount(&_SubscribeV2.TransactOpts, _limitReserveAmount)
}

// SetLimitReserveAmount is a paid mutator transaction binding the contract method 0x235816da.
//
// Solidity: function setLimitReserveAmount(uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetLimitReserveAmount(_limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetLimitReserveAmount(&_SubscribeV2.TransactOpts, _limitReserveAmount)
}

// SetNFTAddr is a paid mutator transaction binding the contract method 0x79c9ecac.
//
// Solidity: function setNFTAddr(address _nftAddr) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetNFTAddr(opts *bind.TransactOpts, _nftAddr common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setNFTAddr", _nftAddr)
}

// SetNFTAddr is a paid mutator transaction binding the contract method 0x79c9ecac.
//
// Solidity: function setNFTAddr(address _nftAddr) returns()
func (_SubscribeV2 *SubscribeV2Session) SetNFTAddr(_nftAddr common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetNFTAddr(&_SubscribeV2.TransactOpts, _nftAddr)
}

// SetNFTAddr is a paid mutator transaction binding the contract method 0x79c9ecac.
//
// Solidity: function setNFTAddr(address _nftAddr) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetNFTAddr(_nftAddr common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetNFTAddr(&_SubscribeV2.TransactOpts, _nftAddr)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2Session) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetPrice(&_SubscribeV2.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetPrice(&_SubscribeV2.TransactOpts, _price)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetReceiver(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setReceiver", _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_SubscribeV2 *SubscribeV2Session) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetReceiver(&_SubscribeV2.TransactOpts, _receiver)
}

// SetReceiver is a paid mutator transaction binding the contract method 0x718da7ee.
//
// Solidity: function setReceiver(address _receiver) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetReceiver(_receiver common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetReceiver(&_SubscribeV2.TransactOpts, _receiver)
}

// SetTimeAndLimit is a paid mutator transaction binding the contract method 0x6b4c6026.
//
// Solidity: function setTimeAndLimit(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetTimeAndLimit(opts *bind.TransactOpts, _startTime *big.Int, _endTime *big.Int, _limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setTimeAndLimit", _startTime, _endTime, _limitReserveAmount)
}

// SetTimeAndLimit is a paid mutator transaction binding the contract method 0x6b4c6026.
//
// Solidity: function setTimeAndLimit(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2Session) SetTimeAndLimit(_startTime *big.Int, _endTime *big.Int, _limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetTimeAndLimit(&_SubscribeV2.TransactOpts, _startTime, _endTime, _limitReserveAmount)
}

// SetTimeAndLimit is a paid mutator transaction binding the contract method 0x6b4c6026.
//
// Solidity: function setTimeAndLimit(uint256 _startTime, uint256 _endTime, uint256 _limitReserveAmount) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetTimeAndLimit(_startTime *big.Int, _endTime *big.Int, _limitReserveAmount *big.Int) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetTimeAndLimit(&_SubscribeV2.TransactOpts, _startTime, _endTime, _limitReserveAmount)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_SubscribeV2 *SubscribeV2Transactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_SubscribeV2 *SubscribeV2Session) SetToken(_token common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetToken(&_SubscribeV2.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.SetToken(&_SubscribeV2.TransactOpts, _token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SubscribeV2 *SubscribeV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SubscribeV2 *SubscribeV2Session) Unpause() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Unpause(&_SubscribeV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _SubscribeV2.Contract.Unpause(&_SubscribeV2.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SubscribeV2 *SubscribeV2Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SubscribeV2 *SubscribeV2Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.UpgradeTo(&_SubscribeV2.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SubscribeV2.Contract.UpgradeTo(&_SubscribeV2.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SubscribeV2 *SubscribeV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SubscribeV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SubscribeV2 *SubscribeV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SubscribeV2.Contract.UpgradeToAndCall(&_SubscribeV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SubscribeV2 *SubscribeV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SubscribeV2.Contract.UpgradeToAndCall(&_SubscribeV2.TransactOpts, newImplementation, data)
}

// SubscribeV2AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SubscribeV2 contract.
type SubscribeV2AdminChangedIterator struct {
	Event *SubscribeV2AdminChanged // Event containing the contract specifics and raw log

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
func (it *SubscribeV2AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2AdminChanged)
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
		it.Event = new(SubscribeV2AdminChanged)
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
func (it *SubscribeV2AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2AdminChanged represents a AdminChanged event raised by the SubscribeV2 contract.
type SubscribeV2AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SubscribeV2 *SubscribeV2Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*SubscribeV2AdminChangedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2AdminChangedIterator{contract: _SubscribeV2.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SubscribeV2 *SubscribeV2Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SubscribeV2AdminChanged) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2AdminChanged)
				if err := _SubscribeV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseAdminChanged(log types.Log) (*SubscribeV2AdminChanged, error) {
	event := new(SubscribeV2AdminChanged)
	if err := _SubscribeV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SubscribeV2 contract.
type SubscribeV2BeaconUpgradedIterator struct {
	Event *SubscribeV2BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SubscribeV2BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2BeaconUpgraded)
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
		it.Event = new(SubscribeV2BeaconUpgraded)
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
func (it *SubscribeV2BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2BeaconUpgraded represents a BeaconUpgraded event raised by the SubscribeV2 contract.
type SubscribeV2BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SubscribeV2 *SubscribeV2Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SubscribeV2BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2BeaconUpgradedIterator{contract: _SubscribeV2.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SubscribeV2 *SubscribeV2Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SubscribeV2BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2BeaconUpgraded)
				if err := _SubscribeV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseBeaconUpgraded(log types.Log) (*SubscribeV2BeaconUpgraded, error) {
	event := new(SubscribeV2BeaconUpgraded)
	if err := _SubscribeV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SubscribeV2 contract.
type SubscribeV2InitializedIterator struct {
	Event *SubscribeV2Initialized // Event containing the contract specifics and raw log

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
func (it *SubscribeV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Initialized)
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
		it.Event = new(SubscribeV2Initialized)
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
func (it *SubscribeV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Initialized represents a Initialized event raised by the SubscribeV2 contract.
type SubscribeV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SubscribeV2 *SubscribeV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*SubscribeV2InitializedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2InitializedIterator{contract: _SubscribeV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SubscribeV2 *SubscribeV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubscribeV2Initialized) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Initialized)
				if err := _SubscribeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseInitialized(log types.Log) (*SubscribeV2Initialized, error) {
	event := new(SubscribeV2Initialized)
	if err := _SubscribeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SubscribeV2 contract.
type SubscribeV2PausedIterator struct {
	Event *SubscribeV2Paused // Event containing the contract specifics and raw log

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
func (it *SubscribeV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Paused)
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
		it.Event = new(SubscribeV2Paused)
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
func (it *SubscribeV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Paused represents a Paused event raised by the SubscribeV2 contract.
type SubscribeV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SubscribeV2 *SubscribeV2Filterer) FilterPaused(opts *bind.FilterOpts) (*SubscribeV2PausedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2PausedIterator{contract: _SubscribeV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SubscribeV2 *SubscribeV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SubscribeV2Paused) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Paused)
				if err := _SubscribeV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParsePaused(log types.Log) (*SubscribeV2Paused, error) {
	event := new(SubscribeV2Paused)
	if err := _SubscribeV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2RefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the SubscribeV2 contract.
type SubscribeV2RefundedIterator struct {
	Event *SubscribeV2Refunded // Event containing the contract specifics and raw log

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
func (it *SubscribeV2RefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Refunded)
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
		it.Event = new(SubscribeV2Refunded)
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
func (it *SubscribeV2RefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2RefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Refunded represents a Refunded event raised by the SubscribeV2 contract.
type SubscribeV2Refunded struct {
	ReserveId *big.Int
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x279ac10417b4a21068fc6d3150dd2e5fc45320cb722f60a07d37a113a3bf3acd.
//
// Solidity: event Refunded(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) FilterRefunded(opts *bind.FilterOpts) (*SubscribeV2RefundedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2RefundedIterator{contract: _SubscribeV2.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x279ac10417b4a21068fc6d3150dd2e5fc45320cb722f60a07d37a113a3bf3acd.
//
// Solidity: event Refunded(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *SubscribeV2Refunded) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Refunded)
				if err := _SubscribeV2.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x279ac10417b4a21068fc6d3150dd2e5fc45320cb722f60a07d37a113a3bf3acd.
//
// Solidity: event Refunded(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) ParseRefunded(log types.Log) (*SubscribeV2Refunded, error) {
	event := new(SubscribeV2Refunded)
	if err := _SubscribeV2.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2ReservedIterator is returned from FilterReserved and is used to iterate over the raw logs and unpacked data for Reserved events raised by the SubscribeV2 contract.
type SubscribeV2ReservedIterator struct {
	Event *SubscribeV2Reserved // Event containing the contract specifics and raw log

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
func (it *SubscribeV2ReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Reserved)
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
		it.Event = new(SubscribeV2Reserved)
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
func (it *SubscribeV2ReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2ReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Reserved represents a Reserved event raised by the SubscribeV2 contract.
type SubscribeV2Reserved struct {
	ReserveId *big.Int
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReserved is a free log retrieval operation binding the contract event 0xe7e5dbbc37f320d2fc0d3ebb0fa8f453db17c9b2e059b4f82d62765b55b8f288.
//
// Solidity: event Reserved(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) FilterReserved(opts *bind.FilterOpts) (*SubscribeV2ReservedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Reserved")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2ReservedIterator{contract: _SubscribeV2.contract, event: "Reserved", logs: logs, sub: sub}, nil
}

// WatchReserved is a free log subscription operation binding the contract event 0xe7e5dbbc37f320d2fc0d3ebb0fa8f453db17c9b2e059b4f82d62765b55b8f288.
//
// Solidity: event Reserved(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) WatchReserved(opts *bind.WatchOpts, sink chan<- *SubscribeV2Reserved) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Reserved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Reserved)
				if err := _SubscribeV2.contract.UnpackLog(event, "Reserved", log); err != nil {
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

// ParseReserved is a log parse operation binding the contract event 0xe7e5dbbc37f320d2fc0d3ebb0fa8f453db17c9b2e059b4f82d62765b55b8f288.
//
// Solidity: event Reserved(uint256 reserveId, address user, uint256 amount, uint256 timestamp)
func (_SubscribeV2 *SubscribeV2Filterer) ParseReserved(log types.Log) (*SubscribeV2Reserved, error) {
	event := new(SubscribeV2Reserved)
	if err := _SubscribeV2.contract.UnpackLog(event, "Reserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SubscribeV2 contract.
type SubscribeV2RoleAdminChangedIterator struct {
	Event *SubscribeV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SubscribeV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2RoleAdminChanged)
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
		it.Event = new(SubscribeV2RoleAdminChanged)
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
func (it *SubscribeV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2RoleAdminChanged represents a RoleAdminChanged event raised by the SubscribeV2 contract.
type SubscribeV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SubscribeV2 *SubscribeV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SubscribeV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2RoleAdminChangedIterator{contract: _SubscribeV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SubscribeV2 *SubscribeV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SubscribeV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2RoleAdminChanged)
				if err := _SubscribeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseRoleAdminChanged(log types.Log) (*SubscribeV2RoleAdminChanged, error) {
	event := new(SubscribeV2RoleAdminChanged)
	if err := _SubscribeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SubscribeV2 contract.
type SubscribeV2RoleGrantedIterator struct {
	Event *SubscribeV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *SubscribeV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2RoleGranted)
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
		it.Event = new(SubscribeV2RoleGranted)
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
func (it *SubscribeV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2RoleGranted represents a RoleGranted event raised by the SubscribeV2 contract.
type SubscribeV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubscribeV2 *SubscribeV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubscribeV2RoleGrantedIterator, error) {

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

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2RoleGrantedIterator{contract: _SubscribeV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubscribeV2 *SubscribeV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SubscribeV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2RoleGranted)
				if err := _SubscribeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseRoleGranted(log types.Log) (*SubscribeV2RoleGranted, error) {
	event := new(SubscribeV2RoleGranted)
	if err := _SubscribeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SubscribeV2 contract.
type SubscribeV2RoleRevokedIterator struct {
	Event *SubscribeV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *SubscribeV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2RoleRevoked)
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
		it.Event = new(SubscribeV2RoleRevoked)
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
func (it *SubscribeV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2RoleRevoked represents a RoleRevoked event raised by the SubscribeV2 contract.
type SubscribeV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubscribeV2 *SubscribeV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubscribeV2RoleRevokedIterator, error) {

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

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2RoleRevokedIterator{contract: _SubscribeV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubscribeV2 *SubscribeV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SubscribeV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2RoleRevoked)
				if err := _SubscribeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseRoleRevoked(log types.Log) (*SubscribeV2RoleRevoked, error) {
	event := new(SubscribeV2RoleRevoked)
	if err := _SubscribeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SubscribeV2 contract.
type SubscribeV2UnpausedIterator struct {
	Event *SubscribeV2Unpaused // Event containing the contract specifics and raw log

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
func (it *SubscribeV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Unpaused)
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
		it.Event = new(SubscribeV2Unpaused)
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
func (it *SubscribeV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Unpaused represents a Unpaused event raised by the SubscribeV2 contract.
type SubscribeV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SubscribeV2 *SubscribeV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*SubscribeV2UnpausedIterator, error) {

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SubscribeV2UnpausedIterator{contract: _SubscribeV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SubscribeV2 *SubscribeV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SubscribeV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Unpaused)
				if err := _SubscribeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseUnpaused(log types.Log) (*SubscribeV2Unpaused, error) {
	event := new(SubscribeV2Unpaused)
	if err := _SubscribeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscribeV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SubscribeV2 contract.
type SubscribeV2UpgradedIterator struct {
	Event *SubscribeV2Upgraded // Event containing the contract specifics and raw log

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
func (it *SubscribeV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscribeV2Upgraded)
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
		it.Event = new(SubscribeV2Upgraded)
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
func (it *SubscribeV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscribeV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscribeV2Upgraded represents a Upgraded event raised by the SubscribeV2 contract.
type SubscribeV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SubscribeV2 *SubscribeV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SubscribeV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SubscribeV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SubscribeV2UpgradedIterator{contract: _SubscribeV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SubscribeV2 *SubscribeV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SubscribeV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SubscribeV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscribeV2Upgraded)
				if err := _SubscribeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SubscribeV2 *SubscribeV2Filterer) ParseUpgraded(log types.Log) (*SubscribeV2Upgraded, error) {
	event := new(SubscribeV2Upgraded)
	if err := _SubscribeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
