package common

import (
	"fmt"
	"time"
)

var (
	DATETIMESTR = "2006-01-02 15:04:05"
	DATESTR     = "2006-01-02"
	CDATESTR    = "2006年01月02日"
	MONTHSTR    = "2006-01"
	CMONTHSTR   = "2006年01月"
)

var CWEEKDAY = map[time.Weekday]string{
	time.Monday:    "周一",
	time.Tuesday:   "周二",
	time.Wednesday: "周三",
	time.Thursday:  "周四",
	time.Friday:    "周五",
	time.Saturday:  "周六",
	time.Sunday:    "周日",
}

func GetChinaWeekday(w time.Weekday) (v string) {
	v, _ = CWEEKDAY[w]
	return
}

func GetToday(format string) string {
	today := time.Now()
	switch format {
	case "year":
		return fmt.Sprintf("%d", today.Year())
	case "day":
		return fmt.Sprintf("%d", today.Day())
	default:
		return Time2Str(today, format)
	}
}

/**
字符串转字符串
timeStr: 日期字符串
format: 时间格式
newformat: 新时间格式
*/
func Str2Str(timeStr, format, newformat string) string {
	theTime := Str2Time(timeStr, format)
	return Time2Str(theTime, newformat)
}

/**
字符串转时间对象
timeStr: 日期字符串
format: 时间格式
*/
func Str2Time(timeStr, format string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(format, timeStr, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

/**
字符串转时间戳
timeStr: 日期字符串
format: 时间格式
*/
func Str2Stamp(timeStr, format string) int64 {
	timeStruct := Str2Time(timeStr, format)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/**
时间对象转字符串
timeObj: 时间对象
format: 时间格式
*/
func Time2Str(timeObj time.Time, format string) string {
	temp := time.Date(timeObj.Year(), timeObj.Month(), timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second(), timeObj.Nanosecond(), time.Local)
	str := temp.Format(format)
	return str
}

/**
时间对象转时间戳
timeObj: 时间对象
*/
func Time2Stamp(timeObj time.Time) int64 {
	millisecond := timeObj.UnixNano() / 1e6
	return millisecond
}

/**
时间戳转字符串
stamp: 时间戳
format:时间格式
*/
func Stamp2Str(stamp int64, format string) string {
	str := time.Unix(stamp/1000, 0).Format(format)
	return str
}

/**
时间戳转时间对象
stamp: 时间戳
format:时间格式
*/
func Stamp2Time(stamp int64, format string) time.Time {
	stampStr := Stamp2Str(stamp, format)
	timer := Str2Time(stampStr, format)
	return timer
}
