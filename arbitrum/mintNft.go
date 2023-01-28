package arbitrum

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gabrielfournier1/bot/contract/Adv3nturers"
	"github.com/gabrielfournier1/bot/utils"
)

func MintRandomNft() {
	client := GetClient()
	NFTContract, err := Adv3nturers.NewAdv3nturers(common.HexToAddress("0x8406F8D5fe3869588d5F895878745078122941E9"), client)

	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(utils.GetUserInfo().PrivateKey, big.NewInt(42161))
	if err != nil {
		log.Fatal(err)
	}

	//generate random number between 1 and 8000

	// for loop over i < 8000
	for i := 1; i < 8000; i++ {
		tokenId := big.NewInt(int64(i))
		tx, err := NFTContract.Claim(auth, tokenId)
		if err != nil {
			fmt.Println("error ir", err)
			continue
		}
		AwaitToConfirmTx(tx)

	}
}
