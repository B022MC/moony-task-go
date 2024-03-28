package utils

import (
	"fmt"
	"strings"
	"time"
)

// ReplaceChineseDateSeparator 将中文日期格式中的“年月日”替换为“-”。
func ReplaceChineseDateSeparator(date string) string {
	date = strings.Replace(date, "年", "-", -1)
	date = strings.Replace(date, "月", "-", -1)
	date = strings.Replace(date, "日", "", -1)
	return date
}

// TimeToDate 时间戳转为日期
func TimeToDate(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02")
}

// TimeToDateTime 时间戳转为datetime格式
func TimeToDateTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// DateTimeToTime 日期转为时间戳
func DateTimeToTime(date string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation("2006-01-02 15:04:05", date, LOC)
	return tim.Unix()
}

// DateToTime 日期转为时间戳
func DateToTime(date string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation("2006-01-02", date, LOC)
	return tim.Unix()
}

// NowDate 当前日期
func NowDate() string {
	return time.Now().Format("2006-01-02")
}

// NowDateTime 当前时间
func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TodayTime 当天开始时间戳
func TodayTime() int64 {
	return DateToTime(NowDate())
}

func TimeToWeek(t int64) string {
	week := time.Unix(t, 0)
	switch week.Weekday() {
	case time.Monday:
		return "星期一"
	case time.Tuesday:
		return "星期二"
	case time.Wednesday:
		return "星期三"
	case time.Thursday:
		return "星期四"
	case time.Friday:
		return "星期五"
	case time.Saturday:
		return "星期六"
	case time.Sunday:
		return "星期天"
	}
	return ""
}

// TimeToTips 时间戳转时间提示
func TimeToTips(timestamp int64) string {
	nowTime := time.Now().Unix()
	span := nowTime - timestamp

	result := ""
	if span < 60 {
		result = "刚刚"
	} else if span <= 1800 {
		result = fmt.Sprintf("%d分钟前", span/60)
	} else if span <= 3600 {
		result = fmt.Sprintf("%d分钟前", 30)
	} else if span <= 86400 {
		result = fmt.Sprintf("%d小时前", span/3600)
	} else if span <= 86400*30 {
		result = fmt.Sprintf("%d天前", span/(86400))
	} else if span <= 86400*30*12 {
		result = fmt.Sprintf("%d月前", span/(86400*30))
	} else {
		result = fmt.Sprintf("%d年前", span/(86400*30*12))
	}

	return result
}
