// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// ForForTokenAsset is an auto generated low-level Go binding around an user-defined struct.
type ForForTokenAsset struct {
	FileID     string
	Permission uint8
}

// ForForTokenUser is an auto generated low-level Go binding around an user-defined struct.
type ForForTokenUser struct {
	Addr    common.Address
	Balance *big.Int
	Assets  []ForForTokenAsset
}

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TopUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fileID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"permission\",\"type\":\"uint8\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"assetsID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"permission\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetsIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"callApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"callTransferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"permission\",\"type\":\"uint8\"}],\"internalType\":\"structForForToken.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"}],\"internalType\":\"structForForToken.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"topup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_fileID\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_perm\",\"type\":\"uint8\"}],\"name\":\"updateAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"sucess\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600381526020017f464f520000000000000000000000000000000000000000000000000000000000815250600190805190602001906200005f92919062000111565b506040518060400160405280600c81526020017f466f72466f7220546f6b656e000000000000000000000000000000000000000081525060009080519060200190620000ad92919062000111565b506012600260006101000a81548160ff021916908360ff160217905550610e10600381905550600260009054906101000a900460ff1660ff16600a620000f491906200021c565b6305f5e10062000105919062000359565b60048190555062000465565b8280546200011f90620003c4565b90600052602060002090601f0160209004810192826200014357600085556200018f565b82601f106200015e57805160ff19168380011785556200018f565b828001600101855582156200018f579182015b828111156200018e57825182559160200191906001019062000171565b5b5090506200019e9190620001a2565b5090565b5b80821115620001bd576000816000905550600101620001a3565b5090565b6000808291508390505b60018511156200021357808604811115620001eb57620001ea620003fa565b5b6001851615620001fb5780820291505b80810290506200020b8562000458565b9450620001cb565b94509492505050565b60006200022982620003ba565b91506200023683620003ba565b9250620002657fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200026d565b905092915050565b6000826200027f576001905062000352565b816200028f576000905062000352565b8160018114620002a85760028114620002b357620002e9565b600191505062000352565b60ff841115620002c857620002c7620003fa565b5b8360020a915084821115620002e257620002e1620003fa565b5b5062000352565b5060208310610133831016604e8410600b8410161715620003235782820a9050838111156200031d576200031c620003fa565b5b62000352565b620003328484846001620001c1565b925090508184048111156200034c576200034b620003fa565b5b81810290505b9392505050565b60006200036682620003ba565b91506200037383620003ba565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620003af57620003ae620003fa565b5b828202905092915050565b6000819050919050565b60006002820490506001821680620003dd57607f821691505b60208210811415620003f457620003f362000429565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60008160011c9050919050565b612bc280620004756000396000f3fe6080604052600436106101395760003560e01c80635c658165116100ab578063bdafb6eb1161006f578063bdafb6eb1461049c578063bffa9258146104da578063c4938a4d14610517578063dd62ed3e14610542578063de70cc2c1461057f578063e4085dec146105bc57610140565b80635c6581651461038f57806370a08231146103cc57806395d89b4114610409578063a035b1fe14610434578063a9059cbb1461045f57610140565b806323b872dd116100fd57806323b872dd1461024057806327e235e31461027d5780632e1a7d4d146102ba57806330db555f146102f7578063313ce56714610327578063321329601461035257610140565b8063047fc9aa1461014557806306fdde0314610170578063095ea7b31461019b57806314acbb27146101d857806318160ddd1461021557610140565b3661014057005b600080fd5b34801561015157600080fd5b5061015a6105f9565b60405161016791906126d8565b60405180910390f35b34801561017c57600080fd5b506101856105ff565b6040516101929190612664565b60405180910390f35b3480156101a757600080fd5b506101c260048036038101906101bd9190612374565b61068d565b6040516101cf9190612649565b60405180910390f35b3480156101e457600080fd5b506101ff60048036038101906101fa919061223d565b610791565b60405161020c9190612649565b60405180910390f35b34801561022157600080fd5b5061022a610ab2565b60405161023791906126d8565b60405180910390f35b34801561024c57600080fd5b50610267600480360381019061026291906122c5565b610abc565b6040516102749190612649565b60405180910390f35b34801561028957600080fd5b506102a4600480360381019061029f91906121d0565b610dd8565b6040516102b191906126d8565b60405180910390f35b3480156102c657600080fd5b506102e160048036038101906102dc91906123b4565b610df0565b6040516102ee9190612649565b60405180910390f35b610311600480360381019061030c91906123b4565b610fd3565b60405161031e91906126d8565b60405180910390f35b34801561033357600080fd5b5061033c6111b6565b60405161034991906126f3565b60405180910390f35b34801561035e57600080fd5b50610379600480360381019061037491906122c5565b6111c9565b6040516103869190612649565b60405180910390f35b34801561039b57600080fd5b506103b660048036038101906103b191906121fd565b611557565b6040516103c391906126d8565b60405180910390f35b3480156103d857600080fd5b506103f360048036038101906103ee91906121d0565b61157c565b60405161040091906126d8565b60405180910390f35b34801561041557600080fd5b5061041e6115c5565b60405161042b9190612664565b60405180910390f35b34801561044057600080fd5b50610449611653565b60405161045691906126d8565b60405180910390f35b34801561046b57600080fd5b5061048660048036038101906104819190612374565b611659565b6040516104939190612649565b60405180910390f35b3480156104a857600080fd5b506104c360048036038101906104be9190612318565b611891565b6040516104d1929190612686565b60405180910390f35b3480156104e657600080fd5b5061050160048036038101906104fc91906121d0565b61196d565b60405161050e91906126d8565b60405180910390f35b34801561052357600080fd5b5061052c611985565b60405161053991906126b6565b60405180910390f35b34801561054e57600080fd5b50610569600480360381019061056491906121fd565b611d9c565b60405161057691906126d8565b60405180910390f35b34801561058b57600080fd5b506105a660048036038101906105a19190612374565b611e23565b6040516105b39190612664565b60405180910390f35b3480156105c857600080fd5b506105e360048036038101906105de9190612374565b611ed0565b6040516105f09190612649565b60405180910390f35b60045481565b6000805461060c90612981565b80601f016020809104026020016040519081016040528092919081815260200182805461063890612981565b80156106855780601f1061065a57610100808354040283529160200191610685565b820191906000526020600020905b81548152906001019060200180831161066857829003601f168201915b505050505081565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461071b91906127d5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161077f91906126d8565b60405180910390a36001905092915050565b600080600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085856040516107e3929190612619565b908152602001604051809103902060010160009054906101000a900460ff1660ff161461089a5781600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020858560405161085a929190612619565b908152602001604051809103902060010160008282829054906101000a900460ff161792506101000a81548160ff021916908360ff160217905550610a28565b6000600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508484600860008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020919061093d929190611fd4565b50600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919061098e906129e4565b91905055506000600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002086866040516109e4929190612619565b908152602001604051809103902090508585826000019190610a07929190611fd4565b50838160010160006101000a81548160ff021916908360ff16021790555050505b8383604051610a38929190612619565b60405180910390208573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f1f718c8dca15ead6170a985e926e90538a163d7c0845a3ab643bcaabe47488d685604051610a9d91906126f3565b60405180910390a46001905095945050505050565b6000600454905090565b600081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610b0a57600080fd5b81600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610b9357600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c1e91906127d5565b1015610c2957600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c7891906128b6565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d0b91906128b6565b9250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d6191906127d5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610dc591906126d8565b60405180910390a3600190509392505050565b60056020528060005260406000206000915090505481565b6000808211610dfe57600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610e4a57600080fd5b600060035483610e5a919061285c565b905060008382610e6a919061282b565b90506003548114610e8057600092505050610fce565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905084600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f1391906128b6565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050610f5f5760009350505050610fce565b8460046000828254610f7191906127d5565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436486604051610fbe91906126d8565b60405180910390a2600193505050505b919050565b6000600454821115610fe457600080fd5b60008211610ff157600080fd5b600060035483611001919061285c565b905060008382611011919061282b565b905060006003548214158061102557508234105b1561103657600093505050506111b1565b8234111561105157823461104a91906128b6565b9050611056565b600090505b3073ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f1935050505015801561109c573d6000803e3d6000fd5b50600081146110ed573373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f193505050501580156110eb573d6000803e3d6000fd5b505b84600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461113c91906127d5565b92505081905550846004600082825461115591906128b6565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df866040516111a291906126d8565b60405180910390a28493505050505b919050565b600260009054906101000a900460ff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156112345750600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614155b61123d57600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561128957600080fd5b81600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561131257600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461139d91906127d5565b10156113a857600080fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113f791906128b6565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461148a91906128b6565b9250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546114e091906127d5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161154491906126d8565b60405180910390a3600190509392505050565b6006602052816000526040600020602052806000526040600020600091509150505481565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600180546115d290612981565b80601f01602080910402602001604051908101604052809291908181526020018280546115fe90612981565b801561164b5780601f106116205761010080835404028352916020019161164b565b820191906000526020600020905b81548152906001019060200180831161162e57829003601f168201915b505050505081565b60035481565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561169457600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156116e057600080fd5b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461176b91906127d5565b101561177657600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117c591906128b6565b9250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461181b91906127d5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161187f91906126d8565b60405180910390a36001905092915050565b600960205281600052604060002081805160208101820180518482526020830160208501208183528095505050505050600091509150508060000180546118d790612981565b80601f016020809104026020016040519081016040528092919081815260200182805461190390612981565b80156119505780601f1061192557610100808354040283529160200191611950565b820191906000526020600020905b81548152906001019060200180831161193357829003601f168201915b5050505050908060010160009054906101000a900460ff16905082565b60076020528060005260406000206000915090505481565b61198d61205a565b61199561205a565b33816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054816020018181525050600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205467ffffffffffffffff811115611a6f57611a6e612ae9565b5b604051908082528060200260200182016040528015611aa857816020015b611a95612091565b815260200190600190039081611a8d5790505b50816040018190525060005b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811015611d94576000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000208054611b5890612981565b80601f0160208091040260200160405190810160405280929190818152602001828054611b8490612981565b8015611bd15780601f10611ba657610100808354040283529160200191611bd1565b820191906000526020600020905b815481529060010190602001808311611bb457829003601f168201915b50505050509050600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081604051611c259190612632565b90815260200160405180910390206000018054611c4190612981565b80601f0160208091040260200160405190810160405280929190818152602001828054611c6d90612981565b8015611cba5780601f10611c8f57610100808354040283529160200191611cba565b820191906000526020600020905b815481529060010190602001808311611c9d57829003601f168201915b505050505083604001518381518110611cd657611cd5612aba565b5b602002602001015160000181905250600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081604051611d329190612632565b908152602001604051809103902060010160009054906101000a900460ff1683604001518381518110611d6857611d67612aba565b5b60200260200101516020019060ff16908160ff1681525050508080611d8c906129e4565b915050611ab4565b508091505090565b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6008602052816000526040600020602052806000526040600020600091509150508054611e4f90612981565b80601f0160208091040260200160405190810160405280929190818152602001828054611e7b90612981565b8015611ec85780601f10611e9d57610100808354040283529160200191611ec8565b820191906000526020600020905b815481529060010190602001808311611eab57829003601f168201915b505050505081565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611f5e91906127d5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051611fc291906126d8565b60405180910390a36001905092915050565b828054611fe090612981565b90600052602060002090601f0160209004810192826120025760008555612049565b82601f1061201b57803560ff1916838001178555612049565b82800160010185558215612049579182015b8281111561204857823582559160200191906001019061202d565b5b50905061205691906120ae565b5090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b604051806040016040528060608152602001600060ff1681525090565b5b808211156120c75760008160009055506001016120af565b5090565b60006120de6120d984612733565b61270e565b9050828152602081018484840111156120fa576120f9612b27565b5b61210584828561293f565b509392505050565b60008135905061211c81612b47565b92915050565b60008083601f84011261213857612137612b1d565b5b8235905067ffffffffffffffff81111561215557612154612b18565b5b60208301915083600182028301111561217157612170612b22565b5b9250929050565b600082601f83011261218d5761218c612b1d565b5b813561219d8482602086016120cb565b91505092915050565b6000813590506121b581612b5e565b92915050565b6000813590506121ca81612b75565b92915050565b6000602082840312156121e6576121e5612b31565b5b60006121f48482850161210d565b91505092915050565b6000806040838503121561221457612213612b31565b5b60006122228582860161210d565b92505060206122338582860161210d565b9150509250929050565b60008060008060006080868803121561225957612258612b31565b5b60006122678882890161210d565b95505060206122788882890161210d565b945050604086013567ffffffffffffffff81111561229957612298612b2c565b5b6122a588828901612122565b935093505060606122b8888289016121bb565b9150509295509295909350565b6000806000606084860312156122de576122dd612b31565b5b60006122ec8682870161210d565b93505060206122fd8682870161210d565b925050604061230e868287016121a6565b9150509250925092565b6000806040838503121561232f5761232e612b31565b5b600061233d8582860161210d565b925050602083013567ffffffffffffffff81111561235e5761235d612b2c565b5b61236a85828601612178565b9150509250929050565b6000806040838503121561238b5761238a612b31565b5b60006123998582860161210d565b92505060206123aa858286016121a6565b9150509250929050565b6000602082840312156123ca576123c9612b31565b5b60006123d8848285016121a6565b91505092915050565b60006123ed8383612550565b905092915050565b6123fe816128ea565b82525050565b600061240f82612774565b6124198185612797565b93508360208202850161242b85612764565b8060005b85811015612467578484038952815161244885826123e1565b94506124538361278a565b925060208a0199505060018101905061242f565b50829750879550505050505092915050565b612482816128fc565b82525050565b600061249483856127ca565b93506124a183858461293f565b82840190509392505050565b60006124b88261277f565b6124c281856127a8565b93506124d281856020860161294e565b6124db81612b36565b840191505092915050565b60006124f18261277f565b6124fb81856127b9565b935061250b81856020860161294e565b61251481612b36565b840191505092915050565b600061252a8261277f565b61253481856127ca565b935061254481856020860161294e565b80840191505092915050565b6000604083016000830151848203600086015261256d82826124ad565b915050602083015161258260208601826125fb565b508091505092915050565b60006060830160008301516125a560008601826123f5565b5060208301516125b860208601826125dd565b50604083015184820360408601526125d08282612404565b9150508091505092915050565b6125e681612928565b82525050565b6125f581612928565b82525050565b61260481612932565b82525050565b61261381612932565b82525050565b6000612626828486612488565b91508190509392505050565b600061263e828461251f565b915081905092915050565b600060208201905061265e6000830184612479565b92915050565b6000602082019050818103600083015261267e81846124e6565b905092915050565b600060408201905081810360008301526126a081856124e6565b90506126af602083018461260a565b9392505050565b600060208201905081810360008301526126d0818461258d565b905092915050565b60006020820190506126ed60008301846125ec565b92915050565b6000602082019050612708600083018461260a565b92915050565b6000612718612729565b905061272482826129b3565b919050565b6000604051905090565b600067ffffffffffffffff82111561274e5761274d612ae9565b5b61275782612b36565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006127e082612928565b91506127eb83612928565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156128205761281f612a2d565b5b828201905092915050565b600061283682612928565b915061284183612928565b92508261285157612850612a5c565b5b828204905092915050565b600061286782612928565b915061287283612928565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156128ab576128aa612a2d565b5b828202905092915050565b60006128c182612928565b91506128cc83612928565b9250828210156128df576128de612a2d565b5b828203905092915050565b60006128f582612908565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561296c578082015181840152602081019050612951565b8381111561297b576000848401525b50505050565b6000600282049050600182168061299957607f821691505b602082108114156129ad576129ac612a8b565b5b50919050565b6129bc82612b36565b810181811067ffffffffffffffff821117156129db576129da612ae9565b5b80604052505050565b60006129ef82612928565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415612a2257612a21612a2d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b612b50816128ea565b8114612b5b57600080fd5b50565b612b6781612928565b8114612b7257600080fd5b50565b612b7e81612932565b8114612b8957600080fd5b5056fea2646970667358221220ed4a4bfe07834b4fc7ee5ea141806d1ddd9d90f9dc7d9c94fe89fb069577959864736f6c63430008070033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Token *TokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Token *TokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, tokenOwner, spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Token *TokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Token *TokenSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowed(&_Token.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Token *TokenCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowed(&_Token.CallOpts, arg0, arg1)
}

