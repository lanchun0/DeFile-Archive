package smartcontract

import (
	"dfa/entity"
	"dfa/smartcontract/behavior"
	"dfa/smartcontract/filesharing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type smartcontract struct {
	behaviorContract *behavior.Behavior
	dataContract     *filesharing.Filesharing

	client      *ethclient.Client
	auth        *bind.TransactOpts
	serverIndex int
	servers     []string
}

// var Client *ethclient.Client
// var ServerIndex int = 0
var servers = [10]string{
	"0xf1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5",
	"0x91821f9af458d612362136648fc8552a47d8289c0f25a8a1bf0860510332cef9",
	"0xbb32062807c162a5243dc9bcf21d8114cb636c376596e1cf2895ec9e5e3e0a68",
	"0x95ce6122165d94aa51b0fcf51021895b39b0ff291aa640c803d5401bd87894d5",
	"0x3af93668029f95d526fc1d2bdefccc120bfe1d26a0462d268e8f6b2f71402ba3",
	"0x3b24a4fdf2e6e1375008c387c5456ce00cb0772435ae1938c2fe833103393b9a",
	"0xcba858feeb49e1ca8053a5213987a22c3ee83d9f9f396e138940a018dd837ebb",
	"0xdf48bfda4cb4b4e094803e923836a8538fbf607da79f6e46d68cdd43fb2f3f88",
	"0x487efb6249a8a4d45a19c8e5d1e5c7d3f6610a7e69f8f81ddcf368f9a0c0d6d5",
	"0xbb4cce73db59f456ea427e5862fdb0d5bc038a7d0b930cbb45e1c4f6d122289e",
}

type SmartContract interface {
	DeployContract() (err error)

	Register() (entity.Behavior, string, error)
	Login(pub, signature string) (entity.Behavior, error)

	CreateFile(data entity.Data) (string, error)
	ReadFile(pub, id, signature string) (entity.Data, error)
	WriteFile(pub string, data entity.Data) (string, error)
	ShareFile(from, to, id string, pL entity.Permission) (string, error)

	QueryFile(id string) (entity.Data, error)
	QueryAllFiles() ([]entity.Data, error)

	changeAuth()
}

func NewConract() SmartContract {
	return &smartcontract{
		serverIndex: 0,
		servers:     servers[:],
	}
}
