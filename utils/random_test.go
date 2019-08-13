package utils

import (
	"fmt"
	"testing"
)

func Test_MakeRandomStrSize(t *testing.T) {
	data := MakeRandomStrSize(10)
	fmt.Println("data:", data)
}

func Test_random(t *testing.T) {
	data := random(10)
	fmt.Println("data:", data)
}

func Test_NewUUID(t *testing.T) {
	data, err := NewUUID()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("data:", data)
}

func Test_MakeRandomStr(t *testing.T) {
	data := MakeRandomStr()
	fmt.Println("data:", data)
}
