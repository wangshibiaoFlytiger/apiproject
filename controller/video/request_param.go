package c_video

import "github.com/bwmarrin/snowflake"

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type ReqPage struct {
	//当前页码
	PageNo int `form:"pageNo,default=1" json:"pageNo" example:"1"`
	//每页大小
	PageSize int `form:"pageSize,default=10" json:"pageSize" example:"10"`
}

type ReqFindVideoByWhere struct {
	TitleLike string `form:"titleLike" json:"titleLike"`
	SiteId    string `form:"siteId" binding:"required" json:"siteId"`
}

type ReqVideoId struct {
	ID snowflake.ID `form:"id" binding:"required" json:"id" example:"1241316493350305792"`
}
