package video

type VideoModel struct {
	ID    string `gorm:"size:255"`
	Title string `gorm:"size:255"`
}

/**
设置表名
*/
func (this VideoModel) TableName() string {
	return "video"
}
