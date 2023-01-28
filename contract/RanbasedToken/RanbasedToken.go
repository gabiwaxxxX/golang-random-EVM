// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RanbasedToken

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

// RanbasedTokenMetaData contains all meta data concerning the RanbasedToken contract.
var RanbasedTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"LogRebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TRANSACTION_LIMIT_STATUS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WALLET_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSED\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RANB_WETH_pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"supplyDelta\",\"type\":\"int256\"}],\"name\":\"rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setMaxTransactionLimitStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"}],\"name\":\"setPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RanbasedTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RanbasedTokenMetaData.ABI instead.
var RanbasedTokenABI = RanbasedTokenMetaData.ABI

// RanbasedToken is an auto generated Go binding around an Ethereum contract.
type RanbasedToken struct {
	RanbasedTokenCaller     // Read-only binding to the contract
	RanbasedTokenTransactor // Write-only binding to the contract
	RanbasedTokenFilterer   // Log filterer for contract events
}

// RanbasedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RanbasedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RanbasedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RanbasedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RanbasedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RanbasedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RanbasedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RanbasedTokenSession struct {
	Contract     *RanbasedToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RanbasedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RanbasedTokenCallerSession struct {
	Contract *RanbasedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RanbasedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RanbasedTokenTransactorSession struct {
	Contract     *RanbasedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RanbasedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RanbasedTokenRaw struct {
	Contract *RanbasedToken // Generic contract binding to access the raw methods on
}

// RanbasedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RanbasedTokenCallerRaw struct {
	Contract *RanbasedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RanbasedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RanbasedTokenTransactorRaw struct {
	Contract *RanbasedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRanbasedToken creates a new instance of RanbasedToken, bound to a specific deployed contract.
func NewRanbasedToken(address common.Address, backend bind.ContractBackend) (*RanbasedToken, error) {
	contract, err := bindRanbasedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RanbasedToken{RanbasedTokenCaller: RanbasedTokenCaller{contract: contract}, RanbasedTokenTransactor: RanbasedTokenTransactor{contract: contract}, RanbasedTokenFilterer: RanbasedTokenFilterer{contract: contract}}, nil
}

// NewRanbasedTokenCaller creates a new read-only instance of RanbasedToken, bound to a specific deployed contract.
func NewRanbasedTokenCaller(address common.Address, caller bind.ContractCaller) (*RanbasedTokenCaller, error) {
	contract, err := bindRanbasedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenCaller{contract: contract}, nil
}

// NewRanbasedTokenTransactor creates a new write-only instance of RanbasedToken, bound to a specific deployed contract.
func NewRanbasedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RanbasedTokenTransactor, error) {
	contract, err := bindRanbasedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenTransactor{contract: contract}, nil
}

// NewRanbasedTokenFilterer creates a new log filterer instance of RanbasedToken, bound to a specific deployed contract.
func NewRanbasedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RanbasedTokenFilterer, error) {
	contract, err := bindRanbasedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenFilterer{contract: contract}, nil
}

// bindRanbasedToken binds a generic wrapper to an already deployed contract.
func bindRanbasedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RanbasedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RanbasedToken *RanbasedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RanbasedToken.Contract.RanbasedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RanbasedToken *RanbasedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RanbasedToken.Contract.RanbasedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RanbasedToken *RanbasedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RanbasedToken.Contract.RanbasedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RanbasedToken *RanbasedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RanbasedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RanbasedToken *RanbasedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RanbasedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RanbasedToken *RanbasedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RanbasedToken.Contract.contract.Transact(opts, method, params...)
}

