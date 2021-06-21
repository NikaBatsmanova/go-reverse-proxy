package main

import (
	"go-reverse-proxy/project/cache"
	"go-reverse-proxy/project/controllers"
	"go-reverse-proxy/project/env"
	"os"
)

func main() {
	cache.InitCache()
	env.Init()
	r := controllers.InitRouter()
	port := os.Getenv("port")
	r.Run("localhost:" + port)
}
