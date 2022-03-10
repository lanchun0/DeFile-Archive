// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package behavior

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Asset is an auto generated low-level Go binding around an user-defined struct.
type Asset struct {
	Fileid     string
	Permission uint8
}

// User is an auto generated low-level Go binding around an user-defined struct.
type User struct {
	PublicKey  string
	PrivateKey string
	Assets     []Asset
	Exist      bool
}

// BehaviorMetaData contains all meta data concerning the Behavior contract.
var BehaviorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_permission\",\"type\":\"uint8\"}],\"name\":\"addAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_permission\",\"type\":\"uint8\"}],\"name\":\"addFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_permission\",\"type\":\"uint8\"}],\"name\":\"authorFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_priv\",\"type\":\"string\"}],\"name\":\"createUser\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"}],\"name\":\"existBehavior\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"}],\"name\":\"existUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"privateKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileid\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"permission\",\"type\":\"uint8\"}],\"internalType\":\"structAsset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"internalType\":\"structUser\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"}],\"name\":\"hasAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"}],\"name\":\"hasFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_permission\",\"type\":\"uint8\"}],\"name\":\"transferAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611879806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806387b162b51161007157806387b162b51461019e5780639dbe87a1146101ce578063abb7acf5146101ff578063b5cb15f71461022f578063b6623e4b1461024d578063bace2caf1461027d576100a9565b80632d34f01f146100ae57806331feb671146100de57806335e1f89d1461010e5780635ba4d07e1461013e578063706ef6ee1461016e575b600080fd5b6100c860048036038101906100c391906110a4565b6102ae565b6040516100d59190611406565b60405180910390f35b6100f860048036038101906100f39190610f0d565b6102c8565b60405161010591906114ca565b60405180910390f35b61012860048036038101906101239190610f0d565b61059f565b6040516101359190611406565b60405180910390f35b61015860048036038101906101539190610fdb565b6105da565b6040516101659190611406565b60405180910390f35b61018860048036038101906101839190610f0d565b610886565b6040516101959190611406565b60405180910390f35b6101b860048036038101906101b39190610f5a565b61089a565b6040516101c591906114ec565b60405180910390f35b6101e860048036038101906101e39190610f5a565b610ad5565b6040516101f6929190611421565b60405180910390f35b610219600480360381019061021491906110a4565b610bf7565b6040516102269190611406565b60405180910390f35b610237610d7d565b60405161024491906114ec565b60405180910390f35b61026760048036038101906102629190610fdb565b610d9b565b6040516102749190611406565b60405180910390f35b61029760048036038101906102929190610f5a565b610db9565b6040516102a5929190611421565b60405180910390f35b60006102bd8686868686610bf7565b905095945050505050565b6102d0610dd5565b6102da838361059f565b610319576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103109061146a565b60405180910390fd5b60006001848460405161032d9291906113d6565b908152602001604051809103902060405180608001604052908160008201805461035690611625565b80601f016020809104026020016040519081016040528092919081815260200182805461038290611625565b80156103cf5780601f106103a4576101008083540402835291602001916103cf565b820191906000526020600020905b8154815290600101906020018083116103b257829003601f168201915b505050505081526020016001820180546103e890611625565b80601f016020809104026020016040519081016040528092919081815260200182805461041490611625565b80156104615780601f1061043657610100808354040283529160200191610461565b820191906000526020600020905b81548152906001019060200180831161044457829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b8282101561057057838290600052602060002090600202016040518060400160405290816000820180546104c290611625565b80601f01602080910402602001604051908101604052809291908181526020018280546104ee90611625565b801561053b5780601f106105105761010080835404028352916020019161053b565b820191906000526020600020905b81548152906001019060200180831161051e57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1660ff1660ff16815250508152602001906001019061048f565b5050505081526020016003820160009054906101000a900460ff16151515158152505090508091505092915050565b6000600183836040516105b39291906113d6565b908152602001604051809103902060030160009054906101000a900460ff16905092915050565b6000806000905060008060006105f28c8c8a8a610ad5565b80935081955050506106068a8a8a8a610ad5565b80925081945050506106188a8a61059f565b1580610622575083155b8061063257508560ff168260ff16105b80610640575060408260ff16105b8061065157508560ff168160ff1610155b1561066357600094505050505061087b565b82610693576106758a8a8a8a8a610bf7565b61068657600094505050505061087b565b600194505050505061087b565b60005b60018b8b6040516106a89291906113d6565b90815260200160405180910390206002018054905081101561087157600060018c8c6040516106d89291906113d6565b908152602001604051809103902060020182815481106106fb576106fa61172f565b5b9060005260206000209060020201600001805461071790611625565b80601f016020809104026020016040519081016040528092919081815260200182805461074390611625565b80156107905780601f1061076557610100808354040283529160200191610790565b820191906000526020600020905b81548152906001019060200180831161077357829003601f168201915b5050505050905089896040516020016107aa9291906113d6565b60405160208183030381529060405280519060200120816040516020016107d191906113ef565b60405160208183030381529060405280519060200120141561085d5782881760018d8d6040516108029291906113d6565b908152602001604051809103902060020183815481106108255761082461172f565b5b906000526020600020906002020160010160006101000a81548160ff021916908360ff1602179055506001965050505050505061087b565b50808061086990611657565b915050610696565b5060009450505050505b979650505050505050565b6000610892838361059f565b905092915050565b60008085856040516020016108b09291906113d6565b604051602081830303815290604052511415610901576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f89061144a565b60405180910390fd5b600083836040516020016109169291906113d6565b604051602081830303815290604052511415610967576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095e9061148a565b60405180910390fd5b610971858561059f565b156109b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a8906114aa565b60405180910390fd5b6002600081819054906101000a900467ffffffffffffffff16809291906109d7906116a0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484600080600260009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000209190610a49929190610dff565b50600060018686604051610a5e9291906113d6565b908152602001604051809103902090508585826000019190610a81929190610dff565b508383826001019190610a95929190610dff565b5060018160030160006101000a81548160ff021916908315150217905550600260009054906101000a900467ffffffffffffffff16915050949350505050565b6000806000848490501480610af15750610aef868661059f565b155b15610b025760008091509150610bee565b6000610b0e87876102c8565b905060008060005b836040015151811015610be357600084604001518281518110610b3c57610b3b61172f565b5b60200260200101516000015190508888604051602001610b5d9291906113d6565b6040516020818303038152906040528051906020012081604051602001610b8491906113ef565b604051602081830303815290604052805190602001201415610bcf576001935084604001518281518110610bbb57610bba61172f565b5b602002602001015160200151925050610be3565b508080610bdb90611657565b915050610b16565b508181945094505050505b94509492505050565b6000610c03868661059f565b1580610c125750600084849050145b15610c205760009050610d74565b6000610c2e87878787610ad5565b50809150508015610c43576000915050610d74565b60018787604051610c559291906113d6565b90815260200160405180910390206002016001816001815401808255809150500390600052602060002090505060006001808989604051610c979291906113d6565b908152602001604051809103902060020180549050610cb69190611578565b9050858560018a8a604051610ccc9291906113d6565b90815260200160405180910390206002018381548110610cef57610cee61172f565b5b90600052602060002090600202016000019190610d0d929190610dff565b508360018989604051610d219291906113d6565b90815260200160405180910390206002018281548110610d4457610d4361172f565b5b906000526020600020906002020160010160006101000a81548160ff021916908360ff1602179055506001925050505b95945050505050565b6000600260009054906101000a900467ffffffffffffffff16905090565b6000610dac888888888888886105da565b9050979650505050505050565b600080610dc886868686610ad5565b9150915094509492505050565b60405180608001604052806060815260200160608152602001606081526020016000151581525090565b828054610e0b90611625565b90600052602060002090601f016020900481019282610e2d5760008555610e74565b82601f10610e4657803560ff1916838001178555610e74565b82800160010185558215610e74579182015b82811115610e73578235825591602001919060010190610e58565b5b509050610e819190610e85565b5090565b5b80821115610e9e576000816000905550600101610e86565b5090565b60008083601f840112610eb857610eb7611763565b5b8235905067ffffffffffffffff811115610ed557610ed461175e565b5b602083019150836001820283011115610ef157610ef0611768565b5b9250929050565b600081359050610f078161182c565b92915050565b60008060208385031215610f2457610f23611772565b5b600083013567ffffffffffffffff811115610f4257610f4161176d565b5b610f4e85828601610ea2565b92509250509250929050565b60008060008060408587031215610f7457610f73611772565b5b600085013567ffffffffffffffff811115610f9257610f9161176d565b5b610f9e87828801610ea2565b9450945050602085013567ffffffffffffffff811115610fc157610fc061176d565b5b610fcd87828801610ea2565b925092505092959194509250565b60008060008060008060006080888a031215610ffa57610ff9611772565b5b600088013567ffffffffffffffff8111156110185761101761176d565b5b6110248a828b01610ea2565b9750975050602088013567ffffffffffffffff8111156110475761104661176d565b5b6110538a828b01610ea2565b9550955050604088013567ffffffffffffffff8111156110765761107561176d565b5b6110828a828b01610ea2565b935093505060606110958a828b01610ef8565b91505092959891949750929550565b6000806000806000606086880312156110c0576110bf611772565b5b600086013567ffffffffffffffff8111156110de576110dd61176d565b5b6110ea88828901610ea2565b9550955050602086013567ffffffffffffffff81111561110d5761110c61176d565b5b61111988828901610ea2565b9350935050604061112c88828901610ef8565b9150509295509295909350565b600061114583836112fb565b905092915050565b600061115882611517565b611162818561153a565b93508360208202850161117485611507565b8060005b858110156111b057848403895281516111918582611139565b945061119c8361152d565b925060208a01995050600181019050611178565b50829750879550505050505092915050565b6111cb816115ac565b82525050565b6111da816115ac565b82525050565b60006111ec838561156d565b93506111f98385846115e3565b82840190509392505050565b600061121082611522565b61121a818561154b565b935061122a8185602086016115f2565b61123381611777565b840191505092915050565b600061124982611522565b611253818561156d565b93506112638185602086016115f2565b80840191505092915050565b600061127c60188361155c565b915061128782611788565b602082019050919050565b600061129f60138361155c565b91506112aa826117b1565b602082019050919050565b60006112c260198361155c565b91506112cd826117da565b602082019050919050565b60006112e560138361155c565b91506112f082611803565b602082019050919050565b600060408301600083015184820360008601526113188282611205565b915050602083015161132d60208601826113b8565b508091505092915050565b600060808301600083015184820360008601526113558282611205565b9150506020830151848203602086015261136f8282611205565b91505060408301518482036040860152611389828261114d565b915050606083015161139e60608601826111c2565b508091505092915050565b6113b2816115c2565b82525050565b6113c1816115d6565b82525050565b6113d0816115d6565b82525050565b60006113e38284866111e0565b91508190509392505050565b60006113fb828461123e565b915081905092915050565b600060208201905061141b60008301846111d1565b92915050565b600060408201905061143660008301856111d1565b61144360208301846113c7565b9392505050565b600060208201905081810360008301526114638161126f565b9050919050565b6000602082019050818103600083015261148381611292565b9050919050565b600060208201905081810360008301526114a3816112b5565b9050919050565b600060208201905081810360008301526114c3816112d8565b9050919050565b600060208201905081810360008301526114e48184611338565b905092915050565b600060208201905061150160008301846113a9565b92915050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611583826115b8565b915061158e836115b8565b9250828210156115a1576115a06116d1565b5b828203905092915050565b60008115159050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156116105780820151818401526020810190506115f5565b8381111561161f576000848401525b50505050565b6000600282049050600182168061163d57607f821691505b6020821081141561165157611650611700565b5b50919050565b6000611662826115b8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611695576116946116d1565b5b600182019050919050565b60006116ab826115c2565b915067ffffffffffffffff8214156116c6576116c56116d1565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f7075626c6963206b65792063616e6e6f74206265206e696c0000000000000000600082015250565b7f7573657220646f6573206e6f7420657869737400000000000000000000000000600082015250565b7f70726976617465206b65792063616e6e6f74206265206e696c00000000000000600082015250565b7f7573657220616c72656164792065786973747300000000000000000000000000600082015250565b611835816115d6565b811461184057600080fd5b5056fea26469706673582212200737cd4d820731a2517777f8c836067cddae4cf6332450f841edc57489c7536064736f6c63430008070033",
}

