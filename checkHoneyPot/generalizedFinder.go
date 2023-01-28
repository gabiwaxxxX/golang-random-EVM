package checkHoneyPot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Pair"
)

func FindYourLP(lpAddress string) {
	var wg sync.WaitGroup
	// wg.Add(1)
	// t1 := time.Now()

	for _, blockchain := range BLOCKCHAIN_LIST {
		// wg := wg
		getCodeAt(&wg, blockchain, lpAddress)
	}

	// wg.Wait()
}

func getCodeAt(wg *sync.WaitGroup, blockchain, lpAddress string) {

	client := GetClient(blockchain)

	// test := curlPost(blockchain, lpAddress)
	// fmt.Println(test)

	CodeAt, err := client.CodeAt(context.Background(), common.HexToAddress(lpAddress), nil)

	if err != nil {
		fmt.Println("err: ", err)
		fmt.Println(err)
		// defer wg.Done()
		FindYourLP(lpAddress)
	}
	if len(CodeAt) > 0 {
		fmt.Println("Found LP on blockchain", blockchain)
		QuerryFactoryOf(blockchain, lpAddress)
		// defer wg.Done()

	}

}

func QuerryFactoryOf(blockchain, lpAddress string) {

	client := GetClient(blockchain)
	log.Println("QuerryFactoryOf", blockchain, lpAddress)

	lpContract, err := UniswapV2Pair.NewUniswapV2PairCaller(common.HexToAddress(lpAddress), client)
	if err != nil {
		fmt.Println(err)
	}

	factoryAddress, err := lpContract.Factory(nil)
	if err != nil {
		fmt.Println("err 1", err)

	}

	token0Address, err := lpContract.Token0(nil)
	if err != nil {
		fmt.Println("err 2", err)

	}

	token1Address, err := lpContract.Token1(nil)
	if err != nil {
		fmt.Println("err 3", err)

	}

	fmt.Println("Factory address: ", factoryAddress.Hex())
	fmt.Println("Token 0 address: ", token0Address.Hex())
	fmt.Println("Token 1 address: ", token1Address.Hex())
	fmt.Println("Router address: ", FACTORY_TO_ROUTER[blockchain][factoryAddress.Hex()])

	if contains(WHITELISTED_ASSET[blockchain], token0Address.Hex()) {
		amount := big.NewInt(10000000000000000)
		SimulateHoneyPotAndGetInternalTaxes3(blockchain, token1Address.String(), token0Address.Hex(), amount, FACTORY_TO_ROUTER[blockchain][factoryAddress.Hex()])
	} else if contains(WHITELISTED_ASSET[blockchain], token1Address.Hex()) {
		amount := big.NewInt(10000000000000000)
		SimulateHoneyPotAndGetInternalTaxes3(blockchain, token0Address.String(), token1Address.Hex(), amount, FACTORY_TO_ROUTER[blockchain][factoryAddress.Hex()])
	}

}

