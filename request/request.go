package request

//Pagination 请求时的分页结构
type RequestPagination struct {
	PageNo   uint16 `json:"pageno"`   //当前的页数
	PageRows uint16 `json:"pagerows"` //每页的行数
}

// Request http 請求的包头
type Request struct {
	Params     interface{}        `json:"params"`
	Pagination *RequestPagination `json:"pagination"`
}
