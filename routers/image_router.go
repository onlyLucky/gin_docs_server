package routers

import (
	"gin_docs_server/api"
	"gin_docs_server/middleware"
)

func (router RouterGroup) ImageRouter() {
	app := api.App.ImageApi
	router.POST("/uploadAvatar", middleware.JwtAdmin(), app.ImageUploadView)       // 上传头像
	router.GET("/uploadAvatarList", middleware.JwtAdmin(), app.ImageListView)      // 上传头像列表
	router.DELETE("/avatarList", middleware.JwtAdmin(), app.ImageRemoveView) // 上传头像列表删除
}
