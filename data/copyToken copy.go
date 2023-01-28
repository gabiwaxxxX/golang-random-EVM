package main

import (

	"log"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"

)


type Tokens struct {
    Tokens []Token `json:"tokens"`
}

type Token struct {
    Address			string 	`json:"address"`
    Symbol		string 	`json:"symbol"`
	Decimals		int 	`json:"decimals"`
	ChainId		int 	`json:"chainId"`

}



func main() {

	// sync pairs
	fmt.Println("loading tokens...")

	jsonFile, err := os.Open("./data/SushiSwapData/tokenList/all.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var Ultimate Tokens
	json.Unmarshal([]byte(byteValue), &Ultimate)


	folders, err := ioutil.ReadDir("./data/SushiSwapData/tokenList")
	if err != nil {
		log.Fatal(err)
	}
	for _, pairFile := range folders {
		pairFileName := pairFile.Name()
		if pairFileName != "all.json" {
			jsonFile, err := os.Open("./data/SushiSwapData/tokenList/" + pairFileName)
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()
		
			byteValue, _ := ioutil.ReadAll(jsonFile)
		
			var tokens Tokens

			json.Unmarshal([]byte(byteValue), &tokens)

			for _, token := range tokens.Tokens {
				var result bool = false
				for _, x := range Ultimate.Tokens {
					if x.Symbol == token.Symbol {
						result = true
						break
					} 
				}
				if result == false {
					if token.ChainId == 1 {
						Ultimate.Tokens = append(Ultimate.Tokens, token)
					}
				}
		
			}

		}

	}

	jsonStr, err := json.Marshal(Ultimate)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("TokensWithDecimals.json", jsonStr, 0644)
	}
}





