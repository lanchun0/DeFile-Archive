package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"dfa/entity"
	"dfa/smartcontract/behavior"
	"dfa/smartcontract/filesharing"

	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type smartcontract struct {
	behaviorContract *behavior.Behavior
	dataContract     *filesharing.Filesharing

	client      *ethclient.Client
	auth        *bind.TransactOpts
	serverIndex int
	servers     []string
}

// var Client *ethclient.Client
// var ServerIndex int = 0
var servers = [10]string{
	"0xf1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5",
	"0x91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9",
	"0xbb32062807c162a5243dc9bcf21d8114cb636c376596e1cf2895ec9e5e3e0a68",
	"0x95ce6122165d94aa51b0fcf51021895b39b0ff291aa640c803d5401bd87894d5",
	"0x3af93668029f95d526fc1d2bdefccc120bfe1d26a0462d268e8f6b2f71402ba3",
	"0x3b24a4fdf2e6e1375008c387c5456ce00cb0772435ae1938c2fe833103393b9a",
	"0xcba858feeb49e1ca8053a5213987a22c3ee83d9f9f396e138940a018dd837ebb",
	"0xdf48bfda4cb4b4e094803e923836a8538fbf607da79f6e46d68cdd43fb2f3f88",
	"0x487efb6249a8a4d45a19c8e5d1e5c7d3f6610a7e69f8f81ddcf368f9a0c0d6d5",
	"0xbb4cce73db59f456ea427e5862fdb0d5bc038a7d0b930cbb45e1c4f6d122289e",
}

type SmartContract interface {
	DeployContract() (err error)
	Register() (entity.Behavior, string, error)
	Login(pub, signature string) (entity.Behavior, error)
	changeAuth()
}

func GetConract() SmartContract {
	return &smartcontract{
		serverIndex: 0,
		servers:     servers[:],
	}
}

func (contract *smartcontract) DeployContract() (err error) {
	contract.client, err = ethclient.Dial("http://localhost:8545")
	if err != nil {
		return err
	}
	contract.changeAuth()
	address_1, tx_1, behaviorContract, err := behavior.DeployBehavior(contract.auth, contract.client)
	if err != nil {
		return err
	}
	fmt.Printf("Behavior contract deployed:\n address: %s\n transaction: %s\n", address_1.Hex(), tx_1.Hash().Hex())

	contract.changeAuth()
	address_2, tx_2, dataContract, err := filesharing.DeployFilesharing(contract.auth, contract.client, address_1)
	if err != nil {
		return err
	}
	fmt.Printf("Data contract deployed:\n address: %s\n transaction: %s\n", address_2.Hex(), tx_2.Hash().Hex())
	contract.behaviorContract, contract.dataContract = behaviorContract, dataContract
	return nil
}

func (contract *smartcontract) changeAuth() {
	contract.serverIndex = (contract.serverIndex + 1) % len(contract.servers)
	priv := contract.servers[contract.serverIndex][2:]
	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println(err)
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := contract.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	gasPrice, err := contract.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	contract.auth = bind.NewKeyedTransactor(privateKey)
	contract.auth.Nonce = big.NewInt(int64(nonce))
	contract.auth.Value = big.NewInt(0)      // in wei
	contract.auth.GasLimit = uint64(6721975) // in units
	contract.auth.GasPrice = gasPrice
}
