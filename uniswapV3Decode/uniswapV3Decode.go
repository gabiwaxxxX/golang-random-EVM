package uniswapV3Decode

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"sync"	

	"github.com/gabrielfournier1/bot/utils"
    "github.com/gabrielfournier1/bot/contract/UniswapV2Factory"
    "github.com/gabrielfournier1/bot/contract/UniswapV2Pair"
    "github.com/gabrielfournier1/bot/eth_const"


	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)


type BinarySearchResult struct {
	MaxETHICanBuy          *big.Int
	AmountTknIWillBuy      *big.Int
	AmountTknVictimWillBuy *big.Int
	Rtkn1                  *big.Int
	Reth1                  *big.Int
	ExpectedProfits        *big.Int
}

var BinaryResult *BinarySearchResult


func DecodeInput(value *big.Int,G_uniV3abi abi.ABI, txInput string) {
	rawData, err   := hex.DecodeString(txInput)
	multicall, err := G_uniV3abi.MethodById(rawData[:4])
	if err != nil {
		log.Fatal(err)
	}
	if multicall.Name == "multicall0" {
		DecodeMulticall(value,G_uniV3abi, rawData)
	} else {
		fmt.Println("not multicall")
	}
}

func DecodeMulticall(value *big.Int ,G_uniV3abi abi.ABI, rawData []byte) {
	multicall, err := G_uniV3abi.MethodById(rawData[:4])
	result, err    := multicall.Inputs.UnpackValues(rawData[4:])

	resLen         := len(result[1].([][]byte))

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < resLen; i++ {
		rawFunc    := result[1].([][]byte)[i]
		funcSig, _ := G_uniV3abi.MethodById(rawFunc[:4])
		inputs, _  := funcSig.Inputs.UnpackValues(rawFunc[4:])
		if funcSig.Name == "swapExactTokensForTokens" {
			swapExactTokensForTokens := swapExactTokensForTokensParams {
				FuncName:    	funcSig.Name,
				AmountIn: 		inputs[0].(*big.Int),
				AmountOutMin: 	inputs[1].(*big.Int),
				Path:         	inputs[2].([]common.Address),
				To:           	inputs[3].(common.Address),
			}
			// if swapExactTokensForTokens.Path[0] == common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") {
			if swapExactTokensForTokens.Path[len(swapExactTokensForTokens.Path)-1] == common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") || swapExactTokensForTokens.Path[len(swapExactTokensForTokens.Path)-1] == common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				fmt.Println("Swapped some USD")
				continue}
			
			fmt.Println(swapExactTokensForTokens)
			// 	Rtkn0, Reth0 := getReservesData(swapExactTokensForTokens.Path)
			// 	success := assessProfitability(swapExactTokensForTokens.Path[1], value, swapExactTokensForTokens.AmountOutMin, Rtkn0, Reth0)
			// 	fmt.Println(success)
			// 	// go utils.SniperUniCake(0.001,0,swapExactTokensForTokens.Path[0].String(),swapExactTokensForTokens.Path[1].String())

			// 	if success {
			// 		fmt.Println(swapExactTokensForTokens)
			// 		utils.SniperUniCake(0.1,0,swapExactTokensForTokens.Path[0].String(),swapExactTokensForTokens.Path[1].String())

			// 	}

			}else if funcSig.Name == "exactInputSingle" {
			var exactInputSingle ExactInputSingleParams
			params := inputs[0].(struct { TokenIn common.Address "json:\"tokenIn\""; 
										  TokenOut common.Address "json:\"tokenOut\"";
										  Fee *big.Int "json:\"fee\"";
										  Recipient common.Address "json:\"recipient\"";
										  AmountIn *big.Int "json:\"amountIn\"";
										  AmountOutMinimum *big.Int "json:\"amountOutMinimum\"";
										  SqrtPriceLimitX96 *big.Int "json:\"sqrtPriceLimitX96\"" 
									    })

			exactInputSingle.FuncName 		   = funcSig.Name
			exactInputSingle.TokenIn           = params.TokenIn
			exactInputSingle.TokenOut          = params.TokenOut
			exactInputSingle.Fee               = params.Fee
			exactInputSingle.Recipient         = params.Recipient
			exactInputSingle.AmountIn          = params.AmountIn
			exactInputSingle.AmountOutMinimum  = params.AmountOutMinimum
			exactInputSingle.SqrtPriceLimitX96 = params.SqrtPriceLimitX96
			if exactInputSingle.TokenOut == common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") || exactInputSingle.TokenOut == common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				fmt.Println("Swapped some USD")
				continue}
			fmt.Println(exactInputSingle)

		} else if funcSig.Name == "selfPermit" {
			rArray := inputs[4].([32]byte)
			sArray := inputs[5].([32]byte)

			selfPermit := SelfPermit {
				FuncName:    funcSig.Name,
				Token: 		inputs[0].(common.Address),
				Value: 		inputs[1].(*big.Int),
				Deadline: 	inputs[2].(*big.Int),
				V:     		inputs[2].(*big.Int),
				R: 	   		hex.EncodeToString(rArray[:]),
				S: 	   		hex.EncodeToString(sArray[:]),
			}
			fmt.Println(selfPermit)

		} else if funcSig.Name == "exactOutput" {
			var exactOutputBytesPath ExactOutPutBytesPath
			exactOutputBytesPath = inputs[0].(struct {
														Path []uint8 "json:\"path\""; 
														Recipient common.Address "json:\"recipient\""; 
														AmountOut *big.Int "json:\"amountOut\""; 
														AmountInMaximum *big.Int "json:\"amountInMaximum\"";
												   	 })
			pathBytes := hex.EncodeToString([]byte(exactOutputBytesPath.Path))
			pathLen   := len(pathBytes)/44
			var pathString []string
			if pathLen < 4 {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			} else {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			}
			exactOutput := ExactOutPutStringPath {
				FuncName: 		 funcSig.Name,
				Path: 			 pathString,
				Recipient: 		 exactOutputBytesPath.Recipient, 
				AmountOut:       exactOutputBytesPath.AmountOut,
				AmountInMaximum: exactOutputBytesPath.AmountInMaximum,
			}
			if exactOutput.Path[len(exactOutput.Path)-1] == "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" || exactOutput.Path[len(exactOutput.Path)-1] == "0xdAC17F958D2ee523a2206206994597C13D831ec7" {
				fmt.Println("Swapped some USD")
				continue}
			fmt.Println(exactOutput)
		} else if funcSig.Name == "exactInput" {
			var exactInputBytesPath ExactInputBytesPath
			exactInputBytesPath = inputs[0].(struct { Path []uint8 "json:\"path\""; 
													  Recipient common.Address "json:\"recipient\""; 
													  AmountIn *big.Int "json:\"amountIn\""; 
													  AmountOutMinimum *big.Int "json:\"amountOutMinimum\"" 
													 })
			
			pathBytes := hex.EncodeToString([]byte(exactInputBytesPath.Path))
			pathLen   := len(pathBytes)/44
			var pathString []string
			if pathLen < 4 {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			} else {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			}
			exactInputs := ExactInputStringPath {
				FuncName: 			funcSig.Name,
				Path: 				pathString,
				Recipient: 			exactInputBytesPath.Recipient,
				AmountIn: 			exactInputBytesPath.AmountIn,
				AmountOutMinimum: 	exactInputBytesPath.AmountOutMinimum,
			}
			if exactInputs.Path[len(exactInputs.Path)-1] == "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" || exactInputs.Path[len(exactInputs.Path)-1] == "0xdAC17F958D2ee523a2206206994597C13D831ec7" {
				fmt.Println("Swapped some USD")
				continue}
			fmt.Println(exactInputs)
		} else if funcSig.Name == "exactOutputSingle" {
			var exactOutpuSingle exactOutpuSingleParams
			params := inputs[0].(struct { TokenIn common.Address "json:\"tokenIn\""; 
										  TokenOut common.Address "json:\"tokenOut\"";
										  Fee *big.Int "json:\"fee\"";
										  Recipient common.Address "json:\"recipient\"";
										  AmountOut *big.Int "json:\"amountOut\"";
										  AmountInMaximum *big.Int "json:\"amountInMaximum\"";
										  SqrtPriceLimitX96 *big.Int "json:\"sqrtPriceLimitX96\"" 
									    })

										exactOutpuSingle.FuncName 		   = funcSig.Name
										exactOutpuSingle.TokenIn           = params.TokenIn
										exactOutpuSingle.TokenOut          = params.TokenOut
										exactOutpuSingle.Fee               = params.Fee
										exactOutpuSingle.Recipient         = params.Recipient
										exactOutpuSingle.AmountOut          = params.AmountOut
										exactOutpuSingle.AmountInMaximum  = params.AmountInMaximum
										exactOutpuSingle.SqrtPriceLimitX96 = params.SqrtPriceLimitX96

			if exactOutpuSingle.TokenOut == common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") || exactOutpuSingle.TokenOut == common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				fmt.Println("Swapped some USD")
				continue}
			fmt.Println(exactOutpuSingle)


		} else if funcSig.Name == "swapTokensForExactTokens" {
			swapTokensForExactTokens := swapTokensForExactTokensParams {
				FuncName:    	funcSig.Name,
				AmountOut: 		inputs[0].(*big.Int),
				AmountInMax: 	inputs[1].(*big.Int),
				Path:         	inputs[2].([]common.Address),
				To:           	inputs[3].(common.Address),
			}
			if swapTokensForExactTokens.Path[len(swapTokensForExactTokens.Path)-1] == common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") || swapTokensForExactTokens.Path[len(swapTokensForExactTokens.Path)-1] == common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				fmt.Println("Swapped some USD")
			
				continue}
			fmt.Println(swapTokensForExactTokens)
		}
	}



}



