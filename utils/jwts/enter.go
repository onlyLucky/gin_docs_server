package jwts

import (
	"gin_docs_server/global"
	"gin_docs_server/models"

	"github.com/dgrijalva/jwt-go/v4"
)

type JwyPayLoad struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
}

var Secret []byte

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}

// 根据token获取当前用户信息
func (c CustomClaims) GetUser() (user *models.UserModel, err error) {
	// 函数返回为models.UserModel的指针，需要user变量初始化
	var newUser models.UserModel
	// user = new(models.UserModel) user变量初始化
	err = global.DB.Take(&newUser, c.UserID).Error
	if err == nil {
		user = &newUser // 如果获取成功，将新建的实例赋值给返回值
	}
	return
}