func SimulateHoneyPotAndGetInternalTaxes2(bc string, _targetToken string, _buyToken string, amount *big.Int, routerContract string) (FunctionOutputsInt, error) {
	time1 := time.Now()
	ctx := context.Background()
	routerContractAddress := common.HexToAddress(routerContract)
	var targetContractAddress common.Address
	var wrappedCoinAddress common.Address
	honeyPotCheckerFile, err := ioutil.ReadFile("./contract/Vyper_contract/Vyper_contract.json")
	if err != nil {
		fmt.Println("Error while reading HoneyPotCheck.json", "err", err)
		return FunctionOutputsInt{}, err
	}
	switch bc {
	case "bsc":
		targetContractAddress = common.HexToAddress("0x5834cC0D87Ac7FeCE6eB28c4a881A7EFa08c5A22")
		// targetContractAddress = common.HexToAddress("0x1840cCcca3f798edCaA77cb8E3607D3CCA9EDE3E")
		wrappedCoinAddress = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
		fmt.Println("BSC")

	case "eth":
		targetContractAddress = common.HexToAddress("0x5a9fa3ab3f0747830a8b1b8813514bf5ca5dd97d")
		wrappedCoinAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	case "matic":
		targetContractAddress = common.HexToAddress("0x11786b639431778FaC0D5eBF4Db727C77519d8f3")
		wrappedCoinAddress = common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270")
	case "arbitrum":
		targetContractAddress = common.HexToAddress("0xB57889B6868892B7e1Db43a20Cd82d4f7e578BD1")
		wrappedCoinAddress = common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1")
	case "avax":
		targetContractAddress = common.HexToAddress("0xfb4147dE76aa455f1b34b9919dFf7889A765AeD5")
		wrappedCoinAddress = common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
	default:
		targetContractAddress = common.HexToAddress("0xefd7e19cc770ddb8b5ec970e7dcf354c3e2f51c6")
		wrappedCoinAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	}
	targetToken := common.HexToAddress(_targetToken)
	buyToken := common.HexToAddress(_buyToken)

	honeyPotCheckerABI, err := abi.JSON(strings.NewReader(string(honeyPotCheckerFile)))
	if err != nil {
		fmt.Println("Error while reading abi", "err", err)
		return FunctionOutputsInt{}, err
	}

	method := honeyPotCheckerABI.Methods["honeyCheck"]
	input, err := honeyPotCheckerABI.Pack(method.Name, buyToken, targetToken, routerContractAddress, amount)
	if err != nil {
		fmt.Println("Error while packing input", "err", err)
		return FunctionOutputsInt{}, err
	}

	if err != nil {
		fmt.Println("Error while packing input", "err", err)
		return FunctionOutputsInt{}, err
	}

	userAddress := common.HexToAddress("Your wallet adress")
	var msg ethereum.CallMsg
	if buyToken.String() == wrappedCoinAddress.String() {
		msg = ethereum.CallMsg{
			From:  userAddress,
			To:    &targetContractAddress,
			Value: amount,
			Data:  input,
		}
	} else {
		msg = ethereum.CallMsg{
			From: userAddress,
			To:   &targetContractAddress,
			Data: input,
		}

	}

	result, err := GetClient(bc).CallContract(ctx, msg, nil)
	if err != nil {
		fmt.Println("Got an error calling ", err)
		return FunctionOutputsInt{}, err
	}
	// fmt.Println("result", result)
	decoded, err := honeyPotCheckerABI.Unpack(method.Name, result)
	if err != nil {

		log.Fatal("you got an error on unmarshaling the result ", err)
	}

	var params FunctionOutputsInt = decoded[0].(struct {
		BuyResult    *big.Int   "json:\"buyResult\""
		TokenBalance *big.Int   "json:\"tokenBalance\""
		SellResult   *big.Int   "json:\"sellResult\""
		BuyCost      *big.Int   "json:\"buyCost\""
		SellCost     *big.Int   "json:\"sellCost\""
		Amounts      []*big.Int "json:\"amounts\""
	})

	var outputs FunctionOutputsInt
	outputs.BuyResult = params.BuyResult
	outputs.TokenBalance = params.TokenBalance
	outputs.SellResult = params.SellResult
	outputs.BuyCost = params.BuyCost
	outputs.SellCost = params.SellCost
	outputs.Amounts = params.Amounts

	fmt.Println("buyGasCost", outputs.BuyCost, "units")
	fmt.Println("sellGasCost", outputs.SellCost, "units")
	fmt.Println("==========================")

	one := big.NewFloat(1)

	buyTaxDiv := new(big.Float).Quo(new(big.Float).SetInt(outputs.BuyResult), new(big.Float).SetInt(outputs.Amounts[1]))
	buyTaxDiv.Sub(one, buyTaxDiv)
	buyTaxDiv.Mul(buyTaxDiv, big.NewFloat(100))
	fmt.Println("buyTax ! ", buyTaxDiv, "%")

	sellTaxDiv := new(big.Float).Quo(new(big.Float).SetInt(outputs.SellResult), new(big.Float).SetInt(amount))
	sellTaxDiv.Sub(one, sellTaxDiv)
	sellTaxDiv.Mul(sellTaxDiv, big.NewFloat(100))
	fmt.Println("sellTax ! ", sellTaxDiv, "%")
	time2 := time.Now()
	fmt.Println("time", time2.Sub(time1))

	return outputs, nil

}

