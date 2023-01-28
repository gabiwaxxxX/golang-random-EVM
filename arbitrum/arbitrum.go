package arbitrum

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gabrielfournier1/bot/contract/RanbasedToken"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Pair"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Router"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Router02"
	"github.com/gabrielfournier1/bot/contract/erc20"
	"github.com/gabrielfournier1/bot/data/Router"
	"github.com/gabrielfournier1/bot/utils"
)

func HexToInt(hex string) *big.Int {
	i := new(big.Int)
	i.SetString(hex, 16)
	return i
}

func GetCurrentNonce(addr common.Address) (*big.Int, error) {
	client := GetClient()
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(nonce)), nil
}

func GetClient() *ethclient.Client {
	rpcAddress := "https://rpc.ankr.com/arbitrum"
	// rpcAddress := "https://arb-mainnet.g.alchemy.com/v2/demo"
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SwapExactTokensForETH() {
	// client := GetClient()

	amountIn := big.NewInt(100000000000000) // 1 ETH
	amountOutMin := big.NewInt(0)
	path := []common.Address{common.HexToAddress("0xC6eb4a24CD2A5914e6714812Aa6593Ae08c4b374"), common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")}
	to := common.HexToAddress("Your wallet adress")
	deadline := big.NewInt(time.Now().Add(15 * time.Minute).Unix())

	router, err := UniswapV2Router02.NewUniswapV2Router02(common.HexToAddress("0x1b02da8cb0d097eb8d57a175b88c7d8b47997506"), GetClient())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	fmt.Println("Gas Price: ", auth.GasPrice)
	fmt.Println("Gas Limit: ", auth.GasLimit)

	// add 10% to auth.GasPrice *big.Int
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))

	tx, err := router.SwapExactTokensForETH(auth, amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		log.Fatal(err)
	}
	// utils.AwaitToConfirmTx(tx)
	bind.WaitMined(context.Background(), GetClient(), tx)
	receipt, err := GetClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("not founded", err)
	}
	fmt.Printf("Transaction receipt: %v", receipt)

}

func CheckPause() {
	tokenRebase := common.HexToAddress("0x7dbf12f21B00F78587d03BF728A4885eA04739Fc")
	contract, err := RanbasedToken.NewRanbasedToken(tokenRebase, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	for {

		pause, err := contract.PAUSED(nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Pause: ", pause)
		if !pause {
			BuyRandFromEth()
		}
		time.Sleep(250 * time.Millisecond)
	}

}

func BuyRandFromEth() {
	fmt.Println("BuyRandFromEth")
	//weth to rand
	tokenIn := common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
	tokenOut := common.HexToAddress("0x7dbf12f21B00F78587d03BF728A4885eA04739Fc")
	// tokenOut := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")

	path := []common.Address{tokenIn, tokenOut}
	to := common.HexToAddress("Your wallet adress")
	deadline := big.NewInt(time.Now().Add(15 * time.Minute).Unix())

	router, err := UniswapV2Router02.NewUniswapV2Router02(common.HexToAddress("0x1b02da8cb0d097eb8d57a175b88c7d8b47997506"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	Value := big.NewInt(100000000000000000) // 0.1 ETH // 0.001 eth to usdc
	// Value := big.NewInt(1000000000000000) // 0.1 ETH // 0.001 eth to usdc
	amountIn := new(big.Int).Mul(Value, big.NewInt(3))

	amountOutMin, err := router.GetAmountsOut(nil, Value, path)
	if err != nil {
		log.Fatal(err)
	}
	if len(amountOutMin) == 0 {
		log.Fatal("amountOutMin is empty")
	}
	fmt.Println("amountOutMin: ", amountOutMin)
	// amountOutMin - 10 %
	amountOutMin2 := new(big.Int).Div(new(big.Int).Mul(amountOutMin[1], big.NewInt(90)), big.NewInt(100))
	fmt.Println("amountOutMin2: ", amountOutMin2)

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 4000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth.Value = Value
	auth.Nonce, err = GetCurrentNonce(common.HexToAddress("Your wallet adress"))
	fmt.Println(" Nonce: ", auth.Nonce)

	if err != nil {
		log.Fatal(err)
	}

	// 0.1 ETH
	wg := sync.WaitGroup{}

	for i := 0; i < 1; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			tx, err := router.SwapExactETHForTokensSupportingFeeOnTransferTokens(auth, amountOutMin2, path, to, deadline)
			if err != nil {
				log.Fatal(err)
			}
			bind.WaitMined(context.Background(), GetClient(), tx)
			receipt, err := GetClient().TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal("not founded", err)
			}
			fmt.Println("Transaction receipt: ", receipt.Status)

			// var Nonce = auth.Nonce
			auth.Nonce = new(big.Int).Add(auth.Nonce, big.NewInt(1))
		}()

		wg.Wait()
	}
	// auth.Nonce = new(big.Int).Add(auth.Nonce, big.NewInt(1))
	ApproveSpenderToken(auth.Nonce, tokenOut.String(), "0x1b02da8cb0d097eb8d57a175b88c7d8b47997506")

	TrackGains(amountIn, []common.Address{tokenOut, tokenIn})

}

func SwapExactETHForTokens() {
	// client := GetClient()

	amountOutMin := big.NewInt(0)
	path := []common.Address{common.HexToAddress("0x58ea7917f74834dbe6b57d0a2a74fb68c1e94c55"), common.HexToAddress("0x30dcBa0405004cF124045793E1933C798Af9E66a")}
	to := common.HexToAddress("Your wallet adress")
	deadline := big.NewInt(time.Now().Add(15 * time.Minute).Unix())

	router, err := UniswapV2Router02.NewUniswapV2Router02(common.HexToAddress("0x1b02da8cb0d097eb8d57a175b88c7d8b47997506"), GetClient())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(400000000000000000) // 1 ETH

	fmt.Println("Gas Price: ", auth.GasPrice)
	fmt.Println("Gas Limit: ", auth.GasLimit)

	// add 10% to auth.GasPrice *big.Int
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))

	tx, err := router.SwapExactETHForTokens(auth, amountOutMin, path, to, deadline)
	if err != nil {
		log.Fatal(err)
	}
	// utils.AwaitToConfirmTx(tx)
	bind.WaitMined(context.Background(), GetClient(), tx)
	receipt, err := GetClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("not founded", err)
	}
	fmt.Printf("Transaction receipt: %v", receipt)

}

