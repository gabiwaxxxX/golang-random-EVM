package main

import (

	"os"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Data struct {
    Data Tokens `json:"data"`
}

type Tokens struct {
    Tokens []Token `json:"tokens"`
}

type Token struct {
    Id			string 	`json:"id"`
    Symbol		string 	`json:"symbol"`

}



func main() {
	fmt.Println("loading tokens...")

	jsonFile, err := os.Open("./Tokens.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var finalTokens Data
	json.Unmarshal([]byte(byteValue), &finalTokens)
	var FreshTokens []Token


	for _, token := range finalTokens.Data.Tokens {
		var result bool = false
		for _, x := range FreshTokens {
			if x == token {
				result = true
				break
			} 
		}
		if result == false {
			FreshTokens = append(FreshTokens, token)
		}

	}

	jsonStr, err := json.Marshal(FreshTokens)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("FreshTokens.json", jsonStr, 0644)
	}
}





