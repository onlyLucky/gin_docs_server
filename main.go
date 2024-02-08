package main

import (
	"gin_docs_server/core"
	"gin_docs_server/flags"
	"gin_docs_server/global"
	"gin_docs_server/routers"
)

func main() {
	global.Log = core.InitLogger(core.LogRequest{
		LogPath: "logs",
		AppName: "gvd",
	})
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis()

	option:=flags.Parse()
	if option.Run(){
		return
	}
	
	addr := global.Config.System.GetAddr()
	route := routers.InitRouter();
	route.Run(addr)
}