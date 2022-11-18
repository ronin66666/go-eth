// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wyvern_exchange

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

// WyvernExchangeMetaData contains all meta data concerning the WyvernExchange contract.
var WyvernExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"registryAddrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"customPersonalSignPrefix\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staticTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"staticSelector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumFill\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"OrderApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFill\",\"type\":\"uint256\"}],\"name\":\"OrderFillChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"firstHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"secondHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstMaker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"secondMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFirstFill\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSecondFill\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approveOrderHash_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staticTarget\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"staticSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maximumFill\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"approveOrder_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[16]\",\"name\":\"uints\",\"type\":\"uint256[16]\"},{\"internalType\":\"bytes4[2]\",\"name\":\"staticSelectors\",\"type\":\"bytes4[2]\"},{\"internalType\":\"bytes\",\"name\":\"firstExtradata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"firstCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"secondExtradata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"secondCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint8[2]\",\"name\":\"howToCalls\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"atomicMatch_\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fills\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staticTarget\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"staticSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maximumFill\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"hashOrder_\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"hashToSign_\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registries\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fill\",\"type\":\"uint256\"}],\"name\":\"setOrderFill_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateOrderAuthorization_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staticTarget\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"staticSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maximumFill\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"validateOrderParameters_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approveFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WyvernExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use WyvernExchangeMetaData.ABI instead.
var WyvernExchangeABI = WyvernExchangeMetaData.ABI

// WyvernExchange is an auto generated Go binding around an Ethereum contract.
type WyvernExchange struct {
	WyvernExchangeCaller     // Read-only binding to the contract
	WyvernExchangeTransactor // Write-only binding to the contract
	WyvernExchangeFilterer   // Log filterer for contract events
}

// WyvernExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WyvernExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WyvernExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WyvernExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WyvernExchangeSession struct {
	Contract     *WyvernExchange   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WyvernExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WyvernExchangeCallerSession struct {
	Contract *WyvernExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WyvernExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WyvernExchangeTransactorSession struct {
	Contract     *WyvernExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WyvernExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WyvernExchangeRaw struct {
	Contract *WyvernExchange // Generic contract binding to access the raw methods on
}

// WyvernExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WyvernExchangeCallerRaw struct {
	Contract *WyvernExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// WyvernExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WyvernExchangeTransactorRaw struct {
	Contract *WyvernExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWyvernExchange creates a new instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchange(address common.Address, backend bind.ContractBackend) (*WyvernExchange, error) {
	contract, err := bindWyvernExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WyvernExchange{WyvernExchangeCaller: WyvernExchangeCaller{contract: contract}, WyvernExchangeTransactor: WyvernExchangeTransactor{contract: contract}, WyvernExchangeFilterer: WyvernExchangeFilterer{contract: contract}}, nil
}

// NewWyvernExchangeCaller creates a new read-only instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeCaller(address common.Address, caller bind.ContractCaller) (*WyvernExchangeCaller, error) {
	contract, err := bindWyvernExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeCaller{contract: contract}, nil
}

// NewWyvernExchangeTransactor creates a new write-only instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*WyvernExchangeTransactor, error) {
	contract, err := bindWyvernExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeTransactor{contract: contract}, nil
}

// NewWyvernExchangeFilterer creates a new log filterer instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*WyvernExchangeFilterer, error) {
	contract, err := bindWyvernExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeFilterer{contract: contract}, nil
}

// bindWyvernExchange binds a generic wrapper to an already deployed contract.
func bindWyvernExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WyvernExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WyvernExchange *WyvernExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WyvernExchange.Contract.WyvernExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WyvernExchange *WyvernExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.Contract.WyvernExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WyvernExchange *WyvernExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WyvernExchange.Contract.WyvernExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WyvernExchange *WyvernExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WyvernExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WyvernExchange *WyvernExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WyvernExchange *WyvernExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WyvernExchange.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0xfac7fc38.
//
// Solidity: function approved(address , bytes32 ) view returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) Approved(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "approved", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approved is a free data retrieval call binding the contract method 0xfac7fc38.
//
// Solidity: function approved(address , bytes32 ) view returns(bool)
func (_WyvernExchange *WyvernExchangeSession) Approved(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.Approved(&_WyvernExchange.CallOpts, arg0, arg1)
}