// MAXTRANSACTIONLIMITSTATUS is a free data retrieval call binding the contract method 0xf90231bc.
//
// Solidity: function MAX_TRANSACTION_LIMIT_STATUS() view returns(bool)
func (_RanbasedToken *RanbasedTokenCaller) MAXTRANSACTIONLIMITSTATUS(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "MAX_TRANSACTION_LIMIT_STATUS")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MAXTRANSACTIONLIMITSTATUS is a free data retrieval call binding the contract method 0xf90231bc.
//
// Solidity: function MAX_TRANSACTION_LIMIT_STATUS() view returns(bool)
func (_RanbasedToken *RanbasedTokenSession) MAXTRANSACTIONLIMITSTATUS() (bool, error) {
	return _RanbasedToken.Contract.MAXTRANSACTIONLIMITSTATUS(&_RanbasedToken.CallOpts)
}

// MAXTRANSACTIONLIMITSTATUS is a free data retrieval call binding the contract method 0xf90231bc.
//
// Solidity: function MAX_TRANSACTION_LIMIT_STATUS() view returns(bool)
func (_RanbasedToken *RanbasedTokenCallerSession) MAXTRANSACTIONLIMITSTATUS() (bool, error) {
	return _RanbasedToken.Contract.MAXTRANSACTIONLIMITSTATUS(&_RanbasedToken.CallOpts)
}

// MAXTXSIZE is a free data retrieval call binding the contract method 0xc03c2cf9.
//
// Solidity: function MAX_TX_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCaller) MAXTXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "MAX_TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTXSIZE is a free data retrieval call binding the contract method 0xc03c2cf9.
//
// Solidity: function MAX_TX_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) MAXTXSIZE() (*big.Int, error) {
	return _RanbasedToken.Contract.MAXTXSIZE(&_RanbasedToken.CallOpts)
}

// MAXTXSIZE is a free data retrieval call binding the contract method 0xc03c2cf9.
//
// Solidity: function MAX_TX_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCallerSession) MAXTXSIZE() (*big.Int, error) {
	return _RanbasedToken.Contract.MAXTXSIZE(&_RanbasedToken.CallOpts)
}

// MAXWALLETSIZE is a free data retrieval call binding the contract method 0x4f054317.
//
// Solidity: function MAX_WALLET_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCaller) MAXWALLETSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "MAX_WALLET_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWALLETSIZE is a free data retrieval call binding the contract method 0x4f054317.
//
// Solidity: function MAX_WALLET_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) MAXWALLETSIZE() (*big.Int, error) {
	return _RanbasedToken.Contract.MAXWALLETSIZE(&_RanbasedToken.CallOpts)
}

// MAXWALLETSIZE is a free data retrieval call binding the contract method 0x4f054317.
//
// Solidity: function MAX_WALLET_SIZE() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCallerSession) MAXWALLETSIZE() (*big.Int, error) {
	return _RanbasedToken.Contract.MAXWALLETSIZE(&_RanbasedToken.CallOpts)
}

// PAUSED is a free data retrieval call binding the contract method 0xa9aad58c.
//
// Solidity: function PAUSED() view returns(bool)
func (_RanbasedToken *RanbasedTokenCaller) PAUSED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "PAUSED")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PAUSED is a free data retrieval call binding the contract method 0xa9aad58c.
//
// Solidity: function PAUSED() view returns(bool)
func (_RanbasedToken *RanbasedTokenSession) PAUSED() (bool, error) {
	return _RanbasedToken.Contract.PAUSED(&_RanbasedToken.CallOpts)
}

// PAUSED is a free data retrieval call binding the contract method 0xa9aad58c.
//
// Solidity: function PAUSED() view returns(bool)
func (_RanbasedToken *RanbasedTokenCallerSession) PAUSED() (bool, error) {
	return _RanbasedToken.Contract.PAUSED(&_RanbasedToken.CallOpts)
}

// RANBWETHPair is a free data retrieval call binding the contract method 0xc091e179.
//
// Solidity: function RANB_WETH_pair() view returns(address)
func (_RanbasedToken *RanbasedTokenCaller) RANBWETHPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "RANB_WETH_pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RANBWETHPair is a free data retrieval call binding the contract method 0xc091e179.
//
// Solidity: function RANB_WETH_pair() view returns(address)
func (_RanbasedToken *RanbasedTokenSession) RANBWETHPair() (common.Address, error) {
	return _RanbasedToken.Contract.RANBWETHPair(&_RanbasedToken.CallOpts)
}

