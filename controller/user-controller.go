package controller

import (
	"dfa/dto"
	"dfa/service"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *dfaController) Register(ctx *gin.Context) {
	user, priv, err := c.contract.Register()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "falied to create new account",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":       user,
		"privatekey": priv,
	})
}

// post
func (c *dfaController) Login(ctx *gin.Context) {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "expired identity",
		})
		return
	}
	//fmt.Println(credentials.PrivateKey)
	b, err := c.contract.Login(credentials.PrivateKey)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to login",
		})
		return
	}
	token := c.jWtService.GenerateToken(credentials.PrivateKey)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":   "success login",
		"user":  b,
		"token": token,
	})
}

func (c dfaController) Topup(ctx *gin.Context) {
	var topupRequest dto.TopupRequest
	err := ctx.ShouldBind(&topupRequest)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid top up parameters",
		})
		return
	}
	priv, err := service.ParseToken(topupRequest.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid identity",
		})
		return
	}
	ok, tx, err := c.contract.Topup(priv, topupRequest.Amount)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to toup",
		})
		return
	}
	msg := "topup status: "
	if ok {
		msg = msg + "succeeded"
	} else {
		msg = msg + "failed"
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         msg,
		"amount":      topupRequest.Amount,
		"transaction": tx,
	})
}

func (c dfaController) WithDraw(ctx *gin.Context) {
	var withdraw dto.WithDrawRequest
	err := ctx.ShouldBind(&withdraw)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid top up parameters",
		})
		return
	}
	priv, err := service.ParseToken(withdraw.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid identity",
		})
		return
	}
	ok, tx, err := c.contract.WithDraw(priv, withdraw.Amount)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to withdraw",
		})
		return
	}
	msg := "withdraw status: "
	if ok {
		msg = msg + "succeeded"
	} else {
		msg = msg + "failed"
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         msg,
		"amount":      withdraw.Amount,
		"transaction": tx,
	})
}

func (c dfaController) Approve(ctx *gin.Context) {
	var approve dto.ApproveRequest
	err := ctx.ShouldBind(&approve)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid top up parameters",
		})
		return
	}
	priv, err := service.ParseToken(approve.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid identity",
		})
		return
	}
	ok, tx, err := c.contract.Approve(priv, approve.Price)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to approve",
		})
		return
	}
	msg := "approve status: "
	if ok {
		msg = msg + "succeeded"
	} else {
		msg = msg + "failed"
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         msg,
		"price":       approve.Price,
		"transaction": tx,
	})
}
