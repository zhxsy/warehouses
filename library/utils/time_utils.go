package utils

import (
	"time"
)

// Local using env TZ, fallback to Asia/Shanghai if TZ not set
var Local *time.Location

func init() {
	//if os.Getenv("TZ") == "" {
	//	Local, _ = time.LoadLocation("Asia/Shanghai")
	//} else {
	//	Local = time.Local
	//}
	Local = time.Local
}

// yyyy-MM-dd HH:mm:ss
const WEB = "2006-01-02 15:04:05"
const Continuity = "20060102150405"

// yyyy-MM-dd
const DATE = "2006-01-02"

// hh:mm:ss
const TIME = "15:04:05"

// yyyyMMdd
const DATE_SHORT = "20060102"

// MM/dd/yyyy
const DATE_EN = "01/02/2006"

// MM/dd/yyyy HH:mm:ss
const DATE_EN_TIME = "01/02/2006 15:04:05"

// ParseTime using Local (env TZ or Asia/Shanghai)
func ParseTime(layout, str string) (time.Time, error) {
	return time.ParseInLocation(layout, str, Local)
}

// FormatTimeLocal using Local (env TZ or Asia/Shanghai)
func FormatTimeLocal(layout string, t time.Time) string {
	if layout == "" {
		layout = WEB
	}
	return t.In(Local).Format(layout)
}

// FormatTimestampLocal using Local (env TZ or Asia/Shanghai)
func FormatTimestampLocal(layout string, tsSec int64) string {
	return FormatTimeLocal(layout, time.Unix(tsSec, 0))
}
