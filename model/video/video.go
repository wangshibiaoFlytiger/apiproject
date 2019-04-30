package m_video

import (
	"time"
)

type Video struct {
	ID         string `gorm:"size:255"`
	Title      string `gorm:"size:255"`
	SiteId     string `gorm:"size:255" form:"siteId" binding:"required"`
	CreateTime time.Time
	UpdateTime time.Time
}

/**
设置表名
*/
func (this Video) TableName() string {
	return "video"
}
