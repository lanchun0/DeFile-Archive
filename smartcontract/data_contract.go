package smartcontract

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract/filesharing"

	"fmt"
	"math/big"
)

func (contract *smartcontract) CreateFile(data entity.Data) (string, error) {
	hashBytes := general.String2Bytes(data.HashDigest)
	ok := general.VerifySignature(data.Owner, data.Signature, hashBytes)
	if !ok {
		return "", fmt.Errorf("invalid signature")
	}
	contract.changeAuth()
	tx1, err := contract.dataContract.CreateFile(contract.auth, data.ID, data.HashDigest, data.Owner, data.PermissionLevel)
	if err != nil {
		return "", fmt.Errorf("failed to record the file on chain: %v", err)
	}
	contract.changeAuth()
	tx2, err := contract.dataContract.WriteFile(contract.auth, data.Owner, data.ID, data.HashDigest, data.Signature)
	if err != nil {
		return "", fmt.Errorf("failed to record the file on chain: %v", err)
	}
	contract.changeAuth()
	timestamp, size := general.Timestamp2Str(data.MeteData.TimeStamp), big.NewInt(int64(data.MeteData.Size))
	tx3, err := contract.dataContract.WriteMeta(contract.auth, data.Owner, data.ID, data.MeteData.FileName, size, timestamp)
	if err != nil {
		return "", fmt.Errorf("failed to record the file on chain: %v", err)
	}
	tx := "transaction 1: " + tx1.Hash().Hex() + "\ntrasaction 2: " + tx2.Hash().Hex() + "\ntrasaction 3: " + tx3.Hash().Hex()
	return tx, nil
}

// To read a file, sign with the id of that file
func (contract *smartcontract) ReadFile(pub, id, signature string) (entity.Data, error) {
	if !general.VerifySignature(pub, signature, general.String2Bytes(id)) {
		return entity.Data{}, fmt.Errorf("invalid reading identity")
	}
	can, err := contract.dataContract.CanRead(nil, pub, id)
	if !can || err != nil {
		return entity.Data{}, fmt.Errorf("no authority to read the file")
	}
	f, err := contract.dataContract.QueryFile(nil, pub)
	if err != nil {
		return entity.Data{}, fmt.Errorf("cannot read the file")
	}
	data := offFile2Data(&f)
	return data, nil
}

// To write a file, sign the hash digest
func (contract *smartcontract) WriteFile(pub string, data entity.Data) (string, error) {
	if !general.VerifySignature(pub, data.Signature, general.String2Bytes(data.HashDigest)) {
		return "", fmt.Errorf("invalid reading identity")
	}
	can, err := contract.dataContract.CanWrite(nil, pub, data.ID)
	if !can || err != nil {
		return "", fmt.Errorf("no authority to write the file")
	}
	contract.changeAuth()
	tx1, err := contract.dataContract.WriteFile(nil, pub, data.ID, data.HashDigest, data.Signature)
	if err != nil {
		return "", fmt.Errorf("failed to write the file")
	}
	contract.changeAuth()
	size, timestamp := big.NewInt(int64(data.MeteData.Size)), general.Timestamp2Str(data.MeteData.TimeStamp)
	tx2, err := contract.dataContract.WriteMeta(nil, pub, data.ID, data.MeteData.FileName, size, timestamp)
	if err != nil {
		return "", fmt.Errorf("failed to write the mete data")
	}
	tx := "transaction 1:" + tx1.Hash().Hex() + "transaction 2:" + tx2.Hash().Hex()
	return tx, nil
}

// To share a file, sign the to+id
func (contract *smartcontract) ShareFile(from, to, id, signature string, pL entity.Permission) (string, error) {
	plaintext := general.String2Bytes(to + id)
	valid := general.VerifySignature(from, signature, plaintext)
	if !valid {
		return "", fmt.Errorf("invalid sharing")
	}
	contract.changeAuth()
	tx, err := contract.dataContract.AddPermission(contract.auth, from, to, id, uint8(pL))
	if err != nil {
		return "", fmt.Errorf("cannot share: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func (contract *smartcontract) QueryFile(id string) (entity.Data, error) {
	f, err := contract.dataContract.QueryFile(nil, id)
	if err != nil {
		return entity.Data{}, fmt.Errorf("failed to query: %v", err)
	}
	if f.PermissionLevel == uint8(entity.L_0) {
		return entity.Data{}, fmt.Errorf("failed to query: file is private")
	}
	data := offFile2Data(&f)
	return data, nil
}

func (contract *smartcontract) QueryAllFiles() ([]entity.Data, error) {
	total, err := contract.dataContract.GetFileCount(nil)
	if err != nil {
		return []entity.Data{}, fmt.Errorf("failed to query: %v", err)
	}
	var i uint64
	res, i := []entity.Data{}, 0
	for i < total {
		f, err := contract.dataContract.QueryFileByIndex(nil, i)
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

func offFile2Data(f *filesharing.OffFile) entity.Data {
	data := entity.Data{
		ID:              f.FileID,
		HashDigest:      f.HashDigest,
		Owner:           f.Owner,
		PermissionLevel: f.PermissionLevel,
		Signature:       f.Signature,
		MeteData: entity.MeteData{
			FileName: f.MeteData.FileName,
		},
		PermissionList: []entity.PermissionList{},
	}
	size := f.MeteData.Size.Int64()
	data.MeteData.TimeStamp, _ = general.Str2Timestamp(f.MeteData.Timestamp)
	data.MeteData.Size = int(size)
	for _, pl := range f.PermissionList {
		pl_new := entity.PermissionList{
			PublicKey:  pl.PublicKey,
			Permission: pl.Permission,
		}
		data.PermissionList = append(data.PermissionList, pl_new)
	}
	return data
}
