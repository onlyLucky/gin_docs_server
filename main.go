package main

import (
	"gin_docs_server/core"
	// _ "gin_docs_server/docs"
	"gin_docs_server/flags"
	"gin_docs_server/global"
	"gin_docs_server/routers"
)

// @title 文档项目api文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	global.Log = core.InitLogger(core.LogRequest{
		LogPath: "logs",
		AppName: "gvd",
	})
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis()

	option := flags.Parse()
	if option.Run() {
		return
	}

	addr := global.Config.System.GetAddr()
	route := routers.InitRouter()
	route.Run(addr)
}
