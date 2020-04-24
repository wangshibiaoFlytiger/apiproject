package c_cron

import (
	"apiproject/log"
	m_cron "apiproject/model/cron"
	s_cron "apiproject/service/cron"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/**
添加定时任务
*/

/************************start swagger api定义注解 **************/
// @Summary 添加定时任务
// @Description 添加定时任务
// @Tags 定时任务
// @Accept  json
// @Produce  json
// @Param cronTask body ReqCronTask true "json对象"
// @Success 200 {object} gin.H
// @Router /api/cronTask/addCronTask [post]
/************************end swagger api定义注解 **************/
func AddCronTask(ctx *gin.Context) {
	para := ReqCronTask{}
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

	para.Status = 1
	if err := s_cron.CronTaskService.AddCronTask(&m_cron.CronTask{
		Type:   para.Type,
		Title:  para.Title,
		Spec:   para.Spec,
		Remark: para.Remark,
		Status: para.Status,
	}); err != nil {
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
}

/**
删除定时任务
*/

/************************start swagger api定义注解 **************/
// @Summary 删除定时任务
// @Description 删除定时任务
// @Tags 定时任务
// @Accept  json
// @Produce  json
// @Param cronTaskId query ReqCronTaskId true "查询参数"
// @Success 200 {object} gin.H
// @Router /api/cronTask/deleteCronTask [delete]
/************************end swagger api定义注解 **************/
func DeleteCronTask(ctx *gin.Context) {
	para := ReqCronTaskId{}
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

	if err := s_cron.CronTaskService.DeleteCronTask(para.ID); err != nil {
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
}

/**
启用定时任务
*/

/************************start swagger api定义注解 **************/
// @Summary 启用定时任务
// @Description 启用定时任务
// @Tags 定时任务
// @Accept  json
// @Produce  json
// @Param cronTaskId body ReqCronTaskId true "json对象"
// @Success 200 {object} gin.H
// @Router /api/cronTask/enableCronTask [post]
/************************end swagger api定义注解 **************/
func EnableCronTask(ctx *gin.Context) {
	para := ReqCronTaskId{}
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

	if err := s_cron.CronTaskService.EnableCronTask(para.ID); err != nil {
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
}

/**
禁用定时任务
*/

/************************start swagger api定义注解 **************/
// @Summary 禁用定时任务
// @Description 禁用定时任务
// @Tags 定时任务
// @Accept  json
// @Produce  json
// @Param cronTaskId body ReqCronTaskId true "json对象"
// @Success 200 {object} gin.H
// @Router /api/cronTask/disableCronTask [post]
/************************end swagger api定义注解 **************/
func DisableCronTask(ctx *gin.Context) {
	para := ReqCronTaskId{}
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

	if err := s_cron.CronTaskService.DisableCronTask(para.ID); err != nil {
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
}

/**
查询定时任务列表
*/

/************************start swagger api定义注解 **************/
// @Summary 查询定时任务列表
// @Description 查询定时任务列表
// @Tags 定时任务
// @Accept  json
// @Produce  json
// @Success 200 {object} gin.H
// @Router /api/cronTask/findCronTaskList [get]
/************************end swagger api定义注解 **************/
func FindCronTaskList(ctx *gin.Context) {
	cronTaskList, err := s_cron.CronTaskService.FindCronTaskList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": cronTaskList,
	})
}
