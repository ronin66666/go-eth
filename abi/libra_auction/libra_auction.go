// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package libra_auction

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

// LibraAuctionMetaData contains all meta data concerning the LibraAuction contract.
var LibraAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registryAddrs\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"AuctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"AuctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"AuctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"interrupt\",\"type\":\"bool\"}],\"name\":\"AuctionInterrupted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"HighestBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"Rebate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coin\",\"type\":\"address\"}],\"name\":\"WithDraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"activeAuctionByTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approveFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionBidDatas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"auctionBidOfIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"auctionTotalCountAndPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionTotalPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getFloorPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getHighester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getHighesterAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRefundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApproveFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"isCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"isEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAuctionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coinAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStep\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStep\",\"type\":\"uint256\"}],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reigstry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_i\",\"type\":\"bool\"}],\"name\":\"interrupt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"marketer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rebate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LibraAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use LibraAuctionMetaData.ABI instead.
var LibraAuctionABI = LibraAuctionMetaData.ABI

// LibraAuction is an auto generated Go binding around an Ethereum contract.
type LibraAuction struct {
	LibraAuctionCaller     // Read-only binding to the contract
	LibraAuctionTransactor // Write-only binding to the contract
	LibraAuctionFilterer   // Log filterer for contract events
}

// LibraAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibraAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibraAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibraAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibraAuctionSession struct {
	Contract     *LibraAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibraAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibraAuctionCallerSession struct {
	Contract *LibraAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LibraAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibraAuctionTransactorSession struct {
	Contract     *LibraAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LibraAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibraAuctionRaw struct {
	Contract *LibraAuction // Generic contract binding to access the raw methods on
}

// LibraAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibraAuctionCallerRaw struct {
	Contract *LibraAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// LibraAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibraAuctionTransactorRaw struct {
	Contract *LibraAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibraAuction creates a new instance of LibraAuction, bound to a specific deployed contract.
func NewLibraAuction(address common.Address, backend bind.ContractBackend) (*LibraAuction, error) {
	contract, err := bindLibraAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibraAuction{LibraAuctionCaller: LibraAuctionCaller{contract: contract}, LibraAuctionTransactor: LibraAuctionTransactor{contract: contract}, LibraAuctionFilterer: LibraAuctionFilterer{contract: contract}}, nil
}

// NewLibraAuctionCaller creates a new read-only instance of LibraAuction, bound to a specific deployed contract.
func NewLibraAuctionCaller(address common.Address, caller bind.ContractCaller) (*LibraAuctionCaller, error) {
	contract, err := bindLibraAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionCaller{contract: contract}, nil
}

// NewLibraAuctionTransactor creates a new write-only instance of LibraAuction, bound to a specific deployed contract.
func NewLibraAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*LibraAuctionTransactor, error) {
	contract, err := bindLibraAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionTransactor{contract: contract}, nil
}

// NewLibraAuctionFilterer creates a new log filterer instance of LibraAuction, bound to a specific deployed contract.
func NewLibraAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*LibraAuctionFilterer, error) {
	contract, err := bindLibraAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionFilterer{contract: contract}, nil
}

// bindLibraAuction binds a generic wrapper to an already deployed contract.
func bindLibraAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibraAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibraAuction *LibraAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibraAuction.Contract.LibraAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibraAuction *LibraAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibraAuction.Contract.LibraAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibraAuction *LibraAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibraAuction.Contract.LibraAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibraAuction *LibraAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibraAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibraAuction *LibraAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibraAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibraAuction *LibraAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibraAuction.Contract.contract.Transact(opts, method, params...)
}

// ActiveAuctionByTokenId is a free data retrieval call binding the contract method 0xa47995da.
//
// Solidity: function activeAuctionByTokenId(address tokenAddr, uint256 tokenId) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) ActiveAuctionByTokenId(opts *bind.CallOpts, tokenAddr common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "activeAuctionByTokenId", tokenAddr, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveAuctionByTokenId is a free data retrieval call binding the contract method 0xa47995da.
//
// Solidity: function activeAuctionByTokenId(address tokenAddr, uint256 tokenId) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) ActiveAuctionByTokenId(tokenAddr common.Address, tokenId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.ActiveAuctionByTokenId(&_LibraAuction.CallOpts, tokenAddr, tokenId)
}

// ActiveAuctionByTokenId is a free data retrieval call binding the contract method 0xa47995da.
//
// Solidity: function activeAuctionByTokenId(address tokenAddr, uint256 tokenId) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) ActiveAuctionByTokenId(tokenAddr common.Address, tokenId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.ActiveAuctionByTokenId(&_LibraAuction.CallOpts, tokenAddr, tokenId)
}

