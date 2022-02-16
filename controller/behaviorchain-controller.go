package controller

import (
	"dfa/entity"

	"github.com/gin-gonic/gin"
)

func (c *blockchainController) Register() (entity.Behavior, error) {
	return c.behaviorService.Register()
}

func (c *blockchainController) RecordAccess(ctx *gin.Context) (entity.Behavior, error) {
	var log entity.LogInfo
	ctx.BindJSON(&log)
	return c.behaviorService.RecordAccess(log)
}

func (c *blockchainController) Login(ctx *gin.Context) (entity.Behavior, error) {
	var keypair entity.Keypair
	ctx.BindJSON(&keypair)
	return c.behaviorService.Login(keypair.PrivateKey, keypair.PublicKey)
}
