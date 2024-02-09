package routers

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*gin.RouterGroup
}

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	apiRouter := router.Group("api")
	routerGroup := RouterGroup{apiRouter}
	// /api 分组
	routerGroup.UserRouter()
	return router
}