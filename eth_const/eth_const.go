package eth_const

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	UniSwapV2RouterAddress  = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	UniSwapV2FactoryAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

const (
	EthRpcAddr = "https://mainnet.infura.io/v3/fe05e55ba1724b0589f19037b0a54a38"
)

const (
	ChainTypeEth = "eth"
)

var (
	ZeroAddress = common.Address{}
	ChainIdMap  = map[string]int64{
		ChainTypeEth: 1,
	}

	UniSwapFactoryContractMap = map[string]common.Address{
		ChainTypeEth: common.HexToAddress(UniSwapV2FactoryAddress),
	}

	UniSwapRouterContractMap = map[string]common.Address{
		ChainTypeEth: common.HexToAddress(UniSwapV2RouterAddress),
	}
)




