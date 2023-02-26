package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
