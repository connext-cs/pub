package utils

import (
	crand "crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func MakeRandomStrSize(length int) (str string) {
	randomStr := MakeRandomStr()
	fmt.Println("randomStr:", randomStr)
	size := len([]byte(randomStr))
	if length > size {
		length = size
	}
	for i := 0; i < length; i++ {
		str += string(randomStr[i])
	}
	return
}

func MakeRandomStr() string {
	r := random(6)
	r += strconv.FormatInt(time.Now().Unix(), 10)
	return Md5([]byte(r))
}

func MakeToken() string {
	r := random(8)
	r += strconv.FormatInt(time.Now().Unix(), 10)
	return Md5([]byte(r))
}

func MakeActivationKey() string {
	r := random(22)
	r += strconv.FormatInt(time.Now().Unix(), 10)
	return Md5([]byte(r))
}

// 生产n位数字，主用作验证码
func Randomint(length int) (str string) {

	var arr []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(arr)
	for i := 0; i < length; i++ {
		str += string(arr[r.Intn(size)])
	}
	return
}

func random(length int) (str string) {

	var arr []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(arr)
	for i := 0; i < length; i++ {
		str += string(arr[r.Intn(size)])
	}
	return
}

func CreateSecureRandom(length int) (string, error) {

	b := make([]byte, length)
	_, err := crand.Read(b)
	if err != nil {
		return "", err
	}
	if len(b) == 0 {
		return "", errors.New("len(b) == 0")
	}
	sr := hex.EncodeToString(b)

	return sr, nil
}

//生成
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
