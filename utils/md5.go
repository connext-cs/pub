package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 返回MD5字符串
func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 单字节异或校验
func ByteXor(data []byte) byte {
	xor := byte(0)
	for _, v := range data {
		xor = xor ^ v
	}
	return xor
}
