package smartcontract

import (
	"dfa/entity"
	"dfa/smartcontract/filesharing"
	"dfa/smartcontract/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type smartcontract struct {
	userContract *token.Token
	dataContract *filesharing.Filesharing
	userAddress  common.Address
	dataAddress  common.Address

	client *ethclient.Client
	// auth        *bind.TransactOpts
	servers []string
}

// type identity struct {
// 	client *ethclient.Client
// 	auth   *bind.TransactOpts
// }

// var Client *ethclient.Client
// var ServerIndex int = 0
var servers = [2]string{
	"0xf1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5",
	"0x91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9",
}

type SmartContract interface {
	DeployContract() (err error)

	Register() (entity.User, string, error)
	Login(priv string) (entity.User, error)

	CreateFile(priv string, data entity.Data) (tx string, err error)
	ReadFile(priv, id string) (entity.Data, error)
	WriteFile(priv string, data entity.Data) (string, error)
	ShareFile(priv, to, id string, pL entity.Permission) (string, error)

	QueryFile(id string) (entity.Data, error)
	QueryAllFiles() ([]entity.Data, error)

	GetDataContractAddress() string
	GetUserContractAddress() string
	parseIdentity(priv string) (*bind.TransactOpts, common.Address, error)
	// changeAuth()
}

func NewConract() SmartContract {
	return &smartcontract{
		servers: servers[:],
	}
}
