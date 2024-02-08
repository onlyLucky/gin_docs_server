package main

import (
	"fmt"
	"gin_docs_server/core"
	"gin_docs_server/global"
	"gin_docs_server/routers"
)

func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	fmt.Println(global.Config)
	addr := global.Config.System.GetAddr()
	route := routers.InitRouter();
	route.Run(addr)
}