// Approved is a free data retrieval call binding the contract method 0xfac7fc38.
//
// Solidity: function approved(address , bytes32 ) view returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) Approved(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.Approved(&_WyvernExchange.CallOpts, arg0, arg1)
}

// Fills is a free data retrieval call binding the contract method 0xffc2ab14.
//
// Solidity: function fills(address , bytes32 ) view returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) Fills(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "fills", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fills is a free data retrieval call binding the contract method 0xffc2ab14.
//
// Solidity: function fills(address , bytes32 ) view returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) Fills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _WyvernExchange.Contract.Fills(&_WyvernExchange.CallOpts, arg0, arg1)
}

// Fills is a free data retrieval call binding the contract method 0xffc2ab14.
//
// Solidity: function fills(address , bytes32 ) view returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) Fills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _WyvernExchange.Contract.Fills(&_WyvernExchange.CallOpts, arg0, arg1)
}

// HashOrder is a free data retrieval call binding the contract method 0xf861975b.
//
// Solidity: function hashOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) pure returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeCaller) HashOrder(opts *bind.CallOpts, registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "hashOrder_", registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0xf861975b.
//
// Solidity: function hashOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) pure returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeSession) HashOrder(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) ([32]byte, error) {
	return _WyvernExchange.Contract.HashOrder(&_WyvernExchange.CallOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)
}

// HashOrder is a free data retrieval call binding the contract method 0xf861975b.
//
// Solidity: function hashOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) pure returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeCallerSession) HashOrder(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) ([32]byte, error) {
	return _WyvernExchange.Contract.HashOrder(&_WyvernExchange.CallOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)
}

// HashToSign is a free data retrieval call binding the contract method 0xd57ad588.
//
// Solidity: function hashToSign_(bytes32 orderHash) view returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeCaller) HashToSign(opts *bind.CallOpts, orderHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "hashToSign_", orderHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashToSign is a free data retrieval call binding the contract method 0xd57ad588.
//
// Solidity: function hashToSign_(bytes32 orderHash) view returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeSession) HashToSign(orderHash [32]byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashToSign(&_WyvernExchange.CallOpts, orderHash)
}

