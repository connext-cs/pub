// Copyright(c)，Shanghai Connext Information Technology Co., Ltd.，All Rights Resevered.

/*
	Time: 2019/07/02
	Author: trump.liu
	File: sql_pagination.go
	Describe: pagination for sql use
*/

package pagination

import "fmt"
import "github.com/connext-cs/pub/response"

const Limit string = " LIMIT %d,%d "

// get sql query limit part
//
// Input
//      pageNo: which page, should be at least 1
//      PageRows: every page records amount
//      totalCount: all records amount
// OutPut
//      sql query limit part string
//
// using the return value as:
// select * from ... where ... limit 0, 10
// if pageNo provided out of page, will return the last page
func GetSqlLimitOption(pageNo, pageRows, totalCnt int) string {
	limitOption := ""
	pageSize := getPageCnt(pageRows, totalCnt)
	if pageNo >= pageSize {
		if pageSize-1 >= 0 {
			limitOption = fmt.Sprintf(Limit, (pageSize-1)*pageRows, pageRows)
		} else {
			return fmt.Sprintf(Limit, 0, pageRows)
		}
	} else {
		limitOption = fmt.Sprintf(Limit, (pageNo-1)*pageRows, pageRows)
	}

	return limitOption
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
func GenerageSqlPagination(limit, totalCount int) *response.Pagination {
	var pagination response.Pagination
	pagination.PageSize = limit
	pagination.Total = totalCount
	if pagination.PageSize > 0 {
		pageCnt := getPageCnt(pagination.PageSize, pagination.Total)
		pagination.Page = pageCnt
	}
	return &pagination
}
