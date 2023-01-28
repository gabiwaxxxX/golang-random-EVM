// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SURGE

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

// SURGEMetaData contains all meta data concerning the SURGE contract.
var SURGEMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beans\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dollarBuy\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beans\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dollarSell\",\"type\":\"uint256\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"_buy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBNBOut\",\"type\":\"uint256\"}],\"name\":\"_sell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candleStickData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"close\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"high\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTeamWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryWallet\",\"type\":\"address\"}],\"name\":\"changeFeeReceivers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newbuyMul\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newsellMul\",\"type\":\"uint256\"}],\"name\":\"changeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exempt\",\"type\":\"bool\"}],\"name\":\"changeIsFeeExempt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exempt\",\"type\":\"bool\"}],\"name\":\"changeIsTxLimitExempt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStablePair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newStableAddress\",\"type\":\"address\"}],\"name\":\"changeStablePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newteamShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newtreasuryShare\",\"type\":\"uint256\"}],\"name\":\"changeTaxDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getBNBAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBNBPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCirculatingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBNBIn\",\"type\":\"uint256\"}],\"name\":\"getTokenAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"getValueOfHoldings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indVol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isFeeExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTxLimitExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liqConst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBag\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareDIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tVol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVolume\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeOpenTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"txTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTaxBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SURGEABI is the input ABI used to generate the binding from.
// Deprecated: Use SURGEMetaData.ABI instead.
var SURGEABI = SURGEMetaData.ABI

// SURGE is an auto generated Go binding around an Ethereum contract.
type SURGE struct {
	SURGECaller     // Read-only binding to the contract
	SURGETransactor // Write-only binding to the contract
	SURGEFilterer   // Log filterer for contract events
}

// SURGECaller is an auto generated read-only Go binding around an Ethereum contract.
type SURGECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SURGETransactor is an auto generated write-only Go binding around an Ethereum contract.
type SURGETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SURGEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SURGEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SURGESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SURGESession struct {
	Contract     *SURGE            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SURGECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SURGECallerSession struct {
	Contract *SURGECaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SURGETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SURGETransactorSession struct {
	Contract     *SURGETransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SURGERaw is an auto generated low-level Go binding around an Ethereum contract.
type SURGERaw struct {
	Contract *SURGE // Generic contract binding to access the raw methods on
}

// SURGECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SURGECallerRaw struct {
	Contract *SURGECaller // Generic read-only contract binding to access the raw methods on
}

// SURGETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SURGETransactorRaw struct {
	Contract *SURGETransactor // Generic write-only contract binding to access the raw methods on
}

// NewSURGE creates a new instance of SURGE, bound to a specific deployed contract.
func NewSURGE(address common.Address, backend bind.ContractBackend) (*SURGE, error) {
	contract, err := bindSURGE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SURGE{SURGECaller: SURGECaller{contract: contract}, SURGETransactor: SURGETransactor{contract: contract}, SURGEFilterer: SURGEFilterer{contract: contract}}, nil
}

// NewSURGECaller creates a new read-only instance of SURGE, bound to a specific deployed contract.
func NewSURGECaller(address common.Address, caller bind.ContractCaller) (*SURGECaller, error) {
	contract, err := bindSURGE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SURGECaller{contract: contract}, nil
}

// NewSURGETransactor creates a new write-only instance of SURGE, bound to a specific deployed contract.
func NewSURGETransactor(address common.Address, transactor bind.ContractTransactor) (*SURGETransactor, error) {
	contract, err := bindSURGE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SURGETransactor{contract: contract}, nil
}

// NewSURGEFilterer creates a new log filterer instance of SURGE, bound to a specific deployed contract.
func NewSURGEFilterer(address common.Address, filterer bind.ContractFilterer) (*SURGEFilterer, error) {
	contract, err := bindSURGE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SURGEFilterer{contract: contract}, nil
}

// bindSURGE binds a generic wrapper to an already deployed contract.
func bindSURGE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SURGEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SURGE *SURGERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SURGE.Contract.SURGECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SURGE *SURGERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SURGE.Contract.SURGETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SURGE *SURGERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SURGE.Contract.SURGETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SURGE *SURGECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SURGE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SURGE *SURGETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SURGE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SURGE *SURGETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SURGE.Contract.contract.Transact(opts, method, params...)
}

// DIVISOR is a free data retrieval call binding the contract method 0x3410fe6e.
//
// Solidity: function DIVISOR() view returns(uint256)
func (_SURGE *SURGECaller) DIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DIVISOR is a free data retrieval call binding the contract method 0x3410fe6e.
//
// Solidity: function DIVISOR() view returns(uint256)
func (_SURGE *SURGESession) DIVISOR() (*big.Int, error) {
	return _SURGE.Contract.DIVISOR(&_SURGE.CallOpts)
}