// AuctionBidDatas is a free data retrieval call binding the contract method 0x2324a800.
//
// Solidity: function auctionBidDatas(uint256 , uint256 ) view returns(address bidAddress, uint256 bidPrice)
func (_LibraAuction *LibraAuctionCaller) AuctionBidDatas(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	BidAddress common.Address
	BidPrice   *big.Int
}, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "auctionBidDatas", arg0, arg1)

	outstruct := new(struct {
		BidAddress common.Address
		BidPrice   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BidPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AuctionBidDatas is a free data retrieval call binding the contract method 0x2324a800.
//
// Solidity: function auctionBidDatas(uint256 , uint256 ) view returns(address bidAddress, uint256 bidPrice)
func (_LibraAuction *LibraAuctionSession) AuctionBidDatas(arg0 *big.Int, arg1 *big.Int) (struct {
	BidAddress common.Address
	BidPrice   *big.Int
}, error) {
	return _LibraAuction.Contract.AuctionBidDatas(&_LibraAuction.CallOpts, arg0, arg1)
}

// AuctionBidDatas is a free data retrieval call binding the contract method 0x2324a800.
//
// Solidity: function auctionBidDatas(uint256 , uint256 ) view returns(address bidAddress, uint256 bidPrice)
func (_LibraAuction *LibraAuctionCallerSession) AuctionBidDatas(arg0 *big.Int, arg1 *big.Int) (struct {
	BidAddress common.Address
	BidPrice   *big.Int
}, error) {
	return _LibraAuction.Contract.AuctionBidDatas(&_LibraAuction.CallOpts, arg0, arg1)
}

// AuctionBidOfIndex is a free data retrieval call binding the contract method 0x6dd641ac.
//
// Solidity: function auctionBidOfIndex(uint256 auctionId, uint256 index) view returns(address, uint256)
func (_LibraAuction *LibraAuctionCaller) AuctionBidOfIndex(opts *bind.CallOpts, auctionId *big.Int, index *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "auctionBidOfIndex", auctionId, index)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// AuctionBidOfIndex is a free data retrieval call binding the contract method 0x6dd641ac.
//
// Solidity: function auctionBidOfIndex(uint256 auctionId, uint256 index) view returns(address, uint256)
func (_LibraAuction *LibraAuctionSession) AuctionBidOfIndex(auctionId *big.Int, index *big.Int) (common.Address, *big.Int, error) {
	return _LibraAuction.Contract.AuctionBidOfIndex(&_LibraAuction.CallOpts, auctionId, index)
}

// AuctionBidOfIndex is a free data retrieval call binding the contract method 0x6dd641ac.
//
// Solidity: function auctionBidOfIndex(uint256 auctionId, uint256 index) view returns(address, uint256)
func (_LibraAuction *LibraAuctionCallerSession) AuctionBidOfIndex(auctionId *big.Int, index *big.Int) (common.Address, *big.Int, error) {
	return _LibraAuction.Contract.AuctionBidOfIndex(&_LibraAuction.CallOpts, auctionId, index)
}

// AuctionCount is a free data retrieval call binding the contract method 0x898c31be.
//
// Solidity: function auctionCount(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) AuctionCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "auctionCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionCount is a free data retrieval call binding the contract method 0x898c31be.
//
// Solidity: function auctionCount(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) AuctionCount(arg0 *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.AuctionCount(&_LibraAuction.CallOpts, arg0)
}

// AuctionCount is a free data retrieval call binding the contract method 0x898c31be.
//
// Solidity: function auctionCount(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) AuctionCount(arg0 *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.AuctionCount(&_LibraAuction.CallOpts, arg0)
}

// AuctionTotalCountAndPrice is a free data retrieval call binding the contract method 0x4fe46088.
//
// Solidity: function auctionTotalCountAndPrice(uint256 auctionId) view returns(uint256, uint256)
func (_LibraAuction *LibraAuctionCaller) AuctionTotalCountAndPrice(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "auctionTotalCountAndPrice", auctionId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// AuctionTotalCountAndPrice is a free data retrieval call binding the contract method 0x4fe46088.
//
// Solidity: function auctionTotalCountAndPrice(uint256 auctionId) view returns(uint256, uint256)
func (_LibraAuction *LibraAuctionSession) AuctionTotalCountAndPrice(auctionId *big.Int) (*big.Int, *big.Int, error) {
	return _LibraAuction.Contract.AuctionTotalCountAndPrice(&_LibraAuction.CallOpts, auctionId)
}

// AuctionTotalCountAndPrice is a free data retrieval call binding the contract method 0x4fe46088.
//
// Solidity: function auctionTotalCountAndPrice(uint256 auctionId) view returns(uint256, uint256)
func (_LibraAuction *LibraAuctionCallerSession) AuctionTotalCountAndPrice(auctionId *big.Int) (*big.Int, *big.Int, error) {
	return _LibraAuction.Contract.AuctionTotalCountAndPrice(&_LibraAuction.CallOpts, auctionId)
}

// AuctionTotalPrice is a free data retrieval call binding the contract method 0x49f791a2.
//
// Solidity: function auctionTotalPrice(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) AuctionTotalPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "auctionTotalPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionTotalPrice is a free data retrieval call binding the contract method 0x49f791a2.
//
// Solidity: function auctionTotalPrice(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) AuctionTotalPrice(arg0 *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.AuctionTotalPrice(&_LibraAuction.CallOpts, arg0)
}

// AuctionTotalPrice is a free data retrieval call binding the contract method 0x49f791a2.
//
// Solidity: function auctionTotalPrice(uint256 ) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) AuctionTotalPrice(arg0 *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.AuctionTotalPrice(&_LibraAuction.CallOpts, arg0)
}

// GetCoin is a free data retrieval call binding the contract method 0xda311588.
//
// Solidity: function getCoin(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCaller) GetCoin(opts *bind.CallOpts, _auctionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getCoin", _auctionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoin is a free data retrieval call binding the contract method 0xda311588.
//
// Solidity: function getCoin(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionSession) GetCoin(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetCoin(&_LibraAuction.CallOpts, _auctionId)
}

// GetCoin is a free data retrieval call binding the contract method 0xda311588.
//
// Solidity: function getCoin(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCallerSession) GetCoin(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetCoin(&_LibraAuction.CallOpts, _auctionId)
}

// GetFloorPrice is a free data retrieval call binding the contract method 0xb2d4cf18.
//
// Solidity: function getFloorPrice(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) GetFloorPrice(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getFloorPrice", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFloorPrice is a free data retrieval call binding the contract method 0xb2d4cf18.
//
// Solidity: function getFloorPrice(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) GetFloorPrice(_auctionId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.GetFloorPrice(&_LibraAuction.CallOpts, _auctionId)
}

