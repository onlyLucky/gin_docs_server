package main

import (
	"fmt"
	"gin_docs_server/core"
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
	
	val,err := global.Redis.Get("name").Result()
	
	global.Log.Infof(val,err)

	fmt.Println(global.Config)
	addr := global.Config.System.GetAddr()
	route := routers.InitRouter();
	route.Run(addr)
}