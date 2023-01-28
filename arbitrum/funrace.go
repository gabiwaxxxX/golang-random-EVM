package arbitrum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/contract/Furnace"
	"github.com/gabrielfournier1/bot/utils"
)

func SpammingAdvance() {
	furnaceAddress := common.HexToAddress("0x054Dcaa740Ee3A9C192A9917ecb178C433Eceb96")
	valueBig := StringToBigInt("1d1661cb61bf5e3066f17f82099786d0fcc49d46")

	furnaceContractTx, err := Furnace.NewFurnace(furnaceAddress, GetClient())
	if err != nil {
		fmt.Println("error 1")

		panic(err)
	}

	nextEpoch := big.NewInt(20)
	nextEpochTime := int64(1669221030)
	epochSquare := new(big.Int).Mul(nextEpoch, nextEpoch)
	valueBig.Mul(valueBig, epochSquare)

	for {

		// get current time in timestamp format
		currentTime := time.Now().Unix()
		// fmt.Println("current time", currentTime)

		if currentTime == nextEpochTime {

			//fmt.Println(valueBig)
			auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
			if err != nil {
				fmt.Println("error 0")

				panic(err)
			}

			auth.GasLimit = 3000000

			auth.GasPrice = big.NewInt(100000000000)
			furnaceContractTx.Advance(auth, valueBig)
		}

	}

}

func StringToBigInt(str string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(str, 16)
	if !ok {
		fmt.Println("error ")
	}
	return n
}
