package models

type UserModel struct {
	Model
	UserName string `gorm:"column:userName;size:36;unique;not null" json:"userName"` // 用户名
	Password string `gorm:"column:password;size:128" json:"password"` // 密码
	Avatar   string `gorm:"column:avatar;size:256" json:"avatar"`   // 头像
	NickName string `gorm:"column:nickName;size:36" json:"nickName"` // 昵称
	Email    string `gorm:"column:email;size:128" json:"email"`    // 邮箱
	Token    string `gorm:"column:token;size:64" json:"token"`    // 其他平台的唯一id
	IP       string `gorm:"column:ip;size:16" json:"ip"`       // ip
	Address  string `gorm:"column:address;size:64" json:"address"`  // 地址
	RoleID   uint   `gorm:"column:roleID" json:"roleID"`   // 用户对应的角色id
}