// GetFloorPrice is a free data retrieval call binding the contract method 0xb2d4cf18.
//
// Solidity: function getFloorPrice(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) GetFloorPrice(_auctionId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.GetFloorPrice(&_LibraAuction.CallOpts, _auctionId)
}

// GetHighester is a free data retrieval call binding the contract method 0x8e055293.
//
// Solidity: function getHighester(uint256 _auctionId) view returns(address, uint256)
func (_LibraAuction *LibraAuctionCaller) GetHighester(opts *bind.CallOpts, _auctionId *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getHighester", _auctionId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetHighester is a free data retrieval call binding the contract method 0x8e055293.
//
// Solidity: function getHighester(uint256 _auctionId) view returns(address, uint256)
func (_LibraAuction *LibraAuctionSession) GetHighester(_auctionId *big.Int) (common.Address, *big.Int, error) {
	return _LibraAuction.Contract.GetHighester(&_LibraAuction.CallOpts, _auctionId)
}

// GetHighester is a free data retrieval call binding the contract method 0x8e055293.
//
// Solidity: function getHighester(uint256 _auctionId) view returns(address, uint256)
func (_LibraAuction *LibraAuctionCallerSession) GetHighester(_auctionId *big.Int) (common.Address, *big.Int, error) {
	return _LibraAuction.Contract.GetHighester(&_LibraAuction.CallOpts, _auctionId)
}

// GetHighesterAccount is a free data retrieval call binding the contract method 0x4f0a44a0.
//
// Solidity: function getHighesterAccount(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCaller) GetHighesterAccount(opts *bind.CallOpts, _auctionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getHighesterAccount", _auctionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHighesterAccount is a free data retrieval call binding the contract method 0x4f0a44a0.
//
// Solidity: function getHighesterAccount(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionSession) GetHighesterAccount(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetHighesterAccount(&_LibraAuction.CallOpts, _auctionId)
}

// GetHighesterAccount is a free data retrieval call binding the contract method 0x4f0a44a0.
//
// Solidity: function getHighesterAccount(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCallerSession) GetHighesterAccount(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetHighesterAccount(&_LibraAuction.CallOpts, _auctionId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCaller) GetOwner(opts *bind.CallOpts, _auctionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getOwner", _auctionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionSession) GetOwner(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetOwner(&_LibraAuction.CallOpts, _auctionId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCallerSession) GetOwner(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetOwner(&_LibraAuction.CallOpts, _auctionId)
}

// GetRefundAmount is a free data retrieval call binding the contract method 0xf53bba09.
//
// Solidity: function getRefundAmount(uint256 _auctionId, address account) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) GetRefundAmount(opts *bind.CallOpts, _auctionId *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getRefundAmount", _auctionId, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRefundAmount is a free data retrieval call binding the contract method 0xf53bba09.
//
// Solidity: function getRefundAmount(uint256 _auctionId, address account) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) GetRefundAmount(_auctionId *big.Int, account common.Address) (*big.Int, error) {
	return _LibraAuction.Contract.GetRefundAmount(&_LibraAuction.CallOpts, _auctionId, account)
}

// GetRefundAmount is a free data retrieval call binding the contract method 0xf53bba09.
//
// Solidity: function getRefundAmount(uint256 _auctionId, address account) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) GetRefundAmount(_auctionId *big.Int, account common.Address) (*big.Int, error) {
	return _LibraAuction.Contract.GetRefundAmount(&_LibraAuction.CallOpts, _auctionId, account)
}

// GetRegistry is a free data retrieval call binding the contract method 0x193c045e.
//
// Solidity: function getRegistry(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCaller) GetRegistry(opts *bind.CallOpts, _auctionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getRegistry", _auctionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x193c045e.
//
// Solidity: function getRegistry(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionSession) GetRegistry(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetRegistry(&_LibraAuction.CallOpts, _auctionId)
}

// GetRegistry is a free data retrieval call binding the contract method 0x193c045e.
//
// Solidity: function getRegistry(uint256 _auctionId) view returns(address)
func (_LibraAuction *LibraAuctionCallerSession) GetRegistry(_auctionId *big.Int) (common.Address, error) {
	return _LibraAuction.Contract.GetRegistry(&_LibraAuction.CallOpts, _auctionId)
}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) GetTokenId(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "getTokenId", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionSession) GetTokenId(_auctionId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.GetTokenId(&_LibraAuction.CallOpts, _auctionId)
}

// GetTokenId is a free data retrieval call binding the contract method 0x14ff5ea3.
//
// Solidity: function getTokenId(uint256 _auctionId) view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) GetTokenId(_auctionId *big.Int) (*big.Int, error) {
	return _LibraAuction.Contract.GetTokenId(&_LibraAuction.CallOpts, _auctionId)
}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_LibraAuction *LibraAuctionCaller) IsApproveFor(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "isApproveFor", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_LibraAuction *LibraAuctionSession) IsApproveFor(operator common.Address) (bool, error) {
	return _LibraAuction.Contract.IsApproveFor(&_LibraAuction.CallOpts, operator)
}

// IsApproveFor is a free data retrieval call binding the contract method 0x0e461a10.
//
// Solidity: function isApproveFor(address operator) view returns(bool)
func (_LibraAuction *LibraAuctionCallerSession) IsApproveFor(operator common.Address) (bool, error) {
	return _LibraAuction.Contract.IsApproveFor(&_LibraAuction.CallOpts, operator)
}