// Check price of a token buy dividing my amount of USDC and the expected amount token if I buy it

func CheckPriceOfTheToken() {

	usdcAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	xcalAddr := common.HexToAddress("0xd2568acCD10A4C98e87c44E9920360031ad89fCB")

	router, err := Router.NewRouter(common.HexToAddress("0x8e72bf5A45F800E182362bDF906DFB13d5D5cb5d"), GetClient())
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	token, err := erc20.NewErc20(usdcAddr, GetClient())
	if err != nil {
		log.Fatal("error 2 ", err)

	}

	decimal1, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("error 3 ", err)

	}

	balance, err := token.BalanceOf(nil, common.HexToAddress("Your wallet adress"))
	if err != nil {
		log.Fatal("error 4 ", err)

	}
	//my balance of USDC
	usdcBalance := utils.WeiToEtherFloatByDecimals(int(decimal1), balance)

	//while loop to compute the price
	var PreviousPrice *big.Float = big.NewFloat(0.43472046526231838410284)
	for {
		amountOut, err := router.GetAmountOut(nil, balance, usdcAddr, xcalAddr)
		if err != nil {
			log.Fatal("error 5 ", err)

		}

		tokenOut, err := erc20.NewErc20(xcalAddr, GetClient())
		if err != nil {
			log.Fatal(err)
		}
		decimal0, err := tokenOut.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		//total amount expected after
		expectedAmount := utils.WeiToEtherFloatByDecimals(int(decimal0), amountOut.Amount)
		//compute the price of the token in usdc
		price := new(big.Float).Quo(usdcBalance, expectedAmount)
		//compute the difference between the previous price and the current price
		diff := new(big.Float).Sub(price, PreviousPrice)
		//compute the percentage of the difference
		percentage := new(big.Float).Quo(diff, PreviousPrice)
		//compute the percentage of the difference in percentage
		percentage = new(big.Float).Mul(percentage, big.NewFloat(100))
		//print the price
		if percentage.Cmp(big.NewFloat(0)) != 0 {
			fmt.Println("Price: ", price)
			//print the difference
			fmt.Println("Difference: ", diff)
			//print the percentage of the difference
			fmt.Println("Percentage: ", percentage)
			//time
			fmt.Println("Time: ", time.Now().Format("15:04:05"))
			fmt.Println("")

			if percentage.Cmp(big.NewFloat(-13)) == -1 {
				fmt.Println("Buy !")
				BuyWhenThePriceDrop()
				fmt.Println("==========================")
				// f, err := os.Open("Rockafeller Skank.mp3")
				// if err != nil {
				// 	log.Fatal(err)
				// }

				// streamer, format, err := mp3.Decode(f)
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}
		}

		//set the previous price to the current price
		PreviousPrice = price

		sleepTime := time.Duration(100) * time.Millisecond
		time.Sleep(sleepTime)
		// fmt.Println("price: ", price, " $")
	}
}

