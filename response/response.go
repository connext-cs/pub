package response

import (
	"encoding/json"
)

// 返回状态
const (
	CodeSuccess = 0
	CodeFail    = 1
)

// Response 返回的内容
type Response struct {
	Code       int         `json:"code"` // 0 成功， 1失败
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
	Pagination *Pagination `json:"pagination"`
}

//NewResponse 普通返回
func NewResponse(code int, errorstr string, data ...interface{}) *Response {
	res := new(Response)
	if len(data) > 0 {
		res.Data = data[0]
	}
	res.Code = code
	res.Error = errorstr
	return res
}

// Pack 转换
func (res *Response) Pack() (data []byte, err error) {
	data, err = json.Marshal(res)
	return
}

// Pagination 分页结构
type Pagination struct {
	PageSize    int `json:"pagesize"`     //每页显示的数据条数, 50
	Page        int `json:"page"`         //总页数 = 总行数/每页显示的数据条数
	CurrentPage int `json:"current_page"` //当前页号
	Total       int `json:"total"`        //总行数
}

// NewResponsePageination 分页返回
func NewResponsePageination(code int, errorstr string, pagination *Pagination, data ...interface{}) *Response {
	res := new(Response)
	if len(data) > 0 {
		res.Data = data[0]
	}
	res.Pagination = pagination
	res.Code = code
	res.Error = errorstr
	return res
}
