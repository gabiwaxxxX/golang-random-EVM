package main

import (

	"log"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"


	"github.com/spf13/viper"
	"github.com/ethereum/go-ethereum/common"
)




type Pairs struct {
    Pairs []Pair `json:"pairs"`
}

type Pair struct {
    Id			string 	`json:"id"`
    Token0		Token 	`json:"token0"`
	Token1		Token 	`json:"token1"`
}

type Token struct {
	Id			string 	`json:"id"`
}


func main() {
	fmt.Println("Hello World")
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
	uniswapAddr :=common.HexToAddress(UNI_SWAP_FACTORY_CONTRACT_ADDRESS)
	var LPLoaded map[common.Address]map[common.Address]map[common.Address]common.Address
	LPLoaded = make(map[common.Address]map[common.Address]map[common.Address]common.Address)

	folders, err := ioutil.ReadDir("./data/pairs/pairs")
	if err != nil {
		log.Fatal(err)
	}

	for _, pairFile := range folders {
		pairFileName := pairFile.Name()
		jsonFile, err := os.Open("./data/pairs/pairs/" + pairFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var pairs Pairs
		json.Unmarshal([]byte(byteValue), &pairs)

		// fmt.Println(pairs)
		for _, pair := range pairs.Pairs {
			token0 := common.HexToAddress(pair.Token0.Id)
			token1 := common.HexToAddress(pair.Token1.Id)
			if _, ok := LPLoaded[uniswapAddr]; !ok {
				LPLoaded[uniswapAddr] = make(map[common.Address]map[common.Address]common.Address)
			}
			if _, ok := LPLoaded[uniswapAddr][token0]; !ok {
				LPLoaded[uniswapAddr][token0] = make(map[common.Address]common.Address)	
			}
			LPLoaded[uniswapAddr][token0][token1] = common.HexToAddress(pair.Id)		
			
		}

	}

	jsonStr, err := json.Marshal(LPLoaded)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("LPLoaded.json", jsonStr, 0644)
	}
}





