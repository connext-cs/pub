package email

import (
	"k8sproject/config"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"k8sproject/plugin/log"	
)

type MailInfo struct {
	loginAuth
	unencryptedAuth
	host    string
	content string
	title   string
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	_, resp, th := a.Auth.Start(&s)
	return "LOGIN", resp, th
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	command := string(fromServer)
	command = strings.TrimSpace(command)
	command = strings.TrimSuffix(command, ":")
	command = strings.ToLower(command)
	if more {
		if command == "username" {
			return []byte(fmt.Sprintf("%s", a.username)), nil
		} else if command == "password" {
			return []byte(fmt.Sprintf("%s", a.password)), nil
		} else {
			// We've already sent everything.
			return nil, fmt.Errorf("unexpected server challenge: %s", command)
		}
	}
	return nil, nil
}

func (mailInfo *MailInfo) InitHostInfo() {
	// mailInfo.SetServer("smtp.office365.com", "587")
	// mailInfo.SetUser("xiang.pan@connext.com.cn", "Px@1234cn")
	// mailInfo.title = "defaultTitle"
	mailInfo.SetServer(config.GetConfig().SonarEMail.Host, config.GetConfig().SonarEMail.GetPort())
	mailInfo.SetUser(config.GetConfig().SonarEMail.User, config.GetConfig().SonarEMail.Password)
	mailInfo.title = config.GetConfig().SonarEMail.DefaultTitle
}

func (mailInfo *MailInfo) SetTitle(title string) {
	mailInfo.title = title
}

func (mailInfo *MailInfo) ClearEMail() {
	mailInfo.content = ""
}

func (mailInfo *MailInfo) AddContent(contentLine string) {
	mailInfo.content = fmt.Sprintf("%s%s\r\n", mailInfo.content, contentLine)
}

func (mailInfo *MailInfo) SendMail(receiveAddr string) error {
	log.Info("receiveAddr:%+v,mailInfo:%+v", receiveAddr, mailInfo)
	//mailInfo.InitHostInfo()
	err := mailInfo.send(receiveAddr)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return err
	}
	return nil
}

func (mailInfo *MailInfo) SetServer(hostname, port string) {
	//hostname="smtp.qq.com"			//qq邮箱服务器地址
	//port="587"						//qq邮箱SMTP服务器端口
	//hostname="smtp.office365.com"		//outlook服务器地址
	//port="587"								//outlookSMTP服务器端口
	mailInfo.host = net.JoinHostPort(hostname, port)
}

func (mailInfo *MailInfo) SetUser(username, password string) {
	mailInfo.Auth = LoginAuth(username, password)
	mailInfo.username = username
}

func (mailInfo *MailInfo) send(receiveAddr string) error {
	to := strings.Split(receiveAddr, ";")
	msg := []byte("To: " + receiveAddr + "\r\n" + "Subject:" + mailInfo.title + "\r\n" + "\r\n" + mailInfo.content + "\r\n")
	//to := []string{receiveAddr}
	err := smtp.SendMail(mailInfo.host, mailInfo.Auth, mailInfo.username, to, msg)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}
	return nil
}

//SonarEMail
var sonarQuberEMail *MailInfo

//GetSonarQuberEMail ...
func GetSonarQuberEMail() *MailInfo {
	if sonarQuberEMail == nil {
		sonarQuberEMail = &MailInfo{}
		sonarQuberEMail.InitHostInfo()
	}
	return sonarQuberEMail
}
