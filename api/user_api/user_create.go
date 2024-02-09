package user_api

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`// 用户名
	Password string `json:"password" binding:"required"`// 密码
	NickName string `json:"nickName"`// 昵称
	RoleID uint `json:"roleID" binding:"required"`// 角色id
}

func (UserApi) UserCreateView(c *gin.Context){
	// 用来储存用户传参
	var cr UserCreateRequest 
	// 绑定json传参
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err,c)
		return 
	}
	// 数据库操作
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名已存在",c)
		return
	}

	if cr.NickName == ""{
		// 昵称不存在
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		cr.NickName = fmt.Sprintf("用户_%d",maxID+1)
	}

	// 判断角色id
	var role models.RoleModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil{
		global.Log.Error(err)
		res.FailWithMsg("角色不存在",c)
		return
	}

	// 创建新用户
	err = global.DB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password:pwd.HashPwd(cr.Password),
		NickName:cr.NickName,
		IP:c.RemoteIP(),
		RoleID:cr.RoleID,
	}).Error
	if err != nil{
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败",c)
		return
	}
	res.OKWithMsg("success",c)
	return
}