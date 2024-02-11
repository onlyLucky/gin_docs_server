package user_api

import (
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

// 用户更新接口
type UserUpdateRequest struct {
	ID uint `json:"id" binding:"required"`
	Password string `json:"password"`             // 密码
	NickName string `json:"nickName"`                                // 昵称
	RoleID   uint   `json:"roleID"`               // 角色id
}

// @Tags 用户管理
// @Summary 用户更新
// @Description 用户更新
// @Param data body UserUpdateRequest true "参数"
// @Param token header string true "token"
// @Router /api/user [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateView(c *gin.Context){
	var cr UserUpdateRequest 
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误",c)
		return 
	}

	// 查询用户信息
	var user models.UserModel
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("用户名不存在",c)
		return
	}

	if cr.Password != ""{
		cr.Password = pwd.HashPwd(cr.Password)
	}

	if cr.RoleID != 0{
		var role models.RoleModel
		err = global.DB.Take(&role, cr.RoleID).Error
		if err != nil{
			global.Log.Error(err)
			res.FailWithMsg("角色不存在",c)
			return
		}

	}

	err = global.DB.Model(&user).Updates(models.UserModel{
		Password: cr.Password,
		NickName: cr.NickName,
		RoleID: cr.RoleID,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户信息更新失败",c)
		return
	}
	res.OKWithMsg("success",c)
	return
}