package main

import (
	"dfa/controller"
	"dfa/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)

	dataService          service.DataService             = service.NewData()
	behaviorService      service.BehaviorService         = service.NewBehavior()
	blockchainController controller.BlockchainController = controller.NewBlockchainController(dataService, behaviorService)
)

func main() {
	server := gin.Default()

	// server.GET("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })

	// server.POST("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.Save(ctx))
	// })

	server.GET("/behavior", func(ctx *gin.Context) {
		data, _ := blockchainController.Register()
		ctx.JSON(200, data)
	})

	server.POST("/behavior", func(ctx *gin.Context) {
		data, err := blockchainController.Login(ctx)
		ctx.JSON(200, data)
		fmt.Println(err)
	})

	server.Run(":9051")
}
