package main

import (
	"dfa/controller"
	"dfa/middlewares"
	"dfa/routers"
	"dfa/service"

	"github.com/gin-gonic/gin"
)

var (
	contractService    service.ContractService       = service.NewContractService()
	jwtService         service.JWTService            = service.NewJWTService()
	contractController controller.ContractController = controller.NewContractController(contractService, jwtService)
)

// var (
// 	dataService          service.DataService             = service.NewData()
// 	behaviorService      service.BehaviorService         = service.NewBehavior()
// 	blockchainController controller.BlockchainController = controller.NewBlockchainController(dataService, behaviorService)
// )

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())
	routers.SetRouter(server, contractController)

	server.Run(":7051")
}
