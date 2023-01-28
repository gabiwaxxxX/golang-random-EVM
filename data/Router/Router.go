// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Router

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

// Routerroute is an auto generated low-level Go binding around an user-defined struct.
type Routerroute struct {
	From   common.Address
	To     common.Address
	Stable bool
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"UNSAFE_swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"isPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSimple\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_Router *RouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

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
func (_Router *RouterSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _Router.Contract.GetAmountOut(&_Router.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_Router *RouterCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _Router.Contract.GetAmountOut(&_Router.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router *RouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, routes []Routerroute) ([]*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getAmountsOut", amountIn, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router *RouterSession) GetAmountsOut(amountIn *big.Int, routes []Routerroute) ([]*big.Int, error) {
	return _Router.Contract.GetAmountsOut(&_Router.CallOpts, amountIn, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_Router *RouterCallerSession) GetAmountsOut(amountIn *big.Int, routes []Routerroute) ([]*big.Int, error) {
	return _Router.Contract.GetAmountsOut(&_Router.CallOpts, amountIn, routes)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Router *RouterCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable)

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
func (_Router *RouterSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Router.Contract.GetReserves(&_Router.CallOpts, tokenA, tokenB, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Router *RouterCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Router.Contract.GetReserves(&_Router.CallOpts, tokenA, tokenB, stable)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router *RouterCaller) IsPair(opts *bind.CallOpts, pair common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isPair", pair)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router *RouterSession) IsPair(pair common.Address) (bool, error) {
	return _Router.Contract.IsPair(&_Router.CallOpts, pair)
}

// IsPair is a free data retrieval call binding the contract method 0xe5e31b13.
//
// Solidity: function isPair(address pair) view returns(bool)
func (_Router *RouterCallerSession) IsPair(pair common.Address) (bool, error) {
	return _Router.Contract.IsPair(&_Router.CallOpts, pair)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router *RouterCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router *RouterSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Router.Contract.PairFor(&_Router.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Router *RouterCallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Router.Contract.PairFor(&_Router.CallOpts, tokenA, tokenB, stable)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router *RouterCaller) QuoteAddLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "quoteAddLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired)

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
func (_Router *RouterSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _Router.Contract.QuoteAddLiquidity(&_Router.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router *RouterCallerSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _Router.Contract.QuoteAddLiquidity(&_Router.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_Router *RouterCaller) QuoteRemoveLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "quoteRemoveLiquidity", tokenA, tokenB, stable, liquidity)

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
func (_Router *RouterSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Router.Contract.QuoteRemoveLiquidity(&_Router.CallOpts, tokenA, tokenB, stable, liquidity)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_Router *RouterCallerSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _Router.Contract.QuoteRemoveLiquidity(&_Router.CallOpts, tokenA, tokenB, stable, liquidity)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Router *RouterCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

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
func (_Router *RouterSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Router.Contract.SortTokens(&_Router.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Router *RouterCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Router.Contract.SortTokens(&_Router.CallOpts, tokenA, tokenB)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterSession) Weth() (common.Address, error) {
	return _Router.Contract.Weth(&_Router.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterCallerSession) Weth() (common.Address, error) {
	return _Router.Contract.Weth(&_Router.CallOpts)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router *RouterTransactor) UNSAFESwapExactTokensForTokens(opts *bind.TransactOpts, amounts []*big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router *RouterSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UNSAFESwapExactTokensForTokens(&_Router.TransactOpts, amounts, routes, to, deadline)
}

// UNSAFESwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x7301e3c8.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[])
func (_Router *RouterTransactorSession) UNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.UNSAFESwapExactTokensForTokens(&_Router.TransactOpts, amounts, routes, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router *RouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Router *RouterSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router *RouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router *RouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Router *RouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETH(&_Router.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterTransactorSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETH(&_Router.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0x448725b4.
//
// Solidity: function removeLiquidityETHWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterTransactor) RemoveLiquidityETHWithPermit(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidityETHWithPermit", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0x448725b4.
//
// Solidity: function removeLiquidityETHWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterSession) RemoveLiquidityETHWithPermit(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETHWithPermit(&_Router.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityETHWithPermit is a paid mutator transaction binding the contract method 0x448725b4.
//
// Solidity: function removeLiquidityETHWithPermit(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (_Router *RouterTransactorSession) RemoveLiquidityETHWithPermit(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidityETHWithPermit(&_Router.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactETHForTokens(&_Router.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactETHForTokens(&_Router.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForETH(&_Router.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForETH(&_Router.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokens(&_Router.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []Routerroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokens(&_Router.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactor) SwapExactTokensForTokensSimple(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapExactTokensForTokensSimple", amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterSession) SwapExactTokensForTokensSimple(amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokensSimple(&_Router.TransactOpts, amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// SwapExactTokensForTokensSimple is a paid mutator transaction binding the contract method 0x13dcfc59.
//
// Solidity: function swapExactTokensForTokensSimple(uint256 amountIn, uint256 amountOutMin, address tokenFrom, address tokenTo, bool stable, address to, uint256 deadline) returns(uint256[] amounts)
func (_Router *RouterTransactorSession) SwapExactTokensForTokensSimple(amountIn *big.Int, amountOutMin *big.Int, tokenFrom common.Address, tokenTo common.Address, stable bool, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapExactTokensForTokensSimple(&_Router.TransactOpts, amountIn, amountOutMin, tokenFrom, tokenTo, stable, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}