func SimulateHoneyPotAndGetInternalTaxes3(bc string, _targetToken string, _buyToken string, amount *big.Int, routerContract string) (FunctionOutputsInt, error) {
	time1 := time.Now()
	ctx := context.Background()
	routerContractAddress := common.HexToAddress(routerContract)
	var targetContractAddress common.Address
	honeyPotCheckerFile, err := ioutil.ReadFile("./contract/Vyper_contract/Vyper_contract_1.json")
	if err != nil {
		fmt.Println("Error while reading HoneyPotCheck.json", "err", err)
		return FunctionOutputsInt{}, err
	}
	switch bc {
	case "bsc":
		targetContractAddress = common.HexToAddress("0x5834cC0D87Ac7FeCE6eB28c4a881A7EFa08c5A22")
	case "eth":
		targetContractAddress = common.HexToAddress("0x5a9fa3ab3f0747830a8b1b8813514bf5ca5dd97d")
	case "matic":
		targetContractAddress = common.HexToAddress("0x11786b639431778FaC0D5eBF4Db727C77519d8f3")
	case "arbitrum":
		targetContractAddress = common.HexToAddress("0xB57889B6868892B7e1Db43a20Cd82d4f7e578BD1")
	case "avax":
		targetContractAddress = common.HexToAddress("0xfb4147dE76aa455f1b34b9919dFf7889A765AeD5")
	case "fantom":
		targetContractAddress = common.HexToAddress("0xD7a25A571d3db5639E3d5C8E6428995199A8dAfe")
	default:
		targetContractAddress = common.HexToAddress("0x5a9fa3ab3f0747830a8b1b8813514bf5ca5dd97d")

	}
	targetToken := common.HexToAddress(_targetToken)
	buyToken := common.HexToAddress(_buyToken)

	honeyPotCheckerABI, err := abi.JSON(strings.NewReader(string(honeyPotCheckerFile)))
	if err != nil {
		fmt.Println("Error while reading abi", "err", err)
		return FunctionOutputsInt{}, err
	}

	method := honeyPotCheckerABI.Methods["honeyCheck"]
	input, err := honeyPotCheckerABI.Pack(method.Name, buyToken, targetToken, routerContractAddress)
	if err != nil {
		fmt.Println("Error while packing input", "err", err)
		return FunctionOutputsInt{}, err
	}

	if err != nil {
		fmt.Println("Error while packing input", "err", err)
		return FunctionOutputsInt{}, err
	}

	userAddress := common.HexToAddress("Your wallet adress")

	msg := ethereum.CallMsg{
		From:  userAddress,
		To:    &targetContractAddress,
		Value: amount,
		Data:  input,
	}

	result, err := GetClient(bc).CallContract(ctx, msg, nil)
	if err != nil {
		fmt.Println("Got an error calling ", err)
		return FunctionOutputsInt{}, err
	}
	// fmt.Println("result", result)
	decoded, err := honeyPotCheckerABI.Unpack(method.Name, result)
	if err != nil {

		log.Fatal("you got an error on unmarshaling the result ", err)
	}

	var params FunctionOutputsInt = decoded[0].(struct {
		BuyResult    *big.Int   "json:\"buyResult\""
		TokenBalance *big.Int   "json:\"tokenBalance\""
		SellResult   *big.Int   "json:\"sellResult\""
		BuyCost      *big.Int   "json:\"buyCost\""
		SellCost     *big.Int   "json:\"sellCost\""
		Amounts      []*big.Int "json:\"amounts\""
	})

	var outputs FunctionOutputsInt
	outputs.BuyResult = params.BuyResult
	outputs.TokenBalance = params.TokenBalance
	outputs.SellResult = params.SellResult
	outputs.BuyCost = params.BuyCost
	outputs.SellCost = params.SellCost
	outputs.Amounts = params.Amounts

	fmt.Println("buyGasCost", outputs.BuyCost, "units")
	fmt.Println("sellGasCost", outputs.SellCost, "units")
	fmt.Println("==========================")

	one := big.NewFloat(1)

	buyTaxDiv := new(big.Float).Quo(new(big.Float).SetInt(outputs.BuyResult), new(big.Float).SetInt(outputs.Amounts[1]))
	buyTaxDiv.Sub(one, buyTaxDiv)
	buyTaxDiv.Mul(buyTaxDiv, big.NewFloat(100))
	fmt.Println("buyTax ! ", buyTaxDiv, "%")

	sellTaxDiv := new(big.Float).Quo(new(big.Float).SetInt(outputs.SellResult), new(big.Float).SetInt(amount))
	sellTaxDiv.Sub(one, sellTaxDiv)
	sellTaxDiv.Mul(sellTaxDiv, big.NewFloat(100))
	fmt.Println("sellTax ! ", sellTaxDiv, "%")
	time2 := time.Now()
	fmt.Println("time", time2.Sub(time1))

	return outputs, nil

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
