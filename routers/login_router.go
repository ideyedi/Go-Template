package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello_route(router *gin.Engine) *gin.Engine {
	loginRoute := router.Group("/login")
	{
		loginRoute.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Hello": "World!"})
		})
	}
	return router
}

func Test() string {
	var hello string
	hello = "hello world!"
	return hello
}
