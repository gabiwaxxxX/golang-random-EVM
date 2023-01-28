package arbitrum

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/contract/Router01"
	"github.com/gabrielfournier1/bot/utils"
)

func SwapExactTokenForToken() {

	client := GetClient()
	routerCaller, err := Router01.NewRouter01Caller(common.HexToAddress("0xfa3e4050dd383b8df1990831e59a82c350379bd7"), GetClient())
	if err != nil {
		panic(err)
	}

	//loop

	//check value of Factory
	factory, err := routerCaller.Factory(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Factory: %v", factory)
	//0x6CC691fbFEf23a669452a21ad8296D1DCCC330eC
	// if factory == common.HexToAddress("0x0000000000000000000000000000000000000000") {
	// 	panic("Factory is not correct")
	// } else {
	// 	fmt.Println("Factory is correct")
	// 	//get pair address
	// 	pair, err := routerCaller.PairFor(nil, common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"), common.HexToAddress("0xFd808d72761Ad2423F100c53Aa2F2c7F133dEF0a"), false)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Pair: %v", pair)
	// 	if pair == common.HexToAddress("0x0000000000000000000000000000000000000000") {
	// 		panic("Pair is not correct")
	// 	} else {
	// 		fmt.Println("Pair is correct")
	// 	}
	// }

	amountIn := big.NewInt(500000000)
	// amountIn := big.NewInt(5000)
	amountOutMin := big.NewInt(0)
	path := Router01.Router01Route{common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"), common.HexToAddress("0xFd808d72761Ad2423F100c53Aa2F2c7F133dEF0a"), false}
	var routes []Router01.Router01Route
	routes = append(routes, path)
	deadline := big.NewInt(time.Now().Add(15 * time.Minute).Unix())

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	router, err := Router01.NewRouter01(common.HexToAddress("0xfa3e4050dd383b8df1990831e59a82c350379bd7"), client)
	if err != nil {
		panic(err)
	}

	tx, err := router.SwapExactTokensForTokensSupportingFeeOnTransferTokens(auth, amountIn, amountOutMin, routes, common.HexToAddress("Your wallet adress"), deadline)
	if err != nil {
		panic(err)
	}

	bind.WaitMined(context.Background(), GetClient(), tx)
	receipt, err := GetClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("not founded", err)
	}
	fmt.Println("Transaction receipt: %v", receipt.Status)
	fmt.Println("Transaction receipt: %v", receipt.TxHash)

}
