package models

// 角色表
type RoleModel struct {
	Model
	Title    string `gorm:"size:16;not null;comment:角色名称" json:"title"`    // 角色名称
	Pwd      string `gorm:"column:pwa;size:64;comment:角色密码" json:"-"`        // 角色密码
	IsSystem bool   `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"` // 是否是系统角色
}