// AssetCount is a free data retrieval call binding the contract method 0xbffa9258.
//
// Solidity: function assetCount(address ) view returns(uint256)
func (_Token *TokenCaller) AssetCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "assetCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetCount is a free data retrieval call binding the contract method 0xbffa9258.
//
// Solidity: function assetCount(address ) view returns(uint256)
func (_Token *TokenSession) AssetCount(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.AssetCount(&_Token.CallOpts, arg0)
}

// AssetCount is a free data retrieval call binding the contract method 0xbffa9258.
//
// Solidity: function assetCount(address ) view returns(uint256)
func (_Token *TokenCallerSession) AssetCount(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.AssetCount(&_Token.CallOpts, arg0)
}

// AssetsID is a free data retrieval call binding the contract method 0xbdafb6eb.
//
// Solidity: function assetsID(address , string ) view returns(string fileID, uint8 permission)
func (_Token *TokenCaller) AssetsID(opts *bind.CallOpts, arg0 common.Address, arg1 string) (struct {
	FileID     string
	Permission uint8
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "assetsID", arg0, arg1)

	outstruct := new(struct {
		FileID     string
		Permission uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileID = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Permission = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// AssetsID is a free data retrieval call binding the contract method 0xbdafb6eb.
//
// Solidity: function assetsID(address , string ) view returns(string fileID, uint8 permission)
func (_Token *TokenSession) AssetsID(arg0 common.Address, arg1 string) (struct {
	FileID     string
	Permission uint8
}, error) {
	return _Token.Contract.AssetsID(&_Token.CallOpts, arg0, arg1)
}

// AssetsID is a free data retrieval call binding the contract method 0xbdafb6eb.
//
// Solidity: function assetsID(address , string ) view returns(string fileID, uint8 permission)
func (_Token *TokenCallerSession) AssetsID(arg0 common.Address, arg1 string) (struct {
	FileID     string
	Permission uint8
}, error) {
	return _Token.Contract.AssetsID(&_Token.CallOpts, arg0, arg1)
}

// AssetsIndex is a free data retrieval call binding the contract method 0xde70cc2c.
//
// Solidity: function assetsIndex(address , uint256 ) view returns(string)
func (_Token *TokenCaller) AssetsIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "assetsIndex", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetsIndex is a free data retrieval call binding the contract method 0xde70cc2c.
//
// Solidity: function assetsIndex(address , uint256 ) view returns(string)
func (_Token *TokenSession) AssetsIndex(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Token.Contract.AssetsIndex(&_Token.CallOpts, arg0, arg1)
}

// AssetsIndex is a free data retrieval call binding the contract method 0xde70cc2c.
//
// Solidity: function assetsIndex(address , uint256 ) view returns(string)
func (_Token *TokenCallerSession) AssetsIndex(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Token.Contract.AssetsIndex(&_Token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Token *TokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Token *TokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, tokenOwner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Token *TokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Token *TokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.Balances(&_Token.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Token *TokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.Balances(&_Token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Token *TokenCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Token *TokenSession) Price() (*big.Int, error) {
	return _Token.Contract.Price(&_Token.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Token *TokenCallerSession) Price() (*big.Int, error) {
	return _Token.Contract.Price(&_Token.CallOpts)
}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((address,uint256,(string,uint8)[]))
func (_Token *TokenCaller) QueryUser(opts *bind.CallOpts) (ForForTokenUser, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "queryUser")

	if err != nil {
		return *new(ForForTokenUser), err
	}

	out0 := *abi.ConvertType(out[0], new(ForForTokenUser)).(*ForForTokenUser)

	return out0, err

}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((address,uint256,(string,uint8)[]))
func (_Token *TokenSession) QueryUser() (ForForTokenUser, error) {
	return _Token.Contract.QueryUser(&_Token.CallOpts)
}

// QueryUser is a free data retrieval call binding the contract method 0xc4938a4d.
//
// Solidity: function queryUser() view returns((address,uint256,(string,uint8)[]))
func (_Token *TokenCallerSession) QueryUser() (ForForTokenUser, error) {
	return _Token.Contract.QueryUser(&_Token.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Token *TokenCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Token *TokenSession) Supply() (*big.Int, error) {
	return _Token.Contract.Supply(&_Token.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Token *TokenCallerSession) Supply() (*big.Int, error) {
	return _Token.Contract.Supply(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Token *TokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Token *TokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, tokens)
}

// CallApprove is a paid mutator transaction binding the contract method 0xe4085dec.
//
// Solidity: function callApprove(address spender, uint256 tokens) returns(bool)
func (_Token *TokenTransactor) CallApprove(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "callApprove", spender, tokens)
}

// CallApprove is a paid mutator transaction binding the contract method 0xe4085dec.
//
// Solidity: function callApprove(address spender, uint256 tokens) returns(bool)
func (_Token *TokenSession) CallApprove(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CallApprove(&_Token.TransactOpts, spender, tokens)
}

// CallApprove is a paid mutator transaction binding the contract method 0xe4085dec.
//
// Solidity: function callApprove(address spender, uint256 tokens) returns(bool)
func (_Token *TokenTransactorSession) CallApprove(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CallApprove(&_Token.TransactOpts, spender, tokens)
}

// CallTransferFrom is a paid mutator transaction binding the contract method 0x32132960.
//
// Solidity: function callTransferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactor) CallTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "callTransferFrom", from, to, tokens)
}

// CallTransferFrom is a paid mutator transaction binding the contract method 0x32132960.
//
// Solidity: function callTransferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenSession) CallTransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CallTransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// CallTransferFrom is a paid mutator transaction binding the contract method 0x32132960.
//
// Solidity: function callTransferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactorSession) CallTransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.CallTransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// Topup is a paid mutator transaction binding the contract method 0x30db555f.
//
// Solidity: function topup(uint256 _tokens) payable returns(uint256 amount)
func (_Token *TokenTransactor) Topup(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "topup", _tokens)
}

// Topup is a paid mutator transaction binding the contract method 0x30db555f.
//
// Solidity: function topup(uint256 _tokens) payable returns(uint256 amount)
func (_Token *TokenSession) Topup(_tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Topup(&_Token.TransactOpts, _tokens)
}

// Topup is a paid mutator transaction binding the contract method 0x30db555f.
//
// Solidity: function topup(uint256 _tokens) payable returns(uint256 amount)
func (_Token *TokenTransactorSession) Topup(_tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Topup(&_Token.TransactOpts, _tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Token *TokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokens)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x14acbb27.
//
// Solidity: function updateAsset(address from, address to, string _fileID, uint8 _perm) returns(bool)
func (_Token *TokenTransactor) UpdateAsset(opts *bind.TransactOpts, from common.Address, to common.Address, _fileID string, _perm uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateAsset", from, to, _fileID, _perm)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x14acbb27.
//
// Solidity: function updateAsset(address from, address to, string _fileID, uint8 _perm) returns(bool)
func (_Token *TokenSession) UpdateAsset(from common.Address, to common.Address, _fileID string, _perm uint8) (*types.Transaction, error) {
	return _Token.Contract.UpdateAsset(&_Token.TransactOpts, from, to, _fileID, _perm)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x14acbb27.
//
// Solidity: function updateAsset(address from, address to, string _fileID, uint8 _perm) returns(bool)
func (_Token *TokenTransactorSession) UpdateAsset(from common.Address, to common.Address, _fileID string, _perm uint8) (*types.Transaction, error) {
	return _Token.Contract.UpdateAsset(&_Token.TransactOpts, from, to, _fileID, _perm)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokens) returns(bool sucess)
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts, tokens *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw", tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokens) returns(bool sucess)
func (_Token *TokenSession) Withdraw(tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokens) returns(bool sucess)
func (_Token *TokenTransactorSession) Withdraw(tokens *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, tokens)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTopUpIterator is returned from FilterTopUp and is used to iterate over the raw logs and unpacked data for TopUp events raised by the Token contract.
type TokenTopUpIterator struct {
	Event *TokenTopUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTopUp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTopUp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTopUp represents a TopUp event raised by the Token contract.
type TokenTopUp struct {
	TokenOwner common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUp is a free log retrieval operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) FilterTopUp(opts *bind.FilterOpts, tokenOwner []common.Address) (*TokenTopUpIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TopUp", tokenOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenTopUpIterator{contract: _Token.contract, event: "TopUp", logs: logs, sub: sub}, nil
}

// WatchTopUp is a free log subscription operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) WatchTopUp(opts *bind.WatchOpts, sink chan<- *TokenTopUp, tokenOwner []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TopUp", tokenOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTopUp)
				if err := _Token.contract.UnpackLog(event, "TopUp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTopUp is a log parse operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) ParseTopUp(log types.Log) (*TokenTopUp, error) {
	event := new(TokenTopUp)
	if err := _Token.contract.UnpackLog(event, "TopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Token contract.
type TokenUpdateIterator struct {
	Event *TokenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUpdate represents a Update event raised by the Token contract.
type TokenUpdate struct {
	From       common.Address
	To         common.Address
	FileID     common.Hash
	Permission uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x1f718c8dca15ead6170a985e926e90538a163d7c0845a3ab643bcaabe47488d6.
//
// Solidity: event Update(address indexed from, address indexed to, string indexed fileID, uint8 permission)
func (_Token *TokenFilterer) FilterUpdate(opts *bind.FilterOpts, from []common.Address, to []common.Address, fileID []string) (*TokenUpdateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fileIDRule []interface{}
	for _, fileIDItem := range fileID {
		fileIDRule = append(fileIDRule, fileIDItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Update", fromRule, toRule, fileIDRule)
	if err != nil {
		return nil, err
	}
	return &TokenUpdateIterator{contract: _Token.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x1f718c8dca15ead6170a985e926e90538a163d7c0845a3ab643bcaabe47488d6.
//
// Solidity: event Update(address indexed from, address indexed to, string indexed fileID, uint8 permission)
func (_Token *TokenFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *TokenUpdate, from []common.Address, to []common.Address, fileID []string) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fileIDRule []interface{}
	for _, fileIDItem := range fileID {
		fileIDRule = append(fileIDRule, fileIDItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Update", fromRule, toRule, fileIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUpdate)
				if err := _Token.contract.UnpackLog(event, "Update", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdate is a log parse operation binding the contract event 0x1f718c8dca15ead6170a985e926e90538a163d7c0845a3ab643bcaabe47488d6.
//
// Solidity: event Update(address indexed from, address indexed to, string indexed fileID, uint8 permission)
func (_Token *TokenFilterer) ParseUpdate(log types.Log) (*TokenUpdate, error) {
	event := new(TokenUpdate)
	if err := _Token.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Token contract.
type TokenWithdrawIterator struct {
	Event *TokenWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdraw represents a Withdraw event raised by the Token contract.
type TokenWithdraw struct {
	TokenOwner common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) FilterWithdraw(opts *bind.FilterOpts, tokenOwner []common.Address) (*TokenWithdrawIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Withdraw", tokenOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawIterator{contract: _Token.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TokenWithdraw, tokenOwner []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Withdraw", tokenOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdraw)
				if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed tokenOwner, uint256 tokens)
func (_Token *TokenFilterer) ParseWithdraw(log types.Log) (*TokenWithdraw, error) {
	event := new(TokenWithdraw)
	if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
