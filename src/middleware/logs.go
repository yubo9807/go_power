package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"server/src/service"
	"server/src/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写 Write([]byte) (int, error)
func (w responseWriter) Write(body []byte) (int, error) {
	w.body.Write(body)                  // 向一个 bytes.buffer 中再写一份数据
	return w.ResponseWriter.Write(body) // 完成 gin.Context.Writer.Write() 原有功能
}

func init() {
	err := os.Mkdir("logs", 0777)
	if err != nil {
		return
	}
}

var currentData string

func Logs(ctx *gin.Context) {
	data, _ := ctx.GetRawData() // body 数据只能被读一次，读完即删
	currentData = string(data)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 回写

	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer

	ctx.Next()

	// response := writer.body.String()
	LogsWrite(ctx, "")
}

func LogsWrite(ctx *gin.Context, append string) {
	filename := utils.DateFormater(time.Now(), "YYYY-MM-DD")

	logSrc := LogsGetSrc("logs/" + filename + ".log")
	log.SetOutput(logSrc)
	log.SetPrefix("\n")

	log.Println(
		service.State.RunTime,
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.RequestURI,
		utils.If(currentData == "", "", "\nbody:"+string(currentData)),
		append,
	)
	currentData = "" // 清理内存
}

func LogsGetSrc(filename string) *os.File {
	src, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logSrc := src
	if err != nil {
		src, _ := os.Create(filename)
		logSrc = src
	}
	return logSrc
}
