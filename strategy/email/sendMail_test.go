package email

import (
	"testing"
)

type a string

func Test_SendMail(t  *testing.T){
	mainInfo := &MailInfo{}
	mainInfo.InitHostInfo()
	mainInfo.AddContent("111sds")
	mainInfo.AddContent("dsfee")
	err := mainInfo.SendMail("799815792@qq.com;panxiang_cn@qq.com")
	if err != nil {
		t.Error("%+v", err)
	}
	ccc := &MailInfo{}
	ccc.InitHostInfo()
	ccc.AddContent("111sds")
	ccc.AddContent("dsfee")
	err = ccc.SendMail("799815792@qq.com;panxiang_cn@qq.com")
	if err != nil {
		t.Error("%+v", err)
	}
}