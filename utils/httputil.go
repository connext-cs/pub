package utils

/**
*网络辅助
 */
import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	QWS_WX_ROOT_URL = "http://wxmsg.360qws.cn"
)

/**Http post 提交请求
*@author Andy.wang
*@param url 请求地址
*@param url  参数
 */
func HttpPost(url, s string) (string, error) {
	var result string
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(s))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return string(body), nil
}

/**下载文件
*@author andy
*@param url :网络地址
*@param name :文件名称
*@param mdr :存储目录
 */
func Download(url, name, mdr string) (err error) {
	out, err := os.Create(mdr + "/" + name)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		return err
	}
	return
}

func GetClientIP(proxyrealip, remoteip string) string {
	if proxyrealip != "" {
		return proxyrealip
	}
	addr := strings.Split(remoteip, ":")
	if len(addr) > 0 {
		return addr[0]
	}
	return ""
}
