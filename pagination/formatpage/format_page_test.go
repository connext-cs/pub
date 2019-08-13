package formatpage

import (
	"fmt"
	"testing"
)

func Test_FormatArray(t *testing.T) {

	var TestData []int
	TestData = make([]int, 100)
	for i := range TestData {
		TestData[i] = i
	}

	currentpage := 6
	totals := 100
	pagerows := 12
	beginIndex, endIndex, page := FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 60) && (endIndex == 72) && (page == 6)) {
		t.Error("FormatPage error")
	}
	var resultList []int
	if (beginIndex == 0) && (endIndex == 0) {
		resultList = TestData
	} else if beginIndex == 0 {
		resultList = TestData[:endIndex]
	} else if endIndex == 0 {
		resultList = TestData[beginIndex:]
	} else {
		resultList = TestData[beginIndex:endIndex]
	}
	fmt.Println(resultList)
	if len(resultList) != 12 {
		t.Error("FormatPage error")
	}

	if !((resultList[0] == 60) && (resultList[11] == 71) && (page == 6)) {
		t.Error("1111FormatPage error")
	}
}

func Test_FormatPage(t *testing.T) {
	currentpage := 1
	totals := 10
	pagerows := 6
	beginIndex, endIndex, page := FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 0) && (endIndex == 6)) && (page == 1) {
		t.Error("FormatPage error")
	}

	currentpage = 2
	totals = 10
	pagerows = 6
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 6) && (endIndex == 0)) && (page == 2) {
		t.Error("FormatPage error")
	}

	currentpage = 3
	totals = 100
	pagerows = 11
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 22) && (endIndex == 33)) && (page == 3) {
		t.Error("FormatPage error")
	}

	currentpage = 10
	totals = 100
	pagerows = 11
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 99) && (endIndex == 0)) && (page == 10) {
		t.Error("FormatPage error")
	}

	currentpage = 9
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 96) && (endIndex == 0)) && (page == 9) {
		t.Error("FormatPage error")
	}

	currentpage = 11
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 96) && (endIndex == 0)) && (page == 9) {
		t.Error("FormatPage error")
	}

	currentpage = 1
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 0) && (endIndex == 12)) && (page == 1) {
		t.Error("FormatPage error")
	}

	currentpage = 0
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 0) && (endIndex == 12)) && (page == 1) {
		t.Error("FormatPage error")
	}

	currentpage = 2
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 12) && (endIndex == 24)) && (page == 2) {
		t.Error("FormatPage error")
	}

	currentpage = 5
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 48) && (endIndex == 60)) && (page == 5) {
		t.Error("FormatPage error")
	}

	currentpage = 6
	totals = 100
	pagerows = 12
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 60) && (endIndex == 72)) && (page == 6) {
		t.Error("FormatPage error")
	}

	currentpage = 1
	totals = 5
	pagerows = 6
	beginIndex, endIndex, page = FormatPage(currentpage, totals, pagerows)
	if !((beginIndex == 0) && (endIndex == 0)) && (page == 1) {
		t.Error("FormatPage error")
	}
}

func Test_calcPage(t *testing.T) {
	currentpage := 1
	totals := 10
	pagerows := 6
	beginIndex, endIndex := calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 0) && (endIndex == 6)) {
		t.Error("FormatPage error")
	}

	currentpage = 2
	totals = 10
	pagerows = 6
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 6) && (endIndex == 0)) {
		t.Error("FormatPage error")
	}

	currentpage = 3
	totals = 100
	pagerows = 11
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 22) && (endIndex == 33)) {
		t.Error("FormatPage error")
	}

	currentpage = 10
	totals = 100
	pagerows = 11
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 99) && (endIndex == 0)) {
		t.Error("FormatPage error")
	}

	currentpage = 9
	totals = 100
	pagerows = 12
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 96) && (endIndex == 0)) {
		t.Error("FormatPage error")
	}

	currentpage = 11
	totals = 100
	pagerows = 12
	beginIndex, endIndex = calcPage(&currentpage, totals, pagerows)
	if !((beginIndex == 96) && (endIndex == 0)) {
		t.Error("FormatPage error")
	}
}
