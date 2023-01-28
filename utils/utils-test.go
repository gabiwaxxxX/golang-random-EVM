package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	// "encoding/hex"

	"github.com/gabrielfournier1/bot/contract/UniswapV2Factory"
	"github.com/gabrielfournier1/bot/contract/UniswapV3Router"
	"github.com/gabrielfournier1/bot/contract/erc20"
	"github.com/gabrielfournier1/bot/data/manageData"
	"github.com/gabrielfournier1/bot/eth_const"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// to use wen you will have to swap that you don't known
func SniperUniCakeTestV0(amountIn float64, amountOutMin float64, token0Addr string, token1Addr string) {
	// calculte time execution
	var times []time.Duration
	for i := 0; i < 100; i++ {
		t1 := time.Now()
		UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetFlash())

		factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), GetClient())
		if err != nil {
			log.Fatal(err)
		}

		pairAddress, err := factory.GetPair(nil, common.HexToAddress(token0Addr), common.HexToAddress(token1Addr))
		if err != nil {
			log.Fatal(err)
		}
		var zeroAdress = common.Address{}
		if pairAddress == zeroAdress {
			log.Printf("pair not createe or not founded")
			//finding path
			fmt.Println("finding path")
			return
		}

		//get decimals for token0
		token0, err := erc20.NewErc20(common.HexToAddress(token0Addr), GetClient())
		if err != nil {
			log.Fatal(err)
		}
		decimal0, err := token0.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		token1, err := erc20.NewErc20(common.HexToAddress(token1Addr), GetClient())
		if err != nil {
			log.Fatal(err)
		}
		decimal1, err := token1.Decimals(&bind.CallOpts{})

		amountInF, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		amountOutF, _ := big.NewFloat(amountOutMin * PowInt(10, int(decimal1))).Int(nil)

		// fmt.Println("amount input : ",amountInF )
		// fmt.Println("amount min output : ",amountOutF)
		// fmt.Println("from address : ",GetUserInfo().FromAddress.String())
		// fmt.Println("token0 : ",token0Addr)
		// fmt.Println("token1 : ",token1Addr)

		// fmt.Println("Preparing data to swap")
		// make a function with this inputs : amountInF / amountOutF / FromAddress / [token0Addr, token1Addr]
		PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)}, "WETH"})
		// fmt.Println("data prepared")

		auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
		if err != nil {
			log.Fatal(err)
		}

		test, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		if token0Addr == eth_const.FindAddressFromTokenSymbol("WETH") {
			auth.Value = test
		}
		auth.GasLimit = 500000
		auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

		//add 10% to auth.GasPrice *big.Int
		auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))

		elapsed := time.Since(t1)
		times = append(times, elapsed)
		// fmt.Println("elapsed time:",elapsed)
		// fmt.Println(data,router)
	}
	fmt.Println("average time : ", average(times))
}

// to use wen you will have to swap known tokens
func SniperUniCakeTestV1(amountIn float64, amountOutMin float64, token0Addr string, token1Addr string) {
	// calculte time execution
	var times []time.Duration
	for i := 0; i < 100; i++ {
		t1 := time.Now()
		UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetFlash())

		factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), GetClient())
		if err != nil {
			log.Fatal(err)
		}

		pairAddress, err := factory.GetPair(nil, common.HexToAddress(token0Addr), common.HexToAddress(token1Addr))
		if err != nil {
			log.Fatal(err)
		}
		var zeroAdress = common.Address{}
		if pairAddress == zeroAdress {
			log.Printf("pair not createe or not founded")
			//finding path
			fmt.Println("finding path")
			return
		}

		//get decimals for token0
		var decimal0 uint8 = manageData.GetTokenDecimals(token0Addr)
		if decimal0 == 0 {
			// log.Printf("token0 not founded")
			token0, err := erc20.NewErc20(common.HexToAddress(token0Addr), GetClient())
			if err != nil {
				log.Fatal(err)
			}
			decimal0, err = token0.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}
		}
		//get decimals for token1
		var decimal1 = manageData.GetTokenDecimals(token1Addr)
		if decimal1 == 0 {
			// log.Printf("token1 not founded")
			token1, err := erc20.NewErc20(common.HexToAddress(token1Addr), GetClient())
			if err != nil {
				log.Fatal(err)
			}
			decimal1, err = token1.Decimals(&bind.CallOpts{})
		}

		amountInF, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		amountOutF, _ := big.NewFloat(amountOutMin * PowInt(10, int(decimal1))).Int(nil)

		// fmt.Println("amount input : ",amountInF )
		// fmt.Println("amount min output : ",amountOutF)
		// fmt.Println("from address : ",GetUserInfo().FromAddress.String())
		// fmt.Println("token0 : ",token0Addr)
		// fmt.Println("token1 : ",token1Addr)

		// fmt.Println("Preparing data to swap")
		// make a function with this inputs : amountInF / amountOutF / FromAddress / [token0Addr, token1Addr]
		PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)}, "WETH"})
		// fmt.Println("data prepared")

		auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
		if err != nil {
			log.Fatal(err)
		}

		test, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		if token0Addr == eth_const.FindAddressFromTokenSymbol("WETH") {
			auth.Value = test
		}
		auth.GasLimit = 500000
		auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

		//add 10% to auth.GasPrice *big.Int
		auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))

		elapsed := time.Since(t1)
		times = append(times, elapsed)
		// fmt.Println("elapsed time:",elapsed)
		// fmt.Println(data,router)
	}
	fmt.Println("average time : ", average(times))
}

