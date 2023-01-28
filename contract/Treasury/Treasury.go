// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Treasury

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

// TreasuryMetaData contains all meta data concerning the Treasury contract.
var TreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mugen\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_administrator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdminRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CapReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommunicator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDepositable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderMinDeposit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_priceFreed\",\"type\":\"address\"}],\"name\":\"DepositableToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Communicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"accurateWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pricefeed\",\"type\":\"address\"}],\"name\":\"addTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminRemoved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateContinuousMintReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"changeTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"checkDepositable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sourceReserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_sourceReserveWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_targetReserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_targetReserveWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"crossReserveTargetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_reserveRatio\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_reserveRatio\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fundSupplyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initMaxExpArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_reserveRatio\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidateReserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mugen\",\"outputs\":[{\"internalType\":\"contractIMugen\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"normalizedWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseN\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseD\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_expN\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_expD\",\"type\":\"uint32\"}],\"name\":\"power\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"contractAggregatorPriceFeeds\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_reserveWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"purchaseTargetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"safeFactors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_reserveWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"saleTargetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_comms\",\"type\":\"address\"}],\"name\":\"setCommunicator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valueDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasuryMetaData.ABI instead.
var TreasuryABI = TreasuryMetaData.ABI

// Treasury is an auto generated Go binding around an Ethereum contract.
type Treasury struct {
	TreasuryCaller     // Read-only binding to the contract
	TreasuryTransactor // Write-only binding to the contract
	TreasuryFilterer   // Log filterer for contract events
}

// TreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasurySession struct {
	Contract     *Treasury         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryCallerSession struct {
	Contract *TreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryTransactorSession struct {
	Contract     *TreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryRaw struct {
	Contract *Treasury // Generic contract binding to access the raw methods on
}

// TreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryCallerRaw struct {
	Contract *TreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryTransactorRaw struct {
	Contract *TreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasury creates a new instance of Treasury, bound to a specific deployed contract.
func NewTreasury(address common.Address, backend bind.ContractBackend) (*Treasury, error) {
	contract, err := bindTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}, TreasuryFilterer: TreasuryFilterer{contract: contract}}, nil
}

// NewTreasuryCaller creates a new read-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryCaller(address common.Address, caller bind.ContractCaller) (*TreasuryCaller, error) {
	contract, err := bindTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryCaller{contract: contract}, nil
}

// NewTreasuryTransactor creates a new write-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryTransactor, error) {
	contract, err := bindTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryTransactor{contract: contract}, nil
}

// NewTreasuryFilterer creates a new log filterer instance of Treasury, bound to a specific deployed contract.
func NewTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryFilterer, error) {
	contract, err := bindTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryFilterer{contract: contract}, nil
}

// bindTreasury binds a generic wrapper to an already deployed contract.
func bindTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.TreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transact(opts, method, params...)
}

// Communicator is a free data retrieval call binding the contract method 0x77ceba2c.
//
// Solidity: function Communicator() view returns(address)
func (_Treasury *TreasuryCaller) Communicator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "Communicator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Communicator is a free data retrieval call binding the contract method 0x77ceba2c.
//
// Solidity: function Communicator() view returns(address)
func (_Treasury *TreasurySession) Communicator() (common.Address, error) {
	return _Treasury.Contract.Communicator(&_Treasury.CallOpts)
}

// Communicator is a free data retrieval call binding the contract method 0x77ceba2c.
//
// Solidity: function Communicator() view returns(address)
func (_Treasury *TreasuryCallerSession) Communicator() (common.Address, error) {
	return _Treasury.Contract.Communicator(&_Treasury.CallOpts)
}

// RESERVERATIO is a free data retrieval call binding the contract method 0xe37dbfd9.
//
// Solidity: function RESERVE_RATIO() view returns(uint256)
func (_Treasury *TreasuryCaller) RESERVERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "RESERVE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESERVERATIO is a free data retrieval call binding the contract method 0xe37dbfd9.
//
// Solidity: function RESERVE_RATIO() view returns(uint256)
func (_Treasury *TreasurySession) RESERVERATIO() (*big.Int, error) {
	return _Treasury.Contract.RESERVERATIO(&_Treasury.CallOpts)
}

// RESERVERATIO is a free data retrieval call binding the contract method 0xe37dbfd9.
//
// Solidity: function RESERVE_RATIO() view returns(uint256)
func (_Treasury *TreasuryCallerSession) RESERVERATIO() (*big.Int, error) {
	return _Treasury.Contract.RESERVERATIO(&_Treasury.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Treasury *TreasuryCaller) SCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Treasury *TreasurySession) SCALE() (*big.Int, error) {
	return _Treasury.Contract.SCALE(&_Treasury.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint256)
func (_Treasury *TreasuryCallerSession) SCALE() (*big.Int, error) {
	return _Treasury.Contract.SCALE(&_Treasury.CallOpts)
}

// AccurateWeights is a free data retrieval call binding the contract method 0x30c7d637.
//
// Solidity: function accurateWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasuryCaller) AccurateWeights(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (uint32, uint32, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "accurateWeights", _a, _b)

	if err != nil {
		return *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// AccurateWeights is a free data retrieval call binding the contract method 0x30c7d637.
//
// Solidity: function accurateWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasurySession) AccurateWeights(_a *big.Int, _b *big.Int) (uint32, uint32, error) {
	return _Treasury.Contract.AccurateWeights(&_Treasury.CallOpts, _a, _b)
}

// AccurateWeights is a free data retrieval call binding the contract method 0x30c7d637.
//
// Solidity: function accurateWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasuryCallerSession) AccurateWeights(_a *big.Int, _b *big.Int) (uint32, uint32, error) {
	return _Treasury.Contract.AccurateWeights(&_Treasury.CallOpts, _a, _b)
}

