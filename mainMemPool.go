package main

import (
	"log"
	"os"

	"github.com/gabrielfournier1/bot/checkHoneyPot"
	"github.com/gabrielfournier1/bot/utils"
	"github.com/spf13/viper"
	// "github.com/gabrielfournier1/bot/mempool"
	// "github.com/gabrielfournier1/bot/utils"
	// "github.com/ethereum/go-ethereum/common"
)

func main() {

	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		if os.IsNotExist(err) {
			if args := os.Args; len(args) >= 2 && args[1] == "init" {
				return
			}
		}
		log.Fatal("init config before sniper")
	}

	// checkHoneyPot.SimulateHoneyPotAndGetInternalTaxes()
	// uniswapV3Decode.FindPool(common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))
	// safeIn.FindDevWallet(common.HexToAddress("0x5d91fCcbE208fc31C204cEF807f654155562ba09"))
	// TOKEN_TO_BUY := os.Args[1]
	// t1 := time.Now()
	// safeIn.FindDevWallet(common.HexToAddress(TOKEN_TO_BUY))
	// t2 := time.Now()
	// log.Println("time: ", t2.Sub(t1))
	// mempool.ListenOnMemPoolUser(ownerAddr)

	// walletBalance.GetAllTransferEventsToAnAddress("Your wallet adress")
	// arbitrum.SwapExactTokenForToken()
	// payload_parsing.ParsePayload("0x8b49e081000000000000000000000000000000000000000000000000000000000000001cdaea27499e15c6d780708a4ee4832b0e0fe576431d033448a8da70fcb1120b6f4ebd72faba125ce4ca3cdf6ce77e1d8d65459121d344bb63b6a078e01eb9d86e00")
	// payload_parsing.ParsePayload("0x8b49e081000000000000000000000000000000000000000000000000000000000000001c611e3852278b9efb5784d34b9de52dda20c6a0dab033cd435b07fe5963b543a03842279d497d7900090102962895cd14b8b9707a919051ec5e12e797a6c858d401")
	// tata, toto := payload_parsing.ParsePayload("0x9278a35a00000000000000000000000098dcfcf075a06e12ba34a4b849aaaad031e179ea0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a4ed8df1640000000000000000000000000000000000000000000000000000000000000040527a3e852c67b61aac552b4d372026f34b48c9afe366e76b6c6fa9d1ebb871e3000000000000000000000000000000000000000000000000000000000000002f466f72207768617420706572736f6e20646f20616c6c206d656e2074616b65206f666620746865697220686174733f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	// fmt.Println(tata, toto)

	// bc := "bsc"
	// targetToken := "0x5693eBF94b7c104D8cf2c1b1cC6DEB53f7ce48Bf"
	// buyToken := "0x55d398326f99059fF775485246999027B3197955"
	// amount := big.NewInt(10000)

	checkHoneyPot.FindYourLP(os.Args[1])

	// bc2 := "eth"
	// targetToken2 := "0x7EC7cFc14d04308540c0f7770f263Fec51D7342d"
	// buyToken := "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"

	// checkHoneyPot.SimulateHoneyPotAndGetInternalTaxes(bc2, targetToken2, buyToken, amount)

	// arbitrum.MintRandomNft()
	// arbitrum.CheckPause()
	// arbitrum.BuyRandFromEth()
	utils.BuyTokenOnSC()

}

// mempool.ListenOnMemPoolTargetContract([]string{"0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"})
