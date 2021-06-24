package models

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime struct
type CustomTime struct {
	time.Time
}

// ctLayout custom layout.
const ctLayout = "2006-01-02T15:04:05.000Z0700"

var (
	nilTime = (time.Time{}).UnixNano()
)

// UnmarshalJSON json CustomTime correctly.
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON json CustomTime correctly.
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
