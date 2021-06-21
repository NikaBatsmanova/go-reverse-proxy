package controllers

import (
	"github.com/gin-gonic/gin"
	"go-reverse-proxy/project/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", handlers.GetProxy)
	return router
}