func CheckingPriceOfTokenkByLp() {

	usdcAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")

	LPAddress := common.HexToAddress("0x2Cc6AC1454490AfA83333Fabc84345FaD751285B")
	LP, err := UniswapV2Pair.NewUniswapV2Pair(LPAddress, GetClient())
	if err != nil {
		log.Fatal("error 1 ", err)
	}

	reservers, err := LP.GetReserves(nil)
	if err != nil {
		log.Fatal("error 2 ", err)
	}

	token0Addr, err := LP.Token0(nil)
	if err != nil {
		log.Fatal("error 3 ", err)
	}

	token0, err := erc20.NewErc20(token0Addr, GetClient())
	if err != nil {
		log.Fatal("error 4 ", err)

	}

	decimal0, err := token0.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("error 5 ", err)

	}

	token1Addr, err := LP.Token1(nil)
	if err != nil {
		log.Fatal("error 6 ", err)
	}

	token1, err := erc20.NewErc20(token1Addr, GetClient())
	if err != nil {
		log.Fatal("error 7 ", err)
	}

	decimal1, err := token1.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("error 8 ", err)
	}
	var previousReserve *big.Float = big.NewFloat(0)
	var otherReserve *big.Float = big.NewFloat(0)
	if token0Addr == usdcAddr {
		fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0))
	} else {
		fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1))
	}

	for {
		reservers, err := LP.GetReserves(nil)
		if err != nil {
			log.Fatal("error 9 ", err)
		}

		if token0Addr == usdcAddr {
			if previousReserve.Cmp(utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0)) != 0 {
				currentReserve := utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0)
				currentOthreserve := utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1)

				diff := new(big.Float).Sub(currentReserve, previousReserve)
				if diff.Cmp(big.NewFloat(0)) == 1 {
					if otherReserve.Cmp(currentOthreserve) == 1 {
						fmt.Println("Someone bought xcal")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					} else {
						fmt.Println("Someone added liquidity")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					}
				}
				if diff.Cmp(big.NewFloat(0)) == -1 {
					if otherReserve.Cmp(currentOthreserve) == -1 {
						fmt.Println("Someone sold xcal")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
						if diff.Cmp(big.NewFloat(-11000)) == -1 {
							fmt.Println("Buy !")
							BuyWhenThePriceDrop()
							fmt.Println("==========================")
						}

					} else {
						fmt.Println("Someone removed liquidity")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					}
				}

				previousReserve = currentReserve
				otherReserve = currentOthreserve

			}
		} else {
			if previousReserve.Cmp(utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1)) != 0 {
				currentReserve := utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1)
				currentOthreserve := utils.WeiToEtherFloatByDecimals(int(decimal0), reservers.Reserve0)
				diff := new(big.Float).Sub(currentReserve, previousReserve)
				if diff.Cmp(big.NewFloat(0)) == 1 {
					if otherReserve.Cmp(currentOthreserve) == 1 {
						color.Green("Someone bought xcal")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					} else {
						color.Blue("Someone added liquidity")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					}
				}
				if diff.Cmp(big.NewFloat(0)) == -1 {
					if otherReserve.Cmp(currentOthreserve) == -1 {
						color.Red("Someone sold xcal")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
						if diff.Cmp(big.NewFloat(-11000)) == -1 {
							fmt.Println("Buy !")
							BuyWhenThePriceDrop()
							fmt.Println("==========================")
						}
					} else {
						color.Yellow("Someone removed liquidity")
						fmt.Println("USDC: ", utils.WeiToEtherFloatByDecimals(int(decimal1), reservers.Reserve1))
						fmt.Println("Difference: ", diff)
						fmt.Println("Time: ", time.Now().Format("15:04:05"))
						fmt.Println("")
					}
				}
				previousReserve = currentReserve
				otherReserve = currentOthreserve

			}
		}

		sleepTime := time.Duration(100) * time.Millisecond
		time.Sleep(sleepTime)
	}

}

