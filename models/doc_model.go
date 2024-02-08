package models

// 文档表
type DocModel struct {
	Model
	Title string `gorm:"comment:文档标题" json:"title"`
	Content string `gorm:"comment:文档内容" json:"-"`
	DiggCount int `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`
	LookCount int `gorm:"comment:阅读量;column:lookCount" json:"lookCount"`
	Key string `gorm:"comment:key;not null;unique" json:"key"`
	ParentID *uint `gorm:"comment:父文档id;column:parentID" json:"parentID"`
	ParentModel *DocModel `gorm:"foreignKey:ParentID" json:"-"` // 父文档
	FreeContent string `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"` // 不是实体表
}

/* 
joinForeignKey 连接的主键id
JoinReferences 关联的主键id 
*/