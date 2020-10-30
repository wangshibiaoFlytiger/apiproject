package model

import (
	"apiproject/entity"
	"github.com/bwmarrin/snowflake"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type BaseModel struct {
	ID        snowflake.ID     `form:"id" binding:"-" json:"id"`
	CreatedAt *entity.JsonTime `json:"createTime" swaggertype:"string" example:"2020-04-11 03:15:52"`
	UpdatedAt *entity.JsonTime `json:"updateTime" swaggertype:"string" example:"2020-04-11 03:15:52"`
	DeletedAt *entity.JsonTime `json:"deleteTime" swaggertype:"string" example:"2020-04-11 03:15:52"`
}
