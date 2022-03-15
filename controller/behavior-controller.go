package controller

import (
	"dfa/dto"

	"dfa/general"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *contractController) Register(ctx *gin.Context) {
	behavior, tx, err := c.contractService.Register()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "falied to create new account",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"behavior":    behavior,
		"transaction": tx,
	})
}

// post
func (c *contractController) Login(ctx *gin.Context) {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	ctx.ShouldBind(&credentials)
	plaintext := general.String2Bytes(credentials.PublicKey)
	signature, err := general.MakeSignature(credentials.PrivateKey, plaintext)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid key pairs",
		})
		return
	}
	b, err := c.contractService.Login(credentials.PublicKey, signature)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "incorrect signature",
		})
		return
	}
	token := c.jWtService.GenerateToken(credentials.PublicKey, credentials.PrivateKey)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":      "success login",
		"behavior": b,
		"token":    token,
	})
}
