package utils

import (
	"math"
)

const EARTH_RADIUS = 6378.137
const X_Pi = math.Pi * 3000.0 / 180.0
const BJ54_a = 6378245.0
const BJ54_ee = 0.00669342162296594323

func rad(d float64) float64 {
	return d * math.Pi / 180.0
}

// 根据经纬度计算距离
func GetDistance(lng1, lat1, lng2, lat2 float64) float64 {
	radLat1 := rad(lat1)
	radLat2 := rad(lat2)
	a := radLat1 - radLat2
	b := rad(lng1) - rad(lng2)
	s := 2 * math.Asin((math.Sqrt(math.Pow(math.Sin(a/2), 2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2))))
	s = s * EARTH_RADIUS
	s = math.Ceil(s*10000) / 10000 // 需要自己写round
	return s
}

// 计算运动方向(返回角度)
func GetDirection(lng1, lat1, lng2, lat2 float64) float64 {
	// 不考虑球体，近似计算
	angle := math.Atan2((lat2-lat1), (lng2-lng1)) / math.Pi * 180
	if angle < 0 {
		return angle + 360
	}
	return angle
}

// 火星坐标转百度坐标
func Bd_encrypt(gcj_longi, gcj_lati float64) (bd_longi, bd_lati float64) {
	x, y := gcj_longi, gcj_lati
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*X_Pi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*X_Pi)
	bd_longi = z*math.Cos(theta) + 0.0065
	bd_lati = z*math.Sin(theta) + 0.006
	return bd_longi, bd_lati
}

// 百度坐标转火星坐标
func Bd_decrypt(bd_longi, bd_lati float64) (gcj_longi, gcj_lati float64) {
	x, y := (bd_longi - 0.0065), (bd_lati - 0.006)
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*X_Pi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*X_Pi)
	gcj_longi = z * math.Cos(theta)
	gcj_lati = z * math.Sin(theta)
	return gcj_longi, gcj_lati
}

// GPS坐标转火星坐标
func Gps2Gcj02(longi, lati float64) (gcj_longi, gcj_lati float64) {
	dLat := transformLat(longi-105.0, lati-35.0)
	dLon := transformLon(longi-105.0, lati-35.0)
	radLat := lati / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - BJ54_ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((BJ54_a * (1 - BJ54_ee)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (BJ54_a / sqrtMagic * math.Cos(radLat) * math.Pi)
	gcj_longi = longi + dLon
	gcj_lati = lati + dLat
	return gcj_longi, gcj_lati
}

// GPS坐标转火星坐标(经度)
func transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

// GPS坐标转火星坐标(纬度)
func transformLon(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

// gps卫星信号分档
func DetermineGpssvsLevel(gpssvs int16) int16 {
	if gpssvs <= 4 {
		return 1
	}
	if gpssvs > 4 && gpssvs <= 6 {
		return 2
	}
	if gpssvs > 6 && gpssvs <= 8 {
		return 3
	}
	if gpssvs > 8 && gpssvs <= 10 {
		return 4
	}
	if gpssvs > 10 {
		return 5
	}
	return 5
}

// 基站信号强度分档
func DetermineBssignalLevel(bssignal int16) int16 {
	if bssignal <= -100 {
		return 1
	}
	if bssignal > -100 && bssignal <= -90 {
		return 2
	}
	if bssignal > -90 && bssignal <= -80 {
		return 3
	}
	if bssignal > -80 && bssignal <= -60 {
		return 4
	}
	if bssignal > -60 {
		return 5
	}
	return 5
}