// AdminRemoved is a free data retrieval call binding the contract method 0x41bdb954.
//
// Solidity: function adminRemoved() view returns(bool)
func (_Treasury *TreasuryCaller) AdminRemoved(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "adminRemoved")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AdminRemoved is a free data retrieval call binding the contract method 0x41bdb954.
//
// Solidity: function adminRemoved() view returns(bool)
func (_Treasury *TreasurySession) AdminRemoved() (bool, error) {
	return _Treasury.Contract.AdminRemoved(&_Treasury.CallOpts)
}

// AdminRemoved is a free data retrieval call binding the contract method 0x41bdb954.
//
// Solidity: function adminRemoved() view returns(bool)
func (_Treasury *TreasuryCallerSession) AdminRemoved() (bool, error) {
	return _Treasury.Contract.AdminRemoved(&_Treasury.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Treasury *TreasuryCaller) Administrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "administrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Treasury *TreasurySession) Administrator() (common.Address, error) {
	return _Treasury.Contract.Administrator(&_Treasury.CallOpts)
}

// Administrator is a free data retrieval call binding the contract method 0xf53d0a8e.
//
// Solidity: function administrator() view returns(address)
func (_Treasury *TreasuryCallerSession) Administrator() (common.Address, error) {
	return _Treasury.Contract.Administrator(&_Treasury.CallOpts)
}

// CalculateContinuousMintReturn is a free data retrieval call binding the contract method 0x28c3d701.
//
// Solidity: function calculateContinuousMintReturn(uint256 _amount) view returns(uint256 mintAmount)
func (_Treasury *TreasuryCaller) CalculateContinuousMintReturn(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "calculateContinuousMintReturn", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateContinuousMintReturn is a free data retrieval call binding the contract method 0x28c3d701.
//
// Solidity: function calculateContinuousMintReturn(uint256 _amount) view returns(uint256 mintAmount)
func (_Treasury *TreasurySession) CalculateContinuousMintReturn(_amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.CalculateContinuousMintReturn(&_Treasury.CallOpts, _amount)
}

// CalculateContinuousMintReturn is a free data retrieval call binding the contract method 0x28c3d701.
//
// Solidity: function calculateContinuousMintReturn(uint256 _amount) view returns(uint256 mintAmount)
func (_Treasury *TreasuryCallerSession) CalculateContinuousMintReturn(_amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.CalculateContinuousMintReturn(&_Treasury.CallOpts, _amount)
}

// CheckDepositable is a free data retrieval call binding the contract method 0x0c31264d.
//
// Solidity: function checkDepositable(address _token) view returns(bool)
func (_Treasury *TreasuryCaller) CheckDepositable(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "checkDepositable", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDepositable is a free data retrieval call binding the contract method 0x0c31264d.
//
// Solidity: function checkDepositable(address _token) view returns(bool)
func (_Treasury *TreasurySession) CheckDepositable(_token common.Address) (bool, error) {
	return _Treasury.Contract.CheckDepositable(&_Treasury.CallOpts, _token)
}

// CheckDepositable is a free data retrieval call binding the contract method 0x0c31264d.
//
// Solidity: function checkDepositable(address _token) view returns(bool)
func (_Treasury *TreasuryCallerSession) CheckDepositable(_token common.Address) (bool, error) {
	return _Treasury.Contract.CheckDepositable(&_Treasury.CallOpts, _token)
}

// CrossReserveTargetAmount is a free data retrieval call binding the contract method 0x94491fab.
//
// Solidity: function crossReserveTargetAmount(uint256 _sourceReserveBalance, uint32 _sourceReserveWeight, uint256 _targetReserveBalance, uint32 _targetReserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) CrossReserveTargetAmount(opts *bind.CallOpts, _sourceReserveBalance *big.Int, _sourceReserveWeight uint32, _targetReserveBalance *big.Int, _targetReserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "crossReserveTargetAmount", _sourceReserveBalance, _sourceReserveWeight, _targetReserveBalance, _targetReserveWeight, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossReserveTargetAmount is a free data retrieval call binding the contract method 0x94491fab.
//
// Solidity: function crossReserveTargetAmount(uint256 _sourceReserveBalance, uint32 _sourceReserveWeight, uint256 _targetReserveBalance, uint32 _targetReserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) CrossReserveTargetAmount(_sourceReserveBalance *big.Int, _sourceReserveWeight uint32, _targetReserveBalance *big.Int, _targetReserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.CrossReserveTargetAmount(&_Treasury.CallOpts, _sourceReserveBalance, _sourceReserveWeight, _targetReserveBalance, _targetReserveWeight, _amount)
}

// CrossReserveTargetAmount is a free data retrieval call binding the contract method 0x94491fab.
//
// Solidity: function crossReserveTargetAmount(uint256 _sourceReserveBalance, uint32 _sourceReserveWeight, uint256 _targetReserveBalance, uint32 _targetReserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) CrossReserveTargetAmount(_sourceReserveBalance *big.Int, _sourceReserveWeight uint32, _targetReserveBalance *big.Int, _targetReserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.CrossReserveTargetAmount(&_Treasury.CallOpts, _sourceReserveBalance, _sourceReserveWeight, _targetReserveBalance, _targetReserveWeight, _amount)
}

