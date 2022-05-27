package main

import (
	"dfa/controller"
	"dfa/middlewares"
	"dfa/routers"
	"dfa/service"
	"dfa/smartcontract"

	"github.com/gin-gonic/gin"
)

var (
	contract      smartcontract.SmartContract = smartcontract.NewConract()
	jwtService    service.JWTService          = service.NewJWTService()
	ipfsService   service.IPFSService         = service.NewIPFSService()
	dfaController controller.DFAController    = controller.NewDFAController(contract, ipfsService, jwtService)
)

// var (
// 	dataService          service.DataService             = service.NewData()
// 	behaviorService      service.BehaviorService         = service.NewBehavior()
// 	blockchainController controller.BlockchainController = controller.NewBlockchainController(dataService, behaviorService)
// )

func main() {
	server := gin.New()
	dfaController.Init()
	server.Use(gin.Recovery(), middlewares.Logger())
	//server.Use(middlewares.Cors())
	routers.SetRouter(server, dfaController)

	server.Run(":7051")
}
