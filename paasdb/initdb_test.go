package paasdb

import (
	"testing"
)

func Test_InitDB(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Error("InitDB,err : ", err.Error())
	}
}
