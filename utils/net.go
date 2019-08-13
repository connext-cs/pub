package utils

import (
	"encoding/binary"
	"reflect"
)

// 转换接口为网络字节流（只包含无符号整数和字符串）
func Any2Byte(i interface{}) []byte {
	var buf, b []byte
	sv := reflect.Indirect(reflect.ValueOf(i))
	for kk := 0; kk < sv.NumField(); kk++ {
		switch sv.Field(kk).Type().String() {
		case "uint64": // 8字节无符号整数
			b = Uint64ToBytes(sv.Field(kk).Uint())
			buf = append(buf, b...)
		case "uint32": // 4字节无符号整数
			b = Uint32ToBytes(uint32(sv.Field(kk).Uint()))
			buf = append(buf, b...)
		case "uint16": // 4字节无符号整数
			b = Uint16ToBytes(uint16(sv.Field(kk).Uint()))
			buf = append(buf, b...)
		case "uint8": // 单字节无符号整数
			buf = append(buf, byte(sv.Field(kk).Uint()))
		case "[]uint8": // 字符串
			buf = append(buf, sv.Field(kk).Bytes()...)
		}
	}
	return buf
}

//  无符号整数到byte
func Uint16ToBytes(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}
func Uint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}
func Uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// 网络字节到无符号整数
func BytesToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}
func BytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// 将字节反序
func ReverseBytes(b []byte) []byte {
	r := make([]byte, len(b))
	for k, v := range b {
		r[len(b)-k-1] = v
	}
	return r
}
