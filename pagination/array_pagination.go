package pagination

import "github.com/connext-cs/pub/response"

/*
入参:
  totals     数据总数
  pagerows   每页数据行数
返回参数:
  pages 开始编号
*/
func CalcPageNum(totals, pagerows int) (pages int) {
	if totals == 0 {
		return 0
	}
	if pagerows == 0 {
		return totals
	}
	pages = totals / pagerows
	if totals%pagerows > 0 {
		pages = pages + 1
	}
	return pages
}

/*
入参:
  currentpage 当前页数
  totals     数据总数
  pagerows   每页数据行数
返回参数:
  beginIndex 开始编号
  endIndex   结束编号
  pagination 分页数据
*/
func FormatPage(currentpage, totals, pagerows int) (beginIndex, endIndex int, pagination *response.Pagination) {
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)

	pageNo := currentpage
	if (beginIndex == 0) && (endIndex == 0) {
		pageNo = 1
	}
	pageNum := CalcPageNum(totals, pagerows)
	pagination = &response.Pagination{
		PageSize:    pagerows,
		Page:        pageNum, //总页数 = 总行数/每页显示的数据条数
		CurrentPage: pageNo,
		Total:       totals,
	}

	return beginIndex, endIndex, pagination
}

func calcPage(currentpage *int, totals, pagerows int) (beginIndex, endIndex int) {
	c := pagerows
	s := totals
	if s == 0 {
		return 0, 0
	}

	cp := s / c
	if cp <= 0 {
		cp = 1
	}

	if s > c {
		if 1 <= (s % c) {
			cp++
		}
	}

	if *currentpage > cp {
		*currentpage = cp
	}

	if *currentpage <= 0 {
		return 0, c
	} else {
		m := (*currentpage) * c
		j := m - c
		if m <= s {
			return j, m
		} else {
			return j, 0
		}
	}
}
