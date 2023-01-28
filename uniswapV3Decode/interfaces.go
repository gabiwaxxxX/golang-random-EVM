package uniswapV3Decode

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ExactInputSingleParams struct {
	FuncName          string
	TokenIn           common.Address "json:\"tokenIn\""
	TokenOut          common.Address "json:\"tokenOut\""
	Fee               *big.Int       "json:\"fee\""
	Recipient         common.Address "json:\"recipient\""
	AmountIn          *big.Int       "json:\"amountIn\""
	AmountOutMinimum  *big.Int       "json:\"amountOutMinimum\""
	SqrtPriceLimitX96 *big.Int       "json:\"sqrtPriceLimitX96\""
}

type exactOutpuSingleParams struct {
	FuncName          string
	TokenIn           common.Address "json:\"tokenIn\""
	TokenOut          common.Address "json:\"tokenOut\""
	Fee               *big.Int       "json:\"fee\""
	Recipient         common.Address "json:\"recipient\""
	AmountOut         *big.Int       "json:\"amountOut\""
	AmountInMaximum   *big.Int       "json:\"amountInMaximum\""
	SqrtPriceLimitX96 *big.Int       "json:\"sqrtPriceLimitX96\""
}

type ExactInputBytesPath struct {
	Path             []uint8        "json:\"path\""
	Recipient        common.Address "json:\"recipient\""
	AmountIn         *big.Int       "json:\"amountIn\""
	AmountOutMinimum *big.Int       "json:\"amountOutMinimum\""
}

type ExactInputStringPath struct {
	FuncName         string
	Path             []string       "json:\"path\""
	Recipient        common.Address "json:\"recipient\""
	AmountIn         *big.Int       "json:\"amountIn\""
	AmountOutMinimum *big.Int       "json:\"amountOutMinimum\""
}

type ExactOutPutBytesPath struct {
	Path            []uint8        "json:\"path\""
	Recipient       common.Address "json:\"recipient\""
	AmountOut       *big.Int       "json:\"amountOut\""
	AmountInMaximum *big.Int       "json:\"amountInMaximum\""
}

type ExactOutPutStringPath struct {
	FuncName        string
	Path            []string       "json:\"path\""
	Recipient       common.Address "json:\"recipient\""
	AmountOut       *big.Int       "json:\"amountOut\""
	AmountInMaximum *big.Int       "json:\"amountInMaximum\""
}

type SelfPermit struct {
	FuncName string
	Token    common.Address
	Value    *big.Int
	Deadline *big.Int
	V        *big.Int
	R        string
	S        string
}

type swapExactTokensForTokensParams struct {
	FuncName     string
	AmountIn     *big.Int         "json:\"amountIn\""
	AmountOutMin *big.Int         "json:\"amountOutMin\""
	Path         []common.Address "json:\"path\""
	To           common.Address   "json:\"to\""
}

type swapTokensForExactTokensParams struct {
	FuncName    string
	AmountOut   *big.Int         "json:\"amountOut\""
	AmountInMax *big.Int         "json:\"amountInMax\""
	Path        []common.Address "json:\"path\""
	To          common.Address   "json:\"to\""
}
