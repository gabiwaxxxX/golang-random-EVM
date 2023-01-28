package utils

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabrielfournier1/bot/contract/SURGE"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Factory"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Router"
	"github.com/gabrielfournier1/bot/contract/UniswapV3Router"
	"github.com/gabrielfournier1/bot/contract/erc20"
	"github.com/gabrielfournier1/bot/eth_const"
	"github.com/gabrielfournier1/bot/wifi_managment"

	// "github.com/gabrielfournier1/bot/data/manageData"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
)

type TransactionInput struct {
	AmountIn        string
	MinAmountOut    string
	To              common.Address
	Path            []common.Address
	TransactionType string // ETH will mean eth to x / X will mean x to eth / USDC will mean usdc to x / X-X will mean x to x
}

type fullyTransactionInput struct {
	deadline int
	data     [][]byte
}

type UserInfo struct {
	PrivateKey  *ecdsa.PrivateKey
	FromAddress common.Address
}

func WeiToEtherFloatByDecimals(decimals int, wei *big.Int) *big.Float {
	if decimals == 0 {
		decimals = 18
	}
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(decimals)))
}

func intToHex(i int) string {
	return fmt.Sprintf("%x", i)
}

func HexToInt(hex string) *big.Int {
	i := new(big.Int)
	i.SetString(hex, 16)
	return i
}

func toHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func PowInt(x, y int) float64 {
	return float64(math.Pow(float64(x), float64(y)))
}

func GetUserInfo() UserInfo {

	PrivateKey, err := crypto.HexToECDSA(viper.GetString("PrivateKey"))
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA, ok := PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return UserInfo{
		PrivateKey:  PrivateKey,
		FromAddress: FromAddress,
	}
}

func Pad(input string) string {
	var output string = input
	for len(output) < 64 {
		output = "0" + output
	}
	return output
}

func GetSwapRouterInput(futurinput TransactionInput) string {
	var input string
	input = input + "0x472b43f3"

	if futurinput.Path[1].String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		input = input + Pad(futurinput.AmountIn)
		input = input + Pad(futurinput.MinAmountOut)
		input = input + Pad(big.NewInt(int64(80)).String())
		input = input + Pad(big.NewInt(int64(2)).String())
		input = input + Pad(big.NewInt(int64(len(futurinput.Path))).String())
		for _, addr := range futurinput.Path {
			input = input + Pad(strings.Split(addr.String(), "x")[1])
		}
	} else {
		input = input + Pad(futurinput.AmountIn)
		input = input + Pad(futurinput.MinAmountOut)
		input = input + Pad(big.NewInt(int64(80)).String())
		input = input + Pad(strings.Split(futurinput.To.String(), "x")[1])
		input = input + Pad(big.NewInt(int64(len(futurinput.Path))).String())
		for _, addr := range futurinput.Path {
			input = input + Pad(strings.Split(addr.String(), "x")[1])
		}
	}
	return input
}

