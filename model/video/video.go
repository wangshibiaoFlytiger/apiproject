package m_video

import (
	"time"
)

type Video struct {
	ID         string `gorm:"size:255"`
	Title      string `gorm:"size:255"`
	CreateTime time.Time
	UpdateTime time.Time
}

/**
设置表名
*/
func (this Video) TableName() string {
	return "video"
}