func BuyWhenThePriceDrop() {
	usdcAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	xcalAddr := common.HexToAddress("0xd2568acCD10A4C98e87c44E9920360031ad89fCB")

	router, err := Router.NewRouter(common.HexToAddress("0x8e72bf5A45F800E182362bDF906DFB13d5D5cb5d"), GetClient())
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	token, err := erc20.NewErc20(usdcAddr, GetClient())
	if err != nil {
		log.Fatal("error 2 ", err)

	}
	balance, err := token.BalanceOf(nil, common.HexToAddress("Your wallet adress"))
	if err != nil {
		log.Fatal("error 4 ", err)

	}
	amountIn := balance
	// amountIn := big.NewInt(1)
	// amountIn := big.NewInt(1000000000)

	amountOut, err := router.GetAmountOut(nil, balance, usdcAddr, xcalAddr)
	if err != nil {
		log.Fatal("error 5 ", err)

	}

	amountOutMin := big.NewInt(0)

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	var route []Router.Routerroute
	var routerroute Router.Routerroute
	routerroute.From = usdcAddr
	routerroute.To = xcalAddr
	routerroute.Stable = amountOut.Stable
	route = append(route, routerroute)
	tx, err := router.SwapExactTokensForTokens(auth, amountIn, amountOutMin, route, common.HexToAddress("Your wallet adress"), big.NewInt(time.Now().Unix()+1000))
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)

	TrackGains(amountIn, []common.Address{xcalAddr, usdcAddr})

	//buy when the price drop

}

func AwaitToConfirmTx(tx *types.Transaction) {
	t1 := time.Now()
	bind.WaitMined(context.Background(), GetClient(), tx)
	receipt, err := GetClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("not founded", err)
	}
	t2 := time.Now()

	diff := t2.Sub(t1)

	if receipt.Status == 1 {
		log.Printf("Transaction has been mined, tx hash: %s, time: %s", tx.Hash().Hex(), diff)

	} else {
		fmt.Println("Transaction has been rejected")
	}

}

// func TrackGains_2(amountIn float64, path []common.Address) {
// 	fmt.Println("tracking gains")
// 	go readInput_1(amountIn, path)
// 	for ok := true; ok; ok = true {
// 		if amountIn > 0 {
// 			expectedAmountOut := CheckExpectedAmountOut(path)
// 			gain_not_string := new(big.Float).Mul(new(big.Float).Sub(new(big.Float).Quo(expectedAmountOut, big.NewFloat(amountIn)), big.NewFloat(1)), big.NewFloat(100))
// 			fmt.Println("gain: ", gain_not_string)
// 			gainPercentage := gain_not_string.String()
// 			fmt.Println(amountIn, "========>", expectedAmountOut.String(), " ======  gains ", gainPercentage, "%")

// 			if gain_not_string.Cmp(big.NewFloat(100)) > 1 {
// 				fmt.Println("gains are more than 100%, selling half")
// 				SellTokensExpress(path)
// 			}
// 			time.Sleep(time.Second * 1)
// 		} else {
// 			expectedAmountOut := CheckExpectedAmountOut(path)
// 			fmt.Println(amountIn, "========>", expectedAmountOut.String(), " ======  gains ")
// 			time.Sleep(time.Second * 1)
// 		}
// 	}
// }

// func TrackGains_1(amountIn *big.Int, path []common.Address) {
// 	fmt.Println("tracking gains")
// 	go readInput()
// 	for ok := true; ok; ok = true {
// 		expectedAmountOut := CheckExpectedAmountOut()
// 		fmt.Println("expectedAmountOut: ", expectedAmountOut)
// 		gain_not_string := new(big.Float).Mul(new(big.Float).Sub(new(big.Float).Quo(expectedAmountOut, new(big.Float).SetInt(amountIn)), big.NewFloat(1)), big.NewFloat(100))
// 		fmt.Println("gain: ", gain_not_string)
// 		if gain_not_string.Cmp(big.NewFloat(10)) == 1 {
// 			fmt.Println("Sell !")
// 			SellTokens()
// 		}

// 		time.Sleep(time.Second * 1)
// 	}
// }

func TrackGains(amountIn *big.Int, path []common.Address) {
	fmt.Println("tracking gains")
	go readInput()
	for ok := true; ok; ok = true {
		expectedAmountOut := CheckExpectedAmountOut()
		fmt.Println("expectedAmountOut: ", expectedAmountOut)
		gain_not_string := new(big.Float).Mul(new(big.Float).Sub(new(big.Float).Quo(expectedAmountOut, new(big.Float).SetInt(amountIn)), big.NewFloat(1)), big.NewFloat(100))
		fmt.Println("gain: ", gain_not_string)
		if gain_not_string.Cmp(big.NewFloat(10)) == 1 {
			fmt.Println("Sell !")
			SellTokens()
		}

		time.Sleep(time.Second * 1)
	}
}

