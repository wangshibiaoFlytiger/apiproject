package c_upload

import (
	"apiproject/log"
	"apiproject/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/**
上传文件
*/

/************************start swagger api定义注解 **************/
// @Summary 上传文件
// @Description 上传文件
// @Tags 上传
// @Accept  json
// @Produce  json
// @Param file formData file true "文件"
// @Success 200 {object} gin.H
// @Router /api/upload/uploadFile [post]
/************************end swagger api定义注解 **************/
func UploadFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		log.Logger.Error("上传文件, 异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	filePath := "/data2/uploadfile/" + fileHeader.Filename
	util.CreateFileDir(filePath)
	if err = ctx.SaveUploadedFile(fileHeader, filePath); err != nil {
		log.Logger.Error("上传文件, 保存文件, 异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
	return
}
