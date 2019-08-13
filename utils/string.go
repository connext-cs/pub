package utils

import (
	"runtime"
	"strings"
)

// 检查字符是否在数组里
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// 检查整数是否在数组里
func Uint16InSlice(a uint16, list []uint16) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// 获取当前函数名字
func CallerFuncName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	names := strings.Split(f.Name(), ".")
	if len(names) == 2 {
		return names[1]
	}
	return ""
}

// 判断interface数组里是否含有空值
func ExistsNil(s []interface{}) bool {
	for _, v := range s {
		if v == nil {
			return true
		}
	}
	return false
}
