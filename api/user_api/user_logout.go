package user_api

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/service/common/res"
	"gin_docs_server/service/redis_service"
	"gin_docs_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

// 退出登录

// @Tags 用户管理
// @Summary 退出登录
// @Description 退出登录
// @Param token header string true "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)

	exp := claims.ExpiresAt
	fmt.Println(exp)
	diff := exp.Time.Sub(time.Now())
	fmt.Println(diff)

	err := redis_service.Logout(token, diff)
	if err != nil {
		global.Log.Errorf("UserLogoutView err:%s", err)
		res.FailWithMsg("退出登录失败", c)
		return
	}

	res.OKWithMsg("用户退出登录成功", c)
	return
}
