package pagination

import (
	"testing"
)

func Test_FormatPage(t *testing.T) {
	pagenum := CalcPageNum(0, 5)
	if pagenum != 0 {
		t.Error("CalcPageNum error")
	}

	pagenum = CalcPageNum(5, 0)
	if pagenum != 5 {
		t.Error("CalcPageNum error")
	}

	pagenum = CalcPageNum(9, 3)
	if pagenum != 3 {
		t.Error("CalcPageNum error")
	}

	pagenum = CalcPageNum(9, 4)
	if pagenum != 3 {
		t.Error("CalcPageNum error")
	}
}
