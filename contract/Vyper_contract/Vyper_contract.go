// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vyper_contract

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	BuyResult    *big.Int
	TokenBalance *big.Int
	SellResult   *big.Int
	BuyCost      *big.Int
	SellCost     *big.Int
	Amounts      []*big.Int
}

// VyperContractMetaData contains all meta data concerning the VyperContract contract.
var VyperContractMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"honeyCheck\",\"inputs\":[{\"name\":\"_buyToken\",\"type\":\"address\"},{\"name\":\"_targetTokenAddress\",\"type\":\"address\"},{\"name\":\"idexRouterAddres\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"buyResult\",\"type\":\"uint256\"},{\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"name\":\"sellResult\",\"type\":\"uint256\"},{\"name\":\"buyCost\",\"type\":\"uint256\"},{\"name\":\"sellCost\",\"type\":\"uint256\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}]}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"recoverETH\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"executor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"approveInfinity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// VyperContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VyperContractMetaData.ABI instead.
var VyperContractABI = VyperContractMetaData.ABI

// VyperContract is an auto generated Go binding around an Ethereum contract.
type VyperContract struct {
	VyperContractCaller     // Read-only binding to the contract
	VyperContractTransactor // Write-only binding to the contract
	VyperContractFilterer   // Log filterer for contract events
}

// VyperContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VyperContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VyperContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VyperContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VyperContractSession struct {
	Contract     *VyperContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VyperContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VyperContractCallerSession struct {
	Contract *VyperContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VyperContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VyperContractTransactorSession struct {
	Contract     *VyperContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VyperContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VyperContractRaw struct {
	Contract *VyperContract // Generic contract binding to access the raw methods on
}

// VyperContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VyperContractCallerRaw struct {
	Contract *VyperContractCaller // Generic read-only contract binding to access the raw methods on
}

// VyperContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VyperContractTransactorRaw struct {
	Contract *VyperContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVyperContract creates a new instance of VyperContract, bound to a specific deployed contract.
func NewVyperContract(address common.Address, backend bind.ContractBackend) (*VyperContract, error) {
	contract, err := bindVyperContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VyperContract{VyperContractCaller: VyperContractCaller{contract: contract}, VyperContractTransactor: VyperContractTransactor{contract: contract}, VyperContractFilterer: VyperContractFilterer{contract: contract}}, nil
}

// NewVyperContractCaller creates a new read-only instance of VyperContract, bound to a specific deployed contract.
func NewVyperContractCaller(address common.Address, caller bind.ContractCaller) (*VyperContractCaller, error) {
	contract, err := bindVyperContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VyperContractCaller{contract: contract}, nil
}

// NewVyperContractTransactor creates a new write-only instance of VyperContract, bound to a specific deployed contract.
func NewVyperContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VyperContractTransactor, error) {
	contract, err := bindVyperContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VyperContractTransactor{contract: contract}, nil
}

// NewVyperContractFilterer creates a new log filterer instance of VyperContract, bound to a specific deployed contract.
func NewVyperContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VyperContractFilterer, error) {
	contract, err := bindVyperContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VyperContractFilterer{contract: contract}, nil
}

// bindVyperContract binds a generic wrapper to an already deployed contract.
func bindVyperContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VyperContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VyperContract *VyperContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VyperContract.Contract.VyperContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VyperContract *VyperContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VyperContract.Contract.VyperContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VyperContract *VyperContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VyperContract.Contract.VyperContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VyperContract *VyperContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VyperContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VyperContract *VyperContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VyperContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VyperContract *VyperContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VyperContract.Contract.contract.Transact(opts, method, params...)
}

