package utils

import (
	"time"
)

// 是否同一天
func SameDay(time1, time2 time.Time) bool {
	_, offset1 := time1.Zone()
	_, offset2 := time2.Zone()
	if offset1 != offset2 { //  必须同时区
		return false
	}
	if time1.Format("20060102") != time2.Format("20060102") {
		return false
	}
	return true
}

// 截取日期到天
func Trunc2Date(ts time.Time) time.Time {
	ts = ts.Local() // 转换为服务器时区
	str := ts.Format("20060102 -0700")
	t, _ := time.Parse("20060102 -0700", str)
	return t
}

// 计算日期间隔天数
func DateDiff(date1, date2 time.Time) int32 {
	date1_d, date2_d := Trunc2Date(date1), Trunc2Date(date2)
	days := date1_d.Sub(date2_d) / (time.Hour * 24)
	return int32(days)
}

// 当天开始时间
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// 当天结束时间
func Bof(t time.Time) time.Time {
	year, month, day := t.Date()
	m, _ := time.ParseDuration("-1ns")
	return time.Date(year, month, day+1, 0, 0, 0, 0, t.Location()).Add(m)
}

// 当月开始时间
func Bom(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 0, 0, 0, 0, 0, t.Location())
}

// 转换日期为本地时区
func Date2Local(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// 返回时间差，开始时间到当前时间差
func TimeCost(start time.Time) time.Duration{
	return time.Since(start)
}
