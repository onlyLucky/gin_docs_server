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
	fmt.Println(global.Config)
	global.Log.Errorf("2333")
	global.Log.Infof("1111")
	global.Log.Errorf("2222")
	addr := global.Config.System.GetAddr()
	route := routers.InitRouter();
	route.Run(addr)
}