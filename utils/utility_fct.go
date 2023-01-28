package utils

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/gabrielfournier1/bot/contract/UniswapV2Pair"
	"github.com/gabrielfournier1/bot/contract/erc20"
	"github.com/gabrielfournier1/bot/rpcClient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/holiman/uint256"
)

// convert WEI to ETH
func FormatEthWeiToEther(etherAmount *big.Int) float64 {
	var base, exponent = big.NewInt(10), big.NewInt(18)
	denominator := base.Exp(base, exponent, nil)
	tokensSentFloat := new(big.Float).SetInt(etherAmount)
	denominatorFloat := new(big.Float).SetInt(denominator)
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	return final
}

// convert ETH to WEI
func FormatEthToWei(etherAmount float64) *big.Int {
	var base, exponent = big.NewInt(10), big.NewInt(18)
	denominator := base.Exp(base, exponent, nil)
	tokensSentFloat := new(big.Float).SetFloat64(etherAmount)
	denominatorFloat := new(big.Float).SetInt(denominator)
	tokensSent, _ := new(big.Float).Mul(tokensSentFloat, denominatorFloat).Int(nil)
	return tokensSent
}

// fetch ERC20 token symbol
func GetTokenSymbol(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	sym, _ := tokenIntance.Symbol(nil)
	return sym
}

func DownloadBlock(blockNumber string) ([]interface{}, string) {

	result, err := rpcClient.HTTPClient.Call("eth_getBlockByNumber", blockNumber, true)
	if err != nil {
		fmt.Println(err)
	}
	returnedBlock := result.Result.(map[string]interface{})["number"].(string)

	txArray := result.Result.(map[string]interface{})["transactions"].([]interface{})

	//return the array of transactions and the block number
	return txArray, returnedBlock
}

func RemoveLeadingZeros(inputString string) string {
	for index := 0; index < len(inputString); index++ {
		// Just return the string as soon as first non zero value is detected
		if string(inputString[index]) != "0" {
			return inputString[index:]
		}
	}
	return "" // Value is only zeros
}

func DeriveReservesFromSlot(slot string) (uint256.Int, uint256.Int) {

	// Yes these names are right, for some reason they are stored in reverse order
	reserve1, _ := uint256.FromHex("0x" + RemoveLeadingZeros(slot[10:38]))

	reserve0, _ := uint256.FromHex("0x" + RemoveLeadingZeros(slot[38:66]))

	return *reserve0, *reserve1
}

func GetReserves(pair common.Address) (*uint256.Int, *uint256.Int) {

	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	token0Address, err := pairToken.Token0(nil)
	if err != nil {
		log.Fatal(err)
	}
	token1Address, err := pairToken.Token1(nil)
	if err != nil {
		log.Fatal(err)
	}
	return GetReservesFromToken(token0Address, pair), GetReservesFromToken(token1Address, pair)
}

func GetReservesFromToken(tokenAddress common.Address, pair common.Address) *uint256.Int {
	token, err := erc20.NewErc20(tokenAddress, GetClient())
	if err != nil {
		if fmt.Sprint(err) == "Post \"https:/_____\": dial tcp 45.55.121.48:443: connect: connection refused" {
			log.Fatal(err)
		}
		fmt.Println("you got an error", err, " => ", tokenAddress.String())
		return uint256.NewInt(0)
	}

	balance, err := token.BalanceOf(nil, pair)
	if err != nil {
		if fmt.Sprint(err) == "Post \"https:/_____\": dial tcp 45.55.121.48:443: connect: connection refused" {

			log.Fatal(err)
		}
		fmt.Println("you got an error", err, " => ", tokenAddress.String())
		return uint256.NewInt(0)
	}
	setUint, o := uint256.FromBig(balance)
	if o {
		log.Fatal("error")
	}
	return setUint
}

func GetReservesFromTokenBig(tokenAddress common.Address, pair common.Address) *big.Int {
	token, err := erc20.NewErc20(tokenAddress, GetClient())
	if err != nil {
		if fmt.Sprint(err) == "Post \"https:/_____\": dial tcp 45.55.121.48:443: connect: connection refused" {
			log.Fatal(err)
		}
		fmt.Println("you got an error", err, " => ", tokenAddress.String())
		return big.NewInt(0)
	}

	balance, err := token.BalanceOf(nil, pair)
	if err != nil {
		if fmt.Sprint(err) == "Post \"https:/_____\": dial tcp 45.55.121.48:443: connect: connection refused" {

			log.Fatal(err)
		}
		fmt.Println("you got an error", err, " => ", tokenAddress.String())
		return big.NewInt(0)

	}
	// setUint, o := uint256.FromBig(balance)
	// if o {
	// 	log.Fatal("error")
	// }
	return balance
}

func CheckRugOnLp(pair common.Address) bool {
	token0, token1 := GetReserves(pair)
	//check if one of them is less than 100
	if token0.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 || token1.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 {
		return true
	}

	return false
}

func CheckRugOnLpWithTokens(pair common.Address, client *ethclient.Client) bool {
	// token0, token1 :=GetReservesFromToken(token0Addr,pair ), GetReservesFromToken(token1Addr,pair)
	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, client)
	if err != nil {
		log.Fatal(err)
	}
	Reserves, err := pairToken.GetReserves(nil)
	if err != nil {
		log.Fatal("error ", err)
	}
	reserve0, o := uint256.FromBig(Reserves.Reserve0)
	if o {
		log.Fatal("error")
	}
	reserve1, o := uint256.FromBig(Reserves.Reserve1)
	if o {
		log.Fatal("error")
	}
	if reserve0.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 || reserve1.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 {
		return true
	}
	return false
}

func CheckLastUsOfLpTokens(pair common.Address, client *ethclient.Client) bool {
	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := pairToken.GetReserves(nil)

	timel, _ := strconv.ParseInt(strconv.Itoa(int(tx.BlockTimestampLast)), 10, 64)
	delta := time.Now().Unix() - timel
	if delta > 604800 {
		return true
	}
	return false
}

func CheckAll(pair common.Address, client *ethclient.Client) bool {
	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, client)
	if err != nil {
		fmt.Println("Got an erro on line 168 ")
		log.Fatal(err)
	}
	Reserves, err := pairToken.GetReserves(nil)
	if err != nil {
		return true

	}
	reserve0, o := uint256.FromBig(Reserves.Reserve0)
	if o {
		log.Fatal("error")
	}
	reserve1, o := uint256.FromBig(Reserves.Reserve1)
	if o {
		log.Fatal("error")
	}
	if reserve0.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 || reserve1.Cmp(uint256.NewInt(0).SetUint64(1000)) == -1 {
		return true
	}
	timel, _ := strconv.ParseInt(strconv.Itoa(int(Reserves.BlockTimestampLast)), 10, 64)
	delta := time.Now().Unix() - timel
	if delta > 604800 {
		return true
	}
	return false
}

func GetReservesSpeed(pair common.Address) (*big.Int, *big.Int) {

	pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, GetClient())
	if err != nil {
		log.Fatal(err)
	}
	Reserves, err := pairToken.GetReserves(nil)
	if err != nil {
		log.Fatal("error ", err)
	}
	return Reserves.Reserve0, Reserves.Reserve1
}
