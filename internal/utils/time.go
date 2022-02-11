package utils

import (
	"strings"
	"time"
)

/*
时间格式工具函数
*/

// CurrentTime 当前时间
func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:06")
}

// CurrentDate 当前日期
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// TransferStrToTime 字符串转时间对象
func TransferStrToTime(s string) (time.Time, error) {
	var formatString = "2006-01-02"
	if strings.Contains(s, ":") {
		formatString = "2006-01-02 15:04:06"
	}
	return time.ParseInLocation(formatString, s, time.Local)
}