// DepositCap is a free data retrieval call binding the contract method 0xdbd5edc7.
//
// Solidity: function depositCap() view returns(uint256)
func (_Treasury *TreasuryCaller) DepositCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "depositCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCap is a free data retrieval call binding the contract method 0xdbd5edc7.
//
// Solidity: function depositCap() view returns(uint256)
func (_Treasury *TreasurySession) DepositCap() (*big.Int, error) {
	return _Treasury.Contract.DepositCap(&_Treasury.CallOpts)
}

// DepositCap is a free data retrieval call binding the contract method 0xdbd5edc7.
//
// Solidity: function depositCap() view returns(uint256)
func (_Treasury *TreasuryCallerSession) DepositCap() (*big.Int, error) {
	return _Treasury.Contract.DepositCap(&_Treasury.CallOpts)
}

// DepositableTokens is a free data retrieval call binding the contract method 0xc14edd79.
//
// Solidity: function depositableTokens(address ) view returns(bool)
func (_Treasury *TreasuryCaller) DepositableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "depositableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositableTokens is a free data retrieval call binding the contract method 0xc14edd79.
//
// Solidity: function depositableTokens(address ) view returns(bool)
func (_Treasury *TreasurySession) DepositableTokens(arg0 common.Address) (bool, error) {
	return _Treasury.Contract.DepositableTokens(&_Treasury.CallOpts, arg0)
}

// DepositableTokens is a free data retrieval call binding the contract method 0xc14edd79.
//
// Solidity: function depositableTokens(address ) view returns(bool)
func (_Treasury *TreasuryCallerSession) DepositableTokens(arg0 common.Address) (bool, error) {
	return _Treasury.Contract.DepositableTokens(&_Treasury.CallOpts, arg0)
}

// FundCost is a free data retrieval call binding the contract method 0xebbb2158.
//
// Solidity: function fundCost(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) FundCost(opts *bind.CallOpts, _supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "fundCost", _supply, _reserveBalance, _reserveRatio, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundCost is a free data retrieval call binding the contract method 0xebbb2158.
//
// Solidity: function fundCost(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) FundCost(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.FundCost(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// FundCost is a free data retrieval call binding the contract method 0xebbb2158.
//
// Solidity: function fundCost(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) FundCost(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.FundCost(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// FundSupplyAmount is a free data retrieval call binding the contract method 0x2f55bdb5.
//
// Solidity: function fundSupplyAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) FundSupplyAmount(opts *bind.CallOpts, _supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "fundSupplyAmount", _supply, _reserveBalance, _reserveRatio, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundSupplyAmount is a free data retrieval call binding the contract method 0x2f55bdb5.
//
// Solidity: function fundSupplyAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) FundSupplyAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.FundSupplyAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// FundSupplyAmount is a free data retrieval call binding the contract method 0x2f55bdb5.
//
// Solidity: function fundSupplyAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) FundSupplyAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.FundSupplyAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// LiquidateReserveAmount is a free data retrieval call binding the contract method 0x8074590a.
//
// Solidity: function liquidateReserveAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) LiquidateReserveAmount(opts *bind.CallOpts, _supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "liquidateReserveAmount", _supply, _reserveBalance, _reserveRatio, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidateReserveAmount is a free data retrieval call binding the contract method 0x8074590a.
//
// Solidity: function liquidateReserveAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) LiquidateReserveAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.LiquidateReserveAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// LiquidateReserveAmount is a free data retrieval call binding the contract method 0x8074590a.
//
// Solidity: function liquidateReserveAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveRatio, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) LiquidateReserveAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveRatio uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.LiquidateReserveAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveRatio, _amount)
}

// Mugen is a free data retrieval call binding the contract method 0x54dfbdf6.
//
// Solidity: function mugen() view returns(address)
func (_Treasury *TreasuryCaller) Mugen(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "mugen")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mugen is a free data retrieval call binding the contract method 0x54dfbdf6.
//
// Solidity: function mugen() view returns(address)
func (_Treasury *TreasurySession) Mugen() (common.Address, error) {
	return _Treasury.Contract.Mugen(&_Treasury.CallOpts)
}

// Mugen is a free data retrieval call binding the contract method 0x54dfbdf6.
//
// Solidity: function mugen() view returns(address)
func (_Treasury *TreasuryCallerSession) Mugen() (common.Address, error) {
	return _Treasury.Contract.Mugen(&_Treasury.CallOpts)
}

