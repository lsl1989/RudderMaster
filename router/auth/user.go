package auth

import (
	"RudderMaster/apis/auth"
	"github.com/gin-gonic/gin"
)

func RegisterUserApi(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	userGroup.GET("/:id", auth.UserInfoApi)
}
