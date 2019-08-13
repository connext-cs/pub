package utils

import (
	"fmt"
	"testing"
)

func Test_ObjectToJsonStr(t *testing.T)  {
	type testStruct struct {
		Name string
		Age int
		Like []string
	}

	var testStructData = testStruct{
		Name:"hel",
		Age:15,
		Like:[]string{"a","b"},
	}

	ch1:=ObjectToJsonStr(testStructData)
	fmt.Println(ch1)
}
