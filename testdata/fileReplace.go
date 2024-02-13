package main

import (
	"fmt"
	"gin_docs_server/api/image_api"
)

func main() {
	newPath := image_api.ReplaceFileName("3.45.png")
	fmt.Println(newPath)
}
