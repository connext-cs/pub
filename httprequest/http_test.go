package httprequest

import (
	"strings"
	"testing"
)

func Test_Post(t *testing.T) {
	data, err := Post("", "", "")
	if err == nil {
		t.Error("Post err")
	}
	tockenstr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjIzMTM3NzgsImlhdCI6MTU2MjE0MDk3OCwianRpIjoiMTU2MjE0MDk3OC0xMDAwMDEyNy4wLjAuMSJ9.z70RaedIcYUWk-vNx2jaseLXhnj4fVQeT6dvhxFJ6jo"
	url := "http://10.128.0.150:8002/backend/HostService/List"
	data, err = Post(tockenstr, url, `{"host_tags":[],"pagination":{"pageno":1,"pagerows":10}}`)
	if err != nil {
		t.Error("Post err:", err.Error())
	}

	if strings.TrimSpace(data) == "" {
		t.Error("Post err:", err.Error())
	}
}
