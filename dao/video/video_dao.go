package d_video

import (
	"apiproject/dao"
)

var VideoDao = &videoDao{}

type videoDao struct {
	dao.BaseDao
}
