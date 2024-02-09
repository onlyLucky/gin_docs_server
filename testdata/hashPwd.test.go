package main

import (
	"fmt"
	"gin_docs_server/utils/pwd"
)

func main() {
	hash := pwd.HashPwd("123456")
	hash1 := pwd.HashPwd("123456")
	fmt.Println(hash) // $2a$04$6UwyMqncUPQpvLecejoSU.hVLShU8ZiJ7q6yK7lYfrMPWj77MElaS 
	fmt.Println(hash1) // $2a$04$NnUmm9Bxqb5NWLaD.J.LLOdvmN0O3K0HHW1HIXvzYrTw9RYiz.cjC
	ok := pwd.CheckPwd("$2a$04$6UwyMqncUPQpvLecejoSU.hVLShU8ZiJ7q6yK7lYfrMPWj77MElaS","123456")
	fmt.Println("hash ok",ok)
	ok = pwd.CheckPwd("$2a$04$NnUmm9Bxqb5NWLaD.J.LLOdvmN0O3K0HHW1HIXvzYrTw9RYiz.cjC","123456")
	fmt.Println("hash1 ok",ok)
}