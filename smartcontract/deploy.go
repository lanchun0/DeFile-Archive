package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"dfa/smartcontract/behavior"
	"dfa/smartcontract/filesharing"

	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