func SniperUniCakeTestV2(amountIn float64, amountOutMin float64, token0Addr string, token1Addr string) {

	// calculte time execution
	var times []time.Duration
	for i := 0; i < 100; i++ {
		time1 := time.Now()
		UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetFlash())

		factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), GetClient())
		if err != nil {
			log.Fatal(err)
		}

		pairAddress, err := factory.GetPair(nil, common.HexToAddress(token0Addr), common.HexToAddress(token1Addr))
		if err != nil {
			log.Fatal(err)
		}
		var zeroAdress = common.Address{}
		if pairAddress == zeroAdress {
			log.Printf("pair not createe or not founded")
			//finding path
			fmt.Println("finding path")
			return
		}

		//get decimals for token0
		var decimal0 uint8 = manageData.GetTokenDecimals(token0Addr)
		if decimal0 == 0 {
			// log.Printf("token0 not founded")
			token0, err := erc20.NewErc20(common.HexToAddress(token0Addr), GetClient())
			if err != nil {
				log.Fatal(err)
			}
			decimal0, err = token0.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}
		}
		//get decimals for token1
		var decimal1 = manageData.GetTokenDecimals(token1Addr)
		if decimal1 == 0 {
			// log.Printf("token1 not founded")
			token1, err := erc20.NewErc20(common.HexToAddress(token1Addr), GetClient())
			if err != nil {
				log.Fatal(err)
			}
			decimal1, err = token1.Decimals(&bind.CallOpts{})
		}

		amountInF, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		amountOutF, _ := big.NewFloat(amountOutMin * PowInt(10, int(decimal1))).Int(nil)
		rawData := PrepareRawDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)}, "WETH"})
		// fmt.Println("raw data : ",rawData)

		nonce, err := GetClient().PendingNonceAt(context.Background(), GetUserInfo().FromAddress)
		if err != nil {
			log.Fatal(err)
		}
		value, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		toAddress := common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
		gasLimit := uint64(500000)
		gasPrice, err := GetClient().SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		gasPrice = new(big.Int).Div(new(big.Int).Mul(gasPrice, big.NewInt(200)), big.NewInt(100))

		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, rawData)

		chainID, err := GetClient().NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		types.SignTx(tx, types.NewEIP155Signer(chainID), GetUserInfo().PrivateKey)
		if err != nil {
			log.Fatal(err)
		}

		// ts := types.Transactions{signedTx}
		// rawTxBytes := ts.EncodeRLP(0)
		// // rawTxHex := hex.EncodeToString(rawTxBytes)

		// txToSend := new(types.Transaction)
		// rlp.DecodeBytes(rawTxBytes, &txToSend)

		elapsed := time.Since(time1)
		// fmt.Println("elapsed time : ",elapsed)
		times = append(times, elapsed)
		// err = GetClient().SendTransaction(context.Background(), signedTx)
		// if err != nil {
		//     log.Fatal(err)
		// }

		// data:=PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)},"WETH"})
		// fmt.Println("data : ",data)
	}
	fmt.Println("average time : ", average(times))
}

