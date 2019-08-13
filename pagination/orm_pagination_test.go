package pagination

import (
	"testing"
	"gotest.tools/assert"
)

func Test_ORMLimit(t *testing.T) {
	// positive case
	limit, offset, err := GetORMLimitOffset(1, 5, 12)
	assert.Assert(t, err == nil)
	assert.Equal(t, limit, 5)
	assert.Equal(t, offset, 0)

	// nagitive case, pagesize overload
	limit, offset, err = GetORMLimitOffset(4, 5, 12)
	assert.Assert(t, err == nil)
	assert.Equal(t, limit, 5)
	assert.Equal(t, offset, 10)
}

func Test_GenerageORMPagination(t *testing.T) {
	// positive case
	paging := GenerageORMPagination(5, 12)
	assert.Assert(t, paging != nil)
	assert.Equal(t, paging.PageSize, 5)
	assert.Equal(t, paging.Page, 3)
	assert.Equal(t, paging.Total, 12)

	// nagitive case
	paging = GenerageORMPagination(5, 2)
	assert.Assert(t, paging != nil)
	assert.Equal(t, paging.PageSize, 5)
	assert.Equal(t, paging.Page, 1)
	assert.Equal(t, paging.Total, 2)
}

