// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Adv3nturers

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

// Adv3nturersMetaData contains all meta data concerning the Adv3nturers contract.
var Adv3nturersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAgility\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getConstitution\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getElement\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getIntuition\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLuck\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPerception\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getStrength\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getWisdom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getZodiac\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Adv3nturersABI is the input ABI used to generate the binding from.
// Deprecated: Use Adv3nturersMetaData.ABI instead.
var Adv3nturersABI = Adv3nturersMetaData.ABI

// Adv3nturers is an auto generated Go binding around an Ethereum contract.
type Adv3nturers struct {
	Adv3nturersCaller     // Read-only binding to the contract
	Adv3nturersTransactor // Write-only binding to the contract
	Adv3nturersFilterer   // Log filterer for contract events
}

// Adv3nturersCaller is an auto generated read-only Go binding around an Ethereum contract.
type Adv3nturersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Adv3nturersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Adv3nturersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Adv3nturersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Adv3nturersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Adv3nturersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Adv3nturersSession struct {
	Contract     *Adv3nturers      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Adv3nturersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Adv3nturersCallerSession struct {
	Contract *Adv3nturersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Adv3nturersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Adv3nturersTransactorSession struct {
	Contract     *Adv3nturersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Adv3nturersRaw is an auto generated low-level Go binding around an Ethereum contract.
type Adv3nturersRaw struct {
	Contract *Adv3nturers // Generic contract binding to access the raw methods on
}

// Adv3nturersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Adv3nturersCallerRaw struct {
	Contract *Adv3nturersCaller // Generic read-only contract binding to access the raw methods on
}

// Adv3nturersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Adv3nturersTransactorRaw struct {
	Contract *Adv3nturersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdv3nturers creates a new instance of Adv3nturers, bound to a specific deployed contract.
func NewAdv3nturers(address common.Address, backend bind.ContractBackend) (*Adv3nturers, error) {
	contract, err := bindAdv3nturers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adv3nturers{Adv3nturersCaller: Adv3nturersCaller{contract: contract}, Adv3nturersTransactor: Adv3nturersTransactor{contract: contract}, Adv3nturersFilterer: Adv3nturersFilterer{contract: contract}}, nil
}

// NewAdv3nturersCaller creates a new read-only instance of Adv3nturers, bound to a specific deployed contract.
func NewAdv3nturersCaller(address common.Address, caller bind.ContractCaller) (*Adv3nturersCaller, error) {
	contract, err := bindAdv3nturers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersCaller{contract: contract}, nil
}

// NewAdv3nturersTransactor creates a new write-only instance of Adv3nturers, bound to a specific deployed contract.
func NewAdv3nturersTransactor(address common.Address, transactor bind.ContractTransactor) (*Adv3nturersTransactor, error) {
	contract, err := bindAdv3nturers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersTransactor{contract: contract}, nil
}

// NewAdv3nturersFilterer creates a new log filterer instance of Adv3nturers, bound to a specific deployed contract.
func NewAdv3nturersFilterer(address common.Address, filterer bind.ContractFilterer) (*Adv3nturersFilterer, error) {
	contract, err := bindAdv3nturers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersFilterer{contract: contract}, nil
}

// bindAdv3nturers binds a generic wrapper to an already deployed contract.
func bindAdv3nturers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Adv3nturersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adv3nturers *Adv3nturersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adv3nturers.Contract.Adv3nturersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adv3nturers *Adv3nturersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Adv3nturersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adv3nturers *Adv3nturersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Adv3nturersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adv3nturers *Adv3nturersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adv3nturers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adv3nturers *Adv3nturersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adv3nturers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adv3nturers *Adv3nturersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adv3nturers.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Adv3nturers *Adv3nturersCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Adv3nturers *Adv3nturersSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Adv3nturers.Contract.BalanceOf(&_Adv3nturers.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Adv3nturers *Adv3nturersCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Adv3nturers.Contract.BalanceOf(&_Adv3nturers.CallOpts, owner)
}

// GetAgility is a free data retrieval call binding the contract method 0x8bbee1f3.
//
// Solidity: function getAgility(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetAgility(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getAgility", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAgility is a free data retrieval call binding the contract method 0x8bbee1f3.
//
// Solidity: function getAgility(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetAgility(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetAgility(&_Adv3nturers.CallOpts, tokenId)
}

// GetAgility is a free data retrieval call binding the contract method 0x8bbee1f3.
//
// Solidity: function getAgility(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetAgility(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetAgility(&_Adv3nturers.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Adv3nturers.Contract.GetApproved(&_Adv3nturers.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Adv3nturers.Contract.GetApproved(&_Adv3nturers.CallOpts, tokenId)
}

