package test

import (
	"server/src/middleware"
	"server/src/service"
	"time"

	"github.com/gin-gonic/gin"
)

// 等待 _ 毫秒后结束
func Delay(ctx *gin.Context) {
	type Params struct {
		Time int `form:"time" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorCapture(406, err.Error())
		return
	}

	date := time.Duration(params.Time * 1000000)
	time.Sleep(date)
	middleware.State.Data = "delay: " + date.String()
}
