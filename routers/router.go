package routers

import (
	"dfa/controller"
	"dfa/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRouter(server *gin.Engine, controller controller.ContractController) {
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/defile")
	apiRoutes.POST("/register", controller.Register)
	apiRoutes.GET("/login", controller.Login)

	apiRoutes.Use(middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/download", controller.DownloadFile)

		apiRoutes.POST("/share", controller.ShareFile)
		apiRoutes.POST("/upload", controller.UploadFile)

		apiRoutes.GET("/write", controller.WriteFile)
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/all", controller.QueryAllFiles)
		viewRoutes.GET("/one", controller.QueryFile)
	}

}