// BehaviorABI is the input ABI used to generate the binding from.
// Deprecated: Use BehaviorMetaData.ABI instead.
var BehaviorABI = BehaviorMetaData.ABI

// BehaviorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BehaviorMetaData.Bin instead.
var BehaviorBin = BehaviorMetaData.Bin

// DeployBehavior deploys a new Ethereum contract, binding an instance of Behavior to it.
func DeployBehavior(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Behavior, error) {
	parsed, err := BehaviorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BehaviorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Behavior{BehaviorCaller: BehaviorCaller{contract: contract}, BehaviorTransactor: BehaviorTransactor{contract: contract}, BehaviorFilterer: BehaviorFilterer{contract: contract}}, nil
}

// Behavior is an auto generated Go binding around an Ethereum contract.
type Behavior struct {
	BehaviorCaller     // Read-only binding to the contract
	BehaviorTransactor // Write-only binding to the contract
	BehaviorFilterer   // Log filterer for contract events
}

// BehaviorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BehaviorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BehaviorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BehaviorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BehaviorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BehaviorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BehaviorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BehaviorSession struct {
	Contract     *Behavior         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BehaviorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BehaviorCallerSession struct {
	Contract *BehaviorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BehaviorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BehaviorTransactorSession struct {
	Contract     *BehaviorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BehaviorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BehaviorRaw struct {
	Contract *Behavior // Generic contract binding to access the raw methods on
}

// BehaviorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BehaviorCallerRaw struct {
	Contract *BehaviorCaller // Generic read-only contract binding to access the raw methods on
}

// BehaviorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BehaviorTransactorRaw struct {
	Contract *BehaviorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBehavior creates a new instance of Behavior, bound to a specific deployed contract.
func NewBehavior(address common.Address, backend bind.ContractBackend) (*Behavior, error) {
	contract, err := bindBehavior(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Behavior{BehaviorCaller: BehaviorCaller{contract: contract}, BehaviorTransactor: BehaviorTransactor{contract: contract}, BehaviorFilterer: BehaviorFilterer{contract: contract}}, nil
}

// NewBehaviorCaller creates a new read-only instance of Behavior, bound to a specific deployed contract.
func NewBehaviorCaller(address common.Address, caller bind.ContractCaller) (*BehaviorCaller, error) {
	contract, err := bindBehavior(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BehaviorCaller{contract: contract}, nil
}

// NewBehaviorTransactor creates a new write-only instance of Behavior, bound to a specific deployed contract.
func NewBehaviorTransactor(address common.Address, transactor bind.ContractTransactor) (*BehaviorTransactor, error) {
	contract, err := bindBehavior(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BehaviorTransactor{contract: contract}, nil
}

// NewBehaviorFilterer creates a new log filterer instance of Behavior, bound to a specific deployed contract.
func NewBehaviorFilterer(address common.Address, filterer bind.ContractFilterer) (*BehaviorFilterer, error) {
	contract, err := bindBehavior(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BehaviorFilterer{contract: contract}, nil
}

// bindBehavior binds a generic wrapper to an already deployed contract.
func bindBehavior(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BehaviorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Behavior *BehaviorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Behavior.Contract.BehaviorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Behavior *BehaviorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Behavior.Contract.BehaviorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Behavior *BehaviorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Behavior.Contract.BehaviorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Behavior *BehaviorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Behavior.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Behavior *BehaviorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Behavior.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Behavior *BehaviorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Behavior.Contract.contract.Transact(opts, method, params...)
}

// ExistBehavior is a free data retrieval call binding the contract method 0x706ef6ee.
//
// Solidity: function existBehavior(string _pub) view returns(bool)
func (_Behavior *BehaviorCaller) ExistBehavior(opts *bind.CallOpts, _pub string) (bool, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "existBehavior", _pub)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistBehavior is a free data retrieval call binding the contract method 0x706ef6ee.
//
// Solidity: function existBehavior(string _pub) view returns(bool)
func (_Behavior *BehaviorSession) ExistBehavior(_pub string) (bool, error) {
	return _Behavior.Contract.ExistBehavior(&_Behavior.CallOpts, _pub)
}

// ExistBehavior is a free data retrieval call binding the contract method 0x706ef6ee.
//
// Solidity: function existBehavior(string _pub) view returns(bool)
func (_Behavior *BehaviorCallerSession) ExistBehavior(_pub string) (bool, error) {
	return _Behavior.Contract.ExistBehavior(&_Behavior.CallOpts, _pub)
}

// ExistUser is a free data retrieval call binding the contract method 0x35e1f89d.
//
// Solidity: function existUser(string _pub) view returns(bool)
func (_Behavior *BehaviorCaller) ExistUser(opts *bind.CallOpts, _pub string) (bool, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "existUser", _pub)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistUser is a free data retrieval call binding the contract method 0x35e1f89d.
//
// Solidity: function existUser(string _pub) view returns(bool)
func (_Behavior *BehaviorSession) ExistUser(_pub string) (bool, error) {
	return _Behavior.Contract.ExistUser(&_Behavior.CallOpts, _pub)
}

// ExistUser is a free data retrieval call binding the contract method 0x35e1f89d.
//
// Solidity: function existUser(string _pub) view returns(bool)
func (_Behavior *BehaviorCallerSession) ExistUser(_pub string) (bool, error) {
	return _Behavior.Contract.ExistUser(&_Behavior.CallOpts, _pub)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _pub) view returns((string,string,(string,uint8)[],bool))
func (_Behavior *BehaviorCaller) GetUser(opts *bind.CallOpts, _pub string) (User, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "getUser", _pub)

	if err != nil {
		return *new(User), err
	}

	out0 := *abi.ConvertType(out[0], new(User)).(*User)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _pub) view returns((string,string,(string,uint8)[],bool))
func (_Behavior *BehaviorSession) GetUser(_pub string) (User, error) {
	return _Behavior.Contract.GetUser(&_Behavior.CallOpts, _pub)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _pub) view returns((string,string,(string,uint8)[],bool))
func (_Behavior *BehaviorCallerSession) GetUser(_pub string) (User, error) {
	return _Behavior.Contract.GetUser(&_Behavior.CallOpts, _pub)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint64)
func (_Behavior *BehaviorCaller) GetUserCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "getUserCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint64)
func (_Behavior *BehaviorSession) GetUserCount() (uint64, error) {
	return _Behavior.Contract.GetUserCount(&_Behavior.CallOpts)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint64)
func (_Behavior *BehaviorCallerSession) GetUserCount() (uint64, error) {
	return _Behavior.Contract.GetUserCount(&_Behavior.CallOpts)
}

// HasAsset is a free data retrieval call binding the contract method 0x9dbe87a1.
//
// Solidity: function hasAsset(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorCaller) HasAsset(opts *bind.CallOpts, _pub string, _fileID string) (bool, uint8, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "hasAsset", _pub, _fileID)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// HasAsset is a free data retrieval call binding the contract method 0x9dbe87a1.
//
// Solidity: function hasAsset(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorSession) HasAsset(_pub string, _fileID string) (bool, uint8, error) {
	return _Behavior.Contract.HasAsset(&_Behavior.CallOpts, _pub, _fileID)
}

// HasAsset is a free data retrieval call binding the contract method 0x9dbe87a1.
//
// Solidity: function hasAsset(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorCallerSession) HasAsset(_pub string, _fileID string) (bool, uint8, error) {
	return _Behavior.Contract.HasAsset(&_Behavior.CallOpts, _pub, _fileID)
}

// HasFile is a free data retrieval call binding the contract method 0xbace2caf.
//
// Solidity: function hasFile(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorCaller) HasFile(opts *bind.CallOpts, _pub string, _fileID string) (bool, uint8, error) {
	var out []interface{}
	err := _Behavior.contract.Call(opts, &out, "hasFile", _pub, _fileID)

	if err != nil {
		return *new(bool), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// HasFile is a free data retrieval call binding the contract method 0xbace2caf.
//
// Solidity: function hasFile(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorSession) HasFile(_pub string, _fileID string) (bool, uint8, error) {
	return _Behavior.Contract.HasFile(&_Behavior.CallOpts, _pub, _fileID)
}

// HasFile is a free data retrieval call binding the contract method 0xbace2caf.
//
// Solidity: function hasFile(string _pub, string _fileID) view returns(bool, uint8)
func (_Behavior *BehaviorCallerSession) HasFile(_pub string, _fileID string) (bool, uint8, error) {
	return _Behavior.Contract.HasFile(&_Behavior.CallOpts, _pub, _fileID)
}

// AddAsset is a paid mutator transaction binding the contract method 0xabb7acf5.
//
// Solidity: function addAsset(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactor) AddAsset(opts *bind.TransactOpts, _pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.contract.Transact(opts, "addAsset", _pub, _fileID, _permission)
}

// AddAsset is a paid mutator transaction binding the contract method 0xabb7acf5.
//
// Solidity: function addAsset(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorSession) AddAsset(_pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AddAsset(&_Behavior.TransactOpts, _pub, _fileID, _permission)
}

