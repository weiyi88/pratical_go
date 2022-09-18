package router

import (
	controller "gin_demo/contorller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	g.GET("/", controller.Index)

	return g
}
