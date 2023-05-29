package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func NumberRandom(num int) int {
	var timestamp = time.Now().UnixNano()
	rand.Seed(timestamp)
	return rand.Intn(num)
}

// 将整数转换为指定进制的字符串
func IntToBase(num int64, base int) string {
	if base < 2 || base > 36 {
		return ""
	}
	return strconv.FormatInt(int64(num), base)
}

// 将指定进制的字符串转换为整数
func BaseToInt(str string, base int) (int64, error) {
	if base < 2 || base > 36 {
		return 0, fmt.Errorf("无效的进制：%d", base)
	}
	return strconv.ParseInt(str, base, 0)
}
