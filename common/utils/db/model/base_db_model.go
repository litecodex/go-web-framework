package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseDBModel struct {
	Id         int64 `json:"id"`
	CreateTime int64 `gorm:"type:bigint" json:"creatTime,omitempty"`
	UpdateTime int64 `gorm:"type:bigint" json:"updateTime,omitempty"`
}

// BeforeCreate 在创建记录之前设置 CreateTime 和 UpdateTime
func (base *BaseDBModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreateTime = time.Now().Unix()
	base.UpdateTime = time.Now().Unix()
	return
}

// BeforeUpdate 在更新记录之前设置 UpdateTime
func (base *BaseDBModel) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdateTime = time.Now().Unix()
	return
}
