package handlers

import "strings"

func  retain_YMD(date *string)  {
    *date = strings.Split(*date, "T")[0]
}
