package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2e7349ea207f44d9b4145310179b673a")

	logFatal(err)

	account := common.HexToAddress("0xC94D85fF952457c3F9F1FeF88eC052df6FAe2F4e")
	balance, err := client.BalanceAt(context.Background(), account, nil)

	logFatal(err)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())

	// convert wei to eth
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Account Balance Eth:", ethValue)
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
