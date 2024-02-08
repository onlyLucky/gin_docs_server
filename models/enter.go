package models

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // 主键
	CreatedAt time.Time `gorm:"column:createdAt;comment:创建时间" json:"createdAt" `           // 创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt" ` // 更新时间
}