package controller

import (
	"dfa/dto"
	"dfa/entity"
	"dfa/general"
	"dfa/service"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST (token)
// param: permissionlevel string
//        price uint
//        tradable bool
//        file file
func (c *dfaController) UploadFile(ctx *gin.Context) {
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to upload a file",
			"err": err,
		})
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	permLevelStr := ctx.DefaultPostForm("permissionlevel", "N.A.")
	permLevel, err := dto.String2PermissionLevel(permLevelStr)
	if err != nil {
		errFunc(err)
		return
	}
	priceStr := ctx.DefaultPostForm("price", "0")
	price, err := strconv.Atoi(priceStr)
	if err != nil || price < 0 {
		if price < 0 {
			err = fmt.Errorf("invalid price: %s", priceStr)
		}
		errFunc(err)
		return
	}
	cantrade, tradable := ctx.DefaultPostForm("tradable", "false"), false
	if cantrade == "true" || cantrade == "T" || cantrade == "t" {
		tradable = true
	}
	// upload the file
	f, err := ctx.FormFile("file")
	if err != nil {
		errFunc(err)
		return
	}
	// store the file
	err = ctx.SaveUploadedFile(f, "./tmp/"+f.Filename)
	if err != nil {
		errFunc(err)
		return
	}
	hashAddr, err := c.ipfsService.UploadFile("./tmp/" + f.Filename)
	if err != nil {
		errFunc(err)
		return
	}
	fmt.Println("hash address: ", hashAddr)
	file, err := os.Open("./tmp/" + f.Filename)
	defer func() {
		file.Close()
		dir, _ := ioutil.ReadDir("./tmp")
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
		}
	}()
	if err != nil {
		errFunc(err)
		return
	}
	//
	// fill in the table and submit to blockchain
	id, _ := general.MakeUUID()
	time := general.GenerateTimeStamp()
	digest, err := general.MakeHashDigest(file)
	if err != nil {
		errFunc(err)
		return
	}
	data := entity.Data{
		ID:              id,
		PermissionLevel: permLevel,
		Tradable:        tradable,
		Price:           uint64(price),
		MeteData: entity.MeteData{
			FileName:   f.Filename,
			HashDigest: digest,
			Size:       uint64(f.Size),
			TimeStamp:  time,
		},
	}
	tx1, err := c.contract.CreateFile(priv, data)
	if err != nil {
		errFunc(err)
		return
	}
	tx2, err := c.contract.WriteFile(priv, id, hashAddr, data.MeteData)
	if err != nil {
		errFunc(err)
		return
	}
	view := dto.Data2View(data)
	b, _ := c.contract.Login(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":       "success uploading",
		"data":      view,
		"txcreated": tx1,
		"txwrited":  tx2,
		"user":      u,
	})

}

// GET (token)
// param: id string
func (c *dfaController) DownloadFile(ctx *gin.Context) {
	//c.contractService.DownloadFile()
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to download a file",
			"err": err,
		})
	}
	deleFunc := func() {
		dir, _ := ioutil.ReadDir("./tmp")
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
		}
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	// id := ctx.DefaultPostForm("id", "nil")
	id := ctx.DefaultQuery("id", "nil")
	fmt.Println("get the id:", id)
	ipfsHash, data, err := c.contract.ReadFile(priv, id)
	if err != nil {
		errFunc(err)
		return
	}
	name, err := c.ipfsService.DownloadFile(ipfsHash)
	if err != nil {
		errFunc(err)
		return
	}
	defer deleFunc()
	err = os.Rename("./"+name, "./tmp/"+data.MeteData.FileName)
	if err != nil {
		errFunc(err)
		return
	}
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", data.MeteData.FileName))
	ctx.File("./tmp/" + data.MeteData.FileName)
}

