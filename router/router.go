package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	registerRouter(r)
	return r
}

func registerRouter(r *gin.Engine) {
	// 健康检查
	health(r)
}

func health(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
