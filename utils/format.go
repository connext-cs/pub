package utils

import (
	"math"
	"time"
)

// 获取时间类型中午字符串
func TimeString(ts time.Time) string {
	return ts.Format("2006年01月02日15时04分05秒")
}

func TimeStringDate(ts time.Time) string {
	s := ts.Format("2006年01月02日15时04分05秒")
	if len(s) < 17 {
		return ""
	}
	return s[:17]
}

// 4字节对齐补零
func Align2fourbytes(b []byte) []byte {
	miss := len(b) % 4
	if miss > 0 {
		for i := 0; i < (4 - miss); i++ {
			b = append(b, byte(0))
		}
	}
	return b
}

// float64准确精度
func Float64toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
