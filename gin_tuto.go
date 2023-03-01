package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routers "go.template/routers"
	_ "log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routers.Hello_route(r)

	fmt.Println(routers.Test())
	// 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
	r.Run()
}
