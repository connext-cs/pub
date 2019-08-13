package common

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"github.com/connext-cs/pub/logs"
)

func GetAppPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("GetAppPath err:", err.Error())
		panic(err)
		return ""
	}
	return path
}

func GetWorkPath() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return workPath
}

func GetURLFileName(urlstr string) (filePath, filename string, err error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		logs.Error(err)
		return "", "", err
	}
	filePath, filename = filepath.Split(u.Path)
	return
}
