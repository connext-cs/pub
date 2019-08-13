package common

import (
	"fmt"
	"testing"
)

func Test_GetURLFileName(t *testing.T) {
	url := "http://10.246.9.125:8002/backend/file/7/137/tst.20190507.bak.zip"
	filePath, filename, err := GetURLFileName(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(filePath)
	fmt.Println(filename)
}
