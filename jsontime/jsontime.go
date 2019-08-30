package jsontime

import (
	"fmt"
	"time"
)


const TimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"


type JsonDate time.Time
type JsonTime time.Time

func (this JsonDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(DateFormat))
	return []byte(stamp), nil
}
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(TimeFormat))
	return []byte(stamp), nil
}