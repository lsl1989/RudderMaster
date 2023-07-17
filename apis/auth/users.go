package auth

import (
	svcAuth "RudderMaster/service/auth"
	"RudderMaster/utils/response"
	"github.com/gin-gonic/gin"
)

func UserInfoApi(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := svcAuth.GetUserInfo(userId)
	if err != nil {
		response.ResWithStatus(ctx, response.UserNotExistsMsg)
		return
	}
	response.Data(ctx, user)
}
