package eth_const

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
	"os"
)

type Tokens struct {
    Tokens []Token `json:"tokens"`
}

type Token struct {
    Symbol		string 	`json:"symbol"`
    Address		string 	`json:"address"`
}

func GetEthTokens() Tokens {
	// Open our jsonFile
	jsonFile, err := os.Open("./eth_const/tokens.uniswap.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tokens Tokens
	json.Unmarshal([]byte(byteValue), &tokens)
	return tokens
}

func FindAddressFromTokenSymbol(symbol string) string {
	tokens := GetEthTokens()
	for _, token := range tokens.Tokens {
		if token.Symbol == symbol {
			return token.Address
		}
	}
	return ""
}

func FindTokenSymbolFromAddress(address string) string {
	tokens := GetEthTokens()
	for _, token := range tokens.Tokens {
		if token.Address == address {
			return token.Symbol
		}
	}
	return ""
}
