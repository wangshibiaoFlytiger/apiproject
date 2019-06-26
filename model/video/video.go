package m_video

import (
	"apiproject/model"
)

type Video struct {
	model.BaseModel
	Title  string `gorm:"size:255" form:"title" json:"title"`
	SiteId string `gorm:"size:255" form:"siteId" binding:"required" json:"siteId"`
}

/**
设置表名
*/
func (this *Video) TableName() string {
	return "video"
}
