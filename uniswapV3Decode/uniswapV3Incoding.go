package uniswapV3Decode

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/contract/SwapRouter"
	"github.com/gabrielfournier1/bot/contract/UniswapV3Factory"
	"github.com/gabrielfournier1/bot/eth_const"
	"github.com/gabrielfournier1/bot/utils"
)

const (
	SwapRouter00Address     = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	UniSwapV3FactoryAddress = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	SwapRouter02Address     = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"
)

var (
	poolFeeTier001 = big.NewInt(1000)
	poolFeeTier003 = big.NewInt(3000)
	poolFeeTier01  = big.NewInt(10000)
)

func FindPool(tokenA common.Address, tokenB common.Address) {
	factory, _ := UniswapV3Factory.NewUniswapV3Factory(common.HexToAddress(UniSwapV3FactoryAddress), utils.GetClient())
	pool, _ := factory.GetPool(nil, tokenA, tokenB, poolFeeTier001)
	if pool != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		println("pool found")
		println(pool.Hex())
		SwapExactInputSingle(tokenA, tokenB, poolFeeTier001, big.NewInt(1000000000000000), big.NewInt(0))
		return
	}
	pool, _ = factory.GetPool(nil, tokenA, tokenB, poolFeeTier003)
	if pool != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		println("pool found")
		println(pool.Hex())
		SwapExactInputSingle(tokenA, tokenB, poolFeeTier003, big.NewInt(1000000000000000), big.NewInt(0))

		return

	}
	pool, _ = factory.GetPool(nil, tokenA, tokenB, poolFeeTier01)
	if pool != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		println("pool found")
		println(pool.Hex())
		SwapExactInputSingle(tokenA, tokenB, poolFeeTier01, big.NewInt(1000000000000000), big.NewInt(0))

		return

	}
}

// swap a fixed amount of one token for a maximum possible amount of another token.
func SwapExactInputSingle(tokenIn, tokenOut common.Address, fee *big.Int, amountIn, amountOutMinimum *big.Int) {
	router, _ := SwapRouter.NewSwapRouter(common.HexToAddress(SwapRouter00Address), utils.GetClient())

	var input SwapRouter.ISwapRouterExactInputSingleParams = SwapRouter.ISwapRouterExactInputSingleParams{
		AmountIn:          amountIn,
		AmountOutMinimum:  amountOutMinimum,
		Recipient:         utils.GetUserInfo().FromAddress,
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		Fee:               fee,
		SqrtPriceLimitX96: new(big.Int).SetInt64(0),
		Deadline:          big.NewInt(time.Now().Add(5 * time.Minute).Unix()),
	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 300000
	auth.GasPrice, err = utils.GetClient().SuggestGasPrice(context.Background())

	if condition := tokenIn.Hex() == eth_const.FindAddressFromTokenSymbol("WETH"); condition {
		auth.Value = amountIn
	}

	tx, err := router.ExactInputSingle(auth, input)
	fmt.Println(common.Bytes2Hex(tx.Data()))
	if err != nil {
		println("error")
		println(err.Error())
	}
	println("tx sent")
	println(tx.Hash().Hex())
	println("waiting for tx to be mined")
	utils.AwaitToConfirmTx(tx)
}

// which swap a minimum possible amount of one token for a fixed amount of another token.
func SwapExactOutputSingle(tokenIn, tokenOut common.Address, fee *big.Int, amountOut, amountInMaximum *big.Int) {
	router, _ := SwapRouter.NewSwapRouter(common.HexToAddress(SwapRouter00Address), utils.GetClient())

	var input SwapRouter.ISwapRouterExactOutputSingleParams = SwapRouter.ISwapRouterExactOutputSingleParams{
		AmountOut:         amountOut,
		AmountInMaximum:   amountInMaximum,
		Recipient:         utils.GetUserInfo().FromAddress,
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		Fee:               fee,
		SqrtPriceLimitX96: new(big.Int).SetInt64(0),
		Deadline:          big.NewInt(time.Now().Add(5 * time.Minute).Unix()),
	}

	tx, err := router.ExactOutputSingle(nil, input)
	fmt.Println(tx.Data())
	if err != nil {
		println("error")
		println(err.Error())
	}
	println("tx sent")
	println(tx.Hash().Hex())
	println("waiting for tx to be mined")
	utils.AwaitToConfirmTx(tx)
}

// unwrapWETH9
// 0x49404b7c
// 00000000000000000000000000000000000000000000000001d577805823ed14
// 0000000000000000000000001d1661cb61bf5e3066f17f82099786d0fcc49d46

// eth to USDC

// 0x53b4df
// 000000000000000000000000caaa39b223fe8d0a0e5c4f27ead9083c756cc2
// 000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
// 00000000000000000000000000000000000000000000000000000000000001f4
// 0000000000000000000000001d1661cb61bf5e3066f17f82099786d0fcc49d46
// 000000000000000000000000000000000000000000000000000000000dbb03c0
// 00000000000000000000000000000000000000000000000001d6b5911b98f0
// 0000000000000000000000000000000000000000000000000000000000000000

// Refund ETH
// 0x12210e8a
