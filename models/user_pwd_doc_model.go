package models

// 用户文档密码表
type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:userID" json:"userID"`
	DocID uint `gorm:"column:docID" json:"docID"`
}