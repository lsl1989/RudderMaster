package auth

import (
	formAuth "RudderMaster/forms/auth"
	svcAuth "RudderMaster/service/auth"
	"RudderMaster/utils/jwtauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var form formAuth.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 4000, "message": err.Error()})
		return
	}
	user, err := svcAuth.LoginCheck(&form)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 4001, "message": err.Error()})
		return
	}
	token, err := jwtauth.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 5000, "message": "系统错误，token生成失败!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 2000, "message": "登录成功", "token": token})
}