// NormalizedWeights is a free data retrieval call binding the contract method 0x76c2d510.
//
// Solidity: function normalizedWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasuryCaller) NormalizedWeights(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (uint32, uint32, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "normalizedWeights", _a, _b)

	if err != nil {
		return *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// NormalizedWeights is a free data retrieval call binding the contract method 0x76c2d510.
//
// Solidity: function normalizedWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasurySession) NormalizedWeights(_a *big.Int, _b *big.Int) (uint32, uint32, error) {
	return _Treasury.Contract.NormalizedWeights(&_Treasury.CallOpts, _a, _b)
}

// NormalizedWeights is a free data retrieval call binding the contract method 0x76c2d510.
//
// Solidity: function normalizedWeights(uint256 _a, uint256 _b) pure returns(uint32, uint32)
func (_Treasury *TreasuryCallerSession) NormalizedWeights(_a *big.Int, _b *big.Int) (uint32, uint32, error) {
	return _Treasury.Contract.NormalizedWeights(&_Treasury.CallOpts, _a, _b)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasurySession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasuryCallerSession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Treasury *TreasuryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Treasury *TreasurySession) Paused() (bool, error) {
	return _Treasury.Contract.Paused(&_Treasury.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Treasury *TreasuryCallerSession) Paused() (bool, error) {
	return _Treasury.Contract.Paused(&_Treasury.CallOpts)
}

// Power is a free data retrieval call binding the contract method 0x32833d51.
//
// Solidity: function power(uint256 _baseN, uint256 _baseD, uint32 _expN, uint32 _expD) view returns(uint256, uint8)
func (_Treasury *TreasuryCaller) Power(opts *bind.CallOpts, _baseN *big.Int, _baseD *big.Int, _expN uint32, _expD uint32) (*big.Int, uint8, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "power", _baseN, _baseD, _expN, _expD)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// Power is a free data retrieval call binding the contract method 0x32833d51.
//
// Solidity: function power(uint256 _baseN, uint256 _baseD, uint32 _expN, uint32 _expD) view returns(uint256, uint8)
func (_Treasury *TreasurySession) Power(_baseN *big.Int, _baseD *big.Int, _expN uint32, _expD uint32) (*big.Int, uint8, error) {
	return _Treasury.Contract.Power(&_Treasury.CallOpts, _baseN, _baseD, _expN, _expD)
}

// Power is a free data retrieval call binding the contract method 0x32833d51.
//
// Solidity: function power(uint256 _baseN, uint256 _baseD, uint32 _expN, uint32 _expD) view returns(uint256, uint8)
func (_Treasury *TreasuryCallerSession) Power(_baseN *big.Int, _baseD *big.Int, _expN uint32, _expD uint32) (*big.Int, uint8, error) {
	return _Treasury.Contract.Power(&_Treasury.CallOpts, _baseN, _baseD, _expN, _expD)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_Treasury *TreasuryCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_Treasury *TreasurySession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.PriceFeeds(&_Treasury.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_Treasury *TreasuryCallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.PriceFeeds(&_Treasury.CallOpts, arg0)
}

// PricePerToken is a free data retrieval call binding the contract method 0x7b1b1de6.
//
// Solidity: function pricePerToken() view returns(uint256)
func (_Treasury *TreasuryCaller) PricePerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "pricePerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerToken is a free data retrieval call binding the contract method 0x7b1b1de6.
//
// Solidity: function pricePerToken() view returns(uint256)
func (_Treasury *TreasurySession) PricePerToken() (*big.Int, error) {
	return _Treasury.Contract.PricePerToken(&_Treasury.CallOpts)
}

// PricePerToken is a free data retrieval call binding the contract method 0x7b1b1de6.
//
// Solidity: function pricePerToken() view returns(uint256)
func (_Treasury *TreasuryCallerSession) PricePerToken() (*big.Int, error) {
	return _Treasury.Contract.PricePerToken(&_Treasury.CallOpts)
}

// PurchaseTargetAmount is a free data retrieval call binding the contract method 0xf3250fe2.
//
// Solidity: function purchaseTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) PurchaseTargetAmount(opts *bind.CallOpts, _supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "purchaseTargetAmount", _supply, _reserveBalance, _reserveWeight, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PurchaseTargetAmount is a free data retrieval call binding the contract method 0xf3250fe2.
//
// Solidity: function purchaseTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) PurchaseTargetAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.PurchaseTargetAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveWeight, _amount)
}

// PurchaseTargetAmount is a free data retrieval call binding the contract method 0xf3250fe2.
//
// Solidity: function purchaseTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) PurchaseTargetAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.PurchaseTargetAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveWeight, _amount)
}

// ReadSupply is a free data retrieval call binding the contract method 0x6df137f6.
//
// Solidity: function readSupply() view returns(uint256)
func (_Treasury *TreasuryCaller) ReadSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "readSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadSupply is a free data retrieval call binding the contract method 0x6df137f6.
//
// Solidity: function readSupply() view returns(uint256)
func (_Treasury *TreasurySession) ReadSupply() (*big.Int, error) {
	return _Treasury.Contract.ReadSupply(&_Treasury.CallOpts)
}

