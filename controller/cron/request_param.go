package c_cron

import "github.com/bwmarrin/snowflake"

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type ReqCronTask struct {
	//任务类型: 1种任务对应后台的一个任务处理函数
	Type   string `json:"type" binding:"required" example:"crawDouyin"`
	Title  string `json:"title" binding:"required" example:"抓取抖音视频"`
	Spec   string `json:"spec" binding:"required" example:"*/10 * * * * *"`
	Remark string `json:"remark" example:"每隔10秒执行1次"`
	//状态:1启用,2禁用, 默认启用
	Status int `json:"status" example:"1"`
}

type ReqCronTaskId struct {
	//定时任务ID
	ID snowflake.ID `form:"id" binding:"required" json:"id" example:"1241316493350305792"`
}
