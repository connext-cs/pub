package pagination

import (
	"testing"
	"gotest.tools/assert"
)

func Test_SQL_Limit(t *testing.T) {
	limitStr := GetSqlLimitOption(1, 5, 12)
	assert.Equal(t, limitStr , " LIMIT 0,5 ")
	limitStr = GetSqlLimitOption(2, 5, 12)
	assert.Equal(t, limitStr , " LIMIT 5,5 ")
	limitStr = GetSqlLimitOption(3, 5, 12)
	assert.Equal(t, limitStr , " LIMIT 10,5 ")
	limitStr = GetSqlLimitOption(4, 5, 12)
	assert.Equal(t, limitStr , " LIMIT 10,5 ")
}

