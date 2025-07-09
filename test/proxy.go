package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 代理服务配置
const (
	origin = "http://localhost:20020"
	prefix = "/permissions"
	openId = "1hendj97f"
)

func main() {
	app := gin.Default()

	// 代理该服务
	power := app.Group(prefix)
	// power.Use(middleware.RoleVerify("0")) // 设置只有管理员有权限访问该地址
	power.Any("/*path", func(ctx *gin.Context) {
		targetURL, _ := url.Parse(origin)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		ctx.Request.URL.Scheme = targetURL.Scheme
		ctx.Request.URL.Host = targetURL.Host
		ctx.Request.Host = targetURL.Host
		ctx.Request.Header.Set("Open-Id", openId) // 如 Open-Id 与 config.yml 中设置不一致，接口则会返回 401

		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	api := app.Group("/busine")
	// 访问每个业务接口前，判断是否有权限
	api.Use(func(ctx *gin.Context) {
		roleId := "0"
		method := ctx.Request.Method
		url := ctx.Request.URL.String()

		reqUrl := origin + prefix + "/v1/api/interface/authority?roleId=" + roleId + "&method=" + method + "&url=" + url
		req, _ := http.NewRequest("GET", reqUrl, nil)
		req.Header.Set("Open-Id", openId)
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "服务异常：" + err.Error(),
			})
			ctx.Abort()
			return
		}

		body, _ := io.ReadAll(res.Body)
		data := struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
		}{}
		json.Unmarshal(body, &data)

		if data.Code == 200 {
			ctx.Next()
		} else {
			ctx.JSON(200, gin.H{
				"code": 404,
				"msg":  "无权限访问该地址：" + url,
			})
			ctx.Abort()
		}
	})

	// 业务接口
	api.GET("/test", func(ctx *gin.Context) {
		// code...
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "业务接口访问权限通过",
		})
	})

	app.Run(":8080")
}
