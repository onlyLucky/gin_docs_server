package middleware

import (
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

// 验证用户是否为超级管理员
func JwtAdmin() func(c *gin.Context){
	return func(c *gin.Context){
		token := c.Request.Header.Get("token")
		if token == ""{
			res.FailWithMsg("请求未携带token",c)
			// 拦截响应
			c.Abort()
			return
		}
		claims,err:=jwts.ParseToken(token)
		if err != nil{
			// token错误
			res.FailWithMsg("请求未授权",c)
			// 拦截响应
			c.Abort()
			return
		}
		if claims.RoleID != 1{
			res.FailWithMsg("权限错误",c)
			c.Abort()
			return
		}
		// 路由设置存储变量
		c.Set("claims",claims)
		// c.Get("claims")
	}
}