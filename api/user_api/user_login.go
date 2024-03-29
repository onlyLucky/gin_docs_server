package user_api

import (
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/plugins/log_stash"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/jwts"
	"gin_docs_server/utils/pwd"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// UserLoginView 用户登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录
// @Param data body UserLoginRequest true "参数"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	// 绑定json传参
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "userName=?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserName)
		res.FailWithMsg("用户名不存在", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("用户密码错误", cr.UserName, cr.Password)
		res.FailWithMsg("用户密码错误", c)
		return
	}

	// 生成token jwt
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		UserName: user.UserName,
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
		return
	}

	// 更新lastLogin
	global.DB.Model(&user).Update("lastLogin", time.Now())
	log_stash.NewSuccessLogin(c)
	res.OKWithData(token, c)
	return
}