func getReservesData(path []common.Address) (*big.Int, *big.Int) {
	factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), utils.GetClient())
	if err != nil {
		log.Fatal(err)
	}

	pairAddress, err := factory.GetPair(nil, path[0], path[1])
	if err != nil {
		log.Fatal(err)
	}

	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pairAddress, utils.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token0Address, err := pairToken.Token0(nil)
	if err != nil {
		log.Fatal(err)
	}

	reservesData, _ := pairToken.GetReserves(nil)
	if reservesData.Reserve0 == nil {
		return nil, nil
	}
	var Rtkn0 = new(big.Int)
	var Reth0 = new(big.Int)

	if token0Address.String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		Reth0 = reservesData.Reserve0
		Rtkn0 = reservesData.Reserve1
	} else {
		Reth0 = reservesData.Reserve1
		Rtkn0 = reservesData.Reserve0
	}
	return Rtkn0, Reth0
}


func _getAmountOut(myMaxBuy, reserveOut, reserveIn *big.Int) *big.Int {

	var myMaxBuy9975 = new(big.Int)
	var z = new(big.Int)
	num := big.NewInt(9975)
	myMaxBuy9975.Mul(num, myMaxBuy)
	num.Mul(myMaxBuy9975, reserveOut)

	den := big.NewInt(10000)
	den.Mul(den, reserveIn)
	den.Add(den, myMaxBuy9975)
	z.Div(num, den)
	return z
}


