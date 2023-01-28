package walletBalance

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/nucleus"
	"github.com/gabrielfournier1/bot/rpcClient"
)

type GetTokenBalanceResponse struct {
	Result struct {
		Address       string         `json:"address"`
		TokenBalances []TokenBalance `json:"tokenBalances"`
	} `json:"result"`
}

type TokenBalance struct {
	ContractAddress string `json:"contractAddress"`
	Balance         string `json:"tokenBalance"`
}

func GetAllTokensOwnedByAnAdress() {

	url := "https://eth-mainnet.alchemyapi.io/v2/RehQnHQOFGkON04oeEKygcqp-nbPI0ob"

	payload := strings.NewReader("{\"id\":2,\"jsonrpc\":\"2.0\",\"method\":\"alchemy_getTokenBalances\",\"params\":[\"Your wallet adress\",\"erc20\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	var resp GetTokenBalanceResponse
	json.Unmarshal(body, &resp)
	// fmt.Println(resp.Result.TokenBalances)

	for _, token := range resp.Result.TokenBalances {
		//get only if balance is not 0 i hex
		if token.Balance != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			fmt.Println(token.ContractAddress, token.Balance)
		}
	}
	//length
	fmt.Println(len(resp.Result.TokenBalances))

}

func GetAllTransferEventsToAnAddress(address string) {
	abis, err := nucleus.ReadABIs()
	if err != nil {
		panic(err)
	}
	rpcClient.Initialize("./rpcClient/carlito.json")
	blockNumber, err := rpcClient.WSClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	event := abis["erc20"].Events["Transfer"]
	fmt.Println(event)

	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{[]common.Hash{event.ID}},
	}
	var wg sync.WaitGroup

	batchSize := uint64(1000000)
	for i := uint64(13670000); i <= blockNumber; i += batchSize {
		i := i
		wg.Add(1)
		go func() {
			fmt.Println("Processing block range", i, i+batchSize)
			defer wg.Done()
			query.FromBlock = big.NewInt(int64(i))
			query.ToBlock = big.NewInt(int64(i + batchSize))
			logs, err := rpcClient.WSClient.FilterLogs(context.Background(), query)
			if err != nil {
				return
			}
			for _, log := range logs {

				fmt.Println(log)
			}
		}()
		wg.Wait()

	}
}

func GetAllTransactionsOfAnAdress(address string) {

}
