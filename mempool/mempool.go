package mempool

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"io/ioutil"
	"log"
	"strings"

	"github.com/gabrielfournier1/bot/uniswapV3Decode"
	"github.com/gabrielfournier1/bot/utils"

	"time"
	// "os"
)

func BaseClient() *rpc.Client {
	// var rpcAddress string = "https:/_____"

	var rpcAddress string = "ws:/_____"
	baseClient, err := rpc.Dial(rpcAddress)
	if err != nil {
		log.Fatalln(err)
	}
	return baseClient
}

func ListenOnMemPool() {
	ctx := context.Background()
	txnsHash := make(chan common.Hash)

	baseClient := BaseClient()

	ethClient := utils.GetClient()
	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	subscriber := gethclient.New(baseClient)
	_, err = subscriber.SubscribePendingTransactions(ctx, txnsHash)

	if err != nil {
		log.Fatalln(err)
	}

	signer := types.NewLondonSigner(chainID)

	defer func() {
		baseClient.Close()
		ethClient.Close()
	}()
	for txnHash := range txnsHash {
		// fmt.Println(txnHash)
		txn, _, err := ethClient.TransactionByHash(ctx, txnHash)
		if err != nil {
			continue
		}

		message, err := txn.AsMessage(signer, nil)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("==========================")
		fmt.Println("to", message.To().Hex())
		fmt.Println("hash", txnHash.Hex())
		fmt.Println("data : ", common.Bytes2Hex(txn.Data()))
		fmt.Println("gas : ", txn.Gas())
		fmt.Println("gasPrice : ", txn.GasPrice())
		fmt.Println("==========================")
	}
}

func ListenOnMemPoolTargetContract(targetContracts []string) {
	ctx := context.Background()
	txnsHash := make(chan common.Hash)
	_abiUniV3JSON, err := ioutil.ReadFile("/Users/GabrielFournier/Desktop/Kualian/botToSnip/go-implementation/main/bot/contract/UniswapV3Router/UniswapV3Router.json")
	if err != nil {
		log.Fatal(err)
	}

	G_uniV3abi, err := abi.JSON(strings.NewReader(string(_abiUniV3JSON)))
	if err != nil {
		log.Fatal(err)
	}

	baseClient := BaseClient()

	ethClient := utils.GetClient()
	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	subscriber := gethclient.New(baseClient)
	_, err = subscriber.SubscribePendingTransactions(ctx, txnsHash)

	if err != nil {
		log.Fatalln(err)
	}

	signer := types.NewLondonSigner(chainID)

	defer func() {
		baseClient.Close()
		ethClient.Close()
	}()
	for txnHash := range txnsHash {
		// fmt.Println(txnHash)
		txn, _, err := ethClient.TransactionByHash(ctx, txnHash)
		if err != nil {
			continue
		}

		message, err := txn.AsMessage(signer, nil)
		if err != nil {
			log.Fatalln(err)
		}

		// fmt.Println("txn",txn)

		for _, targetContract := range targetContracts {

			if message.To() == nil {
				continue
			}
			if message.To().Hex() == targetContract {
				if txn.Value().Cmp(utils.FormatEthToWei(1)) > 0 {
					fmt.Println("==========================")
					fmt.Println("hash", txnHash.Hex())
					t := time.Now()
					fmt.Println(t.Format("2006-01-02 15:04:05"))

					fmt.Println("Value", utils.FormatEthWeiToEther(txn.Value()), " $ETH")
					fmt.Println("From", message.From().Hex())

					go uniswapV3Decode.DecodeInput(txn.Value(), G_uniV3abi, common.Bytes2Hex(txn.Data()))
				}
				// fmt.Println("data",common.Bytes2Hex(txn.Data()))
			}
		}
	}
}
func ListenOnMemPoolUser(userToWatch common.Address) {
	fmt.Println("Listening on mempool for user : ", userToWatch)
	ctx := context.Background()
	txnsHash := make(chan common.Hash)

	baseClient := BaseClient()

	ethClient := utils.GetClient()
	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	subscriber := gethclient.New(baseClient)
	_, err = subscriber.SubscribePendingTransactions(ctx, txnsHash)

	if err != nil {
		log.Fatalln(err)
	}

	signer := types.NewLondonSigner(chainID)

	defer func() {
		baseClient.Close()
		ethClient.Close()
	}()
	for txnHash := range txnsHash {
		// fmt.Println(txnHash)
		txn, _, err := ethClient.TransactionByHash(ctx, txnHash)
		if err != nil {
			continue
		}

		message, err := txn.AsMessage(signer, nil)
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println("from",message.From().Hex())
		if message.From() == userToWatch {
			t := time.Now()
			fmt.Println(t.Format("2006-01-02 15:04:05"))
			fmt.Println("to", message.To().Hex())
			fmt.Println("==========================")
			// fmt.Println("hash", txnHash.Hex())
			fmt.Println("data : ", common.Bytes2Hex(txn.Data()))
			fmt.Println("gas : ", txn.Gas())
			fmt.Println("gasPrice : ", txn.GasPrice())
			fmt.Println("==========================")
		}

		// fmt.Println("transaction hash ",txnHash.String())
	}
}