// perform the binary search to determine optimal amount of WBNB to engage on the sandwich without breaking victim's slippage
func _binarySearch(amountToTest, Rtkn0, Reth0, txValue, amountOutMinVictim *big.Int) {

	amountTknImBuying1 := _getAmountOut(amountToTest, Rtkn0, Reth0)
	var Rtkn1 = new(big.Int)
	var Reth1 = new(big.Int)
	Rtkn1.Sub(Rtkn0, amountTknImBuying1)
	Reth1.Add(Reth0, amountToTest)
	amountTknVictimWillBuy1 := _getAmountOut(txValue, Rtkn1, Reth1)

	// check if this amountToTest is really the best we can have
	// 1) we don't break victim's slippage with amountToTest
	mul10pow18, _ := new(big.Int).SetString("1000000000000000000", 10)
	maxbound := big.NewInt(int64(2))
	maxbound.Mul(maxbound, mul10pow18)
	MAXBOUND := maxbound

	var Sandwicher_baseunit = 0.02 //  BN
	mul10pow14, _ := new(big.Int).SetString("100000000000000", 10) 
	Sandwicher_baseunitx1000 := 10000 * Sandwicher_baseunit
	baseUnit := big.NewInt(int64(Sandwicher_baseunitx1000))
	baseUnit.Mul(baseUnit, mul10pow14)
	BASE_UNIT := baseUnit

	if amountTknVictimWillBuy1.Cmp(amountOutMinVictim) == 1 {
		// 2) engage MAXBOUND on the sandwich if MAXBOUND doesn't break slippage
		if amountToTest.Cmp(MAXBOUND) == 0 {
			BinaryResult = &BinarySearchResult{MAXBOUND, amountTknImBuying1, amountTknVictimWillBuy1, Rtkn1, Reth1, big.NewInt(0)}
			return
		}
		myMaxBuy := amountToTest.Add(amountToTest, BASE_UNIT)
		amountTknImBuying2 := _getAmountOut(myMaxBuy, Rtkn0, Reth0)
		var Rtkn1Test = new(big.Int)
		var Reth1Test = new(big.Int)
		Rtkn1Test.Sub(Rtkn0, amountTknImBuying2)
		Reth1Test.Add(Reth0, myMaxBuy)
		amountTknVictimWillBuy2 := _getAmountOut(txValue, Rtkn1Test, Reth1Test)
		// 3) if we go 1 step further on the ladder and it breaks the slippage, that means that amountToTest is really the amount of WBNB that we can engage and milk the maximum of profits from the sandwich.
		if amountTknVictimWillBuy2.Cmp(amountOutMinVictim) == -1 {
			BinaryResult = &BinarySearchResult{amountToTest, amountTknImBuying1, amountTknVictimWillBuy1, Rtkn1, Reth1, big.NewInt(0)}
		}
	}
	return
}

