package utils

import (
	"strconv"
	"strings"
	"time"
)

var CN, _ = time.LoadLocation("Asia/Shanghai")

const date_layout = "2006-01-02"

func China_date(t time.Time) string {
	return t.In(CN).Format(date_layout)
}

func Parse_order_number(s string) int32 {
	i, _ := strconv.Atoi(strings.Fields(s)[1])
	return int32(i)
}

func Parse_business_number(s string) int32 {
	i, _ := strconv.Atoi(s[3:])
	return int32(i)
}

// 前提：lhs >= rhs
func Sub_day(lhs time.Time, rhs time.Time) int32 {
	return int32(lhs.Sub(rhs).Hours()/24) + 1 // 把 lhs 当日也算进去
}
