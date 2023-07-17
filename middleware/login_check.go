package middleware

import (
	"RudderMaster/utils/jwtauth"
	"RudderMaster/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization-Token")
		if token == "" {
			response.ResWithStatus(context, response.LoginMissingMsg)
			context.Abort()
			return
		}
		claims, err := jwtauth.ValidToken(token)
		if err != nil {
			response.ResWithStatus(context, response.TokenCheckFailedMsg.AddDetail(err.Error()))
			context.Abort()
			return
		}
		fmt.Println(claims)
		context.Next()
	}
}
