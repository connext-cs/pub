package utils

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

var zlibWriterPool = sync.Pool{
	New: func() interface{} {
		return zlib.NewWriter(nil)
	},
}

var zlibReaderPool = sync.Pool{}

type LocatorVersion struct {
	Version string
	Sn      string
	Imsi    string
	Imei    string
}

// 转换GPS卫星时间
func gps2utc(gpsweek, gpssec int32) time.Time {
	basetime, _ := time.Parse("2006-Jan-02", "1980-Jan-06")
	utctime := basetime.AddDate(0, 0, int(gpsweek*7))
	utctime = utctime.Add(time.Duration(gpssec * 1e9))
	return utctime
}

// qws zip
func QwsZipWrite(origin []byte) (zip []byte) {
	var b bytes.Buffer
	zw := zlib.NewWriter(&b)
	origin = append(origin, 0) // 为兼容c对字符串的解析，末尾加0
	zw.Write(origin)
	zw.Close()
	olen := ReverseBytes(Uint32ToBytes(uint32(len(origin))))
	nlen := ReverseBytes(Uint32ToBytes(uint32(b.Len())))
	zip = append([]byte{0xFD, 0xFD, 0xFE, 0xFE}, nlen...)
	zip = append(zip, olen...)
	zip = append(zip, b.Bytes()...)
	return
}

func QwsZIPEncode(in []byte) ([]byte, error) {
	var data bytes.Buffer
	data.Write(in)
	data.WriteByte(0) // 为兼容c对字符串的解析，末尾加0

	var z = bytes.NewBuffer(make([]byte, 0, len(in)))
	w := zlibWriterPool.Get().(*zlib.Writer)
	w.Reset(z)
	w.Write(data.Bytes())
	w.Close()
	zlibWriterPool.Put(w)

	olen := data.Len()
	nlen := z.Len()

	out := bytes.NewBuffer(make([]byte, 0, 12+nlen))
	out.Write([]byte{0xFD, 0xFD, 0xFE, 0xFE})
	//out := bytes.NewBuffer([]byte{0xFD, 0xFD, 0xFE, 0xFE})

	out.WriteByte(byte(nlen))
	out.WriteByte(byte(nlen >> 8))
	out.WriteByte(byte(nlen >> 16))
	out.WriteByte(byte(nlen >> 24))

	out.WriteByte(byte(olen))
	out.WriteByte(byte(olen >> 8))
	out.WriteByte(byte(olen >> 16))
	out.WriteByte(byte(olen >> 24))

	out.Write(z.Bytes())

	return out.Bytes(), nil
}

func QwsZIPDecode(data []byte) ([]byte, error) {
	b := bytes.NewReader(data[12:])

	r := zlibReaderPool.Get()
	var r2 io.ReadCloser
	var err error
	if r == nil {
		r2, err = zlib.NewReader(b)
		if err != nil {
			return nil, err
		}
	} else {
		r2 = r.(io.ReadCloser)
		resetter := r.(zlib.Resetter)
		resetter.Reset(b, nil)
	}
	defer func() {
		r2.Close()
		zlibReaderPool.Put(r2)
	}()

	var rlt []byte
	rlt, err = ioutil.ReadAll(r2)
	if err != nil {
		return nil, err
	}
	rlt = rlt[:len(rlt)-1] // 去除c字符串末尾0
	return rlt, nil
}

// 天气编码
func EncodeWeather(weather string) string {
	switch weather {
	case "晴":
		return "00"
	case "多云":
		return "01"
	case "阴":
		return "02"
	case "阵雨":
		return "03"
	case "雷阵雨":
		return "04"
	case "雷阵雨并伴有冰雹":
		return "05"
	case "雨夹雪":
		return "06"
	case "小雨":
		return "07"
	case "中雨":
		return "08"
	case "大雨":
		return "09"
	case "暴雨":
		return "10"
	case "大暴雨":
		return "11"
	case "特大暴雨":
		return "12"
	case "阵雪":
		return "13"
	case "小雪":
		return "14"
	case "中雪":
		return "15"
	case "大雪":
		return "16"
	case "暴雪":
		return "17"
	case "雾":
		return "18"
	case "冻雨":
		return "19"
	case "沙尘暴":
		return "20"
	case "小雨-中雨":
		return "21"
	case "中雨-大雨":
		return "22"
	case "大雨-暴雨":
		return "23"
	case "暴雨-大暴雨":
		return "24"
	case "大暴雨-特大暴雨":
		return "25"
	case "小雪-中雪":
		return "26"
	case "中雪-大雪":
		return "27"
	case "大雪-暴雪":
		return "28"
	case "浮尘":
		return "29"
	case "扬沙":
		return "30"
	case "强沙尘暴":
		return "31"
	case "飑":
		return "32"
	case "龙卷风":
		return "33"
	case "弱高吹雪":
		return "34"
	case "轻雾":
		return "35"
	case "霾":
		return "53"
	default:
		return "NA"
	}
}

// 天气转码（详细转基本）
func WeatherAll2Base(all string) (base string) {
	switch all {
	case "00", "01":
		return "00" // 晴
	case "02", "18", "20", "29", "30", "31", "33", "35", "53":
		return "01" // 阴
	case "03", "04", "05", "07", "08", "09", "10", "11", "12", "19", "21", "22", "23", "24", "25", "32":
		return "07" // 小雨
	case "06", "13", "14", "15", "16", "17", "26", "27", "28", "34":
		return "14" // 小雪
	default:
		return "NA"
	}
}

// 解析定位器版本信息
func ParseLocatorVersion(ver string) (version LocatorVersion) {
	json.Unmarshal([]byte(ver), &version)
	version.Version = strings.Trim(strings.Replace(version.Version, "Project Version:", "", -1), " ")
	return
}

func Rune2String(brand []rune) (r string) {
	for _, v := range brand {
		if v == 0 {
			break
		}
		r += string(v)
	}
	return
}
