package models

import "fmt"

type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户id" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileName;size:64;comment:文件名" json:"fileName"`
	Size      int64     `gorm:"column:size;comment:文件大小，单位字节" json:"size"`
	Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`
	Hash      string    `gorm:"column:hash;size:64;comment:文件唯一hash" json:"hash"` // 用户可能上传路径一样
}

func (image ImageModel) WebPath() string {
	return fmt.Sprintf("/%s", image.Path)
}