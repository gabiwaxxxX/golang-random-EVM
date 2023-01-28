package nucleus

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/gabrielfournier1/bot/contract/UniswapV2Factory"
	"github.com/gabrielfournier1/bot/contract/UniswapV2Pair"
	"github.com/gabrielfournier1/bot/rpcClient"
	"github.com/gabrielfournier1/bot/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient/gethclient"
)

const UNI_SWAP_FACTORY_CONTRACT_ADDRESS = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const SUSHI_FACTORY_ADDRESS = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
const CONTRACT_DIR = "contract"

// read ABIs from abis folder
func ReadABIs() (map[string]abi.ABI, error) {
	folders, err := ioutil.ReadDir(CONTRACT_DIR)
	if err != nil {
		log.Fatal(err)
	}
	abis := make(map[string]abi.ABI)
	for _, folder := range folders {
		if folder.IsDir() {
			abiFiles, err := ioutil.ReadDir(CONTRACT_DIR + "/" + folder.Name())
			if err != nil {
				log.Fatal(err)
			}

			for _, abiFile := range abiFiles {
				//check if json file
				if abiFile.Name()[len(abiFile.Name())-5:] == ".json" {
					fmt.Printf("abi file: %s\n", abiFile.Name())
					file, err := ioutil.ReadFile(CONTRACT_DIR + "/" + folder.Name() + "/" + abiFile.Name())
					// 	file, err := ioutil.ReadFile(ABIS_DIR + "/" + name)
					_abi, err := abi.JSON(strings.NewReader(string(file)))
					if err != nil {
						log.Fatal(err)
					}
					abiName := abiFile.Name()[:len(abiFile.Name())-5]
					abis[abiName] = _abi
				}
			}
		}
	}
	return abis, nil
}

// compare addresses
func compareAddresses(a common.Address, b common.Address) int {
	return bytes.Compare(a.Bytes(), b.Bytes())
}

func LoadUniToSushi() (map[string]string, error) {
	uniFactoryAddress := common.HexToAddress(UNI_SWAP_FACTORY_CONTRACT_ADDRESS)
	sushiFactoryAddress := common.HexToAddress(SUSHI_FACTORY_ADDRESS)

	// read ABIs
	abis, err := ReadABIs()
	if err != nil {
		return nil, err
	}
	rpcClient.Initialize("./rpcClient/config.json")
	blockNumber, err := rpcClient.WSClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	// get PairCreated logs
	// batch to avoid errors
	event := abis["UniswapV2Factory"].Events["PairCreated"]
	query := ethereum.FilterQuery{
		Addresses: []common.Address{uniFactoryAddress, sushiFactoryAddress},
		Topics:    [][]common.Hash{[]common.Hash{event.ID}},
	}
	batchSize := uint64(1000000)

	var wg sync.WaitGroup
	var pairsMu sync.Mutex
	pairs := make(map[common.Address]map[common.Address]common.Address)
	var sushiPairsMu sync.Mutex
	sushiPairs := make(map[common.Address]map[common.Address]common.Address)
	for i := uint64(0); i <= blockNumber; i += batchSize {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
			query := query
			query.FromBlock = new(big.Int).SetUint64(i)
			query.ToBlock = new(big.Int).SetUint64(i + batchSize - 1)

			logs, err := rpcClient.WSClient.FilterLogs(context.Background(), query)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, vLog := range logs {
				// get factory address
				factoryAddress := vLog.Address

				// get pair address
				data, err := event.Inputs.Unpack(vLog.Data)
				if err != nil {
					fmt.Println(err)
					continue
				}
				pairAddress := data[0].(common.Address)

				// get token addresses
				token0Address := common.HexToAddress(vLog.Topics[1].Hex())
				token1Address := common.HexToAddress(vLog.Topics[2].Hex())

				// store pair
				if compareAddresses(factoryAddress, uniFactoryAddress) == 0 {
					pairsMu.Lock()
					if _, ok := pairs[token0Address]; !ok {
						pairs[token0Address] = make(map[common.Address]common.Address)
					}
					pairs[token0Address][token1Address] = pairAddress
					pairsMu.Unlock()
				} else if compareAddresses(factoryAddress, sushiFactoryAddress) == 0 {
					sushiPairsMu.Lock()
					if _, ok := sushiPairs[token0Address]; !ok {
						sushiPairs[token0Address] = make(map[common.Address]common.Address)
					}
					sushiPairs[token0Address][token1Address] = pairAddress
					sushiPairsMu.Unlock()
				} else {
					fmt.Println("?")
					continue
				}
			}
		}()
	}
	wg.Wait()

	uniToSushi := make(map[string]string)
	for token0Address := range pairs {
		for token1Address := range pairs[token0Address] {
			uni := pairs[token0Address][token1Address]
			sushi, ok := sushiPairs[token0Address][token1Address]
			if !ok {
				sushi = uni
			} else {
				fmt.Printf("founded uni pair: %s and sushi pair: %s\n", uni, sushi)
			}
			uniToSushi[strings.ToLower(uni.String())] = sushi.String()
		}
	}
	return uniToSushi, nil
}