// GetConstitution is a free data retrieval call binding the contract method 0xf3a4aba9.
//
// Solidity: function getConstitution(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetConstitution(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getConstitution", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetConstitution is a free data retrieval call binding the contract method 0xf3a4aba9.
//
// Solidity: function getConstitution(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetConstitution(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetConstitution(&_Adv3nturers.CallOpts, tokenId)
}

// GetConstitution is a free data retrieval call binding the contract method 0xf3a4aba9.
//
// Solidity: function getConstitution(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetConstitution(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetConstitution(&_Adv3nturers.CallOpts, tokenId)
}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetElement(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getElement", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetElement(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetElement(&_Adv3nturers.CallOpts, tokenId)
}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetElement(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetElement(&_Adv3nturers.CallOpts, tokenId)
}

// GetIntuition is a free data retrieval call binding the contract method 0xaa430941.
//
// Solidity: function getIntuition(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetIntuition(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getIntuition", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIntuition is a free data retrieval call binding the contract method 0xaa430941.
//
// Solidity: function getIntuition(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetIntuition(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetIntuition(&_Adv3nturers.CallOpts, tokenId)
}

// GetIntuition is a free data retrieval call binding the contract method 0xaa430941.
//
// Solidity: function getIntuition(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetIntuition(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetIntuition(&_Adv3nturers.CallOpts, tokenId)
}

// GetLuck is a free data retrieval call binding the contract method 0xc7511545.
//
// Solidity: function getLuck(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetLuck(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getLuck", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLuck is a free data retrieval call binding the contract method 0xc7511545.
//
// Solidity: function getLuck(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetLuck(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetLuck(&_Adv3nturers.CallOpts, tokenId)
}

// GetLuck is a free data retrieval call binding the contract method 0xc7511545.
//
// Solidity: function getLuck(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetLuck(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetLuck(&_Adv3nturers.CallOpts, tokenId)
}

// GetPerception is a free data retrieval call binding the contract method 0x8a007d3d.
//
// Solidity: function getPerception(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetPerception(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getPerception", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPerception is a free data retrieval call binding the contract method 0x8a007d3d.
//
// Solidity: function getPerception(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetPerception(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetPerception(&_Adv3nturers.CallOpts, tokenId)
}

// GetPerception is a free data retrieval call binding the contract method 0x8a007d3d.
//
// Solidity: function getPerception(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetPerception(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetPerception(&_Adv3nturers.CallOpts, tokenId)
}

// GetStrength is a free data retrieval call binding the contract method 0xe53fbda6.
//
// Solidity: function getStrength(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetStrength(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getStrength", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetStrength is a free data retrieval call binding the contract method 0xe53fbda6.
//
// Solidity: function getStrength(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetStrength(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetStrength(&_Adv3nturers.CallOpts, tokenId)
}

// GetStrength is a free data retrieval call binding the contract method 0xe53fbda6.
//
// Solidity: function getStrength(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetStrength(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetStrength(&_Adv3nturers.CallOpts, tokenId)
}

// GetWisdom is a free data retrieval call binding the contract method 0x30799dc6.
//
// Solidity: function getWisdom(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetWisdom(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getWisdom", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetWisdom is a free data retrieval call binding the contract method 0x30799dc6.
//
// Solidity: function getWisdom(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetWisdom(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetWisdom(&_Adv3nturers.CallOpts, tokenId)
}

// GetWisdom is a free data retrieval call binding the contract method 0x30799dc6.
//
// Solidity: function getWisdom(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetWisdom(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetWisdom(&_Adv3nturers.CallOpts, tokenId)
}

// GetZodiac is a free data retrieval call binding the contract method 0x0c9066dd.
//
// Solidity: function getZodiac(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) GetZodiac(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "getZodiac", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetZodiac is a free data retrieval call binding the contract method 0x0c9066dd.
//
// Solidity: function getZodiac(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) GetZodiac(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetZodiac(&_Adv3nturers.CallOpts, tokenId)
}

