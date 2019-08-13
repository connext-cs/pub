package handler

import (
	"net/http"
	"github.com/connext-cs/pub/response"
)

func ReponseSuccessData(responseWriter http.ResponseWriter, data interface{}) {
	var outdata []byte
	outdata, _ = response.NewResponse(0, "", data).Pack()
	//增加跨域问题
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(outdata)
}

func ReponsePaginationSuccessData(responseWriter http.ResponseWriter, pagination *response.Pagination, data interface{}) {
	var outdata []byte
	//outdata, _ = common.NewResponse(0, "", data).Pack()
	outdata, _ = response.NewResponsePageination(0, "", pagination, data).Pack()
	//增加跨域问题
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(outdata)
}

func ReponseFailData(responseWriter http.ResponseWriter, errtext string) {
	var outdata []byte
	outdata, _ = response.NewResponse(1, errtext, "").Pack()
	//增加跨域问题
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(outdata)
}
