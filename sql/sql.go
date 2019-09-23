package sql

import "fmt"
import "strings"


func GetBulkInsertSqlStr(table string, columns []string, values ...[]interface{}) string {
	columnStr := strings.Join(columns, ",")
	sqlStr := fmt.Sprintf("INSERT INTO `%s` (%s) values ", table, columnStr)

	columnCnt := len(values) - 1
	for i := 0; i <= columnCnt; i++ {
		oneColumn := ""
		oneColumn += "("
		cnt := len(values[i]) - 1
		for j := 0; j <= cnt; j++ {
			val := values[i][j]
			switch val.(type) {
			case string:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%s',", val.(string))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%s'),", val.(string))
					} else {
						oneColumn += fmt.Sprintf("'%s');", val.(string))
					}
				}
				break
			case int:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(int))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(int))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(int))
					}
				}
				break
			case int32:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(int32))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(int32))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(int32))
					}
				}
				break
			case int16:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(int16))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(int16))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(int16))
					}
				}
				break
			case int8:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(int8))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(int8))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(int8))
					}
				}
				break
			case int64:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(int64))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(int64))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(int64))
					}
				}
				break
			case float32:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%f',", val.(float32))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%f'),", val.(float32))
					} else {
						oneColumn += fmt.Sprintf("'%f');", val.(float32))
					}
				}
				break
			case float64:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%f',", val.(float64))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%f'),", val.(float64))
					} else {
						oneColumn += fmt.Sprintf("'%f');", val.(float64))
					}
				}
				break
			case uint, uint16, uint32, uint64:
				if j < cnt {
					oneColumn += fmt.Sprintf("'%d',", val.(uint))
				} else {
					if i < columnCnt {
						oneColumn += fmt.Sprintf("'%d'),", val.(uint))
					} else {
						oneColumn += fmt.Sprintf("'%d');", val.(uint))
					}
				}
				break
			}

		}
		sqlStr += oneColumn
	}
	fmt.Print("insert sql is :\n", sqlStr)
	return sqlStr
}