// ReadSupply is a free data retrieval call binding the contract method 0x6df137f6.
//
// Solidity: function readSupply() view returns(uint256)
func (_Treasury *TreasuryCallerSession) ReadSupply() (*big.Int, error) {
	return _Treasury.Contract.ReadSupply(&_Treasury.CallOpts)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_Treasury *TreasuryCaller) ReserveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "reserveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_Treasury *TreasurySession) ReserveBalance() (*big.Int, error) {
	return _Treasury.Contract.ReserveBalance(&_Treasury.CallOpts)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_Treasury *TreasuryCallerSession) ReserveBalance() (*big.Int, error) {
	return _Treasury.Contract.ReserveBalance(&_Treasury.CallOpts)
}

// STotalSupply is a free data retrieval call binding the contract method 0x6a74a0fc.
//
// Solidity: function s_totalSupply() view returns(uint256)
func (_Treasury *TreasuryCaller) STotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "s_totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalSupply is a free data retrieval call binding the contract method 0x6a74a0fc.
//
// Solidity: function s_totalSupply() view returns(uint256)
func (_Treasury *TreasurySession) STotalSupply() (*big.Int, error) {
	return _Treasury.Contract.STotalSupply(&_Treasury.CallOpts)
}

// STotalSupply is a free data retrieval call binding the contract method 0x6a74a0fc.
//
// Solidity: function s_totalSupply() view returns(uint256)
func (_Treasury *TreasuryCallerSession) STotalSupply() (*big.Int, error) {
	return _Treasury.Contract.STotalSupply(&_Treasury.CallOpts)
}

// SafeFactors is a free data retrieval call binding the contract method 0xa67453b0.
//
// Solidity: function safeFactors(uint256 _a, uint256 _b) pure returns(uint256, uint256)
func (_Treasury *TreasuryCaller) SafeFactors(opts *bind.CallOpts, _a *big.Int, _b *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "safeFactors", _a, _b)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SafeFactors is a free data retrieval call binding the contract method 0xa67453b0.
//
// Solidity: function safeFactors(uint256 _a, uint256 _b) pure returns(uint256, uint256)
func (_Treasury *TreasurySession) SafeFactors(_a *big.Int, _b *big.Int) (*big.Int, *big.Int, error) {
	return _Treasury.Contract.SafeFactors(&_Treasury.CallOpts, _a, _b)
}

// SafeFactors is a free data retrieval call binding the contract method 0xa67453b0.
//
// Solidity: function safeFactors(uint256 _a, uint256 _b) pure returns(uint256, uint256)
func (_Treasury *TreasuryCallerSession) SafeFactors(_a *big.Int, _b *big.Int) (*big.Int, *big.Int, error) {
	return _Treasury.Contract.SafeFactors(&_Treasury.CallOpts, _a, _b)
}

// SaleTargetAmount is a free data retrieval call binding the contract method 0x76cf0b56.
//
// Solidity: function saleTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCaller) SaleTargetAmount(opts *bind.CallOpts, _supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "saleTargetAmount", _supply, _reserveBalance, _reserveWeight, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleTargetAmount is a free data retrieval call binding the contract method 0x76cf0b56.
//
// Solidity: function saleTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasurySession) SaleTargetAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.SaleTargetAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveWeight, _amount)
}

// SaleTargetAmount is a free data retrieval call binding the contract method 0x76cf0b56.
//
// Solidity: function saleTargetAmount(uint256 _supply, uint256 _reserveBalance, uint32 _reserveWeight, uint256 _amount) view returns(uint256)
func (_Treasury *TreasuryCallerSession) SaleTargetAmount(_supply *big.Int, _reserveBalance *big.Int, _reserveWeight uint32, _amount *big.Int) (*big.Int, error) {
	return _Treasury.Contract.SaleTargetAmount(&_Treasury.CallOpts, _supply, _reserveBalance, _reserveWeight, _amount)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Treasury *TreasuryCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Treasury *TreasurySession) Treasury() (common.Address, error) {
	return _Treasury.Contract.Treasury(&_Treasury.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Treasury *TreasuryCallerSession) Treasury() (common.Address, error) {
	return _Treasury.Contract.Treasury(&_Treasury.CallOpts)
}

// ValueDeposited is a free data retrieval call binding the contract method 0xe1d20a84.
//
// Solidity: function valueDeposited() view returns(uint256)
func (_Treasury *TreasuryCaller) ValueDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "valueDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueDeposited is a free data retrieval call binding the contract method 0xe1d20a84.
//
// Solidity: function valueDeposited() view returns(uint256)
func (_Treasury *TreasurySession) ValueDeposited() (*big.Int, error) {
	return _Treasury.Contract.ValueDeposited(&_Treasury.CallOpts)
}

// ValueDeposited is a free data retrieval call binding the contract method 0xe1d20a84.
//
// Solidity: function valueDeposited() view returns(uint256)
func (_Treasury *TreasuryCallerSession) ValueDeposited() (*big.Int, error) {
	return _Treasury.Contract.ValueDeposited(&_Treasury.CallOpts)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xf8d9c0ae.
//
// Solidity: function addTokenInfo(address _token, address _pricefeed) returns()
func (_Treasury *TreasuryTransactor) AddTokenInfo(opts *bind.TransactOpts, _token common.Address, _pricefeed common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "addTokenInfo", _token, _pricefeed)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xf8d9c0ae.
//
// Solidity: function addTokenInfo(address _token, address _pricefeed) returns()
func (_Treasury *TreasurySession) AddTokenInfo(_token common.Address, _pricefeed common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddTokenInfo(&_Treasury.TransactOpts, _token, _pricefeed)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xf8d9c0ae.
//
// Solidity: function addTokenInfo(address _token, address _pricefeed) returns()
func (_Treasury *TreasuryTransactorSession) AddTokenInfo(_token common.Address, _pricefeed common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddTokenInfo(&_Treasury.TransactOpts, _token, _pricefeed)
}

// ChangeTreasury is a paid mutator transaction binding the contract method 0xb14f2a39.
//
// Solidity: function changeTreasury(address _treasury) returns()
func (_Treasury *TreasuryTransactor) ChangeTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "changeTreasury", _treasury)
}

// ChangeTreasury is a paid mutator transaction binding the contract method 0xb14f2a39.
//
// Solidity: function changeTreasury(address _treasury) returns()
func (_Treasury *TreasurySession) ChangeTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.ChangeTreasury(&_Treasury.TransactOpts, _treasury)
}

