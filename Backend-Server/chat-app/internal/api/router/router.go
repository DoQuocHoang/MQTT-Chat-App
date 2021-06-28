package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Engine *gin.Engine

}

func (r *Router) InitRouter() {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
}

func (r *Router) SetupHandler() {

	r.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "OKKK",
		})
	})
}