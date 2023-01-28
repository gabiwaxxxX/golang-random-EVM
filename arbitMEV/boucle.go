package arbitMEV

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gabrielfournier1/bot/contract/Mugen"
	"github.com/gabrielfournier1/bot/contract/Treasury"
	"github.com/gabrielfournier1/bot/contract/UniswapV3Router"
	"github.com/spf13/viper"
)

type DexScreenerResp struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	PriceUsd string `json:"priceUsd"`
}
type userInfo struct {
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

func GetUserInfo() userInfo {

	privateKey, err := crypto.HexToECDSA(viper.GetString("privateKey"))
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return userInfo{
		privateKey:  privateKey,
		fromAddress: fromAddress,
	}
}
func GetArbiClient() *ethclient.Client {
	var rpcAddress string = "https://rpc.ankr.com/arbitrum"
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetBondPrice() *big.Float {
	treasuryAddr := common.HexToAddress("0xf7bE8476AE27d27eBc236e33020150B23a86F3Dd")
	TreasuryContract, err := Treasury.NewTreasury(treasuryAddr, GetArbiClient())
	if err != nil {
		log.Fatal(err)
	}
	tokenBondPrice, err := TreasuryContract.PricePerToken(nil)
	if err != nil {
		log.Fatal(err)
	}
	//price in USD
	priceInUSD := new(big.Float).Quo(new(big.Float).SetInt(tokenBondPrice), big.NewFloat(100))
	// fmt.Println("Price of bond atm", priceInUSD)
	return priceInUSD
}

func GetPriceFromDexScreen() (*big.Float, *big.Float) {
	base, err := url.Parse("https://api.dexscreener.com/latest/dex/tokens/0xFc77b86F3ADe71793E1EEc1E7944DB074922856e")
	if err != nil {
		return nil, nil
	}

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("Oh no! Got: ")

		body, err := io.ReadAll(resp.Body)
		log.Println(string(body))
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var dexScreenerResp DexScreenerResp
	err = json.Unmarshal(body, &dexScreenerResp)
	if err != nil {
		log.Fatal(err)

	}

	price1, _ := new(big.Float).SetString(dexScreenerResp.Pairs[0].PriceUsd)
	price2, _ := new(big.Float).SetString(dexScreenerResp.Pairs[1].PriceUsd)

	return price1, price2

}

func CheckPossibleArbitrage() {
	//get price of bond
	//get price of token on dex
	//compare
	//if price of token on dex is lower than price of bond
	//buy bond
	//sell on dex
	//profit
	for {
		bondPrice := GetBondPrice()
		price1, price2 := GetPriceFromDexScreen()

		if bondPrice.Cmp(price1) == -1 {
			fmt.Println("Arbitrage possible on pair 1")
			arbitrage := new(big.Float).Quo(new(big.Float).Sub(price1, bondPrice), price1)
			arbitrage.Mul(arbitrage, big.NewFloat(100))
			fmt.Println("Arbitrage %", arbitrage)
			if arbitrage.Cmp(big.NewFloat(1)) == 1 {
				BuyBondMugen()
				return
			}

		}
		if bondPrice.Cmp(price2) == -1 {
			fmt.Println("Arbitrage possible on pair 2")
			arbitrage := new(big.Float).Quo(new(big.Float).Sub(price2, bondPrice), price2)
			arbitrage.Mul(arbitrage, big.NewFloat(100))
			fmt.Println("Arbitrage %", arbitrage)
			if arbitrage.Cmp(big.NewFloat(1)) == 1 {
				BuyBondMugen()
				return
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func BuyBondMugen() {
	treasuryAddr := common.HexToAddress("0xf7bE8476AE27d27eBc236e33020150B23a86F3Dd")
	TreasuryContract, err := Treasury.NewTreasury(treasuryAddr, GetArbiClient())
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(100000000) //100 USDC
	USDC := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")

	auth, err := bind.NewKeyedTransactorWithChainID(GetUserInfo().privateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasPrice, err = GetArbiClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = new(big.Int).Div(new(big.Int).Mul(auth.GasPrice, big.NewInt(200)), big.NewInt(100))
	auth.GasLimit = 1000000

	tx, err := TreasuryContract.Deposit(auth, USDC, amount)
	if err != nil {
		log.Fatal(err)
	}
	awaitToConfirmTx(tx)
}

func awaitToConfirmTx(tx *types.Transaction) {

	bind.WaitMined(context.Background(), GetArbiClient(), tx)
	receipt, err := GetArbiClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal("not founded", err)
	}

	if receipt.Status == 1 {
		log.Printf("Transaction has been mined")

	} else {
		fmt.Println("Transaction has been rejected")
	}
}

func SellTokenEthPool() {

	uniV3Router, err := UniswapV3Router.NewUniswapV3Router(common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"), GetArbiClient())
	if err != nil {
		log.Fatal(err)
	}
	params := UniswapV3Router.IV3SwapRouterExactInputSingleParams{}
	params.TokenIn = common.HexToAddress("0xFc77b86F3ADe71793E1EEc1E7944DB074922856e")
	params.TokenOut = common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
	params.Fee = big.NewInt(1000)
	params.Recipient = common.HexToAddress("0x0000000000000000000000000000000000000002")

	MugenContract, err := Mugen.NewMugen(common.HexToAddress("0xFc77b86F3ADe71793E1EEc1E7944DB074922856e"), GetArbiClient())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := MugenContract.BalanceOf(nil, GetUserInfo().fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance of token", balance)
	params.AmountIn = balance
	params.AmountOutMinimum = big.NewInt(0)
	params.SqrtPriceLimitX96 = big.NewInt(0)

	tx, err := uniV3Router.ExactInputSingle(nil, params)
	if err != nil {
		fmt.Println("Error", err)
		log.Fatal(err)
	}
	fmt.Println(tx.Data())

}

// 0x04e45aaf
// 000000000000000000000000fc77b86f3ade71793e1eec1e7944db074922856e
// 00000000000000000000000082af49447d8a07e3bd95bd0d56f35241523fbab1
// 0000000000000000000000000000000000000000000000000000000000002710
// 0000000000000000000000000000000000000000000000000000000000000002
// 00000000000000000000000000000000000000000000000029a2241af62c0000
// 0000000000000000000000000000000000000000000000000372538f502d508f
// 0000000000000000000000000000000000000000000000000000000000000000

// 0x49404b7c
// 0000000000000000000000000000000000000000000000000372538f502d508f
// 0000000000000000000000006e20692aab07c684ee72edc598878c87cc92217d
