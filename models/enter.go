package models

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`                            // 主键
	CreatedAt time.Time `gorm:"column:createdAt;comment:创建时间" json:"createdAt" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt" ` // 更新时间
}

// 分页传参处理
type Pagination struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
	Sort  string `json:"sort" form:"sort"`
}

type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required" label:"用户ids"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" binding:"required" label:"用户id"`
}