// IsCanceled is a free data retrieval call binding the contract method 0x39e3af9b.
//
// Solidity: function isCanceled(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionCaller) IsCanceled(opts *bind.CallOpts, _auctionId *big.Int) (bool, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "isCanceled", _auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCanceled is a free data retrieval call binding the contract method 0x39e3af9b.
//
// Solidity: function isCanceled(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionSession) IsCanceled(_auctionId *big.Int) (bool, error) {
	return _LibraAuction.Contract.IsCanceled(&_LibraAuction.CallOpts, _auctionId)
}

// IsCanceled is a free data retrieval call binding the contract method 0x39e3af9b.
//
// Solidity: function isCanceled(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionCallerSession) IsCanceled(_auctionId *big.Int) (bool, error) {
	return _LibraAuction.Contract.IsCanceled(&_LibraAuction.CallOpts, _auctionId)
}

// IsEnded is a free data retrieval call binding the contract method 0xb66f7a8b.
//
// Solidity: function isEnded(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionCaller) IsEnded(opts *bind.CallOpts, _auctionId *big.Int) (bool, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "isEnded", _auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnded is a free data retrieval call binding the contract method 0xb66f7a8b.
//
// Solidity: function isEnded(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionSession) IsEnded(_auctionId *big.Int) (bool, error) {
	return _LibraAuction.Contract.IsEnded(&_LibraAuction.CallOpts, _auctionId)
}

// IsEnded is a free data retrieval call binding the contract method 0xb66f7a8b.
//
// Solidity: function isEnded(uint256 _auctionId) view returns(bool)
func (_LibraAuction *LibraAuctionCallerSession) IsEnded(_auctionId *big.Int) (bool, error) {
	return _LibraAuction.Contract.IsEnded(&_LibraAuction.CallOpts, _auctionId)
}

// LatestAuctionId is a free data retrieval call binding the contract method 0x4026d63e.
//
// Solidity: function latestAuctionId() view returns(uint256)
func (_LibraAuction *LibraAuctionCaller) LatestAuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "latestAuctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAuctionId is a free data retrieval call binding the contract method 0x4026d63e.
//
// Solidity: function latestAuctionId() view returns(uint256)
func (_LibraAuction *LibraAuctionSession) LatestAuctionId() (*big.Int, error) {
	return _LibraAuction.Contract.LatestAuctionId(&_LibraAuction.CallOpts)
}

