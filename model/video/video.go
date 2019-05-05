package m_video

import (
	"time"
)

type Video struct {
	ID         string    `gorm:"size:255" json:"id"`
	Title      string    `gorm:"size:255" form:"title" json:"title"`
	SiteId     string    `gorm:"size:255" form:"siteId" binding:"required" json:"siteId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

/**
设置表名
*/
func (this Video) TableName() string {
	return "video"
}
