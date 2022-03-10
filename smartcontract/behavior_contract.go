package smartcontract

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract/behavior"
	"fmt"
)

func (contract *smartcontract) Register() (entity.Behavior, string, error) {
	contract.changeAuth()
	pub, priv, addr := general.GenerateECSDAKey()
	fmt.Println("Generated new user, address: ", addr)
	tx, err := contract.behaviorContract.CreateUser(contract.auth, pub, priv)
	if err != nil {
		return entity.Behavior{}, "", err
	}
	fmt.Println("Transaction: ", tx.Hash().Hex())

	user, err := contract.behaviorContract.GetUser(nil, pub)
	if err != nil {
		return entity.Behavior{}, "", err
	}
	b := user2Behavior(&user)
	return b, tx.Hash().Hex(), nil
}

// sign the public key to login
func (contract *smartcontract) Login(pub, signature string) (entity.Behavior, error) {
	pubBytes := general.String2Bytes(pub)
	if !general.VerifySignature(pub, signature, pubBytes) {
		return entity.Behavior{}, fmt.Errorf("invalid signature")
	}
	user, err := contract.behaviorContract.GetUser(nil, pub)
	if err != nil {
		return entity.Behavior{}, fmt.Errorf("failed to query on behavior chain: %v", err)
	}
	b := user2Behavior(&user)
	return b, nil
}

func user2Behavior(user *behavior.User) entity.Behavior {
	b := entity.Behavior{
		PrivateKey: user.PrivateKey,
		PublicKey:  user.PublicKey,
	}
	for _, asset := range user.Assets {
		a := entity.Asset{
			FileID:     asset.Fileid,
			Permission: asset.Permission,
		}
		b.Assets = append(b.Assets, a)
	}
	return b
}
