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
	Bin: "0x608060405234801561001057600080fd5b5061166a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806387b162b51161007157806387b162b51461019e5780639dbe87a1146101ce578063abb7acf5146101ff578063b5cb15f71461022f578063b6623e4b1461024d578063bace2caf1461027d576100a9565b80632d34f01f146100ae57806331feb671146100de57806335e1f89d1461010e5780635ba4d07e1461013e578063706ef6ee1461016e575b600080fd5b6100c860048036038101906100c39190611061565b6102ae565b6040516100d5919061142b565b60405180910390f35b6100f860048036038101906100f39190610eee565b6102c8565b60405161010591906114ef565b60405180910390f35b61012860048036038101906101239190610eee565b6105d5565b604051610135919061142b565b60405180910390f35b61015860048036038101906101539190610fa8565b610610565b604051610165919061142b565b60405180910390f35b61018860048036038101906101839190610eee565b6108b8565b604051610195919061142b565b60405180910390f35b6101b860048036038101906101b39190610f33565b6108cc565b6040516101c59190611511565b60405180910390f35b6101e860048036038101906101e39190610f33565b610aed565b6040516101f6929190611446565b60405180910390f35b61021960048036038101906102149190611061565b610c01565b604051610226919061142b565b60405180910390f35b610237610d70565b6040516102449190611511565b60405180910390f35b61026760048036038101906102629190610fa8565b610d8e565b604051610274919061142b565b60405180910390f35b61029760048036038101906102929190610f33565b610dac565b6040516102a5929190611446565b60405180910390f35b60006102bd8686868686610c01565b905095945050505050565b6102d0610dc8565b6102da83836105d5565b610319576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103109061148f565b60405180910390fd5b610321610dc8565b600184846040516103339291906113fb565b9081526020016040518091039020604051806080016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e55780601f106103ba576101008083540402835291602001916103e5565b820191906000526020600020905b8154815290600101906020018083116103c857829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104875780601f1061045c57610100808354040283529160200191610487565b820191906000526020600020905b81548152906001019060200180831161046a57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156105a65783829060005260206000209060020201604051806040016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105715780601f1061054657610100808354040283529160200191610571565b820191906000526020600020905b81548152906001019060200180831161055457829003601f168201915b505050505081526020016001820160009054906101000a900460ff1660ff1660ff1681525050815260200190600101906104b5565b5050505081526020016003820160009054906101000a900460ff16151515158152505090508091505092915050565b6000600183836040516105e99291906113fb565b908152602001604051809103902060030160009054906101000a900460ff16905092915050565b6000806000905060008060006106288c8c8a8a610aed565b809350819550505061063c8a8a8a8a610aed565b809250819450505061064e8a8a6105d5565b1580610658575083155b8061066857508560ff168260ff16105b80610676575060408260ff16105b8061068757508560ff168160ff1610155b156106995760009450505050506108ad565b826106c9576106ab8a8a8a8a8a610c01565b6106bc5760009450505050506108ad565b60019450505050506108ad565b60005b60018b8b6040516106de9291906113fb565b9081526020016040518091039020600201805490508110156108a357606060018c8c60405161070e9291906113fb565b9081526020016040518091039020600201828154811061072a57fe5b90600052602060002090600202016000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107cf5780601f106107a4576101008083540402835291602001916107cf565b820191906000526020600020905b8154815290600101906020018083116107b257829003601f168201915b5050505050905089896040516020016107e99291906113fb565b60405160208183030381529060405280519060200120816040516020016108109190611414565b6040516020818303038152906040528051906020012014156108955782881760018d8d6040516108419291906113fb565b9081526020016040518091039020600201838154811061085d57fe5b906000526020600020906002020160010160006101000a81548160ff021916908360ff160217905550600196505050505050506108ad565b5080806001019150506106cc565b5060009450505050505b979650505050505050565b60006108c483836105d5565b905092915050565b60008085856040516020016108e29291906113fb565b604051602081830303815290604052511415610933576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092a9061146f565b60405180910390fd5b600083836040516020016109489291906113fb565b604051602081830303815290604052511415610999576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610990906114af565b60405180910390fd5b6109a385856105d5565b156109e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109da906114cf565b60405180910390fd5b6002600081819054906101000a900467ffffffffffffffff168092919060010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484600080600260009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000209190610a75929190610df2565b50600060018686604051610a8a9291906113fb565b908152602001604051809103902090508383826001019190610aad929190610df2565b5060018160030160006101000a81548160ff021916908315150217905550600260009054906101000a900467ffffffffffffffff16915050949350505050565b6000806000848490501480610b095750610b0786866105d5565b155b15610b1a5760008091509150610bf8565b610b22610dc8565b610b2c87876102c8565b905060008060005b836040015151811015610bed57606084604001518281518110610b5357fe5b60200260200101516000015190508888604051602001610b749291906113fb565b6040516020818303038152906040528051906020012081604051602001610b9b9190611414565b604051602081830303815290604052805190602001201415610bdf576001935084604001518281518110610bcb57fe5b602002602001015160200151925050610bed565b508080600101915050610b34565b508181945094505050505b94509492505050565b6000610c0d86866105d5565b1580610c1c5750600084849050145b15610c2a5760009050610d67565b6000610c3887878787610aed565b50809150508015610c4d576000915050610d67565b60018787604051610c5f9291906113fb565b90815260200160405180910390206002016001816001815401808255809150500390600052602060002090505060006001808989604051610ca19291906113fb565b908152602001604051809103902060020180549050039050858560018a8a604051610ccd9291906113fb565b90815260200160405180910390206002018381548110610ce957fe5b90600052602060002090600202016000019190610d07929190610df2565b508360018989604051610d1b9291906113fb565b90815260200160405180910390206002018281548110610d3757fe5b906000526020600020906002020160010160006101000a81548160ff021916908360ff1602179055506001925050505b95945050505050565b6000600260009054906101000a900467ffffffffffffffff16905090565b6000610d9f88888888888888610610565b9050979650505050505050565b600080610dbb86868686610aed565b9150915094509492505050565b60405180608001604052806060815260200160608152602001606081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e3357803560ff1916838001178555610e61565b82800160010185558215610e61579182015b82811115610e60578235825591602001919060010190610e45565b5b509050610e6e9190610e72565b5090565b5b80821115610e8b576000816000905550600101610e73565b5090565b60008083601f840112610ea157600080fd5b8235905067ffffffffffffffff811115610eba57600080fd5b602083019150836001820283011115610ed257600080fd5b9250929050565b600081359050610ee88161161d565b92915050565b60008060208385031215610f0157600080fd5b600083013567ffffffffffffffff811115610f1b57600080fd5b610f2785828601610e8f565b92509250509250929050565b60008060008060408587031215610f4957600080fd5b600085013567ffffffffffffffff811115610f6357600080fd5b610f6f87828801610e8f565b9450945050602085013567ffffffffffffffff811115610f8e57600080fd5b610f9a87828801610e8f565b925092505092959194509250565b60008060008060008060006080888a031215610fc357600080fd5b600088013567ffffffffffffffff811115610fdd57600080fd5b610fe98a828b01610e8f565b9750975050602088013567ffffffffffffffff81111561100857600080fd5b6110148a828b01610e8f565b9550955050604088013567ffffffffffffffff81111561103357600080fd5b61103f8a828b01610e8f565b935093505060606110528a828b01610ed9565b91505092959891949750929550565b60008060008060006060868803121561107957600080fd5b600086013567ffffffffffffffff81111561109357600080fd5b61109f88828901610e8f565b9550955050602086013567ffffffffffffffff8111156110be57600080fd5b6110ca88828901610e8f565b935093505060406110dd88828901610ed9565b9150509295509295909350565b60006110f68383611320565b905092915050565b60006111098261153c565b611113818561155f565b9350836020820285016111258561152c565b8060005b85811015611161578484038952815161114285826110ea565b945061114d83611552565b925060208a01995050600181019050611129565b50829750879550505050505092915050565b61117c8161159d565b82525050565b61118b8161159d565b82525050565b600061119d8385611592565b93506111aa8385846115ca565b82840190509392505050565b60006111c182611547565b6111cb8185611570565b93506111db8185602086016115d9565b6111e48161160c565b840191505092915050565b60006111fa82611547565b6112048185611592565b93506112148185602086016115d9565b80840191505092915050565b600061122d601883611581565b91507f7075626c6963206b65792063616e6e6f74206265206e696c00000000000000006000830152602082019050919050565b600061126d601383611581565b91507f7573657220646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b60006112ad601983611581565b91507f70726976617465206b65792063616e6e6f74206265206e696c000000000000006000830152602082019050919050565b60006112ed601383611581565b91507f7573657220616c726561647920657869737473000000000000000000000000006000830152602082019050919050565b6000604083016000830151848203600086015261133d82826111b6565b915050602083015161135260208601826113dd565b508091505092915050565b6000608083016000830151848203600086015261137a82826111b6565b9150506020830151848203602086015261139482826111b6565b915050604083015184820360408601526113ae82826110fe565b91505060608301516113c36060860182611173565b508091505092915050565b6113d7816115a9565b82525050565b6113e6816115bd565b82525050565b6113f5816115bd565b82525050565b6000611408828486611191565b91508190509392505050565b600061142082846111ef565b915081905092915050565b60006020820190506114406000830184611182565b92915050565b600060408201905061145b6000830185611182565b61146860208301846113ec565b9392505050565b6000602082019050818103600083015261148881611220565b9050919050565b600060208201905081810360008301526114a881611260565b9050919050565b600060208201905081810360008301526114c8816112a0565b9050919050565b600060208201905081810360008301526114e8816112e0565b9050919050565b60006020820190508181036000830152611509818461135d565b905092915050565b600060208201905061152660008301846113ce565b92915050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008115159050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156115f75780820151818401526020810190506115dc565b83811115611606576000848401525b50505050565b6000601f19601f8301169050919050565b611626816115bd565b811461163157600080fd5b5056fea26469706673582212209444af9fe409d717a9395c9d656ea3a363c115b6fd260b78b8beb258140d9e8764736f6c63430007000033",
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

