package m_video

type Video struct {
	ID    string `gorm:"size:255"`
	Title string `gorm:"size:255"`
}

/**
设置表名
*/
func (this Video) TableName() string {
	return "video"
}
