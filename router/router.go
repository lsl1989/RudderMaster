package router

import (
	"RudderMaster/apis/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	registerRouter(r)
	return r
}

func registerRouter(r *gin.Engine) {
	// 健康检查
	health(r)
	// 登录
	r.POST("/login", auth.Login)
}

func health(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
