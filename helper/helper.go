package helper

import (
	"time"
)

// FormatTime 格式化时间戳
func FormatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
