package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/connext-cs/pub/common"
	"net"
	"os"
	"regexp"
	"strings"
)

var localIpAddr = ""

func GenerateTlsConfig(servCertPEM, servKeyPEM string) (config *tls.Config, err error) {
	servTLSCert, err := tls.X509KeyPair([]byte(servCertPEM), []byte(servKeyPEM))
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
	}, nil
}

func localIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func GetLocalIp() string {
	if localIpAddr == "" {
		var err error
		if localIpAddr, err = localIp(); err != nil {
			fmt.Println(err)
			os.Exit(common.OsExitSignal)
		}
	}
	return localIpAddr
}
//数据为空判断
func IsEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

//正则匹配
func IsMatch(content string, regex string) bool {
	return regexp.MustCompile(regex).MatchString(content)
}
func IsNotMatch(content string, regex string) bool {
	return !IsMatch(content, regex)
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return false
}

func IsNotNil(v interface{}) bool {
	return !IsNil(v)
}

func GetMapValueByKey(mapKey string, mapData map[string]interface{}) interface{}  {
	if _,ok:=mapData[mapKey];ok{
		return mapData[mapKey]
	}
	return nil
}

