package common

import (
	"testing"
)

func Test_ValidIP(t *testing.T) {
	err := ValidIP("192.168.110.57")
	if err != nil {
		t.Error("ValidIP err:", err.Error())
	}
}

func Test_GetFileName(t *testing.T) {
	fullFilename := "/Users/itfanr/Documents/test.txt"
	filename := GetFileName(fullFilename)
	if filename != "test" {
		t.Error("GetFileName error")
	}

	fullFilename = "test.0619.zip"
	filename = GetFileName(fullFilename)
	if filename != "test" {
		t.Error("GetFileName error")
	}
}
