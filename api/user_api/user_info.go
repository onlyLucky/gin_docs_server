package user_api

import (
	"fmt"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	models.UserModel
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

func (UserApi) UserInfoView(c *gin.Context) {
	// 从token中获取用户的id
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	fmt.Println(user.RoleModel)
	info := UserInfoResponse{
		UserModel: *user,
		UserName:  user.UserName,
		Role:      user.RoleModel.Title,
	}
	res.OKWithData(info, c)
	return
}
