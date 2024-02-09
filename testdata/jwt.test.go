package main

import (
	"fmt"
	"gin_docs_server/core"
	"gin_docs_server/global"
	"gin_docs_server/utils/jwts"
)

func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()

	token,err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: "枫枫",
	})
	fmt.Println(token,err)
	claims, err1 := jwts.ParseToken(token)
	fmt.Println(claims, err1)
}