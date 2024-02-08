package models

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // 主键
	CreatedAt time.Time `json:"createdAt" `           // 创建时间
	UpdatedAt time.Time `json:"updatedAt" ` // 更新时间
}