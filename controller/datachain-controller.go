package controller

import (
	"dfa/entity"

	"github.com/gin-gonic/gin"
)

func (c *blockchainController) FindAllData() ([]entity.Data, error) {
	return c.dataService.FindAll()
}
func (c *blockchainController) SaveData(ctx *gin.Context) (entity.Data, error) {
	var data entity.Data
	ctx.BindJSON(&data)
	return c.dataService.Save(data)
}
