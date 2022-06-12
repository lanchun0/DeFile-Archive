package controller

import (
	"dfa/service"
	"dfa/smartcontract"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DFAController interface {
	Init() error

	// POST (no token)
	// no params
	Register(ctx *gin.Context)

	// POST (no token)
	// param: privatekey string
	Login(ctx *gin.Context)

	// POST(token)
	// param: amount uint
	Topup(ctx *gin.Context)

	// POST(token)
	// param: amount uint
	WithDraw(ctx *gin.Context)

	// POST(token)
	// param: amount uint
	Approve(ctx *gin.Context)

	// POST (token)
	// param: permissionlevel string
	//        price uint
	//        tradable bool
	//        file file
	UploadFile(ctx *gin.Context)

	// POST (token)
	// param: id string
	DownloadFile(ctx *gin.Context)

	// POST (token)
	// id string
	// file file
	WriteFile(ctx *gin.Context)

	// POST (token)
	// param: id string
	//        to string
	//        permission string
	ShareFile(ctx *gin.Context)

	// POSt (token)
	// param: id string
	BuyAFile(ctx *gin.Context)

	// GET no token
	// param id string
	QueryFile(ctx *gin.Context)

	// GET no token
	// no param
	QueryAllFiles(ctx *gin.Context)

	// GET (token)
	// no param
	GetAllowance(ctx *gin.Context)

	// GET (token)
	// no param
	GetApproved(ctx *gin.Context)

	// POST (token)
	// param amount uint
	TransferFromContract(ctx *gin.Context)

	// GET (no token)
	// no param
	GetAddress(ctx *gin.Context)
}

type dfaController struct {
	// contractService service.ContractService
	contract    smartcontract.SmartContract
	ipfsService service.IPFSService
	jWtService  service.JWTService
}

func NewDFAController(contract smartcontract.SmartContract, ipfsService service.IPFSService, jwtService service.JWTService) DFAController {
	return &dfaController{
		// contractService: contractService,
		contract:    contract,
		jWtService:  jwtService,
		ipfsService: ipfsService,
	}
}

func (c *dfaController) Init() error {
	err := c.contract.DeployContract()
	if err != nil {
		return fmt.Errorf("failed to deploy the smart contract: %v", err)
	}
	fmt.Println("Succeeded deploying smart contract")
	return nil
}