// ChangeTreasury is a paid mutator transaction binding the contract method 0xb14f2a39.
//
// Solidity: function changeTreasury(address _treasury) returns()
func (_Treasury *TreasuryTransactorSession) ChangeTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.ChangeTreasury(&_Treasury.TransactOpts, _treasury)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Treasury *TreasuryTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Treasury *TreasurySession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Deposit(&_Treasury.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_Treasury *TreasuryTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Deposit(&_Treasury.TransactOpts, _token, _amount)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Treasury *TreasuryTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Treasury *TreasurySession) Init() (*types.Transaction, error) {
	return _Treasury.Contract.Init(&_Treasury.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Treasury *TreasuryTransactorSession) Init() (*types.Transaction, error) {
	return _Treasury.Contract.Init(&_Treasury.TransactOpts)
}

// InitMaxExpArray is a paid mutator transaction binding the contract method 0xaa248a46.
//
// Solidity: function initMaxExpArray() returns()
func (_Treasury *TreasuryTransactor) InitMaxExpArray(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "initMaxExpArray")
}

// InitMaxExpArray is a paid mutator transaction binding the contract method 0xaa248a46.
//
// Solidity: function initMaxExpArray() returns()
func (_Treasury *TreasurySession) InitMaxExpArray() (*types.Transaction, error) {
	return _Treasury.Contract.InitMaxExpArray(&_Treasury.TransactOpts)
}

