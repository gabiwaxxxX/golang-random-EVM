package uniswapV3Decode

import (
	"fmt"
)

func (exactInputSingle ExactInputSingleParams) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  TokenIn: %v\n  TokenOut: %v\n  Fee: %v\n  Recipient: %v\n "+
		"  AmountIn: %v\n  AmountOutMinimum: %v\n",
		exactInputSingle.FuncName, exactInputSingle.TokenIn, exactInputSingle.TokenOut, exactInputSingle.Fee,
		exactInputSingle.Recipient, exactInputSingle.AmountIn,
		exactInputSingle.AmountOutMinimum)
}

func (exactOutputSingle exactOutpuSingleParams) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  TokenIn: %v\n  TokenOut: %v\n  Fee: %v\n  Recipient: %v\n "+
		"  AmountOut: %v\n  AmountInMaximum: %v\n",
		exactOutputSingle.FuncName, exactOutputSingle.TokenIn, exactOutputSingle.TokenOut, exactOutputSingle.Fee,
		exactOutputSingle.Recipient, exactOutputSingle.AmountOut,
		exactOutputSingle.AmountInMaximum)
}

func (exactInput ExactInputStringPath) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  Path: %v\n  Recipient: %v\n AmountIn: %v\n"+
		"  AmountOutMinimum: %v\n", exactInput.FuncName, exactInput.Path, exactInput.Recipient,
		exactInput.AmountIn, exactInput.AmountOutMinimum)
}

func (exactOutput ExactOutPutStringPath) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  Path: %v\n  Recipient: %v\n "+
		"  AmountOut: %v\n  AmountInMaximum: %v\n", exactOutput.FuncName,
		exactOutput.Path, exactOutput.Recipient,
		exactOutput.AmountOut, exactOutput.AmountInMaximum)
}

func (selfPermit SelfPermit) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  Token: %v\n  Value: %v\n  Deadline: %v\n  V: %v\n  R: %v\n  S: %v\n",
		selfPermit.FuncName, selfPermit.Token, selfPermit.Value, selfPermit.Deadline,
		selfPermit.V, selfPermit.R, selfPermit.S)
}

func (swapExactTokensForTokens swapExactTokensForTokensParams) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  AmountIn: %v\n  AmountOutMin: %v\n  Path: %v\n  To: %v\n",
		swapExactTokensForTokens.FuncName, swapExactTokensForTokens.AmountIn, swapExactTokensForTokens.AmountOutMin,
		swapExactTokensForTokens.Path, swapExactTokensForTokens.To)
}

func (swapTokensForExactTokens swapTokensForExactTokensParams) String() string {
	return fmt.Sprintf("Function: %v\n Parameters: \n  AmountOut: %v\n  AmountInMax: %v\n  Path: %v\n  To: %v\n",
		swapTokensForExactTokens.FuncName, swapTokensForExactTokens.AmountOut, swapTokensForExactTokens.AmountInMax,
		swapTokensForExactTokens.Path, swapTokensForExactTokens.To)
}
