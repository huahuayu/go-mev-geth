package bundle

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mev-geth/common/http"
	"math/big"
)

var (
	Relay               = "https://relay.flashbots.net"
	TestnetRelay        = "https://relay-goerli.flashbots.net"
	method_simulate     = "eth_callBundle"
	method_send         = "eth_sendBundle"
	method_getUserStats = "flashbots_getUserStats"
)

type (
	JsonRpc struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int64         `json:"id"`
	}

	Bundle struct {
		Signer             string   `json:"signer"`
		SignedTransactions []string `json:"signedTransactions"`
		BlockNumber        string   `json:"blocknumber"`
		MinTimestamp       int      `json:"minTimestamp"`
		MaxTimestamp       int      `json:"maxTimestamp"`
	}

	Response interface{}

	ErrResponse struct {
		Error struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}

	SimulationResponseSuccess struct {
		bundleHash   string
		coinbaseDiff *big.Int
		results      []interface{}
		totalGasUsed uint64
		firstRevert  []interface{}
	}
)

func NewBundle(signer string, signedTransactions []string, blockNumber uint64) *Bundle {
	return &Bundle{Signer: signer, SignedTransactions: signedTransactions, BlockNumber: "0x" + fmt.Sprintf("%x", blockNumber)}
}

func (b *Bundle) Send() (res Response, err error) {
	payload := b.prepareRequest(method_send)
	signature, err := b.sign(payload)
	if err != nil {
		return nil, err
	}
	signerAddress, err := b.signerAddress()
	if err != nil {
		return nil, err
	}
	err = http.Post(Relay, payload, res, map[string]string{"X-Flashbots-Signature": signerAddress + ":" + signature})
	return res, err
}

func signHash(hashString string) []byte {
	return crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n66" + hashString))
}

func (b *Bundle) sign(jsonRpc JsonRpc) (signature string, err error) {
	marshal, err := json.Marshal(jsonRpc)
	if err != nil {
		return "", err
	}
	fmt.Println(string(marshal))
	ecdsaPrivateKey, err := crypto.HexToECDSA(b.Signer)
	if err != nil {
		return "", err
	}
	signatureBytes, err := crypto.Sign(signHash(hexutil.Encode(crypto.Keccak256(marshal))), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signatureBytes), nil
}

func (b *Bundle) Simulate() (res Response, err error) {
	payload := b.prepareRequest(method_simulate)
	signature, err := b.sign(payload)
	if err != nil {
		return nil, err
	}
	singerAddress, err := b.signerAddress()
	if err != nil {
		return nil, err
	}
	err = http.Post(Relay, payload, res, map[string]string{"X-Flashbots-Signature": singerAddress + ":" + signature})
	return res, err
}

func (b *Bundle) prepareRequest(method string) JsonRpc {
	return JsonRpc{Jsonrpc: "2.0", Method: method, Params: []interface{}{b.SignedTransactions, b.BlockNumber, b.MinTimestamp, b.MaxTimestamp}, ID: 1}
}

func (b *Bundle) signerAddress() (address string, err error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(b.Signer)
	if err != nil {
		return "", err
	}
	publicKey := ecdsaPrivateKey.PublicKey
	addr := crypto.PubkeyToAddress(publicKey)
	return addr.String(), nil
}