func SniperUniCakeTestV3(amountIn float64, amountOutMin float64, token0Addr string, token1Addr string) {
	// calculte time execution
	var times []time.Duration
	for i := 0; i < 100; i++ {

		t1 := time.Now()
		UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetFlash())

		factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), GetClient())
		if err != nil {
			log.Fatal(err)
		}

		pairAddress, err := factory.GetPair(nil, common.HexToAddress(token0Addr), common.HexToAddress(token1Addr))
		if err != nil {
			log.Fatal(err)
		}
		var zeroAdress = common.Address{}
		if pairAddress == zeroAdress {
			log.Printf("pair not createe or not founded")
			//finding path
			fmt.Println("finding path")
			return
		}

		token0, err := erc20.NewErc20(common.HexToAddress(token0Addr), GetClient())
		if err != nil {
			log.Fatal(err)
		}
		decimal0, err := token0.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		token1, err := erc20.NewErc20(common.HexToAddress(token1Addr), GetClient())
		if err != nil {
			log.Fatal(err)
		}
		decimal1, err := token1.Decimals(&bind.CallOpts{})

		amountInF, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		amountOutF, _ := big.NewFloat(amountOutMin * PowInt(10, int(decimal1))).Int(nil)

		// fmt.Println("amount input : ",amountInF )
		// fmt.Println("amount min output : ",amountOutF)
		// fmt.Println("from address : ",GetUserInfo().FromAddress.String())
		// fmt.Println("token0 : ",token0Addr)
		// fmt.Println("token1 : ",token1Addr)

		// make a function with this inputs : amountInF / amountOutF / FromAddress / [token0Addr, token1Addr]
		PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)}, "WETH"})

		auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
		if err != nil {
			log.Fatal(err)
		}

		test, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
		if token0Addr == eth_const.FindAddressFromTokenSymbol("WETH") {
			auth.Value = test
		}
		auth.GasLimit = 500000
		auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

		//add 10% to auth.GasPrice *big.Int
		auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))

		elapsed := time.Since(t1)
		// fmt.Println("elapsed time:",elapsed)
		times = append(times, elapsed)
		//instance of router will be with the flashbot RPC
		// tx, err := router.Multicall0(auth,big.NewInt(time.Now().Add(5*time.Minute).Unix()), data)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// awaitToConfirmTx(tx)

		// if(token1Addr != eth_const.FindAddressFromTokenSymbol("WETH") || token1Addr != eth_const.FindAddressFromTokenSymbol("DAI")){
		// 	ApproveSpenderToken(token1Addr,tx.To().Hex())
		// 	TrackGains(amountIn,[]common.Address{common.HexToAddress(token1Addr), common.HexToAddress(token0Addr)})

		// }
	}
	fmt.Println("average time : ", average(times))

}

func PrepareDataToSwapTest(input TransactionInput) [][]byte {
	var data [][]byte
	// fmt.Println("Preparing data to swap",GetSwapRouterInput(input))

	var line []byte = common.Hex2Bytes(strings.Split(GetSwapRouterInput(input), "x")[1])
	data = append(data, line)
	if input.Path[1].String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		var secondInput string
		secondInput = secondInput + "0x49404b7c"
		secondInput = secondInput + Pad(input.MinAmountOut)
		secondInput = secondInput + Pad(strings.Split(input.To.String(), "x")[1])
		data = append(data, common.Hex2Bytes(strings.Split(secondInput, "x")[1]))
	}
	return data
}

func PrepareRawDataToSwap(input TransactionInput) []byte {
	//firstfunction signature : 5ae401dc
	var data string
	data = data + "5ae401dc"
	deadline := time.Now().Add(5 * time.Minute).Unix()
	data = data + pad32Bytes(toHexInt64(int(deadline)))
	data = data + pad32Bytes(toHexInt64(64))
	data = data + pad32Bytes(toHexInt64(1))
	data = data + pad32Bytes(toHexInt64(32))
	data = data + pad32Bytes(toHexInt64(228))
	data = data + "472b43f3"
	data = data + pad32Bytes(input.AmountIn)
	data = data + pad32Bytes(input.MinAmountOut)
	data = data + pad32Bytes(toHexInt64(128))
	if input.Path[1].String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		data = data + pad32Bytes(toHexInt64(2))
	} else {
		data = data + pad32Bytes(strings.Split(input.To.String(), "x")[1])
	}
	data = data + pad32Bytes(toHexInt64(len(input.Path)))
	for _, value := range input.Path {
		data = data + pad32Bytes(strings.Split(value.String(), "x")[1])
	}
	if input.Path[1].String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		data = data + pad32Bytes(toHexInt64(0))
		data = data + pad32Bytes("4449404b7c")
		data = data + pad32Bytes(input.MinAmountOut)
		data = data + pad32Bytes(strings.Split(input.To.String(), "x")[1])
	}
	return common.Hex2Bytes(data)

}

func average(times []time.Duration) time.Duration {
	var sum time.Duration
	for _, time := range times {
		sum += time
	}
	return sum / time.Duration(len(times))
}

func toHexInt64(n int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func pad32Bytes(data string) string {
	return strings.Repeat("0", 64-len(data)) + data
}
