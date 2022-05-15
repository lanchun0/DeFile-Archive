package controller

import (
	"dfa/service"
	"dfa/smartcontract"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DFAController interface {
	Init() error
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Topup(ctx *gin.Context)
	WithDraw(ctx *gin.Context)
	Approve(ctx *gin.Context)

	UploadFile(ctx *gin.Context)
	DownloadFile(ctx *gin.Context)
	WriteFile(ctx *gin.Context)
	ShareFile(ctx *gin.Context)
	BuyAFile(ctx *gin.Context)

	QueryFile(ctx *gin.Context)
	QueryAllFiles(ctx *gin.Context)
	GetAllowance(ctx *gin.Context)
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
