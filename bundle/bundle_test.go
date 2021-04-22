package bundle

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mev-geth/eth"
	"go-mev-geth/eth/contract/mevtransfer"
	"math/big"
	"testing"
)

var (
	sender   = "$privateKey"
	contract = "0x5392b5fE24367d4D8079DE9C3c6e2b7B2eB632a7"
	signer   = "1313a69334ed0bebe685b690233bfd5a8137c727eb09233506fab901f7f6ff80" // random zero-balance EOA
)

func TestBundle_Simulate(t *testing.T) {
	eth.Init()
	nonce, err := getNonce(sender)
	if err != nil {
		t.Fatal(err)
	}
	blockNumber, err := eth.EthClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	signedTx1 := signTx(t, nonce, 1, 0)
	signedTx2 := signTx(t, nonce+1, 10000000000000001, 10000000000000000)
	bundle := NewBundle(signer, []string{signedTx1, signedTx2}, blockNumber+1)
	res, err := bundle.Simulate()
	if err != nil {
		t.Log(res)
		t.Fatal(err)
	}
	t.Log(res)
}

func getTransactor(privateKey string) (transactor *bind.TransactOpts, err error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := ecdsaPrivateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)
	chainID, err := eth.EthClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	transactor, err = bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	transactor.Context = context.Background()
	transactor.Value = big.NewInt(0)
	transactor.GasPrice = big.NewInt(0)
	transactor.GasLimit = 0
	transactor.From = fromAddress
	return transactor, nil
}

func signTx(t *testing.T, nonce uint64, value int64, bribe int64) (signedTx string) {
	instance, err := mevtransfer.NewMevtransfer(common.HexToAddress(contract), eth.EthClient)
	if err != nil {
		t.Fatal(err)
	}
	transactor, err := getTransactor(sender)
	if err != nil {
		t.Fatal(err)
	}
	transactor.NoSend = true
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(value)
	tx, err := instance.Transfer(transactor, big.NewInt(bribe))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	return hexutil.Encode(bytes)
}

func getNonce(privateKey string) (nonce uint64, err error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return 0, err
	}
	publicKey := ecdsaPrivateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, err = eth.EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return 0, err
	}
	return nonce, err
}
