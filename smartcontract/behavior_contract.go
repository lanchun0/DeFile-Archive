package smartcontract

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract/token"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