// HashToSign is a free data retrieval call binding the contract method 0xd57ad588.
//
// Solidity: function hashToSign_(bytes32 orderHash) view returns(bytes32 hash)
func (_WyvernExchange *WyvernExchangeCallerSession) HashToSign(orderHash [32]byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashToSign(&_WyvernExchange.CallOpts, orderHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WyvernExchange *WyvernExchangeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WyvernExchange *WyvernExchangeSession) Name() (string, error) {
	return _WyvernExchange.Contract.Name(&_WyvernExchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WyvernExchange *WyvernExchangeCallerSession) Name() (string, error) {
	return _WyvernExchange.Contract.Name(&_WyvernExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WyvernExchange *WyvernExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WyvernExchange *WyvernExchangeSession) Owner() (common.Address, error) {
	return _WyvernExchange.Contract.Owner(&_WyvernExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) Owner() (common.Address, error) {
	return _WyvernExchange.Contract.Owner(&_WyvernExchange.CallOpts)
}

// Registries is a free data retrieval call binding the contract method 0xcaed80df.
//
// Solidity: function registries(address ) view returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) Registries(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0xcaed80df.
//
// Solidity: function registries(address ) view returns(bool)
func (_WyvernExchange *WyvernExchangeSession) Registries(arg0 common.Address) (bool, error) {
	return _WyvernExchange.Contract.Registries(&_WyvernExchange.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0xcaed80df.
//
// Solidity: function registries(address ) view returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) Registries(arg0 common.Address) (bool, error) {
	return _WyvernExchange.Contract.Registries(&_WyvernExchange.CallOpts, arg0)
}

// ValidateOrderAuthorization is a free data retrieval call binding the contract method 0x13b219e2.
//
// Solidity: function validateOrderAuthorization_(bytes32 hash, address maker, bytes signature) view returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) ValidateOrderAuthorization(opts *bind.CallOpts, hash [32]byte, maker common.Address, signature []byte) (bool, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "validateOrderAuthorization_", hash, maker, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrderAuthorization is a free data retrieval call binding the contract method 0x13b219e2.
//
// Solidity: function validateOrderAuthorization_(bytes32 hash, address maker, bytes signature) view returns(bool)
func (_WyvernExchange *WyvernExchangeSession) ValidateOrderAuthorization(hash [32]byte, maker common.Address, signature []byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderAuthorization(&_WyvernExchange.CallOpts, hash, maker, signature)
}

// ValidateOrderAuthorization is a free data retrieval call binding the contract method 0x13b219e2.
//
// Solidity: function validateOrderAuthorization_(bytes32 hash, address maker, bytes signature) view returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) ValidateOrderAuthorization(hash [32]byte, maker common.Address, signature []byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderAuthorization(&_WyvernExchange.CallOpts, hash, maker, signature)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xf9b0bc20.
//
// Solidity: function validateOrderParameters_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) view returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) ValidateOrderParameters(opts *bind.CallOpts, registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) (bool, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "validateOrderParameters_", registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xf9b0bc20.
//
// Solidity: function validateOrderParameters_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) view returns(bool)
func (_WyvernExchange *WyvernExchangeSession) ValidateOrderParameters(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderParameters(&_WyvernExchange.CallOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xf9b0bc20.
//
// Solidity: function validateOrderParameters_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt) view returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) ValidateOrderParameters(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderParameters(&_WyvernExchange.CallOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WyvernExchange *WyvernExchangeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WyvernExchange.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WyvernExchange *WyvernExchangeSession) Version() (string, error) {
	return _WyvernExchange.Contract.Version(&_WyvernExchange.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_WyvernExchange *WyvernExchangeCallerSession) Version() (string, error) {
	return _WyvernExchange.Contract.Version(&_WyvernExchange.CallOpts)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address registry, bool approved) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ApproveFor(opts *bind.TransactOpts, registry common.Address, approved bool) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "approveFor", registry, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address registry, bool approved) returns()
func (_WyvernExchange *WyvernExchangeSession) ApproveFor(registry common.Address, approved bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveFor(&_WyvernExchange.TransactOpts, registry, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address registry, bool approved) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ApproveFor(registry common.Address, approved bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveFor(&_WyvernExchange.TransactOpts, registry, approved)
}

// ApproveOrderHash is a paid mutator transaction binding the contract method 0x0812226e.
//
// Solidity: function approveOrderHash_(bytes32 hash) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ApproveOrderHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "approveOrderHash_", hash)
}

// ApproveOrderHash is a paid mutator transaction binding the contract method 0x0812226e.
//
// Solidity: function approveOrderHash_(bytes32 hash) returns()
func (_WyvernExchange *WyvernExchangeSession) ApproveOrderHash(hash [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrderHash(&_WyvernExchange.TransactOpts, hash)
}

// ApproveOrderHash is a paid mutator transaction binding the contract method 0x0812226e.
//
// Solidity: function approveOrderHash_(bytes32 hash) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ApproveOrderHash(hash [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrderHash(&_WyvernExchange.TransactOpts, hash)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x995e8195.
//
// Solidity: function approveOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ApproveOrder(opts *bind.TransactOpts, registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "approveOrder_", registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x995e8195.
//
// Solidity: function approveOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeSession) ApproveOrder(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrder(&_WyvernExchange.TransactOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x995e8195.
//
// Solidity: function approveOrder_(address registry, address maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ApproveOrder(registry common.Address, maker common.Address, staticTarget common.Address, staticSelector [4]byte, staticExtradata []byte, maximumFill *big.Int, listingTime *big.Int, expirationTime *big.Int, salt *big.Int, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrder(&_WyvernExchange.TransactOpts, registry, maker, staticTarget, staticSelector, staticExtradata, maximumFill, listingTime, expirationTime, salt, orderbookInclusionDesired)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x9a5168db.
//
// Solidity: function atomicMatch_(uint256[16] uints, bytes4[2] staticSelectors, bytes firstExtradata, bytes firstCalldata, bytes secondExtradata, bytes secondCalldata, uint8[2] howToCalls, bytes32 metadata, bytes signatures) payable returns()
func (_WyvernExchange *WyvernExchangeTransactor) AtomicMatch(opts *bind.TransactOpts, uints [16]*big.Int, staticSelectors [2][4]byte, firstExtradata []byte, firstCalldata []byte, secondExtradata []byte, secondCalldata []byte, howToCalls [2]uint8, metadata [32]byte, signatures []byte) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "atomicMatch_", uints, staticSelectors, firstExtradata, firstCalldata, secondExtradata, secondCalldata, howToCalls, metadata, signatures)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x9a5168db.
//
// Solidity: function atomicMatch_(uint256[16] uints, bytes4[2] staticSelectors, bytes firstExtradata, bytes firstCalldata, bytes secondExtradata, bytes secondCalldata, uint8[2] howToCalls, bytes32 metadata, bytes signatures) payable returns()
func (_WyvernExchange *WyvernExchangeSession) AtomicMatch(uints [16]*big.Int, staticSelectors [2][4]byte, firstExtradata []byte, firstCalldata []byte, secondExtradata []byte, secondCalldata []byte, howToCalls [2]uint8, metadata [32]byte, signatures []byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.AtomicMatch(&_WyvernExchange.TransactOpts, uints, staticSelectors, firstExtradata, firstCalldata, secondExtradata, secondCalldata, howToCalls, metadata, signatures)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x9a5168db.
//
// Solidity: function atomicMatch_(uint256[16] uints, bytes4[2] staticSelectors, bytes firstExtradata, bytes firstCalldata, bytes secondExtradata, bytes secondCalldata, uint8[2] howToCalls, bytes32 metadata, bytes signatures) payable returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) AtomicMatch(uints [16]*big.Int, staticSelectors [2][4]byte, firstExtradata []byte, firstCalldata []byte, secondExtradata []byte, secondCalldata []byte, howToCalls [2]uint8, metadata [32]byte, signatures []byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.AtomicMatch(&_WyvernExchange.TransactOpts, uints, staticSelectors, firstExtradata, firstCalldata, secondExtradata, secondCalldata, howToCalls, metadata, signatures)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_WyvernExchange *WyvernExchangeTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_WyvernExchange *WyvernExchangeSession) Attack() (*types.Transaction, error) {
	return _WyvernExchange.Contract.Attack(&_WyvernExchange.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) Attack() (*types.Transaction, error) {
	return _WyvernExchange.Contract.Attack(&_WyvernExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _WyvernExchange.Contract.RenounceOwnership(&_WyvernExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WyvernExchange.Contract.RenounceOwnership(&_WyvernExchange.TransactOpts)
}

// SetOrderFill is a paid mutator transaction binding the contract method 0x14127f13.
//
// Solidity: function setOrderFill_(bytes32 hash, uint256 fill) returns()
func (_WyvernExchange *WyvernExchangeTransactor) SetOrderFill(opts *bind.TransactOpts, hash [32]byte, fill *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "setOrderFill_", hash, fill)
}

// SetOrderFill is a paid mutator transaction binding the contract method 0x14127f13.
//
// Solidity: function setOrderFill_(bytes32 hash, uint256 fill) returns()
func (_WyvernExchange *WyvernExchangeSession) SetOrderFill(hash [32]byte, fill *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.SetOrderFill(&_WyvernExchange.TransactOpts, hash, fill)
}

// SetOrderFill is a paid mutator transaction binding the contract method 0x14127f13.
//
// Solidity: function setOrderFill_(bytes32 hash, uint256 fill) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) SetOrderFill(hash [32]byte, fill *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.SetOrderFill(&_WyvernExchange.TransactOpts, hash, fill)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.TransferOwnership(&_WyvernExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.TransferOwnership(&_WyvernExchange.TransactOpts, newOwner)
}

// WyvernExchangeOrderApprovedIterator is returned from FilterOrderApproved and is used to iterate over the raw logs and unpacked data for OrderApproved events raised by the WyvernExchange contract.
type WyvernExchangeOrderApprovedIterator struct {
	Event *WyvernExchangeOrderApproved // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrderApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrderApproved)
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
		it.Event = new(WyvernExchangeOrderApproved)
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
func (it *WyvernExchangeOrderApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrderApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrderApproved represents a OrderApproved event raised by the WyvernExchange contract.
type WyvernExchangeOrderApproved struct {
	Hash                      [32]byte
	Registry                  common.Address
	Maker                     common.Address
	StaticTarget              common.Address
	StaticSelector            [4]byte
	StaticExtradata           []byte
	MaximumFill               *big.Int
	ListingTime               *big.Int
	ExpirationTime            *big.Int
	Salt                      *big.Int
	OrderbookInclusionDesired bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterOrderApproved is a free log retrieval operation binding the contract event 0xbace5d282ba59cba3de1ac539b0ed46ae5a3d8ad434a3ac9f07ce45d16614818.
//
// Solidity: event OrderApproved(bytes32 indexed hash, address registry, address indexed maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrderApproved(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address) (*WyvernExchangeOrderApprovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrderApproved", hashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrderApprovedIterator{contract: _WyvernExchange.contract, event: "OrderApproved", logs: logs, sub: sub}, nil
}

// WatchOrderApproved is a free log subscription operation binding the contract event 0xbace5d282ba59cba3de1ac539b0ed46ae5a3d8ad434a3ac9f07ce45d16614818.
//
// Solidity: event OrderApproved(bytes32 indexed hash, address registry, address indexed maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrderApproved(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrderApproved, hash [][32]byte, maker []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrderApproved", hashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrderApproved)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrderApproved", log); err != nil {
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

// ParseOrderApproved is a log parse operation binding the contract event 0xbace5d282ba59cba3de1ac539b0ed46ae5a3d8ad434a3ac9f07ce45d16614818.
//
// Solidity: event OrderApproved(bytes32 indexed hash, address registry, address indexed maker, address staticTarget, bytes4 staticSelector, bytes staticExtradata, uint256 maximumFill, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_WyvernExchange *WyvernExchangeFilterer) ParseOrderApproved(log types.Log) (*WyvernExchangeOrderApproved, error) {
	event := new(WyvernExchangeOrderApproved)
	if err := _WyvernExchange.contract.UnpackLog(event, "OrderApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WyvernExchangeOrderFillChangedIterator is returned from FilterOrderFillChanged and is used to iterate over the raw logs and unpacked data for OrderFillChanged events raised by the WyvernExchange contract.
type WyvernExchangeOrderFillChangedIterator struct {
	Event *WyvernExchangeOrderFillChanged // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrderFillChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrderFillChanged)
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
		it.Event = new(WyvernExchangeOrderFillChanged)
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
func (it *WyvernExchangeOrderFillChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrderFillChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrderFillChanged represents a OrderFillChanged event raised by the WyvernExchange contract.
type WyvernExchangeOrderFillChanged struct {
	Hash    [32]byte
	Maker   common.Address
	NewFill *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFillChanged is a free log retrieval operation binding the contract event 0x826cd05b52546b47590739469989cacd5336e0a9e6f9392de2dee5cfb19c8aab.
//
// Solidity: event OrderFillChanged(bytes32 indexed hash, address indexed maker, uint256 newFill)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrderFillChanged(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address) (*WyvernExchangeOrderFillChangedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrderFillChanged", hashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrderFillChangedIterator{contract: _WyvernExchange.contract, event: "OrderFillChanged", logs: logs, sub: sub}, nil
}

// WatchOrderFillChanged is a free log subscription operation binding the contract event 0x826cd05b52546b47590739469989cacd5336e0a9e6f9392de2dee5cfb19c8aab.
//
// Solidity: event OrderFillChanged(bytes32 indexed hash, address indexed maker, uint256 newFill)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrderFillChanged(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrderFillChanged, hash [][32]byte, maker []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrderFillChanged", hashRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrderFillChanged)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrderFillChanged", log); err != nil {
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

// ParseOrderFillChanged is a log parse operation binding the contract event 0x826cd05b52546b47590739469989cacd5336e0a9e6f9392de2dee5cfb19c8aab.
//
// Solidity: event OrderFillChanged(bytes32 indexed hash, address indexed maker, uint256 newFill)
func (_WyvernExchange *WyvernExchangeFilterer) ParseOrderFillChanged(log types.Log) (*WyvernExchangeOrderFillChanged, error) {
	event := new(WyvernExchangeOrderFillChanged)
	if err := _WyvernExchange.contract.UnpackLog(event, "OrderFillChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WyvernExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the WyvernExchange contract.
type WyvernExchangeOrdersMatchedIterator struct {
	Event *WyvernExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrdersMatched)
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
		it.Event = new(WyvernExchangeOrdersMatched)
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
func (it *WyvernExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrdersMatched represents a OrdersMatched event raised by the WyvernExchange contract.
type WyvernExchangeOrdersMatched struct {
	FirstHash     [32]byte
	SecondHash    [32]byte
	FirstMaker    common.Address
	SecondMaker   common.Address
	NewFirstFill  *big.Int
	NewSecondFill *big.Int
	Metadata      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0xe8447df45ce371d4c2f6cfcb9342b06165f71b244e809b909977f0dab191a009.
//
// Solidity: event OrdersMatched(bytes32 firstHash, bytes32 secondHash, address indexed firstMaker, address indexed secondMaker, uint256 newFirstFill, uint256 newSecondFill, bytes32 indexed metadata)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, firstMaker []common.Address, secondMaker []common.Address, metadata [][32]byte) (*WyvernExchangeOrdersMatchedIterator, error) {

	var firstMakerRule []interface{}
	for _, firstMakerItem := range firstMaker {
		firstMakerRule = append(firstMakerRule, firstMakerItem)
	}
	var secondMakerRule []interface{}
	for _, secondMakerItem := range secondMaker {
		secondMakerRule = append(secondMakerRule, secondMakerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrdersMatched", firstMakerRule, secondMakerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrdersMatchedIterator{contract: _WyvernExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0xe8447df45ce371d4c2f6cfcb9342b06165f71b244e809b909977f0dab191a009.
//
// Solidity: event OrdersMatched(bytes32 firstHash, bytes32 secondHash, address indexed firstMaker, address indexed secondMaker, uint256 newFirstFill, uint256 newSecondFill, bytes32 indexed metadata)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrdersMatched, firstMaker []common.Address, secondMaker []common.Address, metadata [][32]byte) (event.Subscription, error) {

	var firstMakerRule []interface{}
	for _, firstMakerItem := range firstMaker {
		firstMakerRule = append(firstMakerRule, firstMakerItem)
	}
	var secondMakerRule []interface{}
	for _, secondMakerItem := range secondMaker {
		secondMakerRule = append(secondMakerRule, secondMakerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrdersMatched", firstMakerRule, secondMakerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrdersMatched)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0xe8447df45ce371d4c2f6cfcb9342b06165f71b244e809b909977f0dab191a009.
//
// Solidity: event OrdersMatched(bytes32 firstHash, bytes32 secondHash, address indexed firstMaker, address indexed secondMaker, uint256 newFirstFill, uint256 newSecondFill, bytes32 indexed metadata)
func (_WyvernExchange *WyvernExchangeFilterer) ParseOrdersMatched(log types.Log) (*WyvernExchangeOrdersMatched, error) {
	event := new(WyvernExchangeOrdersMatched)
	if err := _WyvernExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WyvernExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WyvernExchange contract.
type WyvernExchangeOwnershipTransferredIterator struct {
	Event *WyvernExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOwnershipTransferred)
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
		it.Event = new(WyvernExchangeOwnershipTransferred)
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
func (it *WyvernExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the WyvernExchange contract.
type WyvernExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WyvernExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOwnershipTransferredIterator{contract: _WyvernExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOwnershipTransferred)
				if err := _WyvernExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WyvernExchange *WyvernExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*WyvernExchangeOwnershipTransferred, error) {
	event := new(WyvernExchangeOwnershipTransferred)
	if err := _WyvernExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
