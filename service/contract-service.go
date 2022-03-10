package service

import (
	"dfa/entity"
	"dfa/smartcontract"
	"fmt"
)

type ContractService interface {
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

func (service *contractService) DownloadFile() {

}
