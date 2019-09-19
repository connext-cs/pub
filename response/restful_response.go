package response

import (
	"github.com/emicklei/go-restful"
	"encoding/json"
)

// RspFailRestData 失败返回 数据
func RspFailRestData(rsp *restful.Response, errMsg string) {
	response := NewResponse(CodeFail, errMsg)
	rsp.WriteAsJson(response)
}

//RspSucRestData 成功返回
func RspSucRestData(rsp *restful.Response, errMsg string, data interface{}) {
	response := NewResponse(CodeSuccess, errMsg, data)
	rsp.WriteAsJson(response)
}

// RspInfoRestData 成功信息返回
func RspInfoRestData(rsp *restful.Response, InfoMsg string, data interface{}) {
	response := NewResponse(CodeSuccess, InfoMsg, data)
	rsp.WriteAsJson(response)
}

// ReponsePaginationSuccessData 分页返回
func ReponsePaginationSuccessData(rsp *restful.Response, pagination *Pagination, data interface{}) {
	outdata := NewResponsePageination(0, "", pagination, data)
	rsp.WriteAsJson(outdata)
}

//responseSuccessRestData 成功返回 数据
func responseSuccessRestData(rsp *restful.Response, data interface{}) {
	//response := response.NewResponse(CODE_SUCCESS, "", data)
	resp := new(Response)
	json.Unmarshal([]byte(data.(string)), &resp)

	//rsp.WriteAsJson(string(jData))

	rsp.WriteEntity(resp)
}

// ResponseSuccessRestData ...
func ResponseSuccessRestData(rsp *restful.Response, data interface{}) {
	response := NewResponse(CodeSuccess, "", data)
	rsp.WriteAsJson(response)
	//rsp.WriteEntity(response)
}

//responseSuccessPaginationRestData 成功分页返回 数据
func responseSuccessPaginationRestData(rsp *restful.Response, pagination *Pagination, data interface{}) {
	response := NewResponse(CodeSuccess, "", data)
	response.Pagination = pagination
	rsp.WriteAsJson(response)
}

// responseFailRestDataDo 失败返回 数据
func responseFailRestDataDo(rsp *restful.Response, errMsg string) {
	response := NewResponse(CodeFail, errMsg)
	rsp.WriteAsJson(response)
}

// ResponseFailRestData ...
func ResponseFailRestData(rsp *restful.Response, errMsg string) {
	responseFailRestDataDo(rsp, errMsg)
}

// responseFailRestDataInfo 失败返回 带数据
func responseFailRestDataInfo(rsp *restful.Response, errMsg string, data interface{}) {
	response := NewResponse(CodeFail, errMsg, data)
	rsp.WriteAsJson(response)
}

func ResponseSuccessRestDataRPC(rsp *restful.Response, data interface{}) {
	responseSuccessHttpData((*rsp).ResponseWriter, data)
}
