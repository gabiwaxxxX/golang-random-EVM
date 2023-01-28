package listing

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	// "context"
	// "log"
	// "strings"
	// "math/big"
	// "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/core/types"
	// "project.main/contract/uniswap"
)

func listing() {
	client, err := ethclient.Dial("__")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

// type LogPairCreated struct {
// 	Token0	common.Address
// 	Token1 	common.Address
// 	Pair   	common.Address
// 	Num 	*big.Int
// }

// type UniswapV2PairCreated struct {
// 	Token0 common.Address
// 	Token1 common.Address
// 	Pair   common.Address
// 	Arg3   *big.Int
// 	Raw    types.Log // Blockchain specific contextual infos
// }

// func main() {
//     client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/RehQnHQOFGkON04oeEKygcqp-nbPI0ob")
//     if err != nil {
//         log.Fatal(err)
//     }

//     contractAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
//     query := ethereum.FilterQuery{
//         Addresses: []common.Address{contractAddress},
//     }

// 	logs, err := client.FilterLogs(context.Background(), query)
//     if err != nil {
//         log.Fatal(err)
//     }

// 	contractAbi, err := abi.JSON(strings.NewReader(string(uniswap.UniswapV2ABI)))
//     if err != nil {
//         log.Fatal(err)
//     }
// 	fmt.Println(contractAbi)

// 	fmt.Println("listing all the events of the uniswap router smart contract ...")

// 	logPairCreatedSig := []byte("PairCreated(address, address, address, uint256)")
//     logPairCreatedHash := crypto.Keccak256Hash(logPairCreatedSig)

// 	for _, vLog := range logs {
// 		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
//         fmt.Printf("Log Index: %d\n", vLog.Index)

// 		switch vLog.Topics[0].Hex() {
// 			case logPairCreatedHash.Hex():
// 				fmt.Printf("Log Name: Pair Created\n")

// 				// var pairCreatedEvent UniswapV2PairCreated

// 				test,err := contractAbi.Unpack("PairCreated", vLog.Data)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				fmt.Printf("%+v\n", test)
// 				// Fprintln(vLog)
// 				// fmt.Printf(vLog)

// 				// pairCreatedEvent.Token0 = common.HexToAddress(vLog.Topics[1].Hex())
// 				// pairCreatedEvent.Token1 = common.HexToAddress(vLog.Topics[2].Hex())

// 				// fmt.Printf("From: %s\n", pairCreatedEvent.From.Hex())
// 				// fmt.Printf("To: %s\n", pairCreatedEvent.To.Hex())
// 				// fmt.Printf("Tokens: %s\n", pairCreatedEvent.Tokens.String())
// 		}

//         // event := struct {
//         //     Token0	common.Address
//         //     Token1 	common.Address
// 		// 	Pair   	common.Address
// 		// 	Num 	*big.Int
//         // }{}
//         // err := contractAbi.Unpack(&event, "PairCreated", vLog.Data)
//         // if err != nil {
//         //     log.Fatal(err)
//         // }

//         // fmt.Println(string(event.Token0[:]))   // foo
//         // fmt.Println(string(event.Token1[:])) // bar
//         // fmt.Println(string(event.Pair[:])) // bar
//         // fmt.Println(string(event.Num[:])) // bar

//         // var topics [4]string
//         // for i := range vLog.Topics {
//         //     topics[i] = vLog.Topics[i].Hex()
//         // }

//         // fmt.Println(topics[0])
//     }

// }
