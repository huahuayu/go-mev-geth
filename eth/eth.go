package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	rpcProvider = "https://mainnet.infura.io/v3/ce918b7f57a24f138ba3973bae59da4f"
	EthClient   *ethclient.Client
)

func Init() {
	var err error
	if EthClient, err = ethclient.DialContext(context.Background(), rpcProvider); err != nil {
		log.Fatal(err)
	}
}
