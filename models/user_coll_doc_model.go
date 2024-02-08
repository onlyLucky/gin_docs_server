package models

// 用户收藏文档表 虚拟表
type UserCollDocModel struct {
	Model
	DocID uint `gorm:"column:"docID" json:"docID"`
	DocModel DocModel `gorm:"foreignKey:DocID"`
	UserID uint `gorm:"column:"userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}