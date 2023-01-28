// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Furnace

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

// FurnaceMetaData contains all meta data concerning the Furnace contract.
var FurnaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Advance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueUnderlying\",\"type\":\"uint256\"}],\"name\":\"Bond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CouponApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponsExpired\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessRedeemable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBonded\",\"type\":\"uint256\"}],\"name\":\"CouponExpiration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dollarAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"CouponPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"CouponRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CouponTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Incentivization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebt\",\"type\":\"uint256\"}],\"name\":\"SupplyDecrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRedeemable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lessDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBonded\",\"type\":\"uint256\"}],\"name\":\"SupplyIncrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"SupplyNeutral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueUnderlying\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"advance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"advanceIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowanceCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfBonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"balanceOfCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfStaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"bond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"bootstrappingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"couponPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"couponsExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dollar\",\"outputs\":[{\"internalType\":\"contractIDollar\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"expiringCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"expiringCouponsAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"fluidUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"outstandingCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dollarAmount\",\"type\":\"uint256\"}],\"name\":\"purchaseCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"couponEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\"}],\"name\":\"redeemCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIDollar\",\"name\":\"_dollar\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumAccount.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threeCRVPeg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"totalBondedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalCoupons\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRedeemable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unbond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unbondUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FurnaceABI is the input ABI used to generate the binding from.
// Deprecated: Use FurnaceMetaData.ABI instead.
var FurnaceABI = FurnaceMetaData.ABI

// Furnace is an auto generated Go binding around an Ethereum contract.
type Furnace struct {
	FurnaceCaller     // Read-only binding to the contract
	FurnaceTransactor // Write-only binding to the contract
	FurnaceFilterer   // Log filterer for contract events
}

// FurnaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FurnaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurnaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FurnaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurnaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FurnaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurnaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FurnaceSession struct {
	Contract     *Furnace          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FurnaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FurnaceCallerSession struct {
	Contract *FurnaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FurnaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FurnaceTransactorSession struct {
	Contract     *FurnaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FurnaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FurnaceRaw struct {
	Contract *Furnace // Generic contract binding to access the raw methods on
}

// FurnaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FurnaceCallerRaw struct {
	Contract *FurnaceCaller // Generic read-only contract binding to access the raw methods on
}

// FurnaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FurnaceTransactorRaw struct {
	Contract *FurnaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFurnace creates a new instance of Furnace, bound to a specific deployed contract.
func NewFurnace(address common.Address, backend bind.ContractBackend) (*Furnace, error) {
	contract, err := bindFurnace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Furnace{FurnaceCaller: FurnaceCaller{contract: contract}, FurnaceTransactor: FurnaceTransactor{contract: contract}, FurnaceFilterer: FurnaceFilterer{contract: contract}}, nil
}

// NewFurnaceCaller creates a new read-only instance of Furnace, bound to a specific deployed contract.
func NewFurnaceCaller(address common.Address, caller bind.ContractCaller) (*FurnaceCaller, error) {
	contract, err := bindFurnace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FurnaceCaller{contract: contract}, nil
}

// NewFurnaceTransactor creates a new write-only instance of Furnace, bound to a specific deployed contract.
func NewFurnaceTransactor(address common.Address, transactor bind.ContractTransactor) (*FurnaceTransactor, error) {
	contract, err := bindFurnace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FurnaceTransactor{contract: contract}, nil
}

// NewFurnaceFilterer creates a new log filterer instance of Furnace, bound to a specific deployed contract.
func NewFurnaceFilterer(address common.Address, filterer bind.ContractFilterer) (*FurnaceFilterer, error) {
	contract, err := bindFurnace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FurnaceFilterer{contract: contract}, nil
}

// bindFurnace binds a generic wrapper to an already deployed contract.
func bindFurnace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FurnaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Furnace *FurnaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Furnace.Contract.FurnaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Furnace *FurnaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furnace.Contract.FurnaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Furnace *FurnaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Furnace.Contract.FurnaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Furnace *FurnaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Furnace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Furnace *FurnaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furnace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Furnace *FurnaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Furnace.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Furnace.Contract.Allowance(&_Furnace.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Furnace.Contract.Allowance(&_Furnace.CallOpts, owner, spender)
}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceCaller) AllowanceCoupons(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "allowanceCoupons", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceSession) AllowanceCoupons(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Furnace.Contract.AllowanceCoupons(&_Furnace.CallOpts, owner, spender)
}

// AllowanceCoupons is a free data retrieval call binding the contract method 0x9f6e1b26.
//
// Solidity: function allowanceCoupons(address owner, address spender) view returns(uint256)
func (_Furnace *FurnaceCallerSession) AllowanceCoupons(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Furnace.Contract.AllowanceCoupons(&_Furnace.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Furnace *FurnaceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Furnace *FurnaceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOf(&_Furnace.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Furnace *FurnaceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOf(&_Furnace.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Furnace *FurnaceCaller) BalanceOfBonded(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "balanceOfBonded", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Furnace *FurnaceSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfBonded(&_Furnace.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Furnace *FurnaceCallerSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfBonded(&_Furnace.CallOpts, account)
}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCaller) BalanceOfCoupons(opts *bind.CallOpts, account common.Address, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "balanceOfCoupons", account, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceSession) BalanceOfCoupons(account common.Address, epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfCoupons(&_Furnace.CallOpts, account, epoch)
}

// BalanceOfCoupons is a free data retrieval call binding the contract method 0xbc7513e2.
//
// Solidity: function balanceOfCoupons(address account, uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCallerSession) BalanceOfCoupons(account common.Address, epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfCoupons(&_Furnace.CallOpts, account, epoch)
}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Furnace *FurnaceCaller) BalanceOfStaged(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "balanceOfStaged", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Furnace *FurnaceSession) BalanceOfStaged(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfStaged(&_Furnace.CallOpts, account)
}

// BalanceOfStaged is a free data retrieval call binding the contract method 0x86cf9f14.
//
// Solidity: function balanceOfStaged(address account) view returns(uint256)
func (_Furnace *FurnaceCallerSession) BalanceOfStaged(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.BalanceOfStaged(&_Furnace.CallOpts, account)
}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Furnace *FurnaceCaller) BootstrappingAt(opts *bind.CallOpts, epoch *big.Int) (bool, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "bootstrappingAt", epoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Furnace *FurnaceSession) BootstrappingAt(epoch *big.Int) (bool, error) {
	return _Furnace.Contract.BootstrappingAt(&_Furnace.CallOpts, epoch)
}

