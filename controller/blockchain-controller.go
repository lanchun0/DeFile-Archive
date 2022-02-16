package controller

import (
	"dfa/entity"
	"dfa/service"

	"github.com/gin-gonic/gin"
)

type BlockchainController interface {
	FindAllData() ([]entity.Data, error)
	SaveData(ctx *gin.Context) (entity.Data, error)

	Register() (entity.Behavior, error)
	RecordAccess(ctx *gin.Context) (entity.Behavior, error)
	Login(ctx *gin.Context) (entity.Behavior, error)
}

type blockchainController struct {
	behaviorService service.BehaviorService
	dataService     service.DataService
}

func NewBlockchainController(ds service.DataService, bs service.BehaviorService) BlockchainController {
	return &blockchainController{
		behaviorService: bs,
		dataService:     ds,
	}
}
