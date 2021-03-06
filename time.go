package wen

import (
	"fmt"
	"time"
)

func Timestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
func TimeFormat(t time.Time, layout string) string {
	return t.Local().Format(layout)
}
func TimeFormatDateTime(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05")
}
func TimeFormatDate(t time.Time) string {
	return t.Local().Format("2006-01-02")
}

// cast.ToTime()
func TimeParse(timeString, layout string) time.Time {
	t, _ := time.ParseInLocation(layout, timeString, time.Local)
	return t
}

// 统计耗时
// defer wen.TimeCost()()
func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}
