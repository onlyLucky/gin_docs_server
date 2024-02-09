package models

// 文档数据表（统计数据）
type DocDataModel struct {
	Model
	DocID uint `gorm:"column:docID" json:"docID"`
	DocTitle string `gorm:"column:docTitle;comment:文档标题" json:"docTitle"`
	DiggCount int `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`
	LookCount int `gorm:"comment:阅读量;column:lookCount" json:"lookCount"`
	CollCount int `gorm:"comment:收藏量;column:collCount" json:"collCount"`
}