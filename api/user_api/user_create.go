package user_api

import (
	"gin_docs_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required"`// 用户名
	Password string `json:"password"`// 密码
	NickName string `json:"nickName"`// 昵称
	RoleID uint `json:"roleID" binding:"required"`// 角色id
}

func (UserApi) UserCreateView(c *gin.Context){
	var cr UserCreateRequest 
	// 绑定json传参
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数校验错误", c)
		return 
	}
	res.OKWithMsg("success",c)
	return
}