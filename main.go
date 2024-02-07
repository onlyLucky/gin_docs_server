package main

import "gin_docs_server/routers"

func main() {
	route := routers.InitRouter();
	route.Run("8080")
}