// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Router01

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

// Router01Route is an auto generated low-level Go binding around an user-defined struct.
type Router01Route struct {
	From   common.Address
	To     common.Address
	Stable bool
}

// Router01MetaData contains all meta data concerning the Router01 contract.
var Router01MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"UNSAFE_swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATICMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityMATIC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wmatic\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quoteLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATICMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityMATIC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATIC\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountFTMMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityMATICSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountFTM\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATICMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityMATICWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMATIC\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountFTMMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityMATICWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountFTM\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactMATICForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactMATICForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForMATIC\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForMATICSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSimple\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter01.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wmatic\",\"outputs\":[{\"internalType\":\"contractIWMATIC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Router01ABI is the input ABI used to generate the binding from.
// Deprecated: Use Router01MetaData.ABI instead.
var Router01ABI = Router01MetaData.ABI

// Router01 is an auto generated Go binding around an Ethereum contract.
type Router01 struct {
	Router01Caller     // Read-only binding to the contract
	Router01Transactor // Write-only binding to the contract
	Router01Filterer   // Log filterer for contract events
}

// Router01Caller is an auto generated read-only Go binding around an Ethereum contract.
type Router01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Router01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Router01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Router01Session struct {
	Contract     *Router01         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Router01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Router01CallerSession struct {
	Contract *Router01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Router01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Router01TransactorSession struct {
	Contract     *Router01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Router01Raw is an auto generated low-level Go binding around an Ethereum contract.
type Router01Raw struct {
	Contract *Router01 // Generic contract binding to access the raw methods on
}

// Router01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Router01CallerRaw struct {
	Contract *Router01Caller // Generic read-only contract binding to access the raw methods on
}

// Router01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Router01TransactorRaw struct {
	Contract *Router01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter01 creates a new instance of Router01, bound to a specific deployed contract.
func NewRouter01(address common.Address, backend bind.ContractBackend) (*Router01, error) {
	contract, err := bindRouter01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router01{Router01Caller: Router01Caller{contract: contract}, Router01Transactor: Router01Transactor{contract: contract}, Router01Filterer: Router01Filterer{contract: contract}}, nil
}

// NewRouter01Caller creates a new read-only instance of Router01, bound to a specific deployed contract.
func NewRouter01Caller(address common.Address, caller bind.ContractCaller) (*Router01Caller, error) {
	contract, err := bindRouter01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Router01Caller{contract: contract}, nil
}

// NewRouter01Transactor creates a new write-only instance of Router01, bound to a specific deployed contract.
func NewRouter01Transactor(address common.Address, transactor bind.ContractTransactor) (*Router01Transactor, error) {
	contract, err := bindRouter01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Router01Transactor{contract: contract}, nil
}

// NewRouter01Filterer creates a new log filterer instance of Router01, bound to a specific deployed contract.
func NewRouter01Filterer(address common.Address, filterer bind.ContractFilterer) (*Router01Filterer, error) {
	contract, err := bindRouter01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Router01Filterer{contract: contract}, nil
}

// bindRouter01 binds a generic wrapper to an already deployed contract.
func bindRouter01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Router01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router01 *Router01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router01.Contract.Router01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router01 *Router01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router01.Contract.Router01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router01 *Router01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router01.Contract.Router01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router01 *Router01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router01 *Router01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router01 *Router01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router01.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router01 *Router01Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router01 *Router01Session) Factory() (common.Address, error) {
	return _Router01.Contract.Factory(&_Router01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router01 *Router01CallerSession) Factory() (common.Address, error) {
	return _Router01.Contract.Factory(&_Router01.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_Router01 *Router01Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

	outstruct := new(struct {
		Amount *big.Int
		Stable bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_Router01 *Router01Session) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _Router01.Contract.GetAmountOut(&_Router01.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_Router01 *Router01CallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _Router01.Contract.GetAmountOut(&_Router01.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router01 *Router01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, routes []Router01Route) ([]*big.Int, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "getAmountsOut", amountIn, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router01 *Router01Session) GetAmountsOut(amountIn *big.Int, routes []Router01Route) ([]*big.Int, error) {
	return _Router01.Contract.GetAmountsOut(&_Router01.CallOpts, amountIn, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router01 *Router01CallerSession) GetAmountsOut(amountIn *big.Int, routes []Router01Route) ([]*big.Int, error) {
	return _Router01.Contract.GetAmountsOut(&_Router01.CallOpts, amountIn, routes)
}

// GetExactAmountOut is a free data retrieval call binding the contract method 0xb6710cb9.
//
// Solidity: function getExactAmountOut(uint256 amountIn, address tokenIn, address tokenOut, bool stable) view returns(uint256)
func (_Router01 *Router01Caller) GetExactAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address, stable bool) (*big.Int, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "getExactAmountOut", amountIn, tokenIn, tokenOut, stable)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExactAmountOut is a free data retrieval call binding the contract method 0xb6710cb9.
//
// Solidity: function getExactAmountOut(uint256 amountIn, address tokenIn, address tokenOut, bool stable) view returns(uint256)
func (_Router01 *Router01Session) GetExactAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address, stable bool) (*big.Int, error) {
	return _Router01.Contract.GetExactAmountOut(&_Router01.CallOpts, amountIn, tokenIn, tokenOut, stable)
}

// GetExactAmountOut is a free data retrieval call binding the contract method 0xb6710cb9.
//
// Solidity: function getExactAmountOut(uint256 amountIn, address tokenIn, address tokenOut, bool stable) view returns(uint256)
func (_Router01 *Router01CallerSession) GetExactAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address, stable bool) (*big.Int, error) {
	return _Router01.Contract.GetExactAmountOut(&_Router01.CallOpts, amountIn, tokenIn, tokenOut, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Router01 *Router01Caller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Router01 *Router01Session) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Router01.Contract.GetReserves(&_Router01.CallOpts, tokenA, tokenB, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Router01 *Router01CallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Router01.Contract.GetReserves(&_Router01.CallOpts, tokenA, tokenB, stable)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router01 *Router01Caller) IsPair(opts *bind.CallOpts, pair common.Address) (bool, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "isPair", pair)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router01 *Router01Session) IsPair(pair common.Address) (bool, error) {
	return _Router01.Contract.IsPair(&_Router01.CallOpts, pair)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router01 *Router01CallerSession) IsPair(pair common.Address) (bool, error) {
	return _Router01.Contract.IsPair(&_Router01.CallOpts, pair)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router01 *Router01Caller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router01 *Router01Session) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Router01.Contract.PairFor(&_Router01.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router01 *Router01CallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Router01.Contract.PairFor(&_Router01.CallOpts, tokenA, tokenB, stable)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01Caller) QuoteAddLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "quoteAddLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired)

	outstruct := new(struct {
		AmountA   *big.Int
		AmountB   *big.Int
		Liquidity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01Session) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _Router01.Contract.QuoteAddLiquidity(&_Router01.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01CallerSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _Router01.Contract.QuoteAddLiquidity(&_Router01.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteLiquidity is a free data retrieval call binding the contract method 0xae568868.
//
// Solidity: function quoteLiquidity(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Router01 *Router01Caller) QuoteLiquidity(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "quoteLiquidity", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteLiquidity is a free data retrieval call binding the contract method 0xae568868.
//
// Solidity: function quoteLiquidity(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Router01 *Router01Session) QuoteLiquidity(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Router01.Contract.QuoteLiquidity(&_Router01.CallOpts, amountA, reserveA, reserveB)
}

// QuoteLiquidity is a free data retrieval call binding the contract method 0xae568868.
//
// Solidity: function quoteLiquidity(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Router01 *Router01CallerSession) QuoteLiquidity(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Router01.Contract.QuoteLiquidity(&_Router01.CallOpts, amountA, reserveA, reserveB)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Caller) QuoteRemoveLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "quoteRemoveLiquidity", tokenA, tokenB, stable, liquidity)

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Session) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Router01.Contract.QuoteRemoveLiquidity(&_Router01.CallOpts, tokenA, tokenB, stable, liquidity)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01CallerSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Router01.Contract.QuoteRemoveLiquidity(&_Router01.CallOpts, tokenA, tokenB, stable, liquidity)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Router01 *Router01Caller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Router01 *Router01Session) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Router01.Contract.SortTokens(&_Router01.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Router01 *Router01CallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Router01.Contract.SortTokens(&_Router01.CallOpts, tokenA, tokenB)
}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Router01 *Router01Caller) Wmatic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router01.contract.Call(opts, &out, "wmatic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Router01 *Router01Session) Wmatic() (common.Address, error) {
	return _Router01.Contract.Wmatic(&_Router01.CallOpts)
}

// Wmatic is a free data retrieval call binding the contract method 0xfb41be16.
//
// Solidity: function wmatic() view returns(address)
func (_Router01 *Router01CallerSession) Wmatic() (common.Address, error) {
	return _Router01.Contract.Wmatic(&_Router01.CallOpts)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router01 *Router01Transactor) UNSAFESwapExactTokensForTokens(opts *bind.TransactOpts, amounts []*big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router01 *Router01Session) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.UNSAFESwapExactTokensForTokens(&_Router01.TransactOpts, amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router01 *Router01TransactorSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.UNSAFESwapExactTokensForTokens(&_Router01.TransactOpts, amounts, routes, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01Session) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.AddLiquidity(&_Router01.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router01 *Router01TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.AddLiquidity(&_Router01.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityMATIC is a paid mutator transaction binding the contract method 0xe2dbe241.
//
// Solidity: function addLiquidityMATIC(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountMATIC, uint256 liquidity)
func (_Router01 *Router01Transactor) AddLiquidityMATIC(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "addLiquidityMATIC", token, stable, amountTokenDesired, amountTokenMin, amountMATICMin, to, deadline)
}

// AddLiquidityMATIC is a paid mutator transaction binding the contract method 0xe2dbe241.
//
// Solidity: function addLiquidityMATIC(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountMATIC, uint256 liquidity)
func (_Router01 *Router01Session) AddLiquidityMATIC(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.AddLiquidityMATIC(&_Router01.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountMATICMin, to, deadline)
}

// AddLiquidityMATIC is a paid mutator transaction binding the contract method 0xe2dbe241.
//
// Solidity: function addLiquidityMATIC(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountMATIC, uint256 liquidity)
func (_Router01 *Router01TransactorSession) AddLiquidityMATIC(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.AddLiquidityMATIC(&_Router01.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountMATICMin, to, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _wmatic) returns()
func (_Router01 *Router01Transactor) Initialize(opts *bind.TransactOpts, _factory common.Address, _wmatic common.Address) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "initialize", _factory, _wmatic)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _wmatic) returns()
func (_Router01 *Router01Session) Initialize(_factory common.Address, _wmatic common.Address) (*types.Transaction, error) {
	return _Router01.Contract.Initialize(&_Router01.TransactOpts, _factory, _wmatic)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _wmatic) returns()
func (_Router01 *Router01TransactorSession) Initialize(_factory common.Address, _wmatic common.Address) (*types.Transaction, error) {
	return _Router01.Contract.Initialize(&_Router01.TransactOpts, _factory, _wmatic)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidity(&_Router01.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidity(&_Router01.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityMATIC is a paid mutator transaction binding the contract method 0xcf14e1b1.
//
// Solidity: function removeLiquidityMATIC(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01Transactor) RemoveLiquidityMATIC(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidityMATIC", token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline)
}

// RemoveLiquidityMATIC is a paid mutator transaction binding the contract method 0xcf14e1b1.
//
// Solidity: function removeLiquidityMATIC(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01Session) RemoveLiquidityMATIC(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATIC(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline)
}

// RemoveLiquidityMATIC is a paid mutator transaction binding the contract method 0xcf14e1b1.
//
// Solidity: function removeLiquidityMATIC(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01TransactorSession) RemoveLiquidityMATIC(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATIC(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline)
}

// RemoveLiquidityMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x765f62a9.
//
// Solidity: function removeLiquidityMATICSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01Transactor) RemoveLiquidityMATICSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidityMATICSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline)
}

// RemoveLiquidityMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x765f62a9.
//
// Solidity: function removeLiquidityMATICSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01Session) RemoveLiquidityMATICSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICSupportingFeeOnTransferTokens(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline)
}

// RemoveLiquidityMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x765f62a9.
//
// Solidity: function removeLiquidityMATICSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01TransactorSession) RemoveLiquidityMATICSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICSupportingFeeOnTransferTokens(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline)
}

// RemoveLiquidityMATICWithPermit is a paid mutator transaction binding the contract method 0xc11efebc.
//
// Solidity: function removeLiquidityMATICWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01Transactor) RemoveLiquidityMATICWithPermit(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidityMATICWithPermit", token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityMATICWithPermit is a paid mutator transaction binding the contract method 0xc11efebc.
//
// Solidity: function removeLiquidityMATICWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01Session) RemoveLiquidityMATICWithPermit(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICWithPermit(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityMATICWithPermit is a paid mutator transaction binding the contract method 0xc11efebc.
//
// Solidity: function removeLiquidityMATICWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountMATICMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountMATIC)
func (_Router01 *Router01TransactorSession) RemoveLiquidityMATICWithPermit(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountMATICMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICWithPermit(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountMATICMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd251deed.
//
// Solidity: function removeLiquidityMATICWithPermitSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01Transactor) RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidityMATICWithPermitSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd251deed.
//
// Solidity: function removeLiquidityMATICWithPermitSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01Session) RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd251deed.
//
// Solidity: function removeLiquidityMATICWithPermitSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountFTMMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountFTM)
func (_Router01 *Router01TransactorSession) RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountFTMMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityMATICWithPermitSupportingFeeOnTransferTokens(&_Router01.TransactOpts, token, stable, liquidity, amountTokenMin, amountFTMMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0xa32b1fcd.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0xa32b1fcd.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityWithPermit(&_Router01.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0xa32b1fcd.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Router01 *Router01TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router01.Contract.RemoveLiquidityWithPermit(&_Router01.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapExactMATICForTokens is a paid mutator transaction binding the contract method 0x92e36441.
//
// Solidity: function swapExactMATICForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router01 *Router01Transactor) SwapExactMATICForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactMATICForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactMATICForTokens is a paid mutator transaction binding the contract method 0x92e36441.
//
// Solidity: function swapExactMATICForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router01 *Router01Session) SwapExactMATICForTokens(amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactMATICForTokens(&_Router01.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactMATICForTokens is a paid mutator transaction binding the contract method 0x92e36441.
//
// Solidity: function swapExactMATICForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router01 *Router01TransactorSession) SwapExactMATICForTokens(amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactMATICForTokens(&_Router01.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactMATICForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x25200f17.
//
// Solidity: function swapExactMATICForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_Router01 *Router01Transactor) SwapExactMATICForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactMATICForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
}

// SwapExactMATICForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x25200f17.
//
// Solidity: function swapExactMATICForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_Router01 *Router01Session) SwapExactMATICForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactMATICForTokensSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactMATICForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x25200f17.
//
// Solidity: function swapExactMATICForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_Router01 *Router01TransactorSession) SwapExactMATICForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactMATICForTokensSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATIC is a paid mutator transaction binding the contract method 0x340dbb0b.
//
// Solidity: function swapExactTokensForMATIC(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Transactor) SwapExactTokensForMATIC(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactTokensForMATIC", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATIC is a paid mutator transaction binding the contract method 0x340dbb0b.
//
// Solidity: function swapExactTokensForMATIC(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Session) SwapExactTokensForMATIC(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForMATIC(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATIC is a paid mutator transaction binding the contract method 0x340dbb0b.
//
// Solidity: function swapExactTokensForMATIC(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01TransactorSession) SwapExactTokensForMATIC(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForMATIC(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfbace415.
//
// Solidity: function swapExactTokensForMATICSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01Transactor) SwapExactTokensForMATICSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactTokensForMATICSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfbace415.
//
// Solidity: function swapExactTokensForMATICSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01Session) SwapExactTokensForMATICSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForMATICSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForMATICSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfbace415.
//
// Solidity: function swapExactTokensForMATICSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01TransactorSession) SwapExactTokensForMATICSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForMATICSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Transactor) SwapExactTokensForTokensSimple(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactTokensForTokensSimple", amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01Session) SwapExactTokensForTokensSimple(amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokensSimple(&_Router01.TransactOpts, amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router01 *Router01TransactorSession) SwapExactTokensForTokensSimple(amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokensSimple(&_Router01.TransactOpts, amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_Router01 *Router01TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Router01Route, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router01.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Router01.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router01 *Router01Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router01.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router01 *Router01Session) Receive() (*types.Transaction, error) {
	return _Router01.Contract.Receive(&_Router01.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router01 *Router01TransactorSession) Receive() (*types.Transaction, error) {
	return _Router01.Contract.Receive(&_Router01.TransactOpts)
}

// Router01InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Router01 contract.
type Router01InitializedIterator struct {
	Event *Router01Initialized // Event containing the contract specifics and raw log

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
func (it *Router01InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router01Initialized)
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
		it.Event = new(Router01Initialized)
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
func (it *Router01InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router01InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router01Initialized represents a Initialized event raised by the Router01 contract.
type Router01Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router01 *Router01Filterer) FilterInitialized(opts *bind.FilterOpts) (*Router01InitializedIterator, error) {

	logs, sub, err := _Router01.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Router01InitializedIterator{contract: _Router01.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router01 *Router01Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Router01Initialized) (event.Subscription, error) {

	logs, sub, err := _Router01.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router01Initialized)
				if err := _Router01.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Router01 *Router01Filterer) ParseInitialized(log types.Log) (*Router01Initialized, error) {
	event := new(Router01Initialized)
	if err := _Router01.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
