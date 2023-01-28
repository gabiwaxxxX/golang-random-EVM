package main

import (

	"log"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"


	"github.com/spf13/viper"
	// "github.com/ethereum/go-ethereum/common"
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
	fmt.Println("loading tokens...")

	jsonFile, err := os.Open("./data/tokens/" + "9.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var finalTokens Data
	json.Unmarshal([]byte(byteValue), &finalTokens)


	folders, err := ioutil.ReadDir("./data/tokens")
	if err != nil {
		log.Fatal(err)
	}
	for _, pairFile := range folders {
		pairFileName := pairFile.Name()
		if pairFileName != "9.json" {
			jsonFile, err := os.Open("./data/tokens/" + pairFileName)
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()
		
			byteValue, _ := ioutil.ReadAll(jsonFile)
		
			var tokens Data

			json.Unmarshal([]byte(byteValue), &tokens)
			finalTokens.Data.Tokens = append(finalTokens.Data.Tokens, tokens.Data.Tokens...)
		}

	}

	jsonStr, err := json.Marshal(finalTokens)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		ioutil.WriteFile("Tokens.json", jsonStr, 0644)
	}
}





