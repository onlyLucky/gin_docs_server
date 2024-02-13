package main

import (
	"fmt"
	"gin_docs_server/utils/hash"
)

func main() {
	hashString := hash.MD5([]byte("123456"))
	fmt.Println(hashString)
}