func CheckExpectedAmountOut_1(path []common.Address) *big.Float {

	router, err := UniswapV2Router.NewUniswapV2Router(common.HexToAddress("0x1b02da8cb0d097eb8d57a175b88c7d8b47997506"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token, err := erc20.NewErc20(path[0], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, utils.GetUserInfo().FromAddress)
	if err != nil {
		fmt.Printf("error 1: %v")

		balance = big.NewInt(1)
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("balance is 0")
		os.Exit(1)
	}

	amountOut, err := router.GetAmountsOut(nil, balance, path)
	if err != nil {
		fmt.Println("error 2: ", err)
		amountOut = []*big.Int{big.NewInt(0), big.NewInt(0)}
	}

	tokenOut, err := erc20.NewErc20(path[1], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	decimal0, err := tokenOut.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return (weiToEtherFloatByDecimals(int(decimal0), amountOut[1]))

}

func weiToEtherFloatByDecimals(decimals int, wei *big.Int) *big.Float {
	if decimals == 0 {
		decimals = 18
	}
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(decimals)))
}

func CheckExpectedAmountOut() *big.Float {

	usdcAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	xcalAddr := common.HexToAddress("0xd2568acCD10A4C98e87c44E9920360031ad89fCB")

	router, err := Router.NewRouter(common.HexToAddress("0x8e72bf5A45F800E182362bDF906DFB13d5D5cb5d"), GetClient())
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	token, err := erc20.NewErc20(xcalAddr, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, common.HexToAddress("Your wallet adress"))
	if err != nil {
		log.Fatal(err)
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("balance is 0")
		os.Exit(1)
	}

	amountOut, err := router.GetAmountOut(nil, balance, xcalAddr, usdcAddr)
	if err != nil {
		log.Fatal(err)
	}

	tokenOut, err := erc20.NewErc20(usdcAddr, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	decimal0, err := tokenOut.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return (utils.WeiToEtherFloatByDecimals(int(decimal0), amountOut.Amount))

}

func readInput_1(amountIn float64, path []common.Address) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("could not process input %v\n", input)
		}
		switch string(input) {
		case "sell":
			fmt.Println("selling your amount of tokens")
			SellTokens()

		case "s":
			fmt.Println("selling your amount of tokens")
			SellTokens()
		}
	}
}

func readInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("could not process input %v\n", input)
		}
		switch string(input) {
		case "sell":
			fmt.Println("selling your amount of tokens")
			SellTokens()

		case "s":
			fmt.Println("selling your amount of tokens")
			SellTokens()
		}
	}
}

func SellTokens() {
	usdcAddr := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	xcalAddr := common.HexToAddress("0xd2568acCD10A4C98e87c44E9920360031ad89fCB")

	router, err := Router.NewRouter(common.HexToAddress("0x8e72bf5A45F800E182362bDF906DFB13d5D5cb5d"), GetClient())
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	token, err := erc20.NewErc20(xcalAddr, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, common.HexToAddress("Your wallet adress"))
	if err != nil {
		log.Fatal(err)
	}
	amountOut, err := router.GetAmountOut(nil, balance, xcalAddr, usdcAddr)
	if err != nil {
		log.Fatal("error 5 ", err)

	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	var route []Router.Routerroute
	var routerroute Router.Routerroute
	routerroute.From = xcalAddr
	routerroute.To = usdcAddr
	routerroute.Stable = amountOut.Stable
	tx, err := router.SwapExactTokensForTokens(auth, balance, amountOut.Amount, route, common.HexToAddress("Your wallet adress"), big.NewInt(time.Now().Unix()+1000))
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)
	CheckingPriceOfTokenkByLp()
}

func ApproveSpenderToken(nonce *big.Int, tokenToApprove string, spender string) {

	fmt.Println("approving spender")

	token, err := erc20.NewErc20(common.HexToAddress(tokenToApprove), GetClient())
	if err != nil {
		log.Fatal(err)
	}

	amount := HexToInt("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

	id, err := GetClient().ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, id)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 1000000
	auth.Nonce = nonce

	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(130)), big.NewInt(100))

	tx, err := token.Approve(auth, common.HexToAddress(spender), amount)
	if err != nil {
		log.Fatal(err)
	}

	AwaitToConfirmTx(tx)

}
