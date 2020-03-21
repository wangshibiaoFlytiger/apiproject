package d_video

import (
	"apiproject/dao"
)

var CronTaskDao = &cronTaskDao{}

type cronTaskDao struct {
	dao.BaseDao
}
