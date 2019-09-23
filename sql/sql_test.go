package sql

import "testing"

func TestGetBulkInsertSqlStr(t *testing.T) {
	var columns []string
	var values [][]interface{}
	columns = append(columns, "`user_name`")
	columns = append(columns, "`user_age`")
	columns = append(columns, "`user_sex`")

	var oneColumn []interface{}
	oneColumn = append(oneColumn, "trump")
	oneColumn = append(oneColumn, 18.1)
	oneColumn = append(oneColumn, 0)
	var twoColumn []interface{}
	twoColumn = append(twoColumn, "trump2")
	twoColumn = append(twoColumn, 18.2)
	twoColumn = append(twoColumn, 1)
	var threeColumn []interface{}
	threeColumn = append(threeColumn, "trump3")
	threeColumn = append(threeColumn, 18.3)
	threeColumn = append(threeColumn, 2)

	values = append(values, oneColumn)
	values = append(values, twoColumn)
	values = append(values, threeColumn)

	type args struct {
		table   string
		columns []string
		values  [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{
			table: "user",
			columns: columns,
			values: values,
		}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBulkInsertSqlStr(tt.args.table, tt.args.columns, tt.args.values...); got == tt.want {
				t.Errorf("GetBulkInsertSqlStr() got empty")
			}
		})
	}
}