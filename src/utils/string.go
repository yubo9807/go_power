package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 生成ID
func CreateID() int64 {
	nowTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	random := strconv.Itoa(NumberRandom(1000000))
	newStr := StringRandomSort(nowTime + random)
	num, _ := strconv.ParseInt(newStr, 10, 64)
	fmt.Println(num)
	return num
}

// 字符串随机排序
func StringRandomSort(str string) string {
	arr := []rune(str)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return string(arr)
}
