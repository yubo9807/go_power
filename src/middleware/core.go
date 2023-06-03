package middleware

import (
	"github.com/gin-gonic/gin"
)

func Core(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
	}
}
