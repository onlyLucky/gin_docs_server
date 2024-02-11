package user_api

import (
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

// 更新用户信息
type UserUpdateInfoRequest struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

func (UserApi) UserUpdateInfoView(c *gin.Context) {
	var cr UserUpdateInfoRequest
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

	if cr.Avatar != "" {
		var image models.ImageModel
		err := global.DB.Debug().Take(&image, "userID = ? and path = ?", claims.UserID, cr.Avatar[1:]).Error
		if err != nil {
			res.FailWithMsg("用户头像不存在", c)
			return
		}
	}

	if !(cr.NickName == "" && cr.Avatar == "") {
		global.DB.Model(user).Updates(models.UserModel{
			Avatar:   cr.Avatar,
			NickName: cr.NickName,
		})
	}

	res.OKWithMsg("用户详情修改成功", c)
	return
}
