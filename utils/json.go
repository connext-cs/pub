package utils

import "encoding/json"

// 对象转json字符串
func ObjectToJsonStr(v interface{}) string {
	str, err:=json.Marshal(v)
	if err != nil {
		str = []byte("change err")
	}
	return string(str)
}