func PrepareDataToSwap(input TransactionInput) [][]byte {
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

func GetClient() *ethclient.Client {
	var rpcAddress string
	wifi := wifi_managment.GetCurrentWifiSSID1()

	if strings.Contains(wifi, "Livebox-00B0") {
		rpcAddress = "___"
	} else if strings.Contains(wifi, "Guest_2.4") {
		rpcAddress = "___"
	} else {
		rpcAddress = "___"
	}

	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetFlash() *ethclient.Client {
	var rpcAddress string
	wifi := wifi_managment.GetCurrentWifiSSID1()

	if strings.Contains(wifi, "Livebox-00B0") {
		rpcAddress = "___"
	} else if strings.Contains(wifi, "Guest_2.4") {
		rpcAddress = "___"
	} else {
		rpcAddress = "___"
	}

	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SniperUniCake(amountIn float64, amountOutMin float64, token0Addr string, token1Addr string) {
	// calculte time execution
	t1 := time.Now()
	router, err := UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetFlash())
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Println("amount input : ", amountInF)
	fmt.Println("amount min output : ", amountOutF)
	fmt.Println("from address : ", GetUserInfo().FromAddress.String())
	fmt.Println("token0 : ", token0Addr)
	fmt.Println("token1 : ", token1Addr)

	fmt.Println("Preparing data to swap")
	// make a function with this inputs : amountInF / amountOutF / FromAddress / [token0Addr, token1Addr]
	var data [][]byte = PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOutF), GetUserInfo().FromAddress, []common.Address{common.HexToAddress(token0Addr), common.HexToAddress(token1Addr)}, "WETH"})
	fmt.Println("data prepared")

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
	fmt.Println("elapsed time:", elapsed)
	//instance of router will be with the flashbot RPC
	tx, err := router.Multicall0(auth, big.NewInt(time.Now().Add(5*time.Minute).Unix()), data)
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)

	if token1Addr != eth_const.FindAddressFromTokenSymbol("WETH") || token1Addr != eth_const.FindAddressFromTokenSymbol("DAI") {
		ApproveSpenderToken(token1Addr, tx.To().Hex())
		TrackGains(amountIn, []common.Address{common.HexToAddress(token1Addr), common.HexToAddress(token0Addr)})

	}

}

func ApproveSpenderToken(tokenToApprove string, spender string) {

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

	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, id)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 500000

	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(130)), big.NewInt(100))

	tx, err := token.Approve(auth, common.HexToAddress(spender), amount)
	if err != nil {
		log.Fatal(err)
	}

	AwaitToConfirmTx(tx)

}

// //function that check the expected amount out if you sell some amount of token0 for token1
func CheckExpectedAmountOut(path []common.Address) *big.Float {

	router, err := UniswapV2Router.NewUniswapV2Router(common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token, err := erc20.NewErc20(path[0], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, GetUserInfo().FromAddress)
	if err != nil {
		log.Fatal(err)
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("balance is 0")
		os.Exit(1)
	}

	amountOut, err := router.GetAmountsOut(nil, balance, path)
	if err != nil {
		log.Fatal(err)
	}

	tokenOut, err := erc20.NewErc20(path[1], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	decimal0, err := tokenOut.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return (WeiToEtherFloatByDecimals(int(decimal0), amountOut[1]))

}

// function that track the gains between the amount of token In and the expected return in case of sell
func TrackGains(amountIn float64, path []common.Address) {
	fmt.Println("tracking gains")
	go readInput(amountIn, path)
	for ok := true; ok; ok = true {
		if amountIn > 0 {
			expectedAmountOut := CheckExpectedAmountOut(path)
			gainPercentage := new(big.Float).Mul(new(big.Float).Sub(new(big.Float).Quo(expectedAmountOut, big.NewFloat(amountIn)), big.NewFloat(1)), big.NewFloat(100)).String()
			fmt.Println(amountIn, "========>", expectedAmountOut.String(), " ======  gains ", gainPercentage, "%")
			time.Sleep(time.Second * 5)
		} else {
			expectedAmountOut := CheckExpectedAmountOut(path)
			fmt.Println(amountIn, "========>", expectedAmountOut.String(), " ======  gains ")
			time.Sleep(time.Second * 5)
		}
	}
}

func readInput(amountInted float64, path []common.Address) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("could not process input %v\n", input)
		}
		var amountIn float64
		if strings.Contains(string(input), "buy") {
			amountIn, _ = strconv.ParseFloat(strings.Split(string(input), " ")[1], 64)
			Buy(amountIn, []common.Address{path[1], path[0]})
			TrackGains(amountIn+amountInted, path)
		}
		switch string(input) {
		case "sell":
			fmt.Println("selling your amount of tokens")
			SellTokens(path)
		case "half":
			fmt.Println("half your amount of tokens")
			SellHalfTokens(path)
			TrackGains(0, path)

		}
	}
}

