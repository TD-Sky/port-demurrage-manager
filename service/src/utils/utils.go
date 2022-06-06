package utils

import (
	"strconv"
	"strings"
	"time"
)

var cn, _ = time.LoadLocation("Asia/Shanghai")

// 将时间转为中国时区的日期字符串
func Date_string(t *time.Time) string {
	return t.In(cn).Format("2006-01-02")
}

func Parse_date(s string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", s, cn)
	return t
}

func Parse_order_number(s string) int32 {
	i, _ := strconv.Atoi(strings.Fields(s)[1])
	return int32(i)
}

func Parse_business_number(s string) int32 {
	i, _ := strconv.Atoi(s[3:])
	return int32(i)
}
