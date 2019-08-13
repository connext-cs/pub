package formatpage

func FormatPage(currentpage, totals, pagerows int) (beginIndex, endIndex, pageNo int) {
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	// if endIndex > 0 {
	// 	endIndex--
	// }

	pageNo = currentpage
	if (beginIndex == 0) && (endIndex == 0) {
		pageNo = 1
	}

	if totals <= endIndex {
		return beginIndex, totals, pageNo
	}

	return beginIndex, endIndex, pageNo
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
