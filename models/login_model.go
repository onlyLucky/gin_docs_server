package models

// 登录记录表 
type LoginModel struct {
	Model
	UserID uint `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	NickName string `gorm:"column:nickName;size:36;comment:昵称" json:"nickName"`
	IP       string `gorm:"column:ip;size:16;comment:ip" json:"ip"`
	Address  string `gorm:"column:address;size:64;comment:地址" json:"address"`
	UA       string `gorm:"column:ua;size:256;comment:ua" json:"ua"`
	Token    string `gorm:"column:token;size:64;comment:其他平台的唯一id" json:"-"`
	Device string `gorm:"column:device;size:256;comment:登录设备" json:"device"`
}