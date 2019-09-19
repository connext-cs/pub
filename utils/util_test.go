package utils

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func Test_copy(t *testing.T)  {

	type Ast struct {
		A string
		B string
	}

	type Bst struct {
		A string
		B string
		C string
	}

	aa := Ast{
		"aa",
		"bb",
	}

	bb := Bst{}

	//StructCopy(bb, aa)

	logs.Info(bb)

}