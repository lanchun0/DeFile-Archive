package controller

import (
	"dfa/dto"
	"dfa/service"
	"log"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST (no token)
// no params
func (c *dfaController) Register(ctx *gin.Context) {
	user, priv, err := c.contract.Register()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "falied to create new account",
			"err": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":       user,
		"privatekey": priv,
	})
}

// POST (no token)
// param: privatekey string
func (c *dfaController) Login(ctx *gin.Context) {
	// var credentials dto.Credentials
	// err := ctx.ShouldBind(&credentials)
	priv := ctx.DefaultPostForm("privatekey", "N.A.")
	b, err := c.contract.Login(priv)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to login",
			"err": err,
		})
		return
	}
	token := c.jWtService.GenerateToken(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":   "success login",
		"user":  u,
		"token": token,
	})
}

// POST(token)
// param: amount uint
func (c dfaController) Topup(ctx *gin.Context) {
	// var topupRequest dto.TopupRequest
	amountStr := ctx.DefaultPostForm("amount", "0")
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to top up",
			"err": err,
		})
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		if amount <= 0 {
			err = fmt.Errorf("invalid amount: %s", amountStr)
		}
		errFunc(err)
		return
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	ok, tx, err := c.contract.Topup(priv, uint64(amount))
	if err != nil {
		errFunc(err)
		return
	}
	msg := "topup status: "
	if ok {
		msg = msg + "succeeded"
	} else {
		msg = msg + "failed"
	}
	b, _ := c.contract.Login(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         msg,
		"amount":      amount,
		"transaction": tx,
		"user":        u,
	})
}

// POST(token)
// param: amount uint
func (c dfaController) WithDraw(ctx *gin.Context) {
	amountStr := ctx.DefaultPostForm("amount", "0")
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusNotModified, gin.H{
			"msg": "failed to withdraw",
			"err": err,
		})
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		if amount <= 0 {
			err = fmt.Errorf("invalid amount: %s", amountStr)
		}
		errFunc(err)
		return
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid identity",
		})
		return
	}
	ok, tx, err := c.contract.WithDraw(priv, uint64(amount))
	if err != nil {
		errFunc(err)
		return
	}
	msg := "withdraw status: "
	if ok {
		msg = msg + "succeeded"
	} else {
		msg = msg + "failed"
	}
	b, _ := c.contract.Login(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         msg,
		"amount":      amount,
		"transaction": tx,
		"user":        u,
	})
}

// POST(token)
// param: amount uint
func (c dfaController) Approve(ctx *gin.Context) {
	amountStr := ctx.DefaultPostForm("amount", "0")
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to approve",
			"err": err,
		})
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		if amount <= 0 {
			err = fmt.Errorf("invalid amount: %s", amountStr)
		}
		errFunc(err)
		return
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}

	ok, tx, err := c.contract.Approve(priv, uint64(amount))
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
		"price":       amount,
		"transaction": tx,
	})
}

// GET (token)
// no param
func (c *dfaController) GetAllowance(ctx *gin.Context) {
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to query allowance",
		})
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	amount, err := c.contract.GetAllowance(priv)
	if err != nil {
		errFunc(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "allowed to spend with ForForToken",
		"amount": amount,
	})

}
