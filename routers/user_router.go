package routers

import (
	"gin_docs_server/api"
	"gin_docs_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	// 添加中间件 middleware.JwtAuth
	// router.POST("/user",middleware.JwtAuth,app.UserCreateView)  // 创建用户
	// router.POST("/user",middleware.JwtAuth(),app.UserCreateView)  // 创建用户
	router.POST("/login",app.UserLoginView)	 // 用户登录
	router.POST("/user",middleware.JwtAdmin(),app.UserCreateView)  // 创建用户
	router.PUT("/user",middleware.JwtAdmin(),app.UserUpdateView)	 // 更新用户信息
	router.GET("/user",middleware.JwtAdmin(),app.UserListView)	 // 获取用户列表分页
}