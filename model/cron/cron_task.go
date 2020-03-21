package m_cron

import (
	"apiproject/model"
	"github.com/robfig/cron/v3"
)

type CronTask struct {
	model.BaseModel
	Type string `json:"type"`
	//该ID由任务管理器生成, 作为任务的唯一标识
	EntryId cron.EntryID `json:"entryId"`
	Title   string       `gorm:"size:255" form:"title" json:"title"`
	Spec    string       `json:"spec"`
	Remark  string       `json:"remark"`
	//状态:1启用,2禁用
	Status int `json:"status"`
}

/**
设置表名
*/
func (this *CronTask) TableName() string {
	return "cron_task"
}