// BootstrappingAt is a free data retrieval call binding the contract method 0x75d5024b.
//
// Solidity: function bootstrappingAt(uint256 epoch) view returns(bool)
func (_Furnace *FurnaceCallerSession) BootstrappingAt(epoch *big.Int) (bool, error) {
	return _Furnace.Contract.BootstrappingAt(&_Furnace.CallOpts, epoch)
}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Furnace *FurnaceCaller) CouponPremium(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "couponPremium", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Furnace *FurnaceSession) CouponPremium(amount *big.Int) (*big.Int, error) {
	return _Furnace.Contract.CouponPremium(&_Furnace.CallOpts, amount)
}

// CouponPremium is a free data retrieval call binding the contract method 0xd8f54138.
//
// Solidity: function couponPremium(uint256 amount) view returns(uint256)
func (_Furnace *FurnaceCallerSession) CouponPremium(amount *big.Int) (*big.Int, error) {
	return _Furnace.Contract.CouponPremium(&_Furnace.CallOpts, amount)
}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCaller) CouponsExpiration(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "couponsExpiration", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceSession) CouponsExpiration(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.CouponsExpiration(&_Furnace.CallOpts, epoch)
}

// CouponsExpiration is a free data retrieval call binding the contract method 0x10e95b6c.
//
// Solidity: function couponsExpiration(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCallerSession) CouponsExpiration(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.CouponsExpiration(&_Furnace.CallOpts, epoch)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Furnace *FurnaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Furnace *FurnaceSession) Decimals() (uint8, error) {
	return _Furnace.Contract.Decimals(&_Furnace.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Furnace *FurnaceCallerSession) Decimals() (uint8, error) {
	return _Furnace.Contract.Decimals(&_Furnace.CallOpts)
}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Furnace *FurnaceCaller) Dollar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "dollar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Furnace *FurnaceSession) Dollar() (common.Address, error) {
	return _Furnace.Contract.Dollar(&_Furnace.CallOpts)
}

// Dollar is a free data retrieval call binding the contract method 0x51adeb57.
//
// Solidity: function dollar() view returns(address)
func (_Furnace *FurnaceCallerSession) Dollar() (common.Address, error) {
	return _Furnace.Contract.Dollar(&_Furnace.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Furnace *FurnaceCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Furnace *FurnaceSession) Epoch() (*big.Int, error) {
	return _Furnace.Contract.Epoch(&_Furnace.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Furnace *FurnaceCallerSession) Epoch() (*big.Int, error) {
	return _Furnace.Contract.Epoch(&_Furnace.CallOpts)
}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Furnace *FurnaceCaller) EpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "epochTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Furnace *FurnaceSession) EpochTime() (*big.Int, error) {
	return _Furnace.Contract.EpochTime(&_Furnace.CallOpts)
}

// EpochTime is a free data retrieval call binding the contract method 0x5053e461.
//
// Solidity: function epochTime() view returns(uint256)
func (_Furnace *FurnaceCallerSession) EpochTime() (*big.Int, error) {
	return _Furnace.Contract.EpochTime(&_Furnace.CallOpts)
}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCaller) ExpiringCoupons(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "expiringCoupons", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceSession) ExpiringCoupons(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.ExpiringCoupons(&_Furnace.CallOpts, epoch)
}

// ExpiringCoupons is a free data retrieval call binding the contract method 0x6a39e328.
//
// Solidity: function expiringCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCallerSession) ExpiringCoupons(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.ExpiringCoupons(&_Furnace.CallOpts, epoch)
}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Furnace *FurnaceCaller) ExpiringCouponsAtIndex(opts *bind.CallOpts, epoch *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "expiringCouponsAtIndex", epoch, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Furnace *FurnaceSession) ExpiringCouponsAtIndex(epoch *big.Int, i *big.Int) (*big.Int, error) {
	return _Furnace.Contract.ExpiringCouponsAtIndex(&_Furnace.CallOpts, epoch, i)
}

// ExpiringCouponsAtIndex is a free data retrieval call binding the contract method 0x4c736099.
//
// Solidity: function expiringCouponsAtIndex(uint256 epoch, uint256 i) view returns(uint256)
func (_Furnace *FurnaceCallerSession) ExpiringCouponsAtIndex(epoch *big.Int, i *big.Int) (*big.Int, error) {
	return _Furnace.Contract.ExpiringCouponsAtIndex(&_Furnace.CallOpts, epoch, i)
}