func LoadUniswapLP(factoryAddress string) (map[common.Address]map[common.Address]common.Address, error) {

	factory, err := UniswapV2Factory.NewUniswapV2Factory(common.HexToAddress(factoryAddress), utils.GetClient())
	if err != nil {
		log.Fatal(err)
	}

	// get pairs
	nbrOfUniswapPairs, err := factory.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	numPairs := nbrOfUniswapPairs.Uint64()
	// get pairs
	pairs := make(map[common.Address]map[common.Address]common.Address)
	var wg sync.WaitGroup
	var pairsMu sync.Mutex
	var batchSize = uint64(5000)

	for i := uint64(0); i <= numPairs; i += batchSize {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := i; j <= i+batchSize && j <= numPairs; j += 1 {

				fmt.Printf("%d/%d\n", j, nbrOfUniswapPairs)
				querryPairNumber := big.NewInt(int64(j))
				pair, err := factory.AllPairs(&bind.CallOpts{}, querryPairNumber)
				if err != nil {
					log.Fatal(err)
				}

				pairToken, err := UniswapV2Pair.NewUniswapV2Pair(pair, utils.GetClient())
				if err != nil {
					log.Fatal(err)
				}
				token0Address, err := pairToken.Token0(nil)
				if err != nil {
					log.Fatal(err)
				}
				token1Address, err := pairToken.Token1(nil)
				if err != nil {
					log.Fatal(err)
				}
				pairsMu.Lock()

				if _, ok := pairs[token0Address]; !ok {
					pairs[token0Address] = make(map[common.Address]common.Address)
				}
				pairs[token0Address][token1Address] = pair
				pairsMu.Unlock()
			}

		}()
	}
	wg.Wait()
	return pairs, nil
}

func FindCommonLp(dex0 map[common.Address]map[common.Address]common.Address, dex1 map[common.Address]map[common.Address]common.Address) (map[string]string, error) {

	common := make(map[string]string)
	//check the length and chose the shortest
	if len(dex0) > len(dex1) {
		dex0, dex1 = dex1, dex0
	}
	for token0Address := range dex0 {
		for token1Address := range dex0[token0Address] {
			zero := dex0[token0Address][token1Address]
			one, ok := dex1[token0Address][token1Address]
			if ok {
				common[strings.ToLower(zero.String())] = one.String()
			}
		}
	}

	// for _, value := range dex0 {
	// 	for _, value2 := range value {
	// 		for _, value3 := range value2 {
	// 			onlyLP = append(onlyLP, value3)
	// 		}
	// 	}
	// }
	return common, nil
}

func LoadLPFromEvent(factorysAddress []string) (map[common.Address]map[common.Address]map[common.Address]common.Address, error) {
	var Addresses []common.Address

	for _, factoryAddress := range factorysAddress {
		Addresses = append(Addresses, common.HexToAddress(factoryAddress))
	}

	// read ABIs
	abis, err := ReadABIs()
	if err != nil {
		return nil, err
	}
	rpcClient.Initialize("./rpcClient/carlito.json")
	blockNumber, err := rpcClient.WSClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	// get PairCreated logs
	// batch to avoid errors
	event := abis["UniswapV2Factory"].Events["PairCreated"]
	query := ethereum.FilterQuery{
		Addresses: Addresses,
		Topics:    [][]common.Hash{[]common.Hash{event.ID}},
	}
	batchSize := uint64(1000000)

	var wg sync.WaitGroup

	allPairs := make(map[common.Address]map[common.Address]map[common.Address]common.Address, len(factorysAddress))
	// listOfMutex := make(map[common.Address]sync.Mutex, len(factorysAddress))
	// for _, address := range Addresses {

	// 	listOfMutex[address] = mut
	// }
	// var client = utils.GetClient()

	// var ruggedCounter uint64
	var mut sync.Mutex
	for i := uint64(0); i <= blockNumber; i += batchSize {
		i := i
		// fmt.Println(i)
		wg.Add(1)
		go func() {
			defer wg.Done()

			query := query
			query.FromBlock = new(big.Int).SetUint64(i)
			query.ToBlock = new(big.Int).SetUint64(i + batchSize - 1)

			logs, err := rpcClient.WSClient.FilterLogs(context.Background(), query)
			if err != nil {
				fmt.Println(" Re init client")
				rpcClient.Initialize("./rpcClient/carlito.json")
				logs, err = rpcClient.WSClient.FilterLogs(context.Background(), query)
				if err != nil {
					fmt.Println(" You got an err on line 318 ", err)
					log.Fatal(err)
				}
			}

			for _, vLog := range logs {
				// get factory address
				factoryAddress := vLog.Address

				// get pair address
				data, err := event.Inputs.Unpack(vLog.Data)
				if err != nil {
					fmt.Println(err)
					fmt.Println(" You got an err on line 327 ", err)
					continue
				}
				pairAddress := data[0].(common.Address)

				// get token addresses
				token0Address := common.HexToAddress(vLog.Topics[1].Hex())
				token1Address := common.HexToAddress(vLog.Topics[2].Hex())
				var added bool = false
				//check if Rug
				// rugged := utils.CheckAll(pairAddress,client)

				// if !rugged {
				for _, addresse := range Addresses {
					if compareAddresses(factoryAddress, addresse) == 0 {
						mut.Lock()
						if _, ok := allPairs[addresse]; !ok {
							allPairs[addresse] = make(map[common.Address]map[common.Address]common.Address)
						}
						if _, ok := allPairs[addresse][token0Address]; !ok {
							allPairs[addresse][token0Address] = make(map[common.Address]common.Address)

						}
						allPairs[addresse][token0Address][token1Address] = pairAddress
						added = true
						mut.Unlock()
					}
				}
				if !added {
					fmt.Println(pairAddress, " ?")
					continue
				}
				// }
				// if rugged {
				// 	ruggedCounter++
				// }

			}
		}()
	}
	wg.Wait()
	// fmt.Println("Rugged LPs: ",ruggedCounter)
	return allPairs, nil
}
