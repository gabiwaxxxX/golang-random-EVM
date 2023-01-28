package manageData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

)


type Tokens struct {
    Tokens []Token `json:"tokens"`
}

type Token struct {
    Address		string 	`json:"address"`
    Symbol		string 	`json:"symbol"`
	Decimals	int 	`json:"decimals"`
	ChainId		int 	`json:"chainId"`
}

func getTokens () Tokens {
	jsonFile, err := os.Open("/Users/GabrielFournier/Desktop/Kualian/botToSnip/go-implementation/main/bot/TokensWithDecimals.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var tokens Tokens
	json.Unmarshal([]byte(byteValue), &tokens)
	return tokens
}

func GetMapTokenAddressToDecimals () map[string]int {
	var tokens Tokens = getTokens()
	decimalsMap := make(map[string]int)
	for _, token := range tokens.Tokens {
		decimalsMap[token.Address] = token.Decimals
	}
	return decimalsMap
}

func GetMapTokenAddressToSymbol () map[string]string {
	var tokens Tokens = getTokens()

	symbolMap := make(map[string]string)
	for _, token := range tokens.Tokens {
		symbolMap[token.Address] = token.Symbol
	}
	return symbolMap
}

func GetTokenDecimals (tokenAddress string) uint8 {

	decimalsMap := GetMapTokenAddressToDecimals()
	if _, ok := decimalsMap[strings.ToLower(tokenAddress)]; !ok {
		//does not exist
		return uint8(0)
	}
	uint8Decimals := uint8(decimalsMap[strings.ToLower(tokenAddress)])
	return uint8Decimals
}

func GetTokenSymbol (tokenAddress string) string {
	symbolMap := GetMapTokenAddressToSymbol()
	if _, ok := symbolMap[strings.ToLower(tokenAddress)]; !ok {
		//does not exist
		return ""
	}
	symbol := symbolMap[strings.ToLower(tokenAddress)]
	return symbol
}

func ReWriteAllTokenAddressToCommonAddress() {
	var tokens Tokens = getTokens()
	for _, token := range tokens.Tokens {
		token.Address = strings.ToLower(token.Address )
	}
	//save tokens
	jsonStr, err := json.Marshal(tokens)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("TokensWithDecimalsChangeds.json", jsonStr, 0644)
	}
}