// AddAsset is a paid mutator transaction binding the contract method 0xabb7acf5.
//
// Solidity: function addAsset(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactorSession) AddAsset(_pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AddAsset(&_Behavior.TransactOpts, _pub, _fileID, _permission)
}

// AddFile is a paid mutator transaction binding the contract method 0x2d34f01f.
//
// Solidity: function addFile(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactor) AddFile(opts *bind.TransactOpts, _pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.contract.Transact(opts, "addFile", _pub, _fileID, _permission)
}

// AddFile is a paid mutator transaction binding the contract method 0x2d34f01f.
//
// Solidity: function addFile(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorSession) AddFile(_pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AddFile(&_Behavior.TransactOpts, _pub, _fileID, _permission)
}

// AddFile is a paid mutator transaction binding the contract method 0x2d34f01f.
//
// Solidity: function addFile(string _pub, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactorSession) AddFile(_pub string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AddFile(&_Behavior.TransactOpts, _pub, _fileID, _permission)
}

// AuthorFile is a paid mutator transaction binding the contract method 0xb6623e4b.
//
// Solidity: function authorFile(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactor) AuthorFile(opts *bind.TransactOpts, _from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.contract.Transact(opts, "authorFile", _from, _to, _fileID, _permission)
}

// AuthorFile is a paid mutator transaction binding the contract method 0xb6623e4b.
//
// Solidity: function authorFile(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorSession) AuthorFile(_from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AuthorFile(&_Behavior.TransactOpts, _from, _to, _fileID, _permission)
}

