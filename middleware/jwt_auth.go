package middleware

import (
	"gin_docs_server/service/common/res"
	"gin_docs_server/service/redis_service"
	"gin_docs_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

// 校验当前请求接口用户是否登录了
// 这里直接返回一个中间件，可以再调用函数中添加额外的传参
func JwtAuth() func(c *gin.Context){
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

		// 检测token是否在redis缓存失效中
		ok := redis_service.CheckLogout(token)
		if ok {
			res.FailWithMsg("token已注销",c)
			c.Abort()
			return
		}
		
		// 路由设置存储变量
		c.Set("claims",claims)
		// c.Get("claims")
	}
}

/* func JwtAuth(c *gin.Context){
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
	// 路由设置存储变量
	c.Set("claims",claims)
	// c.Get("claims")
} */