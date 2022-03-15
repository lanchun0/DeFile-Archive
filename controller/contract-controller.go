package controller

import (
	"dfa/service"

	"github.com/gin-gonic/gin"
)

type ContractController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)

	UploadFile(ctx *gin.Context)
	DownloadFile(ctx *gin.Context)
	WriteFile(ctx *gin.Context)
	ShareFile(ctx *gin.Context)

	QueryFile(ctx *gin.Context)
	QueryAllFiles(ctx *gin.Context)
}

type contractController struct {
	contractService service.ContractService
	jWtService      service.JWTService
}

func NewContractController(contractService service.ContractService, jwtService service.JWTService) ContractController {
	return &contractController{
		contractService: contractService,
		jWtService:      jwtService,
	}
}
