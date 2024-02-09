package main

import (
	"fmt"
	"gin_docs_server/utils/jwts"
)

func main() {
	claims,err:=jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrTmFtZSI6ImZmZiIsInJvbGVJRCI6MiwidXNlcklEIjoxLCJleHAiOjE3MDc0NjM2OTAuMTE5MjQ4LCJpc3MiOiJmZmYifQ.Ycvbyyf8Lit7HjyvgBXz7gVFFHX8vjmilFujuhSbtoA")
	fmt.Println(claims,err)
}