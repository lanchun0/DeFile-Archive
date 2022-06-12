package smartcontract

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract/token"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (contract *smartcontract) Register() (entity.User, string, error) {
	pub, priv, addr := general.GenerateECSDAKey()
	fmt.Println("Generated new user, address: ", addr)
	fmt.Println("Public key: ", pub)
	b := entity.User{
		Address: addr,
		Balance: 0,
	}
	return b, priv, nil
}

// sign the public key to login
func (contract *smartcontract) Login(priv string) (entity.User, error) {
	_, addr, err := contract.parseIdentity(priv)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to login: %v", err)
	}
	opts := &bind.CallOpts{From: addr}
	user, err := contract.userContract.QueryUser(opts)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to query on token blockchain: %v", err)
	}
	u := token2User(&user)
	return u, nil
}

func (contract *smartcontract) Approve(priv string, price uint64) (bool, string, error) {
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot approve: invalid identity: %s", priv)
	}
	priceBig := new(big.Int).SetUint64(price)
	tx, err := contract.userContract.Approve(auth, contract.dataAddress, priceBig)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot approve: %v", err)
	}
	fmt.Println("Successful approvement: ", tx.Hash().Hex())
	return true, tx.Hash().Hex(), nil
}

func (contract *smartcontract) Topup(priv string, amount uint64) (bool, string, error) {
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot topup: invalid identity: %s", priv)
	}
	value := amount*3600 + 1
	auth.Value = new(big.Int).SetUint64(value)
	tokenValue := new(big.Int).SetUint64(amount)
	tx, err := contract.userContract.Topup(auth, tokenValue)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot topup: %v", err)
	}
	fmt.Println("Successful topup: ", tx.Hash().Hex())
	return true, tx.Hash().Hex(), nil
}

func (contract *smartcontract) WithDraw(priv string, amount uint64) (bool, string, error) {
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot withdraw: invalid identity: %s", priv)
	}
	// auth.Value = new(big.Int).SetUint64(amount)
	value := new(big.Int).SetUint64(amount)
	tx, err := contract.userContract.Withdraw(auth, value)
	if err != nil {
		return false, "", fmt.Errorf("error: cannot withdraw: %v", err)
	}
	fmt.Println("Successful withdraw: ", tx.Hash().Hex())
	return true, tx.Hash().Hex(), nil
}

func (contract *smartcontract) GetAllowance(priv string) (uint64, error) {
	_, addr, err := contract.parseIdentity(priv)
	if err != nil {
		return 0, fmt.Errorf("error: cannot query allowance: invalid identity: %s", priv)
	}
	amount, err := contract.userContract.Allowance(&bind.CallOpts{
		From: common.HexToAddress(contract.wallet[1]),
	}, contract.dataAddress, addr)
	if err != nil {
		return 0, fmt.Errorf("error: cannot query allowance:  %v", err)
	}
	return amount.Uint64(), nil
}

func (contract *smartcontract) GetApproved(priv string) (uint64, error) {
	_, addr, err := contract.parseIdentity(priv)
	if err != nil {
		return 0, fmt.Errorf("error: cannot query approved: invalid identity: %s", priv)
	}
	amount, err := contract.userContract.Allowance(&bind.CallOpts{
		From: common.HexToAddress(contract.wallet[1]),
	}, addr, contract.dataAddress)
	if err != nil {
		return 0, fmt.Errorf("error: cannot query approved:  %v", err)
	}
	return amount.Uint64(), nil
}

func (contract *smartcontract) TransferFrom(priv string, amount uint64) (string, error) {
	auth, addr, err := contract.parseIdentity(priv)
	if err != nil {
		return "", fmt.Errorf("error: cannot transfer from Filesharing Contract: invalid identity: %s", priv)
	}
	total := new(big.Int).SetUint64(amount)
	tx, err := contract.userContract.TransferFrom(auth, contract.dataAddress, addr, total)
	if err != nil {
		return "", fmt.Errorf("error: cannot transfer from Filesharing Contract:  %v", err)
	}
	return tx.To().Hex(), nil
}

func token2User(user *token.ForForTokenUser) entity.User {
	user.Addr.Hex()
	b := entity.User{
		Address: user.Addr.Hex(),
		Balance: user.Balance.Uint64(),
	}
	for _, asset := range user.Assets {
		a := entity.Asset{
			FileID:     asset.FileID,
			Permission: asset.Permission,
		}
		b.Assets = append(b.Assets, a)
	}
	return b
}