// test if we break victim's slippage with MNBOUND WBNB engaged
func _testMinbound(Rtkn, Reth, txValue, amountOutMinVictim *big.Int) int {

	mul10pow18, _ := new(big.Int).SetString("1000000000000000000", 10)
	minbound := big.NewInt(int64(1))
	minbound.Mul(minbound, mul10pow18)
	MINBOUND := minbound

	amountTknImBuying := _getAmountOut(MINBOUND, Rtkn, Reth)
	var Rtkn1 = new(big.Int)
	var Reth1 = new(big.Int)
	Rtkn1.Sub(Rtkn, amountTknImBuying)
	Reth1.Add(Reth, MINBOUND)
	amountTknVictimWillBuy := _getAmountOut(txValue, Rtkn1, Reth1)
	return amountTknVictimWillBuy.Cmp(amountOutMinVictim)
}

func getMyMaxBuyAmount2(Rtkn0, Reth0, txValue, amountOutMinVictim *big.Int, arrayOfInterest []*big.Int) {
	var wg = sync.WaitGroup{}
	// test with the minimum value we consent to engage. If we break victim's slippage with our MINBOUND, we don't go further.
	if _testMinbound(Rtkn0, Reth0, txValue, amountOutMinVictim) == 1 {
		for _, amountToTest := range arrayOfInterest {
			wg.Add(1)
			go func() {
				_binarySearch(amountToTest, Rtkn0, Reth0, txValue, amountOutMinVictim)
				wg.Done()
			}()
			wg.Wait()
		}
		return
	} else {
		BinaryResult = &BinarySearchResult{}
	}
}


func assessProfitability( tkn_adddress common.Address, txValue, amountOutMinVictim, Rtkn0, Reth0 *big.Int) bool {
	var expectedProfit = new(big.Int)

	mul10pow18, _ := new(big.Int).SetString("1000000000000000000", 10)
	minbound := big.NewInt(int64(1))
	minbound.Mul(minbound, mul10pow18)
	MINBOUND := minbound

	maxbound := big.NewInt(int64(2))
	maxbound.Mul(maxbound, mul10pow18)
	MAXBOUND := maxbound

	var Sandwicher_baseunit = 0.02 //  BN
	mul10pow14, _ := new(big.Int).SetString("100000000000000", 10) 
	Sandwicher_baseunitx1000 := 10000 * Sandwicher_baseunit
	baseUnit := big.NewInt(int64(Sandwicher_baseunitx1000))
	baseUnit.Mul(baseUnit, mul10pow14)
	BASE_UNIT := baseUnit

	var SANDWICHER_LADDER []*big.Int
	counter := new(big.Int).Set(MINBOUND)
	SANDWICHER_LADDER = append(SANDWICHER_LADDER, MINBOUND) // initial value of 1 BNB
	for counter.Cmp(MAXBOUND) != 1 {
		counter.Add(counter, BASE_UNIT)
		var toIncrement = new(big.Int).Set(counter)
		SANDWICHER_LADDER = append(SANDWICHER_LADDER, toIncrement)
	}

	arrayOfInterest := SANDWICHER_LADDER

	// only purpose of this function is to complete the struct BinaryResult via a binary search performed on the sandwich ladder we initialised in the config file. If we cannot even buy 1 BNB without breaking victim slippage, BinaryResult will be nil
	getMyMaxBuyAmount2(Rtkn0, Reth0, txValue, amountOutMinVictim, arrayOfInterest)

	if BinaryResult.MaxETHICanBuy != nil {
		var Rtkn2 = new(big.Int)
		var Reth2 = new(big.Int)
		Rtkn2.Sub(BinaryResult.Rtkn1, BinaryResult.AmountTknVictimWillBuy)
		Reth2.Add(BinaryResult.Reth1, txValue)

		// r0 --> I buy --> r1 --> victim buy --> r2 --> i sell
		// at this point of execution, we just did r2 so the "i sell" phase remains to be done
		bnbAfterSell := _getAmountOut(BinaryResult.AmountTknIWillBuy, Reth2, Rtkn2)
		expectedProfit.Sub(bnbAfterSell, BinaryResult.MaxETHICanBuy)
		fmt.Println("expectedProfit: ", expectedProfit)

		mul10pow14, _ := new(big.Int).SetString("100000000000000", 10) 
		var Sandwicher_minprofit = 0.015 
		Sandwicher_minprofitx1000 := 10000 * Sandwicher_minprofit
		minProfit := big.NewInt(int64(Sandwicher_minprofitx1000))
		minProfit.Mul(minProfit, mul10pow14)
		MINPROFIT := minProfit

		if expectedProfit.Cmp(MINPROFIT) == 1 {
			BinaryResult.ExpectedProfits = expectedProfit
			return true
		}
	}
	return false
}

