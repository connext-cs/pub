package common

import (
	"errors"
	"path"
	"github.com/connext-cs/pub/log"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

//ValidIPSegment ... 是否是正确的IP地址段格式 172.10.16.0/32
func ValidIPSegment(value string) error {
	const MatchKey = `^(?:(?:1[0-9][0-9]\.)|(?:2[0-4][0-9]\.)|(?:25[0-5]\.)|(?:[1-9][0-9]\.)|(?:[0-9]\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))(\/(?:[0-9]|[12][0-9]|3[012]))$`
	match, err := regexp.MatchString(MatchKey, value)
	if !match {
		err = errors.New(value + " is invalid ValidIPSegment.")
	}
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

//是否是正确的IP地址格式
func ValidIP(ip string) error {
	const MatchKey = `^(?:(?:1[0-9][0-9]\.)|(?:2[0-4][0-9]\.)|(?:25[0-5]\.)|(?:[1-9][0-9]\.)|(?:[0-9]\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))$`
	match, err := regexp.MatchString(MatchKey, ip)
	if !match {
		err = errors.New(ip + " is invalid IP.")
	}
	if err != nil {
		beego.Info(err.Error())
	}
	return err
}

//只能输入小写字符和数字
func ValidPcName(value string) error {
	var (
		match bool
		err   error
	)
	const MatchKey = "^[a-z0-9]+$"
	match, err = regexp.MatchString(MatchKey, value)
	if !match {
		err = errors.New(value + " 格式错误, 只能小写英文字母和数字组成,不能带空格.")
	}
	if err != nil {
		beego.Info(err.Error())
	}
	return err
}

/*判断 int 是否 小于等于 0*/
func ValidInt(value int, label string) (err error) {
	if value <= 0 {
		beego.Error(label + " is less or equal 0")
		err = errors.New(label + " is less or equal 0")
		return err
	}
	return nil
}

/*判断 int64 是否 小于等于 0*/
func ValidInt64(value int64, label string) (err error) {
	if value <= 0 {
		beego.Error(label + " is less or equal 0")
		err = errors.New(label + " is less or equal 0")
		return err
	}
	return nil
}

/*判断字符串是否为空*/
func ValidString(value string, label string) (err error) {
	if len(strings.TrimSpace(value)) == 0 {
		beego.Error(label + " is '' ")
		err = errors.New(label + " is '' ")
		return err
	}
	return nil
}

/*获取路径最后文件名 例如   /var/log/2.log      返回 2.log*/
func LastStringFileName(p string) (string, error) {
	err := ValidString(p, "log 输入字符串不能为空")
	if err != nil {
		return "", err
	}

	var data []string
	data = strings.Split(p, "/")
	if len(data) == 0 {
		return "", err
	}
	dataindex := len(data) - 1
	return data[dataindex], nil
}

//input fullfilename get filename
func GetFileName(fullFilename string) string {
	if strings.TrimSpace(fullFilename) == "" {
		return ""
	}
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	return filenameOnly
}

func FormatDBTime(timestr string) string {
	timestr = strings.Replace(timestr, "T", " ", 100)
	timestr = strings.Replace(timestr, "Z", " ", 100)
	return timestr
}
