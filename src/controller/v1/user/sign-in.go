package user

import (
	"server/configs"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	type Params struct {
		Username string `form:"username" binding:"required"`
		Passwrod string `form:"passwrod" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	if params.Username == configs.Config.Username && params.Passwrod == configs.Config.Passwrod {
		userInfo := make(map[string]interface{})
		userInfo["username"] = params.Username
		userInfo["passwrod"] = params.Passwrod
		tokenString := service.Jwt.Publish(userInfo)
		service.State.SuccessData(tokenString)
	} else {
		service.State.ErrorCustom("用户名或密码错误")
	}

}
