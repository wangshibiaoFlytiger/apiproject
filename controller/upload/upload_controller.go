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
// @Param File formData file true "文件"
// @Success 200 {object} gin.H
// @Router /api/upload/uploadFile [post]
/************************end swagger api定义注解 **************/
func UploadFile(ctx *gin.Context) {
	para := ReqUploadFile{}
	if err := ctx.ShouldBind(&para); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("para", para))

	filePath := "/data2/uploadfile/" + para.File.Filename
	if err := util.UploadFile(ctx, para.File, filePath); err != nil {
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
