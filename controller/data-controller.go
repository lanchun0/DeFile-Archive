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
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to parse token",
		})
		return
	}
	permission, err := dto.String2Permission(request.PermissionLevel)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    err,
		})
		return
	}

	// upload the file
	f, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to request file",
		})
		return
	}
	// store the file
	err = ctx.SaveUploadedFile(f, "./tmp/"+f.Filename)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to save the file",
		})
		return
	}
	hashAddr, err := c.ipfsService.UploadFile("./tmp/" + f.Filename)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to upload the file",
		})
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
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to save the file",
		})
		return
	}
	//
	// fill in the table and submit to blockchain
	id, _ := general.MakeUUID()
	time := general.GenerateTimeStamp()
	digest, err := general.MakeHashDigest(file)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to create digest",
		})
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
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "falied to create the file on blockchain",
		})
	}
	tx2, err := c.contract.WriteFile(priv, id, data.MeteData)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "falied to write the file on blockchain",
		})
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
}

func (c *dfaController) WriteFile(ctx *gin.Context) {
	//c.contractService.WriteFile()
	var request dto.WriteFileRequest
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    err,
		})
		return
	}
	_, err = service.ParseToken(request.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    "token expired",
		})
		return
	}

}

func (c *dfaController) ShareFile(ctx *gin.Context) {
	var request dto.ShareFileRequest
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    err,
		})
		return
	}
	priv, err := service.ParseToken(request.Token)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    "token expired",
		})
		return
	}
	permission, err := dto.String2Permission(request.Permission)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    err,
		})
		return
	}
	tx, err := c.contract.ShareFile(priv, request.To, request.ID, entity.Permission(permission))
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"status": http.StatusNotAcceptable,
			"msg":    err,
		})
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
		"status": http.StatusOK,
		"data":   view,
	})

}

func (c *dfaController) QueryAllFiles(ctx *gin.Context) {
	data, err := c.contract.QueryAllFiles()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
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

func (c *dfaController) PurchaseFile(ctx *gin.Context) {

}

func (c *dfaController) GetAddress(ctx *gin.Context) {

}

func (c *dfaController) GetAllowance(ctx *gin.Context) {

}