// InitMaxExpArray is a paid mutator transaction binding the contract method 0xaa248a46.
//
// Solidity: function initMaxExpArray() returns()
func (_Treasury *TreasuryTransactorSession) InitMaxExpArray() (*types.Transaction, error) {
	return _Treasury.Contract.InitMaxExpArray(&_Treasury.TransactOpts)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_Treasury *TreasuryTransactor) PauseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "pauseDeposits")
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_Treasury *TreasurySession) PauseDeposits() (*types.Transaction, error) {
	return _Treasury.Contract.PauseDeposits(&_Treasury.TransactOpts)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_Treasury *TreasuryTransactorSession) PauseDeposits() (*types.Transaction, error) {
	return _Treasury.Contract.PauseDeposits(&_Treasury.TransactOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x072ac8be.
//
// Solidity: function receiveMessage(uint256 _amount) returns(uint256)
func (_Treasury *TreasuryTransactor) ReceiveMessage(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "receiveMessage", _amount)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x072ac8be.
//
// Solidity: function receiveMessage(uint256 _amount) returns(uint256)
func (_Treasury *TreasurySession) ReceiveMessage(_amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.ReceiveMessage(&_Treasury.TransactOpts, _amount)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x072ac8be.
//
// Solidity: function receiveMessage(uint256 _amount) returns(uint256)
func (_Treasury *TreasuryTransactorSession) ReceiveMessage(_amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.ReceiveMessage(&_Treasury.TransactOpts, _amount)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Treasury *TreasuryTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Treasury *TreasurySession) RemoveAdmin() (*types.Transaction, error) {
	return _Treasury.Contract.RemoveAdmin(&_Treasury.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Treasury *TreasuryTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Treasury.Contract.RemoveAdmin(&_Treasury.TransactOpts)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address _token) returns()
func (_Treasury *TreasuryTransactor) RemoveTokenInfo(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "removeTokenInfo", _token)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address _token) returns()
func (_Treasury *TreasurySession) RemoveTokenInfo(_token common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.RemoveTokenInfo(&_Treasury.TransactOpts, _token)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address _token) returns()
func (_Treasury *TreasuryTransactorSession) RemoveTokenInfo(_token common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.RemoveTokenInfo(&_Treasury.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasuryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasurySession) RenounceOwnership() (*types.Transaction, error) {
	return _Treasury.Contract.RenounceOwnership(&_Treasury.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasuryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Treasury.Contract.RenounceOwnership(&_Treasury.TransactOpts)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdmin) returns()
func (_Treasury *TreasuryTransactor) SetAdministrator(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setAdministrator", newAdmin)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdmin) returns()
func (_Treasury *TreasurySession) SetAdministrator(newAdmin common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetAdministrator(&_Treasury.TransactOpts, newAdmin)
}

// SetAdministrator is a paid mutator transaction binding the contract method 0xdf8089ef.
//
// Solidity: function setAdministrator(address newAdmin) returns()
func (_Treasury *TreasuryTransactorSession) SetAdministrator(newAdmin common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetAdministrator(&_Treasury.TransactOpts, newAdmin)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 _amount) returns()
func (_Treasury *TreasuryTransactor) SetCap(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setCap", _amount)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 _amount) returns()
func (_Treasury *TreasurySession) SetCap(_amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.SetCap(&_Treasury.TransactOpts, _amount)
}

// SetCap is a paid mutator transaction binding the contract method 0x47786d37.
//
// Solidity: function setCap(uint256 _amount) returns()
func (_Treasury *TreasuryTransactorSession) SetCap(_amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.SetCap(&_Treasury.TransactOpts, _amount)
}

// SetCommunicator is a paid mutator transaction binding the contract method 0x73101e2b.
//
// Solidity: function setCommunicator(address _comms) returns()
func (_Treasury *TreasuryTransactor) SetCommunicator(opts *bind.TransactOpts, _comms common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setCommunicator", _comms)
}

// SetCommunicator is a paid mutator transaction binding the contract method 0x73101e2b.
//
// Solidity: function setCommunicator(address _comms) returns()
func (_Treasury *TreasurySession) SetCommunicator(_comms common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetCommunicator(&_Treasury.TransactOpts, _comms)
}

// SetCommunicator is a paid mutator transaction binding the contract method 0x73101e2b.
//
// Solidity: function setCommunicator(address _comms) returns()
func (_Treasury *TreasuryTransactorSession) SetCommunicator(_comms common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetCommunicator(&_Treasury.TransactOpts, _comms)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasuryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasurySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.TransferOwnership(&_Treasury.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasuryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.TransferOwnership(&_Treasury.TransactOpts, newOwner)
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_Treasury *TreasuryTransactor) UnpauseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "unpauseDeposits")
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_Treasury *TreasurySession) UnpauseDeposits() (*types.Transaction, error) {
	return _Treasury.Contract.UnpauseDeposits(&_Treasury.TransactOpts)
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_Treasury *TreasuryTransactorSession) UnpauseDeposits() (*types.Transaction, error) {
	return _Treasury.Contract.UnpauseDeposits(&_Treasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Treasury *TreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Treasury *TreasurySession) Receive() (*types.Transaction, error) {
	return _Treasury.Contract.Receive(&_Treasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Treasury *TreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _Treasury.Contract.Receive(&_Treasury.TransactOpts)
}

// TreasuryDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Treasury contract.
type TreasuryDepositIterator struct {
	Event *TreasuryDeposit // Event containing the contract specifics and raw log

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
func (it *TreasuryDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryDeposit)
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
		it.Event = new(TreasuryDeposit)
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
func (it *TreasuryDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryDeposit represents a Deposit event raised by the Treasury contract.
type TreasuryDeposit struct {
	Depositor common.Address
	Token     common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed _depositor, address indexed _token, uint256 _value)
func (_Treasury *TreasuryFilterer) FilterDeposit(opts *bind.FilterOpts, _depositor []common.Address, _token []common.Address) (*TreasuryDepositIterator, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Deposit", _depositorRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryDepositIterator{contract: _Treasury.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed _depositor, address indexed _token, uint256 _value)
func (_Treasury *TreasuryFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TreasuryDeposit, _depositor []common.Address, _token []common.Address) (event.Subscription, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Deposit", _depositorRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryDeposit)
				if err := _Treasury.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed _depositor, address indexed _token, uint256 _value)
func (_Treasury *TreasuryFilterer) ParseDeposit(log types.Log) (*TreasuryDeposit, error) {
	event := new(TreasuryDeposit)
	if err := _Treasury.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryDepositableTokenIterator is returned from FilterDepositableToken and is used to iterate over the raw logs and unpacked data for DepositableToken events raised by the Treasury contract.
type TreasuryDepositableTokenIterator struct {
	Event *TreasuryDepositableToken // Event containing the contract specifics and raw log

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
func (it *TreasuryDepositableTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryDepositableToken)
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
		it.Event = new(TreasuryDepositableToken)
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
func (it *TreasuryDepositableTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryDepositableTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryDepositableToken represents a DepositableToken event raised by the Treasury contract.
type TreasuryDepositableToken struct {
	Token      common.Address
	PriceFreed common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositableToken is a free log retrieval operation binding the contract event 0x26086898d7fa4d8da4244a39784155e5561eb409d2397a7c23e2f315183a3190.
//
// Solidity: event DepositableToken(address indexed _token, address indexed _priceFreed)
func (_Treasury *TreasuryFilterer) FilterDepositableToken(opts *bind.FilterOpts, _token []common.Address, _priceFreed []common.Address) (*TreasuryDepositableTokenIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _priceFreedRule []interface{}
	for _, _priceFreedItem := range _priceFreed {
		_priceFreedRule = append(_priceFreedRule, _priceFreedItem)
	}

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "DepositableToken", _tokenRule, _priceFreedRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryDepositableTokenIterator{contract: _Treasury.contract, event: "DepositableToken", logs: logs, sub: sub}, nil
}

// WatchDepositableToken is a free log subscription operation binding the contract event 0x26086898d7fa4d8da4244a39784155e5561eb409d2397a7c23e2f315183a3190.
//
// Solidity: event DepositableToken(address indexed _token, address indexed _priceFreed)
func (_Treasury *TreasuryFilterer) WatchDepositableToken(opts *bind.WatchOpts, sink chan<- *TreasuryDepositableToken, _token []common.Address, _priceFreed []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _priceFreedRule []interface{}
	for _, _priceFreedItem := range _priceFreed {
		_priceFreedRule = append(_priceFreedRule, _priceFreedItem)
	}

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "DepositableToken", _tokenRule, _priceFreedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryDepositableToken)
				if err := _Treasury.contract.UnpackLog(event, "DepositableToken", log); err != nil {
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

// ParseDepositableToken is a log parse operation binding the contract event 0x26086898d7fa4d8da4244a39784155e5561eb409d2397a7c23e2f315183a3190.
//
// Solidity: event DepositableToken(address indexed _token, address indexed _priceFreed)
func (_Treasury *TreasuryFilterer) ParseDepositableToken(log types.Log) (*TreasuryDepositableToken, error) {
	event := new(TreasuryDepositableToken)
	if err := _Treasury.contract.UnpackLog(event, "DepositableToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Treasury contract.
type TreasuryOwnershipTransferredIterator struct {
	Event *TreasuryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TreasuryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryOwnershipTransferred)
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
		it.Event = new(TreasuryOwnershipTransferred)
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
func (it *TreasuryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryOwnershipTransferred represents a OwnershipTransferred event raised by the Treasury contract.
type TreasuryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Treasury *TreasuryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasuryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryOwnershipTransferredIterator{contract: _Treasury.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Treasury *TreasuryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasuryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryOwnershipTransferred)
				if err := _Treasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Treasury *TreasuryFilterer) ParseOwnershipTransferred(log types.Log) (*TreasuryOwnershipTransferred, error) {
	event := new(TreasuryOwnershipTransferred)
	if err := _Treasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Treasury contract.
type TreasuryPausedIterator struct {
	Event *TreasuryPaused // Event containing the contract specifics and raw log

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
func (it *TreasuryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryPaused)
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
		it.Event = new(TreasuryPaused)
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
func (it *TreasuryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryPaused represents a Paused event raised by the Treasury contract.
type TreasuryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Treasury *TreasuryFilterer) FilterPaused(opts *bind.FilterOpts) (*TreasuryPausedIterator, error) {

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TreasuryPausedIterator{contract: _Treasury.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Treasury *TreasuryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TreasuryPaused) (event.Subscription, error) {

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryPaused)
				if err := _Treasury.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Treasury *TreasuryFilterer) ParsePaused(log types.Log) (*TreasuryPaused, error) {
	event := new(TreasuryPaused)
	if err := _Treasury.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Treasury contract.
type TreasuryTokenRemovedIterator struct {
	Event *TreasuryTokenRemoved // Event containing the contract specifics and raw log

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
func (it *TreasuryTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryTokenRemoved)
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
		it.Event = new(TreasuryTokenRemoved)
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
func (it *TreasuryTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryTokenRemoved represents a TokenRemoved event raised by the Treasury contract.
type TreasuryTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _token)
func (_Treasury *TreasuryFilterer) FilterTokenRemoved(opts *bind.FilterOpts, _token []common.Address) (*TreasuryTokenRemovedIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "TokenRemoved", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryTokenRemovedIterator{contract: _Treasury.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _token)
func (_Treasury *TreasuryFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *TreasuryTokenRemoved, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "TokenRemoved", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryTokenRemoved)
				if err := _Treasury.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _token)
func (_Treasury *TreasuryFilterer) ParseTokenRemoved(log types.Log) (*TreasuryTokenRemoved, error) {
	event := new(TreasuryTokenRemoved)
	if err := _Treasury.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Treasury contract.
type TreasuryUnpausedIterator struct {
	Event *TreasuryUnpaused // Event containing the contract specifics and raw log

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
func (it *TreasuryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryUnpaused)
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
		it.Event = new(TreasuryUnpaused)
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
func (it *TreasuryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryUnpaused represents a Unpaused event raised by the Treasury contract.
type TreasuryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Treasury *TreasuryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TreasuryUnpausedIterator, error) {

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TreasuryUnpausedIterator{contract: _Treasury.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Treasury *TreasuryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TreasuryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryUnpaused)
				if err := _Treasury.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Treasury *TreasuryFilterer) ParseUnpaused(log types.Log) (*TreasuryUnpaused, error) {
	event := new(TreasuryUnpaused)
	if err := _Treasury.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
