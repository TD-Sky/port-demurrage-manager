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

func Fee(time_interval float64, mass float64) float64 {
	if time_interval <= 14 {
		return 0
	} else {
		return (time_interval - 14) * 0.5 * mass
	}
}

// 前提：lhs >= rhs
func Sub_day(lhs time.Time, rhs time.Time) float64 {
	return float64(lhs.Sub(rhs).Hours()/24) + 1 // 把 lhs 当日也算进去
}
