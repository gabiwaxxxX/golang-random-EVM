package checkHoneyPot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
}

// curl -X POST -H "Content-Type: application/json" --data '{"method": "eth_getCode", "params": ["0x10393c20975cF177a3513071bC110f7962CD67da", "latest"],"id":1,"jsonrpc":"2.0"}' https://rpc.ankr.com/arbitrum
func curlPost(bc string, address string) string {
	var rpcAddress string
	switch bc {
	case "bsc":
		rpcAddress = "https://bsc-dataseed3.defibit.io"
	case "eth":
		rpcAddress = "https:/_____"
	case "arbitrum":
		rpcAddress = "https://rpc.ankr.com/arbitrum"
	case "matic":
		rpcAddress = "https://rpc-mainnet.matic.quiknode.pro"
	case "ftm":
		rpcAddress = "https://rpc.ftm.tools"
	case "avax":
		rpcAddress = "https://rpc.ankr.com/avalanche"
	default:
		rpcAddress = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	}
	payload := Payload{
		Method:  "eth_getCode",
		Params:  []string{address, "latest"},
		ID:      1,
		Jsonrpc: "2.0",
	}
	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(rpcAddress, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
