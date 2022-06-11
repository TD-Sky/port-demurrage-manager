package utils

import (
	"strconv"
	"strings"
	"time"
)

var CN, _ = time.LoadLocation("Asia/Shanghai")

// 转换到中国时区
func CST(t *time.Time) time.Time {
	return t.In(CN)
}

func Parse_order_number(s string) int32 {
	i, _ := strconv.Atoi(strings.Fields(s)[1])
	return int32(i)
}

func Parse_business_number(s string) int32 {
	i, _ := strconv.Atoi(s[3:])
	return int32(i)
}
