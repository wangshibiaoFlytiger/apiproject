package d_video

import (
	"apiproject/dao"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

var VideoDao = &videoDao{}

type videoDao struct {
	dao.BaseDao
}
