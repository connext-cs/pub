// Copyright(c)，Shanghai Connext Information Technology Co., Ltd.，All Rights Resevered.

/*
	Time: 2019/6/20
	Author: trump.liu
	File: orm_pagination.go
	Describe: pagination for orm use
*/

package pagination

import "github.com/connext-cs/pub/response"
import "errors"

// get orm limit and offset
//
// Input
// 		pageNo: which page, should be at least 1
//	    PageRows: every page records amount
//      totalCount: all records amount
// OutPut
//      limit and offset
//
// using the return value as:
// orm.Limit(limit, offset).Find(&example)
// if pageNo provided out of page, will return the offset of the last page
func GetORMLimitOffset(pageNo, pageRows, totalCount int) (int, int, error) {
	limit := 0
	offset := 0

	if pageNo < 1 {
		return 0, 0, errors.New("page rows need more then 1 or equals 1")
	} else {
		limit = pageRows
	}

	pageNo--
	if pageNo > 0 {
		offset = pageNo * pageRows
	} else {
		offset = 0
	}

	pageSize := getPageCnt(limit, totalCount)
	if pageNo >= pageSize {
		if pageSize-1 >= 0 {
			offset = (pageSize - 1) * pageRows
		}
	}

	return limit, offset, nil

}

// generage ORM Pagination
//
// Input
// 		limit: every page records amount
//      totalCount: all records amount
// OutPut
//      pagination.PageSize, every page records amount
//      pagination.Total, all records amount
//      pagination.Page, total page amount
func GenerageORMPagination(limit, totalCount int) *response.Pagination {
	var pagination response.Pagination
	pagination.PageSize = limit
	pagination.Total = totalCount
	if pagination.PageSize > 0 {
		pageCnt := getPageCnt(pagination.PageSize, pagination.Total)
		pagination.Page = pageCnt
	}
	return &pagination
}

// get the page count
// for example total 32 records and every page has 5, the result is  7
func getPageCnt(limit, totalCount int) int {
	mod := totalCount % limit
	pageCnt := totalCount / limit
	if mod > 0 {
		pageCnt++
	}
	return pageCnt
}
