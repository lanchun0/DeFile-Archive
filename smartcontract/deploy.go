package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"dfa/smartcontract/filesharing"
	"dfa/smartcontract/token"

	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (contract *smartcontract) DeployContract() (err error) {
	contract.client, err = ethclient.Dial("http://localhost:8545")
	if err != nil {
		return err
	}
	// contract.changeAuth()
	auth, _, err := contract.parseIdentity(contract.servers[0])
	if err != nil {
		return fmt.Errorf("failed to deploy token contract: %v", err)
	}
	address_1, tx_1, userContract, err := token.DeployToken(auth, contract.client)
	if err != nil {
		return err
	}
	fmt.Printf("For For Token contract deployed:\n address: %s\n transaction: %s\n", address_1.Hex(), tx_1.Hash().Hex())

	auth, _, err = contract.parseIdentity(contract.servers[1])
	if err != nil {
		return fmt.Errorf("failed to deploy file contract: %v", err)
	}
	address_2, tx_2, dataContract, err := filesharing.DeployFilesharing(auth, contract.client, address_1)
	if err != nil {
		return err
	}
	contract.dataAddress, contract.userAddress = address_2, address_1
	fmt.Printf("Data contract deployed:\n address: %s\n transaction: %s\n", address_2.Hex(), tx_2.Hash().Hex())
	contract.userContract, contract.dataContract = userContract, dataContract
	return nil
}

func (contract *smartcontract) parseIdentity(priv string) (*bind.TransactOpts, common.Address, error) {
	if len(priv) <= 2 {
		return nil, common.Address{}, fmt.Errorf("invalid private key")
	}
	privateKey, err := crypto.HexToECDSA(priv[2:])
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("invalid private key")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println(err)
		return nil, common.Address{}, fmt.Errorf("cannot parse the public key from %s", priv)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := contract.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, common.Address{}, err
	}

	gasPrice, err := contract.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, common.Address{}, fmt.Errorf("cannot set gas fee: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice
	return auth, fromAddress, nil
}

func (contract *smartcontract) GetDataContractAddress() string {
	return contract.dataAddress.Hex()
}

func (contract *smartcontract) GetUserContractAddress() string {
	return contract.userAddress.Hex()
}
