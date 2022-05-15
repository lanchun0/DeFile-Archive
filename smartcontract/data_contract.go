package smartcontract

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract/filesharing"

	"github.com/ethereum/go-ethereum/common"

	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (contract *smartcontract) CreateFile(priv string, data entity.Data) (string, error) {
	if len(data.ID) == 0 {
		return "", fmt.Errorf("file id cannot be nil")
	}
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	price := new(big.Int).SetUint64(data.Price)
	tx, err := contract.dataContract.CreateFile(auth, data.ID, data.PermissionLevel, data.Tradable, price)
	if err != nil {
		err = fmt.Errorf("failed tx: %v", err)
	} else {
		fmt.Println("transaction: ", tx.Hash().Hex())
	}
	return tx.Hash().Hex(), err
}

// To read a file, sign with the id of that file
func (contract *smartcontract) ReadFile(priv, id string) (string, entity.Data, error) {
	_, addr, err := contract.parseIdentity(priv)
	if err != nil {
		return "", entity.Data{}, fmt.Errorf("failed to read the file: %v", err)
	}
	opts := &bind.CallOpts{From: addr}
	ipfsHash, f, err := contract.dataContract.ReadFile(opts, id)
	if err != nil {
		return ipfsHash, entity.Data{}, fmt.Errorf("failed to read file: %v", err)
	}
	data := offFile2Data(&f)
	return ipfsHash, data, nil
}

// To write a file, sign the hash digest
func (contract *smartcontract) WriteFile(priv, id, ipfsHash string, data entity.MeteData) (string, error) {
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return "", fmt.Errorf("failed to write the file: %v", err)
	}
	name, digest := data.FileName, data.HashDigest
	time, size := general.Timestamp2Str(data.TimeStamp), new(big.Int).SetUint64(data.Size)
	tx, err := contract.dataContract.WriteFile(auth, id, digest, name, size, time, ipfsHash)
	if err != nil {
		return "", fmt.Errorf("failed to write the file: %v", err)
	}
	fmt.Println("tx: ", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

// To share a file
func (contract *smartcontract) ShareFile(priv, to, id string, pL entity.Permission) (string, error) {
	// plaintext := general.String2Bytes(to + id)
	// valid := general.VerifySignature(from, signature, plaintext)
	// if !valid {
	// 	return "", fmt.Errorf("invalid sharing")
	// }
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return "", fmt.Errorf("failed to share the file: %v", err)
	}
	toAddr := common.HexToAddress(to)
	tx, err := contract.dataContract.Authorize(auth, toAddr, id, uint8(pL))
	if err != nil {
		return "", fmt.Errorf("failed to share the file: %v", err)
	}

	return tx.Hash().Hex(), nil
}

func (contract *smartcontract) PurchaseFile(priv, id string) (string, bool, error) {
	auth, _, err := contract.parseIdentity(priv)
	if err != nil {
		return "", false, fmt.Errorf("failed to purchase the file: %v", err)
	}
	tx, err := contract.dataContract.TradeFile(auth, id)
	if err != nil {
		return "", false, fmt.Errorf("failed to purchase the file: %v", err)
	}
	return tx.Hash().Hex(), true, nil
}

func (contract *smartcontract) QueryFile(id string) (entity.Data, error) {
	f, err := contract.dataContract.QueryFile(nil, id)
	if err != nil {
		return entity.Data{}, fmt.Errorf("failed to query: %v", err)
	}
	data := offFile2Data(&f)
	return data, nil
}

func (contract *smartcontract) QueryAllFiles() ([]entity.Data, error) {
	total, err := contract.dataContract.GetFileCount(nil)
	if err != nil {
		return []entity.Data{}, fmt.Errorf("failed to query: %v", err)
	}
	amount := total.Uint64()
	var i uint64
	res, i := []entity.Data{}, 0
	for i < amount {
		index := big.NewInt(int64(i))
		f, err := contract.dataContract.QueryFileByIndex(nil, index)
		if err != nil {
			return res, err
		}
		if f.PermissionLevel != uint8(entity.L_0) {
			data := offFile2Data(&f)
			res = append(res, data)
		}
		i++
	}
	return res, nil
}

func offFile2Data(f *filesharing.QueryFile) entity.Data {
	data := entity.Data{
		ID:              f.FileID,
		Owner:           f.Owner.Hex(),
		PermissionLevel: f.PermissionLevel,
		Tradable:        f.Tradable,
		Price:           f.Price.Uint64(),
		MeteData: entity.MeteData{
			FileName:   f.MeteData.FileName,
			HashDigest: f.MeteData.HashDigest,
			Signer:     f.MeteData.Signer.Hex(),
			Size:       f.MeteData.Size.Uint64(),
		},
		PermissionList: []entity.PermissionList{},
	}
	data.MeteData.TimeStamp, _ = general.Str2Timestamp(f.MeteData.Timestamp)
	for _, pl := range f.PermissionList {
		pl_new := entity.PermissionList{
			Address:    pl.User.Hex(),
			Permission: pl.Permission,
		}
		data.PermissionList = append(data.PermissionList, pl_new)
	}
	return data
}
