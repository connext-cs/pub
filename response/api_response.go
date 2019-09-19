package response

import (
	api "github.com/micro/go-api/proto")

func reponseSuccessData(rsp *api.Response, data interface{}) {
	var outdata []byte
	outdata, _ = NewResponse(0, "", data).Pack()
	//增加跨域问题
	// responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// responseWriter.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	// responseWriter.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-types,Authorization")
	rsp.Body = string(outdata)
}

func reponseSuccessList(rsp *api.Response, pagination *Pagination, data interface{}) {
	var outdata []byte
	//outdata, _ = response.NewResponse(0, "", data).Pack()
	response := NewResponse(0, "", data)
	response.Pagination = pagination
	outdata, _ = response.Pack()
	//增加跨域问题
	// responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// responseWriter.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	// responseWriter.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-types,Authorization")
	rsp.Body = string(outdata)
}

func reponsePaginationSuccessData(rsp *api.Response, pagination *Pagination, data interface{}) {
	var outdata []byte
	outdata, _ = NewResponsePageination(0, "", pagination, data).Pack()
	rsp.Body = string(outdata)
}

func reponseFailData(rsp *api.Response, errtext string) {
	var outdata []byte
	outdata, _ = NewResponse(1, errtext, "").Pack()
	rsp.Body = string(outdata)
}

// error is a interface
func reponseFailDataErrorInfo(rsp *api.Response, errtext string, data interface{}) {
	var outdata []byte
	outdata, _ = NewResponse(1, errtext, data).Pack()
	rsp.Body = string(outdata)
}
