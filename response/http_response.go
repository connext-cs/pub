package response

import (
	"net/http"
	"encoding/json"
	util "github.com/connext-cs/pub/utils"
	"github.com/connext-cs/pub/logs"
)

func responseSuccessHttpData(rsp http.ResponseWriter, data interface{}) {
	response, _ := NewResponse(CodeSuccess, "", data).Pack()
	rsp.Write(response)
}
func ResponseSuccessHttpFileData(rsp http.ResponseWriter, data interface{}) {
	//response, _ := response.NewResponse(CODE_SUCCESS, "", data).Pack()
	//rsp.Write(response)

	resp := new(Response)
	err := json.Unmarshal([]byte(data.(string)), &resp)
	if util.IsNotNil(err) {
		logs.Error(err)
	}
	reData, err := json.Marshal(resp)
	if util.IsNotNil(err) {
		logs.Error(err)
	}

	rsp.Write(reData)
}

func ResponseFailHttpData(rsp http.ResponseWriter, errMsg string) {
	response, _ := NewResponse(CodeFail, errMsg).Pack()
	rsp.Write(response)
}
func ResponseFailHttpFileData(rsp http.ResponseWriter, errMsg string) {
	//response, _ := response.NewResponse(CODE_FAIL, errMsg).Pack()
	//rsp.Write(response)
	resp := new(Response)
	err := json.Unmarshal([]byte(errMsg), &resp)
	if util.IsNotNil(err) {
		logs.Error(err)
	}
	reData, err := json.Marshal(resp)
	if util.IsNotNil(err) {
		logs.Error(err)
	}

	rsp.Write(reData)
}
