package routers

import "gin_docs_server/api"

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("/user",app.UserCreateView)
}