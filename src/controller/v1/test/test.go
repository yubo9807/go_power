package test

import (
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func Gain(ctx *gin.Context) {
	middleware.State.Data = "hello"
}
