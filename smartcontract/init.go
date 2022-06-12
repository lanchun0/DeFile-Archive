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
	wallet  []string
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

var wallet = [2]string{
	"0xE280029a7867BA5C9154434886c241775ea87e53",
	"0x68dB32D26d9529B2a142927c6f1af248fc6Ba7e9",
}

type SmartContract interface {
	DeployContract() (err error)

	TransferFrom(priv string, amount uint64) (tx string, err error)
	GetAllowance(priv string) (amount uint64, err error)
	GetApproved(priv string) (amount uint64, err error)
	Register() (entity.User, string, error)
	Login(priv string) (entity.User, error)
	Approve(priv string, price uint64) (success bool, tx string, err error)
	Topup(priv string, amount uint64) (success bool, tx string, err error)
	WithDraw(priv string, amount uint64) (success bool, tx string, err error)

	CreateFile(priv string, data entity.Data) (tx string, err error)
	ReadFile(priv, id string) (string, entity.Data, error)
	WriteFile(priv, id, ipfsHash string, data entity.MeteData) (string, error)
	ShareFile(priv, to, id string, pL entity.Permission) (string, error)
	PurchaseFile(priv, id string) (tx string, success bool, err error)

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
		wallet:  wallet[:],
	}
}