// RANBWETHPair is a free data retrieval call binding the contract method 0xc091e179.
//
// Solidity: function RANB_WETH_pair() view returns(address)
func (_RanbasedToken *RanbasedTokenCallerSession) RANBWETHPair() (common.Address, error) {
	return _RanbasedToken.Contract.RANBWETHPair(&_RanbasedToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_RanbasedToken *RanbasedTokenCaller) Allowance(opts *bind.CallOpts, owner_ common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "allowance", owner_, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) Allowance(owner_ common.Address, spender common.Address) (*big.Int, error) {
	return _RanbasedToken.Contract.Allowance(&_RanbasedToken.CallOpts, owner_, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_RanbasedToken *RanbasedTokenCallerSession) Allowance(owner_ common.Address, spender common.Address) (*big.Int, error) {
	return _RanbasedToken.Contract.Allowance(&_RanbasedToken.CallOpts, owner_, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_RanbasedToken *RanbasedTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _RanbasedToken.Contract.BalanceOf(&_RanbasedToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_RanbasedToken *RanbasedTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _RanbasedToken.Contract.BalanceOf(&_RanbasedToken.CallOpts, who)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_RanbasedToken *RanbasedTokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_RanbasedToken *RanbasedTokenSession) Controller() (common.Address, error) {
	return _RanbasedToken.Contract.Controller(&_RanbasedToken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_RanbasedToken *RanbasedTokenCallerSession) Controller() (common.Address, error) {
	return _RanbasedToken.Contract.Controller(&_RanbasedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RanbasedToken *RanbasedTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RanbasedToken *RanbasedTokenSession) Decimals() (uint8, error) {
	return _RanbasedToken.Contract.Decimals(&_RanbasedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RanbasedToken *RanbasedTokenCallerSession) Decimals() (uint8, error) {
	return _RanbasedToken.Contract.Decimals(&_RanbasedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RanbasedToken *RanbasedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RanbasedToken *RanbasedTokenSession) Name() (string, error) {
	return _RanbasedToken.Contract.Name(&_RanbasedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RanbasedToken *RanbasedTokenCallerSession) Name() (string, error) {
	return _RanbasedToken.Contract.Name(&_RanbasedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RanbasedToken *RanbasedTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RanbasedToken *RanbasedTokenSession) Owner() (common.Address, error) {
	return _RanbasedToken.Contract.Owner(&_RanbasedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RanbasedToken *RanbasedTokenCallerSession) Owner() (common.Address, error) {
	return _RanbasedToken.Contract.Owner(&_RanbasedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RanbasedToken *RanbasedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RanbasedToken *RanbasedTokenSession) Symbol() (string, error) {
	return _RanbasedToken.Contract.Symbol(&_RanbasedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RanbasedToken *RanbasedTokenCallerSession) Symbol() (string, error) {
	return _RanbasedToken.Contract.Symbol(&_RanbasedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RanbasedToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) TotalSupply() (*big.Int, error) {
	return _RanbasedToken.Contract.TotalSupply(&_RanbasedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RanbasedToken *RanbasedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RanbasedToken.Contract.TotalSupply(&_RanbasedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Approve(&_RanbasedToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Approve(&_RanbasedToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.DecreaseAllowance(&_RanbasedToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.DecreaseAllowance(&_RanbasedToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.IncreaseAllowance(&_RanbasedToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.IncreaseAllowance(&_RanbasedToken.TransactOpts, spender, addedValue)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_RanbasedToken *RanbasedTokenTransactor) Rebase(opts *bind.TransactOpts, epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "rebase", epoch, supplyDelta)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_RanbasedToken *RanbasedTokenSession) Rebase(epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Rebase(&_RanbasedToken.TransactOpts, epoch, supplyDelta)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_RanbasedToken *RanbasedTokenTransactorSession) Rebase(epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Rebase(&_RanbasedToken.TransactOpts, epoch, supplyDelta)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RanbasedToken *RanbasedTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RanbasedToken *RanbasedTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _RanbasedToken.Contract.RenounceOwnership(&_RanbasedToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RanbasedToken.Contract.RenounceOwnership(&_RanbasedToken.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_RanbasedToken *RanbasedTokenTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_RanbasedToken *RanbasedTokenSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetController(&_RanbasedToken.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetController(&_RanbasedToken.TransactOpts, _controller)
}

// SetMaxTransactionLimitStatus is a paid mutator transaction binding the contract method 0x6360de2d.
//
// Solidity: function setMaxTransactionLimitStatus(bool _status) returns()
func (_RanbasedToken *RanbasedTokenTransactor) SetMaxTransactionLimitStatus(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "setMaxTransactionLimitStatus", _status)
}

// SetMaxTransactionLimitStatus is a paid mutator transaction binding the contract method 0x6360de2d.
//
// Solidity: function setMaxTransactionLimitStatus(bool _status) returns()
func (_RanbasedToken *RanbasedTokenSession) SetMaxTransactionLimitStatus(_status bool) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetMaxTransactionLimitStatus(&_RanbasedToken.TransactOpts, _status)
}

// SetMaxTransactionLimitStatus is a paid mutator transaction binding the contract method 0x6360de2d.
//
// Solidity: function setMaxTransactionLimitStatus(bool _status) returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) SetMaxTransactionLimitStatus(_status bool) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetMaxTransactionLimitStatus(&_RanbasedToken.TransactOpts, _status)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address _pair) returns()
func (_RanbasedToken *RanbasedTokenTransactor) SetPair(opts *bind.TransactOpts, _pair common.Address) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "setPair", _pair)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address _pair) returns()
func (_RanbasedToken *RanbasedTokenSession) SetPair(_pair common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetPair(&_RanbasedToken.TransactOpts, _pair)
}

// SetPair is a paid mutator transaction binding the contract method 0x8187f516.
//
// Solidity: function setPair(address _pair) returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) SetPair(_pair common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetPair(&_RanbasedToken.TransactOpts, _pair)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_RanbasedToken *RanbasedTokenTransactor) SetPaused(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "setPaused", _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_RanbasedToken *RanbasedTokenSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetPaused(&_RanbasedToken.TransactOpts, _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _RanbasedToken.Contract.SetPaused(&_RanbasedToken.TransactOpts, _status)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Transfer(&_RanbasedToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.Transfer(&_RanbasedToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.TransferFrom(&_RanbasedToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RanbasedToken *RanbasedTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RanbasedToken.Contract.TransferFrom(&_RanbasedToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RanbasedToken *RanbasedTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RanbasedToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RanbasedToken *RanbasedTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.TransferOwnership(&_RanbasedToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RanbasedToken *RanbasedTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RanbasedToken.Contract.TransferOwnership(&_RanbasedToken.TransactOpts, newOwner)
}

// RanbasedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RanbasedToken contract.
type RanbasedTokenApprovalIterator struct {
	Event *RanbasedTokenApproval // Event containing the contract specifics and raw log

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
func (it *RanbasedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RanbasedTokenApproval)
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
		it.Event = new(RanbasedTokenApproval)
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
func (it *RanbasedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RanbasedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RanbasedTokenApproval represents a Approval event raised by the RanbasedToken contract.
type RanbasedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RanbasedToken *RanbasedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RanbasedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RanbasedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenApprovalIterator{contract: _RanbasedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RanbasedToken *RanbasedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RanbasedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RanbasedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RanbasedTokenApproval)
				if err := _RanbasedToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RanbasedToken *RanbasedTokenFilterer) ParseApproval(log types.Log) (*RanbasedTokenApproval, error) {
	event := new(RanbasedTokenApproval)
	if err := _RanbasedToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RanbasedTokenLogRebaseIterator is returned from FilterLogRebase and is used to iterate over the raw logs and unpacked data for LogRebase events raised by the RanbasedToken contract.
type RanbasedTokenLogRebaseIterator struct {
	Event *RanbasedTokenLogRebase // Event containing the contract specifics and raw log

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
func (it *RanbasedTokenLogRebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RanbasedTokenLogRebase)
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
		it.Event = new(RanbasedTokenLogRebase)
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
func (it *RanbasedTokenLogRebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RanbasedTokenLogRebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RanbasedTokenLogRebase represents a LogRebase event raised by the RanbasedToken contract.
type RanbasedTokenLogRebase struct {
	Epoch       *big.Int
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRebase is a free log retrieval operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_RanbasedToken *RanbasedTokenFilterer) FilterLogRebase(opts *bind.FilterOpts, epoch []*big.Int) (*RanbasedTokenLogRebaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RanbasedToken.contract.FilterLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenLogRebaseIterator{contract: _RanbasedToken.contract, event: "LogRebase", logs: logs, sub: sub}, nil
}

// WatchLogRebase is a free log subscription operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_RanbasedToken *RanbasedTokenFilterer) WatchLogRebase(opts *bind.WatchOpts, sink chan<- *RanbasedTokenLogRebase, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RanbasedToken.contract.WatchLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RanbasedTokenLogRebase)
				if err := _RanbasedToken.contract.UnpackLog(event, "LogRebase", log); err != nil {
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

// ParseLogRebase is a log parse operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_RanbasedToken *RanbasedTokenFilterer) ParseLogRebase(log types.Log) (*RanbasedTokenLogRebase, error) {
	event := new(RanbasedTokenLogRebase)
	if err := _RanbasedToken.contract.UnpackLog(event, "LogRebase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RanbasedTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RanbasedToken contract.
type RanbasedTokenOwnershipTransferredIterator struct {
	Event *RanbasedTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RanbasedTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RanbasedTokenOwnershipTransferred)
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
		it.Event = new(RanbasedTokenOwnershipTransferred)
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
func (it *RanbasedTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RanbasedTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RanbasedTokenOwnershipTransferred represents a OwnershipTransferred event raised by the RanbasedToken contract.
type RanbasedTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RanbasedToken *RanbasedTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RanbasedTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RanbasedToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenOwnershipTransferredIterator{contract: _RanbasedToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RanbasedToken *RanbasedTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RanbasedTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RanbasedToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RanbasedTokenOwnershipTransferred)
				if err := _RanbasedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RanbasedToken *RanbasedTokenFilterer) ParseOwnershipTransferred(log types.Log) (*RanbasedTokenOwnershipTransferred, error) {
	event := new(RanbasedTokenOwnershipTransferred)
	if err := _RanbasedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RanbasedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RanbasedToken contract.
type RanbasedTokenTransferIterator struct {
	Event *RanbasedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RanbasedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RanbasedTokenTransfer)
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
		it.Event = new(RanbasedTokenTransfer)
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
func (it *RanbasedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RanbasedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RanbasedTokenTransfer represents a Transfer event raised by the RanbasedToken contract.
type RanbasedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RanbasedToken *RanbasedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RanbasedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RanbasedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RanbasedTokenTransferIterator{contract: _RanbasedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RanbasedToken *RanbasedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RanbasedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RanbasedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RanbasedTokenTransfer)
				if err := _RanbasedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RanbasedToken *RanbasedTokenFilterer) ParseTransfer(log types.Log) (*RanbasedTokenTransfer, error) {
	event := new(RanbasedTokenTransfer)
	if err := _RanbasedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
