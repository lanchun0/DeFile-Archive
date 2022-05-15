package routers

import (
	"dfa/controller"
	"dfa/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouter(server *gin.Engine, controller controller.DFAController) {
	// server.Static("/css", "./templates/css")
	// server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/defile")
	apiRoutes.POST("/register", controller.Register)
	apiRoutes.POST("/login", controller.Login)

	apiRoutes.Use(middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/download", controller.DownloadFile)
		apiRoutes.GET("/write", controller.WriteFile)
		apiRoutes.POST("/share", controller.ShareFile)
		apiRoutes.POST("/upload", controller.UploadFile)

		apiRoutes.POST("/buy", controller.BuyAFile)
		apiRoutes.POST("/topup", controller.Topup)
		apiRoutes.POST("/withdraw", controller.WithDraw)
		apiRoutes.POST("/approve", controller.Approve)
		apiRoutes.GET("/allowance", controller.GetAllowance)
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/all", controller.QueryAllFiles)
		viewRoutes.GET("/one", controller.QueryFile)
		viewRoutes.GET("/address", controller.GetAddress)
	}

}