// POST (token)
// id string
// file file
func (c *dfaController) WriteFile(ctx *gin.Context) {
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to upload a file",
			"err": err,
		})
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	id := ctx.DefaultPostForm("id", "N.A.")
	// save the file locally
	f, err := ctx.FormFile("file")
	if err != nil {
		errFunc(err)
		return
	}
	err = ctx.SaveUploadedFile(f, "./tmp/"+f.Filename)
	if err != nil {
		errFunc(err)
		return
	}
	hashAddr, err := c.ipfsService.UploadFile("./tmp/" + f.Filename)
	if err != nil {
		errFunc(err)
		return
	}
	fmt.Println("hash address: ", hashAddr)
	file, err := os.Open("./tmp/" + f.Filename)
	defer func() {
		file.Close()
		dir, _ := ioutil.ReadDir("./tmp")
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
		}
	}()
	if err != nil {
		errFunc(err)
		return
	}
	time := general.GenerateTimeStamp()
	digest, err := general.MakeHashDigest(file)
	if err != nil {
		errFunc(err)
		return
	}
	data := entity.MeteData{
		FileName:   f.Filename,
		HashDigest: digest,
		Size:       uint64(f.Size),
		TimeStamp:  time,
	}
	tx, err := c.contract.WriteFile(priv, id, hashAddr, data)
	if err != nil {
		errFunc(err)
		return
	}
	b, _ := c.contract.Login(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "succeeded in writing file: " + id,
		"transaction": tx,
		"user":        u,
	})
}

// POST (token)
// param: id string
//        to string
//        permission string
func (c *dfaController) ShareFile(ctx *gin.Context) {
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to share a file",
			"err": err,
		})
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	id := ctx.DefaultPostForm("id", "N.A.")
	to := ctx.DefaultPostForm("to", "N.A.")
	perm := ctx.DefaultPostForm("permission", "N.A.")
	permission, err := dto.String2Permission(perm)
	if err != nil {
		errFunc(err)
		return
	}
	tx, err := c.contract.ShareFile(priv, to, id, entity.Permission(permission))
	if err != nil {
		errFunc(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"trasaction": tx,
	})
}

// GET no token
// param id string
func (c *dfaController) QueryFile(ctx *gin.Context) {
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "failed to query a file",
		})
	}
	id := ctx.DefaultQuery("id", "N.A.")
	data, err := c.contract.QueryFile(id)
	if err != nil {
		errFunc(err)
		return
	}
	view := dto.Data2View(data)
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"data":   view,
	})

}

// GET no token
// no param
func (c *dfaController) QueryAllFiles(ctx *gin.Context) {
	data, err := c.contract.QueryAllFiles()
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to query files",
			"err": err,
		})
	}
	if err != nil {
		errFunc(err)
		return
	}
	views := []dto.ViewData{}
	for _, d := range data {
		view := dto.Data2View(d)
		views = append(views, view)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "inquirying all files",
		"data": views,
	})
}

// POSt (token)
// param: id string
func (c *dfaController) BuyAFile(ctx *gin.Context) {
	errFunc := func(err error) {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to buy a file",
			"err": err,
		})
	}
	token := ctx.GetHeader("Authorization")
	priv, err := service.ParseToken(token)
	if err != nil {
		errFunc(err)
		return
	}
	id := ctx.DefaultPostForm("id", "N.A.")
	tx, ok, err := c.contract.PurchaseFile(priv, id)
	if err != nil {
		errFunc(err)
		return
	}
	if !ok {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to buy a file: " + err.Error(),
		})
		return
	}
	b, _ := c.contract.Login(priv)
	u := dto.Behavior2View(b)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "succeeded in buying a file",
		"transaction": tx,
		"user":        u,
	})
}

// GET (no token)
// no param
func (c *dfaController) GetAddress(ctx *gin.Context) {
	dataAddr := c.contract.GetDataContractAddress()
	userAddr := c.contract.GetUserContractAddress()
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":       "smart contract address on Etherueum",
		"fileaddr":  dataAddr,
		"tokenaddr": userAddr,
	})

}
