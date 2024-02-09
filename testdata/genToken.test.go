package main

import (
	"fmt"
	"gin_docs_server/utils/jwts"
)

func main() {
	token,err := jwts.GenToken(jwts.JwyPayLoad{
		RoleID: 2,
		UserID: 1,
		NickName: "fff",
	})
	fmt.Println(token,err)
}