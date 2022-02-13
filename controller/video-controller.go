package controller

import (
	"dfa/entity"
	"dfa/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}
