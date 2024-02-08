package flags

import (
	"gin_docs_server/global"
)

// 修改程序运行端口
func Port(port int){
	global.Config.System.Port = port
}