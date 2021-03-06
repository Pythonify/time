package stdTime

import (
	"strings"
	"time"
	"fmt"
)

const (
	StandardDatetime = "2006-01-02 15:04:05"
	StandardDate     = "2006-01-02"
)

var (
	nilTime = (time.Time{}).UnixNano()
)

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(StandardDatetime, s)
	return
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	if t.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format(StandardDatetime))), nil
}

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(StandardDate, s)
	return
}

func (t *Date) MarshalJSON() ([]byte, error) {
	if t.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format(StandardDate))), nil
}

func NewDateTime(time *time.Time) *DateTime {
	return &DateTime{*time}
}

func NewDate(time *time.Time) *Date {
	return &Date{*time}
}