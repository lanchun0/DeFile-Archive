package service

import (
	"dfa/entity"
	"dfa/smartcontract"
	"fmt"
	"strings"
)

type ContractService interface {
	Register() (entity.Behavior, string, error)
	Login(pub, signature string) (entity.Behavior, error)

	QueryAllFiles() []entity.Data
	QueryFile(id string) (entity.Data, error)
	DownloadFile(pub, id, signature string) (entity.Data, error)
	UploadFile(data entity.Data) (string, error)
	ShareFile(from, to, id, plStr string) (string, error)
	WriteFile(pub string, data entity.Data) (string, error)
}

type contractService struct {
	contract smartcontract.SmartContract
}

func NewContractService() ContractService {
	contract := smartcontract.NewConract()
	err := contract.DeployContract()
	if err != nil {
		panic(err)
	}
	return &contractService{
		contract: contract,
	}
}

func (service *contractService) Register() (entity.Behavior, string, error) {
	return service.contract.Register()
}

func (service *contractService) Login(pub, signature string) (entity.Behavior, error) {
	return service.contract.Login(pub, signature)
}

func (service *contractService) QueryAllFiles() []entity.Data {
	res, err := service.contract.QueryAllFiles()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func (service *contractService) QueryFile(id string) (entity.Data, error) {
	res, err := service.contract.QueryFile(id)
	if err != nil {
		fmt.Println(err)
		return entity.Data{}, fmt.Errorf("cannot query the file: %v", err)
	}
	return res, nil
}

// Query from blockcahin, and return the actual file if able to read
func (service *contractService) DownloadFile(pub, id, signature string) (entity.Data, error) {
	data, err := service.contract.ReadFile(pub, id, signature)
	if err != nil {
		return entity.Data{}, fmt.Errorf("cannot download the file: %v", err)
	}
	// fethch the file and return

	return data, nil
}

// return the transaction id if successful
func (service *contractService) UploadFile(data entity.Data) (string, error) {
	return service.contract.CreateFile(data)
}

func (service *contractService) ShareFile(from, to, id, plStr string) (string, error) {
	var pL entity.Permission
	if strings.EqualFold("owner", plStr) {
		pL = entity.Owner
	} else if strings.EqualFold("writer", plStr) {
		pL = entity.Writer
	} else if strings.EqualFold("administrator", plStr) {
		pL = entity.Administrator
	} else if strings.EqualFold("reader", plStr) {
		pL = entity.Reader
	} else {
		return "", fmt.Errorf("Invalid permission input: %s", plStr)
	}

	return service.contract.ShareFile(from, to, id, pL)
}

// modify both off-chain and on-chain data
// return transaction id if successful
func (service *contractService) WriteFile(pub string, data entity.Data) (string, error) {
	//off-chain

	// on-chain
	return service.contract.WriteFile(pub, data)
}
