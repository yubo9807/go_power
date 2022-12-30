package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"server/src/utils"
	"strings"
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

func Logs(ctx *gin.Context) {
	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer

	ctx.Next()

	response := writer.body.String()
	LogsWrite(ctx, "\nResponse: "+response)
}

func LogsWrite(ctx *gin.Context, append string) {
	filename := utils.DateFormater(time.Now(), "YYYY-MM-DD")

	logSrc := LogsGetSrc("logs/" + filename + ".log")
	log.SetOutput(logSrc)
	log.SetPrefix("\n")

	header, _ := json.Marshal(ctx.Request.Header)
	headerStr := strings.ReplaceAll(strings.ReplaceAll(string(header), "\\", ""), "\"\"", "\"")
	data, _ := json.Marshal(ctx.Request.Body)

	log.Println(
		State.RunTime,
		ctx.ClientIP(),
		"->",
		ctx.Request.Host,
		ctx.Request.Method,
		ctx.Request.RequestURI,
		"\nData: "+string(data),
		"\nHeaders: "+headerStr,
		append,
	)
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
