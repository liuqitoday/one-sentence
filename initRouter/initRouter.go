package initRouter

import (
	"github.com/gin-gonic/gin"
	"one-sentence/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("template/**")
	group := router.Group("/sentence")
	group.Static("/statics","./statics")
	group.GET("/", handler.Index)
	return router
}
