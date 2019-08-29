package jsontime

import (
	"fmt"
	"time"
	"database/sql/driver"
)

// JSONTime format json time field by myself
type JsonTime struct {
    time.Time
}

const TimeFormat = "2006-01-02 15:04:05"


func (t JsonTime) MarshalJSON() ([]byte, error) {
    formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
    return []byte(formatted), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
    now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
    *t = JsonTime{Time: now}
    return err
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
    var zeroTime time.Time
    if t.Time.UnixNano() == zeroTime.UnixNano() {
        return nil, nil
    }
    return t.Time, nil
}

// Scan valueof time.Time
func (t *JsonTime) Scan(v interface{}) error {
    value, ok := v.(time.Time)
    if ok {
        *t = JsonTime{Time: value}
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}