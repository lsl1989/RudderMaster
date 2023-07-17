package response

import "github.com/gin-gonic/gin"

type ResData struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResPage struct {
	ResData
	Total int64 `json:"total"`
}

func Data(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, ResData{
		SuccessDataMsg,
		data,
	})
}

func PagingData(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(200, ResPage{
		ResData{SuccessDataMsg, data},
		total,
	})
}

func ActionSuccess(ctx *gin.Context) {
	ctx.JSON(200, ResData{
		Status: SuccessMsg,
	})
}

func ResWithStatus(ctx *gin.Context, status Status) {
	ctx.JSON(200, ResData{
		Status: status,
	})
}
