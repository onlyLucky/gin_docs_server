package log_stash

import "time"

type LogModel struct {
	ID          uint      `json:"id" gorm:"primaryKey"`                            // 主键
	CreatedAt   time.Time `gorm:"column:createdAt;comment:创建时间" json:"createdAt" ` // 创建时间
	UpdatedAt   time.Time `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt" ` // 更新时间
	IP          string    `gorm:"column:ip;comment:ip" json:"ip"`                  // ip
	Address     string    `gorm:"column:address;comment:地址" json:"address"`        // 地址
	Level       string    `gorm:"column:level;comment:等级" json:"level"`            // 等级
	Title       string    `gorm:"comment:文档标题" json:"title"`
	Content     string    `gorm:"comment:文档内容" json:"content"`
	UserID      uint      `gorm:"column:userID" json:"userID"`
	UserName    string    `gorm:"column:userName;comment:用户名" json:"-"`               // 用户名
	NickName    string    `gorm:"column:nickName;comment:昵称" json:"nickName"`         // 昵称
	ServiceName string    `gorm:"column:serviceName;comment:服务名称" json:"serviceName"` // 昵称
	Type        LogType   `gorm:"column:type;comment:类型" json:"type"`                 // 日志类型： 1.登录 2操作 3运行
	Status      bool      `gorm:"column:status;comment:状态" json:"status"`             //
}
