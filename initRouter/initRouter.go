package initRouter

import (
	"github.com/gin-gonic/gin"
	"one-sentence/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("template/**")
	router.Static("/statics","./statics")
	router.GET("/", handler.Index)
	return router
}
