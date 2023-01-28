package main

import (

	"log"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/spf13/viper"
	"github.com/ethereum/go-ethereum/common"
)

// const UNI_SWAP_FACTORY_CONTRACT_ADDRESS="0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const SUSHI_FACTORY_ADDRESS="0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"



type PairsInfos struct {
    PairsInfos []PairInfo `json:"pairsInfo"`
}

type PairInfo struct {
   Pair 		Pair 	`json:"pair"`
//    APY 	string 	`json:"apy"`
   Liquidity string 	`json:"liquidity"`
}

type Pair struct {
	Token0			Token 	`json:"token0"`
	Token1			Token 	`json:"token1"`
	Id				string 	`json:"id"`
}

type Token struct {
	Id				string 	`json:"id"`
}


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
	// sync pairs
	fmt.Println("loading pairs...")
	uniswapAddr :=common.HexToAddress(SUSHI_FACTORY_ADDRESS)
	var LPLoaded map[common.Address]map[common.Address]map[common.Address]common.Address
	LPLoaded = make(map[common.Address]map[common.Address]map[common.Address]common.Address)

	folders, err := ioutil.ReadDir("./SushiSwapData")
	if err != nil {
		log.Fatal(err)
	}

	for _, pairFile := range folders {
		pairFileName := pairFile.Name()
		jsonFile, err := os.Open("./SushiSwapData/" + pairFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var pairsInfo PairsInfos
		json.Unmarshal([]byte(byteValue), &pairsInfo)

		fmt.Println(pairsInfo)
		for _, PairInfo := range pairsInfo.PairsInfos {
			IntLiquidity, _ :=  strconv.ParseFloat(PairInfo.Liquidity, 32)

			if IntLiquidity < 10000 { continue } 

			token0 := common.HexToAddress(PairInfo.Pair.Token0.Id)
			token1 := common.HexToAddress(PairInfo.Pair.Token1.Id)
			if _, ok := LPLoaded[uniswapAddr]; !ok {
				LPLoaded[uniswapAddr] = make(map[common.Address]map[common.Address]common.Address)
			}
			if _, ok := LPLoaded[uniswapAddr][token0]; !ok {
				LPLoaded[uniswapAddr][token0] = make(map[common.Address]common.Address)	
			}
			LPLoaded[uniswapAddr][token0][token1] = common.HexToAddress(PairInfo.Pair.Id)		
			
		}

	}

	jsonStr, err := json.Marshal(LPLoaded)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("LPLoadedShushi.json", jsonStr, 0644)
	}
}





