package controller

import (
	"dfa/entity"
	"dfa/general"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *contractController) UploadFile(ctx *gin.Context) {

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
	id, _ := general.MakeUUID()
	time := general.GenerateTimeStamp()
	data := entity.Data{
		ID:              id,
		HashDigest:      "hash",
		Owner:           "pub",
		PermissionLevel: 0x80,
		Signature:       "sign",
		MeteData: entity.MeteData{
			FileName:  header.Filename,
			Size:      int(header.Size),
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

	ctx.JSON(http.StatusAccepted, gin.H{
		"msg":         "success uploading",
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
