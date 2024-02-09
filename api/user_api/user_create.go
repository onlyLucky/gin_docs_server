package user_api

import "github.com/gin-gonic/gin"

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required"`// 用户名
	Password string `json:"password"`// 密码
	NickName string `json:"nickName"`// 昵称
	RoleID string `json:"roleID" binding:"required"`// 角色id
}

func (UserApi) UserCreateView(c *gin.Context){
	var cr UserCreateRequest 
	// 绑定json传参
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		c.JSON(200, gin.H{"msg": "error"})
		return 
	}
	c.JSON(200, gin.H{"msg": "success"})
	return
}