// GetZodiac is a free data retrieval call binding the contract method 0x0c9066dd.
//
// Solidity: function getZodiac(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) GetZodiac(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.GetZodiac(&_Adv3nturers.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Adv3nturers *Adv3nturersCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Adv3nturers *Adv3nturersSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Adv3nturers.Contract.IsApprovedForAll(&_Adv3nturers.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Adv3nturers *Adv3nturersCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Adv3nturers.Contract.IsApprovedForAll(&_Adv3nturers.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Adv3nturers *Adv3nturersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Adv3nturers *Adv3nturersSession) Name() (string, error) {
	return _Adv3nturers.Contract.Name(&_Adv3nturers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) Name() (string, error) {
	return _Adv3nturers.Contract.Name(&_Adv3nturers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adv3nturers *Adv3nturersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adv3nturers *Adv3nturersSession) Owner() (common.Address, error) {
	return _Adv3nturers.Contract.Owner(&_Adv3nturers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adv3nturers *Adv3nturersCallerSession) Owner() (common.Address, error) {
	return _Adv3nturers.Contract.Owner(&_Adv3nturers.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Adv3nturers.Contract.OwnerOf(&_Adv3nturers.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Adv3nturers *Adv3nturersCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Adv3nturers.Contract.OwnerOf(&_Adv3nturers.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Adv3nturers *Adv3nturersCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Adv3nturers *Adv3nturersSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Adv3nturers.Contract.SupportsInterface(&_Adv3nturers.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Adv3nturers *Adv3nturersCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Adv3nturers.Contract.SupportsInterface(&_Adv3nturers.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Adv3nturers *Adv3nturersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Adv3nturers *Adv3nturersSession) Symbol() (string, error) {
	return _Adv3nturers.Contract.Symbol(&_Adv3nturers.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) Symbol() (string, error) {
	return _Adv3nturers.Contract.Symbol(&_Adv3nturers.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Adv3nturers.Contract.TokenByIndex(&_Adv3nturers.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Adv3nturers.Contract.TokenByIndex(&_Adv3nturers.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Adv3nturers.Contract.TokenOfOwnerByIndex(&_Adv3nturers.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Adv3nturers *Adv3nturersCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Adv3nturers.Contract.TokenOfOwnerByIndex(&_Adv3nturers.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.TokenURI(&_Adv3nturers.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Adv3nturers *Adv3nturersCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Adv3nturers.Contract.TokenURI(&_Adv3nturers.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Adv3nturers *Adv3nturersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Adv3nturers.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Adv3nturers *Adv3nturersSession) TotalSupply() (*big.Int, error) {
	return _Adv3nturers.Contract.TotalSupply(&_Adv3nturers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Adv3nturers *Adv3nturersCallerSession) TotalSupply() (*big.Int, error) {
	return _Adv3nturers.Contract.TotalSupply(&_Adv3nturers.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Approve(&_Adv3nturers.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Approve(&_Adv3nturers.TransactOpts, to, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactor) Claim(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "claim", tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersSession) Claim(tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Claim(&_Adv3nturers.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) Claim(tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.Claim(&_Adv3nturers.TransactOpts, tokenId)
}

// OwnerClaim is a paid mutator transaction binding the contract method 0x434f48c4.
//
// Solidity: function ownerClaim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactor) OwnerClaim(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "ownerClaim", tokenId)
}

// OwnerClaim is a paid mutator transaction binding the contract method 0x434f48c4.
//
// Solidity: function ownerClaim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersSession) OwnerClaim(tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.OwnerClaim(&_Adv3nturers.TransactOpts, tokenId)
}

// OwnerClaim is a paid mutator transaction binding the contract method 0x434f48c4.
//
// Solidity: function ownerClaim(uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) OwnerClaim(tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.OwnerClaim(&_Adv3nturers.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adv3nturers *Adv3nturersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adv3nturers *Adv3nturersSession) RenounceOwnership() (*types.Transaction, error) {
	return _Adv3nturers.Contract.RenounceOwnership(&_Adv3nturers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adv3nturers *Adv3nturersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Adv3nturers.Contract.RenounceOwnership(&_Adv3nturers.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SafeTransferFrom(&_Adv3nturers.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SafeTransferFrom(&_Adv3nturers.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Adv3nturers *Adv3nturersTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Adv3nturers *Adv3nturersSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SafeTransferFrom0(&_Adv3nturers.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SafeTransferFrom0(&_Adv3nturers.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Adv3nturers *Adv3nturersTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Adv3nturers *Adv3nturersSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SetApprovalForAll(&_Adv3nturers.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Adv3nturers.Contract.SetApprovalForAll(&_Adv3nturers.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.TransferFrom(&_Adv3nturers.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Adv3nturers.Contract.TransferFrom(&_Adv3nturers.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adv3nturers *Adv3nturersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Adv3nturers.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adv3nturers *Adv3nturersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Adv3nturers.Contract.TransferOwnership(&_Adv3nturers.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adv3nturers *Adv3nturersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Adv3nturers.Contract.TransferOwnership(&_Adv3nturers.TransactOpts, newOwner)
}

// Adv3nturersApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Adv3nturers contract.
type Adv3nturersApprovalIterator struct {
	Event *Adv3nturersApproval // Event containing the contract specifics and raw log

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
func (it *Adv3nturersApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Adv3nturersApproval)
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
		it.Event = new(Adv3nturersApproval)
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
func (it *Adv3nturersApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Adv3nturersApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Adv3nturersApproval represents a Approval event raised by the Adv3nturers contract.
type Adv3nturersApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Adv3nturersApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Adv3nturers.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersApprovalIterator{contract: _Adv3nturers.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Adv3nturersApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Adv3nturers.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Adv3nturersApproval)
				if err := _Adv3nturers.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) ParseApproval(log types.Log) (*Adv3nturersApproval, error) {
	event := new(Adv3nturersApproval)
	if err := _Adv3nturers.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Adv3nturersApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Adv3nturers contract.
type Adv3nturersApprovalForAllIterator struct {
	Event *Adv3nturersApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Adv3nturersApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Adv3nturersApprovalForAll)
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
		it.Event = new(Adv3nturersApprovalForAll)
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
func (it *Adv3nturersApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Adv3nturersApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Adv3nturersApprovalForAll represents a ApprovalForAll event raised by the Adv3nturers contract.
type Adv3nturersApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Adv3nturers *Adv3nturersFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Adv3nturersApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Adv3nturers.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersApprovalForAllIterator{contract: _Adv3nturers.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Adv3nturers *Adv3nturersFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Adv3nturersApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Adv3nturers.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Adv3nturersApprovalForAll)
				if err := _Adv3nturers.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Adv3nturers *Adv3nturersFilterer) ParseApprovalForAll(log types.Log) (*Adv3nturersApprovalForAll, error) {
	event := new(Adv3nturersApprovalForAll)
	if err := _Adv3nturers.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Adv3nturersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Adv3nturers contract.
type Adv3nturersOwnershipTransferredIterator struct {
	Event *Adv3nturersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Adv3nturersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Adv3nturersOwnershipTransferred)
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
		it.Event = new(Adv3nturersOwnershipTransferred)
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
func (it *Adv3nturersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Adv3nturersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Adv3nturersOwnershipTransferred represents a OwnershipTransferred event raised by the Adv3nturers contract.
type Adv3nturersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Adv3nturers *Adv3nturersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Adv3nturersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Adv3nturers.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersOwnershipTransferredIterator{contract: _Adv3nturers.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Adv3nturers *Adv3nturersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Adv3nturersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Adv3nturers.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Adv3nturersOwnershipTransferred)
				if err := _Adv3nturers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Adv3nturers *Adv3nturersFilterer) ParseOwnershipTransferred(log types.Log) (*Adv3nturersOwnershipTransferred, error) {
	event := new(Adv3nturersOwnershipTransferred)
	if err := _Adv3nturers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Adv3nturersTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Adv3nturers contract.
type Adv3nturersTransferIterator struct {
	Event *Adv3nturersTransfer // Event containing the contract specifics and raw log

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
func (it *Adv3nturersTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Adv3nturersTransfer)
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
		it.Event = new(Adv3nturersTransfer)
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
func (it *Adv3nturersTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Adv3nturersTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Adv3nturersTransfer represents a Transfer event raised by the Adv3nturers contract.
type Adv3nturersTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Adv3nturersTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Adv3nturers.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Adv3nturersTransferIterator{contract: _Adv3nturers.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Adv3nturersTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Adv3nturers.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Adv3nturersTransfer)
				if err := _Adv3nturers.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Adv3nturers *Adv3nturersFilterer) ParseTransfer(log types.Log) (*Adv3nturersTransfer, error) {
	event := new(Adv3nturersTransfer)
	if err := _Adv3nturers.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
