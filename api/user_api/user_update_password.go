package user_api

import (
	"gin_docs_server/global"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/jwts"
	"gin_docs_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

// 更改用户密码
type UserUpdatePasswordRequest struct {
	OldPwd   string `json:"oldPwd" binding:"required" label:"原来的密码"`
	Password string `json:"password" binding:"required" label:"新密码"`
}

func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UserUpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 从token中获取用户的id
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMsg("原来的密码不正确", c)
		return
	}

	hashPwd := pwd.HashPwd(cr.Password)
	global.DB.Model(user).Update("password", hashPwd)

	res.OKWithMsg("用户密码修改成功", c)
	return

}
