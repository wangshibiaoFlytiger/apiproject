package model

import (
	"apiproject/entity"
)

type BaseModel struct {
	ID        string           `gorm:"size:255" json:"id"`
	CreatedAt entity.JsonTime  `json:"createTime"`
	UpdatedAt entity.JsonTime  `json:"updateTime"`
	DeletedAt *entity.JsonTime `json:"deleteTime"`
}
