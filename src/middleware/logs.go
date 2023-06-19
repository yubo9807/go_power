package middleware

import (
	"bytes"
	"fmt"
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
		fmt.Println(err.Error())
	}
}

func Logs(ctx *gin.Context) {

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
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.SetOutput(logSrc)
	log.SetPrefix("\n")

	state := service.State.GetStateStore(ctx)
	sql := ""
	lastIndex := len(state.Sqls) - 1
	for i, val := range state.Sqls {
		mark := utils.If(i == lastIndex, "└─ ", "├─ ")
		args := utils.If(val[1] == "", "", "\n  └─ "+val[1])
		sql += "\n" + mark + val[0] + args
	}

	body := string(state.Body)

	log.Println(
		state.RunTime,
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.RequestURI,
		utils.If(body == "", "", "\nbody:"+body),
		sql,
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
