package routers

import "github.com/gin-gonic/gin"

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	return router
}