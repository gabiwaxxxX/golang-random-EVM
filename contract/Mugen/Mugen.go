// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Mugen

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

// MugenMetaData contains all meta data concerning the Mugen contract.
var MugenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MinterSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"ReceiveFromChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"SendToChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"sendFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MugenABI is the input ABI used to generate the binding from.
// Deprecated: Use MugenMetaData.ABI instead.
var MugenABI = MugenMetaData.ABI

// Mugen is an auto generated Go binding around an Ethereum contract.
type Mugen struct {
	MugenCaller     // Read-only binding to the contract
	MugenTransactor // Write-only binding to the contract
	MugenFilterer   // Log filterer for contract events
}

// MugenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MugenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MugenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MugenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MugenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MugenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MugenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MugenSession struct {
	Contract     *Mugen            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MugenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MugenCallerSession struct {
	Contract *MugenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MugenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MugenTransactorSession struct {
	Contract     *MugenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MugenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MugenRaw struct {
	Contract *Mugen // Generic contract binding to access the raw methods on
}

// MugenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MugenCallerRaw struct {
	Contract *MugenCaller // Generic read-only contract binding to access the raw methods on
}

// MugenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MugenTransactorRaw struct {
	Contract *MugenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMugen creates a new instance of Mugen, bound to a specific deployed contract.
func NewMugen(address common.Address, backend bind.ContractBackend) (*Mugen, error) {
	contract, err := bindMugen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mugen{MugenCaller: MugenCaller{contract: contract}, MugenTransactor: MugenTransactor{contract: contract}, MugenFilterer: MugenFilterer{contract: contract}}, nil
}

// NewMugenCaller creates a new read-only instance of Mugen, bound to a specific deployed contract.
func NewMugenCaller(address common.Address, caller bind.ContractCaller) (*MugenCaller, error) {
	contract, err := bindMugen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MugenCaller{contract: contract}, nil
}

// NewMugenTransactor creates a new write-only instance of Mugen, bound to a specific deployed contract.
func NewMugenTransactor(address common.Address, transactor bind.ContractTransactor) (*MugenTransactor, error) {
	contract, err := bindMugen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MugenTransactor{contract: contract}, nil
}

// NewMugenFilterer creates a new log filterer instance of Mugen, bound to a specific deployed contract.
func NewMugenFilterer(address common.Address, filterer bind.ContractFilterer) (*MugenFilterer, error) {
	contract, err := bindMugen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MugenFilterer{contract: contract}, nil
}

// bindMugen binds a generic wrapper to an already deployed contract.
func bindMugen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MugenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mugen *MugenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mugen.Contract.MugenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mugen *MugenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mugen.Contract.MugenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mugen *MugenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mugen.Contract.MugenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mugen *MugenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mugen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mugen *MugenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mugen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mugen *MugenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mugen.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mugen *MugenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mugen *MugenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mugen.Contract.Allowance(&_Mugen.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mugen *MugenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mugen.Contract.Allowance(&_Mugen.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mugen *MugenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mugen *MugenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mugen.Contract.BalanceOf(&_Mugen.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mugen *MugenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mugen.Contract.BalanceOf(&_Mugen.CallOpts, account)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns(uint256)
func (_Mugen *MugenCaller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "circulatingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns(uint256)
func (_Mugen *MugenSession) CirculatingSupply() (*big.Int, error) {
	return _Mugen.Contract.CirculatingSupply(&_Mugen.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() view returns(uint256)
func (_Mugen *MugenCallerSession) CirculatingSupply() (*big.Int, error) {
	return _Mugen.Contract.CirculatingSupply(&_Mugen.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mugen *MugenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mugen *MugenSession) Decimals() (uint8, error) {
	return _Mugen.Contract.Decimals(&_Mugen.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mugen *MugenCallerSession) Decimals() (uint8, error) {
	return _Mugen.Contract.Decimals(&_Mugen.CallOpts)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Mugen *MugenCaller) EstimateSendFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "estimateSendFee", _dstChainId, _toAddress, _amount, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Mugen *MugenSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Mugen.Contract.EstimateSendFee(&_Mugen.CallOpts, _dstChainId, _toAddress, _amount, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Mugen *MugenCallerSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Mugen.Contract.EstimateSendFee(&_Mugen.CallOpts, _dstChainId, _toAddress, _amount, _useZro, _adapterParams)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Mugen *MugenCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Mugen *MugenSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Mugen.Contract.FailedMessages(&_Mugen.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Mugen *MugenCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Mugen.Contract.FailedMessages(&_Mugen.CallOpts, arg0, arg1, arg2)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Mugen *MugenCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Mugen *MugenSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Mugen.Contract.GetConfig(&_Mugen.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Mugen *MugenCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Mugen.Contract.GetConfig(&_Mugen.CallOpts, _version, _chainId, arg2, _configType)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Mugen *MugenCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Mugen *MugenSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Mugen.Contract.IsTrustedRemote(&_Mugen.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Mugen *MugenCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Mugen.Contract.IsTrustedRemote(&_Mugen.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Mugen *MugenCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Mugen *MugenSession) LzEndpoint() (common.Address, error) {
	return _Mugen.Contract.LzEndpoint(&_Mugen.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Mugen *MugenCallerSession) LzEndpoint() (common.Address, error) {
	return _Mugen.Contract.LzEndpoint(&_Mugen.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Mugen *MugenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Mugen *MugenSession) Minter() (common.Address, error) {
	return _Mugen.Contract.Minter(&_Mugen.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Mugen *MugenCallerSession) Minter() (common.Address, error) {
	return _Mugen.Contract.Minter(&_Mugen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mugen *MugenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mugen *MugenSession) Name() (string, error) {
	return _Mugen.Contract.Name(&_Mugen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mugen *MugenCallerSession) Name() (string, error) {
	return _Mugen.Contract.Name(&_Mugen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mugen *MugenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mugen *MugenSession) Owner() (common.Address, error) {
	return _Mugen.Contract.Owner(&_Mugen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mugen *MugenCallerSession) Owner() (common.Address, error) {
	return _Mugen.Contract.Owner(&_Mugen.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mugen *MugenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mugen *MugenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mugen.Contract.SupportsInterface(&_Mugen.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mugen *MugenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mugen.Contract.SupportsInterface(&_Mugen.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mugen *MugenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mugen *MugenSession) Symbol() (string, error) {
	return _Mugen.Contract.Symbol(&_Mugen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mugen *MugenCallerSession) Symbol() (string, error) {
	return _Mugen.Contract.Symbol(&_Mugen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mugen *MugenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mugen *MugenSession) TotalSupply() (*big.Int, error) {
	return _Mugen.Contract.TotalSupply(&_Mugen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mugen *MugenCallerSession) TotalSupply() (*big.Int, error) {
	return _Mugen.Contract.TotalSupply(&_Mugen.CallOpts)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Mugen *MugenCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Mugen.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Mugen *MugenSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Mugen.Contract.TrustedRemoteLookup(&_Mugen.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Mugen *MugenCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Mugen.Contract.TrustedRemoteLookup(&_Mugen.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mugen *MugenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mugen *MugenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Approve(&_Mugen.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mugen *MugenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Approve(&_Mugen.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mugen *MugenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mugen *MugenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.DecreaseAllowance(&_Mugen.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mugen *MugenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.DecreaseAllowance(&_Mugen.TransactOpts, spender, subtractedValue)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.Contract.ForceResumeReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.Contract.ForceResumeReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mugen *MugenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mugen *MugenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.IncreaseAllowance(&_Mugen.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mugen *MugenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.IncreaseAllowance(&_Mugen.TransactOpts, spender, addedValue)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.LzReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.LzReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Mugen *MugenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Mugen *MugenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Mint(&_Mugen.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Mugen *MugenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Mint(&_Mugen.TransactOpts, _to, _amount)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.NonblockingLzReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Mugen *MugenTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.NonblockingLzReceive(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mugen *MugenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mugen *MugenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mugen.Contract.RenounceOwnership(&_Mugen.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mugen *MugenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mugen.Contract.RenounceOwnership(&_Mugen.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Mugen *MugenTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Mugen *MugenSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.RetryMessage(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Mugen *MugenTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Mugen.Contract.RetryMessage(&_Mugen.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Mugen *MugenTransactor) SendFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "sendFrom", _from, _dstChainId, _toAddress, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Mugen *MugenSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SendFrom(&_Mugen.TransactOpts, _from, _dstChainId, _toAddress, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Mugen *MugenTransactorSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SendFrom(&_Mugen.TransactOpts, _from, _dstChainId, _toAddress, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Mugen *MugenTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Mugen *MugenSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SetConfig(&_Mugen.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Mugen *MugenTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SetConfig(&_Mugen.TransactOpts, _version, _chainId, _configType, _config)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Mugen *MugenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Mugen *MugenSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Mugen.Contract.SetMinter(&_Mugen.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Mugen *MugenTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Mugen.Contract.SetMinter(&_Mugen.TransactOpts, _minter)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Mugen *MugenTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Mugen *MugenSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Mugen.Contract.SetReceiveVersion(&_Mugen.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Mugen *MugenTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Mugen.Contract.SetReceiveVersion(&_Mugen.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Mugen *MugenTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Mugen *MugenSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Mugen.Contract.SetSendVersion(&_Mugen.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Mugen *MugenTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Mugen.Contract.SetSendVersion(&_Mugen.TransactOpts, _version)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenSession) SetTrustedRemote(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SetTrustedRemote(&_Mugen.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Mugen *MugenTransactorSession) SetTrustedRemote(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Mugen.Contract.SetTrustedRemote(&_Mugen.TransactOpts, _srcChainId, _srcAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mugen *MugenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mugen *MugenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Transfer(&_Mugen.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mugen *MugenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.Transfer(&_Mugen.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mugen *MugenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mugen *MugenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.TransferFrom(&_Mugen.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Mugen *MugenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mugen.Contract.TransferFrom(&_Mugen.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mugen *MugenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mugen.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mugen *MugenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mugen.Contract.TransferOwnership(&_Mugen.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mugen *MugenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mugen.Contract.TransferOwnership(&_Mugen.TransactOpts, newOwner)
}

// MugenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mugen contract.
type MugenApprovalIterator struct {
	Event *MugenApproval // Event containing the contract specifics and raw log

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
func (it *MugenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenApproval)
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
		it.Event = new(MugenApproval)
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
func (it *MugenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenApproval represents a Approval event raised by the Mugen contract.
type MugenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mugen *MugenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MugenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MugenApprovalIterator{contract: _Mugen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mugen *MugenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MugenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenApproval)
				if err := _Mugen.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mugen *MugenFilterer) ParseApproval(log types.Log) (*MugenApproval, error) {
	event := new(MugenApproval)
	if err := _Mugen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Mugen contract.
type MugenMessageFailedIterator struct {
	Event *MugenMessageFailed // Event containing the contract specifics and raw log

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
func (it *MugenMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenMessageFailed)
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
		it.Event = new(MugenMessageFailed)
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
func (it *MugenMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenMessageFailed represents a MessageFailed event raised by the Mugen contract.
type MugenMessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe6f254030bcb01ffd20558175c13fcaed6d1520be7becee4c961b65f79243b0d.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload)
func (_Mugen *MugenFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*MugenMessageFailedIterator, error) {

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &MugenMessageFailedIterator{contract: _Mugen.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe6f254030bcb01ffd20558175c13fcaed6d1520be7becee4c961b65f79243b0d.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload)
func (_Mugen *MugenFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *MugenMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenMessageFailed)
				if err := _Mugen.contract.UnpackLog(event, "MessageFailed", log); err != nil {
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

// ParseMessageFailed is a log parse operation binding the contract event 0xe6f254030bcb01ffd20558175c13fcaed6d1520be7becee4c961b65f79243b0d.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload)
func (_Mugen *MugenFilterer) ParseMessageFailed(log types.Log) (*MugenMessageFailed, error) {
	event := new(MugenMessageFailed)
	if err := _Mugen.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mugen contract.
type MugenOwnershipTransferredIterator struct {
	Event *MugenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MugenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenOwnershipTransferred)
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
		it.Event = new(MugenOwnershipTransferred)
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
func (it *MugenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenOwnershipTransferred represents a OwnershipTransferred event raised by the Mugen contract.
type MugenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mugen *MugenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MugenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MugenOwnershipTransferredIterator{contract: _Mugen.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mugen *MugenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MugenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenOwnershipTransferred)
				if err := _Mugen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mugen *MugenFilterer) ParseOwnershipTransferred(log types.Log) (*MugenOwnershipTransferred, error) {
	event := new(MugenOwnershipTransferred)
	if err := _Mugen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenReceiveFromChainIterator is returned from FilterReceiveFromChain and is used to iterate over the raw logs and unpacked data for ReceiveFromChain events raised by the Mugen contract.
type MugenReceiveFromChainIterator struct {
	Event *MugenReceiveFromChain // Event containing the contract specifics and raw log

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
func (it *MugenReceiveFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenReceiveFromChain)
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
		it.Event = new(MugenReceiveFromChain)
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
func (it *MugenReceiveFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenReceiveFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenReceiveFromChain represents a ReceiveFromChain event raised by the Mugen contract.
type MugenReceiveFromChain struct {
	SrcChainId uint16
	SrcAddress common.Hash
	ToAddress  common.Address
	Amount     *big.Int
	Nonce      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveFromChain is a free log retrieval operation binding the contract event 0x64e10c37f404d128982dce114f5d233c14c5c7f6d8db93099e3d99dacb9e27ba.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) FilterReceiveFromChain(opts *bind.FilterOpts, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (*MugenReceiveFromChainIterator, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &MugenReceiveFromChainIterator{contract: _Mugen.contract, event: "ReceiveFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveFromChain is a free log subscription operation binding the contract event 0x64e10c37f404d128982dce114f5d233c14c5c7f6d8db93099e3d99dacb9e27ba.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) WatchReceiveFromChain(opts *bind.WatchOpts, sink chan<- *MugenReceiveFromChain, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (event.Subscription, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenReceiveFromChain)
				if err := _Mugen.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
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

// ParseReceiveFromChain is a log parse operation binding the contract event 0x64e10c37f404d128982dce114f5d233c14c5c7f6d8db93099e3d99dacb9e27ba.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) ParseReceiveFromChain(log types.Log) (*MugenReceiveFromChain, error) {
	event := new(MugenReceiveFromChain)
	if err := _Mugen.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenSendToChainIterator is returned from FilterSendToChain and is used to iterate over the raw logs and unpacked data for SendToChain events raised by the Mugen contract.
type MugenSendToChainIterator struct {
	Event *MugenSendToChain // Event containing the contract specifics and raw log

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
func (it *MugenSendToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenSendToChain)
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
		it.Event = new(MugenSendToChain)
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
func (it *MugenSendToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenSendToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenSendToChain represents a SendToChain event raised by the Mugen contract.
type MugenSendToChain struct {
	Sender     common.Address
	DstChainId uint16
	ToAddress  common.Hash
	Amount     *big.Int
	Nonce      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendToChain is a free log retrieval operation binding the contract event 0x024797cc77ce15dc717112d54fb1df125fdfd8c81344fb046c5e074427ce1543.
//
// Solidity: event SendToChain(address indexed _sender, uint16 indexed _dstChainId, bytes indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) FilterSendToChain(opts *bind.FilterOpts, _sender []common.Address, _dstChainId []uint16, _toAddress [][]byte) (*MugenSendToChainIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "SendToChain", _senderRule, _dstChainIdRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &MugenSendToChainIterator{contract: _Mugen.contract, event: "SendToChain", logs: logs, sub: sub}, nil
}

// WatchSendToChain is a free log subscription operation binding the contract event 0x024797cc77ce15dc717112d54fb1df125fdfd8c81344fb046c5e074427ce1543.
//
// Solidity: event SendToChain(address indexed _sender, uint16 indexed _dstChainId, bytes indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) WatchSendToChain(opts *bind.WatchOpts, sink chan<- *MugenSendToChain, _sender []common.Address, _dstChainId []uint16, _toAddress [][]byte) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "SendToChain", _senderRule, _dstChainIdRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenSendToChain)
				if err := _Mugen.contract.UnpackLog(event, "SendToChain", log); err != nil {
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

// ParseSendToChain is a log parse operation binding the contract event 0x024797cc77ce15dc717112d54fb1df125fdfd8c81344fb046c5e074427ce1543.
//
// Solidity: event SendToChain(address indexed _sender, uint16 indexed _dstChainId, bytes indexed _toAddress, uint256 _amount, uint64 _nonce)
func (_Mugen *MugenFilterer) ParseSendToChain(log types.Log) (*MugenSendToChain, error) {
	event := new(MugenSendToChain)
	if err := _Mugen.contract.UnpackLog(event, "SendToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Mugen contract.
type MugenSetTrustedRemoteIterator struct {
	Event *MugenSetTrustedRemote // Event containing the contract specifics and raw log

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
func (it *MugenSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenSetTrustedRemote)
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
		it.Event = new(MugenSetTrustedRemote)
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
func (it *MugenSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenSetTrustedRemote represents a SetTrustedRemote event raised by the Mugen contract.
type MugenSetTrustedRemote struct {
	SrcChainId uint16
	SrcAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _srcChainId, bytes _srcAddress)
func (_Mugen *MugenFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*MugenSetTrustedRemoteIterator, error) {

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &MugenSetTrustedRemoteIterator{contract: _Mugen.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _srcChainId, bytes _srcAddress)
func (_Mugen *MugenFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *MugenSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenSetTrustedRemote)
				if err := _Mugen.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
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

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _srcChainId, bytes _srcAddress)
func (_Mugen *MugenFilterer) ParseSetTrustedRemote(log types.Log) (*MugenSetTrustedRemote, error) {
	event := new(MugenSetTrustedRemote)
	if err := _Mugen.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MugenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mugen contract.
type MugenTransferIterator struct {
	Event *MugenTransfer // Event containing the contract specifics and raw log

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
func (it *MugenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MugenTransfer)
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
		it.Event = new(MugenTransfer)
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
func (it *MugenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MugenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MugenTransfer represents a Transfer event raised by the Mugen contract.
type MugenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mugen *MugenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MugenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mugen.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MugenTransferIterator{contract: _Mugen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mugen *MugenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MugenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mugen.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MugenTransfer)
				if err := _Mugen.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mugen *MugenFilterer) ParseTransfer(log types.Log) (*MugenTransfer, error) {
	event := new(MugenTransfer)
	if err := _Mugen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
