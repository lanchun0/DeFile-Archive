package controller

import (
	"dfa/dto"
	"dfa/entity"
	"dfa/general"
	"dfa/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *contractController) UploadFile(ctx *gin.Context) {
	var request dto.FileCreateRequest
	ctx.ShouldBind(&request)
	pub, priv, err := service.ParseToken(request.Token)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to parse token",
		})
		return
	}

	// upload the file
	file, header, err := ctx.Request.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to request file",
		})
		return
	}
	// store the file
	defer file.Close()
	//
	// fill in the table and submit to blockchain
	id, _ := general.MakeUUID()
	time := general.GenerateTimeStamp()
	hash, err := general.MakeHashDigest(file)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to create digest",
		})
		return
	}
	hashBytes := general.String2Bytes(hash)
	signature, err := general.MakeSignature(priv, hashBytes)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "failed to make signature",
		})
		return
	}
	permission, err := dto.String2Permission(request.Permission)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"msg": err,
		})
		return
	}
	data := entity.Data{
		ID:              id,
		HashDigest:      hash,
		Owner:           pub,
		PermissionLevel: permission,
		Signature:       signature,
		MeteData: entity.MeteData{
			FileName:  header.Filename,
			Size:      header.Size,
			TimeStamp: time,
		},
	}
	tx, err := c.contractService.UploadFile(data)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "falied to upload file",
		})
	}
	view := dto.Data2View(data)
	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "success uploading",
		"data":        view,
		"transaction": tx,
	})

}

func (c *contractController) DownloadFile(ctx *gin.Context) {
	//c.contractService.DownloadFile()
}

func (c *contractController) WriteFile(ctx *gin.Context) {
	//c.contractService.WriteFile()
}

func (c *contractController) ShareFile(ctx *gin.Context) {
	//c.contractService.ShareFile()
}

func (c *contractController) QueryFile(ctx *gin.Context) {
	//c.contractService.QueryFile()
}

func (c *contractController) QueryAllFiles(ctx *gin.Context) {
	//c.contractService.QueryAllFiles()
}
