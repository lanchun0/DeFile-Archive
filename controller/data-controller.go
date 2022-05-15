package controller

import (
	"dfa/dto"
	"dfa/entity"
	"dfa/general"
	"dfa/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func (c *dfaController) UploadFile(ctx *gin.Context) {
	var request dto.UploadFileRequest
	ctx.ShouldBind(&request)
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to upload a file",
		})
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		errFunc(err)
		return
	}
	permission, err := dto.String2Permission(request.PermissionLevel)
	if err != nil {
		errFunc(err)
		return
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
		dir, _ := ioutil.ReadDir("/tmp")
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
		PermissionLevel: permission,
		Tradable:        request.Tradable,
		Price:           request.Price,
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
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "success uploading",
		"data":        view,
		"transaction": "creat transaction: " + tx1 + "\nwrite transaction: " + tx2,
	})

}

func (c *dfaController) DownloadFile(ctx *gin.Context) {
	//c.contractService.DownloadFile()
	var request dto.DownloadFileRequest
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to download a file",
		})
	}
	err := ctx.ShouldBind(request)
	if err != nil {
		errFunc(err)
		return
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		errFunc(err)
		return
	}
	ipfsHash, data, err := c.contract.ReadFile(priv, request.ID)
	if err != nil {
		errFunc(err)
		return
	}
	name, err := c.ipfsService.DownloadFile(ipfsHash)
	if err != nil {
		errFunc(err)
		return
	}
	err = os.Rename("./tmp"+name, data.MeteData.FileName)
	if err != nil {
		errFunc(err)
		return
	}
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", data.MeteData.FileName))
	ctx.File("./tmp" + data.MeteData.FileName)
	dir, _ := ioutil.ReadDir("/tmp")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
	}
}

func (c *dfaController) WriteFile(ctx *gin.Context) {
	//c.contractService.WriteFile()
	var request dto.WriteFileRequest
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to upload a file",
		})
	}
	err := ctx.ShouldBind(request)
	if err != nil {
		errFunc(err)
		return
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		errFunc(err)
		return
	}
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
		dir, _ := ioutil.ReadDir("/tmp")
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
	tx, err := c.contract.WriteFile(priv, request.ID, hashAddr, data)
	if err != nil {
		errFunc(err)
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "succeeded in writing file: " + request.ID,
		"transaction": tx,
	})
}

func (c *dfaController) ShareFile(ctx *gin.Context) {
	var request dto.ShareFileRequest
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to share a file",
		})
	}
	err := ctx.ShouldBind(request)
	if err != nil {
		errFunc(err)
		return
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		errFunc(err)
		return
	}
	permission, err := dto.String2Permission(request.Permission)
	if err != nil {
		errFunc(err)
		return
	}
	tx, err := c.contract.ShareFile(priv, request.To, request.ID, entity.Permission(permission))
	if err != nil {
		errFunc(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"trasaction": tx,
	})
}

func (c *dfaController) QueryFile(ctx *gin.Context) {
	var request dto.QueryFileRequest
	err := ctx.ShouldBind(request)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to request file",
		})
		return
	}

	data, err := c.contract.QueryFile(request.ID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": err,
		})
		return
	}
	view := dto.Data2View(data)
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"data":   view,
	})

}

func (c *dfaController) QueryAllFiles(ctx *gin.Context) {
	data, err := c.contract.QueryAllFiles()
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to query files",
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

func (c *dfaController) BuyAFile(ctx *gin.Context) {
	var request dto.BuyAFileRequest
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to buy a file",
		})
	}
	err := ctx.ShouldBind(request)
	if err != nil {
		errFunc(err)
		return
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		errFunc(err)
		return
	}
	tx, ok, err := c.contract.PurchaseFile(priv, request.ID)
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
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "succeeded in buying a file",
		"transaction": tx,
	})
}

func (c *dfaController) GetAddress(ctx *gin.Context) {
	dataAddr := c.contract.GetDataContractAddress()
	userAddr := c.contract.GetUserContractAddress()
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":   "smart contract address on Etherueum",
		"addr1": "file sharing contract address:\n" + dataAddr,
		"addr2": "user & ForFortoken contract address:\n" + userAddr,
	})

}

func (c *dfaController) GetAllowance(ctx *gin.Context) {
	var request dto.TokenIdentity
	errFunc := func(err error) {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to query allowance",
		})
	}
	err := ctx.ShouldBind(request)
	if err != nil {
		errFunc(err)
		return
	}
	priv, err := service.ParseToken(request.Token)
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