func SellTokens(path []common.Address) {
	router, err := UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token, err := erc20.NewErc20(path[0], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, GetUserInfo().FromAddress)
	if err != nil {
		log.Fatal(err)
	}
	amountOut := big.NewInt(0)
	var data [][]byte = PrepareDataToSwap(TransactionInput{toHexInt(balance), toHexInt(amountOut), GetUserInfo().FromAddress, path, "WETH"})

	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 500000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(130)), big.NewInt(100))

	tx, err := router.Multicall0(auth, big.NewInt(time.Now().Add(5*time.Minute).Unix()), data)
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)
}

func CheckAllowanceAndApprove(tokenAddr string, amount *big.Int, spender string) {

	token, err := erc20.NewErc20(common.HexToAddress(tokenAddr), GetClient())
	if err != nil {
		log.Fatal(err)
	}

	allowance, err := token.Allowance(nil, GetUserInfo().FromAddress, common.HexToAddress(spender))
	if err != nil {
		log.Fatal(err)
	}
	//check if allowance token is higher than amountIn
	if allowance.Cmp(amount) < 0 {
		fmt.Println("Your allowance is not high enough, you need to approve more tokens")
		ApproveSpenderToken(tokenAddr, spender)

	}
}

func SellHalfTokens(path []common.Address) {
	router, err := UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token, err := erc20.NewErc20(path[0], GetClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := token.BalanceOf(nil, GetUserInfo().FromAddress)
	if err != nil {
		log.Fatal(err)
	}
	halfBalance := new(big.Int).Div(balance, big.NewInt(2))
	amountOut := big.NewInt(0)
	var data [][]byte = PrepareDataToSwap(TransactionInput{toHexInt(halfBalance), toHexInt(amountOut), GetUserInfo().FromAddress, path, "WETH"})

	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = 500000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(130)), big.NewInt(100))

	tx, err := router.Multicall0(auth, big.NewInt(time.Now().Add(5*time.Minute).Unix()), data)
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)
}

func Buy(amountIn float64, path []common.Address) {
	router, err := UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token, err := erc20.NewErc20(path[0], GetClient())
	if err != nil {
		log.Fatal(err)
	}

	decimal0, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	amountInF, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)

	amountOut := big.NewInt(0)
	var data [][]byte = PrepareDataToSwap(TransactionInput{toHexInt(amountInF), toHexInt(amountOut), GetUserInfo().FromAddress, path, "WETH"})
	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	test, _ := big.NewFloat(amountIn * PowInt(10, int(decimal0))).Int(nil)
	if path[0].String() == eth_const.FindAddressFromTokenSymbol("WETH") {
		auth.Value = test
	}

	auth.GasLimit = 500000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())

	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))
	tx, err := router.Multicall0(auth, big.NewInt(time.Now().Add(5*time.Minute).Unix()), data)
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)
}

func BuyTokenOnSC() {
	client := GetClient()

	//get the contract address
	contractAddress := common.HexToAddress("0x564ea1fcf1Ba9Ebd151ec3f34208CFaF1062F524")

	//get the contract instance
	contract, err := SURGE.NewSURGE(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance of the contract is: ", balance)
	// check if balance is more than 4 eth
	// if balance.Cmp(big.NewInt(4*1e18)) < 0 {
	// 	fmt.Println("Balance is not enough to buy tokens")
	// 	return
	// } else {

	// buyAmount is 0.1 eth
	buyAmount := big.NewInt(1e17)
	minTokenOut := big.NewInt(0)
	deadline := big.NewInt(time.Now().Add(5 * time.Minute).Unix())

	tokenCaller, err := SURGE.NewSURGECaller(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	expectedAmountOut, err := tokenCaller.GetTokenAmountOut(nil, buyAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Expected amount out is: ", expectedAmountOut)

	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().PrivateKey, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	auth.Value = buyAmount
	auth.GasLimit = 3000000
	auth.GasPrice, err = GetClient().SuggestGasPrice(context.Background())
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(130)), big.NewInt(100))
	tx, err := contract.Buy(auth, minTokenOut, deadline)
	if err != nil {
		log.Fatal(err)
	}
	AwaitToConfirmTx(tx)

	// }

}
