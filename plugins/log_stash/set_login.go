package log_stash

import (
	"gin_docs_server/global"

	"github.com/gin-gonic/gin"
)

func NewSuccessLogin(c *gin.Context) {
	token := c.Request.Header.Get("token")
	jwyPayLoad := ParseToken(token)
	saveLoginLog(c, jwyPayLoad.UserID, jwyPayLoad.UserName, jwyPayLoad.NickName, true)
}

func saveLoginLog(c *gin.Context, userID uint, userName string, nickName string, status bool) {
	ip := c.RemoteIP()
	address := "xxx"
	global.DB.Create(&LogModel{
		IP:       ip,
		Address:  address,
		Title:    "登录成功",
		Content:  "---",
		UserID:   userID,
		UserName: userName,
		NickName: nickName,
		Status:   status,
		Type:     LoginType,
	})
}