// FluidUntil is a free data retrieval call binding the contract method 0x51bf21d8.
//
// Solidity: function fluidUntil(address account) view returns(uint256)
func (_Furnace *FurnaceCaller) FluidUntil(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "fluidUntil", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FluidUntil is a free data retrieval call binding the contract method 0x51bf21d8.
//
// Solidity: function fluidUntil(address account) view returns(uint256)
func (_Furnace *FurnaceSession) FluidUntil(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.FluidUntil(&_Furnace.CallOpts, account)
}

// FluidUntil is a free data retrieval call binding the contract method 0x51bf21d8.
//
// Solidity: function fluidUntil(address account) view returns(uint256)
func (_Furnace *FurnaceCallerSession) FluidUntil(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.FluidUntil(&_Furnace.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Furnace *FurnaceCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Furnace *FurnaceSession) IsOwner() (bool, error) {
	return _Furnace.Contract.IsOwner(&_Furnace.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Furnace *FurnaceCallerSession) IsOwner() (bool, error) {
	return _Furnace.Contract.IsOwner(&_Furnace.CallOpts)
}

// LockedUntil is a free data retrieval call binding the contract method 0x9bc289f1.
//
// Solidity: function lockedUntil(address account) view returns(uint256)
func (_Furnace *FurnaceCaller) LockedUntil(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "lockedUntil", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedUntil is a free data retrieval call binding the contract method 0x9bc289f1.
//
// Solidity: function lockedUntil(address account) view returns(uint256)
func (_Furnace *FurnaceSession) LockedUntil(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.LockedUntil(&_Furnace.CallOpts, account)
}

// LockedUntil is a free data retrieval call binding the contract method 0x9bc289f1.
//
// Solidity: function lockedUntil(address account) view returns(uint256)
func (_Furnace *FurnaceCallerSession) LockedUntil(account common.Address) (*big.Int, error) {
	return _Furnace.Contract.LockedUntil(&_Furnace.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Furnace *FurnaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Furnace *FurnaceSession) Name() (string, error) {
	return _Furnace.Contract.Name(&_Furnace.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Furnace *FurnaceCallerSession) Name() (string, error) {
	return _Furnace.Contract.Name(&_Furnace.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Furnace *FurnaceCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Furnace *FurnaceSession) Oracle() (common.Address, error) {
	return _Furnace.Contract.Oracle(&_Furnace.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Furnace *FurnaceCallerSession) Oracle() (common.Address, error) {
	return _Furnace.Contract.Oracle(&_Furnace.CallOpts)
}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCaller) OutstandingCoupons(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "outstandingCoupons", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceSession) OutstandingCoupons(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.OutstandingCoupons(&_Furnace.CallOpts, epoch)
}

// OutstandingCoupons is a free data retrieval call binding the contract method 0xc9aff70c.
//
// Solidity: function outstandingCoupons(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCallerSession) OutstandingCoupons(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.OutstandingCoupons(&_Furnace.CallOpts, epoch)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Furnace *FurnaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Furnace *FurnaceSession) Owner() (common.Address, error) {
	return _Furnace.Contract.Owner(&_Furnace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Furnace *FurnaceCallerSession) Owner() (common.Address, error) {
	return _Furnace.Contract.Owner(&_Furnace.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Furnace *FurnaceCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Furnace *FurnaceSession) Pool() (common.Address, error) {
	return _Furnace.Contract.Pool(&_Furnace.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Furnace *FurnaceCallerSession) Pool() (common.Address, error) {
	return _Furnace.Contract.Pool(&_Furnace.CallOpts)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Furnace *FurnaceCaller) StatusOf(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "statusOf", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Furnace *FurnaceSession) StatusOf(account common.Address) (uint8, error) {
	return _Furnace.Contract.StatusOf(&_Furnace.CallOpts, account)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(address account) view returns(uint8)
func (_Furnace *FurnaceCallerSession) StatusOf(account common.Address) (uint8, error) {
	return _Furnace.Contract.StatusOf(&_Furnace.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Furnace *FurnaceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Furnace *FurnaceSession) Symbol() (string, error) {
	return _Furnace.Contract.Symbol(&_Furnace.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Furnace *FurnaceCallerSession) Symbol() (string, error) {
	return _Furnace.Contract.Symbol(&_Furnace.CallOpts)
}

// ThreeCRVPeg is a free data retrieval call binding the contract method 0xbce98dba.
//
// Solidity: function threeCRVPeg() view returns(uint256 value)
func (_Furnace *FurnaceCaller) ThreeCRVPeg(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "threeCRVPeg")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThreeCRVPeg is a free data retrieval call binding the contract method 0xbce98dba.
//
// Solidity: function threeCRVPeg() view returns(uint256 value)
func (_Furnace *FurnaceSession) ThreeCRVPeg() (*big.Int, error) {
	return _Furnace.Contract.ThreeCRVPeg(&_Furnace.CallOpts)
}

// ThreeCRVPeg is a free data retrieval call binding the contract method 0xbce98dba.
//
// Solidity: function threeCRVPeg() view returns(uint256 value)
func (_Furnace *FurnaceCallerSession) ThreeCRVPeg() (*big.Int, error) {
	return _Furnace.Contract.ThreeCRVPeg(&_Furnace.CallOpts)
}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalBonded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalBonded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Furnace *FurnaceSession) TotalBonded() (*big.Int, error) {
	return _Furnace.Contract.TotalBonded(&_Furnace.CallOpts)
}

// TotalBonded is a free data retrieval call binding the contract method 0x44d96e95.
//
// Solidity: function totalBonded() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalBonded() (*big.Int, error) {
	return _Furnace.Contract.TotalBonded(&_Furnace.CallOpts)
}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCaller) TotalBondedAt(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalBondedAt", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceSession) TotalBondedAt(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.TotalBondedAt(&_Furnace.CallOpts, epoch)
}

// TotalBondedAt is a free data retrieval call binding the contract method 0xffbe3b73.
//
// Solidity: function totalBondedAt(uint256 epoch) view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalBondedAt(epoch *big.Int) (*big.Int, error) {
	return _Furnace.Contract.TotalBondedAt(&_Furnace.CallOpts, epoch)
}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalCoupons(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalCoupons")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Furnace *FurnaceSession) TotalCoupons() (*big.Int, error) {
	return _Furnace.Contract.TotalCoupons(&_Furnace.CallOpts)
}

// TotalCoupons is a free data retrieval call binding the contract method 0x9a649edc.
//
// Solidity: function totalCoupons() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalCoupons() (*big.Int, error) {
	return _Furnace.Contract.TotalCoupons(&_Furnace.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Furnace *FurnaceSession) TotalDebt() (*big.Int, error) {
	return _Furnace.Contract.TotalDebt(&_Furnace.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalDebt() (*big.Int, error) {
	return _Furnace.Contract.TotalDebt(&_Furnace.CallOpts)
}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalNet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalNet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Furnace *FurnaceSession) TotalNet() (*big.Int, error) {
	return _Furnace.Contract.TotalNet(&_Furnace.CallOpts)
}

// TotalNet is a free data retrieval call binding the contract method 0xa6c409f1.
//
// Solidity: function totalNet() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalNet() (*big.Int, error) {
	return _Furnace.Contract.TotalNet(&_Furnace.CallOpts)
}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalRedeemable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalRedeemable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Furnace *FurnaceSession) TotalRedeemable() (*big.Int, error) {
	return _Furnace.Contract.TotalRedeemable(&_Furnace.CallOpts)
}

// TotalRedeemable is a free data retrieval call binding the contract method 0x1edbcf6c.
//
// Solidity: function totalRedeemable() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalRedeemable() (*big.Int, error) {
	return _Furnace.Contract.TotalRedeemable(&_Furnace.CallOpts)
}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalStaged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalStaged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Furnace *FurnaceSession) TotalStaged() (*big.Int, error) {
	return _Furnace.Contract.TotalStaged(&_Furnace.CallOpts)
}

// TotalStaged is a free data retrieval call binding the contract method 0xcf023779.
//
// Solidity: function totalStaged() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalStaged() (*big.Int, error) {
	return _Furnace.Contract.TotalStaged(&_Furnace.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Furnace *FurnaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Furnace.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Furnace *FurnaceSession) TotalSupply() (*big.Int, error) {
	return _Furnace.Contract.TotalSupply(&_Furnace.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Furnace *FurnaceCallerSession) TotalSupply() (*big.Int, error) {
	return _Furnace.Contract.TotalSupply(&_Furnace.CallOpts)
}

// Advance is a paid mutator transaction binding the contract method 0x71d6ddd6.
//
// Solidity: function advance(uint256 key) returns()
func (_Furnace *FurnaceTransactor) Advance(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "advance", key)
}

// Advance is a paid mutator transaction binding the contract method 0x71d6ddd6.
//
// Solidity: function advance(uint256 key) returns()
func (_Furnace *FurnaceSession) Advance(key *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Advance(&_Furnace.TransactOpts, key)
}

// Advance is a paid mutator transaction binding the contract method 0x71d6ddd6.
//
// Solidity: function advance(uint256 key) returns()
func (_Furnace *FurnaceTransactorSession) Advance(key *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Advance(&_Furnace.TransactOpts, key)
}

// AdvanceIncentive is a paid mutator transaction binding the contract method 0xb5702cfb.
//
// Solidity: function advanceIncentive() returns(uint256)
func (_Furnace *FurnaceTransactor) AdvanceIncentive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "advanceIncentive")
}

// AdvanceIncentive is a paid mutator transaction binding the contract method 0xb5702cfb.
//
// Solidity: function advanceIncentive() returns(uint256)
func (_Furnace *FurnaceSession) AdvanceIncentive() (*types.Transaction, error) {
	return _Furnace.Contract.AdvanceIncentive(&_Furnace.TransactOpts)
}

// AdvanceIncentive is a paid mutator transaction binding the contract method 0xb5702cfb.
//
// Solidity: function advanceIncentive() returns(uint256)
func (_Furnace *FurnaceTransactorSession) AdvanceIncentive() (*types.Transaction, error) {
	return _Furnace.Contract.AdvanceIncentive(&_Furnace.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Furnace *FurnaceSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Approve(&_Furnace.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Approve(&_Furnace.TransactOpts, spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Furnace *FurnaceTransactor) ApproveCoupons(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "approveCoupons", spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Furnace *FurnaceSession) ApproveCoupons(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.ApproveCoupons(&_Furnace.TransactOpts, spender, amount)
}

// ApproveCoupons is a paid mutator transaction binding the contract method 0x2f7f889e.
//
// Solidity: function approveCoupons(address spender, uint256 amount) returns()
func (_Furnace *FurnaceTransactorSession) ApproveCoupons(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.ApproveCoupons(&_Furnace.TransactOpts, spender, amount)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Furnace *FurnaceTransactor) Bond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "bond", value)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Furnace *FurnaceSession) Bond(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Bond(&_Furnace.TransactOpts, value)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 value) returns()
func (_Furnace *FurnaceTransactorSession) Bond(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Bond(&_Furnace.TransactOpts, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Furnace *FurnaceTransactor) Deposit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "deposit", value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Furnace *FurnaceSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Deposit(&_Furnace.TransactOpts, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Furnace *FurnaceTransactorSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Deposit(&_Furnace.TransactOpts, value)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Furnace *FurnaceTransactor) PurchaseCoupons(opts *bind.TransactOpts, dollarAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "purchaseCoupons", dollarAmount)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Furnace *FurnaceSession) PurchaseCoupons(dollarAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.PurchaseCoupons(&_Furnace.TransactOpts, dollarAmount)
}

// PurchaseCoupons is a paid mutator transaction binding the contract method 0xe5f55c7e.
//
// Solidity: function purchaseCoupons(uint256 dollarAmount) returns(uint256)
func (_Furnace *FurnaceTransactorSession) PurchaseCoupons(dollarAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.PurchaseCoupons(&_Furnace.TransactOpts, dollarAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Furnace *FurnaceTransactor) RedeemCoupons(opts *bind.TransactOpts, couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "redeemCoupons", couponEpoch, couponAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Furnace *FurnaceSession) RedeemCoupons(couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.RedeemCoupons(&_Furnace.TransactOpts, couponEpoch, couponAmount)
}

// RedeemCoupons is a paid mutator transaction binding the contract method 0xd6a9cf08.
//
// Solidity: function redeemCoupons(uint256 couponEpoch, uint256 couponAmount) returns()
func (_Furnace *FurnaceTransactorSession) RedeemCoupons(couponEpoch *big.Int, couponAmount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.RedeemCoupons(&_Furnace.TransactOpts, couponEpoch, couponAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Furnace *FurnaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Furnace *FurnaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Furnace.Contract.RenounceOwnership(&_Furnace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Furnace *FurnaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Furnace.Contract.RenounceOwnership(&_Furnace.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x77b8b1c7.
//
// Solidity: function setup(address _dollar, address _oracle, address _pool) returns()
func (_Furnace *FurnaceTransactor) Setup(opts *bind.TransactOpts, _dollar common.Address, _oracle common.Address, _pool common.Address) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "setup", _dollar, _oracle, _pool)
}

// Setup is a paid mutator transaction binding the contract method 0x77b8b1c7.
//
// Solidity: function setup(address _dollar, address _oracle, address _pool) returns()
func (_Furnace *FurnaceSession) Setup(_dollar common.Address, _oracle common.Address, _pool common.Address) (*types.Transaction, error) {
	return _Furnace.Contract.Setup(&_Furnace.TransactOpts, _dollar, _oracle, _pool)
}

// Setup is a paid mutator transaction binding the contract method 0x77b8b1c7.
//
// Solidity: function setup(address _dollar, address _oracle, address _pool) returns()
func (_Furnace *FurnaceTransactorSession) Setup(_dollar common.Address, _oracle common.Address, _pool common.Address) (*types.Transaction, error) {
	return _Furnace.Contract.Setup(&_Furnace.TransactOpts, _dollar, _oracle, _pool)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Transfer(&_Furnace.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Transfer(&_Furnace.TransactOpts, recipient, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Furnace *FurnaceTransactor) TransferCoupons(opts *bind.TransactOpts, sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "transferCoupons", sender, recipient, epoch, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Furnace *FurnaceSession) TransferCoupons(sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.TransferCoupons(&_Furnace.TransactOpts, sender, recipient, epoch, amount)
}

// TransferCoupons is a paid mutator transaction binding the contract method 0x005edd37.
//
// Solidity: function transferCoupons(address sender, address recipient, uint256 epoch, uint256 amount) returns()
func (_Furnace *FurnaceTransactorSession) TransferCoupons(sender common.Address, recipient common.Address, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.TransferCoupons(&_Furnace.TransactOpts, sender, recipient, epoch, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.TransferFrom(&_Furnace.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Furnace *FurnaceTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.TransferFrom(&_Furnace.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Furnace *FurnaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Furnace *FurnaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Furnace.Contract.TransferOwnership(&_Furnace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Furnace *FurnaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Furnace.Contract.TransferOwnership(&_Furnace.TransactOpts, newOwner)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Furnace *FurnaceTransactor) Unbond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "unbond", value)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Furnace *FurnaceSession) Unbond(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Unbond(&_Furnace.TransactOpts, value)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 value) returns()
func (_Furnace *FurnaceTransactorSession) Unbond(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Unbond(&_Furnace.TransactOpts, value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Furnace *FurnaceTransactor) UnbondUnderlying(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "unbondUnderlying", value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Furnace *FurnaceSession) UnbondUnderlying(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.UnbondUnderlying(&_Furnace.TransactOpts, value)
}

// UnbondUnderlying is a paid mutator transaction binding the contract method 0xdf9a2b1c.
//
// Solidity: function unbondUnderlying(uint256 value) returns()
func (_Furnace *FurnaceTransactorSession) UnbondUnderlying(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.UnbondUnderlying(&_Furnace.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Furnace *FurnaceTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Furnace.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Furnace *FurnaceSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Withdraw(&_Furnace.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Furnace *FurnaceTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Furnace.Contract.Withdraw(&_Furnace.TransactOpts, value)
}

// FurnaceAdvanceIterator is returned from FilterAdvance and is used to iterate over the raw logs and unpacked data for Advance events raised by the Furnace contract.
type FurnaceAdvanceIterator struct {
	Event *FurnaceAdvance // Event containing the contract specifics and raw log

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
func (it *FurnaceAdvanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceAdvance)
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
		it.Event = new(FurnaceAdvance)
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
func (it *FurnaceAdvanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceAdvanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceAdvance represents a Advance event raised by the Furnace contract.
type FurnaceAdvance struct {
	Epoch     *big.Int
	Block     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAdvance is a free log retrieval operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Furnace *FurnaceFilterer) FilterAdvance(opts *bind.FilterOpts, epoch []*big.Int) (*FurnaceAdvanceIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Advance", epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceAdvanceIterator{contract: _Furnace.contract, event: "Advance", logs: logs, sub: sub}, nil
}

// WatchAdvance is a free log subscription operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Furnace *FurnaceFilterer) WatchAdvance(opts *bind.WatchOpts, sink chan<- *FurnaceAdvance, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Advance", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceAdvance)
				if err := _Furnace.contract.UnpackLog(event, "Advance", log); err != nil {
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

// ParseAdvance is a log parse operation binding the contract event 0xc30b728d1c19e5db3678b8ea9e9a063a5655071e1a325c2f7fdbca48baa90600.
//
// Solidity: event Advance(uint256 indexed epoch, uint256 block, uint256 timestamp)
func (_Furnace *FurnaceFilterer) ParseAdvance(log types.Log) (*FurnaceAdvance, error) {
	event := new(FurnaceAdvance)
	if err := _Furnace.contract.UnpackLog(event, "Advance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceBondIterator is returned from FilterBond and is used to iterate over the raw logs and unpacked data for Bond events raised by the Furnace contract.
type FurnaceBondIterator struct {
	Event *FurnaceBond // Event containing the contract specifics and raw log

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
func (it *FurnaceBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceBond)
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
		it.Event = new(FurnaceBond)
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
func (it *FurnaceBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceBond represents a Bond event raised by the Furnace contract.
type FurnaceBond struct {
	Account         common.Address
	Start           *big.Int
	Value           *big.Int
	ValueUnderlying *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBond is a free log retrieval operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) FilterBond(opts *bind.FilterOpts, account []common.Address) (*FurnaceBondIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Bond", accountRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceBondIterator{contract: _Furnace.contract, event: "Bond", logs: logs, sub: sub}, nil
}

// WatchBond is a free log subscription operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) WatchBond(opts *bind.WatchOpts, sink chan<- *FurnaceBond, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Bond", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceBond)
				if err := _Furnace.contract.UnpackLog(event, "Bond", log); err != nil {
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

// ParseBond is a log parse operation binding the contract event 0x44002fdef5a0c2d2e4e05572e9780b95aef97e0e93ffd7cc076b09fa78ff2b46.
//
// Solidity: event Bond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) ParseBond(log types.Log) (*FurnaceBond, error) {
	event := new(FurnaceBond)
	if err := _Furnace.contract.UnpackLog(event, "Bond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceCouponApprovalIterator is returned from FilterCouponApproval and is used to iterate over the raw logs and unpacked data for CouponApproval events raised by the Furnace contract.
type FurnaceCouponApprovalIterator struct {
	Event *FurnaceCouponApproval // Event containing the contract specifics and raw log

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
func (it *FurnaceCouponApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceCouponApproval)
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
		it.Event = new(FurnaceCouponApproval)
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
func (it *FurnaceCouponApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceCouponApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceCouponApproval represents a CouponApproval event raised by the Furnace contract.
type FurnaceCouponApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCouponApproval is a free log retrieval operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Furnace *FurnaceFilterer) FilterCouponApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FurnaceCouponApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "CouponApproval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceCouponApprovalIterator{contract: _Furnace.contract, event: "CouponApproval", logs: logs, sub: sub}, nil
}

// WatchCouponApproval is a free log subscription operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Furnace *FurnaceFilterer) WatchCouponApproval(opts *bind.WatchOpts, sink chan<- *FurnaceCouponApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "CouponApproval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceCouponApproval)
				if err := _Furnace.contract.UnpackLog(event, "CouponApproval", log); err != nil {
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

// ParseCouponApproval is a log parse operation binding the contract event 0x8ff27e6b95060c1ca851e7c2c28af8b413eb1a8bcb637b0290da9543a709cce3.
//
// Solidity: event CouponApproval(address indexed owner, address indexed spender, uint256 value)
func (_Furnace *FurnaceFilterer) ParseCouponApproval(log types.Log) (*FurnaceCouponApproval, error) {
	event := new(FurnaceCouponApproval)
	if err := _Furnace.contract.UnpackLog(event, "CouponApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceCouponExpirationIterator is returned from FilterCouponExpiration and is used to iterate over the raw logs and unpacked data for CouponExpiration events raised by the Furnace contract.
type FurnaceCouponExpirationIterator struct {
	Event *FurnaceCouponExpiration // Event containing the contract specifics and raw log

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
func (it *FurnaceCouponExpirationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceCouponExpiration)
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
		it.Event = new(FurnaceCouponExpiration)
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
func (it *FurnaceCouponExpirationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceCouponExpirationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceCouponExpiration represents a CouponExpiration event raised by the Furnace contract.
type FurnaceCouponExpiration struct {
	Epoch          *big.Int
	CouponsExpired *big.Int
	LessRedeemable *big.Int
	LessDebt       *big.Int
	NewBonded      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCouponExpiration is a free log retrieval operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) FilterCouponExpiration(opts *bind.FilterOpts, epoch []*big.Int) (*FurnaceCouponExpirationIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "CouponExpiration", epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceCouponExpirationIterator{contract: _Furnace.contract, event: "CouponExpiration", logs: logs, sub: sub}, nil
}

// WatchCouponExpiration is a free log subscription operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) WatchCouponExpiration(opts *bind.WatchOpts, sink chan<- *FurnaceCouponExpiration, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "CouponExpiration", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceCouponExpiration)
				if err := _Furnace.contract.UnpackLog(event, "CouponExpiration", log); err != nil {
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

// ParseCouponExpiration is a log parse operation binding the contract event 0x753df65b37159bf237ae1fca97ba1bd57cf83bc9498f271a514a4d7bafe87bda.
//
// Solidity: event CouponExpiration(uint256 indexed epoch, uint256 couponsExpired, uint256 lessRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) ParseCouponExpiration(log types.Log) (*FurnaceCouponExpiration, error) {
	event := new(FurnaceCouponExpiration)
	if err := _Furnace.contract.UnpackLog(event, "CouponExpiration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceCouponPurchaseIterator is returned from FilterCouponPurchase and is used to iterate over the raw logs and unpacked data for CouponPurchase events raised by the Furnace contract.
type FurnaceCouponPurchaseIterator struct {
	Event *FurnaceCouponPurchase // Event containing the contract specifics and raw log

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
func (it *FurnaceCouponPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceCouponPurchase)
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
		it.Event = new(FurnaceCouponPurchase)
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
func (it *FurnaceCouponPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceCouponPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceCouponPurchase represents a CouponPurchase event raised by the Furnace contract.
type FurnaceCouponPurchase struct {
	Account      common.Address
	Epoch        *big.Int
	DollarAmount *big.Int
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCouponPurchase is a free log retrieval operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) FilterCouponPurchase(opts *bind.FilterOpts, account []common.Address, epoch []*big.Int) (*FurnaceCouponPurchaseIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "CouponPurchase", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceCouponPurchaseIterator{contract: _Furnace.contract, event: "CouponPurchase", logs: logs, sub: sub}, nil
}

// WatchCouponPurchase is a free log subscription operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) WatchCouponPurchase(opts *bind.WatchOpts, sink chan<- *FurnaceCouponPurchase, account []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "CouponPurchase", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceCouponPurchase)
				if err := _Furnace.contract.UnpackLog(event, "CouponPurchase", log); err != nil {
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

// ParseCouponPurchase is a log parse operation binding the contract event 0xbce252db29f761f815dc2e9ea60005af065efba6eb619d2a0b2a113fdeb61414.
//
// Solidity: event CouponPurchase(address indexed account, uint256 indexed epoch, uint256 dollarAmount, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) ParseCouponPurchase(log types.Log) (*FurnaceCouponPurchase, error) {
	event := new(FurnaceCouponPurchase)
	if err := _Furnace.contract.UnpackLog(event, "CouponPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceCouponRedemptionIterator is returned from FilterCouponRedemption and is used to iterate over the raw logs and unpacked data for CouponRedemption events raised by the Furnace contract.
type FurnaceCouponRedemptionIterator struct {
	Event *FurnaceCouponRedemption // Event containing the contract specifics and raw log

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
func (it *FurnaceCouponRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceCouponRedemption)
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
		it.Event = new(FurnaceCouponRedemption)
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
func (it *FurnaceCouponRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceCouponRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceCouponRedemption represents a CouponRedemption event raised by the Furnace contract.
type FurnaceCouponRedemption struct {
	Account      common.Address
	Epoch        *big.Int
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCouponRedemption is a free log retrieval operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) FilterCouponRedemption(opts *bind.FilterOpts, account []common.Address, epoch []*big.Int) (*FurnaceCouponRedemptionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "CouponRedemption", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceCouponRedemptionIterator{contract: _Furnace.contract, event: "CouponRedemption", logs: logs, sub: sub}, nil
}

// WatchCouponRedemption is a free log subscription operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) WatchCouponRedemption(opts *bind.WatchOpts, sink chan<- *FurnaceCouponRedemption, account []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "CouponRedemption", accountRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceCouponRedemption)
				if err := _Furnace.contract.UnpackLog(event, "CouponRedemption", log); err != nil {
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

// ParseCouponRedemption is a log parse operation binding the contract event 0x46e9903ae8ac9e9f0c9bc321b05965c1c036e7d4783758703f5cdfc4133c51b6.
//
// Solidity: event CouponRedemption(address indexed account, uint256 indexed epoch, uint256 couponAmount)
func (_Furnace *FurnaceFilterer) ParseCouponRedemption(log types.Log) (*FurnaceCouponRedemption, error) {
	event := new(FurnaceCouponRedemption)
	if err := _Furnace.contract.UnpackLog(event, "CouponRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceCouponTransferIterator is returned from FilterCouponTransfer and is used to iterate over the raw logs and unpacked data for CouponTransfer events raised by the Furnace contract.
type FurnaceCouponTransferIterator struct {
	Event *FurnaceCouponTransfer // Event containing the contract specifics and raw log

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
func (it *FurnaceCouponTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceCouponTransfer)
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
		it.Event = new(FurnaceCouponTransfer)
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
func (it *FurnaceCouponTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceCouponTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceCouponTransfer represents a CouponTransfer event raised by the Furnace contract.
type FurnaceCouponTransfer struct {
	From  common.Address
	To    common.Address
	Epoch *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCouponTransfer is a free log retrieval operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Furnace *FurnaceFilterer) FilterCouponTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, epoch []*big.Int) (*FurnaceCouponTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "CouponTransfer", fromRule, toRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceCouponTransferIterator{contract: _Furnace.contract, event: "CouponTransfer", logs: logs, sub: sub}, nil
}

// WatchCouponTransfer is a free log subscription operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Furnace *FurnaceFilterer) WatchCouponTransfer(opts *bind.WatchOpts, sink chan<- *FurnaceCouponTransfer, from []common.Address, to []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "CouponTransfer", fromRule, toRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceCouponTransfer)
				if err := _Furnace.contract.UnpackLog(event, "CouponTransfer", log); err != nil {
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

// ParseCouponTransfer is a log parse operation binding the contract event 0x0f1dbb1ccbe57a1590c7baad7b01d581b730c9ebc535dcde4345e6db424063d8.
//
// Solidity: event CouponTransfer(address indexed from, address indexed to, uint256 indexed epoch, uint256 value)
func (_Furnace *FurnaceFilterer) ParseCouponTransfer(log types.Log) (*FurnaceCouponTransfer, error) {
	event := new(FurnaceCouponTransfer)
	if err := _Furnace.contract.UnpackLog(event, "CouponTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Furnace contract.
type FurnaceDepositIterator struct {
	Event *FurnaceDeposit // Event containing the contract specifics and raw log

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
func (it *FurnaceDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceDeposit)
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
		it.Event = new(FurnaceDeposit)
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
func (it *FurnaceDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceDeposit represents a Deposit event raised by the Furnace contract.
type FurnaceDeposit struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*FurnaceDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceDepositIterator{contract: _Furnace.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FurnaceDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceDeposit)
				if err := _Furnace.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) ParseDeposit(log types.Log) (*FurnaceDeposit, error) {
	event := new(FurnaceDeposit)
	if err := _Furnace.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceIncentivizationIterator is returned from FilterIncentivization and is used to iterate over the raw logs and unpacked data for Incentivization events raised by the Furnace contract.
type FurnaceIncentivizationIterator struct {
	Event *FurnaceIncentivization // Event containing the contract specifics and raw log

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
func (it *FurnaceIncentivizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceIncentivization)
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
		it.Event = new(FurnaceIncentivization)
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
func (it *FurnaceIncentivizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceIncentivizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceIncentivization represents a Incentivization event raised by the Furnace contract.
type FurnaceIncentivization struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentivization is a free log retrieval operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Furnace *FurnaceFilterer) FilterIncentivization(opts *bind.FilterOpts, account []common.Address) (*FurnaceIncentivizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Incentivization", accountRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceIncentivizationIterator{contract: _Furnace.contract, event: "Incentivization", logs: logs, sub: sub}, nil
}

// WatchIncentivization is a free log subscription operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Furnace *FurnaceFilterer) WatchIncentivization(opts *bind.WatchOpts, sink chan<- *FurnaceIncentivization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Incentivization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceIncentivization)
				if err := _Furnace.contract.UnpackLog(event, "Incentivization", log); err != nil {
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

// ParseIncentivization is a log parse operation binding the contract event 0xbb4f656853bc420ad6e4321622c07eefb4ed40e3f91b35553ce14a6dff4c0981.
//
// Solidity: event Incentivization(address indexed account, uint256 amount)
func (_Furnace *FurnaceFilterer) ParseIncentivization(log types.Log) (*FurnaceIncentivization, error) {
	event := new(FurnaceIncentivization)
	if err := _Furnace.contract.UnpackLog(event, "Incentivization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Furnace contract.
type FurnaceOwnershipTransferredIterator struct {
	Event *FurnaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FurnaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceOwnershipTransferred)
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
		it.Event = new(FurnaceOwnershipTransferred)
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
func (it *FurnaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceOwnershipTransferred represents a OwnershipTransferred event raised by the Furnace contract.
type FurnaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Furnace *FurnaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FurnaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceOwnershipTransferredIterator{contract: _Furnace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Furnace *FurnaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FurnaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceOwnershipTransferred)
				if err := _Furnace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Furnace *FurnaceFilterer) ParseOwnershipTransferred(log types.Log) (*FurnaceOwnershipTransferred, error) {
	event := new(FurnaceOwnershipTransferred)
	if err := _Furnace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceSupplyDecreaseIterator is returned from FilterSupplyDecrease and is used to iterate over the raw logs and unpacked data for SupplyDecrease events raised by the Furnace contract.
type FurnaceSupplyDecreaseIterator struct {
	Event *FurnaceSupplyDecrease // Event containing the contract specifics and raw log

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
func (it *FurnaceSupplyDecreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceSupplyDecrease)
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
		it.Event = new(FurnaceSupplyDecrease)
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
func (it *FurnaceSupplyDecreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceSupplyDecreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceSupplyDecrease represents a SupplyDecrease event raised by the Furnace contract.
type FurnaceSupplyDecrease struct {
	Epoch   *big.Int
	Price   *big.Int
	NewDebt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupplyDecrease is a free log retrieval operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Furnace *FurnaceFilterer) FilterSupplyDecrease(opts *bind.FilterOpts, epoch []*big.Int) (*FurnaceSupplyDecreaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "SupplyDecrease", epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceSupplyDecreaseIterator{contract: _Furnace.contract, event: "SupplyDecrease", logs: logs, sub: sub}, nil
}

// WatchSupplyDecrease is a free log subscription operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Furnace *FurnaceFilterer) WatchSupplyDecrease(opts *bind.WatchOpts, sink chan<- *FurnaceSupplyDecrease, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "SupplyDecrease", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceSupplyDecrease)
				if err := _Furnace.contract.UnpackLog(event, "SupplyDecrease", log); err != nil {
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

// ParseSupplyDecrease is a log parse operation binding the contract event 0x5e139d4b8080a4a00dcc151e8217694aeebae893936326aa22096924a9906677.
//
// Solidity: event SupplyDecrease(uint256 indexed epoch, uint256 price, uint256 newDebt)
func (_Furnace *FurnaceFilterer) ParseSupplyDecrease(log types.Log) (*FurnaceSupplyDecrease, error) {
	event := new(FurnaceSupplyDecrease)
	if err := _Furnace.contract.UnpackLog(event, "SupplyDecrease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceSupplyIncreaseIterator is returned from FilterSupplyIncrease and is used to iterate over the raw logs and unpacked data for SupplyIncrease events raised by the Furnace contract.
type FurnaceSupplyIncreaseIterator struct {
	Event *FurnaceSupplyIncrease // Event containing the contract specifics and raw log

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
func (it *FurnaceSupplyIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceSupplyIncrease)
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
		it.Event = new(FurnaceSupplyIncrease)
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
func (it *FurnaceSupplyIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceSupplyIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceSupplyIncrease represents a SupplyIncrease event raised by the Furnace contract.
type FurnaceSupplyIncrease struct {
	Epoch         *big.Int
	Price         *big.Int
	NewRedeemable *big.Int
	LessDebt      *big.Int
	NewBonded     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSupplyIncrease is a free log retrieval operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) FilterSupplyIncrease(opts *bind.FilterOpts, epoch []*big.Int) (*FurnaceSupplyIncreaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "SupplyIncrease", epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceSupplyIncreaseIterator{contract: _Furnace.contract, event: "SupplyIncrease", logs: logs, sub: sub}, nil
}

// WatchSupplyIncrease is a free log subscription operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) WatchSupplyIncrease(opts *bind.WatchOpts, sink chan<- *FurnaceSupplyIncrease, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "SupplyIncrease", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceSupplyIncrease)
				if err := _Furnace.contract.UnpackLog(event, "SupplyIncrease", log); err != nil {
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

// ParseSupplyIncrease is a log parse operation binding the contract event 0x32fcaa1e76ed9517f4749d8ec9a77dd5e7329456d740b9bf9665d900eef5e283.
//
// Solidity: event SupplyIncrease(uint256 indexed epoch, uint256 price, uint256 newRedeemable, uint256 lessDebt, uint256 newBonded)
func (_Furnace *FurnaceFilterer) ParseSupplyIncrease(log types.Log) (*FurnaceSupplyIncrease, error) {
	event := new(FurnaceSupplyIncrease)
	if err := _Furnace.contract.UnpackLog(event, "SupplyIncrease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceSupplyNeutralIterator is returned from FilterSupplyNeutral and is used to iterate over the raw logs and unpacked data for SupplyNeutral events raised by the Furnace contract.
type FurnaceSupplyNeutralIterator struct {
	Event *FurnaceSupplyNeutral // Event containing the contract specifics and raw log

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
func (it *FurnaceSupplyNeutralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceSupplyNeutral)
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
		it.Event = new(FurnaceSupplyNeutral)
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
func (it *FurnaceSupplyNeutralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceSupplyNeutralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceSupplyNeutral represents a SupplyNeutral event raised by the Furnace contract.
type FurnaceSupplyNeutral struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSupplyNeutral is a free log retrieval operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Furnace *FurnaceFilterer) FilterSupplyNeutral(opts *bind.FilterOpts, epoch []*big.Int) (*FurnaceSupplyNeutralIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "SupplyNeutral", epochRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceSupplyNeutralIterator{contract: _Furnace.contract, event: "SupplyNeutral", logs: logs, sub: sub}, nil
}

// WatchSupplyNeutral is a free log subscription operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Furnace *FurnaceFilterer) WatchSupplyNeutral(opts *bind.WatchOpts, sink chan<- *FurnaceSupplyNeutral, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "SupplyNeutral", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceSupplyNeutral)
				if err := _Furnace.contract.UnpackLog(event, "SupplyNeutral", log); err != nil {
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

// ParseSupplyNeutral is a log parse operation binding the contract event 0xff7db5a0dc69b02c191ba632db46961b7d0daa1bd30709ddba9b80ad0a15d2c0.
//
// Solidity: event SupplyNeutral(uint256 indexed epoch)
func (_Furnace *FurnaceFilterer) ParseSupplyNeutral(log types.Log) (*FurnaceSupplyNeutral, error) {
	event := new(FurnaceSupplyNeutral)
	if err := _Furnace.contract.UnpackLog(event, "SupplyNeutral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Furnace contract.
type FurnaceTransferIterator struct {
	Event *FurnaceTransfer // Event containing the contract specifics and raw log

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
func (it *FurnaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceTransfer)
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
		it.Event = new(FurnaceTransfer)
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
func (it *FurnaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceTransfer represents a Transfer event raised by the Furnace contract.
type FurnaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Furnace *FurnaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FurnaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceTransferIterator{contract: _Furnace.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Furnace *FurnaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FurnaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceTransfer)
				if err := _Furnace.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Furnace *FurnaceFilterer) ParseTransfer(log types.Log) (*FurnaceTransfer, error) {
	event := new(FurnaceTransfer)
	if err := _Furnace.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Furnace contract.
type FurnaceUnbondIterator struct {
	Event *FurnaceUnbond // Event containing the contract specifics and raw log

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
func (it *FurnaceUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceUnbond)
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
		it.Event = new(FurnaceUnbond)
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
func (it *FurnaceUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceUnbond represents a Unbond event raised by the Furnace contract.
type FurnaceUnbond struct {
	Account         common.Address
	Start           *big.Int
	Value           *big.Int
	ValueUnderlying *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) FilterUnbond(opts *bind.FilterOpts, account []common.Address) (*FurnaceUnbondIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Unbond", accountRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceUnbondIterator{contract: _Furnace.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *FurnaceUnbond, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Unbond", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceUnbond)
				if err := _Furnace.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0x93530ac0ee8c50e696e13c5ac62355d0c0ba4bd943620d5bda1eb08b64ae7512.
//
// Solidity: event Unbond(address indexed account, uint256 start, uint256 value, uint256 valueUnderlying)
func (_Furnace *FurnaceFilterer) ParseUnbond(log types.Log) (*FurnaceUnbond, error) {
	event := new(FurnaceUnbond)
	if err := _Furnace.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FurnaceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Furnace contract.
type FurnaceWithdrawIterator struct {
	Event *FurnaceWithdraw // Event containing the contract specifics and raw log

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
func (it *FurnaceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FurnaceWithdraw)
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
		it.Event = new(FurnaceWithdraw)
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
func (it *FurnaceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FurnaceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FurnaceWithdraw represents a Withdraw event raised by the Furnace contract.
type FurnaceWithdraw struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*FurnaceWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &FurnaceWithdrawIterator{contract: _Furnace.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FurnaceWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Furnace.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FurnaceWithdraw)
				if err := _Furnace.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 value)
func (_Furnace *FurnaceFilterer) ParseWithdraw(log types.Log) (*FurnaceWithdraw, error) {
	event := new(FurnaceWithdraw)
	if err := _Furnace.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