// DIVISOR is a free data retrieval call binding the contract method 0x3410fe6e.
//
// Solidity: function DIVISOR() view returns(uint256)
func (_SURGE *SURGECallerSession) DIVISOR() (*big.Int, error) {
	return _SURGE.Contract.DIVISOR(&_SURGE.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_SURGE *SURGECaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_SURGE *SURGESession) Balances(arg0 common.Address) (*big.Int, error) {
	return _SURGE.Contract.Balances(&_SURGE.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x6ebcf607.
//
// Solidity: function _balances(address ) view returns(uint256)
func (_SURGE *SURGECallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _SURGE.Contract.Balances(&_SURGE.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_SURGE *SURGECaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_SURGE *SURGESession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _SURGE.Contract.Allowance(&_SURGE.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_SURGE *SURGECallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _SURGE.Contract.Allowance(&_SURGE.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SURGE *SURGECaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SURGE *SURGESession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SURGE.Contract.BalanceOf(&_SURGE.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SURGE *SURGECallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SURGE.Contract.BalanceOf(&_SURGE.CallOpts, account)
}

// BuyMul is a free data retrieval call binding the contract method 0xa0e571a2.
//
// Solidity: function buyMul() view returns(uint256)
func (_SURGE *SURGECaller) BuyMul(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "buyMul")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyMul is a free data retrieval call binding the contract method 0xa0e571a2.
//
// Solidity: function buyMul() view returns(uint256)
func (_SURGE *SURGESession) BuyMul() (*big.Int, error) {
	return _SURGE.Contract.BuyMul(&_SURGE.CallOpts)
}

// BuyMul is a free data retrieval call binding the contract method 0xa0e571a2.
//
// Solidity: function buyMul() view returns(uint256)
func (_SURGE *SURGECallerSession) BuyMul() (*big.Int, error) {
	return _SURGE.Contract.BuyMul(&_SURGE.CallOpts)
}

// CalculatePrice is a free data retrieval call binding the contract method 0xd348b409.
//
// Solidity: function calculatePrice() view returns(uint256)
func (_SURGE *SURGECaller) CalculatePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "calculatePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePrice is a free data retrieval call binding the contract method 0xd348b409.
//
// Solidity: function calculatePrice() view returns(uint256)
func (_SURGE *SURGESession) CalculatePrice() (*big.Int, error) {
	return _SURGE.Contract.CalculatePrice(&_SURGE.CallOpts)
}

// CalculatePrice is a free data retrieval call binding the contract method 0xd348b409.
//
// Solidity: function calculatePrice() view returns(uint256)
func (_SURGE *SURGECallerSession) CalculatePrice() (*big.Int, error) {
	return _SURGE.Contract.CalculatePrice(&_SURGE.CallOpts)
}

// CandleStickData is a free data retrieval call binding the contract method 0xa8fb66b4.
//
// Solidity: function candleStickData(uint256 ) view returns(uint256 time, uint256 open, uint256 close, uint256 high, uint256 low)
func (_SURGE *SURGECaller) CandleStickData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Time  *big.Int
	Open  *big.Int
	Close *big.Int
	High  *big.Int
	Low   *big.Int
}, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "candleStickData", arg0)

	outstruct := new(struct {
		Time  *big.Int
		Open  *big.Int
		Close *big.Int
		High  *big.Int
		Low   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Time = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Open = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Close = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.High = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Low = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CandleStickData is a free data retrieval call binding the contract method 0xa8fb66b4.
//
// Solidity: function candleStickData(uint256 ) view returns(uint256 time, uint256 open, uint256 close, uint256 high, uint256 low)
func (_SURGE *SURGESession) CandleStickData(arg0 *big.Int) (struct {
	Time  *big.Int
	Open  *big.Int
	Close *big.Int
	High  *big.Int
	Low   *big.Int
}, error) {
	return _SURGE.Contract.CandleStickData(&_SURGE.CallOpts, arg0)
}

// CandleStickData is a free data retrieval call binding the contract method 0xa8fb66b4.
//
// Solidity: function candleStickData(uint256 ) view returns(uint256 time, uint256 open, uint256 close, uint256 high, uint256 low)
func (_SURGE *SURGECallerSession) CandleStickData(arg0 *big.Int) (struct {
	Time  *big.Int
	Open  *big.Int
	Close *big.Int
	High  *big.Int
	Low   *big.Int
}, error) {
	return _SURGE.Contract.CandleStickData(&_SURGE.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SURGE *SURGECaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SURGE *SURGESession) Decimals() (uint8, error) {
	return _SURGE.Contract.Decimals(&_SURGE.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SURGE *SURGECallerSession) Decimals() (uint8, error) {
	return _SURGE.Contract.Decimals(&_SURGE.CallOpts)
}

// GetBNBAmountOut is a free data retrieval call binding the contract method 0x1f1a73d5.
//
// Solidity: function getBNBAmountOut(uint256 amountIn) view returns(uint256)
func (_SURGE *SURGECaller) GetBNBAmountOut(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getBNBAmountOut", amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBNBAmountOut is a free data retrieval call binding the contract method 0x1f1a73d5.
//
// Solidity: function getBNBAmountOut(uint256 amountIn) view returns(uint256)
func (_SURGE *SURGESession) GetBNBAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _SURGE.Contract.GetBNBAmountOut(&_SURGE.CallOpts, amountIn)
}

// GetBNBAmountOut is a free data retrieval call binding the contract method 0x1f1a73d5.
//
// Solidity: function getBNBAmountOut(uint256 amountIn) view returns(uint256)
func (_SURGE *SURGECallerSession) GetBNBAmountOut(amountIn *big.Int) (*big.Int, error) {
	return _SURGE.Contract.GetBNBAmountOut(&_SURGE.CallOpts, amountIn)
}

// GetBNBPrice is a free data retrieval call binding the contract method 0xf2220c9e.
//
// Solidity: function getBNBPrice() view returns(uint256)
func (_SURGE *SURGECaller) GetBNBPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getBNBPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBNBPrice is a free data retrieval call binding the contract method 0xf2220c9e.
//
// Solidity: function getBNBPrice() view returns(uint256)
func (_SURGE *SURGESession) GetBNBPrice() (*big.Int, error) {
	return _SURGE.Contract.GetBNBPrice(&_SURGE.CallOpts)
}

// GetBNBPrice is a free data retrieval call binding the contract method 0xf2220c9e.
//
// Solidity: function getBNBPrice() view returns(uint256)
func (_SURGE *SURGECallerSession) GetBNBPrice() (*big.Int, error) {
	return _SURGE.Contract.GetBNBPrice(&_SURGE.CallOpts)
}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_SURGE *SURGECaller) GetCirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getCirculatingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_SURGE *SURGESession) GetCirculatingSupply() (*big.Int, error) {
	return _SURGE.Contract.GetCirculatingSupply(&_SURGE.CallOpts)
}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_SURGE *SURGECallerSession) GetCirculatingSupply() (*big.Int, error) {
	return _SURGE.Contract.GetCirculatingSupply(&_SURGE.CallOpts)
}

// GetLiquidity is a free data retrieval call binding the contract method 0x0910a510.
//
// Solidity: function getLiquidity() view returns(uint256)
func (_SURGE *SURGECaller) GetLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidity is a free data retrieval call binding the contract method 0x0910a510.
//
// Solidity: function getLiquidity() view returns(uint256)
func (_SURGE *SURGESession) GetLiquidity() (*big.Int, error) {
	return _SURGE.Contract.GetLiquidity(&_SURGE.CallOpts)
}

// GetLiquidity is a free data retrieval call binding the contract method 0x0910a510.
//
// Solidity: function getLiquidity() view returns(uint256)
func (_SURGE *SURGECallerSession) GetLiquidity() (*big.Int, error) {
	return _SURGE.Contract.GetLiquidity(&_SURGE.CallOpts)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x90825c28.
//
// Solidity: function getMarketCap() view returns(uint256)
func (_SURGE *SURGECaller) GetMarketCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getMarketCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketCap is a free data retrieval call binding the contract method 0x90825c28.
//
// Solidity: function getMarketCap() view returns(uint256)
func (_SURGE *SURGESession) GetMarketCap() (*big.Int, error) {
	return _SURGE.Contract.GetMarketCap(&_SURGE.CallOpts)
}

// GetMarketCap is a free data retrieval call binding the contract method 0x90825c28.
//
// Solidity: function getMarketCap() view returns(uint256)
func (_SURGE *SURGECallerSession) GetMarketCap() (*big.Int, error) {
	return _SURGE.Contract.GetMarketCap(&_SURGE.CallOpts)
}

// GetTokenAmountOut is a free data retrieval call binding the contract method 0x71073b38.
//
// Solidity: function getTokenAmountOut(uint256 amountBNBIn) view returns(uint256)
func (_SURGE *SURGECaller) GetTokenAmountOut(opts *bind.CallOpts, amountBNBIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getTokenAmountOut", amountBNBIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountOut is a free data retrieval call binding the contract method 0x71073b38.
//
// Solidity: function getTokenAmountOut(uint256 amountBNBIn) view returns(uint256)
func (_SURGE *SURGESession) GetTokenAmountOut(amountBNBIn *big.Int) (*big.Int, error) {
	return _SURGE.Contract.GetTokenAmountOut(&_SURGE.CallOpts, amountBNBIn)
}

// GetTokenAmountOut is a free data retrieval call binding the contract method 0x71073b38.
//
// Solidity: function getTokenAmountOut(uint256 amountBNBIn) view returns(uint256)
func (_SURGE *SURGECallerSession) GetTokenAmountOut(amountBNBIn *big.Int) (*big.Int, error) {
	return _SURGE.Contract.GetTokenAmountOut(&_SURGE.CallOpts, amountBNBIn)
}

// GetValueOfHoldings is a free data retrieval call binding the contract method 0x1f02a29c.
//
// Solidity: function getValueOfHoldings(address holder) view returns(uint256)
func (_SURGE *SURGECaller) GetValueOfHoldings(opts *bind.CallOpts, holder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "getValueOfHoldings", holder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueOfHoldings is a free data retrieval call binding the contract method 0x1f02a29c.
//
// Solidity: function getValueOfHoldings(address holder) view returns(uint256)
func (_SURGE *SURGESession) GetValueOfHoldings(holder common.Address) (*big.Int, error) {
	return _SURGE.Contract.GetValueOfHoldings(&_SURGE.CallOpts, holder)
}

// GetValueOfHoldings is a free data retrieval call binding the contract method 0x1f02a29c.
//
// Solidity: function getValueOfHoldings(address holder) view returns(uint256)
func (_SURGE *SURGECallerSession) GetValueOfHoldings(holder common.Address) (*big.Int, error) {
	return _SURGE.Contract.GetValueOfHoldings(&_SURGE.CallOpts, holder)
}

// IndVol is a free data retrieval call binding the contract method 0x9cbd09ac.
//
// Solidity: function indVol(address ) view returns(uint256)
func (_SURGE *SURGECaller) IndVol(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "indVol", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndVol is a free data retrieval call binding the contract method 0x9cbd09ac.
//
// Solidity: function indVol(address ) view returns(uint256)
func (_SURGE *SURGESession) IndVol(arg0 common.Address) (*big.Int, error) {
	return _SURGE.Contract.IndVol(&_SURGE.CallOpts, arg0)
}

// IndVol is a free data retrieval call binding the contract method 0x9cbd09ac.
//
// Solidity: function indVol(address ) view returns(uint256)
func (_SURGE *SURGECallerSession) IndVol(arg0 common.Address) (*big.Int, error) {
	return _SURGE.Contract.IndVol(&_SURGE.CallOpts, arg0)
}

// IsFeeExempt is a free data retrieval call binding the contract method 0x3f4218e0.
//
// Solidity: function isFeeExempt(address ) view returns(bool)
func (_SURGE *SURGECaller) IsFeeExempt(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "isFeeExempt", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeeExempt is a free data retrieval call binding the contract method 0x3f4218e0.
//
// Solidity: function isFeeExempt(address ) view returns(bool)
func (_SURGE *SURGESession) IsFeeExempt(arg0 common.Address) (bool, error) {
	return _SURGE.Contract.IsFeeExempt(&_SURGE.CallOpts, arg0)
}

// IsFeeExempt is a free data retrieval call binding the contract method 0x3f4218e0.
//
// Solidity: function isFeeExempt(address ) view returns(bool)
func (_SURGE *SURGECallerSession) IsFeeExempt(arg0 common.Address) (bool, error) {
	return _SURGE.Contract.IsFeeExempt(&_SURGE.CallOpts, arg0)
}

// IsTxLimitExempt is a free data retrieval call binding the contract method 0x8b42507f.
//
// Solidity: function isTxLimitExempt(address ) view returns(bool)
func (_SURGE *SURGECaller) IsTxLimitExempt(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "isTxLimitExempt", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTxLimitExempt is a free data retrieval call binding the contract method 0x8b42507f.
//
// Solidity: function isTxLimitExempt(address ) view returns(bool)
func (_SURGE *SURGESession) IsTxLimitExempt(arg0 common.Address) (bool, error) {
	return _SURGE.Contract.IsTxLimitExempt(&_SURGE.CallOpts, arg0)
}

// IsTxLimitExempt is a free data retrieval call binding the contract method 0x8b42507f.
//
// Solidity: function isTxLimitExempt(address ) view returns(bool)
func (_SURGE *SURGECallerSession) IsTxLimitExempt(arg0 common.Address) (bool, error) {
	return _SURGE.Contract.IsTxLimitExempt(&_SURGE.CallOpts, arg0)
}

// LiqConst is a free data retrieval call binding the contract method 0xe8f620b1.
//
// Solidity: function liqConst() view returns(uint256)
func (_SURGE *SURGECaller) LiqConst(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "liqConst")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiqConst is a free data retrieval call binding the contract method 0xe8f620b1.
//
// Solidity: function liqConst() view returns(uint256)
func (_SURGE *SURGESession) LiqConst() (*big.Int, error) {
	return _SURGE.Contract.LiqConst(&_SURGE.CallOpts)
}

// LiqConst is a free data retrieval call binding the contract method 0xe8f620b1.
//
// Solidity: function liqConst() view returns(uint256)
func (_SURGE *SURGECallerSession) LiqConst() (*big.Int, error) {
	return _SURGE.Contract.LiqConst(&_SURGE.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint256)
func (_SURGE *SURGECaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint256)
func (_SURGE *SURGESession) Liquidity() (*big.Int, error) {
	return _SURGE.Contract.Liquidity(&_SURGE.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint256)
func (_SURGE *SURGECallerSession) Liquidity() (*big.Int, error) {
	return _SURGE.Contract.Liquidity(&_SURGE.CallOpts)
}

// MaxBag is a free data retrieval call binding the contract method 0xe0cfcc95.
//
// Solidity: function maxBag() view returns(uint256)
func (_SURGE *SURGECaller) MaxBag(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "maxBag")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBag is a free data retrieval call binding the contract method 0xe0cfcc95.
//
// Solidity: function maxBag() view returns(uint256)
func (_SURGE *SURGESession) MaxBag() (*big.Int, error) {
	return _SURGE.Contract.MaxBag(&_SURGE.CallOpts)
}

// MaxBag is a free data retrieval call binding the contract method 0xe0cfcc95.
//
// Solidity: function maxBag() view returns(uint256)
func (_SURGE *SURGECallerSession) MaxBag() (*big.Int, error) {
	return _SURGE.Contract.MaxBag(&_SURGE.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SURGE *SURGECaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SURGE *SURGESession) Name() (string, error) {
	return _SURGE.Contract.Name(&_SURGE.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SURGE *SURGECallerSession) Name() (string, error) {
	return _SURGE.Contract.Name(&_SURGE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SURGE *SURGECaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SURGE *SURGESession) Owner() (common.Address, error) {
	return _SURGE.Contract.Owner(&_SURGE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SURGE *SURGECallerSession) Owner() (common.Address, error) {
	return _SURGE.Contract.Owner(&_SURGE.CallOpts)
}

// SellMul is a free data retrieval call binding the contract method 0x9d755026.
//
// Solidity: function sellMul() view returns(uint256)
func (_SURGE *SURGECaller) SellMul(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "sellMul")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellMul is a free data retrieval call binding the contract method 0x9d755026.
//
// Solidity: function sellMul() view returns(uint256)
func (_SURGE *SURGESession) SellMul() (*big.Int, error) {
	return _SURGE.Contract.SellMul(&_SURGE.CallOpts)
}

// SellMul is a free data retrieval call binding the contract method 0x9d755026.
//
// Solidity: function sellMul() view returns(uint256)
func (_SURGE *SURGECallerSession) SellMul() (*big.Int, error) {
	return _SURGE.Contract.SellMul(&_SURGE.CallOpts)
}

// ShareDIVISOR is a free data retrieval call binding the contract method 0x26763a6c.
//
// Solidity: function shareDIVISOR() view returns(uint256)
func (_SURGE *SURGECaller) ShareDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "shareDIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareDIVISOR is a free data retrieval call binding the contract method 0x26763a6c.
//
// Solidity: function shareDIVISOR() view returns(uint256)
func (_SURGE *SURGESession) ShareDIVISOR() (*big.Int, error) {
	return _SURGE.Contract.ShareDIVISOR(&_SURGE.CallOpts)
}

// ShareDIVISOR is a free data retrieval call binding the contract method 0x26763a6c.
//
// Solidity: function shareDIVISOR() view returns(uint256)
func (_SURGE *SURGECallerSession) ShareDIVISOR() (*big.Int, error) {
	return _SURGE.Contract.ShareDIVISOR(&_SURGE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SURGE *SURGECaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SURGE *SURGESession) Symbol() (string, error) {
	return _SURGE.Contract.Symbol(&_SURGE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SURGE *SURGECallerSession) Symbol() (string, error) {
	return _SURGE.Contract.Symbol(&_SURGE.CallOpts)
}

// TVol is a free data retrieval call binding the contract method 0x49ce234c.
//
// Solidity: function tVol(uint256 ) view returns(uint256)
func (_SURGE *SURGECaller) TVol(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "tVol", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TVol is a free data retrieval call binding the contract method 0x49ce234c.
//
// Solidity: function tVol(uint256 ) view returns(uint256)
func (_SURGE *SURGESession) TVol(arg0 *big.Int) (*big.Int, error) {
	return _SURGE.Contract.TVol(&_SURGE.CallOpts, arg0)
}

// TVol is a free data retrieval call binding the contract method 0x49ce234c.
//
// Solidity: function tVol(uint256 ) view returns(uint256)
func (_SURGE *SURGECallerSession) TVol(arg0 *big.Int) (*big.Int, error) {
	return _SURGE.Contract.TVol(&_SURGE.CallOpts, arg0)
}

// TaxBalance is a free data retrieval call binding the contract method 0xaa98e163.
//
// Solidity: function taxBalance() view returns(uint256)
func (_SURGE *SURGECaller) TaxBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "taxBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBalance is a free data retrieval call binding the contract method 0xaa98e163.
//
// Solidity: function taxBalance() view returns(uint256)
func (_SURGE *SURGESession) TaxBalance() (*big.Int, error) {
	return _SURGE.Contract.TaxBalance(&_SURGE.CallOpts)
}

// TaxBalance is a free data retrieval call binding the contract method 0xaa98e163.
//
// Solidity: function taxBalance() view returns(uint256)
func (_SURGE *SURGECallerSession) TaxBalance() (*big.Int, error) {
	return _SURGE.Contract.TaxBalance(&_SURGE.CallOpts)
}

// TeamShare is a free data retrieval call binding the contract method 0xea6ef2fe.
//
// Solidity: function teamShare() view returns(uint256)
func (_SURGE *SURGECaller) TeamShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "teamShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamShare is a free data retrieval call binding the contract method 0xea6ef2fe.
//
// Solidity: function teamShare() view returns(uint256)
func (_SURGE *SURGESession) TeamShare() (*big.Int, error) {
	return _SURGE.Contract.TeamShare(&_SURGE.CallOpts)
}

// TeamShare is a free data retrieval call binding the contract method 0xea6ef2fe.
//
// Solidity: function teamShare() view returns(uint256)
func (_SURGE *SURGECallerSession) TeamShare() (*big.Int, error) {
	return _SURGE.Contract.TeamShare(&_SURGE.CallOpts)
}

// TeamWallet is a free data retrieval call binding the contract method 0x59927044.
//
// Solidity: function teamWallet() view returns(address)
func (_SURGE *SURGECaller) TeamWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "teamWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeamWallet is a free data retrieval call binding the contract method 0x59927044.
//
// Solidity: function teamWallet() view returns(address)
func (_SURGE *SURGESession) TeamWallet() (common.Address, error) {
	return _SURGE.Contract.TeamWallet(&_SURGE.CallOpts)
}

// TeamWallet is a free data retrieval call binding the contract method 0x59927044.
//
// Solidity: function teamWallet() view returns(address)
func (_SURGE *SURGECallerSession) TeamWallet() (common.Address, error) {
	return _SURGE.Contract.TeamWallet(&_SURGE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGECaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGESession) TotalSupply() (*big.Int, error) {
	return _SURGE.Contract.TotalSupply(&_SURGE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGECallerSession) TotalSupply() (*big.Int, error) {
	return _SURGE.Contract.TotalSupply(&_SURGE.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGECaller) TotalSupply0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "totalSupply0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGESession) TotalSupply0() (*big.Int, error) {
	return _SURGE.Contract.TotalSupply0(&_SURGE.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SURGE *SURGECallerSession) TotalSupply0() (*big.Int, error) {
	return _SURGE.Contract.TotalSupply0(&_SURGE.CallOpts)
}

// TotalTx is a free data retrieval call binding the contract method 0x7220cf39.
//
// Solidity: function totalTx() view returns(uint256)
func (_SURGE *SURGECaller) TotalTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "totalTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTx is a free data retrieval call binding the contract method 0x7220cf39.
//
// Solidity: function totalTx() view returns(uint256)
func (_SURGE *SURGESession) TotalTx() (*big.Int, error) {
	return _SURGE.Contract.TotalTx(&_SURGE.CallOpts)
}

// TotalTx is a free data retrieval call binding the contract method 0x7220cf39.
//
// Solidity: function totalTx() view returns(uint256)
func (_SURGE *SURGECallerSession) TotalTx() (*big.Int, error) {
	return _SURGE.Contract.TotalTx(&_SURGE.CallOpts)
}

// TotalVolume is a free data retrieval call binding the contract method 0x5f81a57c.
//
// Solidity: function totalVolume() view returns(uint256)
func (_SURGE *SURGECaller) TotalVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "totalVolume")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVolume is a free data retrieval call binding the contract method 0x5f81a57c.
//
// Solidity: function totalVolume() view returns(uint256)
func (_SURGE *SURGESession) TotalVolume() (*big.Int, error) {
	return _SURGE.Contract.TotalVolume(&_SURGE.CallOpts)
}

// TotalVolume is a free data retrieval call binding the contract method 0x5f81a57c.
//
// Solidity: function totalVolume() view returns(uint256)
func (_SURGE *SURGECallerSession) TotalVolume() (*big.Int, error) {
	return _SURGE.Contract.TotalVolume(&_SURGE.CallOpts)
}

// TradeOpenTime is a free data retrieval call binding the contract method 0x12fbbbe4.
//
// Solidity: function tradeOpenTime() view returns(uint256)
func (_SURGE *SURGECaller) TradeOpenTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "tradeOpenTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradeOpenTime is a free data retrieval call binding the contract method 0x12fbbbe4.
//
// Solidity: function tradeOpenTime() view returns(uint256)
func (_SURGE *SURGESession) TradeOpenTime() (*big.Int, error) {
	return _SURGE.Contract.TradeOpenTime(&_SURGE.CallOpts)
}

// TradeOpenTime is a free data retrieval call binding the contract method 0x12fbbbe4.
//
// Solidity: function tradeOpenTime() view returns(uint256)
func (_SURGE *SURGECallerSession) TradeOpenTime() (*big.Int, error) {
	return _SURGE.Contract.TradeOpenTime(&_SURGE.CallOpts)
}

// TreasuryShare is a free data retrieval call binding the contract method 0x7796ff37.
//
// Solidity: function treasuryShare() view returns(uint256)
func (_SURGE *SURGECaller) TreasuryShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "treasuryShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryShare is a free data retrieval call binding the contract method 0x7796ff37.
//
// Solidity: function treasuryShare() view returns(uint256)
func (_SURGE *SURGESession) TreasuryShare() (*big.Int, error) {
	return _SURGE.Contract.TreasuryShare(&_SURGE.CallOpts)
}

// TreasuryShare is a free data retrieval call binding the contract method 0x7796ff37.
//
// Solidity: function treasuryShare() view returns(uint256)
func (_SURGE *SURGECallerSession) TreasuryShare() (*big.Int, error) {
	return _SURGE.Contract.TreasuryShare(&_SURGE.CallOpts)
}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_SURGE *SURGECaller) TreasuryWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "treasuryWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_SURGE *SURGESession) TreasuryWallet() (common.Address, error) {
	return _SURGE.Contract.TreasuryWallet(&_SURGE.CallOpts)
}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_SURGE *SURGECallerSession) TreasuryWallet() (common.Address, error) {
	return _SURGE.Contract.TreasuryWallet(&_SURGE.CallOpts)
}

// TxTimeStamp is a free data retrieval call binding the contract method 0xf7a62fe9.
//
// Solidity: function txTimeStamp(uint256 ) view returns(uint256)
func (_SURGE *SURGECaller) TxTimeStamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SURGE.contract.Call(opts, &out, "txTimeStamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxTimeStamp is a free data retrieval call binding the contract method 0xf7a62fe9.
//
// Solidity: function txTimeStamp(uint256 ) view returns(uint256)
func (_SURGE *SURGESession) TxTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _SURGE.Contract.TxTimeStamp(&_SURGE.CallOpts, arg0)
}

// TxTimeStamp is a free data retrieval call binding the contract method 0xf7a62fe9.
//
// Solidity: function txTimeStamp(uint256 ) view returns(uint256)
func (_SURGE *SURGECallerSession) TxTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _SURGE.Contract.TxTimeStamp(&_SURGE.CallOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0x8f0d3b8b.
//
// Solidity: function _buy(uint256 minTokenOut, uint256 deadline) payable returns(bool)
func (_SURGE *SURGETransactor) Buy(opts *bind.TransactOpts, minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "_buy", minTokenOut, deadline)
}

// Buy is a paid mutator transaction binding the contract method 0x8f0d3b8b.
//
// Solidity: function _buy(uint256 minTokenOut, uint256 deadline) payable returns(bool)
func (_SURGE *SURGESession) Buy(minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Buy(&_SURGE.TransactOpts, minTokenOut, deadline)
}

// Buy is a paid mutator transaction binding the contract method 0x8f0d3b8b.
//
// Solidity: function _buy(uint256 minTokenOut, uint256 deadline) payable returns(bool)
func (_SURGE *SURGETransactorSession) Buy(minTokenOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Buy(&_SURGE.TransactOpts, minTokenOut, deadline)
}

// Sell is a paid mutator transaction binding the contract method 0xb37659a4.
//
// Solidity: function _sell(uint256 tokenAmount, uint256 deadline, uint256 minBNBOut) payable returns(bool)
func (_SURGE *SURGETransactor) Sell(opts *bind.TransactOpts, tokenAmount *big.Int, deadline *big.Int, minBNBOut *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "_sell", tokenAmount, deadline, minBNBOut)
}

// Sell is a paid mutator transaction binding the contract method 0xb37659a4.
//
// Solidity: function _sell(uint256 tokenAmount, uint256 deadline, uint256 minBNBOut) payable returns(bool)
func (_SURGE *SURGESession) Sell(tokenAmount *big.Int, deadline *big.Int, minBNBOut *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Sell(&_SURGE.TransactOpts, tokenAmount, deadline, minBNBOut)
}

// Sell is a paid mutator transaction binding the contract method 0xb37659a4.
//
// Solidity: function _sell(uint256 tokenAmount, uint256 deadline, uint256 minBNBOut) payable returns(bool)
func (_SURGE *SURGETransactorSession) Sell(tokenAmount *big.Int, deadline *big.Int, minBNBOut *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Sell(&_SURGE.TransactOpts, tokenAmount, deadline, minBNBOut)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() payable returns()
func (_SURGE *SURGETransactor) AddLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "addLiquidity")
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() payable returns()
func (_SURGE *SURGESession) AddLiquidity() (*types.Transaction, error) {
	return _SURGE.Contract.AddLiquidity(&_SURGE.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() payable returns()
func (_SURGE *SURGETransactorSession) AddLiquidity() (*types.Transaction, error) {
	return _SURGE.Contract.AddLiquidity(&_SURGE.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SURGE *SURGETransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SURGE *SURGESession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Approve(&_SURGE.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SURGE *SURGETransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Approve(&_SURGE.TransactOpts, spender, amount)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_SURGE *SURGETransactor) ApproveMax(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "approveMax", spender)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_SURGE *SURGESession) ApproveMax(spender common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ApproveMax(&_SURGE.TransactOpts, spender)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_SURGE *SURGETransactorSession) ApproveMax(spender common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ApproveMax(&_SURGE.TransactOpts, spender)
}

// ChangeFeeReceivers is a paid mutator transaction binding the contract method 0x1f2c80f1.
//
// Solidity: function changeFeeReceivers(address newTeamWallet, address newTreasuryWallet) returns()
func (_SURGE *SURGETransactor) ChangeFeeReceivers(opts *bind.TransactOpts, newTeamWallet common.Address, newTreasuryWallet common.Address) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeFeeReceivers", newTeamWallet, newTreasuryWallet)
}

// ChangeFeeReceivers is a paid mutator transaction binding the contract method 0x1f2c80f1.
//
// Solidity: function changeFeeReceivers(address newTeamWallet, address newTreasuryWallet) returns()
func (_SURGE *SURGESession) ChangeFeeReceivers(newTeamWallet common.Address, newTreasuryWallet common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeFeeReceivers(&_SURGE.TransactOpts, newTeamWallet, newTreasuryWallet)
}

// ChangeFeeReceivers is a paid mutator transaction binding the contract method 0x1f2c80f1.
//
// Solidity: function changeFeeReceivers(address newTeamWallet, address newTreasuryWallet) returns()
func (_SURGE *SURGETransactorSession) ChangeFeeReceivers(newTeamWallet common.Address, newTreasuryWallet common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeFeeReceivers(&_SURGE.TransactOpts, newTeamWallet, newTreasuryWallet)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x21ecff5b.
//
// Solidity: function changeFees(uint256 newbuyMul, uint256 newsellMul) returns()
func (_SURGE *SURGETransactor) ChangeFees(opts *bind.TransactOpts, newbuyMul *big.Int, newsellMul *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeFees", newbuyMul, newsellMul)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x21ecff5b.
//
// Solidity: function changeFees(uint256 newbuyMul, uint256 newsellMul) returns()
func (_SURGE *SURGESession) ChangeFees(newbuyMul *big.Int, newsellMul *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeFees(&_SURGE.TransactOpts, newbuyMul, newsellMul)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x21ecff5b.
//
// Solidity: function changeFees(uint256 newbuyMul, uint256 newsellMul) returns()
func (_SURGE *SURGETransactorSession) ChangeFees(newbuyMul *big.Int, newsellMul *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeFees(&_SURGE.TransactOpts, newbuyMul, newsellMul)
}

// ChangeIsFeeExempt is a paid mutator transaction binding the contract method 0xa3a2e89e.
//
// Solidity: function changeIsFeeExempt(address holder, bool exempt) returns()
func (_SURGE *SURGETransactor) ChangeIsFeeExempt(opts *bind.TransactOpts, holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeIsFeeExempt", holder, exempt)
}

// ChangeIsFeeExempt is a paid mutator transaction binding the contract method 0xa3a2e89e.
//
// Solidity: function changeIsFeeExempt(address holder, bool exempt) returns()
func (_SURGE *SURGESession) ChangeIsFeeExempt(holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeIsFeeExempt(&_SURGE.TransactOpts, holder, exempt)
}

// ChangeIsFeeExempt is a paid mutator transaction binding the contract method 0xa3a2e89e.
//
// Solidity: function changeIsFeeExempt(address holder, bool exempt) returns()
func (_SURGE *SURGETransactorSession) ChangeIsFeeExempt(holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeIsFeeExempt(&_SURGE.TransactOpts, holder, exempt)
}

// ChangeIsTxLimitExempt is a paid mutator transaction binding the contract method 0xfabe6283.
//
// Solidity: function changeIsTxLimitExempt(address holder, bool exempt) returns()
func (_SURGE *SURGETransactor) ChangeIsTxLimitExempt(opts *bind.TransactOpts, holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeIsTxLimitExempt", holder, exempt)
}

// ChangeIsTxLimitExempt is a paid mutator transaction binding the contract method 0xfabe6283.
//
// Solidity: function changeIsTxLimitExempt(address holder, bool exempt) returns()
func (_SURGE *SURGESession) ChangeIsTxLimitExempt(holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeIsTxLimitExempt(&_SURGE.TransactOpts, holder, exempt)
}

// ChangeIsTxLimitExempt is a paid mutator transaction binding the contract method 0xfabe6283.
//
// Solidity: function changeIsTxLimitExempt(address holder, bool exempt) returns()
func (_SURGE *SURGETransactorSession) ChangeIsTxLimitExempt(holder common.Address, exempt bool) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeIsTxLimitExempt(&_SURGE.TransactOpts, holder, exempt)
}

// ChangeStablePair is a paid mutator transaction binding the contract method 0x1ab6ab24.
//
// Solidity: function changeStablePair(address newStablePair, address newStableAddress) returns()
func (_SURGE *SURGETransactor) ChangeStablePair(opts *bind.TransactOpts, newStablePair common.Address, newStableAddress common.Address) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeStablePair", newStablePair, newStableAddress)
}

// ChangeStablePair is a paid mutator transaction binding the contract method 0x1ab6ab24.
//
// Solidity: function changeStablePair(address newStablePair, address newStableAddress) returns()
func (_SURGE *SURGESession) ChangeStablePair(newStablePair common.Address, newStableAddress common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeStablePair(&_SURGE.TransactOpts, newStablePair, newStableAddress)
}

// ChangeStablePair is a paid mutator transaction binding the contract method 0x1ab6ab24.
//
// Solidity: function changeStablePair(address newStablePair, address newStableAddress) returns()
func (_SURGE *SURGETransactorSession) ChangeStablePair(newStablePair common.Address, newStableAddress common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeStablePair(&_SURGE.TransactOpts, newStablePair, newStableAddress)
}

// ChangeTaxDistribution is a paid mutator transaction binding the contract method 0xf4034e4b.
//
// Solidity: function changeTaxDistribution(uint256 newteamShare, uint256 newtreasuryShare) returns()
func (_SURGE *SURGETransactor) ChangeTaxDistribution(opts *bind.TransactOpts, newteamShare *big.Int, newtreasuryShare *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeTaxDistribution", newteamShare, newtreasuryShare)
}

// ChangeTaxDistribution is a paid mutator transaction binding the contract method 0xf4034e4b.
//
// Solidity: function changeTaxDistribution(uint256 newteamShare, uint256 newtreasuryShare) returns()
func (_SURGE *SURGESession) ChangeTaxDistribution(newteamShare *big.Int, newtreasuryShare *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeTaxDistribution(&_SURGE.TransactOpts, newteamShare, newtreasuryShare)
}

// ChangeTaxDistribution is a paid mutator transaction binding the contract method 0xf4034e4b.
//
// Solidity: function changeTaxDistribution(uint256 newteamShare, uint256 newtreasuryShare) returns()
func (_SURGE *SURGETransactorSession) ChangeTaxDistribution(newteamShare *big.Int, newtreasuryShare *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeTaxDistribution(&_SURGE.TransactOpts, newteamShare, newtreasuryShare)
}

// ChangeWalletLimit is a paid mutator transaction binding the contract method 0x7db1342c.
//
// Solidity: function changeWalletLimit(uint256 newLimit) returns()
func (_SURGE *SURGETransactor) ChangeWalletLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "changeWalletLimit", newLimit)
}

// ChangeWalletLimit is a paid mutator transaction binding the contract method 0x7db1342c.
//
// Solidity: function changeWalletLimit(uint256 newLimit) returns()
func (_SURGE *SURGESession) ChangeWalletLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeWalletLimit(&_SURGE.TransactOpts, newLimit)
}

// ChangeWalletLimit is a paid mutator transaction binding the contract method 0x7db1342c.
//
// Solidity: function changeWalletLimit(uint256 newLimit) returns()
func (_SURGE *SURGETransactorSession) ChangeWalletLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.ChangeWalletLimit(&_SURGE.TransactOpts, newLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SURGE *SURGETransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SURGE *SURGESession) RenounceOwnership() (*types.Transaction, error) {
	return _SURGE.Contract.RenounceOwnership(&_SURGE.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SURGE *SURGETransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SURGE.Contract.RenounceOwnership(&_SURGE.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGETransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGESession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Transfer(&_SURGE.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGETransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.Transfer(&_SURGE.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGETransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGESession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.TransferFrom(&_SURGE.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SURGE *SURGETransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SURGE.Contract.TransferFrom(&_SURGE.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SURGE *SURGETransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SURGE *SURGESession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.TransferOwnership(&_SURGE.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SURGE *SURGETransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SURGE.Contract.TransferOwnership(&_SURGE.TransactOpts, newOwner)
}

// WithdrawTaxBalance is a paid mutator transaction binding the contract method 0x1d6bdb5b.
//
// Solidity: function withdrawTaxBalance() payable returns()
func (_SURGE *SURGETransactor) WithdrawTaxBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SURGE.contract.Transact(opts, "withdrawTaxBalance")
}

// WithdrawTaxBalance is a paid mutator transaction binding the contract method 0x1d6bdb5b.
//
// Solidity: function withdrawTaxBalance() payable returns()
func (_SURGE *SURGESession) WithdrawTaxBalance() (*types.Transaction, error) {
	return _SURGE.Contract.WithdrawTaxBalance(&_SURGE.TransactOpts)
}

// WithdrawTaxBalance is a paid mutator transaction binding the contract method 0x1d6bdb5b.
//
// Solidity: function withdrawTaxBalance() payable returns()
func (_SURGE *SURGETransactorSession) WithdrawTaxBalance() (*types.Transaction, error) {
	return _SURGE.Contract.WithdrawTaxBalance(&_SURGE.TransactOpts)
}

// SURGEApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SURGE contract.
type SURGEApprovalIterator struct {
	Event *SURGEApproval // Event containing the contract specifics and raw log

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
func (it *SURGEApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SURGEApproval)
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
		it.Event = new(SURGEApproval)
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
func (it *SURGEApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SURGEApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SURGEApproval represents a Approval event raised by the SURGE contract.
type SURGEApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SURGE *SURGEFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SURGEApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SURGE.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SURGEApprovalIterator{contract: _SURGE.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SURGE *SURGEFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SURGEApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SURGE.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SURGEApproval)
				if err := _SURGE.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SURGE *SURGEFilterer) ParseApproval(log types.Log) (*SURGEApproval, error) {
	event := new(SURGEApproval)
	if err := _SURGE.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SURGEBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the SURGE contract.
type SURGEBoughtIterator struct {
	Event *SURGEBought // Event containing the contract specifics and raw log

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
func (it *SURGEBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SURGEBought)
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
		it.Event = new(SURGEBought)
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
func (it *SURGEBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SURGEBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SURGEBought represents a Bought event raised by the SURGE contract.
type SURGEBought struct {
	From      common.Address
	To        common.Address
	Tokens    *big.Int
	Beans     *big.Int
	DollarBuy *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarBuy)
func (_SURGE *SURGEFilterer) FilterBought(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SURGEBoughtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.FilterLogs(opts, "Bought", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SURGEBoughtIterator{contract: _SURGE.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarBuy)
func (_SURGE *SURGEFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *SURGEBought, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.WatchLogs(opts, "Bought", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SURGEBought)
				if err := _SURGE.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x7ce543d1780f3bdc3dac42da06c95da802653cd1b212b8d74ec3e3c33ad7095c.
//
// Solidity: event Bought(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarBuy)
func (_SURGE *SURGEFilterer) ParseBought(log types.Log) (*SURGEBought, error) {
	event := new(SURGEBought)
	if err := _SURGE.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SURGEOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SURGE contract.
type SURGEOwnershipTransferredIterator struct {
	Event *SURGEOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SURGEOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SURGEOwnershipTransferred)
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
		it.Event = new(SURGEOwnershipTransferred)
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
func (it *SURGEOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SURGEOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SURGEOwnershipTransferred represents a OwnershipTransferred event raised by the SURGE contract.
type SURGEOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SURGE *SURGEFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SURGEOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SURGE.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SURGEOwnershipTransferredIterator{contract: _SURGE.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SURGE *SURGEFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SURGEOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SURGE.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SURGEOwnershipTransferred)
				if err := _SURGE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SURGE *SURGEFilterer) ParseOwnershipTransferred(log types.Log) (*SURGEOwnershipTransferred, error) {
	event := new(SURGEOwnershipTransferred)
	if err := _SURGE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SURGESoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the SURGE contract.
type SURGESoldIterator struct {
	Event *SURGESold // Event containing the contract specifics and raw log

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
func (it *SURGESoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SURGESold)
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
		it.Event = new(SURGESold)
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
func (it *SURGESoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SURGESoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SURGESold represents a Sold event raised by the SURGE contract.
type SURGESold struct {
	From       common.Address
	To         common.Address
	Tokens     *big.Int
	Beans      *big.Int
	DollarSell *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0x9be8a5ca22b7e6e81f04b5879f0248227bb770114291bd47dfaee4c3a82ad60e.
//
// Solidity: event Sold(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarSell)
func (_SURGE *SURGEFilterer) FilterSold(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SURGESoldIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.FilterLogs(opts, "Sold", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SURGESoldIterator{contract: _SURGE.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0x9be8a5ca22b7e6e81f04b5879f0248227bb770114291bd47dfaee4c3a82ad60e.
//
// Solidity: event Sold(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarSell)
func (_SURGE *SURGEFilterer) WatchSold(opts *bind.WatchOpts, sink chan<- *SURGESold, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.WatchLogs(opts, "Sold", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SURGESold)
				if err := _SURGE.contract.UnpackLog(event, "Sold", log); err != nil {
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

// ParseSold is a log parse operation binding the contract event 0x9be8a5ca22b7e6e81f04b5879f0248227bb770114291bd47dfaee4c3a82ad60e.
//
// Solidity: event Sold(address indexed from, address indexed to, uint256 tokens, uint256 beans, uint256 dollarSell)
func (_SURGE *SURGEFilterer) ParseSold(log types.Log) (*SURGESold, error) {
	event := new(SURGESold)
	if err := _SURGE.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SURGETransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SURGE contract.
type SURGETransferIterator struct {
	Event *SURGETransfer // Event containing the contract specifics and raw log

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
func (it *SURGETransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SURGETransfer)
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
		it.Event = new(SURGETransfer)
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
func (it *SURGETransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SURGETransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SURGETransfer represents a Transfer event raised by the SURGE contract.
type SURGETransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SURGE *SURGEFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SURGETransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SURGETransferIterator{contract: _SURGE.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SURGE *SURGEFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SURGETransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SURGE.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SURGETransfer)
				if err := _SURGE.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SURGE *SURGEFilterer) ParseTransfer(log types.Log) (*SURGETransfer, error) {
	event := new(SURGETransfer)
	if err := _SURGE.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