// LatestAuctionId is a free data retrieval call binding the contract method 0x4026d63e.
//
// Solidity: function latestAuctionId() view returns(uint256)
func (_LibraAuction *LibraAuctionCallerSession) LatestAuctionId() (*big.Int, error) {
	return _LibraAuction.Contract.LatestAuctionId(&_LibraAuction.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_LibraAuction *LibraAuctionCaller) Members(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_LibraAuction *LibraAuctionSession) Members(arg0 common.Address) (bool, error) {
	return _LibraAuction.Contract.Members(&_LibraAuction.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_LibraAuction *LibraAuctionCallerSession) Members(arg0 common.Address) (bool, error) {
	return _LibraAuction.Contract.Members(&_LibraAuction.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LibraAuction *LibraAuctionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LibraAuction *LibraAuctionSession) Name() (string, error) {
	return _LibraAuction.Contract.Name(&_LibraAuction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LibraAuction *LibraAuctionCallerSession) Name() (string, error) {
	return _LibraAuction.Contract.Name(&_LibraAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LibraAuction *LibraAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LibraAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LibraAuction *LibraAuctionSession) Owner() (common.Address, error) {
	return _LibraAuction.Contract.Owner(&_LibraAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LibraAuction *LibraAuctionCallerSession) Owner() (common.Address, error) {
	return _LibraAuction.Contract.Owner(&_LibraAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0xe50701f4.
//
// Solidity: function abort(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionTransactor) Abort(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "abort", _auctionId)
}

// Abort is a paid mutator transaction binding the contract method 0xe50701f4.
//
// Solidity: function abort(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionSession) Abort(_auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Abort(&_LibraAuction.TransactOpts, _auctionId)
}

// Abort is a paid mutator transaction binding the contract method 0xe50701f4.
//
// Solidity: function abort(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionTransactorSession) Abort(_auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Abort(&_LibraAuction.TransactOpts, _auctionId)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_LibraAuction *LibraAuctionTransactor) ApproveFor(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "approveFor", operator, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_LibraAuction *LibraAuctionSession) ApproveFor(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LibraAuction.Contract.ApproveFor(&_LibraAuction.TransactOpts, operator, approved)
}

// ApproveFor is a paid mutator transaction binding the contract method 0xfb4c4bcd.
//
// Solidity: function approveFor(address operator, bool approved) returns()
func (_LibraAuction *LibraAuctionTransactorSession) ApproveFor(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LibraAuction.Contract.ApproveFor(&_LibraAuction.TransactOpts, operator, approved)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_LibraAuction *LibraAuctionTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_LibraAuction *LibraAuctionSession) Attack() (*types.Transaction, error) {
	return _LibraAuction.Contract.Attack(&_LibraAuction.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() returns()
func (_LibraAuction *LibraAuctionTransactorSession) Attack() (*types.Transaction, error) {
	return _LibraAuction.Contract.Attack(&_LibraAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0xa60bc562.
//
// Solidity: function bid(uint256 auctionId, address bidder, address reigstry, uint256 amount) returns()
func (_LibraAuction *LibraAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidder common.Address, reigstry common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "bid", auctionId, bidder, reigstry, amount)
}

// Bid is a paid mutator transaction binding the contract method 0xa60bc562.
//
// Solidity: function bid(uint256 auctionId, address bidder, address reigstry, uint256 amount) returns()
func (_LibraAuction *LibraAuctionSession) Bid(auctionId *big.Int, bidder common.Address, reigstry common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Bid(&_LibraAuction.TransactOpts, auctionId, bidder, reigstry, amount)
}

// Bid is a paid mutator transaction binding the contract method 0xa60bc562.
//
// Solidity: function bid(uint256 auctionId, address bidder, address reigstry, uint256 amount) returns()
func (_LibraAuction *LibraAuctionTransactorSession) Bid(auctionId *big.Int, bidder common.Address, reigstry common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Bid(&_LibraAuction.TransactOpts, auctionId, bidder, reigstry, amount)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionTransactor) Cancel(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "cancel", _auctionId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionSession) Cancel(_auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Cancel(&_LibraAuction.TransactOpts, _auctionId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _auctionId) returns()
func (_LibraAuction *LibraAuctionTransactorSession) Cancel(_auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Cancel(&_LibraAuction.TransactOpts, _auctionId)
}

// Interrupt is a paid mutator transaction binding the contract method 0xf7208089.
//
// Solidity: function interrupt(uint256 _auctionId, bool _i) returns()
func (_LibraAuction *LibraAuctionTransactor) Interrupt(opts *bind.TransactOpts, _auctionId *big.Int, _i bool) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "interrupt", _auctionId, _i)
}

// Interrupt is a paid mutator transaction binding the contract method 0xf7208089.
//
// Solidity: function interrupt(uint256 _auctionId, bool _i) returns()
func (_LibraAuction *LibraAuctionSession) Interrupt(_auctionId *big.Int, _i bool) (*types.Transaction, error) {
	return _LibraAuction.Contract.Interrupt(&_LibraAuction.TransactOpts, _auctionId, _i)
}

// Interrupt is a paid mutator transaction binding the contract method 0xf7208089.
//
// Solidity: function interrupt(uint256 _auctionId, bool _i) returns()
func (_LibraAuction *LibraAuctionTransactorSession) Interrupt(_auctionId *big.Int, _i bool) (*types.Transaction, error) {
	return _LibraAuction.Contract.Interrupt(&_LibraAuction.TransactOpts, _auctionId, _i)
}

// Rebate is a paid mutator transaction binding the contract method 0xfbdfc9cc.
//
// Solidity: function rebate(uint256 auctionId, address marketer, address to, uint256 amount) returns()
func (_LibraAuction *LibraAuctionTransactor) Rebate(opts *bind.TransactOpts, auctionId *big.Int, marketer common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "rebate", auctionId, marketer, to, amount)
}

// Rebate is a paid mutator transaction binding the contract method 0xfbdfc9cc.
//
// Solidity: function rebate(uint256 auctionId, address marketer, address to, uint256 amount) returns()
func (_LibraAuction *LibraAuctionSession) Rebate(auctionId *big.Int, marketer common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Rebate(&_LibraAuction.TransactOpts, auctionId, marketer, to, amount)
}

// Rebate is a paid mutator transaction binding the contract method 0xfbdfc9cc.
//
// Solidity: function rebate(uint256 auctionId, address marketer, address to, uint256 amount) returns()
func (_LibraAuction *LibraAuctionTransactorSession) Rebate(auctionId *big.Int, marketer common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Rebate(&_LibraAuction.TransactOpts, auctionId, marketer, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LibraAuction *LibraAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LibraAuction *LibraAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _LibraAuction.Contract.RenounceOwnership(&_LibraAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LibraAuction *LibraAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LibraAuction.Contract.RenounceOwnership(&_LibraAuction.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x9d0162a3.
//
// Solidity: function start(address owner, address registry, address tokenAddr, address coinAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep) returns(uint256)
func (_LibraAuction *LibraAuctionTransactor) Start(opts *bind.TransactOpts, owner common.Address, registry common.Address, tokenAddr common.Address, coinAddr common.Address, tokenId *big.Int, price *big.Int, startTime *big.Int, endTime *big.Int, minStep *big.Int, maxStep *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "start", owner, registry, tokenAddr, coinAddr, tokenId, price, startTime, endTime, minStep, maxStep)
}

// Start is a paid mutator transaction binding the contract method 0x9d0162a3.
//
// Solidity: function start(address owner, address registry, address tokenAddr, address coinAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep) returns(uint256)
func (_LibraAuction *LibraAuctionSession) Start(owner common.Address, registry common.Address, tokenAddr common.Address, coinAddr common.Address, tokenId *big.Int, price *big.Int, startTime *big.Int, endTime *big.Int, minStep *big.Int, maxStep *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Start(&_LibraAuction.TransactOpts, owner, registry, tokenAddr, coinAddr, tokenId, price, startTime, endTime, minStep, maxStep)
}

// Start is a paid mutator transaction binding the contract method 0x9d0162a3.
//
// Solidity: function start(address owner, address registry, address tokenAddr, address coinAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep) returns(uint256)
func (_LibraAuction *LibraAuctionTransactorSession) Start(owner common.Address, registry common.Address, tokenAddr common.Address, coinAddr common.Address, tokenId *big.Int, price *big.Int, startTime *big.Int, endTime *big.Int, minStep *big.Int, maxStep *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.Start(&_LibraAuction.TransactOpts, owner, registry, tokenAddr, coinAddr, tokenId, price, startTime, endTime, minStep, maxStep)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LibraAuction *LibraAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LibraAuction *LibraAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LibraAuction.Contract.TransferOwnership(&_LibraAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LibraAuction *LibraAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LibraAuction.Contract.TransferOwnership(&_LibraAuction.TransactOpts, newOwner)
}

// WithDraw is a paid mutator transaction binding the contract method 0x191e2e83.
//
// Solidity: function withDraw(address account, uint256 auctionId) returns()
func (_LibraAuction *LibraAuctionTransactor) WithDraw(opts *bind.TransactOpts, account common.Address, auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.contract.Transact(opts, "withDraw", account, auctionId)
}

// WithDraw is a paid mutator transaction binding the contract method 0x191e2e83.
//
// Solidity: function withDraw(address account, uint256 auctionId) returns()
func (_LibraAuction *LibraAuctionSession) WithDraw(account common.Address, auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.WithDraw(&_LibraAuction.TransactOpts, account, auctionId)
}

// WithDraw is a paid mutator transaction binding the contract method 0x191e2e83.
//
// Solidity: function withDraw(address account, uint256 auctionId) returns()
func (_LibraAuction *LibraAuctionTransactorSession) WithDraw(account common.Address, auctionId *big.Int) (*types.Transaction, error) {
	return _LibraAuction.Contract.WithDraw(&_LibraAuction.TransactOpts, account, auctionId)
}

// LibraAuctionAuctionCanceledIterator is returned from FilterAuctionCanceled and is used to iterate over the raw logs and unpacked data for AuctionCanceled events raised by the LibraAuction contract.
type LibraAuctionAuctionCanceledIterator struct {
	Event *LibraAuctionAuctionCanceled // Event containing the contract specifics and raw log

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
func (it *LibraAuctionAuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionAuctionCanceled)
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
		it.Event = new(LibraAuctionAuctionCanceled)
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
func (it *LibraAuctionAuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionAuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionAuctionCanceled represents a AuctionCanceled event raised by the LibraAuction contract.
type LibraAuctionAuctionCanceled struct {
	AuctionId *big.Int
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCanceled is a free log retrieval operation binding the contract event 0xe933b4aae08ab9bd538fc4c8f819965cfc818a17985d76a5b9d706b24c032c4e.
//
// Solidity: event AuctionCanceled(uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) FilterAuctionCanceled(opts *bind.FilterOpts, AuctionId []*big.Int) (*LibraAuctionAuctionCanceledIterator, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "AuctionCanceled", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionAuctionCanceledIterator{contract: _LibraAuction.contract, event: "AuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchAuctionCanceled is a free log subscription operation binding the contract event 0xe933b4aae08ab9bd538fc4c8f819965cfc818a17985d76a5b9d706b24c032c4e.
//
// Solidity: event AuctionCanceled(uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) WatchAuctionCanceled(opts *bind.WatchOpts, sink chan<- *LibraAuctionAuctionCanceled, AuctionId []*big.Int) (event.Subscription, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "AuctionCanceled", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionAuctionCanceled)
				if err := _LibraAuction.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
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

// ParseAuctionCanceled is a log parse operation binding the contract event 0xe933b4aae08ab9bd538fc4c8f819965cfc818a17985d76a5b9d706b24c032c4e.
//
// Solidity: event AuctionCanceled(uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) ParseAuctionCanceled(log types.Log) (*LibraAuctionAuctionCanceled, error) {
	event := new(LibraAuctionAuctionCanceled)
	if err := _LibraAuction.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the LibraAuction contract.
type LibraAuctionAuctionEndedIterator struct {
	Event *LibraAuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *LibraAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionAuctionEnded)
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
		it.Event = new(LibraAuctionAuctionEnded)
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
func (it *LibraAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionAuctionEnded represents a AuctionEnded event raised by the LibraAuction contract.
type LibraAuctionAuctionEnded struct {
	Winner    common.Address
	Bid       *big.Int
	AuctionId *big.Int
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x2f54d857cdbaa677948ee3e8b7f06f0064a2e99dd79ddab0373560d9bf5b647e.
//
// Solidity: event AuctionEnded(address winner, uint256 bid, uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts, AuctionId []*big.Int) (*LibraAuctionAuctionEndedIterator, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "AuctionEnded", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionAuctionEndedIterator{contract: _LibraAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x2f54d857cdbaa677948ee3e8b7f06f0064a2e99dd79ddab0373560d9bf5b647e.
//
// Solidity: event AuctionEnded(address winner, uint256 bid, uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *LibraAuctionAuctionEnded, AuctionId []*big.Int) (event.Subscription, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "AuctionEnded", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionAuctionEnded)
				if err := _LibraAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0x2f54d857cdbaa677948ee3e8b7f06f0064a2e99dd79ddab0373560d9bf5b647e.
//
// Solidity: event AuctionEnded(address winner, uint256 bid, uint256 indexed AuctionId, uint256 tokenId)
func (_LibraAuction *LibraAuctionFilterer) ParseAuctionEnded(log types.Log) (*LibraAuctionAuctionEnded, error) {
	event := new(LibraAuctionAuctionEnded)
	if err := _LibraAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionAuctionInterruptedIterator is returned from FilterAuctionInterrupted and is used to iterate over the raw logs and unpacked data for AuctionInterrupted events raised by the LibraAuction contract.
type LibraAuctionAuctionInterruptedIterator struct {
	Event *LibraAuctionAuctionInterrupted // Event containing the contract specifics and raw log

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
func (it *LibraAuctionAuctionInterruptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionAuctionInterrupted)
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
		it.Event = new(LibraAuctionAuctionInterrupted)
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
func (it *LibraAuctionAuctionInterruptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionAuctionInterruptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionAuctionInterrupted represents a AuctionInterrupted event raised by the LibraAuction contract.
type LibraAuctionAuctionInterrupted struct {
	AuctionId *big.Int
	TokenId   *big.Int
	Interrupt bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionInterrupted is a free log retrieval operation binding the contract event 0x444b031daa559663f0cd3a7b92fa83c70707589f5eabba93d4eee601066eb474.
//
// Solidity: event AuctionInterrupted(uint256 indexed AuctionId, uint256 tokenId, bool interrupt)
func (_LibraAuction *LibraAuctionFilterer) FilterAuctionInterrupted(opts *bind.FilterOpts, AuctionId []*big.Int) (*LibraAuctionAuctionInterruptedIterator, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "AuctionInterrupted", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionAuctionInterruptedIterator{contract: _LibraAuction.contract, event: "AuctionInterrupted", logs: logs, sub: sub}, nil
}

// WatchAuctionInterrupted is a free log subscription operation binding the contract event 0x444b031daa559663f0cd3a7b92fa83c70707589f5eabba93d4eee601066eb474.
//
// Solidity: event AuctionInterrupted(uint256 indexed AuctionId, uint256 tokenId, bool interrupt)
func (_LibraAuction *LibraAuctionFilterer) WatchAuctionInterrupted(opts *bind.WatchOpts, sink chan<- *LibraAuctionAuctionInterrupted, AuctionId []*big.Int) (event.Subscription, error) {

	var AuctionIdRule []interface{}
	for _, AuctionIdItem := range AuctionId {
		AuctionIdRule = append(AuctionIdRule, AuctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "AuctionInterrupted", AuctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionAuctionInterrupted)
				if err := _LibraAuction.contract.UnpackLog(event, "AuctionInterrupted", log); err != nil {
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

// ParseAuctionInterrupted is a log parse operation binding the contract event 0x444b031daa559663f0cd3a7b92fa83c70707589f5eabba93d4eee601066eb474.
//
// Solidity: event AuctionInterrupted(uint256 indexed AuctionId, uint256 tokenId, bool interrupt)
func (_LibraAuction *LibraAuctionFilterer) ParseAuctionInterrupted(log types.Log) (*LibraAuctionAuctionInterrupted, error) {
	event := new(LibraAuctionAuctionInterrupted)
	if err := _LibraAuction.contract.UnpackLog(event, "AuctionInterrupted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the LibraAuction contract.
type LibraAuctionAuctionStartedIterator struct {
	Event *LibraAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *LibraAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionAuctionStarted)
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
		it.Event = new(LibraAuctionAuctionStarted)
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
func (it *LibraAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionAuctionStarted represents a AuctionStarted event raised by the LibraAuction contract.
type LibraAuctionAuctionStarted struct {
	Owner     common.Address
	TokenAddr common.Address
	TokenId   *big.Int
	Price     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	MinStep   *big.Int
	MaxStep   *big.Int
	Registry  common.Address
	Coin      common.Address
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x1979a83c5b9f52924d80f5b8972feb639204f2143ab564da4d43587ca762f57c.
//
// Solidity: event AuctionStarted(address owner, address tokenAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep, address registry, address coin, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, auctionId []*big.Int) (*LibraAuctionAuctionStartedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "AuctionStarted", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionAuctionStartedIterator{contract: _LibraAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x1979a83c5b9f52924d80f5b8972feb639204f2143ab564da4d43587ca762f57c.
//
// Solidity: event AuctionStarted(address owner, address tokenAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep, address registry, address coin, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *LibraAuctionAuctionStarted, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "AuctionStarted", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionAuctionStarted)
				if err := _LibraAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0x1979a83c5b9f52924d80f5b8972feb639204f2143ab564da4d43587ca762f57c.
//
// Solidity: event AuctionStarted(address owner, address tokenAddr, uint256 tokenId, uint256 price, uint256 startTime, uint256 endTime, uint256 minStep, uint256 maxStep, address registry, address coin, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) ParseAuctionStarted(log types.Log) (*LibraAuctionAuctionStarted, error) {
	event := new(LibraAuctionAuctionStarted)
	if err := _LibraAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionHighestBidIterator is returned from FilterHighestBid and is used to iterate over the raw logs and unpacked data for HighestBid events raised by the LibraAuction contract.
type LibraAuctionHighestBidIterator struct {
	Event *LibraAuctionHighestBid // Event containing the contract specifics and raw log

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
func (it *LibraAuctionHighestBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionHighestBid)
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
		it.Event = new(LibraAuctionHighestBid)
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
func (it *LibraAuctionHighestBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionHighestBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionHighestBid represents a HighestBid event raised by the LibraAuction contract.
type LibraAuctionHighestBid struct {
	Bidder    common.Address
	Bid       *big.Int
	TokenId   *big.Int
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHighestBid is a free log retrieval operation binding the contract event 0x7cced2bec86085b46d4dc8fe08e75dbbcac6b9abe42d76bf5373182e6d8fbb90.
//
// Solidity: event HighestBid(address bidder, uint256 bid, uint256 tokenId, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) FilterHighestBid(opts *bind.FilterOpts, auctionId []*big.Int) (*LibraAuctionHighestBidIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "HighestBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionHighestBidIterator{contract: _LibraAuction.contract, event: "HighestBid", logs: logs, sub: sub}, nil
}

// WatchHighestBid is a free log subscription operation binding the contract event 0x7cced2bec86085b46d4dc8fe08e75dbbcac6b9abe42d76bf5373182e6d8fbb90.
//
// Solidity: event HighestBid(address bidder, uint256 bid, uint256 tokenId, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) WatchHighestBid(opts *bind.WatchOpts, sink chan<- *LibraAuctionHighestBid, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "HighestBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionHighestBid)
				if err := _LibraAuction.contract.UnpackLog(event, "HighestBid", log); err != nil {
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

// ParseHighestBid is a log parse operation binding the contract event 0x7cced2bec86085b46d4dc8fe08e75dbbcac6b9abe42d76bf5373182e6d8fbb90.
//
// Solidity: event HighestBid(address bidder, uint256 bid, uint256 tokenId, uint256 indexed auctionId)
func (_LibraAuction *LibraAuctionFilterer) ParseHighestBid(log types.Log) (*LibraAuctionHighestBid, error) {
	event := new(LibraAuctionHighestBid)
	if err := _LibraAuction.contract.UnpackLog(event, "HighestBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LibraAuction contract.
type LibraAuctionOwnershipTransferredIterator struct {
	Event *LibraAuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LibraAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionOwnershipTransferred)
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
		it.Event = new(LibraAuctionOwnershipTransferred)
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
func (it *LibraAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the LibraAuction contract.
type LibraAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibraAuction *LibraAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LibraAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionOwnershipTransferredIterator{contract: _LibraAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibraAuction *LibraAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LibraAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionOwnershipTransferred)
				if err := _LibraAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LibraAuction *LibraAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*LibraAuctionOwnershipTransferred, error) {
	event := new(LibraAuctionOwnershipTransferred)
	if err := _LibraAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionRebateIterator is returned from FilterRebate and is used to iterate over the raw logs and unpacked data for Rebate events raised by the LibraAuction contract.
type LibraAuctionRebateIterator struct {
	Event *LibraAuctionRebate // Event containing the contract specifics and raw log

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
func (it *LibraAuctionRebateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionRebate)
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
		it.Event = new(LibraAuctionRebate)
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
func (it *LibraAuctionRebateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionRebateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionRebate represents a Rebate event raised by the LibraAuction contract.
type LibraAuctionRebate struct {
	AuctionId *big.Int
	Amount    *big.Int
	From      common.Address
	To        common.Address
	Coin      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRebate is a free log retrieval operation binding the contract event 0x7fde9e75bd6b0aef16d611ff25dea887cb9999073654703103e55b1861df1d08.
//
// Solidity: event Rebate(uint256 indexed auctionId, uint256 amount, address from, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) FilterRebate(opts *bind.FilterOpts, auctionId []*big.Int) (*LibraAuctionRebateIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "Rebate", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionRebateIterator{contract: _LibraAuction.contract, event: "Rebate", logs: logs, sub: sub}, nil
}

// WatchRebate is a free log subscription operation binding the contract event 0x7fde9e75bd6b0aef16d611ff25dea887cb9999073654703103e55b1861df1d08.
//
// Solidity: event Rebate(uint256 indexed auctionId, uint256 amount, address from, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) WatchRebate(opts *bind.WatchOpts, sink chan<- *LibraAuctionRebate, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "Rebate", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionRebate)
				if err := _LibraAuction.contract.UnpackLog(event, "Rebate", log); err != nil {
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

// ParseRebate is a log parse operation binding the contract event 0x7fde9e75bd6b0aef16d611ff25dea887cb9999073654703103e55b1861df1d08.
//
// Solidity: event Rebate(uint256 indexed auctionId, uint256 amount, address from, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) ParseRebate(log types.Log) (*LibraAuctionRebate, error) {
	event := new(LibraAuctionRebate)
	if err := _LibraAuction.contract.UnpackLog(event, "Rebate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraAuctionWithDrawIterator is returned from FilterWithDraw and is used to iterate over the raw logs and unpacked data for WithDraw events raised by the LibraAuction contract.
type LibraAuctionWithDrawIterator struct {
	Event *LibraAuctionWithDraw // Event containing the contract specifics and raw log

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
func (it *LibraAuctionWithDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraAuctionWithDraw)
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
		it.Event = new(LibraAuctionWithDraw)
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
func (it *LibraAuctionWithDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraAuctionWithDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraAuctionWithDraw represents a WithDraw event raised by the LibraAuction contract.
type LibraAuctionWithDraw struct {
	AuctionId *big.Int
	Amount    *big.Int
	To        common.Address
	Coin      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithDraw is a free log retrieval operation binding the contract event 0x9a23a8ea9c98a4c7be088300ca3aa71452bb8a0701192f32840ceb38d60e28a9.
//
// Solidity: event WithDraw(uint256 indexed auctionId, uint256 amount, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) FilterWithDraw(opts *bind.FilterOpts, auctionId []*big.Int) (*LibraAuctionWithDrawIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.FilterLogs(opts, "WithDraw", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &LibraAuctionWithDrawIterator{contract: _LibraAuction.contract, event: "WithDraw", logs: logs, sub: sub}, nil
}

// WatchWithDraw is a free log subscription operation binding the contract event 0x9a23a8ea9c98a4c7be088300ca3aa71452bb8a0701192f32840ceb38d60e28a9.
//
// Solidity: event WithDraw(uint256 indexed auctionId, uint256 amount, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) WatchWithDraw(opts *bind.WatchOpts, sink chan<- *LibraAuctionWithDraw, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _LibraAuction.contract.WatchLogs(opts, "WithDraw", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraAuctionWithDraw)
				if err := _LibraAuction.contract.UnpackLog(event, "WithDraw", log); err != nil {
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

// ParseWithDraw is a log parse operation binding the contract event 0x9a23a8ea9c98a4c7be088300ca3aa71452bb8a0701192f32840ceb38d60e28a9.
//
// Solidity: event WithDraw(uint256 indexed auctionId, uint256 amount, address to, address coin)
func (_LibraAuction *LibraAuctionFilterer) ParseWithDraw(log types.Log) (*LibraAuctionWithDraw, error) {
	event := new(LibraAuctionWithDraw)
	if err := _LibraAuction.contract.UnpackLog(event, "WithDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