// ApproveInfinity is a free data retrieval call binding the contract method 0x1578e9e7.
//
// Solidity: function approveInfinity() view returns(uint256)
func (_VyperContract *VyperContractCaller) ApproveInfinity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VyperContract.contract.Call(opts, &out, "approveInfinity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproveInfinity is a free data retrieval call binding the contract method 0x1578e9e7.
//
// Solidity: function approveInfinity() view returns(uint256)
func (_VyperContract *VyperContractSession) ApproveInfinity() (*big.Int, error) {
	return _VyperContract.Contract.ApproveInfinity(&_VyperContract.CallOpts)
}

// ApproveInfinity is a free data retrieval call binding the contract method 0x1578e9e7.
//
// Solidity: function approveInfinity() view returns(uint256)
func (_VyperContract *VyperContractCallerSession) ApproveInfinity() (*big.Int, error) {
	return _VyperContract.Contract.ApproveInfinity(&_VyperContract.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_VyperContract *VyperContractCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VyperContract.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_VyperContract *VyperContractSession) Executor() (common.Address, error) {
	return _VyperContract.Contract.Executor(&_VyperContract.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_VyperContract *VyperContractCallerSession) Executor() (common.Address, error) {
	return _VyperContract.Contract.Executor(&_VyperContract.CallOpts)
}

// HoneyCheck is a paid mutator transaction binding the contract method 0x4380a176.
//
// Solidity: function honeyCheck(address _buyToken, address _targetTokenAddress, address idexRouterAddres, uint256 _amount) payable returns((uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_VyperContract *VyperContractTransactor) HoneyCheck(opts *bind.TransactOpts, _buyToken common.Address, _targetTokenAddress common.Address, idexRouterAddres common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VyperContract.contract.Transact(opts, "honeyCheck", _buyToken, _targetTokenAddress, idexRouterAddres, _amount)
}

// HoneyCheck is a paid mutator transaction binding the contract method 0x4380a176.
//
// Solidity: function honeyCheck(address _buyToken, address _targetTokenAddress, address idexRouterAddres, uint256 _amount) payable returns((uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_VyperContract *VyperContractSession) HoneyCheck(_buyToken common.Address, _targetTokenAddress common.Address, idexRouterAddres common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VyperContract.Contract.HoneyCheck(&_VyperContract.TransactOpts, _buyToken, _targetTokenAddress, idexRouterAddres, _amount)
}

// HoneyCheck is a paid mutator transaction binding the contract method 0x4380a176.
//
// Solidity: function honeyCheck(address _buyToken, address _targetTokenAddress, address idexRouterAddres, uint256 _amount) payable returns((uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_VyperContract *VyperContractTransactorSession) HoneyCheck(_buyToken common.Address, _targetTokenAddress common.Address, idexRouterAddres common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VyperContract.Contract.HoneyCheck(&_VyperContract.TransactOpts, _buyToken, _targetTokenAddress, idexRouterAddres, _amount)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x1d6dc932.
//
// Solidity: function recoverETH(uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractTransactor) RecoverETH(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.contract.Transact(opts, "recoverETH", _amount, _to)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x1d6dc932.
//
// Solidity: function recoverETH(uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractSession) RecoverETH(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.Contract.RecoverETH(&_VyperContract.TransactOpts, _amount, _to)
}

// RecoverETH is a paid mutator transaction binding the contract method 0x1d6dc932.
//
// Solidity: function recoverETH(uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractTransactorSession) RecoverETH(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.Contract.RecoverETH(&_VyperContract.TransactOpts, _amount, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x61b0a56e.
//
// Solidity: function recoverTokens(address _token, uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractTransactor) RecoverTokens(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.contract.Transact(opts, "recoverTokens", _token, _amount, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x61b0a56e.
//
// Solidity: function recoverTokens(address _token, uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractSession) RecoverTokens(_token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.Contract.RecoverTokens(&_VyperContract.TransactOpts, _token, _amount, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x61b0a56e.
//
// Solidity: function recoverTokens(address _token, uint256 _amount, address _to) returns()
func (_VyperContract *VyperContractTransactorSession) RecoverTokens(_token common.Address, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _VyperContract.Contract.RecoverTokens(&_VyperContract.TransactOpts, _token, _amount, _to)
}
