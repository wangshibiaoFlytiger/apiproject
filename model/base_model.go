package model

import (
	"time"
)

type BaseModel struct {
	ID        string     `gorm:"size:255" json:"id"`
	CreatedAt time.Time  `json:"createTime"`
	UpdatedAt time.Time  `json:"updateTime"`
	DeletedAt *time.Time `json:"deleteTime"`
}
