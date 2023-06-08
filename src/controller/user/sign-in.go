package user

import (
	"server/configs"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	type Params struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	if params.Username == configs.Config.Username && params.Password == configs.Config.Password {
		userInfo := make(map[string]interface{})
		userInfo["username"] = params.Username
		tokenString := service.Jwt.Publish(userInfo)
		service.Jwt.StorageSetToken(params.Username, tokenString)
		service.State.SuccessData(ctx, tokenString)
	} else {
		service.State.ErrorCustom(ctx, "用户名或密码错误")
	}

}
