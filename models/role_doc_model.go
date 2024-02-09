package models

// 角色文档表
type RoleDocModel struct {
	Model
	RoleID uint `gorm:"column:roleID;comment:角色id" json:"roleID"`
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"-"`
	DocID uint `gorm:"column:docID;comment:文档id" json:"docID"`
	DocModel DocModel `gorm:"foreignKey:DocID" json:"-"`
	Pwd *string `gorm:"column:pwd;comment:密码配置" json:"pwd"` // null "" "有值"  优先级：角色文档密码>角色密码
	FreeContent *string `gorm:"comment:试看配置;column:freeContent" json:"freeContent"` // 试看配置 优先级： 角色文档试看 > 文档试看字段 > 文档按照特殊字符试看
	Sort int `gorm:"column:sort;comment:排序" json:"sort"` // 排序
}