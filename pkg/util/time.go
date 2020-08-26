package util

import (
	"strconv"
	"time"
)

// 获取当前时间的之前或者之后的时间
func GetDurationTime(s string) time.Time {
	now := time.Now()
	duration, _ := time.ParseDuration(s)
	return now.Add(duration)
}

// 时间戳转Time
func UnixToTime(unixStr string) time.Time {
	unix, _ := strconv.ParseInt(unixStr, 10, 64)
	return time.Unix(unix, 0)
}
