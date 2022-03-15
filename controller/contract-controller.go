package controller

import (
	"dfa/service"

	"github.com/gin-gonic/gin"
)

type ContractController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)

	// FindAllData() ([]entity.Data, error)
	// SaveData(ctx *gin.Context) (entity.Data, error)

	// Register() (entity.Behavior, error)
	// RecordAccess(ctx *gin.Context) (entity.Behavior, error)
	// Login(ctx *gin.Context) (entity.Behavior, error)
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
