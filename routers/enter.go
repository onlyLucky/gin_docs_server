package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	// 添加 swag 文档
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouter := router.Group("api")
	routerGroup := RouterGroup{apiRouter}
	// 线上如果有nginx,可以省略静态资源配置
	router.Static("/uploads", "uploads")

	// /api 分组
	routerGroup.UserRouter()
	routerGroup.ImageRouter()
	return router
}
