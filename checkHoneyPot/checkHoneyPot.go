package checkHoneyPot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum"
)

// BSC : 0xf12903b79ecafdaf79827bc61f6ad7c927ea1708
// pancake : 0x10ED43C718714eb63d5aA57B78B54704E256024E
// wbnb : 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c

// eth : 0x
// uniswap : 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
type FunctionOutputsInt struct {
	BuyResult    *big.Int   "json:\"buyResult\""
	TokenBalance *big.Int   "json:\"tokenBalance\""
	SellResult   *big.Int   "json:\"sellResult\""
	BuyCost      *big.Int   "json:\"buyCost\""
	SellCost     *big.Int   "json:\"sellCost\""
	Amounts      []*big.Int "json:\"amounts\""
}

func GetClient(bc string) *ethclient.Client {
	var rpcAddress string

	switch bc {
	case "bsc":
		rpcAddress = "https://bsc-dataseed3.defibit.io"
	case "eth":
		rpcAddress = "https:/_____"
	case "arbitrum":
		rpcAddress = "https://rpc.ankr.com/arbitrum"
	case "matic":
		rpcAddress = "https://rpc-mainnet.matic.quiknode.pro"
	case "ftm":
		rpcAddress = "https://rpc.ftm.tools"
	case "avax":
		rpcAddress = "https://rpc.ankr.com/avalanche"
	default:
		rpcAddress = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	}
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SimulateHoneyPotAndGetInternalTaxes(bc string, _targetToken string, _buyToken string, amount *big.Int) (FunctionOutputsInt, error) {
	time1 := time.Now()
	ctx := context.Background()
	var targetContractAddress common.Address
	var routerContractAddress common.Address
	var wrappedCoinAddress common.Address
	var honeyPotCheckerFile []byte
	var err error
	switch bc {
	case "bsc":
		targetContractAddress = common.HexToAddress("0x1840cCcca3f798edCaA77cb8E3607D3CCA9EDE3E")
		routerContractAddress = common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
		wrappedCoinAddress = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
		honeyPotCheckerFile, err = ioutil.ReadFile("./contract/Vyper_contract/Vyper_contract.json")
		if err != nil {
			fmt.Println("Error while reading HoneyPotCheck.json", "err", err)
			return FunctionOutputsInt{}, err
		}
	case "eth":
		targetContractAddress = common.HexToAddress("0xefd7e19cc770ddb8b5ec970e7dcf354c3e2f51c6")
		routerContractAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
		wrappedCoinAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
		honeyPotCheckerFile, err = ioutil.ReadFile("./contract/Vyper_contract/Vyper_contract.json")
		if err != nil {
			fmt.Println("Error while reading HoneyPotCheck.json", "err", err)
			return FunctionOutputsInt{}, err
		}
	default:
		targetContractAddress = common.HexToAddress("0xefd7e19cc770ddb8b5ec970e7dcf354c3e2f51c6")
		routerContractAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
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