// AuthorFile is a paid mutator transaction binding the contract method 0xb6623e4b.
//
// Solidity: function authorFile(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactorSession) AuthorFile(_from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.AuthorFile(&_Behavior.TransactOpts, _from, _to, _fileID, _permission)
}

// CreateUser is a paid mutator transaction binding the contract method 0x87b162b5.
//
// Solidity: function createUser(string _pub, string _priv) returns(uint64)
func (_Behavior *BehaviorTransactor) CreateUser(opts *bind.TransactOpts, _pub string, _priv string) (*types.Transaction, error) {
	return _Behavior.contract.Transact(opts, "createUser", _pub, _priv)
}

// CreateUser is a paid mutator transaction binding the contract method 0x87b162b5.
//
// Solidity: function createUser(string _pub, string _priv) returns(uint64)
func (_Behavior *BehaviorSession) CreateUser(_pub string, _priv string) (*types.Transaction, error) {
	return _Behavior.Contract.CreateUser(&_Behavior.TransactOpts, _pub, _priv)
}

// CreateUser is a paid mutator transaction binding the contract method 0x87b162b5.
//
// Solidity: function createUser(string _pub, string _priv) returns(uint64)
func (_Behavior *BehaviorTransactorSession) CreateUser(_pub string, _priv string) (*types.Transaction, error) {
	return _Behavior.Contract.CreateUser(&_Behavior.TransactOpts, _pub, _priv)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5ba4d07e.
//
// Solidity: function transferAsset(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactor) TransferAsset(opts *bind.TransactOpts, _from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.contract.Transact(opts, "transferAsset", _from, _to, _fileID, _permission)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5ba4d07e.
//
// Solidity: function transferAsset(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorSession) TransferAsset(_from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.TransferAsset(&_Behavior.TransactOpts, _from, _to, _fileID, _permission)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5ba4d07e.
//
// Solidity: function transferAsset(string _from, string _to, string _fileID, uint8 _permission) returns(bool)
func (_Behavior *BehaviorTransactorSession) TransferAsset(_from string, _to string, _fileID string, _permission uint8) (*types.Transaction, error) {
	return _Behavior.Contract.TransferAsset(&_Behavior.TransactOpts, _from, _to, _fileID, _permission)
}
